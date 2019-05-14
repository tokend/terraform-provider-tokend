package regources

// CreateAtomicSwapRequest - represents details of the `atomic swap` reviewable request
type CreateAtomicSwapRequest struct {
	Key
	Attributes    CreateAtomicSwapRequestAttrs     `json:"attributes"`
	Relationships CreateAtomicSwapRequestRelations `json:"relationships"`
}

// CreateAtomicSwapRequestAttrs - attributes of the `atomic swap` reviewable request
type CreateAtomicSwapRequestAttrs struct {
	BaseAmount     Amount  `json:"base_amount"`
	CreatorDetails Details `json:"creator_details"`
}

// CreateAtomicSwapRequestRelations - relationships of the `atomic swap` reviewable request
type CreateAtomicSwapRequestRelations struct {
	Bid        *Relation `json:"bid"`
	QuoteAsset *Relation `json:"quote_asset"`
}
