/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CancelDataUpdateRequestOp struct {
	Key
	Relationships CancelDataUpdateRequestOpRelationships `json:"relationships"`
}
type CancelDataUpdateRequestOpResponse struct {
	Data     CancelDataUpdateRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type CancelDataUpdateRequestOpListResponse struct {
	Data     []CancelDataUpdateRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *CancelDataUpdateRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CancelDataUpdateRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCancelDataUpdateRequestOp - returns CancelDataUpdateRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCancelDataUpdateRequestOp(key Key) *CancelDataUpdateRequestOp {
	var cancelDataUpdateRequestOp CancelDataUpdateRequestOp
	if c.tryFindEntry(key, &cancelDataUpdateRequestOp) {
		return &cancelDataUpdateRequestOp
	}
	return nil
}
