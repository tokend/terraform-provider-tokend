/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"database/sql/driver"
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Poll struct {
	Key
	Attributes    PollAttributes    `json:"attributes"`
	Relationships PollRelationships `json:"relationships"`
}
type PollResponse struct {
	Data     Poll     `json:"data"`
	Included Included `json:"included"`
}

type PollListResponse struct {
	Data     []Poll   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustPoll - returns Poll from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPoll(key Key) *Poll {
	var poll Poll
	if c.tryFindEntry(key, &poll) {
		return &poll
	}
	return nil
}

//Value - implements db driver method for auto marshal
func (r Poll) Value() (driver.Value, error) {
	result, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal Poll data")
	}

	return result, nil
}

//Scan - implements db driver method for auto unmarshal
func (r *Poll) Scan(src interface{}) error {
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
		return errors.Wrap(err, "failed to unmarshal Poll data")
	}

	return nil
}
