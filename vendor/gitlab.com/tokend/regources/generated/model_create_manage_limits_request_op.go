/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateManageLimitsRequestOp struct {
	Key
	Attributes    CreateManageLimitsRequestOpAttributes    `json:"attributes"`
	Relationships CreateManageLimitsRequestOpRelationships `json:"relationships"`
}
type CreateManageLimitsRequestOpResponse struct {
	Data     CreateManageLimitsRequestOp `json:"data"`
	Included Included                    `json:"included"`
}

type CreateManageLimitsRequestOpListResponse struct {
	Data     []CreateManageLimitsRequestOp `json:"data"`
	Included Included                      `json:"included"`
	Links    *Links                        `json:"links"`
}

// MustCreateManageLimitsRequestOp - returns CreateManageLimitsRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateManageLimitsRequestOp(key Key) *CreateManageLimitsRequestOp {
	var createManageLimitsRequestOp CreateManageLimitsRequestOp
	if c.tryFindEntry(key, &createManageLimitsRequestOp) {
		return &createManageLimitsRequestOp
	}
	return nil
}
