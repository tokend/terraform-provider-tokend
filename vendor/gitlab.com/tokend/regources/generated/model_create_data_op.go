/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateDataOp struct {
	Key
	Attributes    CreateDataOpAttributes    `json:"attributes"`
	Relationships CreateDataOpRelationships `json:"relationships"`
}
type CreateDataOpResponse struct {
	Data     CreateDataOp `json:"data"`
	Included Included     `json:"included"`
}

type CreateDataOpListResponse struct {
	Data     []CreateDataOp  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *CreateDataOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateDataOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateDataOp - returns CreateDataOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateDataOp(key Key) *CreateDataOp {
	var createDataOp CreateDataOp
	if c.tryFindEntry(key, &createDataOp) {
		return &createDataOp
	}
	return nil
}
