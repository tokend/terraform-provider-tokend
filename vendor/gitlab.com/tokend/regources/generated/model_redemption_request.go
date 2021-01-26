/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type RedemptionRequest struct {
	Key
	Attributes    RedemptionRequestAttributes    `json:"attributes"`
	Relationships RedemptionRequestRelationships `json:"relationships"`
}
type RedemptionRequestResponse struct {
	Data     RedemptionRequest `json:"data"`
	Included Included          `json:"included"`
}

type RedemptionRequestListResponse struct {
	Data     []RedemptionRequest `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *RedemptionRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *RedemptionRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustRedemptionRequest - returns RedemptionRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRedemptionRequest(key Key) *RedemptionRequest {
	var redemptionRequest RedemptionRequest
	if c.tryFindEntry(key, &redemptionRequest) {
		return &redemptionRequest
	}
	return nil
}
