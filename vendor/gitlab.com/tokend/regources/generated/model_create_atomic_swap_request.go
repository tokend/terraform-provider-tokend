/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateAtomicSwapRequest struct {
	Key
	Attributes    CreateAtomicSwapRequestAttributes    `json:"attributes"`
	Relationships CreateAtomicSwapRequestRelationships `json:"relationships"`
}
type CreateAtomicSwapRequestResponse struct {
	Data     CreateAtomicSwapRequest `json:"data"`
	Included Included                `json:"included"`
}

type CreateAtomicSwapRequestsResponse struct {
	Data     []CreateAtomicSwapRequest `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
}

// MustCreateAtomicSwapRequest - returns CreateAtomicSwapRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAtomicSwapRequest(key Key) *CreateAtomicSwapRequest {
	var createAtomicSwapRequest CreateAtomicSwapRequest
	if c.tryFindEntry(key, &createAtomicSwapRequest) {
		return &createAtomicSwapRequest
	}
	return nil
}
