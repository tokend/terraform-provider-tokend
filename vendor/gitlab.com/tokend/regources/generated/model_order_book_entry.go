/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type OrderBookEntry struct {
	Key
	Attributes    OrderBookEntryAttributes    `json:"attributes"`
	Relationships OrderBookEntryRelationships `json:"relationships"`
}
type OrderBookEntryResponse struct {
	Data     OrderBookEntry `json:"data"`
	Included Included       `json:"included"`
}

type OrderBookEntryListResponse struct {
	Data     []OrderBookEntry `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustOrderBookEntry - returns OrderBookEntry from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustOrderBookEntry(key Key) *OrderBookEntry {
	var orderBookEntry OrderBookEntry
	if c.tryFindEntry(key, &orderBookEntry) {
		return &orderBookEntry
	}
	return nil
}
