/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ConvertedBalanceState struct {
	Key
	Attributes    ConvertedBalanceStateAttributes    `json:"attributes"`
	Relationships ConvertedBalanceStateRelationships `json:"relationships"`
}
type ConvertedBalanceStateResponse struct {
	Data     ConvertedBalanceState `json:"data"`
	Included Included              `json:"included"`
}

type ConvertedBalanceStateListResponse struct {
	Data     []ConvertedBalanceState `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
	Meta     json.RawMessage         `json:"meta,omitempty"`
}

func (r *ConvertedBalanceStateListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ConvertedBalanceStateListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustConvertedBalanceState - returns ConvertedBalanceState from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustConvertedBalanceState(key Key) *ConvertedBalanceState {
	var convertedBalanceState ConvertedBalanceState
	if c.tryFindEntry(key, &convertedBalanceState) {
		return &convertedBalanceState
	}
	return nil
}
