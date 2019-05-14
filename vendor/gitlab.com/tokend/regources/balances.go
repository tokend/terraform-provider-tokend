package regources

type BalancesReport struct {
	TotalAccountsCount struct {
		ZeroBalance     int `json:"zero_balance"`
		PositiveBalance int `json:"positive_balance"`
		AboveThreshold  int `json:"above_threshold"`
		BelowThreshold  int `json:"below_threshold"`
	} `json:"total_accounts_count"`
}
