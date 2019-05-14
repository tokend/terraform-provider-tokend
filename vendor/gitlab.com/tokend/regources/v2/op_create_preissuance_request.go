package regources

//CreatePreIssuanceRequestOp - details of corresponding op
type CreatePreIssuanceRequestOp struct {
	Key
	Attributes    CreatePreIssuanceRequestOpAttrs     `json:"attributes"`
	Relationships CreatePreIssuanceRequestOpRelations `json:"relationships"`
}

//CreatePreIssuanceRequestOpAttrs - details of corresponding op
type CreatePreIssuanceRequestOpAttrs struct {
	Amount         Amount  `json:"amount"`
	CreatorDetails Details `json:"creator_details"`
}

type CreatePreIssuanceRequestOpRelations struct {
	Asset   *Relation `json:"asset"`
	Request *Relation `json:"request"`
}
