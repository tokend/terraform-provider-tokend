package regources

import "time"

//CreateSaleRequestOp -details of corresponding op
type CreateSaleRequestOp struct {
	Key
	Attributes    CreateSaleRequestOpAttrs     `json:"attributes"`
	Relationships CreateSaleRequestOpRelations `json:"relationships"`
}

//CreateSaleRequestOpAttrs -details of corresponding op
type CreateSaleRequestOpAttrs struct {
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	SoftCap        Amount    `json:"soft_cap"`
	HardCap        Amount    `json:"hard_cap"`
	CreatorDetails Details   `json:"creator_details"`
}

// CreateSaleRequestOpAttrs - relations of corresponding op
type CreateSaleRequestOpRelations struct {
	Request           *Relation           `json:"request"`
	QuoteAssets       *RelationCollection `json:"quote_assets"`
	BaseAsset         *Relation           `json:"base_asset"`
	DefaultQuoteAsset *Relation           `json:"default_quote_asset"`
}
