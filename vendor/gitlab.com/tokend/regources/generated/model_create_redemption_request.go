/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateRedemptionRequest struct {
	Key
	Attributes    CreateRedemptionRequestAttributes    `json:"attributes"`
	Relationships CreateRedemptionRequestRelationships `json:"relationships"`
}
type CreateRedemptionRequestResponse struct {
	Data     CreateRedemptionRequest `json:"data"`
	Included Included                `json:"included"`
}

type CreateRedemptionRequestListResponse struct {
	Data     []CreateRedemptionRequest `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
	Meta     json.RawMessage           `json:"meta,omitempty"`
}

func (r *CreateRedemptionRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateRedemptionRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateRedemptionRequest - returns CreateRedemptionRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateRedemptionRequest(key Key) *CreateRedemptionRequest {
	var createRedemptionRequest CreateRedemptionRequest
	if c.tryFindEntry(key, &createRedemptionRequest) {
		return &createRedemptionRequest
	}
	return nil
}
