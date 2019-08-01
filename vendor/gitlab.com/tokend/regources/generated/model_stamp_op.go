/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type StampOp struct {
	Key
	Attributes StampOpAttributes `json:"attributes"`
}
type StampOpResponse struct {
	Data     StampOp  `json:"data"`
	Included Included `json:"included"`
}

type StampOpListResponse struct {
	Data     []StampOp `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustStampOp - returns StampOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustStampOp(key Key) *StampOp {
	var stampOp StampOp
	if c.tryFindEntry(key, &stampOp) {
		return &stampOp
	}
	return nil
}
