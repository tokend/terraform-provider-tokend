/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type PollState int

const (
	PollStateOpen PollState = iota + 1
	PollStatePassed
	PollStateFailed
	PollStateCancelled
)

var pollStateStr = map[PollState]string{
	PollStateOpen:      "open",
	PollStatePassed:    "passed",
	PollStateFailed:    "failed",
	PollStateCancelled: "cancelled",
}

func (s PollState) String() string {
	return pollStateStr[s]
}

func (s PollState) MarshalJSON() ([]byte, error) {
	return json.Marshal(Flag{
		Name:  pollStateStr[s],
		Value: int32(s),
	})
}

func (s *PollState) UnmarshalJSON(b []byte) error {
	var res Flag
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}

	*s = PollState(res.Value)
	return nil
}

func (s PollState) IsFlag() bool {
	return true
}
