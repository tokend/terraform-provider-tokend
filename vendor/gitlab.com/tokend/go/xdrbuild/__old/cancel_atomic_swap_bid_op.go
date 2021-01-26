package __old

import (
	"github.com/go-ozzo/ozzo-validation"
	"gitlab.com/tokend/go/xdr"
)

type CancelAtomicSwapBidOp struct {
	BidID uint64
}

func (op CancelAtomicSwapBidOp) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.BidID, validation.Required),
	)
}

func (op CancelAtomicSwapBidOp) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCancelAswapBid,
			CancelASwapBidOp: &xdr.CancelASwapBidOp{
				BidId: xdr.Uint64(op.BidID),
			},
		},
	}, nil
}
