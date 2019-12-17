/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type CreateSignerOp struct {
	Key
	Attributes    CreateSignerOpAttributes    `json:"attributes"`
	Relationships CreateSignerOpRelationships `json:"relationships"`
}
type CreateSignerOpResponse struct {
	Data     CreateSignerOp `json:"data"`
	Included Included       `json:"included"`
}

type CreateSignerOpListResponse struct {
	Data     []CreateSignerOp `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
	Meta     json.RawMessage  `json:"meta,omitempty"`
}

func (r *CreateSignerOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateSignerOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateSignerOp - returns CreateSignerOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateSignerOp(key Key) *CreateSignerOp {
	var createSignerOp CreateSignerOp
	if c.tryFindEntry(key, &createSignerOp) {
		return &createSignerOp
	}
	return nil
}
