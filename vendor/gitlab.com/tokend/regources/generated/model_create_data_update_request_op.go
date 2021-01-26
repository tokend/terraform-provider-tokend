/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateDataUpdateRequestOp struct {
	Key
	Attributes    CreateDataUpdateRequestOpAttributes    `json:"attributes"`
	Relationships CreateDataUpdateRequestOpRelationships `json:"relationships"`
}
type CreateDataUpdateRequestOpResponse struct {
	Data     CreateDataUpdateRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type CreateDataUpdateRequestOpListResponse struct {
	Data     []CreateDataUpdateRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *CreateDataUpdateRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateDataUpdateRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateDataUpdateRequestOp - returns CreateDataUpdateRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateDataUpdateRequestOp(key Key) *CreateDataUpdateRequestOp {
	var createDataUpdateRequestOp CreateDataUpdateRequestOp
	if c.tryFindEntry(key, &createDataUpdateRequestOp) {
		return &createDataUpdateRequestOp
	}
	return nil
}
