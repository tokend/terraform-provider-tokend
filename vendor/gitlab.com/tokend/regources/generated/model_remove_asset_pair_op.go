/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type RemoveAssetPairOp struct {
	Key
	Relationships RemoveAssetPairOpRelationships `json:"relationships"`
}
type RemoveAssetPairOpResponse struct {
	Data     RemoveAssetPairOp `json:"data"`
	Included Included          `json:"included"`
}

type RemoveAssetPairOpListResponse struct {
	Data     []RemoveAssetPairOp `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustRemoveAssetPairOp - returns RemoveAssetPairOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRemoveAssetPairOp(key Key) *RemoveAssetPairOp {
	var removeAssetPairOp RemoveAssetPairOp
	if c.tryFindEntry(key, &removeAssetPairOp) {
		return &removeAssetPairOp
	}
	return nil
}
