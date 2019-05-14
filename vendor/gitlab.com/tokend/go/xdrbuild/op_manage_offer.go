package xdrbuild

import (
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type RemoveOffer struct {
	OrderBookID uint64
	OfferID     uint64
}

func (op *RemoveOffer) XDR() (*xdr.Operation, error) {
	zeroBalance, err := xdr.NewPublicKey(xdr.CryptoKeyTypeKeyTypeEd25519, xdr.Uint256{})
	if err != nil {
		panic(errors.Wrap(err, "failed to create zero balanceID"))
	}
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageOffer,
			ManageOfferOp: &xdr.ManageOfferOp{
				OfferId:      xdr.Uint64(op.OfferID),
				OrderBookId:  xdr.Uint64(op.OrderBookID),
				Amount:       xdr.Int64(0),
				BaseBalance:  xdr.BalanceId(zeroBalance),
				QuoteBalance: xdr.BalanceId(zeroBalance),
			},
		},
	}, nil
}

type CreateOffer struct {
	OrderBookID  uint64
	BaseBalance  string
	QuoteBalance string
	IsBuy        bool
	Amount       int64
	Price        int64
	Fee          int64
}

func (op *CreateOffer) XDR() (*xdr.Operation, error) {
	manageOffer := xdr.ManageOfferOp{
		OfferId:     xdr.Uint64(0),
		OrderBookId: xdr.Uint64(op.OrderBookID),
		Amount:      xdr.Int64(op.Amount),
		Fee:         xdr.Int64(op.Fee),
		Price:       xdr.Int64(op.Price),
		IsBuy:       op.IsBuy,
	}

	err := manageOffer.BaseBalance.SetString(op.BaseBalance)
	if err != nil {
		return nil, errors.Wrap(err, "invalid base balance")
	}

	err = manageOffer.QuoteBalance.SetString(op.QuoteBalance)
	if err != nil {
		return nil, errors.Wrap(err, "invalid quote balance")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:          xdr.OperationTypeManageOffer,
			ManageOfferOp: &manageOffer,
		},
	}, nil
}
