/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateSaleRequestOp struct {
	Key
	Attributes    CreateSaleRequestOpAttributes    `json:"attributes"`
	Relationships CreateSaleRequestOpRelationships `json:"relationships"`
}
type CreateSaleRequestOpResponse struct {
	Data     CreateSaleRequestOp `json:"data"`
	Included Included            `json:"included"`
}

type CreateSaleRequestOpListResponse struct {
	Data     []CreateSaleRequestOp `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
	Meta     json.RawMessage       `json:"meta,omitempty"`
}

func (r *CreateSaleRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateSaleRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateSaleRequestOp - returns CreateSaleRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateSaleRequestOp(key Key) *CreateSaleRequestOp {
	var createSaleRequestOp CreateSaleRequestOp
	if c.tryFindEntry(key, &createSaleRequestOp) {
		return &createSaleRequestOp
	}
	return nil
}
