/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type AccountSpecificRule struct {
	Key
	Attributes AccountSpecificRuleAttributes `json:"attributes"`
}
type AccountSpecificRuleResponse struct {
	Data     AccountSpecificRule `json:"data"`
	Included Included            `json:"included"`
}

type AccountSpecificRuleListResponse struct {
	Data     []AccountSpecificRule `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustAccountSpecificRule - returns AccountSpecificRule from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAccountSpecificRule(key Key) *AccountSpecificRule {
	var accountSpecificRule AccountSpecificRule
	if c.tryFindEntry(key, &accountSpecificRule) {
		return &accountSpecificRule
	}
	return nil
}
