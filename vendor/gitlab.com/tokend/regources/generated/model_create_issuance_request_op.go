/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateIssuanceRequestOp struct {
	Key
	Attributes    CreateIssuanceRequestOpAttributes    `json:"attributes"`
	Relationships CreateIssuanceRequestOpRelationships `json:"relationships"`
}
type CreateIssuanceRequestOpResponse struct {
	Data     CreateIssuanceRequestOp `json:"data"`
	Included Included                `json:"included"`
}

type CreateIssuanceRequestOpListResponse struct {
	Data     []CreateIssuanceRequestOp `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
}

// MustCreateIssuanceRequestOp - returns CreateIssuanceRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateIssuanceRequestOp(key Key) *CreateIssuanceRequestOp {
	var createIssuanceRequestOp CreateIssuanceRequestOp
	if c.tryFindEntry(key, &createIssuanceRequestOp) {
		return &createIssuanceRequestOp
	}
	return nil
}
