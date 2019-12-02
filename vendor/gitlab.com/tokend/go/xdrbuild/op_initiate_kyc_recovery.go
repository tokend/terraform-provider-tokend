package xdrbuild

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type InitiateKycRecovery struct {
	Account string
	Signer  string
}

func (op *InitiateKycRecovery) XDR() (*xdr.Operation, error) {
	initOp := xdr.InitiateKycRecoveryOp{}
	err := initOp.Account.SetAddress(op.Account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set account address")
	}
	err = initOp.Signer.FromString(op.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set signer address")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:                  xdr.OperationTypeInitiateKycRecovery,
			InitiateKycRecoveryOp: &initOp,
		},
	}, nil
}
