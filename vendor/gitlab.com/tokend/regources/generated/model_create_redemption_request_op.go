/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateRedemptionRequestOp struct {
	Key
	Attributes    CreateRedemptionRequestOpAttributes    `json:"attributes"`
	Relationships CreateRedemptionRequestOpRelationships `json:"relationships"`
}
type CreateRedemptionRequestOpResponse struct {
	Data     CreateRedemptionRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type CreateRedemptionRequestOpListResponse struct {
	Data     []CreateRedemptionRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *CreateRedemptionRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateRedemptionRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateRedemptionRequestOp - returns CreateRedemptionRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateRedemptionRequestOp(key Key) *CreateRedemptionRequestOp {
	var createRedemptionRequestOp CreateRedemptionRequestOp
	if c.tryFindEntry(key, &createRedemptionRequestOp) {
		return &createRedemptionRequestOp
	}
	return nil
}
