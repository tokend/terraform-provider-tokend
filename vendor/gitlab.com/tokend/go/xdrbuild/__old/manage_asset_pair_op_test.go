package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAssetPair(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := CreateAssetPair{
			Base:                    "ETH",
			Quote:                   "USD",
			PhysicalPrice:           1,
			PhysicalPriceCorrection: 1,
			Policies:                2,
			MaxPriceStep:            3,
		}

		got, err := op.XDR()
		assert.NoError(t, err)
		assert.EqualValues(t, op.Base, got.Body.ManageAssetPairOp.Base)
		assert.EqualValues(t, op.Quote, got.Body.ManageAssetPairOp.Quote)
		assert.EqualValues(t, op.PhysicalPrice, got.Body.ManageAssetPairOp.PhysicalPrice)
		assert.EqualValues(t, op.PhysicalPriceCorrection, got.Body.ManageAssetPairOp.PhysicalPriceCorrection)
		assert.EqualValues(t, op.Policies, got.Body.ManageAssetPairOp.Policies)
		assert.EqualValues(t, op.MaxPriceStep, got.Body.ManageAssetPairOp.MaxPriceStep)
	})

	t.Run("zero price", func(t *testing.T) {
		op := CreateAssetPair{
			Base:          "ETH",
			Quote:         "USD",
			PhysicalPrice: 0,
		}

		err := op.Validate()
		assert.Error(t, err)
	})
}

func TestUpdateAssetPairPrice(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := UpdateAssetPairPrice{
			Base:          "ETH",
			Quote:         "USD",
			PhysicalPrice: 1,
		}

		got, err := op.XDR()
		assert.NoError(t, err)
		assert.EqualValues(t, op.Base, got.Body.ManageAssetPairOp.Base)
		assert.EqualValues(t, op.Quote, got.Body.ManageAssetPairOp.Quote)
		assert.EqualValues(t, op.PhysicalPrice, got.Body.ManageAssetPairOp.PhysicalPrice)
	})
	t.Run("zero price", func(t *testing.T) {
		op := UpdateAssetPairPrice{
			Base:          "ETH",
			Quote:         "USD",
			PhysicalPrice: 0,
		}

		err := op.Validate()
		assert.Error(t, err)
	})
}

func TestUpdateAssetPairPolicies(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := UpdateAssetPairPolicies{
			Base:  "ETH",
			Quote: "USD",
			PhysicalPriceCorrection: 1,
			Policies:                2,
			MaxPriceStep:            3,
		}

		got, err := op.XDR()
		assert.NoError(t, err)
		assert.EqualValues(t, op.Base, got.Body.ManageAssetPairOp.Base)
		assert.EqualValues(t, op.Quote, got.Body.ManageAssetPairOp.Quote)
		assert.EqualValues(t, op.PhysicalPriceCorrection, got.Body.ManageAssetPairOp.PhysicalPriceCorrection)
		assert.EqualValues(t, op.Policies, got.Body.ManageAssetPairOp.Policies)
		assert.EqualValues(t, op.MaxPriceStep, got.Body.ManageAssetPairOp.MaxPriceStep)
	})
}
