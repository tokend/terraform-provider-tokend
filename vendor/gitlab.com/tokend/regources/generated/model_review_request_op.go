/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ReviewRequestOp struct {
	Key
	Attributes ReviewRequestOpAttributes `json:"attributes"`
}
type ReviewRequestOpResponse struct {
	Data     ReviewRequestOp `json:"data"`
	Included Included        `json:"included"`
}

type ReviewRequestOpListResponse struct {
	Data     []ReviewRequestOp `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustReviewRequestOp - returns ReviewRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustReviewRequestOp(key Key) *ReviewRequestOp {
	var reviewRequestOp ReviewRequestOp
	if c.tryFindEntry(key, &reviewRequestOp) {
		return &reviewRequestOp
	}
	return nil
}
