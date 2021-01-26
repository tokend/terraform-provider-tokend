package __old

import (
	"github.com/go-ozzo/ozzo-validation"
	"gitlab.com/tokend/go/xdr"
)

type CreateSaleRequestOp struct {
	RequestID         uint64
	BaseAsset         string
	DefaultQuoteAsset string
	StartTime         uint64
	EndTime           uint64
	SoftCap           uint64
	HardCap           uint64
	Details           string
	QuoteAssets       []xdr.SaleCreationRequestQuoteAsset
}

func (op CreateSaleRequestOp) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.BaseAsset, validation.Required),
		validation.Field(&op.DefaultQuoteAsset, validation.Required),
		validation.Field(&op.StartTime, validation.Required),
		validation.Field(&op.EndTime, validation.Required),
		validation.Field(&op.SoftCap, validation.Required),
		validation.Field(&op.HardCap, validation.Required),
		validation.Field(&op.QuoteAssets, validation.Required),
	)
}

func (op CreateSaleRequestOp) XDR() (*xdr.Operation, error) {
	xdrOp := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateSaleRequest,
			CreateSaleCreationRequestOp: &xdr.CreateSaleCreationRequestOp{
				RequestId: xdr.Uint64(op.RequestID),
				Request: xdr.SaleCreationRequest{
					BaseAsset:         xdr.AssetCode(op.BaseAsset),
					DefaultQuoteAsset: xdr.AssetCode(op.DefaultQuoteAsset),
					StartTime:         xdr.Uint64(op.StartTime),
					EndTime:           xdr.Uint64(op.EndTime),
					SoftCap:           xdr.Uint64(op.SoftCap),
					HardCap:           xdr.Uint64(op.HardCap),
					CreatorDetails:    xdr.Longstring(op.Details),
					QuoteAssets:       op.QuoteAssets,
				},
			},
		},
	}

	return xdrOp, nil
}
