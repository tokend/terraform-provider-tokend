/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type UpdateRuleOp struct {
	Key
	Attributes    UpdateRuleOpAttributes    `json:"attributes"`
	Relationships UpdateRuleOpRelationships `json:"relationships"`
}
type UpdateRuleOpResponse struct {
	Data     UpdateRuleOp `json:"data"`
	Included Included     `json:"included"`
}

type UpdateRuleOpListResponse struct {
	Data     []UpdateRuleOp  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *UpdateRuleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateRuleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateRuleOp - returns UpdateRuleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateRuleOp(key Key) *UpdateRuleOp {
	var updateRuleOp UpdateRuleOp
	if c.tryFindEntry(key, &updateRuleOp) {
		return &updateRuleOp
	}
	return nil
}
