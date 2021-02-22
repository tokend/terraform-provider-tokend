/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreatePreIssuanceRequestOp struct {
	Key
	Attributes    CreatePreIssuanceRequestOpAttributes    `json:"attributes"`
	Relationships CreatePreIssuanceRequestOpRelationships `json:"relationships"`
}
type CreatePreIssuanceRequestOpResponse struct {
	Data     CreatePreIssuanceRequestOp `json:"data"`
	Included Included                   `json:"included"`
}

type CreatePreIssuanceRequestOpListResponse struct {
	Data     []CreatePreIssuanceRequestOp `json:"data"`
	Included Included                     `json:"included"`
	Links    *Links                       `json:"links"`
	Meta     json.RawMessage              `json:"meta,omitempty"`
}

func (r *CreatePreIssuanceRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreatePreIssuanceRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreatePreIssuanceRequestOp - returns CreatePreIssuanceRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreatePreIssuanceRequestOp(key Key) *CreatePreIssuanceRequestOp {
	var createPreIssuanceRequestOp CreatePreIssuanceRequestOp
	if c.tryFindEntry(key, &createPreIssuanceRequestOp) {
		return &createPreIssuanceRequestOp
	}
	return nil
}
