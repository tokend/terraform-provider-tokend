/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type HorizonState struct {
	Key
	Attributes HorizonStateAttributes `json:"attributes"`
}
type HorizonStateResponse struct {
	Data     HorizonState `json:"data"`
	Included Included     `json:"included"`
}

type HorizonStateListResponse struct {
	Data     []HorizonState  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *HorizonStateListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *HorizonStateListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustHorizonState - returns HorizonState from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustHorizonState(key Key) *HorizonState {
	var horizonState HorizonState
	if c.tryFindEntry(key, &horizonState) {
		return &horizonState
	}
	return nil
}
