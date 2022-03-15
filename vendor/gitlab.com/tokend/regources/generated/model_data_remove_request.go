/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type DataRemoveRequest struct {
	Key
	Attributes    DataRemoveRequestAttributes    `json:"attributes"`
	Relationships DataRemoveRequestRelationships `json:"relationships"`
}
type DataRemoveRequestResponse struct {
	Data     DataRemoveRequest `json:"data"`
	Included Included          `json:"included"`
}

type DataRemoveRequestListResponse struct {
	Data     []DataRemoveRequest `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *DataRemoveRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DataRemoveRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDataRemoveRequest - returns DataRemoveRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDataRemoveRequest(key Key) *DataRemoveRequest {
	var dataRemoveRequest DataRemoveRequest
	if c.tryFindEntry(key, &dataRemoveRequest) {
		return &dataRemoveRequest
	}
	return nil
}
