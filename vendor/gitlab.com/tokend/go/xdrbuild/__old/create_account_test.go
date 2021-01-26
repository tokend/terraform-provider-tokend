package __old

import (
	"testing"

	"gitlab.com/tokend/go/xdr"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/keypair"
)

func TestCreateAccountOp_XDR(t *testing.T) {
	one, _ := keypair.Random()
	three, _ := keypair.Random()

	t.Run("valid w/o referrer", func(t *testing.T) {
		op := CreateAccountOp{
			Address:     one.Address(),
			RoleID:      1,
			SignersData: []xdr.UpdateSignerData{{}},
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateAccountOp
		assert.EqualValues(t, op.RoleID, body.RoleId)
		assert.EqualValues(t, op.Address, body.Destination.Address())
		assert.Nil(t, body.Referrer)
	})

	t.Run("valid with referrer", func(t *testing.T) {
		referrer := three.Address()
		op := CreateAccountOp{
			Address:     one.Address(),
			RoleID:      1,
			Referrer:    &referrer,
			SignersData: []xdr.UpdateSignerData{{}},
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateAccountOp
		assert.EqualValues(t, op.RoleID, body.RoleId)
		assert.EqualValues(t, op.Address, body.Destination.Address())
		assert.EqualValues(t, &referrer, op.Referrer)
	})

	t.Run("missing address", func(t *testing.T) {
		op := CreateAccountOp{
			RoleID:      1,
			SignersData: []xdr.UpdateSignerData{{}},
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing signers data", func(t *testing.T) {
		op := CreateAccountOp{
			Address: one.Address(),
			RoleID:  1,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing account type", func(t *testing.T) {
		op := CreateAccountOp{
			Address:     one.Address(),
			SignersData: []xdr.UpdateSignerData{{}},
		}
		assert.Error(t, op.Validate())
	})
}
