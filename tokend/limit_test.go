package tokend

import (
	"testing"

	"gitlab.com/tokend/go/xdrbuild"

	"gitlab.com/tokend/keypair"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/tokend/terraform-provider-tokend/tokend/data/mocks"
)

var limitSchema = map[string]*schema.Schema{

	"role": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"account_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"stats_type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"asset_code": {
		Type:     schema.TypeString,
		Required: true,
	},
	"convert": {
		Type:     schema.TypeBool,
		Required: true,
	},
	"daily_out": {
		Type:     schema.TypeInt,
		Required: true,
	},
	"weekly_out": {
		Type:     schema.TypeInt,
		Required: true,
	},
	"monthly_out": {
		Type:     schema.TypeInt,
		Required: true,
	},
	"annual_out": {
		Type:     schema.TypeInt,
		Required: true,
	},
}

func TestCreateLimit_XDR(t *testing.T) {
	sign, err := keypair.Random()
	if err != nil {

	}

	connector := mocks.Connector{}
	m := Meta{Connector: &connector, Signer: sign, Source: keypair.MustParseAddress("GBA4EX43M25UPV4WIE6RRMQOFTWXZZRIPFAI5VPY6Z2ZVVXVWZ6NEOOB")}
	builder := xdrbuild.NewBuilder("g", 1)
	m.Builder = *builder

	t.Run("valid", func(t *testing.T) {
		c := map[string]interface{}{
			"account_id":  "GBA4EX43M25UPV4WIE6RRMQOFTWXZZRIPFAI5VPY6Z2ZVVXVWZ6NEOOB",
			"stats_type":  "withdraw",
			"asset_code":  "BTC",
			"convert":     true,
			"daily_out":   3,
			"weekly_out":  12,
			"monthly_out": 35,
			"annual_out":  125,
		}

		d := schema.TestResourceDataRaw(t, limitSchema, c)

		err := resourceLimitsCreate(d, m)
		assert.NoError(t, err)

	})

	t.Run("invalid", func(t *testing.T) {
		c := map[string]interface{}{
			"account_id":  "GBA4EX43M25UPV4WIE6RRMQOFTWXZZRIPFAI5VPY6Z2ZVVXVWZ6NEOOB",
			"stats_type":  "withdraw",
			"asset_code":  "BTC",
			"convert":     true,
			"daily_out":   3,
			"weekly_out":  12,
			"monthly_out": 35,
			"annual_out":  125,
		}

		d := schema.TestResourceDataRaw(t, limitSchema, c)

		err := resourceLimitsCreate(d, m)
		assert.NoError(t, err)

	})
}
