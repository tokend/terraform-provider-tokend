package xdrbuild

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type RemoveData struct {
	ID uint64
}

func (op *RemoveData) XDR() (*xdr.Operation, error) {
	if op.ID == 0 {
		return nil, errors.New("id cannot be zero")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeRemoveData,
			RemoveDataOp: &xdr.RemoveDataOp{
				DataId: xdr.Uint64(op.ID),
			},
		},
	}, nil
}
