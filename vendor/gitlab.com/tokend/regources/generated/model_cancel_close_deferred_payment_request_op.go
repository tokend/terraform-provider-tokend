/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CancelCloseDeferredPaymentRequestOp struct {
	Key
	Relationships CancelCloseDeferredPaymentRequestOpRelationships `json:"relationships"`
}
type CancelCloseDeferredPaymentRequestOpResponse struct {
	Data     CancelCloseDeferredPaymentRequestOp `json:"data"`
	Included Included                            `json:"included"`
}

type CancelCloseDeferredPaymentRequestOpListResponse struct {
	Data     []CancelCloseDeferredPaymentRequestOp `json:"data"`
	Included Included                              `json:"included"`
	Links    *Links                                `json:"links"`
	Meta     json.RawMessage                       `json:"meta,omitempty"`
}

func (r *CancelCloseDeferredPaymentRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CancelCloseDeferredPaymentRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCancelCloseDeferredPaymentRequestOp - returns CancelCloseDeferredPaymentRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCancelCloseDeferredPaymentRequestOp(key Key) *CancelCloseDeferredPaymentRequestOp {
	var cancelCloseDeferredPaymentRequestOp CancelCloseDeferredPaymentRequestOp
	if c.tryFindEntry(key, &cancelCloseDeferredPaymentRequestOp) {
		return &cancelCloseDeferredPaymentRequestOp
	}
	return nil
}
