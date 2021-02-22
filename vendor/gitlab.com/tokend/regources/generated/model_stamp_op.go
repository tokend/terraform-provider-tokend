/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type StampOp struct {
	Key
	Attributes StampOpAttributes `json:"attributes"`
}
type StampOpResponse struct {
	Data     StampOp  `json:"data"`
	Included Included `json:"included"`
}

type StampOpListResponse struct {
	Data     []StampOp       `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *StampOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *StampOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustStampOp - returns StampOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustStampOp(key Key) *StampOp {
	var stampOp StampOp
	if c.tryFindEntry(key, &stampOp) {
		return &stampOp
	}
	return nil
}
