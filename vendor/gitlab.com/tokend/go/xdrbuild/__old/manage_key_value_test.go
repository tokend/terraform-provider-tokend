package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

func TestManageKeyValueOp_XDR(t *testing.T) {
	t.Run("valid with int32", func(t *testing.T) {
		v := uint32(6)
		op := ManageKeyValueOp{
			Key:    "Key",
			Uint32: &v,
		}
		assert.NoError(t, op.Validate())
		xdrOp, err := op.XDR()
		assert.NoError(t, err)
		assert.Equal(t, xdr.ManageKvActionPut, xdrOp.Body.MustManageKeyValueOp().Action.Action)
		assert.Equal(t, xdr.KeyValueEntryTypeUint32, xdrOp.Body.MustManageKeyValueOp().Action.Value.Type)
		assert.Equal(t, xdr.Uint32(v), xdrOp.Body.MustManageKeyValueOp().Action.Value.MustUi32Value())
	})
	t.Run("valid with int64", func(t *testing.T) {
		v := uint64(6)
		op := ManageKeyValueOp{
			Key:    "Key",
			Uint64: &v,
		}
		assert.NoError(t, op.Validate())
		xdrOp, err := op.XDR()
		assert.NoError(t, err)
		assert.Equal(t, xdr.ManageKvActionPut, xdrOp.Body.MustManageKeyValueOp().Action.Action)
		assert.Equal(t, xdr.KeyValueEntryTypeUint64, xdrOp.Body.MustManageKeyValueOp().Action.Value.Type)
		assert.Equal(t, xdr.Uint64(v), xdrOp.Body.MustManageKeyValueOp().Action.Value.MustUi64Value())
	})
	t.Run("valid with string", func(t *testing.T) {
		str := "TaskFaceValidation"
		op := ManageKeyValueOp{
			Key:    "Key",
			String: &str,
		}
		assert.NoError(t, op.Validate())
		xdrOp, err := op.XDR()
		assert.NoError(t, err)
		assert.Equal(t, xdr.ManageKvActionPut, xdrOp.Body.MustManageKeyValueOp().Action.Action)
		assert.Equal(t, xdr.KeyValueEntryTypeString, xdrOp.Body.MustManageKeyValueOp().Action.Value.Type)
		assert.Equal(t, str, xdrOp.Body.MustManageKeyValueOp().Action.Value.MustStringValue())
	})
	t.Run("valid remove", func(t *testing.T) {
		op := ManageKeyValueOp{
			Key: "Key",
		}
		assert.NoError(t, op.Validate())
		xdrOp, err := op.XDR()
		assert.NoError(t, err)
		assert.Equal(t, xdr.ManageKvActionRemove, xdrOp.Body.MustManageKeyValueOp().Action.Action)
	})
	t.Run("invalid struct", func(t *testing.T) {
		v := uint32(6)
		str := "TaskFaceValidation"
		op := ManageKeyValueOp{
			String: &str,
			Uint32: &v,
		}
		assert.Error(t, op.Validate())
	})
	t.Run("valid with empty string", func(t *testing.T) {
		str := ""
		op := ManageKeyValueOp{
			Key:    "Key",
			String: &str,
		}
		assert.NoError(t, op.Validate())
	})
	t.Run("valid with 0", func(t *testing.T) {
		v := uint32(0)
		op := ManageKeyValueOp{
			Key:    "Key",
			Uint32: &v,
		}
		assert.NoError(t, op.Validate())
	})
}
