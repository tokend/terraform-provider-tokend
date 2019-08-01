/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ChangeRoleRequest struct {
	Key
	Attributes    ChangeRoleRequestAttributes    `json:"attributes"`
	Relationships ChangeRoleRequestRelationships `json:"relationships"`
}
type ChangeRoleRequestResponse struct {
	Data     ChangeRoleRequest `json:"data"`
	Included Included          `json:"included"`
}

type ChangeRoleRequestListResponse struct {
	Data     []ChangeRoleRequest `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustChangeRoleRequest - returns ChangeRoleRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustChangeRoleRequest(key Key) *ChangeRoleRequest {
	var changeRoleRequest ChangeRoleRequest
	if c.tryFindEntry(key, &changeRoleRequest) {
		return &changeRoleRequest
	}
	return nil
}
