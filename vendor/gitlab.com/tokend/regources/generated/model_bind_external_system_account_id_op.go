/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type BindExternalSystemAccountIdOp struct {
	Key
	Attributes BindExternalSystemAccountIdOpAttributes `json:"attributes"`
}
type BindExternalSystemAccountIdOpResponse struct {
	Data     BindExternalSystemAccountIdOp `json:"data"`
	Included Included                      `json:"included"`
}

type BindExternalSystemAccountIdOpListResponse struct {
	Data     []BindExternalSystemAccountIdOp `json:"data"`
	Included Included                        `json:"included"`
	Links    *Links                          `json:"links"`
	Meta     json.RawMessage                 `json:"meta,omitempty"`
}

func (r *BindExternalSystemAccountIdOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *BindExternalSystemAccountIdOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustBindExternalSystemAccountIdOp - returns BindExternalSystemAccountIdOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBindExternalSystemAccountIdOp(key Key) *BindExternalSystemAccountIdOp {
	var bindExternalSystemAccountIdOp BindExternalSystemAccountIdOp
	if c.tryFindEntry(key, &bindExternalSystemAccountIdOp) {
		return &bindExternalSystemAccountIdOp
	}
	return nil
}
