package xdrbuild

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type UpdateData struct {
	ID    uint64
	Value string
}

func (op *UpdateData) XDR() (*xdr.Operation, error) {
	ok := json.Valid([]byte(op.Value))
	if !ok {
		return nil, errors.New("data has invalid json struct")
	}
	if op.ID == 0 {
		return nil, errors.New("id cannot be zero")
	}
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeUpdateData,
			UpdateDataOp: &xdr.UpdateDataOp{
				DataId: xdr.Uint64(op.ID), Value: xdr.Longstring(op.Value),
			},
		},
	}, nil
}
