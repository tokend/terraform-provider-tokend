/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type FeeRecord struct {
	Key
	Attributes    FeeRecordAttributes    `json:"attributes"`
	Relationships FeeRecordRelationships `json:"relationships"`
}
type FeeRecordResponse struct {
	Data     FeeRecord `json:"data"`
	Included Included  `json:"included"`
}

type FeeRecordListResponse struct {
	Data     []FeeRecord     `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *FeeRecordListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *FeeRecordListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustFeeRecord - returns FeeRecord from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustFeeRecord(key Key) *FeeRecord {
	var feeRecord FeeRecord
	if c.tryFindEntry(key, &feeRecord) {
		return &feeRecord
	}
	return nil
}
