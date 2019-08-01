/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageKeyValueOp struct {
	Key
	Attributes ManageKeyValueOpAttributes `json:"attributes"`
}
type ManageKeyValueOpResponse struct {
	Data     ManageKeyValueOp `json:"data"`
	Included Included         `json:"included"`
}

type ManageKeyValueOpListResponse struct {
	Data     []ManageKeyValueOp `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustManageKeyValueOp - returns ManageKeyValueOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageKeyValueOp(key Key) *ManageKeyValueOp {
	var manageKeyValueOp ManageKeyValueOp
	if c.tryFindEntry(key, &manageKeyValueOp) {
		return &manageKeyValueOp
	}
	return nil
}
