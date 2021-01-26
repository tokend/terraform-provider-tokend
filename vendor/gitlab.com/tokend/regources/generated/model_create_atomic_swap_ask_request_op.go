/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateAtomicSwapAskRequestOp struct {
	Key
	Attributes    CreateAtomicSwapAskRequestOpAttributes    `json:"attributes"`
	Relationships CreateAtomicSwapAskRequestOpRelationships `json:"relationships"`
}
type CreateAtomicSwapAskRequestOpResponse struct {
	Data     CreateAtomicSwapAskRequestOp `json:"data"`
	Included Included                     `json:"included"`
}

type CreateAtomicSwapAskRequestOpListResponse struct {
	Data     []CreateAtomicSwapAskRequestOp `json:"data"`
	Included Included                       `json:"included"`
	Links    *Links                         `json:"links"`
	Meta     json.RawMessage                `json:"meta,omitempty"`
}

func (r *CreateAtomicSwapAskRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateAtomicSwapAskRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateAtomicSwapAskRequestOp - returns CreateAtomicSwapAskRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAtomicSwapAskRequestOp(key Key) *CreateAtomicSwapAskRequestOp {
	var createAtomicSwapAskRequestOp CreateAtomicSwapAskRequestOp
	if c.tryFindEntry(key, &createAtomicSwapAskRequestOp) {
		return &createAtomicSwapAskRequestOp
	}
	return nil
}
