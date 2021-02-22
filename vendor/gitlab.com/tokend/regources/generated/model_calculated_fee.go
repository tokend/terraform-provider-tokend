/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CalculatedFee struct {
	Key
	Attributes Fee `json:"attributes"`
}
type CalculatedFeeResponse struct {
	Data     CalculatedFee `json:"data"`
	Included Included      `json:"included"`
}

type CalculatedFeeListResponse struct {
	Data     []CalculatedFee `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *CalculatedFeeListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CalculatedFeeListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCalculatedFee - returns CalculatedFee from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCalculatedFee(key Key) *CalculatedFee {
	var calculatedFee CalculatedFee
	if c.tryFindEntry(key, &calculatedFee) {
		return &calculatedFee
	}
	return nil
}
