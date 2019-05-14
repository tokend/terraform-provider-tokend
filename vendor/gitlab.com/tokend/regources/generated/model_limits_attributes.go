/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type LimitsAttributes struct {
	AnnualOut Amount `json:"annual_out"`
	DailyOut  Amount `json:"daily_out"`
	// if `true`, this limit is applied to operations in other assets except for those specified in `asset` relationships
	IsConvertNeeded bool   `json:"is_convert_needed"`
	MonthlyOut      Amount `json:"monthly_out"`
	// defines the type of operation for which a limit is applied. TODO: provide a list of such operations
	StatsOpType int32  `json:"stats_op_type"`
	WeeklyOut   Amount `json:"weekly_out"`
}
