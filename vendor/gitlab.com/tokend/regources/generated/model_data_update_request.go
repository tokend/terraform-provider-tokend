/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type DataUpdateRequest struct {
	Key
	Attributes    DataUpdateRequestAttributes    `json:"attributes"`
	Relationships DataUpdateRequestRelationships `json:"relationships"`
}
type DataUpdateRequestResponse struct {
	Data     DataUpdateRequest `json:"data"`
	Included Included          `json:"included"`
}

type DataUpdateRequestListResponse struct {
	Data     []DataUpdateRequest `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *DataUpdateRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DataUpdateRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDataUpdateRequest - returns DataUpdateRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDataUpdateRequest(key Key) *DataUpdateRequest {
	var dataUpdateRequest DataUpdateRequest
	if c.tryFindEntry(key, &dataUpdateRequest) {
		return &dataUpdateRequest
	}
	return nil
}
