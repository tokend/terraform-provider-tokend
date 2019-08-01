/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type SaleRequest struct {
	Key
	Attributes    *SaleRequestAttributes    `json:"attributes,omitempty"`
	Relationships *SaleRequestRelationships `json:"relationships,omitempty"`
}
type SaleRequestResponse struct {
	Data     SaleRequest `json:"data"`
	Included Included    `json:"included"`
}

type SaleRequestListResponse struct {
	Data     []SaleRequest `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustSaleRequest - returns SaleRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSaleRequest(key Key) *SaleRequest {
	var saleRequest SaleRequest
	if c.tryFindEntry(key, &saleRequest) {
		return &saleRequest
	}
	return nil
}
