/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Asset struct {
	Key
	Attributes    AssetAttributes    `json:"attributes"`
	Relationships AssetRelationships `json:"relationships"`
}
type AssetResponse struct {
	Data     Asset    `json:"data"`
	Included Included `json:"included"`
}

type AssetListResponse struct {
	Data     []Asset         `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *AssetListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *AssetListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustAsset - returns Asset from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAsset(key Key) *Asset {
	var asset Asset
	if c.tryFindEntry(key, &asset) {
		return &asset
	}
	return nil
}
