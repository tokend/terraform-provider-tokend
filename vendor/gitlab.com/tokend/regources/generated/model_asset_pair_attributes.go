/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type AssetPairAttributes struct {
	// defines the maximal allowed price change on offer matching
	MaxPriceStep Amount `json:"max_price_step"`
	// minimum value for which the pair can be traded
	PhysicalPriceCorrection Amount              `json:"physical_price_correction"`
	Policies                xdr.AssetPairPolicy `json:"policies"`
	// defines an asset pair price as quote asset divided by base asset (i.e., amount of quote asset per 1 base asset)
	Price Amount `json:"price"`
}
