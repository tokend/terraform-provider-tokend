package tokend

import (
	"context"
	"github.com/tokend/terraform-provider-tokend/tokend/connector"
	"gitlab.com/tokend/go/xdr"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers/validation"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceAccountSigner() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountSignerCreate,
		Update: resourceAccountSignerUpdate,
		Read:   resourceAccountSignerRead,
		Delete: resourceAccountSignerDelete,
		Schema: map[string]*schema.Schema{
			"public_key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.ValidateSource,
			},
			"weight": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identity": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role_ids": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
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

func resourceAccountSignerCreate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(Meta)

	publicKey := d.Get("public_key").(string)
	rawWeight := d.Get("weight")
	weight, err := cast.ToUint32E(rawWeight)
	if err != nil {
		return errors.Wrap(err, "failed to cast weight to uint32")
	}
	rawIdentity := d.Get("identity")
	identity, err := cast.ToUint32E(rawIdentity)
	if err != nil {
		return errors.Wrap(err, "failed to cast identity to uint32")
	}

	rawRoleIDs := d.Get("role_ids").([]interface{})
	roles := make([]uint64, 0, len(rawRoleIDs))

	for _, r := range rawRoleIDs {
		roleID, err := cast.ToUint64E(r)
		if err != nil {
			return errors.Wrap(err, "failed to cast roleID to uint64")
		}

		roles = append(roles, roleID)
	}

	rawDetails := d.Get("details")
	details, err := helpers.DetailsFromRaw(rawDetails)
	if err != nil {
		return errors.Wrap(err, "failed to get details")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateSigner{
		SignerData: xdrbuild.SignerData{
			PublicKey: publicKey,
			Weight:    weight,
			Identity:  identity,
			RoleIDs:   roles,
			Details:   details,
		},
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

	d.SetId(publicKey)
	return nil
}

func resourceAccountSignerUpdate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_signer update is not implemented")
}

func resourceAccountSignerDelete(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_signer delete is not implemented")
}

func resourceAccountSignerRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
