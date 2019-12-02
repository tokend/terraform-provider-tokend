package xdrbuild

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type KycRecovery struct {
	Details json.Marshaler
	Account string
	Signers []SignerData
}

func (op *KycRecovery) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal json details")
	}

	recoverOp := xdr.KycRecoveryOp{
		CreatorDetails: xdr.Longstring(details),
	}

	err = recoverOp.TargetAccount.SetAddress(op.Account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set address")
	}

	recoverOp.SignersData = make([]xdr.SignerData, 0, len(op.Signers))
	for _, signer := range op.Signers {
		data, err := signer.XDR()
		if err != nil {
			return nil, errors.Wrap(err, "failed to build signer data")
		}
		recoverOp.SignersData = append(recoverOp.SignersData, *data)
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:          xdr.OperationTypeKycRecovery,
			KycRecoveryOp: &recoverOp,
		},
	}, nil
}
