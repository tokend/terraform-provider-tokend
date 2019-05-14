/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type SetFeeOpAttributes struct {
	AccountAddress *string     `json:"account_address,omitempty"`
	AccountRole    *xdr.Uint64 `json:"account_role,omitempty"`
	// Unique identifier of the asset
	AssetCode string `json:"asset_code"`
	// * 0: \"payment_fee\" * 1: \"offer_fee\" * 2: \"withdrawal_fee\" * 3: \"issuance_fee\" * 4: \"invest_fee\" * 5: \"capital_deployment_fee\" * 6: \"operation_fee\" * 7: \"payout_fee\" * 8: \"atomic_swap_sale_fee\" * 9: \"atomic_swap_purchase_fee\"
	FeeType xdr.FeeType `json:"fee_type"`
	// Fixed amount to pay
	FixedFee Amount `json:"fixed_fee"`
	IsDelete bool   `json:"is_delete"`
	// Lower bound of fee applicability
	LowerBound Amount `json:"lower_bound"`
	// Percent to pay
	PercentFee Amount `json:"percent_fee"`
	// Subtype of the fee
	Subtype int64 `json:"subtype"`
	// Upper bound of fee applicability
	UpperBound Amount `json:"upper_bound"`
}
