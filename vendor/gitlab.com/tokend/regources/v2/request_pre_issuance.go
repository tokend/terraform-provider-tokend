package regources

// CreatePreIssuanceRequest - represents details of the `pre-issuance` reviewable request
type CreatePreIssuanceRequest struct {
	Key
	Attributes    CreatePreIssuanceRequestAttrs     `json:"attributes"`
	Relationships CreatePreIssuanceRequestRelations `json:"relationships"`
}

// CreatePreIssuanceRequestAttrs - attributes of the `pre_issuance` reviewable request
type CreatePreIssuanceRequestAttrs struct {
	Amount         Amount  `json:"amount"`
	Signature      string  `json:"signature"`
	Reference      string  `json:"reference"`
	CreatorDetails Details `json:"creator_details"`
}

// CreatePreIssuanceRequestRelations - relationships of the `pre_issuance` reviewable request
type CreatePreIssuanceRequestRelations struct {
	Asset *Relation `json:"asset"`
}
