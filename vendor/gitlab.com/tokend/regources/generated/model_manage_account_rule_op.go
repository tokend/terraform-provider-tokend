/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageAccountRuleOp struct {
	Key
	Attributes    *ManageAccountRuleOpAttributes    `json:"attributes,omitempty"`
	Relationships *ManageAccountRuleOpRelationships `json:"relationships,omitempty"`
}
type ManageAccountRuleOpResponse struct {
	Data     ManageAccountRuleOp `json:"data"`
	Included Included            `json:"included"`
}

type ManageAccountRuleOpListResponse struct {
	Data     []ManageAccountRuleOp `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustManageAccountRuleOp - returns ManageAccountRuleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageAccountRuleOp(key Key) *ManageAccountRuleOp {
	var manageAccountRuleOp ManageAccountRuleOp
	if c.tryFindEntry(key, &manageAccountRuleOp) {
		return &manageAccountRuleOp
	}
	return nil
}
