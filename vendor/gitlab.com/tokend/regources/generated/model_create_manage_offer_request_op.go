/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateManageOfferRequestOp struct {
	Key
	Attributes    CreateManageOfferRequestOpAttributes    `json:"attributes"`
	Relationships CreateManageOfferRequestOpRelationships `json:"relationships"`
}
type CreateManageOfferRequestOpResponse struct {
	Data     CreateManageOfferRequestOp `json:"data"`
	Included Included                   `json:"included"`
}

type CreateManageOfferRequestOpListResponse struct {
	Data     []CreateManageOfferRequestOp `json:"data"`
	Included Included                     `json:"included"`
	Links    *Links                       `json:"links"`
	Meta     json.RawMessage              `json:"meta,omitempty"`
}

func (r *CreateManageOfferRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateManageOfferRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateManageOfferRequestOp - returns CreateManageOfferRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateManageOfferRequestOp(key Key) *CreateManageOfferRequestOp {
	var createManageOfferRequestOp CreateManageOfferRequestOp
	if c.tryFindEntry(key, &createManageOfferRequestOp) {
		return &createManageOfferRequestOp
	}
	return nil
}
