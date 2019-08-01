/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageAssetOp struct {
	Key
	Attributes    ManageAssetOpAttributes    `json:"attributes"`
	Relationships ManageAssetOpRelationships `json:"relationships"`
}
type ManageAssetOpResponse struct {
	Data     ManageAssetOp `json:"data"`
	Included Included      `json:"included"`
}

type ManageAssetOpListResponse struct {
	Data     []ManageAssetOp `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustManageAssetOp - returns ManageAssetOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageAssetOp(key Key) *ManageAssetOp {
	var manageAssetOp ManageAssetOp
	if c.tryFindEntry(key, &manageAssetOp) {
		return &manageAssetOp
	}
	return nil
}
