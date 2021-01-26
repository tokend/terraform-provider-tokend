package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateASwapBidCreationRequestOp_XDR(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := CreateAtomicSwapBidCreationRequestOp{
			BaseBalance: "BDNYSW4QCZPPP4IMT67HJHFKRYLKD2PEAO3VTNQ7MVVN2SXUHCNWTFH5",
			BaseAmount:  1234,
			Details:     "Some details must be json struct",
			QuoteAssets: []QuoteAsset{
				{Price: 10, Asset: "XRP"},
				{Price: 20, Asset: "LTC"},
			},
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		request := got.Body.MustCreateASwapBidCreationRequestOp().Request
		assert.EqualValues(t, op.BaseBalance, request.BaseBalance.AsString())
		assert.EqualValues(t, op.BaseAmount, request.Amount)
		assert.EqualValues(t, op.Details, request.CreatorDetails)
		for i := range op.QuoteAssets {
			assert.EqualValues(t, op.QuoteAssets[i].Price, request.QuoteAssets[i].Price)
			assert.EqualValues(t, op.QuoteAssets[i].Asset, request.QuoteAssets[i].QuoteAsset)
		}
	})

	t.Run("base balance invalid", func(t *testing.T) {
		op := CreateAtomicSwapBidCreationRequestOp{
			BaseBalance: "",
			BaseAmount:  1234,
			Details:     "Some details must be json struct",
			QuoteAssets: []QuoteAsset{
				{Price: 10, Asset: "XRP"},
				{Price: 20, Asset: "LTC"},
			},
		}
		assert.Error(t, op.Validate())

		op = CreateAtomicSwapBidCreationRequestOp{
			BaseBalance: "INVALIDBALANCEID",
			BaseAmount:  1234,
			Details:     "Some details must be json struct",
			QuoteAssets: []QuoteAsset{
				{Price: 10, Asset: "XRP"},
				{Price: 20, Asset: "LTC"},
			},
		}
		assert.NoError(t, op.Validate())
		_, err := op.XDR()
		assert.Error(t, err)
	})

	t.Run("base amount invalid", func(t *testing.T) {
		op := CreateAtomicSwapBidCreationRequestOp{
			BaseBalance: "BDNYSW4QCZPPP4IMT67HJHFKRYLKD2PEAO3VTNQ7MVVN2SXUHCNWTFH5",
			BaseAmount:  0,
			Details:     "Some details must be json struct",
			QuoteAssets: []QuoteAsset{
				{Price: 10, Asset: "XRP"},
				{Price: 20, Asset: "LTC"},
			},
		}
		assert.Error(t, op.Validate())
	})

	t.Run("invalid details", func(t *testing.T) {
		op := CreateAtomicSwapBidCreationRequestOp{
			BaseBalance: "BDNYSW4QCZPPP4IMT67HJHFKRYLKD2PEAO3VTNQ7MVVN2SXUHCNWTFH5",
			BaseAmount:  1234,
			Details:     "",
			QuoteAssets: []QuoteAsset{
				{Price: 10, Asset: "XRP"},
				{Price: 20, Asset: "LTC"},
			},
		}
		assert.Error(t, op.Validate())
	})

	t.Run("invalid quote assets", func(t *testing.T) {
		op := CreateAtomicSwapBidCreationRequestOp{
			BaseBalance: "BDNYSW4QCZPPP4IMT67HJHFKRYLKD2PEAO3VTNQ7MVVN2SXUHCNWTFH5",
			BaseAmount:  1234,
			Details:     "Some details must be json struct",
			QuoteAssets: []QuoteAsset{
				{Price: 0, Asset: "XRP"},
				{Price: 20, Asset: "LTC"},
			},
		}
		assert.Error(t, op.Validate())

		op = CreateAtomicSwapBidCreationRequestOp{
			BaseBalance: "BDNYSW4QCZPPP4IMT67HJHFKRYLKD2PEAO3VTNQ7MVVN2SXUHCNWTFH5",
			BaseAmount:  1234,
			Details:     "Some details must be json struct",
			QuoteAssets: []QuoteAsset{
				{Price: 10, Asset: ""},
				{Price: 20, Asset: "LTC"},
			},
		}
		assert.Error(t, op.Validate())

		op = CreateAtomicSwapBidCreationRequestOp{
			BaseBalance: "BDNYSW4QCZPPP4IMT67HJHFKRYLKD2PEAO3VTNQ7MVVN2SXUHCNWTFH5",
			BaseAmount:  1234,
			Details:     "Some details must be json struct",
			QuoteAssets: []QuoteAsset{},
		}
		assert.Error(t, op.Validate())
	})
}
