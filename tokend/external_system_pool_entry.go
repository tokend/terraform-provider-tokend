package tokend

import (
	"context"
	"fmt"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

const (
	externalSystemTypeStellarKey  = "external_system_type:stellar"
	externalSystemTypeEthereumKey = "external_system_type:ethereum"
)

func resourceExternalSystemPoolEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceExternalSystemPoolEntryCreate,
		Update: resourceExternalSystemPoolEntryUpdate,
		Read:   resourceExternalSystemPoolEntryRead,
		Delete: resourceExternalSystemPoolEntryDelete,
		Schema: map[string]*schema.Schema{
			"target_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"external_system_type": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data": {
				Type:      schema.TypeMap,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceExternalSystemPoolEntryCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	stellarExternalSystemType, err := m.Connector.KeyValues().Value(externalSystemTypeStellarKey)
	if err != nil {
		return errors.Wrap(err, "failed to get stellar external system type")
	}
	if stellarExternalSystemType == nil {
		return errors.New("stellar external system type key value not set")
	}

	ethereumExternalSystemType, err := m.Connector.KeyValues().Value(externalSystemTypeEthereumKey)
	if err != nil {
		return errors.Wrap(err, "failed to get ethereum external system type")
	}
	if ethereumExternalSystemType == nil {
		return errors.New("ethereum external system type key value not set")
	}

	externalType := d.Get("external_system_type").(int)
	count := d.Get("target_count").(int)

	var envelopes []string
	switch uint32(externalType) {
	case *stellarExternalSystemType.U32:
		envelopes, err = helpers.StellarExternalTypeEnvelopes(count, uint32(externalType), d, m.Builder, m.Signer)
		if err != nil {
			return errors.Wrap(err, "failed to create transaction envelope")
		}
	case *ethereumExternalSystemType.U32:
		envelopes, err = helpers.EthereumExternalTypeEnvelopes(count, uint32(externalType), d, m.Builder, m.Signer)
		if err != nil {
			return errors.Wrap(err, "failed to create transaction envelope")
		}
	default:
		return errors.New(fmt.Sprintf("unknown external system type: %d", externalType))
	}

	if len(envelopes) == 0 {
		return errors.New("empty transaction envelopes")
	}
	for _, envelope := range envelopes {
		result := m.Horizon.Submitter().Submit(context.TODO(), envelope)
		if result.Err != nil {
			return errors.Wrapf(result.Err, "failed to submit tx: %s %q", result.TXCode, result.OpCodes)
		}
	}
	return nil
}

func resourceExternalSystemPoolEntryUpdate(_ *schema.ResourceData, _ interface{}) error {
	return errors.New("tokend_external_system_pool_entry update is not implemented")
}

func resourceExternalSystemPoolEntryRead(_ *schema.ResourceData, _ interface{}) error {
	return nil
}

func resourceExternalSystemPoolEntryDelete(_ *schema.ResourceData, _ interface{}) error {
	return errors.New("tokend_external_system_pool_entry delete is not implemented")
}
