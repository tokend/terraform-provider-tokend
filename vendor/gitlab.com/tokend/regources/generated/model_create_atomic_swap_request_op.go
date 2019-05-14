/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateAtomicSwapRequestOp struct {
	Key
	Attributes    CreateAtomicSwapRequestOpAttributes    `json:"attributes"`
	Relationships CreateAtomicSwapRequestOpRelationships `json:"relationships"`
}
type CreateAtomicSwapRequestOpResponse struct {
	Data     CreateAtomicSwapRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type CreateAtomicSwapRequestOpsResponse struct {
	Data     []CreateAtomicSwapRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
}

// MustCreateAtomicSwapRequestOp - returns CreateAtomicSwapRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAtomicSwapRequestOp(key Key) *CreateAtomicSwapRequestOp {
	var createAtomicSwapRequestOp CreateAtomicSwapRequestOp
	if c.tryFindEntry(key, &createAtomicSwapRequestOp) {
		return &createAtomicSwapRequestOp
	}
	return nil
}
