/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type SignerRule struct {
	Key
	Attributes    SignerRuleAttributes    `json:"attributes"`
	Relationships SignerRuleRelationships `json:"relationships"`
}
type SignerRuleResponse struct {
	Data     SignerRule `json:"data"`
	Included Included   `json:"included"`
}

type SignerRuleListResponse struct {
	Data     []SignerRule `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustSignerRule - returns SignerRule from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSignerRule(key Key) *SignerRule {
	var signerRule SignerRule
	if c.tryFindEntry(key, &signerRule) {
		return &signerRule
	}
	return nil
}
