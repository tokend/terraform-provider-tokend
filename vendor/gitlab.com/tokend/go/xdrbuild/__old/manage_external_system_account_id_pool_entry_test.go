package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

func TestCreateExternalPoolEntry(t *testing.T) {
	t.Run("valid create entry", func(t *testing.T) {
		op := CreateExternalPoolEntry(1, "foobar", 2)
		got, err := op.XDR()
		assert.NoError(t, err)
		assert.EqualValues(t, xdr.OperationTypeManageExternalSystemAccountIdPoolEntry, got.Body.Type)
		body := got.Body.ManageExternalSystemAccountIdPoolEntryOp
		assert.NotNil(t, body)
		assert.EqualValues(t, xdr.ManageExternalSystemAccountIdPoolEntryActionCreate, body.ActionInput.Action)
		input := body.ActionInput.CreateExternalSystemAccountIdPoolEntryActionInput
		assert.NotNil(t, input)
		assert.EqualValues(t, 1, input.ExternalSystemType)
		assert.EqualValues(t, "foobar", input.Data)
		assert.EqualValues(t, 2, input.Parent)
	})
}
