/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type EffectMatched struct {
	Key
	Attributes EffectMatchedAttributes `json:"attributes"`
}
type EffectMatchedResponse struct {
	Data     EffectMatched `json:"data"`
	Included Included      `json:"included"`
}

type EffectMatchedListResponse struct {
	Data     []EffectMatched `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustEffectMatched - returns EffectMatched from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustEffectMatched(key Key) *EffectMatched {
	var effectMatched EffectMatched
	if c.tryFindEntry(key, &effectMatched) {
		return &effectMatched
	}
	return nil
}
