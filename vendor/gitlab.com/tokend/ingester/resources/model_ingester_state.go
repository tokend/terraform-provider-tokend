/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type IngesterState struct {
	Key
	Attributes IngesterStateAttributes `json:"attributes"`
}
type IngesterStateResponse struct {
	Data     IngesterState `json:"data"`
	Included Included      `json:"included"`
}

type IngesterStateListResponse struct {
	Data     []IngesterState `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *IngesterStateListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *IngesterStateListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustIngesterState - returns IngesterState from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustIngesterState(key Key) *IngesterState {
	var ingesterState IngesterState
	if c.tryFindEntry(key, &ingesterState) {
		return &ingesterState
	}
	return nil
}
