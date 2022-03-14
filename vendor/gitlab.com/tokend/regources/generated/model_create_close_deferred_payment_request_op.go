/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateCloseDeferredPaymentRequestOp struct {
	Key
	Attributes    CreateCloseDeferredPaymentRequestOpAttributes    `json:"attributes"`
	Relationships CreateCloseDeferredPaymentRequestOpRelationships `json:"relationships"`
}
type CreateCloseDeferredPaymentRequestOpResponse struct {
	Data     CreateCloseDeferredPaymentRequestOp `json:"data"`
	Included Included                            `json:"included"`
}

type CreateCloseDeferredPaymentRequestOpListResponse struct {
	Data     []CreateCloseDeferredPaymentRequestOp `json:"data"`
	Included Included                              `json:"included"`
	Links    *Links                                `json:"links"`
	Meta     json.RawMessage                       `json:"meta,omitempty"`
}

func (r *CreateCloseDeferredPaymentRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateCloseDeferredPaymentRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateCloseDeferredPaymentRequestOp - returns CreateCloseDeferredPaymentRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateCloseDeferredPaymentRequestOp(key Key) *CreateCloseDeferredPaymentRequestOp {
	var createCloseDeferredPaymentRequestOp CreateCloseDeferredPaymentRequestOp
	if c.tryFindEntry(key, &createCloseDeferredPaymentRequestOp) {
		return &createCloseDeferredPaymentRequestOp
	}
	return nil
}
