/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type LedgerEntryChange struct {
	Key
	Attributes LedgerEntryChangeAttributes `json:"attributes"`
}
type LedgerEntryChangeResponse struct {
	Data     LedgerEntryChange `json:"data"`
	Included Included          `json:"included"`
}

type LedgerEntryChangeListResponse struct {
	Data     []LedgerEntryChange `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustLedgerEntryChange - returns LedgerEntryChange from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustLedgerEntryChange(key Key) *LedgerEntryChange {
	var ledgerEntryChange LedgerEntryChange
	if c.tryFindEntry(key, &ledgerEntryChange) {
		return &ledgerEntryChange
	}
	return nil
}
