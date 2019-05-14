package tokend

import (
	"context"
	"fmt"

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
			"price": {
				Type:     schema.TypeString,
				Required: true,
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

func resourceAssetPairUpdate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("tokend_asset_pair update is not implemented")
}

func resourceAssetPairRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAssetPairDelete(d *schema.ResourceData, _m interface{}) error {
	return errors.New("tokend_asset_pair delete is not implemented")
}
