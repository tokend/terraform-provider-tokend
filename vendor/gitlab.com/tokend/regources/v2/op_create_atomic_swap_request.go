package regources

//CreateAtomicSwapRequestOp - details of corresponding op
type CreateAtomicSwapRequestOp struct {
	Key
	Attributes    CreateAtomicSwapRequestOpAttrs     `json:"attributes"`
	Relationships CreateAtomicSwapRequestOpRelations `json:"relationships"`
}

//CreateAtomicSwapRequestOpAttrs - details of corresponding op
type CreateAtomicSwapRequestOpAttrs struct {
	BaseAmount Amount `json:"base_amount"`
}

//CreateAtomicSwapRequestOpRelations - relationships of the operation
type CreateAtomicSwapRequestOpRelations struct {
	Bid        *Relation `json:"bid"`
	Request    *Relation `json:"request"`
	QuoteAsset *Relation `json:"quote_asset"`
}
