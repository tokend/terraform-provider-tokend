package regources

//ManageOfferOp - details of corresponding op
type ManageOfferOp struct {
	Key
	Attributes    ManageOfferOpAttrs     `json:"attributes"`
	Relationships ManageOfferOpRelations `json:"relationships"`
}

//ManageOfferOpAttrs - details of corresponding op
type ManageOfferOpAttrs struct {
	OfferID     int64  `json:"offer_id,omitempty"`
	OrderBookID int64  `json:"order_book_id"`
	BaseAmount  Amount `json:"base_amount"`
	Price       Amount `json:"price"`
	IsBuy       bool   `json:"is_buy"`
	Fee         Fee    `json:"fee"`
	IsDeleted   bool   `json:"is_deleted"`
}

type ManageOfferOpRelations struct {
	BaseAsset  *Relation `json:"base_asset"`
	QuoteAsset *Relation `json:"quote_asset"`
}
