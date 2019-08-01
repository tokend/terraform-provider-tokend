/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type LedgerInfo struct {
	// time at which latest known ledger has been increased last time
	LastLedgerIncreaseTime time.Time `json:"last_ledger_increase_time"`
	// latest known ledger
	Latest uint64 `json:"latest"`
	// sequence of oldest ledger available
	OldestOnStart uint64 `json:"oldest_on_start"`
}
