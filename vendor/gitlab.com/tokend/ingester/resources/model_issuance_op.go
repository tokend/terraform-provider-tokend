/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type IssuanceOp struct {
	Key
	Attributes    IssuanceOpAttributes    `json:"attributes"`
	Relationships IssuanceOpRelationships `json:"relationships"`
}
type IssuanceOpResponse struct {
	Data     IssuanceOp `json:"data"`
	Included Included   `json:"included"`
}

type IssuanceOpListResponse struct {
	Data     []IssuanceOp    `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *IssuanceOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *IssuanceOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustIssuanceOp - returns IssuanceOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustIssuanceOp(key Key) *IssuanceOp {
	var issuanceOp IssuanceOp
	if c.tryFindEntry(key, &issuanceOp) {
		return &issuanceOp
	}
	return nil
}
