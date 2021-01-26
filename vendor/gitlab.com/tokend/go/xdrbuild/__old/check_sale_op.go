package __old

import (
	"github.com/go-ozzo/ozzo-validation"
	"gitlab.com/tokend/go/xdr"
)

type CheckSaleOp struct {
	SaleID uint64
}

func (op CheckSaleOp) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.SaleID, validation.Required),
	)
}

func (op CheckSaleOp) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCheckSaleState,
			CheckSaleStateOp: &xdr.CheckSaleStateOp{
				SaleId: xdr.Uint64(op.SaleID),
			},
		},
	}, nil
}
