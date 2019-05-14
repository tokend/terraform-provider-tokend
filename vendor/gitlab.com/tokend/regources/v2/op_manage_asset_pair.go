package regources

import "gitlab.com/tokend/go/xdr"

//ManageAssetPairOp - details of corresponding op
type ManageAssetPairOp struct {
	Key
	Attributes    ManageAssetPairOpAttrs     `json:"attributes"`
	Relationships ManageAssetPairOpRelations `json:"relationships"`
}

//ManageAssetPairOpAttrs - details of corresponding op
type ManageAssetPairOpAttrs struct {
	PhysicalPrice           Amount              `json:"physical_price"`
	PhysicalPriceCorrection Amount              `json:"physical_price_correction"`
	MaxPriceStep            Amount              `json:"max_price_step"`
	Policies                xdr.AssetPairPolicy `json:"policies"`
}

//ManageAssetPairRelations - relationships of the operation
type ManageAssetPairOpRelations struct {
	BaseAsset  *Relation `json:"base_asset"`
	QuoteAsset *Relation `json:"quote_asset"`
}
