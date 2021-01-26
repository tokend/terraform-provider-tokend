/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type StatisticsAttributes struct {
	AnnualOut Amount `json:"annual_out"`
	DailyOut  Amount `json:"daily_out"`
	// if `true`, this statistics entry exists in connection with operations in other assets except for those specified in `asset` relationships
	IsConvertNeeded bool   `json:"is_convert_needed"`
	MonthlyOut      Amount `json:"monthly_out"`
	// defines the type of operation statistics which entry exists for
	OperationType int32     `json:"operation_type"`
	UpdatedAt     time.Time `json:"updated_at"`
	WeeklyOut     Amount    `json:"weekly_out"`
}
