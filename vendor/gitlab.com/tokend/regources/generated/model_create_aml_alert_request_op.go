/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateAmlAlertRequestOp struct {
	Key
	Attributes    CreateAmlAlertRequestOpAttributes    `json:"attributes"`
	Relationships CreateAmlAlertRequestOpRelationships `json:"relationships"`
}
type CreateAmlAlertRequestOpResponse struct {
	Data     CreateAmlAlertRequestOp `json:"data"`
	Included Included                `json:"included"`
}

type CreateAmlAlertRequestOpListResponse struct {
	Data     []CreateAmlAlertRequestOp `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
}

// MustCreateAmlAlertRequestOp - returns CreateAmlAlertRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAmlAlertRequestOp(key Key) *CreateAmlAlertRequestOp {
	var createAmlAlertRequestOp CreateAmlAlertRequestOp
	if c.tryFindEntry(key, &createAmlAlertRequestOp) {
		return &createAmlAlertRequestOp
	}
	return nil
}
