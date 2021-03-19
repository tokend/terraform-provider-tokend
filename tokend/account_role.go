package tokend

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceAccountRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountRoleCreate,
		Update: resourceAccountRoleUpdate,
		Read:   resourceAccountRoleRead,
		Delete: resourceAccountRoleDelete,
		Schema: map[string]*schema.Schema{
			"details": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceAccountRoleCreate(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	var rules []uint64

	rawRules := d.Get("rules").([]interface{})
	for _, rawRule := range rawRules {
		rule, err := cast.ToUint64E(rawRule)
		if err != nil {
			return errors.Wrap(err, "failed to cast raw rule")
		}

		rules = append(rules, rule)
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateAccountRole{
		Details: VoidDetails{},
		Rules:   rules,
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
	roleID := txCodes[0].Tr.ManageAccountRoleResult.Success.RoleId

	d.SetId(fmt.Sprintf("%d", roleID))

	return nil
}

func resourceAccountRoleUpdate(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast account role id")
	}

	var rules []uint64

	rawRules := d.Get("rules").([]interface{})
	for _, rawRule := range rawRules {
		rule, err := cast.ToUint64E(rawRule)
		if err != nil {
			return errors.Wrap(err, "failed to cast raw rule")
		}

		rules = append(rules, rule)
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.UpdateAccountRole{
		ID:      id,
		Details: VoidDetails{},
		Rules:   rules,
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

	return nil
}

func resourceAccountRoleRead(d *schema.ResourceData, _m interface{}) error {
	return nil
}

func resourceAccountRoleDelete(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast account role id")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.RemoveAccountRole{
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
	roleID := txCodes[0].Tr.ManageAccountRoleResult.Success.RoleId

	d.SetId(fmt.Sprintf("%d", roleID))

	return nil
}
