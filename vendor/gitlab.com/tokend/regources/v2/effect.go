package regources

//EffectBalanceChange - base effect used for balance change
type EffectBalanceChange struct {
	Key
	Attributes EffectBalanceChangeAttrs `json:"attributes"`
}

//EffectBalanceChangeAttrs - attributes for balance change effect
type EffectBalanceChangeAttrs struct {
	Amount Amount `json:"amount"`
	Fee    Fee    `json:"fee"`
}

//EffectMatched - effect for orders match
type EffectMatched struct {
	Key
	Attributes EffectMatchedAttrs `json:"attributes"`
}

//EffectMatchedAttrs - attributes of matched effect
type EffectMatchedAttrs struct {
	OfferID     int64                         `json:"offer_id"`
	OrderBookID int64                         `json:"order_book_id"`
	Price       Amount                        `json:"price"`
	Charged     ParticularBalanceChangeEffect `json:"charged"`
	Funded      ParticularBalanceChangeEffect `json:"funded"`
}

// ParticularBalanceChangeEffect - describes movement of fund for particular balance
type ParticularBalanceChangeEffect struct {
	EffectBalanceChangeAttrs
	BalanceAddress string `json:"balance_address"`
	AssetCode      string `json:"asset_code"`
}
