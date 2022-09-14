package tokend

import (
	"context"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers/validation"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
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
			"role_id": {
				Type:     schema.TypeString,
				Required: true,
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
						"role_id": {
							Type:     schema.TypeString,
							Required: true,
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

	rawRoleID := d.Get("role_id")
	roleID, err := cast.ToUint64E(rawRoleID)
	if err != nil {
		return errors.Wrap(err, "failed to cast roleID to uint64")
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
		RoleID:      roleID,
		Signers:     signers,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}

	result := m.Horizon.Submitter().Submit(context.TODO(), env)
	if result.Err != nil {
		return errors.Wrap(result.Err, "failed to submit tx", logan.F{
			"tx_code":  result.TXCode,
			"op_codes": result.OpCodes,
		})
	}

	var txResult xdr.TransactionResult
	if err := xdr.SafeUnmarshalBase64(result.ResultXDR, &txResult); err != nil {
		return errors.Wrap(err, "failed to decode result")
	}

	d.SetId(destination)

	return nil
}

func resourceAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_account update is not implemented")
}

func resourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	return nil
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

	rawRoleID := d["role_id"]
	roleID, err := cast.ToUint64E(rawRoleID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast roleID to uint64")
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
		RoleID:    roleID,
		Details:   details,
	}, nil
}
