package regources

type Asset struct {
	Code                 string       `json:"code"`
	Owner                string       `json:"owner"`
	AvailableForIssuance Amount       `json:"available_for_issuance"`
	Details              AssetDetails `json:"details"`
	Issued               Amount       `json:"issued"`
	Type                 uint64       `json:"type"`
}
