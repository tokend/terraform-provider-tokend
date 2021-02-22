/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

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
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *ChangeRoleRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ChangeRoleRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
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
