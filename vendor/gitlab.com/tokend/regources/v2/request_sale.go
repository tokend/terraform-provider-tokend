package regources

import (
	"gitlab.com/tokend/go/xdr"
	"time"
)

// CreateSaleRequest - represents details of the `sale` reviewable request
type CreateSaleRequest struct {
	Key
	Attributes    CreateSaleRequestAttrs     `json:"attributes"`
	Relationships CreateSaleRequestRelations `json:"relationships"`
}

// CreateSaleRequestAttrs - attributes of the `sale` reviewable request
type CreateSaleRequestAttrs struct {
	BaseAssetForHardCap Amount       `json:"base_asset_for_hard_cap"`
	StartTime           time.Time    `json:"start_time"`
	EndTime             time.Time    `json:"end_time"`
	SaleType            xdr.SaleType `json:"sale_type"`
	CreatorDetails      Details      `json:"creator_details"`
}

// CreateSaleRequestRelations - attributes of the `sale` reviewable request
type CreateSaleRequestRelations struct {
	BaseAsset         *Relation           `json:"base_asset"`
	QuoteAssets       *RelationCollection `json:"quote_assets"`
	DefaultQuoteAsset *Relation           `json:"default_quote_asset"`
}
