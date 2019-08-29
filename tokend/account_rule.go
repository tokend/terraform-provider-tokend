package tokend

import (
	"context"
	"fmt"

	"github.com/spf13/cast"

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

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.UpdateAccountRule{
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

func resourceAccountRuleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAccountRuleDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(Meta)
	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast account rule id")
	}
	env, err := m.Builder.Transaction(m.Source).Op(&RemoveAccountRule{
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
	ruleID := txCodes[0].Tr.ManageAccountRuleResult.Success.RuleId
	d.SetId(fmt.Sprintf("%d", ruleID))
	return nil
}

type RemoveAccountRule struct {
	ID uint64
}

func (op *RemoveAccountRule) XDR() (*xdr.Operation, error) {

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAccountRule,
			ManageAccountRuleOp: &xdr.ManageAccountRuleOp{
				Data: xdr.ManageAccountRuleOpData{
					Action: xdr.ManageAccountRuleActionRemove,
					RemoveData: &xdr.RemoveAccountRuleData{
						RuleId: xdr.Uint64(op.ID),
					},
				},
			},
		},
	}, nil
}
