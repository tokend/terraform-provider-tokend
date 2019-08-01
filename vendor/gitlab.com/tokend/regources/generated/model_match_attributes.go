/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type MatchAttributes struct {
	// defines quote amount of the match
	BaseAmount Amount `json:"base_amount"`
	// time when the match was happened
	CreatedAt time.Time `json:"created_at"`
	// defines price of the match
	Price Amount `json:"price"`
	// defines base amount of the match
	QuoteAmount Amount `json:"quote_amount"`
}
