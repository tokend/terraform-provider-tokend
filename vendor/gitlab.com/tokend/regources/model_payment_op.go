/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type PaymentOp struct {
	Key
	Attributes    PaymentOpAttributes    `json:"attributes"`
	Relationships PaymentOpRelationships `json:"relationships"`
}
type PaymentOpResponse struct {
	Data     PaymentOp `json:"data"`
	Included Included  `json:"included"`
}

type PaymentOpListResponse struct {
	Data     []PaymentOp     `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *PaymentOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *PaymentOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustPaymentOp - returns PaymentOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPaymentOp(key Key) *PaymentOp {
	var paymentOp PaymentOp
	if c.tryFindEntry(key, &paymentOp) {
		return &paymentOp
	}
	return nil
}
