/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageSignerRoleOp struct {
	Key
	Attributes    *ManageSignerRoleOpAttributes   `json:"attributes,omitempty"`
	Relationships ManageSignerRoleOpRelationships `json:"relationships"`
}
type ManageSignerRoleOpResponse struct {
	Data     ManageSignerRoleOp `json:"data"`
	Included Included           `json:"included"`
}

type ManageSignerRoleOpListResponse struct {
	Data     []ManageSignerRoleOp `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
}

// MustManageSignerRoleOp - returns ManageSignerRoleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageSignerRoleOp(key Key) *ManageSignerRoleOp {
	var manageSignerRoleOp ManageSignerRoleOp
	if c.tryFindEntry(key, &manageSignerRoleOp) {
		return &manageSignerRoleOp
	}
	return nil
}
