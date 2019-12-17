/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type CreateRoleOp struct {
	Key
	Attributes    CreateRoleOpAttributes    `json:"attributes"`
	Relationships CreateRoleOpRelationships `json:"relationships"`
}
type CreateRoleOpResponse struct {
	Data     CreateRoleOp `json:"data"`
	Included Included     `json:"included"`
}

type CreateRoleOpListResponse struct {
	Data     []CreateRoleOp  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *CreateRoleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateRoleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateRoleOp - returns CreateRoleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateRoleOp(key Key) *CreateRoleOp {
	var createRoleOp CreateRoleOp
	if c.tryFindEntry(key, &createRoleOp) {
		return &createRoleOp
	}
	return nil
}
