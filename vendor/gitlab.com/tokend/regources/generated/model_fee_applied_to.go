/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type FeeAppliedTo struct {
	// Unique identifier of an asset
	Asset string `json:"asset"`
	// Type of the fee
	FeeType         int32       `json:"fee_type"`
	FeeTypeExtended xdr.FeeType `json:"fee_type_extended"`
	// lower bound for the fee
	LowerBound Amount `json:"lower_bound"`
	// Subtype of the fee
	Subtype int64 `json:"subtype"`
	// upper bound for the fee
	UpperBound Amount `json:"upper_bound"`
}
