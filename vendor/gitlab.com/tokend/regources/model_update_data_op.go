/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type UpdateDataOp struct {
	Key
	Attributes    UpdateDataOpAttributes    `json:"attributes"`
	Relationships UpdateDataOpRelationships `json:"relationships"`
}
type UpdateDataOpResponse struct {
	Data     UpdateDataOp `json:"data"`
	Included Included     `json:"included"`
}

type UpdateDataOpListResponse struct {
	Data     []UpdateDataOp  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *UpdateDataOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateDataOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateDataOp - returns UpdateDataOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateDataOp(key Key) *UpdateDataOp {
	var updateDataOp UpdateDataOp
	if c.tryFindEntry(key, &updateDataOp) {
		return &updateDataOp
	}
	return nil
}
