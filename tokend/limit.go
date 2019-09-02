package tokend

import (
	"context"
	"fmt"

	"gitlab.com/tokend/go/xdrbuild"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers"
	"gitlab.com/tokend/go/xdr"
)

func resourceLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceLimitsCreate,
		Read:   resourceLimitsRead,
		Update: resourceLimitsUpdate,
		Delete: resourceLimitsDelete,

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceLimitsCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	rawAccountID := d.Get("account_id")
	accountID, err := helpers.AccountIDFromRaw(rawAccountID)
	if err != nil {
		return errors.Wrap(err, "failed to get account id")
	}

	accountRole := xdr.Uint64(d.Get("role").(int))

	rawType := d.Get("stats_type")
	typesCode, err := helpers.StatsOpTypeFromRaw(rawType)
	if err != nil {
		return errors.Wrap(err, "failed to get type")
	}

	dailyOut := d.Get("daily_out").(int)

	weeklyOut := d.Get("weekly_out").(int)

	monthlyOut := d.Get("monthly_out").(int)

	annualOut := d.Get("annual_out").(int)

	if !helpers.ValidateLimits(dailyOut, weeklyOut, monthlyOut, annualOut) {
		return errors.New("failed to set limits - incorrect values")
	}

	assetCode := d.Get("asset_code").(string)

	convertNeed := d.Get("convert").(bool)

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateLimit{
		Action:     xdr.ManageLimitsActionCreate,
		Role:       &accountRole,
		Id:         accountID,
		Type:       xdr.StatsOpType(typesCode),
		Code:       xdr.AssetCode(assetCode),
		Convert:    convertNeed,
		DailyOut:   xdr.Uint64(dailyOut),
		WeeklyOut:  xdr.Uint64(weeklyOut),
		MonthlyOut: xdr.Uint64(monthlyOut),
		AnnualOut:  xdr.Uint64(annualOut),
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	result := m.Horizon.Submitter().Submit(context.TODO(), env)
	if result.Err != nil {
		return errors.Wrapf(result.Err, "failed to submit tx: %s %q", result.TXCode, result.OpCodes)
	}
	return nil
}

func resourceLimitsRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceLimitsUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceLimitsDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(Meta)
	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast limit id")
	}
	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.RemoveLimit{
		ID: id,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	result := m.Horizon.Submitter().Submit(context.TODO(), env)
	if result.Err != nil {
		return errors.Wrapf(result.Err, "failed to submit tx: %s %q", result.TXCode, result.OpCodes)
	}
	var txResult xdr.TransactionResult
	if err := xdr.SafeUnmarshalBase64(result.ResultXDR, &txResult); err != nil {
		return errors.Wrap(err, "failed to decode result")
	}
	txCodes := *(txResult.Result.Results)
	reply := txCodes[0].Tr.ManageLimitsResult.Success.Details.Id
	d.SetId(fmt.Sprintf("%d", reply))
	return nil
}
