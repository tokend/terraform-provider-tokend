/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ConvertedBalancesCollection struct {
	Key
	Relationships ConvertedBalancesCollectionRelationships `json:"relationships"`
}
type ConvertedBalancesCollectionResponse struct {
	Data     ConvertedBalancesCollection `json:"data"`
	Included Included                    `json:"included"`
}

type ConvertedBalancesCollectionListResponse struct {
	Data     []ConvertedBalancesCollection `json:"data"`
	Included Included                      `json:"included"`
	Links    *Links                        `json:"links"`
}

// MustConvertedBalancesCollection - returns ConvertedBalancesCollection from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustConvertedBalancesCollection(key Key) *ConvertedBalancesCollection {
	var convertedBalancesCollection ConvertedBalancesCollection
	if c.tryFindEntry(key, &convertedBalancesCollection) {
		return &convertedBalancesCollection
	}
	return nil
}
