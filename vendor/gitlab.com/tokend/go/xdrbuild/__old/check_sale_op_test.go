package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSaleOp_XDR(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := CheckSaleOp{
			SaleID: 10,
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		assert.EqualValues(t, op.SaleID, got.Body.CheckSaleStateOp.SaleId)
	})

	t.Run("missing sale id", func(t *testing.T) {
		op := CheckSaleOp{}
		assert.Error(t, op.Validate())
	})
}
