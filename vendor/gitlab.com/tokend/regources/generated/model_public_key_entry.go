/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

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
	Meta     json.RawMessage  `json:"meta,omitempty"`
}

func (r *PublicKeyEntryListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *PublicKeyEntryListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
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
