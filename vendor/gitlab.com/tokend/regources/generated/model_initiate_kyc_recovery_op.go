/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type InitiateKycRecoveryOp struct {
	Key
	Attributes    *InitiateKycRecoveryOpAttributes    `json:"attributes,omitempty"`
	Relationships *InitiateKycRecoveryOpRelationships `json:"relationships,omitempty"`
}
type InitiateKycRecoveryOpResponse struct {
	Data     InitiateKycRecoveryOp `json:"data"`
	Included Included              `json:"included"`
}

type InitiateKycRecoveryOpListResponse struct {
	Data     []InitiateKycRecoveryOp `json:"data"`
	Included Included                `json:"included"`
	Links    *Links                  `json:"links"`
}

// MustInitiateKycRecoveryOp - returns InitiateKycRecoveryOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustInitiateKycRecoveryOp(key Key) *InitiateKycRecoveryOp {
	var initiateKYCRecoveryOp InitiateKycRecoveryOp
	if c.tryFindEntry(key, &initiateKYCRecoveryOp) {
		return &initiateKYCRecoveryOp
	}
	return nil
}
