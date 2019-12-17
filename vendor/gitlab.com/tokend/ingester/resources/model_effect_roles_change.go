/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type EffectRolesChange struct {
	Key
	Attributes EffectRolesChangeAttributes `json:"attributes"`
}
type EffectRolesChangeResponse struct {
	Data     EffectRolesChange `json:"data"`
	Included Included          `json:"included"`
}

type EffectRolesChangeListResponse struct {
	Data     []EffectRolesChange `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *EffectRolesChangeListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *EffectRolesChangeListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustEffectRolesChange - returns EffectRolesChange from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEffectRolesChange(key Key) *EffectRolesChange {
	var effectRolesChange EffectRolesChange
	if c.tryFindEntry(key, &effectRolesChange) {
		return &effectRolesChange
	}
	return nil
}
