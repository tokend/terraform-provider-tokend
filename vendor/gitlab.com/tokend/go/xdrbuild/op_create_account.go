package xdrbuild

import (
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateAccount struct {
	Destination string
	Referrer    *string
	RoleID      uint64
	Signers     []SignerData
}

func (op *CreateAccount) XDR() (*xdr.Operation, error) {
	var destination xdr.AccountId
	if err := destination.SetAddress(op.Destination); err != nil {
		return nil, errors.Wrap(err, "failed to set destination address")
	}
	var referrer *xdr.AccountId
	if op.Referrer != nil {
		referrer = &xdr.AccountId{}
		if err := referrer.SetAddress(*op.Referrer); err != nil {
			return nil, errors.Wrap(err, "failed to set referrer address")
		}
	}
	signers := make([]xdr.UpdateSignerData, 0, len(op.Signers))
	for _, signer := range op.Signers {
		data, err := signer.XDR()
		if err != nil {
			return nil, errors.Wrap(err, "failed to build signer data")
		}
		signers = append(signers, *data)
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateAccount,
			CreateAccountOp: &xdr.CreateAccountOp{
				Destination: destination,
				Referrer:    referrer,
				RoleId:      xdr.Uint64(op.RoleID),
				SignersData: signers,
			},
		},
	}, nil
}
