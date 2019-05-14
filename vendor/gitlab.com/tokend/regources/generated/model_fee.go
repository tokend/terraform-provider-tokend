/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

// Fee - describes fee happened on balance. Direction of fee depends on the operation (depending on effect might be charged, locked, unlocked, for all incoming effects but unlocked it's always charged)
type Fee struct {
	// percent of the operation amount
	CalculatedPercent Amount `json:"calculated_percent"`
	// fixed fee charged for the operation
	Fixed Amount `json:"fixed"`
}
