/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Balance struct {
	Key
	Relationships *BalanceRelationships `json:"relationships,omitempty"`
}
type BalanceResponse struct {
	Data     Balance  `json:"data"`
	Included Included `json:"included"`
}

type BalanceListResponse struct {
	Data     []Balance       `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *BalanceListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *BalanceListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustBalance - returns Balance from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBalance(key Key) *Balance {
	var balance Balance
	if c.tryFindEntry(key, &balance) {
		return &balance
	}
	return nil
}
