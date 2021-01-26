package __old

import (
	. "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type ManageBalanceOp struct {
	Action      xdr.ManageBalanceAction
	Destination string
	Asset       string
}

func (mb ManageBalanceOp) Validate() error {
	return ValidateStruct(&mb,
		Field(&mb.Action, Required),
		Field(&mb.Asset, Required),
		Field(&mb.Destination, Required),
	)
}

func (mb ManageBalanceOp) XDR() (*xdr.Operation, error) {
	var destination xdr.AccountId
	if err := destination.SetAddress(mb.Destination); err != nil {
		return nil, errors.Wrap(err, "failed to set destination")
	}
	op := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageBalance,
			ManageBalanceOp: &xdr.ManageBalanceOp{
				Action:      mb.Action,
				Asset:       xdr.AssetCode(mb.Asset),
				Destination: destination,
				Ext:         xdr.ManageBalanceOpExt{V: xdr.LedgerVersionEmptyVersion},
			},
		},
	}
	return op, nil
}
