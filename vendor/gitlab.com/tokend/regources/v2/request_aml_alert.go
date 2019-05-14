package regources

// CreateAmlAlertRequest - represents details of the `aml alert` reviewable request
type CreateAmlAlertRequest struct {
	Key
	Attributes    CreateAmlAlertRequestAttrs     `json:"attributes"`
	Relationships CreateAmlAlertRequestRelations `json:"relationships"`
}

// CreateAmlAlertRequestAttrs - attributes of the `aml alert` reviewable request
type CreateAmlAlertRequestAttrs struct {
	Amount         Amount  `json:"amount"`
	CreatorDetails Details `json:"creator_details"`
}

// CreateAmlAlertRequestRelations - relationships of the `aml alert` reviewable request
type CreateAmlAlertRequestRelations struct {
	Balance *Relation `json:"balance"`
}
