/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type OpenSwapOpAttributes struct {
	Amount         Amount  `json:"amount"`
	DestinationFee Fee     `json:"destination_fee"`
	Details        Details `json:"details"`
	// time when swap can be cancelled
	LockTime time.Time `json:"lock_time"`
	// hash of the swap secret in hexadecimal format
	SecretHash string `json:"secret_hash"`
	SourceFee  Fee    `json:"source_fee"`
	// Whether source of the swap should pay destination fee
	SourcePayForDestination bool `json:"source_pay_for_destination"`
}
