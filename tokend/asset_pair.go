package tokend

import (
	"context"
	"fmt"
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
	m := meta.(Meta)
	actionRaw := d.Get("action").(string)
	var action xdr.ManageAssetPairAction
	switch actionRaw {
	case "CREATE":
		action = xdr.ManageAssetPairActionCreate

	case "UPDATE_PRICE":
		action = xdr.ManageAssetPairActionUpdatePrice

	case "UPDATE_POLICIES":
		action = xdr.ManageAssetPairActionUpdatePolicies

	default:
		action = xdr.ManageAssetPairAction(xdr.ManageAssetPairResultCodeInvalidAction)
		return fmt.Errorf("INVALID_ACTION %s", action)

	}
	base := d.Get("base").(string)
	quote := d.Get("quote").(string)

	priceRaw := d.Get("price").(string)
	price, err := amount.Parse(priceRaw)
	if err != nil {
		return errors.Wrap(err, "failed to parse price")
	}

	physicalPriceCorrection := d.Get("physicalPriceCorrection").(string)
	priceCorrection, err := amount.Parse(physicalPriceCorrection)
	if err != nil {
		return errors.Wrap(err, "failed to parse correction")
	}

	maxPriceStep := d.Get("maxPriceStep").(string)
	maxStep, err := amount.Parse(maxPriceStep)
	if err != nil {
		return errors.Wrap(err, "failed to parse step")
	}

	if price < 0 || priceCorrection < 0 || maxStep < 0 || maxStep > 100 {
		return errors.New("MALFORMED")
	}
	policies := d.Get("policies").(int32)
	if policies < 0 {
		return errors.New("INVALID_POLICIES")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.UpdateAssetPairPolicies{
		Base:  base,
		Quote: quote,
		//PhysicalPrice:           price,
		PhysicalPriceCorrection: priceCorrection,
		MaxPriceStep:            maxStep,
		Policies:                policies,
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	result := m.Horizon.Submitter().Submit(context.TODO(), env)
	if result.Err != nil {
		return errors.Wrapf(result.Err, "failed to submit tx: %s %q", result.TXCode, result.OpCodes)
	}

	return nil
	//return errors.New("tokend_asset_pair update is not implemented")
}

func resourceAssetPairRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAssetPairDelete(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)
	base := d.Get("base").(string)
	quote := d.Get("quote").(string)

	env, err := m.Builder.Transaction(m.Source).Op(&RemoveAssetPair{
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
	//return errors.New("tokend_asset_pair delete is not implemented")
}

type RemoveAssetPair struct {
	Base  string
	Quote string
}

func (ap RemoveAssetPair) XDR() (*xdr.Operation, error) {
	op := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeRemoveAssetPair,
			RemoveAssetPairOp: &xdr.RemoveAssetPairOp{
				Base:  xdr.AssetCode(ap.Base),
				Quote: xdr.AssetCode(ap.Quote),
			},
		},
	}
	return op, nil
}
