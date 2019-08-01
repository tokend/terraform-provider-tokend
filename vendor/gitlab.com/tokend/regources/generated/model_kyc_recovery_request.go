/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type KycRecoveryRequest struct {
	Key
	Attributes    KycRecoveryRequestAttributes    `json:"attributes"`
	Relationships KycRecoveryRequestRelationships `json:"relationships"`
}
type KycRecoveryRequestResponse struct {
	Data     KycRecoveryRequest `json:"data"`
	Included Included           `json:"included"`
}

type KycRecoveryRequestListResponse struct {
	Data     []KycRecoveryRequest `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
}

// MustKycRecoveryRequest - returns KycRecoveryRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustKycRecoveryRequest(key Key) *KycRecoveryRequest {
	var kYCRecoveryRequest KycRecoveryRequest
	if c.tryFindEntry(key, &kYCRecoveryRequest) {
		return &kYCRecoveryRequest
	}
	return nil
}
