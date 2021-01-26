/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type CreateDataCreationRequestOp struct {
	Key
	Attributes    CreateDataCreationRequestOpAttributes    `json:"attributes"`
	Relationships CreateDataCreationRequestOpRelationships `json:"relationships"`
}
type CreateDataCreationRequestOpResponse struct {
	Data     CreateDataCreationRequestOp `json:"data"`
	Included Included                    `json:"included"`
}

type CreateDataCreationRequestOpListResponse struct {
	Data     []CreateDataCreationRequestOp `json:"data"`
	Included Included                      `json:"included"`
	Links    *Links                        `json:"links"`
	Meta     json.RawMessage               `json:"meta,omitempty"`
}

func (r *CreateDataCreationRequestOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *CreateDataCreationRequestOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustCreateDataCreationRequestOp - returns CreateDataCreationRequestOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateDataCreationRequestOp(key Key) *CreateDataCreationRequestOp {
	var createDataCreationRequestOp CreateDataCreationRequestOp
	if c.tryFindEntry(key, &createDataCreationRequestOp) {
		return &createDataCreationRequestOp
	}
	return nil
}
