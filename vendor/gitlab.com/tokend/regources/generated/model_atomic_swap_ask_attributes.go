/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type AtomicSwapAskAttributes struct {
	// Amount that can be bought through atomic swap bid request
	AvailableAmount Amount `json:"available_amount"`
	// time when the atomic swap ask was created
	CreatedAt time.Time `json:"created_at"`
	Details   Details   `json:"details"`
	// defines whether creating atomic swap requests for this ask is forbidden
	IsCanceled bool `json:"is_canceled"`
	// Amount that that is being processed now through atomic swap bid requests
	LockedAmount Amount `json:"locked_amount"`
}
