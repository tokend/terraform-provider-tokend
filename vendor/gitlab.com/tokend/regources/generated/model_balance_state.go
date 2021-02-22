/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type BalanceState struct {
	Key
	Attributes *BalanceStateAttributes `json:"attributes,omitempty"`
}
type BalanceStateResponse struct {
	Data     BalanceState `json:"data"`
	Included Included     `json:"included"`
}

type BalanceStateListResponse struct {
	Data     []BalanceState  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *BalanceStateListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *BalanceStateListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustBalanceState - returns BalanceState from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBalanceState(key Key) *BalanceState {
	var balanceState BalanceState
	if c.tryFindEntry(key, &balanceState) {
		return &balanceState
	}
	return nil
}
