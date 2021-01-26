/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ManageSignerRoleOp struct {
	Key
	Attributes    *ManageSignerRoleOpAttributes   `json:"attributes,omitempty"`
	Relationships ManageSignerRoleOpRelationships `json:"relationships"`
}
type ManageSignerRoleOpResponse struct {
	Data     ManageSignerRoleOp `json:"data"`
	Included Included           `json:"included"`
}

type ManageSignerRoleOpListResponse struct {
	Data     []ManageSignerRoleOp `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *ManageSignerRoleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ManageSignerRoleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustManageSignerRoleOp - returns ManageSignerRoleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageSignerRoleOp(key Key) *ManageSignerRoleOp {
	var manageSignerRoleOp ManageSignerRoleOp
	if c.tryFindEntry(key, &manageSignerRoleOp) {
		return &manageSignerRoleOp
	}
	return nil
}
