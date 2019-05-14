package xdrbuild

import (
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type ManageBalanceOp struct {
	Action      xdr.ManageBalanceAction
	Destination string
	AssetCode   string
}

func (op *ManageBalanceOp) XDR() (*xdr.Operation, error) {
	var dest xdr.AccountId
	err := dest.SetAddress(op.Destination)
	if err != nil {
		return nil, errors.Wrap(err, "invalid destination")
	}
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageBalance,
			ManageBalanceOp: &xdr.ManageBalanceOp{
				Action:      op.Action,
				Destination: dest,
				Asset:       xdr.AssetCode(op.AssetCode),
			},
		},
	}, nil
}
