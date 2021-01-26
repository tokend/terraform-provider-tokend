package internal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

func TestOperationFunc_XDR(t *testing.T) {
	expectedErr := errors.New("sup")

	f := func() (*xdr.Operation, error) {
		return nil, expectedErr
	}

	builder := OperationFunc(f)
	assert.Implements(t, (*Operation)(nil), builder, "OperationFunc expected to implement Operation interface")
	_, got := f()
	assert.Equal(t, expectedErr, got)
}
