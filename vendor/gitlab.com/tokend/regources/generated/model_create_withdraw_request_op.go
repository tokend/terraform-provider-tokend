/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateWithdrawRequestOp struct {
	Key
	Attributes    CreateWithdrawRequestOpAttributes    `json:"attributes"`
	Relationships CreateWithdrawRequestOpRelationships `json:"relationships"`
}
type CreateWithdrawRequestOpResponse struct {
	Data     CreateWithdrawRequestOp `json:"data"`
	Included Included                `json:"included"`
}

type CreateWithdrawRequestOpListResponse struct {
	Data     []CreateWithdrawRequestOp `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
	Meta     json.RawMessage           `json:"meta,omitempty"`
}

func (r *CreateWithdrawRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateWithdrawRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateWithdrawRequestOp - returns CreateWithdrawRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateWithdrawRequestOp(key Key) *CreateWithdrawRequestOp {
	var createWithdrawRequestOp CreateWithdrawRequestOp
	if c.tryFindEntry(key, &createWithdrawRequestOp) {
		return &createWithdrawRequestOp
	}
	return nil
}
