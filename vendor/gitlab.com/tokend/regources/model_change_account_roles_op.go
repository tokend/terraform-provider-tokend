/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ChangeAccountRolesOp struct {
	Key
	Attributes    ChangeAccountRolesOpAttributes    `json:"attributes"`
	Relationships ChangeAccountRolesOpRelationships `json:"relationships"`
}
type ChangeAccountRolesOpResponse struct {
	Data     ChangeAccountRolesOp `json:"data"`
	Included Included             `json:"included"`
}

type ChangeAccountRolesOpListResponse struct {
	Data     []ChangeAccountRolesOp `json:"data"`
	Included Included               `json:"included"`
	Links    *Links                 `json:"links"`
	Meta     json.RawMessage        `json:"meta,omitempty"`
}

func (r *ChangeAccountRolesOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ChangeAccountRolesOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustChangeAccountRolesOp - returns ChangeAccountRolesOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustChangeAccountRolesOp(key Key) *ChangeAccountRolesOp {
	var changeAccountRolesOp ChangeAccountRolesOp
	if c.tryFindEntry(key, &changeAccountRolesOp) {
		return &changeAccountRolesOp
	}
	return nil
}
