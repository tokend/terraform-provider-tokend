/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ManageAccountRoleOp struct {
	Key
	Attributes    *ManageAccountRoleOpAttributes   `json:"attributes,omitempty"`
	Relationships ManageAccountRoleOpRelationships `json:"relationships"`
}
type ManageAccountRoleOpResponse struct {
	Data     ManageAccountRoleOp `json:"data"`
	Included Included            `json:"included"`
}

type ManageAccountRoleOpListResponse struct {
	Data     []ManageAccountRoleOp `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
	Meta     json.RawMessage       `json:"meta,omitempty"`
}

func (r *ManageAccountRoleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ManageAccountRoleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustManageAccountRoleOp - returns ManageAccountRoleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageAccountRoleOp(key Key) *ManageAccountRoleOp {
	var manageAccountRoleOp ManageAccountRoleOp
	if c.tryFindEntry(key, &manageAccountRoleOp) {
		return &manageAccountRoleOp
	}
	return nil
}
