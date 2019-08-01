/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageSignerOp struct {
	Key
	Attributes    *ManageSignerOpAttributes    `json:"attributes,omitempty"`
	Relationships *ManageSignerOpRelationships `json:"relationships,omitempty"`
}
type ManageSignerOpResponse struct {
	Data     ManageSignerOp `json:"data"`
	Included Included       `json:"included"`
}

type ManageSignerOpListResponse struct {
	Data     []ManageSignerOp `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustManageSignerOp - returns ManageSignerOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageSignerOp(key Key) *ManageSignerOp {
	var manageSignerOp ManageSignerOp
	if c.tryFindEntry(key, &manageSignerOp) {
		return &manageSignerOp
	}
	return nil
}
