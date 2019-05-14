package regources

//CreateAtomicSwapBidRequestOp - details of corresponding op
type CreateAtomicSwapBidRequestOp struct {
	Key
	Attributes    CreateAtomicSwapBidRequestOpAttrs     `json:"attributes"`
	Relationships CreateAtomicSwapBidRequestOpRelations `json:"relationships"`
}

//CreateAtomicSwapBidRequestOpAttrs - details of corresponding op
type CreateAtomicSwapBidRequestOpAttrs struct {
	Amount  Amount  `json:"amount"`
	Details Details `json:"details"`
}

type CreateAtomicSwapBidRequestOpRelations struct {
	Request     *Relation           `json:"request"`
	BaseBalance *Relation           `json:"base_balance"`
	QuoteAssets *RelationCollection `json:"quote_assets"`
}
