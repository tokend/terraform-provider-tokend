/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type LimitsWithStats struct {
	Key
	Attributes    LimitsWithStatsAttributes    `json:"attributes"`
	Relationships LimitsWithStatsRelationships `json:"relationships"`
}
type LimitsWithStatsResponse struct {
	Data     LimitsWithStats `json:"data"`
	Included Included        `json:"included"`
}

type LimitsWithStatsListResponse struct {
	Data     []LimitsWithStats `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustLimitsWithStats - returns LimitsWithStats from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustLimitsWithStats(key Key) *LimitsWithStats {
	var limitsWithStats LimitsWithStats
	if c.tryFindEntry(key, &limitsWithStats) {
		return &limitsWithStats
	}
	return nil
}
