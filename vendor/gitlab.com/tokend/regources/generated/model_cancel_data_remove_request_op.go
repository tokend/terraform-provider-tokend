/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CancelDataRemoveRequestOp struct {
	Key
	Relationships CancelDataRemoveRequestOpRelationships `json:"relationships"`
}
type CancelDataRemoveRequestOpResponse struct {
	Data     CancelDataRemoveRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type CancelDataRemoveRequestOpListResponse struct {
	Data     []CancelDataRemoveRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *CancelDataRemoveRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CancelDataRemoveRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCancelDataRemoveRequestOp - returns CancelDataRemoveRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCancelDataRemoveRequestOp(key Key) *CancelDataRemoveRequestOp {
	var cancelDataRemoveRequestOp CancelDataRemoveRequestOp
	if c.tryFindEntry(key, &cancelDataRemoveRequestOp) {
		return &cancelDataRemoveRequestOp
	}
	return nil
}
