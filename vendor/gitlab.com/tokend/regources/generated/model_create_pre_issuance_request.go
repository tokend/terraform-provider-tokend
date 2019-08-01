/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreatePreIssuanceRequest struct {
	Key
	Attributes    CreatePreIssuanceRequestAttributes    `json:"attributes"`
	Relationships CreatePreIssuanceRequestRelationships `json:"relationships"`
}
type CreatePreIssuanceRequestResponse struct {
	Data     CreatePreIssuanceRequest `json:"data"`
	Included Included                 `json:"included"`
}

type CreatePreIssuanceRequestListResponse struct {
	Data     []CreatePreIssuanceRequest `json:"data"`
	Included Included                   `json:"included"`
	Links    *Links                     `json:"links"`
}

// MustCreatePreIssuanceRequest - returns CreatePreIssuanceRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreatePreIssuanceRequest(key Key) *CreatePreIssuanceRequest {
	var createPreIssuanceRequest CreatePreIssuanceRequest
	if c.tryFindEntry(key, &createPreIssuanceRequest) {
		return &createPreIssuanceRequest
	}
	return nil
}
