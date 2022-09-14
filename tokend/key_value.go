package tokend

import (
	"context"
	"fmt"

	"gitlab.com/tokend/go/xdr"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers/validation"
	"gitlab.com/tokend/go/xdrbuild"
)

var keyValueSchema = map[string]*schema.Schema{
	"key": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.NonEmptyString,
	},
	"value_type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.NonEmptyString,
	},
	"value": {
		Type:     schema.TypeString,
		Required: true,
	},
}

func resourceKeyValue() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeyValueCreate,
		Update: resourceKeyValueUpdate,
		Read:   resourceKeyValueRead,
		Delete: resourceKeyValueDelete,
		Schema: keyValueSchema,
	}
}

func resourceKeyValueCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	value := d.Get("value")

	switch t := d.Get("value_type").(string); t {
	case "string":
	case "uint32":
		value, err = cast.ToUint32E(value)
		if err != nil {
			return errors.Wrap(err, "failed to cast value to uint32")
		}
	case "uint64":
		value, err = cast.ToUint64E(value)
		if err != nil {
			return errors.Wrap(err, "failed to cast value to uint64")
		}
	default:
		return fmt.Errorf("value type is not supported: %s", t)
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.PutKeyValue{
		Key:   d.Get("key").(string),
		Value: value,
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

	d.SetId(d.Get("key").(string))

	return nil
}

func resourceKeyValueUpdate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	value := d.Get("value")

	switch t := d.Get("value_type").(string); t {
	case "string":
	case "uint32":
		value, err = cast.ToUint32E(value)
		if err != nil {
			return errors.Wrap(err, "failed to cast value to uint32")
		}
	case "uint64":
		value, err = cast.ToUint64E(value)
		if err != nil {
			return errors.Wrap(err, "failed to cast value to uint64")
		}
	default:
		return fmt.Errorf("value type is not supported: %s", t)
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.PutKeyValue{
		Key:   d.Get("key").(string),
		Value: value,
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

	return nil
}

func resourceKeyValueRead(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	key := d.Get("key").(string)

	rawValue, err := m.Horizon.KeyValues().Value(key)
	if err != nil {
		return errors.Wrap(err, "failed to get value")
	}

	err = helpers.UpdateKeyValue(d, rawValue)
	if err != nil {
		return errors.Wrap(err, "failed to update value")
	}

	return nil
}

func resourceKeyValueDelete(d *schema.ResourceData, _m interface{}) error {
	m := _m.(Meta)

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.RemoveKeyValue{
		Key: d.Get("key").(string),
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}

	result := m.Horizon.Submitter().Submit(context.TODO(), env)
	if result.Err != nil {
		return errors.Wrap(result.Err, "failed to submit tx")
	}

	var txResult xdr.TransactionResult
	if err := xdr.SafeUnmarshalBase64(result.ResultXDR, &txResult); err != nil {
		return errors.Wrap(err, "failed to decode result")
	}

	return nil
}
