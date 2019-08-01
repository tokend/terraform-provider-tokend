package xdrbuild

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/tokend/go/xdr"
)

type BindExternalSystemAccountIDOp struct {
	ExternalSystem int32
}

func (op *BindExternalSystemAccountIDOp) Validate() error {
	return validation.ValidateStruct(op,
		validation.Field(&op.ExternalSystem, validation.Required, validation.Min(1)),
	)
}

func (op *BindExternalSystemAccountIDOp) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeBindExternalSystemAccountId,
			BindExternalSystemAccountIdOp: &xdr.BindExternalSystemAccountIdOp{
				ExternalSystemType: xdr.Int32(op.ExternalSystem),
			},
		},
	}, nil
}
