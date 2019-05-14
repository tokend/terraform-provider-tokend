package regources

// CreateAtomicSwapBidRequest - represents details of the `atomic swap bid` reviewable request
type CreateAtomicSwapBidRequest struct {
	Key
	Attributes    CreateAtomicSwapBidRequestAttrs     `json:"attributes"`
	Relationships CreateAtomicSwapBidRequestRelations `json:"relationships"`
}

// CreateAtomicSwapBidRequestAttrs - attributes of the `atomic swap bid` reviewable request
type CreateAtomicSwapBidRequestAttrs struct {
	BaseAmount     Amount  `json:"base_amount"`
	CreatorDetails Details `json:"creator_details"`
}

// AtomicSwapBidRequestRelations - relationships of the `atomic swap bid` reviewable request
type CreateAtomicSwapBidRequestRelations struct {
	BaseBalance *Relation           `json:"base_balance"`
	QuoteAssets *RelationCollection `json:"quote_assets"`
}
