package regources

type FeeEntry struct {
	Asset       string `json:"asset"`
	Fixed       string `json:"fixed"`
	Percent     string `json:"percent"`
	FeeType     int    `json:"fee_type"`
	Subtype     int64  `json:"subtype"`
	AccountID   string `json:"account_id"`
	AccountRole uint64 `json:"account_type"`
	LowerBound  string `json:"lower_bound"`
	UpperBound  string `json:"upper_bound"`
	FeeAsset    string `json:"fee_asset"`
	Exists      bool   `json:"exists"`
}
