/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type KeyValueEntry struct {
	Key
	Attributes KeyValueEntryAttributes `json:"attributes"`
}
type KeyValueEntryResponse struct {
	Data     KeyValueEntry `json:"data"`
	Included Included      `json:"included"`
}

type KeyValueEntryListResponse struct {
	Data     []KeyValueEntry `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustKeyValueEntry - returns KeyValueEntry from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustKeyValueEntry(key Key) *KeyValueEntry {
	var keyValueEntry KeyValueEntry
	if c.tryFindEntry(key, &keyValueEntry) {
		return &keyValueEntry
	}
	return nil
}
