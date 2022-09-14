package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
	regources "gitlab.com/tokend/regources/generated"
)

func UpdateKeyValue(d *schema.ResourceData, rawValue *regources.KeyValueEntryValue) error {
	if rawValue == nil {
		d.SetId("")
		return nil
	}

	switch rawValue.Type {
	case xdr.KeyValueEntryTypeString:
		err := d.Set("value_type", "string")
		if err != nil {
			return errors.Wrap(err, "failed to set string value type")
		}
		err = d.Set("value", *rawValue.Str)
		if err != nil {
			return errors.Wrap(err, "failed to set string value")
		}
	case xdr.KeyValueEntryTypeUint32:
		err := d.Set("value_type", "uint32")
		if err != nil {
			return errors.Wrap(err, "failed to set uint32 value type")
		}

		err = d.Set("value", fmt.Sprintf("%d", *rawValue.U32))
		if err != nil {
			return errors.Wrap(err, "failed to set uint32 value")
		}
	case xdr.KeyValueEntryTypeUint64:
		err := d.Set("value_type", "uint64")
		if err != nil {
			return errors.Wrap(err, "failed to set uint64 value type")
		}

		err = d.Set("value", fmt.Sprintf("%d", *rawValue.U64))
		if err != nil {
			return errors.Wrap(err, "failed to set uin64 value")
		}
	}
	return nil
}
