package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIssuanceRequestOp_XDR(t *testing.T) {
	balance := "BCGLIDURX7ENQXETLGQXB4RZHX3NE27AZU7SWVGBNBELJ56V36RU2DDC"

	t.Run("valid", func(t *testing.T) {
		op := CreateIssuanceRequestOp{
			Reference: "foobar",
			Receiver:  balance,
			Asset:     "YOBA",
			Amount:    42,
			Details:   "{}",
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateIssuanceRequestOp
		assert.EqualValues(t, op.Reference, body.Reference)
		assert.EqualValues(t, op.Receiver, body.Request.Receiver.AsString())
		assert.EqualValues(t, op.Asset, body.Request.Asset)
		assert.EqualValues(t, op.Amount, body.Request.Amount)
		assert.EqualValues(t, op.Details, body.Request.ExternalDetails)
		assert.Nil(t, op.AllTasks, nil)
	})

	t.Run("valid with tasks", func(t *testing.T) {
		var allTasks uint32 = 8
		op := CreateIssuanceRequestOp{
			Reference: "foobar",
			Receiver:  balance,
			Asset:     "YOBA",
			Amount:    42,
			Details:   "{}",
			AllTasks:  &allTasks,
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateIssuanceRequestOp
		assert.EqualValues(t, op.Reference, body.Reference)
		assert.EqualValues(t, op.Receiver, body.Request.Receiver.AsString())
		assert.EqualValues(t, op.Asset, body.Request.Asset)
		assert.EqualValues(t, op.Amount, body.Request.Amount)
		assert.EqualValues(t, op.Details, body.Request.ExternalDetails)
		assert.EqualValues(t, *op.AllTasks, uint32(*body.AllTasks))
	})

	cases := []struct {
		name string
		op   CreateIssuanceRequestOp
	}{
		{
			"missing reference",
			CreateIssuanceRequestOp{
				Receiver: balance,
				Asset:    "YOBA",
				Amount:   42,
				Details:  "{}",
			},
		},
		{
			"reference too long",
			CreateIssuanceRequestOp{
				Reference: "foobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobar",
				Receiver:  balance,
				Asset:     "YOBA",
				Amount:    42,
				Details:   "{}",
			},
		},
		{
			"missing receiver",
			CreateIssuanceRequestOp{
				Reference: "foobar",
				Asset:     "YOBA",
				Amount:    42,
				Details:   "{}",
			},
		},
		{
			"missing asset",
			CreateIssuanceRequestOp{
				Reference: "foobar",
				Receiver:  balance,
				Amount:    42,
				Details:   "{}",
			},
		},
		{
			"asset too long",
			CreateIssuanceRequestOp{
				Reference: "foobar",
				Receiver:  balance,
				Asset:     "YOBAYOBAYOBAYOBAYOBA",
				Amount:    42,
				Details:   "{}",
			},
		},
		{
			"missing amount",
			CreateIssuanceRequestOp{
				Reference: "foobar",
				Receiver:  balance,
				Asset:     "YOBA",
				Details:   "{}",
			},
		},
		{
			"missing details",
			CreateIssuanceRequestOp{
				Reference: "foobar",
				Receiver:  balance,
				Asset:     "YOBA",
				Amount:    42,
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Error(t, tc.op.Validate())
		})
	}
}
