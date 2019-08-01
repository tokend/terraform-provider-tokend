/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type UpdateLimitsRequest struct {
	Key
	Attributes UpdateLimitsRequestAttributes `json:"attributes"`
}
type UpdateLimitsRequestResponse struct {
	Data     UpdateLimitsRequest `json:"data"`
	Included Included            `json:"included"`
}

type UpdateLimitsRequestListResponse struct {
	Data     []UpdateLimitsRequest `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustUpdateLimitsRequest - returns UpdateLimitsRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateLimitsRequest(key Key) *UpdateLimitsRequest {
	var updateLimitsRequest UpdateLimitsRequest
	if c.tryFindEntry(key, &updateLimitsRequest) {
		return &updateLimitsRequest
	}
	return nil
}
