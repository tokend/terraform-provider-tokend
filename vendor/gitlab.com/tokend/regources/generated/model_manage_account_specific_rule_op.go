/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ManageAccountSpecificRuleOp struct {
	Key
	Attributes    *ManageAccountSpecificRuleOpAttributes    `json:"attributes,omitempty"`
	Relationships *ManageAccountSpecificRuleOpRelationships `json:"relationships,omitempty"`
}
type ManageAccountSpecificRuleOpResponse struct {
	Data     ManageAccountSpecificRuleOp `json:"data"`
	Included Included                    `json:"included"`
}

type ManageAccountSpecificRuleOpListResponse struct {
	Data     []ManageAccountSpecificRuleOp `json:"data"`
	Included Included                      `json:"included"`
	Links    *Links                        `json:"links"`
	Meta     json.RawMessage               `json:"meta,omitempty"`
}

func (r *ManageAccountSpecificRuleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ManageAccountSpecificRuleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustManageAccountSpecificRuleOp - returns ManageAccountSpecificRuleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageAccountSpecificRuleOp(key Key) *ManageAccountSpecificRuleOp {
	var manageAccountSpecificRuleOp ManageAccountSpecificRuleOp
	if c.tryFindEntry(key, &manageAccountSpecificRuleOp) {
		return &manageAccountSpecificRuleOp
	}
	return nil
}
