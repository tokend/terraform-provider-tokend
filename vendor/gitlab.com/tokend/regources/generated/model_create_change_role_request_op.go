/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateChangeRoleRequestOp struct {
	Key
	Attributes    CreateChangeRoleRequestOpAttributes    `json:"attributes"`
	Relationships CreateChangeRoleRequestOpRelationships `json:"relationships"`
}
type CreateChangeRoleRequestOpResponse struct {
	Data     CreateChangeRoleRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type CreateChangeRoleRequestOpListResponse struct {
	Data     []CreateChangeRoleRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *CreateChangeRoleRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateChangeRoleRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateChangeRoleRequestOp - returns CreateChangeRoleRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateChangeRoleRequestOp(key Key) *CreateChangeRoleRequestOp {
	var createChangeRoleRequestOp CreateChangeRoleRequestOp
	if c.tryFindEntry(key, &createChangeRoleRequestOp) {
		return &createChangeRoleRequestOp
	}
	return nil
}
