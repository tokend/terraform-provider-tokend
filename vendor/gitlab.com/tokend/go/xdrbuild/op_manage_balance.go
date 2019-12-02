package xdrbuild

import (
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateBalance struct {
	Destination string
	AssetCode   string
}

func (op *CreateBalance) XDR() (*xdr.Operation, error) {
	var dest xdr.AccountId
	err := dest.SetAddress(op.Destination)
	if err != nil {
		return nil, errors.Wrap(err, "invalid destination")
	}
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateBalance,
			CreateBalanceOp: &xdr.CreateBalanceOp{
				Destination: dest,
				Asset:       xdr.AssetCode(op.AssetCode),
			},
		},
	}, nil
}
