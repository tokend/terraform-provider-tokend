/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManagePollOp struct {
	Key
	Attributes    ManagePollOpAttributes    `json:"attributes"`
	Relationships ManagePollOpRelationships `json:"relationships"`
}
type ManagePollOpResponse struct {
	Data     ManagePollOp `json:"data"`
	Included Included     `json:"included"`
}

type ManagePollOpListResponse struct {
	Data     []ManagePollOp `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustManagePollOp - returns ManagePollOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManagePollOp(key Key) *ManagePollOp {
	var managePollOp ManagePollOp
	if c.tryFindEntry(key, &managePollOp) {
		return &managePollOp
	}
	return nil
}
