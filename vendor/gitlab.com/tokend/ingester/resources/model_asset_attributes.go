/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AssetAttributes struct {
	Details Details `json:"details"`
	// Asset volume that is currently in circulation
	Issued Amount `json:"issued"`
	// Max volume of an asset that can be in circulation
	MaxIssuanceAmount Amount `json:"max_issuance_amount"`
	// Asset volume to be distributed via [asset sale↪](https://tokend.gitbook.io/knowledge-base/platform-features/crowdfunding) but currently locked by the system
	PendingIssuance Amount `json:"pending_issuance"`
	// Security type of an asset
	SecurityType uint32 `json:"security_type"`
	// State of an asset
	State uint32 `json:"state"`
	// Number of significant digits after the point
	TrailingDigits uint32 `json:"trailing_digits"`
}
