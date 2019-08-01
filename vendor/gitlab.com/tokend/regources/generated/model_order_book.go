/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type OrderBook struct {
	Key
	Relationships OrderBookRelationships `json:"relationships"`
}
type OrderBookResponse struct {
	Data     OrderBook `json:"data"`
	Included Included  `json:"included"`
}

type OrderBookListResponse struct {
	Data     []OrderBook `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustOrderBook - returns OrderBook from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustOrderBook(key Key) *OrderBook {
	var orderBook OrderBook
	if c.tryFindEntry(key, &orderBook) {
		return &orderBook
	}
	return nil
}
