package tokend

import (
	"context"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers/validation"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
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
			"public_key": {
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
				Elem: &schema.Schema{
					Type:         schema.TypeMap,
					ValidateFunc: validation.ValidateSignerRole,
				},
			},
		},
	}
}

func resourceAccountCreate(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	rawDestination := d.Get("public_key")
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
		signerMap, err := cast.ToStringMapE(rawSigner)
		if err != nil {
			return errors.Wrap(err, "failed to cast signer")
		}

		var role struct {
			ID uint64 `fig:"role_id"`
		}
		if err := figure.Out(&role).From(signerMap).Please(); err != nil {
			return errors.Wrap(err, "failed to figure out signer data")
		}

		signer := xdrbuild.SignerData{
			PublicKey: destination,
			Weight:    1000,
			RoleID:    role.ID,
			Details:   VoidDetails{},
		}

		signers = append(signers, signer)
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
	return errors.New("tokend_account read is not implemented")
}

func resourceAccountDelete(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_account delete is not implemented")
}
