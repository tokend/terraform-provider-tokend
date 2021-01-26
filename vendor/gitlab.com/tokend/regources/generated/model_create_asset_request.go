/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateAssetRequest struct {
	Key
	Attributes    CreateAssetRequestAttributes    `json:"attributes"`
	Relationships CreateAssetRequestRelationships `json:"relationships"`
}
type CreateAssetRequestResponse struct {
	Data     CreateAssetRequest `json:"data"`
	Included Included           `json:"included"`
}

type CreateAssetRequestListResponse struct {
	Data     []CreateAssetRequest `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *CreateAssetRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateAssetRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateAssetRequest - returns CreateAssetRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAssetRequest(key Key) *CreateAssetRequest {
	var createAssetRequest CreateAssetRequest
	if c.tryFindEntry(key, &createAssetRequest) {
		return &createAssetRequest
	}
	return nil
}
