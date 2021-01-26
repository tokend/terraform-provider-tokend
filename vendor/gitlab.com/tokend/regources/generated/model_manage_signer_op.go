/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ManageSignerOp struct {
	Key
	Attributes    *ManageSignerOpAttributes    `json:"attributes,omitempty"`
	Relationships *ManageSignerOpRelationships `json:"relationships,omitempty"`
}
type ManageSignerOpResponse struct {
	Data     ManageSignerOp `json:"data"`
	Included Included       `json:"included"`
}

type ManageSignerOpListResponse struct {
	Data     []ManageSignerOp `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
	Meta     json.RawMessage  `json:"meta,omitempty"`
}

func (r *ManageSignerOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ManageSignerOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustManageSignerOp - returns ManageSignerOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageSignerOp(key Key) *ManageSignerOp {
	var manageSignerOp ManageSignerOp
	if c.tryFindEntry(key, &manageSignerOp) {
		return &manageSignerOp
	}
	return nil
}
