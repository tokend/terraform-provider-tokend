package xdrbuild

import "gitlab.com/tokend/go/xdr"

type CheckSaleState struct {
	ID uint64
}

func (op CheckSaleState) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCheckSaleState,
			CheckSaleStateOp: &xdr.CheckSaleStateOp{
				SaleId: xdr.Uint64(op.ID),
			},
		},
	}, nil
}
