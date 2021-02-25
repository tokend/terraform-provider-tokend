package tokend

import (
	"context"
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
			"role_id": {
				Type:     schema.TypeString,
				Required: true,
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

	rawRoleID := d.Get("role_id")
	roleID, err := cast.ToUint64E(rawRoleID)
	if err != nil {
		return errors.Wrap(err, "failed to cast roleID to uint64")
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
			RoleID:    roleID,
			Details:   details,
		},
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
