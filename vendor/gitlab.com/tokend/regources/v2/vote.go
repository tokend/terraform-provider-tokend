package regources

import (
	"database/sql/driver"
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

//VoteResponse - response for vote handler
type VoteResponse struct {
	Data     Vote     `json:"data"`
	Included Included `json:"included"`
}

type VotesResponse struct {
	Links    *Links   `json:"links"`
	Data     []Vote   `json:"data"`
	Included Included `json:"included"`
}

// Vote - Resource object representing VoteEntry
type Vote struct {
	Key
	Attributes    VoteAttributes `json:"attributes"`
	Relationships VoteRelations  `json:"relationships"`
}

type VoteAttributes struct {
	VoteData VoteData `json:"vote_data"`
}

type VoteRelations struct {
	Voter *Relation `json:"voter"`
	Poll  *Relation `json:"poll"`
}

type VoteData struct {
	PollType     xdr.PollType `json:"poll_type"`
	SingleChoice *uint64      `json:"single_choice,omitempty"`
}

//Value - implements db driver method for auto marshal
func (r VoteData) Value() (driver.Value, error) {
	result, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal vote data")
	}

	return result, nil
}

//Scan - implements db driver method for auto unmarshal
func (r *VoteData) Scan(src interface{}) error {
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
		return errors.Wrap(err, "failed to unmarshal vote data")
	}

	return nil
}
