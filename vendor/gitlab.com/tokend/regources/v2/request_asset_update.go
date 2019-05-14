package regources

// UpdateAssetRequest - represents details of the `asset update` reviewable request
type UpdateAssetRequest struct {
	Key
	Attributes    AssetUpdateRequestAttrs     `json:"attributes"`
	Relationships UpdateAssetRequestRelations `json:"relationships"`
}

// AssetUpdateRequestAttrs - attributes of the `asset update` reviewable request
type AssetUpdateRequestAttrs struct {
	Policies       int32   `json:"policies"`
	CreatorDetails Details `json:"creator_details"`
}

// UpdateAssetRequestRelations - attributes of the `asset update` reviewable request
type UpdateAssetRequestRelations struct {
	Asset *Relation `json:"asset"`
}
