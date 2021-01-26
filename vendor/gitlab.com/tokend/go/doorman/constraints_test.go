package doorman

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/resources"
	"gitlab.com/tokend/go/signcontrol"
)

type YesRoleConstraint struct {
	roleID uint64
}

func (c YesRoleConstraint) Check(signer resources.Signer) bool {
	return c.roleID == signer.Role
}

func NewYesRoleConstraint(roleID uint64) *YesRoleConstraint {
	return &YesRoleConstraint{
		roleID: roleID,
	}
}

func TestSignerOf(t *testing.T) {
	kp := keypair.MustParse("SBRU6ADAQTY7I4WRZ6KMGM64C3DTUK6PIYAWSV2UXC5AE3GPNWGYUIXM")
	r, _ := http.NewRequest("GET", "/", nil)
	require.NoError(t, signcontrol.SignRequest(r, kp))

	var q AccountQ
	doorman := New(false, &q)
	t.Run("master not a signer", func(t *testing.T) {
		q.On("Signers", kp.Address()).Return([]resources.Signer{}, nil).Once()
		defer q.AssertExpectations(t)

		err := SignerOf(kp.Address())(r, doorman)
		assert.Equal(t, ErrNotAllowed, err)
	})

	t.Run("magic_kyc_recovery_role", func(t *testing.T) {
		t.Run("signer with needed role exists", func(t *testing.T) {
			q.On("Signers", kp.Address()).Return(
				[]resources.Signer{
					{
						AccountID: "GDNIIABPH5YPL2ROG5EODYX2EOZSUPWUAI3QGQKAK2EKH3QXLQIFBWN5", // signer of the request
						Role:      8,
						Weight:    1000,
					},
				},
				nil,
			).Once()
			defer q.AssertExpectations(t)

			err := SignerOf(kp.Address(), NewYesRoleConstraint(8))(r, doorman)

			assert.Equal(t, nil, err)
		})

		t.Run("signer with needed role does not exist", func(t *testing.T) {
			q.On("Signers", kp.Address()).Return(
				[]resources.Signer{
					{
						Role: 9,
					},
				},
				nil,
			).Once()
			defer q.AssertExpectations(t)

			err := SignerOf(kp.Address(), NewYesRoleConstraint(8))(r, doorman)
			assert.Equal(t, ErrNotAllowed, err)
		})
	})

	t.Run("default role constraint", func(t *testing.T) {
		t.Run("with configured default (role allowed)", func(t *testing.T) {
			drm := NewWithOpts(false, &q, SignerOfOpts{
				Constraints: []SignerOfExt{
					&RoleConstraint{
						RoleID: 8,
					},
				},
			})

			q.On("Signers", kp.Address()).Return(
				[]resources.Signer{
					{
						AccountID: "GDNIIABPH5YPL2ROG5EODYX2EOZSUPWUAI3QGQKAK2EKH3QXLQIFBWN5",
						Weight:    1000,
						Role:      8,
					},
				},
				nil,
			).Once()

			err := SignerOf(kp.Address())(r, drm)
			assert.Equal(t, nil, err)
		})

		t.Run("with configured default (role not allowed)", func(t *testing.T) {
			drm := NewWithOpts(false, &q, SignerOfOpts{
				Constraints: []SignerOfExt{
					&RoleConstraint{
						RoleID: 8,
					},
				},
			})

			q.On("Signers", kp.Address()).Return(
				[]resources.Signer{
					{
						AccountID: "GDNIIABPH5YPL2ROG5EODYX2EOZSUPWUAI3QGQKAK2EKH3QXLQIFBWN5",
						Weight:    1000,
						Role:      9,
					},
				},
				nil,
			).Once()

			err := SignerOf(kp.Address())(r, drm)
			assert.Equal(t, ErrNotAllowed, err)
		})

		t.Run("without_configured_default", func(t *testing.T) {
			q.On("Signers", kp.Address()).Return(
				[]resources.Signer{
					{
						AccountID: "GDNIIABPH5YPL2ROG5EODYX2EOZSUPWUAI3QGQKAK2EKH3QXLQIFBWN5",
						Weight:    1000,
						Role:      8,
					},
				},
				nil,
			).Once()

			err := SignerOf(kp.Address())(r, doorman)
			assert.Equal(t, nil, err)
		})
	})
}
