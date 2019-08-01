/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"
import (
	"time"
)

type CreateSaleRequestAttributes struct {
	// indicates access definition type of the sale * 0 - none * 1 - whitelist * 2 - blacklist
	AccessDefinitionType SaleAccessDefinitionType `json:"access_definition_type"`
	// Maximum amount of base asset to be sold
	BaseAssetForHardCap Amount  `json:"base_asset_for_hard_cap"`
	CreatorDetails      Details `json:"creator_details"`
	// Time when a sale should end
	EndTime time.Time `json:"end_time"`
	// Maximal amount in base asset to be sold on sale
	HardCap Amount `json:"hard_cap"`
	// Defines sale type: * 1 - **basic sale** * 2 - **crowdfunding** sale * 3 - **fixed price** sale
	SaleType xdr.SaleType `json:"sale_type"`
	// Minimal amount in base asset for sale to reach to be considered successful
	SoftCap Amount `json:"soft_cap"`
	// Time when a sale should start
	StartTime time.Time `json:"start_time"`
}
