/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"
import (
	"time"
)

type CreateSaleRequestAttributes struct {
	// Maximum amount of base asset to be sold
	BaseAssetForHardCap Amount  `json:"base_asset_for_hard_cap"`
	CreatorDetails      Details `json:"creator_details"`
	// Time when a sale should end
	EndTime time.Time `json:"end_time"`
	// Defines sale type: * 1 - **basic sale** * 2 - **crowdfunding** sale * 3 - **fixed price** sale
	SaleType xdr.SaleType `json:"sale_type"`
	// Time when a sale should start
	StartTime time.Time `json:"start_time"`
}
