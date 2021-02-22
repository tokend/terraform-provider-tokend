package xdrbuild

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type SaleQuoteAsset struct {
	Code  string
	Price uint64
}

type CreateSaleCreationRequest struct {
	CreatorAccountID  string
	SoftCap           uint64
	HardCap           uint64
	StartTime         time.Time
	EndTime           time.Time
	BaseAsset         string
	DefaultQuoteAsset string
	SaleType          uint64
	QuoteAssets       []SaleQuoteAsset
	Details           json.Marshaler
	AllTasks          *uint32
	SaleAmount        uint64
}

func (op CreateSaleCreationRequest) XDR() (*xdr.Operation, error) {
	var creator xdr.AccountId
	if err := creator.SetAddress(op.CreatorAccountID); err != nil {
		return nil, errors.Wrap(err, "failed to set creator account address")
	}

	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details to json")
	}

	quoteAssets := make([]xdr.SaleCreationRequestQuoteAsset, 0, len(op.QuoteAssets))
	for _, qa := range op.QuoteAssets {
		quoteAssets = append(quoteAssets, xdr.SaleCreationRequestQuoteAsset{
			QuoteAsset: xdr.AssetCode(qa.Code),
			Price:      xdr.Uint64(qa.Price),
		})
	}
	return &xdr.Operation{
		SourceAccount: &creator,
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateSaleRequest,
			CreateSaleCreationRequestOp: &xdr.CreateSaleCreationRequestOp{
				Request: xdr.SaleCreationRequest{
					SaleType:                    xdr.Uint64(op.SaleType),
					BaseAsset:                   xdr.AssetCode(op.BaseAsset),
					DefaultQuoteAsset:           xdr.AssetCode(op.DefaultQuoteAsset),
					StartTime:                   xdr.Uint64(op.StartTime.Unix()),
					EndTime:                     xdr.Uint64(op.EndTime.Unix()),
					SoftCap:                     xdr.Uint64(op.SoftCap),
					HardCap:                     xdr.Uint64(op.HardCap),
					CreatorDetails:              xdr.Longstring(details),
					RequiredBaseAssetForHardCap: xdr.Uint64(op.SaleAmount),
					QuoteAssets:                 quoteAssets,
				},
				AllTasks: (*xdr.Uint32)(op.AllTasks),
			},
		},
	}, nil
}
