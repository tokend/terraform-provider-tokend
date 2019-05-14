package regources

//CreateWithdrawRequestOp - details of corresponding op
type CreateWithdrawRequestOp struct {
	Key
	Attributes    CreateWithdrawRequestOpAttrs     `json:"attributes"`
	Relationships CreateWithdrawRequestOpRelations `json:"relationships"`
}

//CreateWithdrawRequestOpAttrs - details of corresponding op
type CreateWithdrawRequestOpAttrs struct {
	Amount         Amount  `json:"amount"`
	Fee            Fee     `json:"fee"`
	CreatorDetails Details `json:"creator_details"`
}

//CreateWithdrawRequestOpRelations - relationships of the operation
type CreateWithdrawRequestOpRelations struct {
	Balance *Relation `json:"balance"`
}
