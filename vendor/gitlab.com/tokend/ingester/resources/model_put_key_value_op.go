/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "encoding/json"

type PutKeyValueOp struct {
	Key
	Attributes PutKeyValueOpAttributes `json:"attributes"`
}
type PutKeyValueOpResponse struct {
	Data     PutKeyValueOp `json:"data"`
	Included Included      `json:"included"`
}

type PutKeyValueOpListResponse struct {
	Data     []PutKeyValueOp `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *PutKeyValueOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *PutKeyValueOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustPutKeyValueOp - returns PutKeyValueOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPutKeyValueOp(key Key) *PutKeyValueOp {
	var putKeyValueOp PutKeyValueOp
	if c.tryFindEntry(key, &putKeyValueOp) {
		return &putKeyValueOp
	}
	return nil
}
