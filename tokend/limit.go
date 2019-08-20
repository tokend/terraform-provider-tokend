package tokend

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers"
	"gitlab.com/tokend/go/xdr"
)

func resourceLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceLimitsCreate,
		Delete: resourceLimitsDelete,
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  1,
			},
			"stats_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"convert": {
				Type:     schema.TypeBool,
				Required: true,
				//Default:  false,
			},
			"daily_out": {
				Type:     schema.TypeInt,
				Required: true,
				//Default: 1,
			},
			"weekly_out": {
				Type:     schema.TypeInt,
				Required: true,
				//Default: 1,
			},
			"monthly_out": {
				Type:     schema.TypeInt,
				Required: true,
				//Default: 1,
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

	rawAccountId := d.Get("account_id")
	accountId, err := helpers.AccountIDFromRaw(rawAccountId)
	if err != nil {
		return errors.Wrap(err, "failed to get account id")
	}

	rawAccountRole := d.Get("role")

	rawStatsType := d.Get("stats_type")

	rawDailyOut := d.Get("daily_out")
	dailyOut, err := cast.ToUint64E(rawDailyOut)
	if err != nil {
		return errors.Wrap(err, "failed to cast daily out")
	}

	rawWeeklyOut := d.Get("weekly_out")
	weeklyOut, err := cast.ToUint64E(rawWeeklyOut)
	if err != nil {
		return errors.Wrap(err, "failed to cast weekly out")
	}

	rawMonthlyOut := d.Get("monthly_out")
	monthlyOut, err := cast.ToUint64E(rawMonthlyOut)
	if err != nil {
		return errors.Wrap(err, "failed to cast monthly out")
	}

	rawAnnualOut := d.Get("annual_out")
	annualOut, err := cast.ToUint64E(rawAnnualOut)
	if err != nil {
		return errors.Wrap(err, "failed to cast annual out")
	}

	boolValidation := helpers.LimitsAreValidated(dailyOut, weeklyOut, monthlyOut, annualOut)
	if boolValidation == true {
		return errors.New("failed to set limits - incorrect values")
	}

	actionRaw := d.Get("action").(string)
	var action xdr.ManageLimitsAction
	if actionRaw == "CREATE" {
		action = xdr.ManageLimitsActionCreate
	} else {
		for _, guess := range xdr.ManageLimitsActionAll {
			fmt.Println(guess.ShortString(), actionRaw)
			if guess.ShortString() == actionRaw {
				action = guess
			}
		}
		if action == 0 {
			return fmt.Errorf("unknown account rule action: %s", actionRaw)
		}
	}

	env, err := m.Builder.Transaction(m.Source).Op(&CreateLimit{
		Action:     action,
		Role:       rawAccountRole.(*xdr.Uint64),
		Id:         accountId,
		Type:       rawStatsType.(xdr.StatsOpType),
		Code:       d.Get("code").(xdr.AssetCode),
		Convert:    d.Get("convert").(bool),
		DailyOut:   rawDailyOut.(xdr.Uint64),
		WeeklyOut:  rawWeeklyOut.(xdr.Uint64),
		MonthlyOut: rawMonthlyOut.(xdr.Uint64),
		AnnualOut:  rawAnnualOut.(xdr.Uint64),
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

type CreateLimit struct {
	Action     xdr.ManageLimitsAction
	Role       *xdr.Uint64
	Id         *xdr.AccountId
	Type       xdr.StatsOpType
	Code       xdr.AssetCode
	Convert    bool
	DailyOut   xdr.Uint64
	WeeklyOut  xdr.Uint64
	MonthlyOut xdr.Uint64
	AnnualOut  xdr.Uint64
}

func (op *CreateLimit) XDR() (*xdr.Operation, error) {

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageLimits,
			ManageLimitsOp: &xdr.ManageLimitsOp{
				Details: xdr.ManageLimitsOpDetails{
					Action: xdr.ManageLimitsActionCreate,
					LimitsCreateDetails: &xdr.LimitsCreateDetails{
						AccountId:       op.Id,
						AccountRole:     op.Role,
						StatsOpType:     op.Type,
						AssetCode:       op.Code,
						IsConvertNeeded: op.Convert,
						DailyOut:        op.DailyOut,
						WeeklyOut:       op.WeeklyOut,
						MonthlyOut:      op.MonthlyOut,
						AnnualOut:       op.AnnualOut,
					},
				},
			},
		},
	}, nil
}

func resourceLimitsDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(Meta)

	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast limit id")
	}
	env, err := m.Builder.Transaction(m.Source).Op(&RemoveLimit{
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

type RemoveLimit struct {
	ID uint64
}

func (op *RemoveLimit) XDR() (*xdr.Operation, error) {
	Id := xdr.Uint64(op.ID)

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageLimits,
			ManageLimitsOp: &xdr.ManageLimitsOp{
				Details: xdr.ManageLimitsOpDetails{
					Action: xdr.ManageLimitsActionRemove,
					Id:     &Id,
				},
			},
		},
	}, nil
}
