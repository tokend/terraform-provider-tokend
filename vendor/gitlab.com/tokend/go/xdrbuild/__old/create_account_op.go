package __old

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateAccountOp struct {
	Address     string
	RoleID      uint64
	Referrer    *string
	SignersData []xdr.UpdateSignerData
}

func (op CreateAccountOp) Validate() error {
	return validation.ValidateStruct(&op,
		// TODO validate address, recovery and referrer are addresses
		validation.Field(&op.Address, validation.Required),
		validation.Field(&op.RoleID, validation.Required),
		validation.Field(&op.Referrer, validation.NilOrNotEmpty),
		validation.Field(&op.SignersData, validation.Required),
	)
}

func (op CreateAccountOp) XDR() (*xdr.Operation, error) {
	var destination xdr.AccountId
	if err := destination.SetAddress(op.Address); err != nil {
		return nil, errors.Wrap(err, "failed to set destination")
	}

	xdrop := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateAccount,
			CreateAccountOp: &xdr.CreateAccountOp{
				Destination: destination,
				RoleId:      xdr.Uint64(op.RoleID),
			},
		},
	}

	if op.Referrer != nil {
		var referrer xdr.AccountId
		referrer.SetAddress(*op.Referrer)
		xdrop.Body.CreateAccountOp.Referrer = &referrer
	}

	return xdrop, nil
}
