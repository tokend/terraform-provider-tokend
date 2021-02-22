/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Account struct {
	Key
	Attributes    AccountAttributes    `json:"attributes"`
	Relationships AccountRelationships `json:"relationships"`
}
type AccountResponse struct {
	Data     Account  `json:"data"`
	Included Included `json:"included"`
}

type AccountListResponse struct {
	Data     []Account       `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *AccountListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *AccountListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustAccount - returns Account from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAccount(key Key) *Account {
	var account Account
	if c.tryFindEntry(key, &account) {
		return &account
	}
	return nil
}
