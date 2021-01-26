/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ManageLimitsOp struct {
	Key
	Attributes ManageLimitsOpAttributes `json:"attributes"`
}
type ManageLimitsOpResponse struct {
	Data     ManageLimitsOp `json:"data"`
	Included Included       `json:"included"`
}

type ManageLimitsOpListResponse struct {
	Data     []ManageLimitsOp `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
	Meta     json.RawMessage  `json:"meta,omitempty"`
}

func (r *ManageLimitsOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ManageLimitsOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustManageLimitsOp - returns ManageLimitsOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageLimitsOp(key Key) *ManageLimitsOp {
	var manageLimitsOp ManageLimitsOp
	if c.tryFindEntry(key, &manageLimitsOp) {
		return &manageLimitsOp
	}
	return nil
}
