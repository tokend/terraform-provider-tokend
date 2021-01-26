/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type UpdateAssetRequest struct {
	Key
	Attributes    UpdateAssetRequestAttributes    `json:"attributes"`
	Relationships UpdateAssetRequestRelationships `json:"relationships"`
}
type UpdateAssetRequestResponse struct {
	Data     UpdateAssetRequest `json:"data"`
	Included Included           `json:"included"`
}

type UpdateAssetRequestListResponse struct {
	Data     []UpdateAssetRequest `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
	Meta     json.RawMessage      `json:"meta,omitempty"`
}

func (r *UpdateAssetRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateAssetRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateAssetRequest - returns UpdateAssetRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateAssetRequest(key Key) *UpdateAssetRequest {
	var updateAssetRequest UpdateAssetRequest
	if c.tryFindEntry(key, &updateAssetRequest) {
		return &updateAssetRequest
	}
	return nil
}
