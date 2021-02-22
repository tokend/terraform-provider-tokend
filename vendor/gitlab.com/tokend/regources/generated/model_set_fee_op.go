/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type SetFeeOp struct {
	Key
	Attributes SetFeeOpAttributes `json:"attributes"`
}
type SetFeeOpResponse struct {
	Data     SetFeeOp `json:"data"`
	Included Included `json:"included"`
}

type SetFeeOpListResponse struct {
	Data     []SetFeeOp      `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *SetFeeOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *SetFeeOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustSetFeeOp - returns SetFeeOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSetFeeOp(key Key) *SetFeeOp {
	var setFeeOp SetFeeOp
	if c.tryFindEntry(key, &setFeeOp) {
		return &setFeeOp
	}
	return nil
}
