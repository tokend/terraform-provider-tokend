/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type IssuanceOpAttributes struct {
	Amount         Amount  `json:"amount"`
	CreatorDetails Details `json:"creator_details"`
	Fee            Fee     `json:"fee"`
	// reference of the request
	Reference string `json:"reference"`
	// Security type of operation
	SecurityType uint32 `json:"security_type"`
}
