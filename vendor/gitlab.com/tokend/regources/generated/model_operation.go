/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type Operation struct {
	Key
	Attributes    OperationAttributes    `json:"attributes"`
	Relationships OperationRelationships `json:"relationships"`
}
type OperationResponse struct {
	Data     Operation `json:"data"`
	Included Included  `json:"included"`
}

type OperationListResponse struct {
	Data     []Operation `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustOperation - returns Operation from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustOperation(key Key) *Operation {
	var operation Operation
	if c.tryFindEntry(key, &operation) {
		return &operation
	}
	return nil
}
