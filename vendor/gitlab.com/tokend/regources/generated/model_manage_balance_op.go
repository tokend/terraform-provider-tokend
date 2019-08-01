/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageBalanceOp struct {
	Key
	Attributes    ManageBalanceOpAttributes    `json:"attributes"`
	Relationships ManageBalanceOpRelationships `json:"relationships"`
}
type ManageBalanceOpResponse struct {
	Data     ManageBalanceOp `json:"data"`
	Included Included        `json:"included"`
}

type ManageBalanceOpListResponse struct {
	Data     []ManageBalanceOp `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustManageBalanceOp - returns ManageBalanceOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageBalanceOp(key Key) *ManageBalanceOp {
	var manageBalanceOp ManageBalanceOp
	if c.tryFindEntry(key, &manageBalanceOp) {
		return &manageBalanceOp
	}
	return nil
}
