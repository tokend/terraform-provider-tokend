/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ManageExternalSystemAccountIdPoolEntryOp struct {
	Key
	Attributes ManageExternalSystemAccountIdPoolEntryOpAttributes `json:"attributes"`
}
type ManageExternalSystemAccountIdPoolEntryOpResponse struct {
	Data     ManageExternalSystemAccountIdPoolEntryOp `json:"data"`
	Included Included                                 `json:"included"`
}

type ManageExternalSystemAccountIdPoolEntryOpListResponse struct {
	Data     []ManageExternalSystemAccountIdPoolEntryOp `json:"data"`
	Included Included                                   `json:"included"`
	Links    *Links                                     `json:"links"`
}

// MustManageExternalSystemAccountIdPoolEntryOp - returns ManageExternalSystemAccountIdPoolEntryOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageExternalSystemAccountIdPoolEntryOp(key Key) *ManageExternalSystemAccountIdPoolEntryOp {
	var manageExternalSystemAccountIDPoolEntryOp ManageExternalSystemAccountIdPoolEntryOp
	if c.tryFindEntry(key, &manageExternalSystemAccountIDPoolEntryOp) {
		return &manageExternalSystemAccountIDPoolEntryOp
	}
	return nil
}
