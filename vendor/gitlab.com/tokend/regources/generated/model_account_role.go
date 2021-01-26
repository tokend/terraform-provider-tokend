/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type AccountRole struct {
	Key
	Attributes    AccountRoleAttributes    `json:"attributes"`
	Relationships AccountRoleRelationships `json:"relationships"`
}
type AccountRoleResponse struct {
	Data     AccountRole `json:"data"`
	Included Included    `json:"included"`
}

type AccountRoleListResponse struct {
	Data     []AccountRole   `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *AccountRoleListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *AccountRoleListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustAccountRole - returns AccountRole from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAccountRole(key Key) *AccountRole {
	var accountRole AccountRole
	if c.tryFindEntry(key, &accountRole) {
		return &accountRole
	}
	return nil
}
