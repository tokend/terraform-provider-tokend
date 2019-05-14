package regources

//CreateManageLimitsRequestOp - details of corresponding op
type CreateManageLimitsRequestOp struct {
	Key
	Attributes    CreateManageLimitsRequestOpAttrs     `json:"attributes"`
	Relationships CreateManageLimitsRequestOpRelations `json:"relationships"`
}

//CreateManageLimitsRequestOpAttrs - details of corresponding op
type CreateManageLimitsRequestOpAttrs struct {
	CreatorDetails Details `json:"creator_details"`
}

//CreateManageLimitsRequestOpRelations - relationships of the operation
type CreateManageLimitsRequestOpRelations struct {
	Request *Relation `json:"request"`
}
