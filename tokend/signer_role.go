package tokend

import (
	"context"
	"fmt"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

	"github.com/spf13/cast"

	"gitlab.com/tokend/go/xdr"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceSignerRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceSignerRoleCreate,
		Update: resourceSignerRoleUpdate,
		Read:   resourceSignerRoleRead,
		Delete: resourceSignerRoleDelete,
		Schema: map[string]*schema.Schema{
			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"details": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceSignerRoleCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	rawDetails := d.Get("details")
	details, err := helpers.DetailsFromRaw(rawDetails)
	if err != nil {
		return errors.Wrap(err, "failed to get details")
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

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateSignerRole{
		Details: details,
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
	roleID := txCodes[0].Tr.ManageSignerRoleResult.Success.RoleId
	d.SetId(fmt.Sprintf("%d", roleID))
	return nil
}

func resourceSignerRoleUpdate(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	rawDetails := d.Get("details")
	details, err := helpers.DetailsFromRaw(rawDetails)
	if err != nil {
		return errors.Wrap(err, "failed to get details")
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

	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast id")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.UpdateSignerRole{
		ID:      id,
		Details: details,
		Rules:   rules,
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

func resourceSignerRoleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSignerRoleDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(Meta)

	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast id")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&RemoveSignerRole{
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
	roleID := txCodes[0].Tr.ManageSignerRoleResult.Success.RoleId
	d.SetId(fmt.Sprintf("%d", roleID))
	return nil

	//return errors.New("tokend_signer_role delete is not implemented")

}

// TODO:- Add this part to op_manage_signer_role.go
type RemoveSignerRole struct {
	ID uint64
}

func (op *RemoveSignerRole) XDR() (*xdr.Operation, error) {

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSignerRole,
			ManageSignerRoleOp: &xdr.ManageSignerRoleOp{
				Data: xdr.ManageSignerRoleOpData{
					Action: xdr.ManageSignerRoleActionRemove,
					RemoveData: &xdr.RemoveSignerRoleData{
						RoleId: xdr.Uint64(op.ID),
					},
				},
			},
		},
	}, nil
}
