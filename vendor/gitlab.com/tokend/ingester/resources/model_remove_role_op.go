/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type RemoveRoleOp struct {
	Key
	Relationships RemoveRoleOpRelationships `json:"relationships"`
}
type RemoveRoleOpResponse struct {
	Data     RemoveRoleOp `json:"data"`
	Included Included     `json:"included"`
}

type RemoveRoleOpListResponse struct {
	Data     []RemoveRoleOp  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *RemoveRoleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *RemoveRoleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustRemoveRoleOp - returns RemoveRoleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRemoveRoleOp(key Key) *RemoveRoleOp {
	var removeRoleOp RemoveRoleOp
	if c.tryFindEntry(key, &removeRoleOp) {
		return &removeRoleOp
	}
	return nil
}
