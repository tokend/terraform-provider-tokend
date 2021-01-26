package __old

import (
	. "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type ManageKeyValueOp struct {
	Key    string
	Uint32 *uint32
	Uint64 *uint64
	String *string
}

func isNil(i interface{}) error {
	_, isNil := Indirect(i)
	if !isNil {
		return errors.New("must be nil")
	}
	return nil
}

func (mkv ManageKeyValueOp) Validate() error {
	errs := Errors{
		"Key": Validate(&mkv.Key, Required),
	}

	if mkv.Uint32 != nil {
		errs["Uint32"] = Validate(mkv.Uint32, NotNil)
		errs["String"] = Validate(mkv.String, By(isNil))
		errs["Uint64"] = Validate(mkv.Uint64, By(isNil))
	}
	if mkv.String != nil {
		errs["String"] = Validate(mkv.String, NotNil)
		errs["Uint32"] = Validate(mkv.Uint32, By(isNil))
		errs["Uint64"] = Validate(mkv.Uint64, By(isNil))
	}
	if mkv.Uint64 != nil {
		errs["Uint64"] = Validate(mkv.Uint64, NotNil)
		errs["String"] = Validate(mkv.String, By(isNil))
		errs["Uint32"] = Validate(mkv.Uint32, By(isNil))
	}
	return errs.Filter()
}

func (mkv ManageKeyValueOp) XDR() (*xdr.Operation, error) {
	manageKvAction := xdr.ManageKvActionRemove
	keyValueEntryValue := (*xdr.KeyValueEntryValue)(nil)
	switch {
	case mkv.Uint32 != nil:
		manageKvAction = xdr.ManageKvActionPut
		val := xdr.Uint32(*mkv.Uint32)
		keyValueEntryValue = &xdr.KeyValueEntryValue{
			Type:      xdr.KeyValueEntryTypeUint32,
			Ui32Value: &val,
		}
	case mkv.Uint64 != nil:
		manageKvAction = xdr.ManageKvActionPut
		val := xdr.Uint64(*mkv.Uint64)
		keyValueEntryValue = &xdr.KeyValueEntryValue{
			Type:      xdr.KeyValueEntryTypeUint64,
			Ui64Value: &val,
		}
	case mkv.String != nil:
		manageKvAction = xdr.ManageKvActionPut
		keyValueEntryValue = &xdr.KeyValueEntryValue{
			Type:        xdr.KeyValueEntryTypeString,
			StringValue: mkv.String,
		}
	}

	if keyValueEntryValue == nil {
		keyValueEntryValue = &xdr.KeyValueEntryValue{}
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageKeyValue,
			ManageKeyValueOp: &xdr.ManageKeyValueOp{
				Key: xdr.Longstring(mkv.Key),
				Action: xdr.ManageKeyValueOpAction{
					Action: manageKvAction,
					Value:  keyValueEntryValue,
				},
			},
		},
	}, nil
}
