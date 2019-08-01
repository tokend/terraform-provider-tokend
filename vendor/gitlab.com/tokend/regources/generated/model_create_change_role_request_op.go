/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateChangeRoleRequestOp struct {
	Key
	Attributes    CreateChangeRoleRequestOpAttributes    `json:"attributes"`
	Relationships CreateChangeRoleRequestOpRelationships `json:"relationships"`
}
type CreateChangeRoleRequestOpResponse struct {
	Data     CreateChangeRoleRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type CreateChangeRoleRequestOpListResponse struct {
	Data     []CreateChangeRoleRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
}

// MustCreateChangeRoleRequestOp - returns CreateChangeRoleRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateChangeRoleRequestOp(key Key) *CreateChangeRoleRequestOp {
	var createChangeRoleRequestOp CreateChangeRoleRequestOp
	if c.tryFindEntry(key, &createChangeRoleRequestOp) {
		return &createChangeRoleRequestOp
	}
	return nil
}
