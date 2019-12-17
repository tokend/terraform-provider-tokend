/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type DestructionOp struct {
	Key
	Attributes    DestructionOpAttributes    `json:"attributes"`
	Relationships DestructionOpRelationships `json:"relationships"`
}
type DestructionOpResponse struct {
	Data     DestructionOp `json:"data"`
	Included Included      `json:"included"`
}

type DestructionOpListResponse struct {
	Data     []DestructionOp `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *DestructionOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DestructionOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDestructionOp - returns DestructionOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDestructionOp(key Key) *DestructionOp {
	var destructionOp DestructionOp
	if c.tryFindEntry(key, &destructionOp) {
		return &destructionOp
	}
	return nil
}
