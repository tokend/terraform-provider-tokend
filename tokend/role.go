package tokend

import (
	"context"
	"fmt"
	"github.com/tokend/terraform-provider-tokend/tokend/connector"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Update: resourceRoleUpdate,
		Read:   resourceRoleRead,
		Delete: resourceRoleDelete,
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

func resourceRoleCreate(d *schema.ResourceData, _m interface{}) error {
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

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateRole{
		Details: VoidDetails{},
		Rules:   rules,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	resp, err := m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(connector.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
	}
	var txResult xdr.TransactionResult
	if err := xdr.SafeUnmarshalBase64(resp.Data.Attributes.ResultXdr, &txResult); err != nil {
		return errors.Wrap(err, "failed to decode result")
	}
	txCodes := *(txResult.Result.Results)
	roleID := txCodes[0].Tr.CreateRoleResult.Success.RoleId
	d.SetId(fmt.Sprintf("%d", roleID))
	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, _m interface{}) error {
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
	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.UpdateRole{
		ID:      id,
		Details: VoidDetails{},
		Rules:   rules,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	_, err = m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(connector.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
	}
	return nil
}

func resourceRoleRead(d *schema.ResourceData, _m interface{}) error {
	return nil
}

func resourceRoleDelete(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)
	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast account role id")
	}
	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.RemoveRole{
		ID: id,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	resp, err := m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(connector.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
	}
	var txResult xdr.TransactionResult
	if err := xdr.SafeUnmarshalBase64(resp.Data.Attributes.ResultXdr, &txResult); err != nil {
		return errors.Wrap(err, "failed to decode result")
	}
	return nil
}
