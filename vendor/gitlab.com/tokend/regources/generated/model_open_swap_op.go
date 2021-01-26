/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type OpenSwapOp struct {
	Key
	Attributes    OpenSwapOpAttributes    `json:"attributes"`
	Relationships OpenSwapOpRelationships `json:"relationships"`
}
type OpenSwapOpResponse struct {
	Data     OpenSwapOp `json:"data"`
	Included Included   `json:"included"`
}

type OpenSwapOpListResponse struct {
	Data     []OpenSwapOp    `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *OpenSwapOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *OpenSwapOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustOpenSwapOp - returns OpenSwapOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustOpenSwapOp(key Key) *OpenSwapOp {
	var openSwapOp OpenSwapOp
	if c.tryFindEntry(key, &openSwapOp) {
		return &openSwapOp
	}
	return nil
}
