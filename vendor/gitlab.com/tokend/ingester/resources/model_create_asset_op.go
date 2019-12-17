/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type CreateAssetOp struct {
	Key
	Attributes    CreateAssetOpAttributes    `json:"attributes"`
	Relationships CreateAssetOpRelationships `json:"relationships"`
}
type CreateAssetOpResponse struct {
	Data     CreateAssetOp `json:"data"`
	Included Included      `json:"included"`
}

type CreateAssetOpListResponse struct {
	Data     []CreateAssetOp `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *CreateAssetOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateAssetOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateAssetOp - returns CreateAssetOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateAssetOp(key Key) *CreateAssetOp {
	var createAssetOp CreateAssetOp
	if c.tryFindEntry(key, &createAssetOp) {
		return &createAssetOp
	}
	return nil
}
