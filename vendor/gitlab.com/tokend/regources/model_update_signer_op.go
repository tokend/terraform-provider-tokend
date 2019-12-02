/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type UpdateSignerOp struct {
	Key
	Attributes    UpdateSignerOpAttributes    `json:"attributes"`
	Relationships UpdateSignerOpRelationships `json:"relationships"`
}
type UpdateSignerOpResponse struct {
	Data     UpdateSignerOp `json:"data"`
	Included Included       `json:"included"`
}

type UpdateSignerOpListResponse struct {
	Data     []UpdateSignerOp `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
	Meta     json.RawMessage  `json:"meta,omitempty"`
}

func (r *UpdateSignerOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateSignerOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateSignerOp - returns UpdateSignerOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateSignerOp(key Key) *UpdateSignerOp {
	var updateSignerOp UpdateSignerOp
	if c.tryFindEntry(key, &updateSignerOp) {
		return &updateSignerOp
	}
	return nil
}
