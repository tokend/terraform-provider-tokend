package xdrbuild

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type InitiateKYCRecovery struct {
	AccountID string
	Signer    string
}

func (op *InitiateKYCRecovery) XDR() (*xdr.Operation, error) {
	var account xdr.AccountId
	err := account.SetAddress(op.AccountID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set account")
	}

	var signer xdr.AccountId
	err = signer.SetAddress(op.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set signer")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeInitiateKycRecovery,
			InitiateKycRecoveryOp: &xdr.InitiateKycRecoveryOp{
				Account: account,
				Signer:  xdr.PublicKey(signer),
			},
		},
	}, nil
}
