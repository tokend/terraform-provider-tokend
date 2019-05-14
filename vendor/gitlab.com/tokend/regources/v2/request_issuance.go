package regources

// CreateIssuanceRequest - represents details of the `issuance` reviewable request
type CreateIssuanceRequest struct {
	Key
	Attributes    CreateIssuanceRequestAttrs     `json:"attributes"`
	Relationships CreateIssuanceRequestRelations `json:"relationships"`
}

// CreateIssuanceRequestAttrs - attributes of the `issuance` reviewable request
type CreateIssuanceRequestAttrs struct {
	Amount         Amount  `json:"amount"`
	CreatorDetails Details `json:"creator_details"`
}

// CreateIssuanceRequestRelations - relationships of the `issuance` reviewable request
type CreateIssuanceRequestRelations struct {
	Asset    *Relation `json:"asset"`
	Receiver *Relation `json:"receiver"`
}
