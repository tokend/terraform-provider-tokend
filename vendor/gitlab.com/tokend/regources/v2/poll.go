package regources

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gitlab.com/tokend/regources"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/tokend/go/xdr"
)

//PollResponse - response for poll handler
type PollResponse struct {
	Data     Poll     `json:"data"`
	Included Included `json:"included"`
}

type PollsResponse struct {
	Links    *Links   `json:"links"`
	Data     []Poll   `json:"data"`
	Included Included `json:"included"`
}

// Poll - Resource object representing PollEntry
type Poll struct {
	Key
	Attributes    PollAttributes `json:"attributes"`
	Relationships PollRelations  `json:"relationships"`
}

type PollAttributes struct {
	PollData                 PollData  `json:"poll_data"`
	PermissionType           uint32    `json:"permission_type"`
	NumberOfChoices          uint32    `json:"number_of_choices"`
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
	VoteConfirmationRequired bool      `json:"vote_confirmation_required"`
	Details                  Details   `json:"details"`
	PollState                PollState `json:"poll_state"`
}

type PollRelations struct {
	Owner          *Relation `json:"owner"`
	ResultProvider *Relation `json:"result_provider"`
	Participation  *Relation `json:"participation"`
}
type PollData struct {
	Type xdr.PollType `json:"type"`
}

//Value - implements db driver method for auto marshal
func (r PollData) Value() (driver.Value, error) {
	result, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal poll data")
	}

	return result, nil
}

//Scan - implements db driver method for auto unmarshal
func (r *PollData) Scan(src interface{}) error {
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
		return errors.Wrap(err, "failed to unmarshal poll data")
	}

	return nil
}

// PollParticipation - Resource object representing outcome of the poll
type PollParticipation struct {
	Key
	Relationships PollParticipationRelations `json:"relationships"`
}

type PollParticipationRelations struct {
	Votes *RelationCollection `json:"votes"`
}

type PollState int

const (
	PollStateOpen PollState = iota + 1
	PollStatePassed
	PollStateFailed
)

var pollStateStr = map[PollState]string{
	PollStateOpen:   "open",
	PollStatePassed: "passed",
	PollStateFailed: "failed",
}

func (s PollState) String() string {
	return pollStateStr[s]
}

func (s PollState) MarshalJSON() ([]byte, error) {
	return json.Marshal(regources.Flag{
		Name:  pollStateStr[s],
		Value: int32(s),
	})
}
