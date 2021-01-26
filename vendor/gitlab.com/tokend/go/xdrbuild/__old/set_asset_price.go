package __old

import (
	"gitlab.com/tokend/go/xdr"
)

type SetAssetPrice struct {
	BaseAsset  string
	QuoteAsset string
	Price      int64
}

func (op SetAssetPrice) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageAssetPair,
			ManageAssetPairOp: &xdr.ManageAssetPairOp{
				Action:        xdr.ManageAssetPairActionUpdatePrice,
				Base:          xdr.AssetCode(op.BaseAsset),
				Quote:         xdr.AssetCode(op.QuoteAsset),
				PhysicalPrice: xdr.Int64(op.Price),
			},
		},
	}, nil
}
