/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

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
	Meta     json.RawMessage   `json:"meta,omitempty"`
}

func (r *ReviewRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ReviewRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
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
