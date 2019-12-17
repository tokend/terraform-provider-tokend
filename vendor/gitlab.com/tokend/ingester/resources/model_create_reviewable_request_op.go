/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type CreateReviewableRequestOp struct {
	Key
	Attributes    CreateReviewableRequestOpAttributes    `json:"attributes"`
	Relationships CreateReviewableRequestOpRelationships `json:"relationships"`
}
type CreateReviewableRequestOpResponse struct {
	Data     CreateReviewableRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type CreateReviewableRequestOpListResponse struct {
	Data     []CreateReviewableRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *CreateReviewableRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateReviewableRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateReviewableRequestOp - returns CreateReviewableRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateReviewableRequestOp(key Key) *CreateReviewableRequestOp {
	var createReviewableRequestOp CreateReviewableRequestOp
	if c.tryFindEntry(key, &createReviewableRequestOp) {
		return &createReviewableRequestOp
	}
	return nil
}
