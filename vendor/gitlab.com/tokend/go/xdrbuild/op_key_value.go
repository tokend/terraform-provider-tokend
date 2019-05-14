package xdrbuild

import (
	"fmt"

	"gitlab.com/tokend/go/xdr"
)

type PutKeyValue struct {
	Key   string
	Value interface{}
}

func (op PutKeyValue) XDR() (*xdr.Operation, error) {
	var value xdr.KeyValueEntryValue
	switch v := op.Value.(type) {
	case uint32:
		value.Type = xdr.KeyValueEntryTypeUint32
		x := xdr.Uint32(v)
		value.Ui32Value = &x
	case uint64:
		value.Type = xdr.KeyValueEntryTypeUint64
		x := xdr.Uint64(v)
		value.Ui64Value = &x
	case string:
		value.Type = xdr.KeyValueEntryTypeString
		value.StringValue = &v
	default:
		return nil, fmt.Errorf("unsupported value type %T", v)
	}
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageKeyValue,
			ManageKeyValueOp: &xdr.ManageKeyValueOp{
				Key: xdr.Longstring(op.Key),
				Action: xdr.ManageKeyValueOpAction{
					Action: xdr.ManageKvActionPut,
					Value:  &value,
				},
			},
		},
	}, nil
}

type RemoveKeyValue struct {
	Key string
}

func (op RemoveKeyValue) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageKeyValue,
			ManageKeyValueOp: &xdr.ManageKeyValueOp{
				Key: xdr.Longstring(op.Key),
				Action: xdr.ManageKeyValueOpAction{
					Action: xdr.ManageKvActionRemove,
				},
			},
		},
	}, nil
}
