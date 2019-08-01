/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type Statistics struct {
	Key
	Attributes    StatisticsAttributes    `json:"attributes"`
	Relationships StatisticsRelationships `json:"relationships"`
}
type StatisticsResponse struct {
	Data     Statistics `json:"data"`
	Included Included   `json:"included"`
}

type StatisticsListResponse struct {
	Data     []Statistics `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustStatistics - returns Statistics from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustStatistics(key Key) *Statistics {
	var statistics Statistics
	if c.tryFindEntry(key, &statistics) {
		return &statistics
	}
	return nil
}
