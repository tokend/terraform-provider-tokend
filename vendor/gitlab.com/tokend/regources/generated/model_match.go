/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Match struct {
	Key
	Attributes    MatchAttributes    `json:"attributes"`
	Relationships MatchRelationships `json:"relationships"`
}
type MatchResponse struct {
	Data     Match    `json:"data"`
	Included Included `json:"included"`
}

type MatchListResponse struct {
	Data     []Match         `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *MatchListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *MatchListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustMatch - returns Match from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustMatch(key Key) *Match {
	var match Match
	if c.tryFindEntry(key, &match) {
		return &match
	}
	return nil
}
