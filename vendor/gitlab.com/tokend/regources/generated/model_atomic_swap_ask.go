/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type AtomicSwapAsk struct {
	Key
	Attributes    AtomicSwapAskAttributes    `json:"attributes"`
	Relationships AtomicSwapAskRelationships `json:"relationships"`
}
type AtomicSwapAskResponse struct {
	Data     AtomicSwapAsk `json:"data"`
	Included Included      `json:"included"`
}

type AtomicSwapAskListResponse struct {
	Data     []AtomicSwapAsk `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustAtomicSwapAsk - returns AtomicSwapAsk from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAtomicSwapAsk(key Key) *AtomicSwapAsk {
	var atomicSwapAsk AtomicSwapAsk
	if c.tryFindEntry(key, &atomicSwapAsk) {
		return &atomicSwapAsk
	}
	return nil
}
