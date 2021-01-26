package __old

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type CreateAtomicSwapBidCreationRequestOp struct {
	BaseBalance string
	BaseAmount  uint64
	Details     string
	QuoteAssets []QuoteAsset
}

type QuoteAsset struct {
	Price uint64
	Asset string
}

func (op CreateAtomicSwapBidCreationRequestOp) Validate() error {
	for _, quoteAsset := range op.QuoteAssets {
		if quoteAsset.Price == 0 {
			return errors.New("quote asset price cannot be zero")
		}

		if quoteAsset.Asset == "" {
			return errors.New("quote asset code cannot be empty")
		}
	}
	return validation.ValidateStruct(&op,
		validation.Field(&op.BaseBalance, validation.Required),
		validation.Field(&op.BaseAmount, validation.Required),
		validation.Field(&op.Details, validation.Required),
		validation.Field(&op.QuoteAssets, validation.Required),
	)
}

func (op CreateAtomicSwapBidCreationRequestOp) XDR() (*xdr.Operation, error) {
	var baseBalance xdr.BalanceId
	if err := baseBalance.SetString(op.BaseBalance); err != nil {
		return nil, errors.Wrap(err, "failed to set base balance")
	}

	var quoteAssets []xdr.ASwapBidQuoteAsset
	for _, quoteAsset := range op.QuoteAssets {
		quoteAssets = append(quoteAssets, xdr.ASwapBidQuoteAsset{
			QuoteAsset: xdr.AssetCode(quoteAsset.Asset),
			Price:      xdr.Uint64(quoteAsset.Price),
			Ext: xdr.ASwapBidQuoteAssetExt{
				V: xdr.LedgerVersionEmptyVersion,
			},
		})
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateAswapBidRequest,
			CreateASwapBidCreationRequestOp: &xdr.CreateASwapBidCreationRequestOp{
				Request: xdr.ASwapBidCreationRequest{
					BaseBalance:    baseBalance,
					Amount:         xdr.Uint64(op.BaseAmount),
					CreatorDetails: xdr.Longstring(op.Details),
					QuoteAssets:    quoteAssets,
				},
			},
		},
	}, nil
}
