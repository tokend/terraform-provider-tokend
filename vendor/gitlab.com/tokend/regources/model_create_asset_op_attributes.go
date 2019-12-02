/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateAssetOpAttributes struct {
	// Unique asset identifier
	Asset          string  `json:"asset"`
	CreatorDetails Details `json:"creator_details"`
	// Maximal amount to be issued
	MaxIssuanceAmount Amount `json:"max_issuance_amount"`
	// Security type of asset
	SecurityType uint32 `json:"security_type"`
	// Number of digits after point (comma)
	TrailingDigitsCount uint32 `json:"trailing_digits_count"`
}
