package regources

//CancelAtomicSwapBidOp - details of corresponding op
type CancelAtomicSwapBidOp struct {
	Key
	Relationships CancelAtomicSwapBidOpRelations `json:"relationships"`
}

// CancelAtomicSwapBidOpRelations - relations of operation
type CancelAtomicSwapBidOpRelations struct {
	Bid *Relation `json:"bid"`
}
