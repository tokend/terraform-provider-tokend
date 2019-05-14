/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"
import (
	"time"
)

type SaleRequestAttributes struct {
	Details *Details `json:"details,omitempty"`
	// End time of the sale
	EndTime *time.Time `json:"end_time,omitempty"`
	// * 1 - open * 2 - closed * 3 - cancelled
	SaleState *SaleState `json:"sale_state,omitempty"`
	// * 1 - basic sale * 2 - crowdfunding sale * 3 - fixed price sale
	SaleType *xdr.SaleType `json:"sale_type,omitempty"`
	// Start time of the sale
	StartTime *time.Time `json:"start_time,omitempty"`
}
