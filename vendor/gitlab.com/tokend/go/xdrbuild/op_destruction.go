package xdrbuild

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type Destruction struct {
	SecurityType uint32
	BalanceID    xdr.BalanceId
	Amount       uint64
	Reference    string
	Details      json.Marshaler
	Source       string
}

func (op *Destruction) XDR() (*xdr.Operation, error) {

	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	var source *xdr.AccountId
	if op.Source != "" {
		source = new(xdr.AccountId)
		err = source.SetAddress(op.Source)
		if err != nil {
			return nil, errors.Wrap(err, "failed to set source account id")
		}
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeDestruction,
			DestructionOp: &xdr.DestructionOp{
				Fee:            xdr.Fee{},
				SecurityType:   xdr.Uint32(op.SecurityType),
				Balance:        op.BalanceID,
				Amount:         xdr.Uint64(op.Amount),
				CreatorDetails: xdr.Longstring(details),
				Ext:            xdr.EmptyExt{},
			},
		},
		SourceAccount: source,
	}, nil
}
