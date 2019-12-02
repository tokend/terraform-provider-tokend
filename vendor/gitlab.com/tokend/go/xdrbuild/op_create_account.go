package xdrbuild

import (
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateAccount struct {
	Destination string
	Referrer    *string
	RoleIDs      []uint64
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
	signers := make([]xdr.SignerData, 0, len(op.Signers))
	for _, signer := range op.Signers {
		data, err := signer.XDR()
		if err != nil {
			return nil, errors.Wrap(err, "failed to build signer data")
		}
		signers = append(signers, *data)
	}

	roleIDs := make([]xdr.Uint64, 0, len(op.RoleIDs))
	for _, roleID := range op.RoleIDs {
		roleIDs = append(roleIDs, xdr.Uint64(roleID))
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateAccount,
			CreateAccountOp: &xdr.CreateAccountOp{
				Destination: destination,
				Referrer:    referrer,
				RoleIDs:     roleIDs,
				SignersData: signers,
			},
		},
	}, nil
}
