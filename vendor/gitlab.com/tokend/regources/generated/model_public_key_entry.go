/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type PublicKeyEntry struct {
	Key
	Relationships PublicKeyEntryRelationships `json:"relationships"`
}
type PublicKeyEntryResponse struct {
	Data     PublicKeyEntry `json:"data"`
	Included Included       `json:"included"`
}

type PublicKeyEntryListResponse struct {
	Data     []PublicKeyEntry `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustPublicKeyEntry - returns PublicKeyEntry from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPublicKeyEntry(key Key) *PublicKeyEntry {
	var publicKeyEntry PublicKeyEntry
	if c.tryFindEntry(key, &publicKeyEntry) {
		return &publicKeyEntry
	}
	return nil
}
