package tokend

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

	"gitlab.com/tokend/go/xdr"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceAccountRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountRuleCreate,
		Update: resourceAccountRuleUpdate,
		Read:   resourceAccountRuleRead,
		Delete: resourceAccountRuleDelete,
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"forbids": {
				Type:     schema.TypeBool,
				Optional: true,
				//Default:  false,
			},
			"details": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"entry_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entry": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

type VoidDetails struct{}

func (VoidDetails) MarshalJSON() ([]byte, error) {
	return []byte(`{}`), nil
}

func resourceAccountRuleCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	resource, err := helpers.AccountRuleEntry(d)
	if err != nil {
		return errors.Wrap(err, "failed to cast entry")
	}

	actionRaw := d.Get("action").(string)
	var action xdr.AccountRuleAction
	if actionRaw == "*" {
		action = xdr.AccountRuleActionAny
	} else {
		for _, guess := range xdr.AccountRuleActionAll {
			fmt.Println(guess.ShortString(), actionRaw)
			if guess.ShortString() == actionRaw {
				action = guess
			}
		}
		if action == 0 {
			return fmt.Errorf("unknown account rule action: %s", actionRaw)
		}
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateAccountRule{
		Resource: *resource,
		Action:   action,
		Forbid:   d.Get("forbids").(bool),
		Details:  VoidDetails{},
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
	ruleID := txCodes[0].Tr.ManageAccountRuleResult.Success.RuleId
	d.SetId(fmt.Sprintf("%d", ruleID))
	return nil
}

func resourceAccountRuleUpdate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	resource, err := helpers.AccountRuleEntry(d)
	if err != nil {
		return errors.Wrap(err, "failed to cast entry")
	}

	actionRaw := d.Get("action").(string)
	var action xdr.AccountRuleAction
	if actionRaw == "*" {
		action = xdr.AccountRuleActionAny
	} else {
		for _, guess := range xdr.AccountRuleActionAll {
			fmt.Println(guess.ShortString(), actionRaw)
			if guess.ShortString() == actionRaw {
				action = guess
			}
		}
		if action == 0 {
			return fmt.Errorf("unknown account rule action: %s", actionRaw)
		}
	}

	env, err := m.Builder.Transaction(m.Source).Op(&UpdateAccountRule{
		Resource: *resource,
		Action:   action,
		Forbid:   d.Get("forbids").(bool),
		Details:  VoidDetails{},
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

type UpdateAccountRule struct {
	ID       uint64
	Resource xdr.AccountRuleResource
	Action   xdr.AccountRuleAction
	Forbid   bool
	Details  json.Marshaler
}

func (op *UpdateAccountRule) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAccountRule,
			ManageAccountRuleOp: &xdr.ManageAccountRuleOp{
				Data: xdr.ManageAccountRuleOpData{
					Action: xdr.ManageAccountRuleActionUpdate,
					UpdateData: &xdr.UpdateAccountRuleData{
						RuleId:   xdr.Uint64(op.ID),
						Resource: op.Resource,
						Action:   op.Action,
						Forbids:  op.Forbid,
						Details:  xdr.Longstring(details),
					},
				},
			},
		},
	}, nil
}

func resourceAccountRuleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAccountRuleDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
	//return errors.New("tokend_account_rule delete is not implemented")
}
