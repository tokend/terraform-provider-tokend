/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type DataCreationRequest struct {
	Key
	Attributes    DataCreationRequestAttributes    `json:"attributes"`
	Relationships DataCreationRequestRelationships `json:"relationships"`
}
type DataCreationRequestResponse struct {
	Data     DataCreationRequest `json:"data"`
	Included Included            `json:"included"`
}

type DataCreationRequestListResponse struct {
	Data     []DataCreationRequest `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
	Meta     json.RawMessage       `json:"meta,omitempty"`
}

func (r *DataCreationRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *DataCreationRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustDataCreationRequest - returns DataCreationRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDataCreationRequest(key Key) *DataCreationRequest {
	var dataCreationRequest DataCreationRequest
	if c.tryFindEntry(key, &dataCreationRequest) {
		return &dataCreationRequest
	}
	return nil
}
