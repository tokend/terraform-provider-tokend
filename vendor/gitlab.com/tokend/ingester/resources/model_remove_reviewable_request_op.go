/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type RemoveReviewableRequestOp struct {
	Key
	Relationships RemoveReviewableRequestOpRelationships `json:"relationships"`
}
type RemoveReviewableRequestOpResponse struct {
	Data     RemoveReviewableRequestOp `json:"data"`
	Included Included                  `json:"included"`
}

type RemoveReviewableRequestOpListResponse struct {
	Data     []RemoveReviewableRequestOp `json:"data"`
	Included Included                    `json:"included"`
	Links    *Links                      `json:"links"`
	Meta     json.RawMessage             `json:"meta,omitempty"`
}

func (r *RemoveReviewableRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *RemoveReviewableRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustRemoveReviewableRequestOp - returns RemoveReviewableRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRemoveReviewableRequestOp(key Key) *RemoveReviewableRequestOp {
	var removeReviewableRequestOp RemoveReviewableRequestOp
	if c.tryFindEntry(key, &removeReviewableRequestOp) {
		return &removeReviewableRequestOp
	}
	return nil
}
