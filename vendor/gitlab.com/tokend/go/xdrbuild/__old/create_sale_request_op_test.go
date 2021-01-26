package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

func TestCreateSaleRequest(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := CreateSaleRequestOp{
			RequestID:         0,
			BaseAsset:         "QTK",
			DefaultQuoteAsset: "USD",
			StartTime:         1542293558,
			EndTime:           1546387200,
			SoftCap:           250000000,
			HardCap:           450000000,
			Details:           "{ 'foo': 'bar' }",
			QuoteAssets: []xdr.SaleCreationRequestQuoteAsset{
				xdr.SaleCreationRequestQuoteAsset{
					Price:      1000000,
					QuoteAsset: "BTC",
				},
				xdr.SaleCreationRequestQuoteAsset{
					Price:      1000000,
					QuoteAsset: "ETH",
				},
			},
		}

		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateSaleCreationRequestOp
		assert.EqualValues(t, op.RequestID, body.RequestId)
		assert.EqualValues(t, op.BaseAsset, body.Request.BaseAsset)
		assert.EqualValues(t, op.DefaultQuoteAsset, body.Request.DefaultQuoteAsset)
		assert.EqualValues(t, op.StartTime, body.Request.StartTime)
		assert.EqualValues(t, op.EndTime, body.Request.EndTime)
		assert.EqualValues(t, op.SoftCap, body.Request.SoftCap)
		assert.EqualValues(t, op.HardCap, body.Request.HardCap)
		assert.EqualValues(t, op.Details, body.Request.CreatorDetails)
	})
}
