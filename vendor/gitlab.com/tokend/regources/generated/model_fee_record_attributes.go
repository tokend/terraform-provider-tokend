/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type FeeRecordAttributes struct {
	AppliedTo FeeAppliedTo `json:"applied_to"`
	// Fixed amount to pay
	Fixed Amount `json:"fixed"`
	// Percent to calculate the fee with (e.g., 1.0 represents 100%, 0.21 represents 21%)
	Percent Amount `json:"percent"`
}
