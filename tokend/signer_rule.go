package tokend

import (
	"context"
	"fmt"

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
	return errors.New("tokend_signer_rule update is not implemented")
}

func resourceSignerRuleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSignerRuleDelete(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_signer_rule delete is not implemented")
}
