/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type DeferredPaymentAttributes struct {
	Amount  Amount  `json:"amount"`
	Details Details `json:"details"`
	// String representation of the payment state
	State string `json:"state"`
	// Integer representation of the payment state
	StateI int32 `json:"state_i"`
}
