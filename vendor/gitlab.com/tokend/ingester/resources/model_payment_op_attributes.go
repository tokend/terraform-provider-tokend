/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PaymentOpAttributes struct {
	Amount         Amount `json:"amount"`
	DestinationFee Fee    `json:"destination_fee"`
	// Reference for the payment
	Reference string `json:"reference"`
	// Security type of operation
	SecurityType uint32 `json:"security_type"`
	SourceFee    Fee    `json:"source_fee"`
	// Whether source of the payment should pay destination fee
	SourcePayForDestination bool `json:"source_pay_for_destination"`
	// Subject of the payment
	Subject string `json:"subject"`
}
