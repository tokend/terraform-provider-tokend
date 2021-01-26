/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type EffectBalanceChange struct {
	Key
	Attributes EffectBalanceChangeAttributes `json:"attributes"`
}
type EffectBalanceChangeResponse struct {
	Data     EffectBalanceChange `json:"data"`
	Included Included            `json:"included"`
}

type EffectBalanceChangeListResponse struct {
	Data     []EffectBalanceChange `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
	Meta     json.RawMessage       `json:"meta,omitempty"`
}

func (r *EffectBalanceChangeListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *EffectBalanceChangeListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustEffectBalanceChange - returns EffectBalanceChange from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEffectBalanceChange(key Key) *EffectBalanceChange {
	var effectBalanceChange EffectBalanceChange
	if c.tryFindEntry(key, &effectBalanceChange) {
		return &effectBalanceChange
	}
	return nil
}
