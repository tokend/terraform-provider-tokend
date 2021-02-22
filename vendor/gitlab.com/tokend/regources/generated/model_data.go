/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"database/sql/driver"
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

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

//Value - implements db driver method for auto marshal
func (r Data) Value() (driver.Value, error) {
	result, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal Data data")
	}

	return result, nil
}

//Scan - implements db driver method for auto unmarshal
func (r *Data) Scan(src interface{}) error {
	var data []byte
	switch rawData := src.(type) {
	case []byte:
		data = rawData
	case string:
		data = []byte(rawData)
	default:
		return errors.New("Unexpected type for jsonb")
	}

	err := json.Unmarshal(data, r)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal Data data")
	}

	return nil
}
