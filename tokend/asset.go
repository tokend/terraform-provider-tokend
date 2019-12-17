package tokend

import (
	"context"
	"gitlab.com/tokend/connectors/submit"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

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
			"max_issuance_amount": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trailing_digits_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
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

	rawSecType := d.Get("security_type").(string)
	secType, err := helpers.WildCardUint32FromRaw(rawSecType)

	rawDetails := d.Get("details")
	details, err := helpers.DetailsFromRaw(rawDetails)
	if err != nil {
		return errors.Wrap(err, "failed to get details")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateAsset{
		Code:                d.Get("code").(string),
		MaxIssuanceAmount:   maxIssuanceAmount,
		TrailingDigitsCount: uint32(d.Get("trailing_digits_count").(int)),
		SecurityType:        secType,
		CreatorDetails:      details,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	_, err = m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(submit.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
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
	return nil
}
