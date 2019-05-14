package regources

type HistoryOffer struct {
	PT                string `json:"paging_token"`
	OfferID           string `json:"offer_id,omitempty"`
	OwnerID           string `json:"owner_id"`
	BaseAsset         string `json:"base_asset"`
	QuoteAsset        string `json:"quote_asset"`
	IsBuy             bool   `json:"is_buy"`
	InitialBaseAmount Amount `json:"initial_base_amount"`
	CurrentBaseAmount Amount `json:"current_base_amount"`
	Price             Amount `json:"price"`
	IsCanceled        bool   `json:"is_canceled"`
	CreatedAt         Time   `json:"created_at"`
}

func (o HistoryOffer) PagingToken() string {
	return o.PT
}
