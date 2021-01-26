package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateASwapRequestOp_XDR(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := CreateAtomicSwapRequestOp{
			BidID:      123,
			BaseAmount: 1234,
			QuoteAsset: "XRP",
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		request := got.Body.MustCreateASwapRequestOp().Request
		assert.EqualValues(t, op.BidID, request.BidId)
		assert.EqualValues(t, op.BaseAmount, request.BaseAmount)
		assert.EqualValues(t, op.QuoteAsset, request.QuoteAsset)
	})

	t.Run("bid id invalid", func(t *testing.T) {
		op := CreateAtomicSwapRequestOp{
			BidID:      0,
			BaseAmount: 1234,
			QuoteAsset: "XRP",
		}
		assert.Error(t, op.Validate())
	})

	t.Run("base amount invalid", func(t *testing.T) {
		op := CreateAtomicSwapRequestOp{
			BidID:      123,
			BaseAmount: 0,
			QuoteAsset: "XRP",
		}
		assert.Error(t, op.Validate())
	})

	t.Run("invalid quote asset", func(t *testing.T) {
		op := CreateAtomicSwapRequestOp{
			BidID:      123,
			BaseAmount: 1234,
			QuoteAsset: "",
		}
		assert.Error(t, op.Validate())
	})
}
