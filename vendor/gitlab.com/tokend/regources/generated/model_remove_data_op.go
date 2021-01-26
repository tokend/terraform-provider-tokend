/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type RemoveDataOp struct {
	Key
	Relationships RemoveDataOpRelationships `json:"relationships"`
}
type RemoveDataOpResponse struct {
	Data     RemoveDataOp `json:"data"`
	Included Included     `json:"included"`
}

type RemoveDataOpListResponse struct {
	Data     []RemoveDataOp  `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *RemoveDataOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *RemoveDataOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustRemoveDataOp - returns RemoveDataOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRemoveDataOp(key Key) *RemoveDataOp {
	var removeDataOp RemoveDataOp
	if c.tryFindEntry(key, &removeDataOp) {
		return &removeDataOp
	}
	return nil
}
