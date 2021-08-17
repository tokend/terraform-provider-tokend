package tokend

import (
	"context"
	"fmt"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

	"github.com/spf13/cast"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers/validation"
	"gitlab.com/tokend/go/xdr"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/amount"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceAsset() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssetCreate,
		Update: resourceAssetUpdate,
		Read:   resourceAssetRead,
		Delete: resourceAssetDelete,
		Schema: map[string]*schema.Schema{
			"code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_issuance_amount": {
				Type:     schema.TypeString,
				Required: true,
			},
			"initial_pre_issuance_amount": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pre_issuance_signer": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.ValidateSource,
			},
			"trailing_digits_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"policies": {
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

func resourceAssetCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	maxIssuanceAmount, err := amount.ParseU(d.Get("max_issuance_amount").(string))
	if err != nil {
		return errors.Wrap(err, "failed to parse max_issuance_amount")
	}

	preIssuanceAmount, err := amount.ParseU(d.Get("initial_pre_issuance_amount").(string))
	if err != nil {
		return errors.Wrap(err, "failed to parse initial_pre_issuance_amount")
	}

	var zero uint32 = 0

	var policies uint32
	for _, policyRaw := range d.Get("policies").([]interface{}) {
		policy, err := cast.ToStringE(policyRaw)
		if err != nil {
			return errors.Wrap(err, "failed to cast policy")
		}
		ok := false
		for _, guess := range xdr.AssetPolicyAll {
			if guess.ShortString() == policy {
				ok = true
				policies |= uint32(guess)
			}
		}
		if !ok {
			panic(errors.Errorf("invalid policy name: %s", policy))
		}
	}

	rawDetails := d.Get("details")
	details, err := helpers.DetailsFromRaw(rawDetails)
	if err != nil {
		return errors.Wrap(err, "failed to get details")
	}

	rawAssetType := d.Get("type")
	var assetType uint64
	if rawAssetType != nil {
		assetType, err = cast.ToUint64E(rawAssetType)
		if err != nil {
			return errors.Wrap(err, "failed to cast asset type")
		}
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateAsset{
		CreatorDetails:           details,
		Code:                     d.Get("code").(string),
		MaxIssuanceAmount:        maxIssuanceAmount,
		PreIssuanceSigner:        d.Get("pre_issuance_signer").(string),
		InitialPreIssuanceAmount: preIssuanceAmount,
		TrailingDigitsCount:      uint32(d.Get("trailing_digits_count").(int)),
		Policies:                 policies,
		AllTasks:                 &zero,
		Type:                     assetType,
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

	d.SetId(d.Get("code").(string))

	return nil
}

func resourceAssetUpdate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_asset update is not implemented")
}

func resourceAssetRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAssetDelete(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	assetCode, err := cast.ToStringE(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast asset code")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.RemoveAsset{
		Code: assetCode,
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
	code := txCodes[0].Tr.RemoveAssetResult.Code

	d.SetId(fmt.Sprintf("%d", code))

	return nil
}
