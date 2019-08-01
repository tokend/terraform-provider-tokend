/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageVoteOp struct {
	Key
	Attributes    ManageVoteOpAttributes    `json:"attributes"`
	Relationships ManageVoteOpRelationships `json:"relationships"`
}
type ManageVoteOpResponse struct {
	Data     ManageVoteOp `json:"data"`
	Included Included     `json:"included"`
}

type ManageVoteOpListResponse struct {
	Data     []ManageVoteOp `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustManageVoteOp - returns ManageVoteOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageVoteOp(key Key) *ManageVoteOp {
	var manageVoteOp ManageVoteOp
	if c.tryFindEntry(key, &manageVoteOp) {
		return &manageVoteOp
	}
	return nil
}
