/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CloseDeferredPaymentRequest struct {
	Key
	Attributes    CloseDeferredPaymentRequestAttributes    `json:"attributes"`
	Relationships CloseDeferredPaymentRequestRelationships `json:"relationships"`
}
type CloseDeferredPaymentRequestResponse struct {
	Data     CloseDeferredPaymentRequest `json:"data"`
	Included Included                    `json:"included"`
}

type CloseDeferredPaymentRequestListResponse struct {
	Data     []CloseDeferredPaymentRequest `json:"data"`
	Included Included                      `json:"included"`
	Links    *Links                        `json:"links"`
	Meta     json.RawMessage               `json:"meta,omitempty"`
}

func (r *CloseDeferredPaymentRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CloseDeferredPaymentRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCloseDeferredPaymentRequest - returns CloseDeferredPaymentRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCloseDeferredPaymentRequest(key Key) *CloseDeferredPaymentRequest {
	var closeDeferredPaymentRequest CloseDeferredPaymentRequest
	if c.tryFindEntry(key, &closeDeferredPaymentRequest) {
		return &closeDeferredPaymentRequest
	}
	return nil
}
