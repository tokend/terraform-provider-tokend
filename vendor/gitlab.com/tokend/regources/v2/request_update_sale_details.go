package regources

// UpdateSaleDetailsRequest - represents details of the `update_sale_details` reviewable request
type UpdateSaleDetailsRequest struct {
	Key
	Attributes    UpdateSaleDetailsRequestAttrs     `json:"attributes"`
	Relationships UpdateSaleDetailsRequestRelations `json:"relationships"`
}

// UpdateSaleDetailsRequestAttrs - attributes of the `update_sale_details` reviewable request
type UpdateSaleDetailsRequestAttrs struct {
	CreatorDetails Details `json:"creator_details"`
}

// UpdateSaleDetailsRequestRelations - relationships of the `update_sale_details` reviewable request
type UpdateSaleDetailsRequestRelations struct {
	Sale *Relation `json:"sale"`
}
