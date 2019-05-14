package regources

//CreateAccountOp - stores details of create account operation
type CreateAccountOp struct {
	Key
	Relationships CreateAccountOpRelation `json:"relationships"`
}

// CreateAccountOpRelation - stores details of create account relations
type CreateAccountOpRelation struct {
	Account *Relation `json:"account"`
	Role    *Relation `json:"role"`
}
