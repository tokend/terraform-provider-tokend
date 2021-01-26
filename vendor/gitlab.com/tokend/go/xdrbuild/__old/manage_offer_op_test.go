package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

func TestManageOfferOp_XDR(t *testing.T) {
	t.Run("valid CreateOffer", func(t *testing.T) {
		op := CreateOffer("BCF36NGODB272WKER4UMBGNTJTGXCR25E3TPAPXFIRPDX6GC4S3ILN2A", "BDZ62WB4FRQKSJ2SJRKQUMEJFPMMDDT2ZYL4JJJATVAIRCBDDGIJCIOM",
			true, 17, 3, 0)
		assert.NoError(t, op.Validate())

		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.ManageOfferOp
		assert.EqualValues(t, op.BaseBalance, body.BaseBalance.AsString())
		assert.EqualValues(t, op.QuoteBalance, body.QuoteBalance.AsString())

		assert.EqualValues(t, op.Amount, int64(body.Amount))
		assert.EqualValues(t, op.Price, int64(body.Price))
		assert.EqualValues(t, op.Fee, int64(body.Fee))

		assert.EqualValues(t, op.OfferID, uint64(0))
		assert.EqualValues(t, uint64(0), uint64(body.OrderBookId))
	})

	t.Run("valid DeleteOffer", func(t *testing.T) {
		op := DeleteOffer(uint64(18))
		assert.NoError(t, op.Validate())

		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.ManageOfferOp

		var baseBalance, quoteBalance xdr.BalanceId
		baseBalance.SetString("BBQJLW43UPKJSS67OXGAHJKZ4RM7JMQ6P7FCJEOMBC5GMYT3XQARZI54")
		quoteBalance.SetString("BBQJLW43UPKJSS67OXGAHJKZ4RM7JMQ6P7FCJEOMBC5GMYT3XQARZI54")

		assert.EqualValues(t, body.BaseBalance, baseBalance)
		assert.EqualValues(t, body.QuoteBalance, quoteBalance)

		assert.EqualValues(t, int64(0), int64(body.Amount))
		assert.EqualValues(t, int64(0), int64(body.Price))
		assert.EqualValues(t, int64(0), int64(body.Fee))

		assert.EqualValues(t, uint64(18), uint64(body.OfferId))
		assert.EqualValues(t, uint64(0), uint64(body.OrderBookId))
	})

	t.Run("missing BaseBalance", func(t *testing.T) {
		op := CreateOffer("", "BDZ62WB4FRQKSJ2SJRKQUMEJFPMMDDT2ZYL4JJJATVAIRCBDDGIJCIOM",
			true, 17, 3, 0)
		assert.Error(t, op.Validate())
	})

	t.Run("missing QuoteBalance", func(t *testing.T) {
		op := CreateOffer("BCF36NGODB272WKER4UMBGNTJTGXCR25E3TPAPXFIRPDX6GC4S3ILN2A", "",
			true, 17, 3, 0)
		assert.Error(t, op.Validate())
	})

	t.Run("zero Amount", func(t *testing.T) {
		op := CreateOffer("BCF36NGODB272WKER4UMBGNTJTGXCR25E3TPAPXFIRPDX6GC4S3ILN2A", "BDZ62WB4FRQKSJ2SJRKQUMEJFPMMDDT2ZYL4JJJATVAIRCBDDGIJCIOM",
			true, 0, 3, 0)
		assert.Error(t, op.Validate())
	})

	t.Run("zero Price", func(t *testing.T) {
		op := CreateOffer("BCF36NGODB272WKER4UMBGNTJTGXCR25E3TPAPXFIRPDX6GC4S3ILN2A", "BDZ62WB4FRQKSJ2SJRKQUMEJFPMMDDT2ZYL4JJJATVAIRCBDDGIJCIOM",
			true, 17, 0, 0)
		assert.Error(t, op.Validate())
	})

	t.Run("Zero OfferID", func(t *testing.T) {
		op := DeleteOffer(0)
		assert.Error(t, op.Validate())
	})
}
