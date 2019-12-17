/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type KycRecoveryOp struct {
	Key
	Attributes    *KycRecoveryOpAttributes    `json:"attributes,omitempty"`
	Relationships *KycRecoveryOpRelationships `json:"relationships,omitempty"`
}
type KycRecoveryOpResponse struct {
	Data     KycRecoveryOp `json:"data"`
	Included Included      `json:"included"`
}

type KycRecoveryOpListResponse struct {
	Data     []KycRecoveryOp `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *KycRecoveryOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *KycRecoveryOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustKycRecoveryOp - returns KycRecoveryOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustKycRecoveryOp(key Key) *KycRecoveryOp {
	var kYCRecoveryOp KycRecoveryOp
	if c.tryFindEntry(key, &kYCRecoveryOp) {
		return &kYCRecoveryOp
	}
	return nil
}
