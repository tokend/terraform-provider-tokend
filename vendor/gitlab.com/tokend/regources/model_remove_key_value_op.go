/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type RemoveKeyValueOp struct {
	Key
	Attributes RemoveKeyValueOpAttributes `json:"attributes"`
}
type RemoveKeyValueOpResponse struct {
	Data     RemoveKeyValueOp `json:"data"`
	Included Included         `json:"included"`
}

type RemoveKeyValueOpListResponse struct {
	Data     []RemoveKeyValueOp `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
	Meta     json.RawMessage    `json:"meta,omitempty"`
}

func (r *RemoveKeyValueOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *RemoveKeyValueOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustRemoveKeyValueOp - returns RemoveKeyValueOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRemoveKeyValueOp(key Key) *RemoveKeyValueOp {
	var removeKeyValueOp RemoveKeyValueOp
	if c.tryFindEntry(key, &removeKeyValueOp) {
		return &removeKeyValueOp
	}
	return nil
}
