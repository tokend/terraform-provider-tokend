/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type SignerRole struct {
	Key
	Attributes    SignerRoleAttributes    `json:"attributes"`
	Relationships SignerRoleRelationships `json:"relationships"`
}
type SignerRoleResponse struct {
	Data     SignerRole `json:"data"`
	Included Included   `json:"included"`
}

type SignerRoleListResponse struct {
	Data     []SignerRole    `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *SignerRoleListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *SignerRoleListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustSignerRole - returns SignerRole from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSignerRole(key Key) *SignerRole {
	var signerRole SignerRole
	if c.tryFindEntry(key, &signerRole) {
		return &signerRole
	}
	return nil
}
