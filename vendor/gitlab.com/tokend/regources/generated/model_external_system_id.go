/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ExternalSystemId struct {
	Key
	Attributes    ExternalSystemIdAttributes    `json:"attributes"`
	Relationships ExternalSystemIdRelationships `json:"relationships"`
}
type ExternalSystemIdResponse struct {
	Data     ExternalSystemId `json:"data"`
	Included Included         `json:"included"`
}

type ExternalSystemIdListResponse struct {
	Data     []ExternalSystemId `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustExternalSystemId - returns ExternalSystemId from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustExternalSystemId(key Key) *ExternalSystemId {
	var externalSystemID ExternalSystemId
	if c.tryFindEntry(key, &externalSystemID) {
		return &externalSystemID
	}
	return nil
}
