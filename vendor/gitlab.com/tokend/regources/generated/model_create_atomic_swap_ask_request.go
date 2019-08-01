/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateAtomicSwapAskRequest struct {
	Key
	Attributes    CreateAtomicSwapAskRequestAttributes    `json:"attributes"`
	Relationships CreateAtomicSwapAskRequestRelationships `json:"relationships"`
}
type CreateAtomicSwapAskRequestResponse struct {
	Data     CreateAtomicSwapAskRequest `json:"data"`
	Included Included                   `json:"included"`
}

type CreateAtomicSwapAskRequestListResponse struct {
	Data     []CreateAtomicSwapAskRequest `json:"data"`
	Included Included                     `json:"included"`
	Links    *Links                       `json:"links"`
}

// MustCreateAtomicSwapAskRequest - returns CreateAtomicSwapAskRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAtomicSwapAskRequest(key Key) *CreateAtomicSwapAskRequest {
	var createAtomicSwapAskRequest CreateAtomicSwapAskRequest
	if c.tryFindEntry(key, &createAtomicSwapAskRequest) {
		return &createAtomicSwapAskRequest
	}
	return nil
}
