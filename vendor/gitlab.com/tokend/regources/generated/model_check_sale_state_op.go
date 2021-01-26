/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CheckSaleStateOp struct {
	Key
	Attributes    CheckSaleStateOpAttributes    `json:"attributes"`
	Relationships CheckSaleStateOpRelationships `json:"relationships"`
}
type CheckSaleStateOpResponse struct {
	Data     CheckSaleStateOp `json:"data"`
	Included Included         `json:"included"`
}

type CheckSaleStateOpListResponse struct {
	Data     []CheckSaleStateOp `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
	Meta     json.RawMessage    `json:"meta,omitempty"`
}

func (r *CheckSaleStateOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CheckSaleStateOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCheckSaleStateOp - returns CheckSaleStateOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCheckSaleStateOp(key Key) *CheckSaleStateOp {
	var checkSaleStateOp CheckSaleStateOp
	if c.tryFindEntry(key, &checkSaleStateOp) {
		return &checkSaleStateOp
	}
	return nil
}
