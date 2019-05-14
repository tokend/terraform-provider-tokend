/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type LicenseOpAttributes struct {
	// Allowed number of admins to set
	AdminCount uint64 `json:"admin_count"`
	// End of the licensed period
	DueDate time.Time `json:"due_date"`
	// Stamped ledger hash
	LedgerHash string `json:"ledger_hash"`
	// Hash of the previous license
	PrevLicenseHash string   `json:"prev_license_hash"`
	Signatures      []string `json:"signatures"`
}
