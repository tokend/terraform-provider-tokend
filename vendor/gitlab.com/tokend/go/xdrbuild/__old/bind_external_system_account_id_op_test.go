package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

func TestBindExternalSystemAccountIDOp_Validate(t *testing.T) {
	t.Run("external type is required", func(t *testing.T) {
		op := BindExternalSystemAccountIDOp{}
		assert.Error(t, op.Validate())
	})

	t.Run("external type should be positive", func(t *testing.T) {
		op := BindExternalSystemAccountIDOp{
			ExternalSystem: -1,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("valid should pass", func(t *testing.T) {
		op := BindExternalSystemAccountIDOp{
			ExternalSystem: 1,
		}
		assert.NoError(t, op.Validate())
	})
}

func TestBindExternalSystemAccountIDOp_XDR(t *testing.T) {
	op := BindExternalSystemAccountIDOp{
		ExternalSystem: 1,
	}
	got, err := op.XDR()
	assert.NoError(t, err)
	assert.Equal(t, xdr.OperationTypeBindExternalSystemAccountId, got.Body.Type)
	assert.NotNil(t, got.Body.BindExternalSystemAccountIdOp)
	assert.EqualValues(t, 1, got.Body.BindExternalSystemAccountIdOp.ExternalSystemType)
}
