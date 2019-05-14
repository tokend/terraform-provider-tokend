package regources

//QuoteAsset - represents resource of ReviewableRequest which defines in which asset it's allowed to invest into sale
type QuoteAsset struct {
	Key
	Attributes QuoteAssetAttrs `json:"attributes"`
}

//QuoteAssetAttrs - attributes of QuoteAsset resource
type QuoteAssetAttrs struct {
	Price Amount `json:"price"`
}

//SaleQuoteAsset - represents resource of SaleQuoteAsset which defines in which asset it's allowed to invest into sale
// and current performance of the sale in this asset
type SaleQuoteAsset struct {
	Key
	Attributes    SaleQuoteAssetAttrs     `json:"attributes"`
	Relationships SaleQuoteAssetRelations `json:"relationships"`
}

//SaleQuoteAssetAttrs - attributes of SaleQuoteAsset resource
type SaleQuoteAssetAttrs struct {
	Price           Amount `json:"price"`
	CurrentCap      Amount `json:"current_cap"`
	TotalCurrentCap Amount `json:"total_current_cap"`
	HardCap         Amount `json:"hard_cap"`
	SoftCap         Amount `json:"soft_cap,omitempty"`
}

//SaleQuoteAssetRelations - represents relationships of SaleQuoteAsset resource
type SaleQuoteAssetRelations struct {
	Asset *Relation `json:"asset"`
}
