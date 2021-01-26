package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/keypair"
)

func TestCreateAsset(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		k, _ := keypair.Random()
		op := CreateAssetOp{
			AssetSigner:       k.Address(),
			MaxIssuanceAmount: 10,
			PreIssuanceAmount: 10,
			Policies:          0,
			Code:              "ETH",
			Details: AssetDetails{
				ExternalSystemType: 10,
				Name:               "Ethereum",
			},
		}

		got, err := op.XDR()
		request := got.Body.ManageAssetOp.Request.CreateAssetCreationRequest
		assert.NoError(t, err)
		assert.Equal(t, op.AssetSigner, request.CreateAsset.PreissuedAssetSigner.Address())
		assert.EqualValues(t, op.Policies, request.CreateAsset.Policies)
		assert.EqualValues(t, op.MaxIssuanceAmount, request.CreateAsset.MaxIssuanceAmount)
		assert.EqualValues(t, op.PreIssuanceAmount, request.CreateAsset.InitialPreissuedAmount)
		assert.EqualValues(t, op.Code, request.CreateAsset.Code)

	})

	t.Run("invalid issuer address", func(t *testing.T) {
		op := CreateAssetOp{
			AssetSigner:       "wrong key",
			MaxIssuanceAmount: 10,
			PreIssuanceAmount: 10,
			Policies:          0,
			Code:              "ETH",
			Details: AssetDetails{
				ExternalSystemType: 10,
				Name:               "Ethereum",
			},
		}
		got, err := op.XDR()
		assert.Error(t, err)
		assert.Nil(t, got)
	})
}
func TestUpdateAsset(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := UpdateAsset{
			Code:     "ETH",
			Policies: 0,
			Details: AssetDetails{
				ExternalSystemType: 10,
				Name:               "Ethereum",
			},
		}

		got, err := op.XDR()
		if assert.NoError(t, err) {
			assert.EqualValues(t, op.Policies, got.Body.ManageAssetOp.Request.CreateAssetUpdateRequest.UpdateAsset.Policies)
			assert.EqualValues(t, op.Code, got.Body.ManageAssetOp.Request.CreateAssetUpdateRequest.UpdateAsset.Code)
		}
	})
}
