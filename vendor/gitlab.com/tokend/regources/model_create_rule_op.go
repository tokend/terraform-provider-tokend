/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateRuleOp struct {
	Key
	Attributes    CreateRuleOpAttributes    `json:"attributes"`
	Relationships CreateRuleOpRelationships `json:"relationships"`
}
type CreateRuleOpResponse struct {
	Data     CreateRuleOp `json:"data"`
	Included Included     `json:"included"`
}

type CreateRuleOpListResponse struct {
	Data     []CreateRuleOp  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *CreateRuleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateRuleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateRuleOp - returns CreateRuleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateRuleOp(key Key) *CreateRuleOp {
	var createRuleOp CreateRuleOp
	if c.tryFindEntry(key, &createRuleOp) {
		return &createRuleOp
	}
	return nil
}
