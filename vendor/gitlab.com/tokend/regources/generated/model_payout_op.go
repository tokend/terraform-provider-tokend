/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type PayoutOp struct {
	Key
	Attributes    PayoutOpAttributes    `json:"attributes"`
	Relationships PayoutOpRelationships `json:"relationships"`
}
type PayoutOpResponse struct {
	Data     PayoutOp `json:"data"`
	Included Included `json:"included"`
}

type PayoutOpListResponse struct {
	Data     []PayoutOp      `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *PayoutOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *PayoutOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustPayoutOp - returns PayoutOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPayoutOp(key Key) *PayoutOp {
	var payoutOp PayoutOp
	if c.tryFindEntry(key, &payoutOp) {
		return &payoutOp
	}
	return nil
}
