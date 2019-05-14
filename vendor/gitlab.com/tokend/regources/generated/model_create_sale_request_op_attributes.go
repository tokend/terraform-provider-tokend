/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type CreateSaleRequestOpAttributes struct {
	CreatorDetails Details `json:"creator_details"`
	// End time of the sale
	EndTime time.Time `json:"end_time"`
	// Maximal amount in base asset to be sold on sale
	HardCap Amount `json:"hard_cap"`
	// Minimal amount in base asset for sale to reach to be considered successful
	SoftCap Amount `json:"soft_cap"`
	// Start time of the sale
	StartTime time.Time `json:"start_time"`
}
