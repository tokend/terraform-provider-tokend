/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreatePollRequest struct {
	Key
	Attributes    CreatePollRequestAttributes    `json:"attributes"`
	Relationships CreatePollRequestRelationships `json:"relationships"`
}
type CreatePollRequestResponse struct {
	Data     CreatePollRequest `json:"data"`
	Included Included          `json:"included"`
}

type CreatePollRequestListResponse struct {
	Data     []CreatePollRequest `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *CreatePollRequestListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreatePollRequestListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreatePollRequest - returns CreatePollRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreatePollRequest(key Key) *CreatePollRequest {
	var createPollRequest CreatePollRequest
	if c.tryFindEntry(key, &createPollRequest) {
		return &createPollRequest
	}
	return nil
}
