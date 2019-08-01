/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type BindExternalSystemAccountIdOp struct {
	Key
	Attributes BindExternalSystemAccountIdOpAttributes `json:"attributes"`
}
type BindExternalSystemAccountIdOpResponse struct {
	Data     BindExternalSystemAccountIdOp `json:"data"`
	Included Included                      `json:"included"`
}

type BindExternalSystemAccountIdOpListResponse struct {
	Data     []BindExternalSystemAccountIdOp `json:"data"`
	Included Included                        `json:"included"`
	Links    *Links                          `json:"links"`
}

// MustBindExternalSystemAccountIdOp - returns BindExternalSystemAccountIdOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBindExternalSystemAccountIdOp(key Key) *BindExternalSystemAccountIdOp {
	var bindExternalSystemAccountIdOp BindExternalSystemAccountIdOp
	if c.tryFindEntry(key, &bindExternalSystemAccountIdOp) {
		return &bindExternalSystemAccountIdOp
	}
	return nil
}
