/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateAtomicSwapBidRequest struct {
	Key
	Attributes    CreateAtomicSwapBidRequestAttributes    `json:"attributes"`
	Relationships CreateAtomicSwapBidRequestRelationships `json:"relationships"`
}
type CreateAtomicSwapBidRequestResponse struct {
	Data     CreateAtomicSwapBidRequest `json:"data"`
	Included Included                   `json:"included"`
}

type CreateAtomicSwapBidRequestListResponse struct {
	Data     []CreateAtomicSwapBidRequest `json:"data"`
	Included Included                     `json:"included"`
	Links    *Links                       `json:"links"`
}

// MustCreateAtomicSwapBidRequest - returns CreateAtomicSwapBidRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAtomicSwapBidRequest(key Key) *CreateAtomicSwapBidRequest {
	var createAtomicSwapBidRequest CreateAtomicSwapBidRequest
	if c.tryFindEntry(key, &createAtomicSwapBidRequest) {
		return &createAtomicSwapBidRequest
	}
	return nil
}
