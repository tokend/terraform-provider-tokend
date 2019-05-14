/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManageAssetPairOpAttributes struct {
	MaxPriceStep            Amount `json:"max_price_step"`
	PhysicalPrice           Amount `json:"physical_price"`
	PhysicalPriceCorrection Amount `json:"physical_price_correction"`
	// Bit mask. * 1: \"tradeable_secondary_market\" * 2: \"physical_price_restriction\" * 4: \"current_price_restriction\"
	Policies xdr.AssetPairPolicy `json:"policies"`
}
