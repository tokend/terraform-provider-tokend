package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/xdr"
)

func TestCreateKYCRequestOp_XDR(t *testing.T) {
	kp, _ := keypair.Random()
	var allTasks uint32 = 3
	t.Run("valid", func(t *testing.T) {
		op := CreateUpdateKYCRequestOp{
			RequestID:          0,
			AccountToUpdateKYC: kp.Address(),
			AccountTypeToSet:   xdr.AccountTypeGeneral,
			KYCLevelToSet:      1,
			KYCData:            "Some KYC data",
			AllTasks:           &allTasks,
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateUpdateKycRequestOp
		assert.EqualValues(t, op.RequestID, body.RequestId)
		assert.EqualValues(t, op.AccountTypeToSet, body.UpdateKycRequestData.AccountTypeToSet)
		assert.EqualValues(t, op.KYCData, body.UpdateKycRequestData.KycData)
		assert.EqualValues(t, op.AccountToUpdateKYC, body.UpdateKycRequestData.AccountToUpdateKyc.Address())
		assert.EqualValues(t, op.KYCLevelToSet, body.UpdateKycRequestData.KycLevelToSet)
		assert.EqualValues(t, op.AllTasks, body.AllTasks)
	})

	t.Run("missing account type", func(t *testing.T) {
		op := CreateUpdateKYCRequestOp{
			RequestID:          0,
			KYCData:            "Some KYC data",
			AccountToUpdateKYC: kp.Address(),
			KYCLevelToSet:      1,
			AllTasks:           nil,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing KYC data", func(t *testing.T) {
		op := CreateUpdateKYCRequestOp{
			RequestID:          0,
			AccountTypeToSet:   xdr.AccountTypeGeneral,
			AccountToUpdateKYC: kp.Address(),
			KYCLevelToSet:      1,
			AllTasks:           nil,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing updated account", func(t *testing.T) {
		op := CreateUpdateKYCRequestOp{
			RequestID:        0,
			AccountTypeToSet: xdr.AccountTypeGeneral,
			KYCData:          "Some KYC data",
			KYCLevelToSet:    1,
			AllTasks:         nil,
		}
		assert.Error(t, op.Validate())
	})
}
