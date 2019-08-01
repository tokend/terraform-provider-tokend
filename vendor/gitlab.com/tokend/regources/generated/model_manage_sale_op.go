/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageSaleOp struct {
	Key
	Attributes ManageSaleOpAttributes `json:"attributes"`
}
type ManageSaleOpResponse struct {
	Data     ManageSaleOp `json:"data"`
	Included Included     `json:"included"`
}

type ManageSaleOpListResponse struct {
	Data     []ManageSaleOp `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustManageSaleOp - returns ManageSaleOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageSaleOp(key Key) *ManageSaleOp {
	var manageSaleOp ManageSaleOp
	if c.tryFindEntry(key, &manageSaleOp) {
		return &manageSaleOp
	}
	return nil
}
