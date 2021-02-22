/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateSaleRequest struct {
	Key
	Attributes    CreateSaleRequestAttributes    `json:"attributes"`
	Relationships CreateSaleRequestRelationships `json:"relationships"`
}
type CreateSaleRequestResponse struct {
	Data     CreateSaleRequest `json:"data"`
	Included Included          `json:"included"`
}

type CreateSaleRequestListResponse struct {
	Data     []CreateSaleRequest `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *CreateSaleRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateSaleRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateSaleRequest - returns CreateSaleRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateSaleRequest(key Key) *CreateSaleRequest {
	var createSaleRequest CreateSaleRequest
	if c.tryFindEntry(key, &createSaleRequest) {
		return &createSaleRequest
	}
	return nil
}
