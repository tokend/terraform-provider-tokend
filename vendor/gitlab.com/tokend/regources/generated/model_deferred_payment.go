/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type DeferredPayment struct {
	Key
	Attributes    DeferredPaymentAttributes    `json:"attributes"`
	Relationships DeferredPaymentRelationships `json:"relationships"`
}
type DeferredPaymentResponse struct {
	Data     DeferredPayment `json:"data"`
	Included Included        `json:"included"`
}

type DeferredPaymentListResponse struct {
	Data     []DeferredPayment `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
	Meta     json.RawMessage   `json:"meta,omitempty"`
}

func (r *DeferredPaymentListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DeferredPaymentListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDeferredPayment - returns DeferredPayment from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDeferredPayment(key Key) *DeferredPayment {
	var deferredPayment DeferredPayment
	if c.tryFindEntry(key, &deferredPayment) {
		return &deferredPayment
	}
	return nil
}
