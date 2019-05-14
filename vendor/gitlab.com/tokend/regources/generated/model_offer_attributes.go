/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type OfferAttributes struct {
	// defines the amount of offer in the base asset
	BaseAmount Amount `json:"base_amount"`
	// defines the time when the offer was created
	CreatedAt time.Time `json:"created_at"`
	Fee       Fee       `json:"fee"`
	// defines whether an offer created is on buying or selling the base_asset, or both
	IsBuy bool `json:"is_buy"`
	// defines whether an offer created is on selling or trading. Could be either `0` (secondary market) or some `saleId` (for specific sale) or `-1`
	OrderBookId string `json:"order_book_id"`
	// defines the price of an offer
	Price Amount `json:"price"`
	// defines the amount of offer in the quote asset
	QuoteAmount Amount `json:"quote_amount"`
}
