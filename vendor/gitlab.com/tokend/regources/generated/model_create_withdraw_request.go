/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateWithdrawRequest struct {
	Key
	Attributes    CreateWithdrawRequestAttributes    `json:"attributes"`
	Relationships CreateWithdrawRequestRelationships `json:"relationships"`
}
type CreateWithdrawRequestResponse struct {
	Data     CreateWithdrawRequest `json:"data"`
	Included Included              `json:"included"`
}

type CreateWithdrawRequestListResponse struct {
	Data     []CreateWithdrawRequest `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
	Meta     json.RawMessage         `json:"meta,omitempty"`
}

func (r *CreateWithdrawRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateWithdrawRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateWithdrawRequest - returns CreateWithdrawRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateWithdrawRequest(key Key) *CreateWithdrawRequest {
	var createWithdrawRequest CreateWithdrawRequest
	if c.tryFindEntry(key, &createWithdrawRequest) {
		return &createWithdrawRequest
	}
	return nil
}
