package regources

import "gitlab.com/tokend/go/xdr"

//ManageLimitsOp - details of corresponding op
type ManageLimitsOp struct {
	Key
	Attributes ManageLimitsOpAttributes `json:"attributes"`
}

//ManageLimitsAttributes - details of the manage limits op
type ManageLimitsOpAttributes struct {
	Action xdr.ManageLimitsAction
	Create *ManageLimitsCreationOp `json:"create"`
	Remove *ManageLimitsRemovalOp  `json:"remove"`
}

//ManageLimitsCreation - details of corresponding op
type ManageLimitsCreationOp struct {
	AccountRole     *xdr.Uint64     `json:"account_role,omitempty"`
	AccountAddress  string          `json:"account_address,omitempty"`
	StatsOpType     xdr.StatsOpType `json:"stats_op_type"`
	AssetCode       string          `json:"asset_code"`
	IsConvertNeeded bool            `json:"is_convert_needed"`
	DailyOut        Amount          `json:"daily_out"`
	WeeklyOut       Amount          `json:"weekly_out"`
	MonthlyOut      Amount          `json:"monthly_out"`
	AnnualOut       Amount          `json:"annual_out"`
}

//ManageLimitsRemoval - details of corresponding op
type ManageLimitsRemovalOp struct {
	LimitsID int64 `json:"limits_id"`
}
