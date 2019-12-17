/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type UpdateReviewableRequestOp struct {
	Key
	Relationships UpdateReviewableRequestOpRelationships `json:"relationships"`
}
type UpdateReviewableRequestOpResponse struct {
	Data     UpdateReviewableRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type UpdateReviewableRequestOpListResponse struct {
	Data     []UpdateReviewableRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *UpdateReviewableRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *UpdateReviewableRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustUpdateReviewableRequestOp - returns UpdateReviewableRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateReviewableRequestOp(key Key) *UpdateReviewableRequestOp {
	var updateReviewableRequestOp UpdateReviewableRequestOp
	if c.tryFindEntry(key, &updateReviewableRequestOp) {
		return &updateReviewableRequestOp
	}
	return nil
}
