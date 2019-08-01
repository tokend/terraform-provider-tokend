/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type AssetPair struct {
	Key
	Attributes    AssetPairAttributes    `json:"attributes"`
	Relationships AssetPairRelationships `json:"relationships"`
}
type AssetPairResponse struct {
	Data     AssetPair `json:"data"`
	Included Included  `json:"included"`
}

type AssetPairListResponse struct {
	Data     []AssetPair `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustAssetPair - returns AssetPair from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAssetPair(key Key) *AssetPair {
	var assetPair AssetPair
	if c.tryFindEntry(key, &assetPair) {
		return &assetPair
	}
	return nil
}
