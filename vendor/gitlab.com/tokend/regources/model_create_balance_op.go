/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateBalanceOp struct {
	Key
	Relationships CreateBalanceOpRelationships `json:"relationships"`
}
type CreateBalanceOpResponse struct {
	Data     CreateBalanceOp `json:"data"`
	Included Included        `json:"included"`
}

type CreateBalanceOpListResponse struct {
	Data     []CreateBalanceOp `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
	Meta     json.RawMessage   `json:"meta,omitempty"`
}

func (r *CreateBalanceOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateBalanceOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateBalanceOp - returns CreateBalanceOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateBalanceOp(key Key) *CreateBalanceOp {
	var createBalanceOp CreateBalanceOp
	if c.tryFindEntry(key, &createBalanceOp) {
		return &createBalanceOp
	}
	return nil
}
