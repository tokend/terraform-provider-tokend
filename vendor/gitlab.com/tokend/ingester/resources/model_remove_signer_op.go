/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type RemoveSignerOp struct {
	Key
	Relationships RemoveSignerOpRelationships `json:"relationships"`
}
type RemoveSignerOpResponse struct {
	Data     RemoveSignerOp `json:"data"`
	Included Included       `json:"included"`
}

type RemoveSignerOpListResponse struct {
	Data     []RemoveSignerOp `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
	Meta     json.RawMessage  `json:"meta,omitempty"`
}

func (r *RemoveSignerOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *RemoveSignerOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustRemoveSignerOp - returns RemoveSignerOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRemoveSignerOp(key Key) *RemoveSignerOp {
	var removeSignerOp RemoveSignerOp
	if c.tryFindEntry(key, &removeSignerOp) {
		return &removeSignerOp
	}
	return nil
}
