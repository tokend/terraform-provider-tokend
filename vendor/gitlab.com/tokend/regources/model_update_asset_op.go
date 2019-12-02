/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type UpdateAssetOp struct {
	Key
	Attributes    UpdateAssetOpAttributes    `json:"attributes"`
	Relationships UpdateAssetOpRelationships `json:"relationships"`
}
type UpdateAssetOpResponse struct {
	Data     UpdateAssetOp `json:"data"`
	Included Included      `json:"included"`
}

type UpdateAssetOpListResponse struct {
	Data     []UpdateAssetOp `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *UpdateAssetOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateAssetOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateAssetOp - returns UpdateAssetOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateAssetOp(key Key) *UpdateAssetOp {
	var updateAssetOp UpdateAssetOp
	if c.tryFindEntry(key, &updateAssetOp) {
		return &updateAssetOp
	}
	return nil
}
