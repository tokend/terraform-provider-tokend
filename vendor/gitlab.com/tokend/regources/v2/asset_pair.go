package regources

import "gitlab.com/tokend/go/xdr"

type AssetPairResponse struct {
	Data     AssetPair `json:"data"`
	Included Included  `json:"included"`
}

type AssetPairsResponse struct {
	Links    *Links      `json:"links"`
	Data     []AssetPair `json:"data"`
	Included Included    `json:"included"`
}

type AssetPair struct {
	Key
	Attributes    AssetPairAttrs     `json:"attributes"`
	Relationships AssetPairRelations `json:"relationships"`
}

type AssetPairAttrs struct {
	Price    Amount              `json:"price"`
	Policies xdr.AssetPairPolicy `json:"policies"`
}

type AssetPairRelations struct {
	BaseAsset  *Relation `json:"base_asset"`
	QuoteAsset *Relation `json:"quote_asset"`
}
