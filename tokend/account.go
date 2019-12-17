package tokend

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/tokend/connectors/submit"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers/validation"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spf13/cast"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountCreate,
		Read:   resourceAccountRead,
		Update: resourceAccountUpdate,
		Delete: resourceAccountDelete,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.ValidateSource,
			},
			"role_ids": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"signers": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"public_key": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.ValidateSource,
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1000,
						},
						"identity": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
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
				},
			},
		},
	}
}

func resourceAccountCreate(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	rawDestination := d.Get("account_id")
	destination, err := cast.ToStringE(rawDestination)
	if err != nil {
		return errors.Wrap(err, "failed to cast public_key to string")
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
	rawSigners := d.Get("signers").(*schema.Set).List()

	var signers []xdrbuild.SignerData
	for _, rawSigner := range rawSigners {
		signer, err := getSigner(rawSigner)
		if err != nil {
			return errors.Wrap(err, "failed to get signer")
		}

		signers = append(signers, *signer)
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateAccount{
		Destination: destination,
		RoleIDs:     roles,
		Signers:     signers,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}

	resp, err := m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(submit.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
	}
	var txResult xdr.TransactionResult
	if err := xdr.SafeUnmarshalBase64(resp.Data.Attributes.ResultXdr, &txResult); err != nil {
		return errors.Wrap(err, "failed to decode result")
	}

	d.SetId(destination)
	return nil
}

func resourceAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_account update is not implemented")
}

func resourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_account read is not implemented")
}

func resourceAccountDelete(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_account delete is not implemented")
}

func getSigner(rawSigner interface{}) (*xdrbuild.SignerData, error) {
	d, err := cast.ToStringMapE(rawSigner)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get signer")
	}

	publicKey := d["public_key"].(string)
	rawWeight := d["weight"]
	weight, err := cast.ToUint32E(rawWeight)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast weight to uint32")
	}
	rawIdentity := d["identity"]
	identity, err := cast.ToUint32E(rawIdentity)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast identity to uint32")
	}

	rawRoleIDs := d["role_ids"].([]interface{})
	roles := make([]uint64, 0, len(rawRoleIDs))

	for _, r := range rawRoleIDs {
		roleID, err := cast.ToUint64E(r)
		if err != nil {
			return nil, errors.Wrap(err, "failed to cast roleID to uint64")
		}

		roles = append(roles, roleID)
	}

	rawDetails := d["details"]
	details, err := helpers.DetailsFromRaw(rawDetails)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get details")
	}

	return &xdrbuild.SignerData{
		PublicKey: publicKey,
		Weight:    weight,
		Identity:  identity,
		RoleIDs:   roles,
		Details:   details,
	}, nil
}
