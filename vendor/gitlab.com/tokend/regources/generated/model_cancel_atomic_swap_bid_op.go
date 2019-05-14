/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CancelAtomicSwapBidOp struct {
	Key
	Relationships CancelAtomicSwapBidOpRelationships `json:"relationships"`
}
type CancelAtomicSwapBidOpResponse struct {
	Data     CancelAtomicSwapBidOp `json:"data"`
	Included Included              `json:"included"`
}

type CancelAtomicSwapBidOpsResponse struct {
	Data     []CancelAtomicSwapBidOp `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
}

// MustCancelAtomicSwapBidOp - returns CancelAtomicSwapBidOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCancelAtomicSwapBidOp(key Key) *CancelAtomicSwapBidOp {
	var cancelAtomicSwapBidOp CancelAtomicSwapBidOp
	if c.tryFindEntry(key, &cancelAtomicSwapBidOp) {
		return &cancelAtomicSwapBidOp
	}
	return nil
}
