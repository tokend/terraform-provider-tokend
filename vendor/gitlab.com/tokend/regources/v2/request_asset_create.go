package regources

// AssetCreationRequest - represents details of the `asset create` reviewable request
type CreateAssetRequest struct {
	Key
	Attributes CreateAssetRequestAttrs `json:"attributes"`
}

// AssetCreationRequestAttrs - attributes of the `asset create` reviewable request
type CreateAssetRequestAttrs struct {
	Asset                  string  `json:"asset"`
	Type                   uint64  `json:"type"`
	Policies               int32   `json:"policies"`
	PreIssuanceAssetSigner string  `json:"pre_issuance_asset_signer"`
	MaxIssuanceAmount      Amount  `json:"max_issuance_amount"`
	InitialPreissuedAmount Amount  `json:"initial_preissued_amount"`
	CreatorDetails         Details `json:"creator_details"`
}
