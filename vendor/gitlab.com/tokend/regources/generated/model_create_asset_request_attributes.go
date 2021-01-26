/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateAssetRequestAttributes struct {
	// Unique asset identifier
	Asset          string  `json:"asset"`
	CreatorDetails Details `json:"creator_details"`
	// Amount to be issued automatically right after the asset creation
	InitialPreissuedAmount Amount `json:"initial_preissued_amount"`
	// Maximal amount to be issued
	MaxIssuanceAmount Amount `json:"max_issuance_amount"`
	// Policies specified for the asset creation
	Policies int32 `json:"policies"`
	// Address of an account that performs pre issuance
	PreIssuanceAssetSigner string `json:"pre_issuance_asset_signer"`
	// Number of digits after point (comma)
	TrailingDigitsCount uint32 `json:"trailing_digits_count"`
	// Numeric type of asset
	Type uint64 `json:"type"`
}
