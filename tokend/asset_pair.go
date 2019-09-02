package tokend

import (
	"context"
	"fmt"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

	"gitlab.com/tokend/go/xdr"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/amount"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceAssetPair() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssetPairCreate,
		Update: resourceAssetPairUpdate,
		Read:   resourceAssetPairRead,
		Delete: resourceAssetPairDelete,
		Schema: map[string]*schema.Schema{
			"base": {
				Type:     schema.TypeString,
				Required: true,
			},
			"quote": {
				Type:     schema.TypeString,
				Required: true,
			},
			"current_price": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"price": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"price_correction": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_price_step": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policies": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceAssetPairCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	priceRaw := d.Get("price").(string)
	price, err := amount.Parse(priceRaw)
	if err != nil {
		return errors.Wrap(err, "failed to parse amount")
	}

	base := d.Get("base").(string)
	quote := d.Get("quote").(string)

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateAssetPair{
		Base:                    base,
		Quote:                   quote,
		PhysicalPrice:           price,
		PhysicalPriceCorrection: 0,
		MaxPriceStep:            0,
		Policies:                1,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	result := m.Horizon.Submitter().Submit(context.TODO(), env)
	if result.Err != nil {
		return errors.Wrapf(result.Err, "failed to submit tx: %s %q", result.TXCode, result.OpCodes)
	}
	d.SetId(fmt.Sprintf("%s-%s", base, quote))
	return nil
}

//FIXME this func must update policy or price fields (now policies only)
func resourceAssetPairUpdate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(Meta)

	base := d.Get("base").(string)
	quote := d.Get("quote").(string)

	maxPriceStep := d.Get("max_price_step").(string)
	maxStep, err := amount.Parse(maxPriceStep)
	if err != nil {
		return errors.Wrap(err, "failed to parse step")
	}

	policyRaw := d.Get("policies").([]interface{})
	policyCode, err := helpers.PoliciesFromRaw(policyRaw)
	if err != nil {
		return errors.Wrap(err, "failed to cast policy")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.UpdateAssetPairPolicies{
		Base:         base,
		Quote:        quote,
		MaxPriceStep: maxStep,
		Policies:     int32(policyCode),
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

func resourceAssetPairUpdatePrice(d *schema.ResourceData, meta interface{}) error {
	m := meta.(Meta)

	base := d.Get("base").(string)
	quote := d.Get("quote").(string)

	physicalPriceRaw := d.Get("price").(string)
	physicalPrice, err := amount.Parse(physicalPriceRaw)
	if err != nil {
		return errors.Wrap(err, "failed to parse price")
	}

	physicalPriceCorrection := d.Get("price_correction").(string)
	priceCorrection, err := amount.Parse(physicalPriceCorrection)
	if err != nil {
		return errors.Wrap(err, "failed to parse correction")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.UpdateAssetPairPrice{
		Base:                    base,
		Quote:                   quote,
		PhysicalPrice:           physicalPrice,
		PhysicalPriceCorrection: priceCorrection,
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
func resourceAssetPairRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAssetPairDelete(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)
	base := d.Get("base").(string)
	quote := d.Get("quote").(string)

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.RemoveAssetPair{
		Base:  base,
		Quote: quote,
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
	CurrentPrice := txCodes[0].Tr.ManageAssetPairResult.Success.CurrentPrice
	d.SetId(fmt.Sprintf("%d", CurrentPrice))
	return nil
}
