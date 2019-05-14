package regources

//CreateIssuanceRequestAttrs - details of corresponding op
type CreateIssuanceRequestOp struct {
	Key
	Attributes    CreateIssuanceRequestOpAttrs     `json:"attributes"`
	Relationships CreateIssuanceRequestOpRelations `json:"relationships"`
}

//CreateIssuanceRequestOpAttrs - details of corresponding op
type CreateIssuanceRequestOpAttrs struct {
	Fee            Fee     `json:"fee"`
	Amount         Amount  `json:"amount"`
	Reference      string  `json:"reference"`
	AllTasks       *int64  `json:"all_tasks,omitempty"`
	CreatorDetails Details `json:"creator_details"`
}

// CreateIssuanceRequestOpRelations - relationships of the operation
type CreateIssuanceRequestOpRelations struct {
	Asset           *Relation `json:"asset"`
	Request         *Relation `json:"request"`
	ReceiverAccount *Relation `json:"receiver_account"`
	ReceiverBalance *Relation `json:"receiver_balance"`
}
