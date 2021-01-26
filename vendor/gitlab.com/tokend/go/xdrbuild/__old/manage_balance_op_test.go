package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

func TestManageBalance(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := ManageBalanceOp{
			Action:      xdr.ManageBalanceActionCreate,
			Asset:       "USD",
			Destination: "GCHAGMRQ6IQEFJPSDNXL3NR2L4L7AIYDUNJVESSZEIA2A7VR4NVKOL2P",
		}

		got, err := op.XDR()
		assert.NoError(t, err)
		assert.EqualValues(t, op.Action, got.Body.ManageBalanceOp.Action)
		assert.EqualValues(t, op.Asset, got.Body.ManageBalanceOp.Asset)
		assert.NotNil(t, got.Body.ManageBalanceOp.Destination)
	})

	t.Run("not valid", func(t *testing.T) {
		op := ManageBalanceOp{
			Action:      xdr.ManageBalanceActionCreate,
			Asset:       "USD",
			Destination: "wrong",
		}

		err := op.Validate()
		assert.Error(t, err)
	})
}
