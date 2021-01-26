package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCancelAtomicSwapBidOp_XDR(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := CancelAtomicSwapBidOp{
			BidID: 123,
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.MustCancelASwapBidOp()
		assert.EqualValues(t, op.BidID, body.BidId)
	})

	t.Run("bid id invalid", func(t *testing.T) {
		op := CancelAtomicSwapBidOp{
			BidID: 0,
		}
		assert.Error(t, op.Validate())
	})
}
