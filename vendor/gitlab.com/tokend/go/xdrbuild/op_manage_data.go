package xdrbuild

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateData struct {
	Type  uint64
	Value json.Marshaler
}

func (c CreateData) XDR() (*xdr.Operation, error) {
	value, err := c.Value.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateData,
			CreateDataOp: &xdr.CreateDataOp{
				Type:  xdr.Uint64(c.Type),
				Value: xdr.Longstring(value),
				Ext:   xdr.EmptyExt{},
			},
		},
	}, nil
}

type UpdateData struct {
	ID    uint64
	Value json.Marshaler
}

func (u UpdateData) XDR() (*xdr.Operation, error) {
	value, err := u.Value.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeUpdateData,
			UpdateDataOp: &xdr.UpdateDataOp{
				DataId: xdr.Uint64(u.ID),
				Value:  xdr.Longstring(value),
				Ext:    xdr.EmptyExt{},
			},
		},
	}, nil
}

type RemoveData struct {
	ID uint64
}

func (r RemoveData) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeRemoveData,
			RemoveDataOp: &xdr.RemoveDataOp{
				DataId: xdr.Uint64(r.ID),
				Ext:    xdr.EmptyExt{},
			},
		},
	}, nil
}
