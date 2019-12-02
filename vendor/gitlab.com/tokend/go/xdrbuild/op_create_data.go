package xdrbuild

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateData struct {
	Value        string
	SecurityType uint64
}

func (op *CreateData) XDR() (*xdr.Operation, error) {
	ok := json.Valid([]byte(op.Value))
	if !ok {
		return nil, errors.New("data has invalid json struct")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateData,
			CreateDataOp: &xdr.CreateDataOp{
				Value:        xdr.Longstring(op.Value),
				SecurityType: xdr.Uint32(op.SecurityType),
			},
		},
	}, nil
}
