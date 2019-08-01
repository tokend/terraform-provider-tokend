/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageOfferOp struct {
	Key
	Attributes    ManageOfferOpAttributes    `json:"attributes"`
	Relationships ManageOfferOpRelationships `json:"relationships"`
}
type ManageOfferOpResponse struct {
	Data     ManageOfferOp `json:"data"`
	Included Included      `json:"included"`
}

type ManageOfferOpListResponse struct {
	Data     []ManageOfferOp `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustManageOfferOp - returns ManageOfferOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageOfferOp(key Key) *ManageOfferOp {
	var manageOfferOp ManageOfferOp
	if c.tryFindEntry(key, &manageOfferOp) {
		return &manageOfferOp
	}
	return nil
}
