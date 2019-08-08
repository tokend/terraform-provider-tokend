package tokend

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceSignerRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSignerRuleCreate,
		Update: resourceSignerRuleUpdate,
		Read:   resourceSignerRuleRead,
		Delete: resourceSignerRuleDelete,
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

func resourceSignerRuleCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)
	resource, err := helpers.SignerRuleEntry(d)
	if err != nil {
		return errors.Wrap(err, "failed to cast entry")
	}
	actionRaw := d.Get("action").(string)
	var action xdr.SignerRuleAction
	if actionRaw == "*" {
		action = xdr.SignerRuleActionAny
	} else {
		for _, guess := range xdr.SignerRuleActionAll {
			fmt.Println(guess.ShortString(), actionRaw)
			if guess.ShortString() == actionRaw {
				action = guess
			}
		}
		if action == 0 {
			return fmt.Errorf("unknown account rule action: %s", actionRaw)
		}
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateSignerRule{
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
	ruleID := txCodes[0].Tr.ManageSignerRuleResult.Success.RuleId
	d.SetId(fmt.Sprintf("%d", ruleID))
	return nil
}

func resourceSignerRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(Meta)

	resource, err := helpers.SignerRuleEntry(d)
	if err != nil {
		return errors.Wrap(err, "failed to cast entry")
	}

	actionRaw := d.Get("action").(string)
	var action xdr.SignerRuleAction
	if actionRaw == "*" {
		action = xdr.SignerRuleActionAny
	} else {
		for _, guess := range xdr.SignerRuleActionAll {
			fmt.Println(guess.ShortString(), actionRaw)
			if guess.ShortString() == actionRaw {
				action = guess
			}
		}
		if action == 0 {
			return fmt.Errorf("unknown signer rule action: %s", actionRaw)
		}
	}

	env, err := m.Builder.Transaction(m.Source).Op(&UpdateSignerRule{
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

// TODO:- Add this part to op_manage_signer_rule.go
type UpdateSignerRule struct {
	ID       uint64
	Resource xdr.SignerRuleResource
	Action   xdr.SignerRuleAction
	Forbid   bool
	Details  json.Marshaler
}

func (op *UpdateSignerRule) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSignerRule,
			ManageSignerRuleOp: &xdr.ManageSignerRuleOp{
				Data: xdr.ManageSignerRuleOpData{
					Action: xdr.ManageSignerRuleActionUpdate,
					UpdateData: &xdr.UpdateSignerRuleData{
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

func resourceSignerRuleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSignerRuleDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(Meta)
	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast account rule id")
	}
	env, err := m.Builder.Transaction(m.Source).Op(&RemoveSignerRule{
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
	ruleID := txCodes[0].Tr.ManageSignerRuleResult.Success.RuleId
	d.SetId(fmt.Sprintf("%d", ruleID))
	return nil
}

// TODO:- Add this part to op_manage_signer_rule.go
type RemoveSignerRule struct {
	ID uint64
}

func (op *RemoveSignerRule) XDR() (*xdr.Operation, error) {

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSignerRule,
			ManageSignerRuleOp: &xdr.ManageSignerRuleOp{
				Data: xdr.ManageSignerRuleOpData{
					Action: xdr.ManageSignerRuleActionRemove,
					RemoveData: &xdr.RemoveSignerRuleData{
						RuleId: xdr.Uint64(op.ID),
					},
				},
			},
		},
	}, nil
}
