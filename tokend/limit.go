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
		Read:   resourceSignerRuleRead,
		Update: resourceSignerRuleUpdate,
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

	rawAccountRole := d.Get("role")
	var typesCode int32
	typeRaw := d.Get("stats_type").(string)
	oneType, err := cast.ToStringE(typeRaw)
	if err != nil {
		return errors.Wrap(err, "failed to cast type")
	}
	ok := false

	for index, guess := range xdr.StatsOpTypeAll {
		if guess.ShortString() == oneType {
			ok = true
			typesCode |= int32(index)
		}
	}
	if !ok {
		panic(errors.Errorf("invalid type code: %s", oneType))
	}

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

	if helpers.ValidateLimits(dailyOut, weeklyOut, monthlyOut, annualOut) == false {
		return errors.New("failed to set limits - incorrect values")
	}
	accountRole := xdr.Uint64((rawAccountRole.(int)))
	env, err := m.Builder.Transaction(m.Source).Op(&CreateLimit{
		Action:     xdr.ManageLimitsActionCreate,
		Role:       &accountRole,
		Id:         accountID,
		Type:       xdr.StatsOpType(typesCode),
		Code:       xdr.AssetCode(d.Get("asset_code").(string)),
		Convert:    d.Get("convert").(bool),
		DailyOut:   xdr.Uint64(rawDailyOut.(int)),
		WeeklyOut:  xdr.Uint64(rawWeeklyOut.(int)),
		MonthlyOut: xdr.Uint64(rawMonthlyOut.(int)),
		AnnualOut:  xdr.Uint64(rawAnnualOut.(int)),
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
	details := xdr.LimitsCreateDetails{
		StatsOpType:     op.Type,
		AssetCode:       op.Code,
		IsConvertNeeded: op.Convert,
		DailyOut:        op.DailyOut,
		WeeklyOut:       op.WeeklyOut,
		MonthlyOut:      op.MonthlyOut,
		AnnualOut:       op.AnnualOut,
	}

	if op.Id == nil {
		details.AccountRole = op.Role
	} else if op.Role == nil {
		details.AccountId = op.Id
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageLimits,
			ManageLimitsOp: &xdr.ManageLimitsOp{
				Details: xdr.ManageLimitsOpDetails{
					Action:              xdr.ManageLimitsActionCreate,
					LimitsCreateDetails: &details,
				},
			},
		},
	}, nil
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
