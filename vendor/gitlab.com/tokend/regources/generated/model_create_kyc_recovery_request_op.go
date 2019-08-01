/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CreateKycRecoveryRequestOp struct {
	Key
	Attributes    *CreateKycRecoveryRequestOpAttributes    `json:"attributes,omitempty"`
	Relationships *CreateKycRecoveryRequestOpRelationships `json:"relationships,omitempty"`
}
type CreateKycRecoveryRequestOpResponse struct {
	Data     CreateKycRecoveryRequestOp `json:"data"`
	Included Included                   `json:"included"`
}

type CreateKycRecoveryRequestOpListResponse struct {
	Data     []CreateKycRecoveryRequestOp `json:"data"`
	Included Included                     `json:"included"`
	Links    *Links                       `json:"links"`
}

// MustCreateKycRecoveryRequestOp - returns CreateKycRecoveryRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateKycRecoveryRequestOp(key Key) *CreateKycRecoveryRequestOp {
	var createKYCRecoveryRequestOp CreateKycRecoveryRequestOp
	if c.tryFindEntry(key, &createKYCRecoveryRequestOp) {
		return &createKYCRecoveryRequestOp
	}
	return nil
}
