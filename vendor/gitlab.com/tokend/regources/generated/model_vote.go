/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"database/sql/driver"
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Vote struct {
	Key
	Attributes    VoteAttributes    `json:"attributes"`
	Relationships VoteRelationships `json:"relationships"`
}
type VoteResponse struct {
	Data     Vote     `json:"data"`
	Included Included `json:"included"`
}

type VoteListResponse struct {
	Data     []Vote   `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustVote - returns Vote from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVote(key Key) *Vote {
	var vote Vote
	if c.tryFindEntry(key, &vote) {
		return &vote
	}
	return nil
}

//Value - implements db driver method for auto marshal
func (r Vote) Value() (driver.Value, error) {
	result, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal Vote data")
	}

	return result, nil
}

//Scan - implements db driver method for auto unmarshal
func (r *Vote) Scan(src interface{}) error {
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
		return errors.Wrap(err, "failed to unmarshal Vote data")
	}

	return nil
}
