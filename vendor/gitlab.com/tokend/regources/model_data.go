/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Data struct {
	Key
	Attributes    DataAttributes    `json:"attributes"`
	Relationships DataRelationships `json:"relationships"`
}
type DataResponse struct {
	Data     Data     `json:"data"`
	Included Included `json:"included"`
}

type DataListResponse struct {
	Data     []Data          `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *DataListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DataListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustData - returns Data from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustData(key Key) *Data {
	var data Data
	if c.tryFindEntry(key, &data) {
		return &data
	}
	return nil
}
