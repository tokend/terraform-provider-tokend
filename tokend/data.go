package tokend

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers"
	"gitlab.com/tokend/go/amount"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceData() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataCreate,
		Update: resourceDataUpdate,
		Read:   resourceDataRead,
		Delete: resourceDataDelete,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"value": {
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

func resourceDataCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	dataType, err := amount.ParseU(d.Get("type").(string))
	if err != nil {
		return errors.Wrap(err, "failed to parse max_issuance_amount")
	}

	rawValue := d.Get("details")
	value, err := helpers.DetailsFromRaw(rawValue)
	if err != nil {
		return errors.Wrap(err, "failed to get details")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateData{
		Type:  dataType,
		Value: value,
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
	id := txCodes[0].Tr.ManageAccountRoleResult.Success.RoleId

	d.SetId(fmt.Sprintf("%d", id))

	return nil
}

func resourceDataUpdate(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast account role id")
	}

	rawValue := d.Get("details")
	value, err := helpers.DetailsFromRaw(rawValue)
	if err != nil {
		return errors.Wrap(err, "failed to get details")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.UpdateData{
		ID:    id,
		Value: value,
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

func resourceDataRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDataDelete(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast account role id")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.RemoveData{
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

	return nil
}
