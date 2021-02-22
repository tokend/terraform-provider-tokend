/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type LicenseInfoAttributes struct {
	// maximal allowed number of admins
	AdminCount int64 `json:"admin_count"`
	// current number of admins
	CurrentAdminCount int64 `json:"current_admin_count"`
	// license expiration date
	DueDate time.Time `json:"due_date"`
}
