/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ReviewableRequest struct {
	Key
	Attributes    ReviewableRequestAttributes    `json:"attributes"`
	Relationships ReviewableRequestRelationships `json:"relationships"`
}
type ReviewableRequestResponse struct {
	Data     ReviewableRequest `json:"data"`
	Included Included          `json:"included"`
}

type ReviewableRequestListResponse struct {
	Data     []ReviewableRequest `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustReviewableRequest - returns ReviewableRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustReviewableRequest(key Key) *ReviewableRequest {
	var reviewableRequest ReviewableRequest
	if c.tryFindEntry(key, &reviewableRequest) {
		return &reviewableRequest
	}
	return nil
}
