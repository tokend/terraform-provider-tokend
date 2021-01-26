/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type SwapAttributes struct {
	Amount Amount `json:"amount"`
	// time when the swap was created
	CreatedAt      time.Time `json:"created_at"`
	DestinationFee Fee       `json:"destination_fee"`
	Details        Details   `json:"details"`
	// time when the swap expires
	LockTime time.Time `json:"lock_time"`
	// secret of the swap in hexadecimal format
	Secret *string `json:"secret,omitempty"`
	// Hash of the swap secret in hexadecimal format
	SecretHash string `json:"secret_hash"`
	SourceFee  Fee    `json:"source_fee"`
	// * 0 - \"open\" * 1 - \"closed\" * 2 - \"cancelled\"
	State SwapState `json:"state"`
}
