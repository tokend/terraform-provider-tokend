/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManageLimitsCreationOp struct {
	AccountAddress string      `json:"account_address"`
	AccountRole    *xdr.Uint64 `json:"account_role,omitempty"`
	AnnualOut      Amount      `json:"annual_out"`
	// Asset for which limits are being set
	AssetCode       string `json:"asset_code"`
	DailyOut        Amount `json:"daily_out"`
	IsConvertNeeded bool   `json:"is_convert_needed"`
	MonthlyOut      Amount `json:"monthly_out"`
	// * 1: \"payment_out\" * 2: \"withdraw\" * 3: \"spend\" * 4: \"deposit\" * 5: \"payout\"
	StatsOpType xdr.StatsOpType `json:"stats_op_type"`
	WeeklyOut   Amount          `json:"weekly_out"`
}
