/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ManageCreatePollRequestOp struct {
	Key
	Attributes    *ManageCreatePollRequestOpAttributes    `json:"attributes,omitempty"`
	Relationships *ManageCreatePollRequestOpRelationships `json:"relationships,omitempty"`
}
type ManageCreatePollRequestOpResponse struct {
	Data     ManageCreatePollRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type ManageCreatePollRequestOpListResponse struct {
	Data     []ManageCreatePollRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *ManageCreatePollRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ManageCreatePollRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustManageCreatePollRequestOp - returns ManageCreatePollRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageCreatePollRequestOp(key Key) *ManageCreatePollRequestOp {
	var manageCreatePollRequestOp ManageCreatePollRequestOp
	if c.tryFindEntry(key, &manageCreatePollRequestOp) {
		return &manageCreatePollRequestOp
	}
	return nil
}
