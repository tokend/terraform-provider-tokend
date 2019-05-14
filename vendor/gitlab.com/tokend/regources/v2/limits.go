package regources

type Limits struct {
	Key
	Attributes    LimitsAttr      `json:"attributes"`
	Relationships LimitsRelations `json:"relationships"`
}

type LimitsAttr struct {
	StatsOpType     int32  `json:"stats_op_type"`
	IsConvertNeeded bool   `json:"is_convert_needed"`
	DailyOut        Amount `json:"daily_out"`
	WeeklyOut       Amount `json:"weekly_out"`
	MonthlyOut      Amount `json:"monthly_out"`
	AnnualOut       Amount `json:"annual_out"`
}

type LimitsRelations struct {
	Account     *Relation `json:"account"`
	AccountRole *Relation `json:"account_role"`
	Asset       *Relation `json:"asset"`
}
