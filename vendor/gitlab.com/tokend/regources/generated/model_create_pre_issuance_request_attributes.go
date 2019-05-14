/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreatePreIssuanceRequestAttributes struct {
	// Amount to be pre issued
	Amount         Amount  `json:"amount"`
	CreatorDetails Details `json:"creator_details"`
	// Reference on the pre issuance request (since it cannot be submitted more than once)
	Reference string `json:"reference"`
	// Signature of the pre issuance signer
	Signature string `json:"signature"`
}
