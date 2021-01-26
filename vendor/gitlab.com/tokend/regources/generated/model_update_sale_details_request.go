/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type UpdateSaleDetailsRequest struct {
	Key
	Attributes    UpdateSaleDetailsRequestAttributes    `json:"attributes"`
	Relationships UpdateSaleDetailsRequestRelationships `json:"relationships"`
}
type UpdateSaleDetailsRequestResponse struct {
	Data     UpdateSaleDetailsRequest `json:"data"`
	Included Included                 `json:"included"`
}

type UpdateSaleDetailsRequestListResponse struct {
	Data     []UpdateSaleDetailsRequest `json:"data"`
	Included Included                   `json:"included"`
	Links    *Links                     `json:"links"`
	Meta     json.RawMessage            `json:"meta,omitempty"`
}

func (r *UpdateSaleDetailsRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateSaleDetailsRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateSaleDetailsRequest - returns UpdateSaleDetailsRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateSaleDetailsRequest(key Key) *UpdateSaleDetailsRequest {
	var updateSaleDetailsRequest UpdateSaleDetailsRequest
	if c.tryFindEntry(key, &updateSaleDetailsRequest) {
		return &updateSaleDetailsRequest
	}
	return nil
}
