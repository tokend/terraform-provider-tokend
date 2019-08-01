/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

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
	Data     []Match  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
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
