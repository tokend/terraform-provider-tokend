package helpers

import (
	"github.com/hashicorp/terraform/helper/schema"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
	regources "gitlab.com/tokend/regources"
)

func UpdateReviewableRequestOperations(d *schema.ResourceData, rawValue *regources.KeyValueEntryValue) error {
	if rawValue == nil {
		d.SetId("")
		return nil
	}

	switch rawValue.Type {
	case xdr.KeyValueEntryTypeString:
		opTypes, err := getOpTypes(*rawValue.Str)
		if err != nil {
			return errors.Wrap(err, "failed to get op types")
		}

		err = d.Set("op_types", opTypes)
		if err != nil {
			return errors.Wrap(err, "failed to set op types")
		}
	default:
		return errors.From(errors.New("Unexpected key value entry type"), logan.F{
			"entry_type": rawValue.Type.String(),
		})
	}
	return nil
}

func getOpTypes(value string) ([]string, error) {
	dest := make([]string, 0)
	err := xdr.SafeUnmarshalBase64(value, dest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal base64 value")
	}

	return dest, nil
}
