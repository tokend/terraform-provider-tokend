package regources

type OfferResponse struct {
	Data     Offer    `json:"data"`
	Included Included `json:"included"`
}

type OffersResponse struct {
	Links    *Links   `json:"links"`
	Data     []Offer  `json:"data"`
	Included Included `json:"included"`
}

type Offer struct {
	Key
	Attributes    OfferAttrs     `json:"attributes"`
	Relationships OfferRelations `json:"relationships"`
}

type OfferAttrs struct {
	IsBuy       bool   `json:"is_buy"`
	OrderBookID string `json:"order_book_id"`
	CreatedAt   string `json:"created_at"`
	BaseAmount  Amount `json:"base_amount"`
	QuoteAmount Amount `json:"quote_amount"`
	Price       Amount `json:"price"`
	Fee         Fee    `json:"fee"`
}

type OfferRelations struct {
	Owner        *Relation `json:"owner"`
	BaseAsset    *Relation `json:"base_asset"`
	BaseBalance  *Relation `json:"base_balance"`
	QuoteAsset   *Relation `json:"quote_asset"`
	QuoteBalance *Relation `json:"quote_balance"`
}
