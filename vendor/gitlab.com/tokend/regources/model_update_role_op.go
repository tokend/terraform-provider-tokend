/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type UpdateRoleOp struct {
	Key
	Attributes    UpdateRoleOpAttributes    `json:"attributes"`
	Relationships UpdateRoleOpRelationships `json:"relationships"`
}
type UpdateRoleOpResponse struct {
	Data     UpdateRoleOp `json:"data"`
	Included Included     `json:"included"`
}

type UpdateRoleOpListResponse struct {
	Data     []UpdateRoleOp  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *UpdateRoleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateRoleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateRoleOp - returns UpdateRoleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateRoleOp(key Key) *UpdateRoleOp {
	var updateRoleOp UpdateRoleOp
	if c.tryFindEntry(key, &updateRoleOp) {
		return &updateRoleOp
	}
	return nil
}
