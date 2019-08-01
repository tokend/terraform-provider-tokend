/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateIssuanceRequest struct {
	Key
	Attributes    CreateIssuanceRequestAttributes    `json:"attributes"`
	Relationships CreateIssuanceRequestRelationships `json:"relationships"`
}
type CreateIssuanceRequestResponse struct {
	Data     CreateIssuanceRequest `json:"data"`
	Included Included              `json:"included"`
}

type CreateIssuanceRequestListResponse struct {
	Data     []CreateIssuanceRequest `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
}

// MustCreateIssuanceRequest - returns CreateIssuanceRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateIssuanceRequest(key Key) *CreateIssuanceRequest {
	var createIssuanceRequest CreateIssuanceRequest
	if c.tryFindEntry(key, &createIssuanceRequest) {
		return &createIssuanceRequest
	}
	return nil
}
