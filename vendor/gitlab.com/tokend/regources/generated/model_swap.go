/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Swap struct {
	Key
	Attributes    SwapAttributes    `json:"attributes"`
	Relationships SwapRelationships `json:"relationships"`
}
type SwapResponse struct {
	Data     Swap     `json:"data"`
	Included Included `json:"included"`
}

type SwapListResponse struct {
	Data     []Swap          `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *SwapListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *SwapListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustSwap - returns Swap from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSwap(key Key) *Swap {
	var swap Swap
	if c.tryFindEntry(key, &swap) {
		return &swap
	}
	return nil
}
