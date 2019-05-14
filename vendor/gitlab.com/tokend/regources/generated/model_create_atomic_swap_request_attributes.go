/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateAtomicSwapRequestAttributes struct {
	// Amount in base asset to perform atomic swap on
	BaseAmount Amount   `json:"base_amount"`
	Details    *Details `json:"details,omitempty"`
}
