/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Rule struct {
	Key
	Attributes RuleAttributes `json:"attributes"`
}
type RuleResponse struct {
	Data     Rule     `json:"data"`
	Included Included `json:"included"`
}

type RuleListResponse struct {
	Data     []Rule          `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *RuleListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *RuleListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustRule - returns Rule from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRule(key Key) *Rule {
	var rule Rule
	if c.tryFindEntry(key, &rule) {
		return &rule
	}
	return nil
}
