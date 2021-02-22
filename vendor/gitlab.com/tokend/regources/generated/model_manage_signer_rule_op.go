/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ManageSignerRuleOp struct {
	Key
	Attributes    *ManageSignerRuleOpAttributes    `json:"attributes,omitempty"`
	Relationships *ManageSignerRuleOpRelationships `json:"relationships,omitempty"`
}
type ManageSignerRuleOpResponse struct {
	Data     ManageSignerRuleOp `json:"data"`
	Included Included           `json:"included"`
}

type ManageSignerRuleOpListResponse struct {
	Data     []ManageSignerRuleOp `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *ManageSignerRuleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ManageSignerRuleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustManageSignerRuleOp - returns ManageSignerRuleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageSignerRuleOp(key Key) *ManageSignerRuleOp {
	var manageSignerRuleOp ManageSignerRuleOp
	if c.tryFindEntry(key, &manageSignerRuleOp) {
		return &manageSignerRuleOp
	}
	return nil
}
