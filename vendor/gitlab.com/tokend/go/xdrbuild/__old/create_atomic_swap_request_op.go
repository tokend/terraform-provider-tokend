package __old

import (
	"github.com/go-ozzo/ozzo-validation"
	"gitlab.com/tokend/go/xdr"
)

type CreateAtomicSwapRequestOp struct {
	BidID      uint64
	BaseAmount uint64
	QuoteAsset string
}

func (op CreateAtomicSwapRequestOp) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.BidID, validation.Required),
		validation.Field(&op.BaseAmount, validation.Required),
		validation.Field(&op.QuoteAsset, validation.Required),
	)
}

func (op CreateAtomicSwapRequestOp) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateAswapRequest,
			CreateASwapRequestOp: &xdr.CreateASwapRequestOp{
				Request: xdr.ASwapRequest{
					BidId:      xdr.Uint64(op.BidID),
					BaseAmount: xdr.Uint64(op.BaseAmount),
					QuoteAsset: xdr.AssetCode(op.QuoteAsset),
				},
			},
		},
	}, nil
}
