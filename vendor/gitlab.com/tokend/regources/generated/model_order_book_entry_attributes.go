/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type OrderBookEntryAttributes struct {
	// defines the base amount of an offer
	BaseAmount Amount `json:"base_amount"`
	// defines the time when the offer was created
	CreatedAt time.Time `json:"created_at"`
	// defines the base amount of current entry plus base amounts of more beneficial entries
	CumulativeBaseAmount Amount `json:"cumulative_base_amount"`
	// defines the quote amount of current entry plus quote amounts of more beneficial entries
	CumulativeQuoteAmount Amount `json:"cumulative_quote_amount"`
	// defines whether an offer is on buying or selling the base_asset
	IsBuy bool `json:"is_buy"`
	// defines the price of an offer
	Price Amount `json:"price"`
	// defines the quote amount of an offer
	QuoteAmount Amount `json:"quote_amount"`
}
