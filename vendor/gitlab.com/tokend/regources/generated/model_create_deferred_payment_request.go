/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateDeferredPaymentRequest struct {
	Key
	Attributes    CreateDeferredPaymentRequestAttributes    `json:"attributes"`
	Relationships CreateDeferredPaymentRequestRelationships `json:"relationships"`
}
type CreateDeferredPaymentRequestResponse struct {
	Data     CreateDeferredPaymentRequest `json:"data"`
	Included Included                     `json:"included"`
}

type CreateDeferredPaymentRequestListResponse struct {
	Data     []CreateDeferredPaymentRequest `json:"data"`
	Included Included                       `json:"included"`
	Links    *Links                         `json:"links"`
	Meta     json.RawMessage                `json:"meta,omitempty"`
}

func (r *CreateDeferredPaymentRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateDeferredPaymentRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateDeferredPaymentRequest - returns CreateDeferredPaymentRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateDeferredPaymentRequest(key Key) *CreateDeferredPaymentRequest {
	var createDeferredPaymentRequest CreateDeferredPaymentRequest
	if c.tryFindEntry(key, &createDeferredPaymentRequest) {
		return &createDeferredPaymentRequest
	}
	return nil
}
