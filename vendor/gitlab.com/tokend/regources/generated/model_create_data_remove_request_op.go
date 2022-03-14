/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateDataRemoveRequestOp struct {
	Key
	Attributes    CreateDataRemoveRequestOpAttributes    `json:"attributes"`
	Relationships CreateDataRemoveRequestOpRelationships `json:"relationships"`
}
type CreateDataRemoveRequestOpResponse struct {
	Data     CreateDataRemoveRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type CreateDataRemoveRequestOpListResponse struct {
	Data     []CreateDataRemoveRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *CreateDataRemoveRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateDataRemoveRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateDataRemoveRequestOp - returns CreateDataRemoveRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateDataRemoveRequestOp(key Key) *CreateDataRemoveRequestOp {
	var createDataRemoveRequestOp CreateDataRemoveRequestOp
	if c.tryFindEntry(key, &createDataRemoveRequestOp) {
		return &createDataRemoveRequestOp
	}
	return nil
}
