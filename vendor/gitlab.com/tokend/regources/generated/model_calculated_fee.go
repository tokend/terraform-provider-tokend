/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type CalculatedFee struct {
	Key
	Attributes Fee `json:"attributes"`
}
type CalculatedFeeResponse struct {
	Data     CalculatedFee `json:"data"`
	Included Included      `json:"included"`
}

type CalculatedFeeListResponse struct {
	Data     []CalculatedFee `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustCalculatedFee - returns CalculatedFee from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCalculatedFee(key Key) *CalculatedFee {
	var calculatedFee CalculatedFee
	if c.tryFindEntry(key, &calculatedFee) {
		return &calculatedFee
	}
	return nil
}
