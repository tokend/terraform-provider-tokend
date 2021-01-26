/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateAmlAlertRequest struct {
	Key
	Attributes    CreateAmlAlertRequestAttributes    `json:"attributes"`
	Relationships CreateAmlAlertRequestRelationships `json:"relationships"`
}
type CreateAmlAlertRequestResponse struct {
	Data     CreateAmlAlertRequest `json:"data"`
	Included Included              `json:"included"`
}

type CreateAmlAlertRequestListResponse struct {
	Data     []CreateAmlAlertRequest `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
	Meta     json.RawMessage         `json:"meta,omitempty"`
}

func (r *CreateAmlAlertRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateAmlAlertRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateAmlAlertRequest - returns CreateAmlAlertRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAmlAlertRequest(key Key) *CreateAmlAlertRequest {
	var createAmlAlertRequest CreateAmlAlertRequest
	if c.tryFindEntry(key, &createAmlAlertRequest) {
		return &createAmlAlertRequest
	}
	return nil
}
