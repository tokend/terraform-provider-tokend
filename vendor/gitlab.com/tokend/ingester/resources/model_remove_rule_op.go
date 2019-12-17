/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type RemoveRuleOp struct {
	Key
	Relationships RemoveRuleOpRelationships `json:"relationships"`
}
type RemoveRuleOpResponse struct {
	Data     RemoveRuleOp `json:"data"`
	Included Included     `json:"included"`
}

type RemoveRuleOpListResponse struct {
	Data     []RemoveRuleOp  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *RemoveRuleOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *RemoveRuleOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustRemoveRuleOp - returns RemoveRuleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRemoveRuleOp(key Key) *RemoveRuleOp {
	var removeRuleOp RemoveRuleOp
	if c.tryFindEntry(key, &removeRuleOp) {
		return &removeRuleOp
	}
	return nil
}
