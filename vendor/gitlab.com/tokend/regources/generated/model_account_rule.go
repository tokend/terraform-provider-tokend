/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type AccountRule struct {
	Key
	Attributes AccountRuleAttributes `json:"attributes"`
}
type AccountRuleResponse struct {
	Data     AccountRule `json:"data"`
	Included Included    `json:"included"`
}

type AccountRuleListResponse struct {
	Data     []AccountRule   `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *AccountRuleListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *AccountRuleListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustAccountRule - returns AccountRule from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAccountRule(key Key) *AccountRule {
	var accountRule AccountRule
	if c.tryFindEntry(key, &accountRule) {
		return &accountRule
	}
	return nil
}
