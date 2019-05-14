/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageOfferOpAttributes struct {
	BaseAmount Amount `json:"base_amount"`
	Fee        Fee    `json:"fee"`
	// Offer direction
	IsBuy     bool `json:"is_buy"`
	IsDeleted bool `json:"is_deleted"`
	// ID of the offer to manage
	OfferId int64 `json:"offer_id"`
	// ID of the order book
	OrderBookId int64  `json:"order_book_id"`
	Price       Amount `json:"price"`
}
