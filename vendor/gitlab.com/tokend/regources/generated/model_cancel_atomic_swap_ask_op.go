/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CancelAtomicSwapAskOp struct {
	Key
	Relationships CancelAtomicSwapAskOpRelationships `json:"relationships"`
}
type CancelAtomicSwapAskOpResponse struct {
	Data     CancelAtomicSwapAskOp `json:"data"`
	Included Included              `json:"included"`
}

type CancelAtomicSwapAskOpListResponse struct {
	Data     []CancelAtomicSwapAskOp `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
}

// MustCancelAtomicSwapAskOp - returns CancelAtomicSwapAskOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCancelAtomicSwapAskOp(key Key) *CancelAtomicSwapAskOp {
	var cancelAtomicSwapAskOp CancelAtomicSwapAskOp
	if c.tryFindEntry(key, &cancelAtomicSwapAskOp) {
		return &cancelAtomicSwapAskOp
	}
	return nil
}
