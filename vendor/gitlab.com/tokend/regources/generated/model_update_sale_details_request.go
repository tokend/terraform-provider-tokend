/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type UpdateSaleDetailsRequest struct {
	Key
	Attributes    UpdateSaleDetailsRequestAttributes    `json:"attributes"`
	Relationships UpdateSaleDetailsRequestRelationships `json:"relationships"`
}
type UpdateSaleDetailsRequestResponse struct {
	Data     UpdateSaleDetailsRequest `json:"data"`
	Included Included                 `json:"included"`
}

type UpdateSaleDetailsRequestListResponse struct {
	Data     []UpdateSaleDetailsRequest `json:"data"`
	Included Included                   `json:"included"`
	Links    *Links                     `json:"links"`
}

// MustUpdateSaleDetailsRequest - returns UpdateSaleDetailsRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateSaleDetailsRequest(key Key) *UpdateSaleDetailsRequest {
	var updateSaleDetailsRequest UpdateSaleDetailsRequest
	if c.tryFindEntry(key, &updateSaleDetailsRequest) {
		return &updateSaleDetailsRequest
	}
	return nil
}
