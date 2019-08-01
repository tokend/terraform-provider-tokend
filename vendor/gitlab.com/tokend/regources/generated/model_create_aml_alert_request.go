/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateAmlAlertRequest struct {
	Key
	Attributes    CreateAmlAlertRequestAttributes    `json:"attributes"`
	Relationships CreateAmlAlertRequestRelationships `json:"relationships"`
}
type CreateAmlAlertRequestResponse struct {
	Data     CreateAmlAlertRequest `json:"data"`
	Included Included              `json:"included"`
}

type CreateAmlAlertRequestListResponse struct {
	Data     []CreateAmlAlertRequest `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
}

// MustCreateAmlAlertRequest - returns CreateAmlAlertRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAmlAlertRequest(key Key) *CreateAmlAlertRequest {
	var createAmlAlertRequest CreateAmlAlertRequest
	if c.tryFindEntry(key, &createAmlAlertRequest) {
		return &createAmlAlertRequest
	}
	return nil
}
