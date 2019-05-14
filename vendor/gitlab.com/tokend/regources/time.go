package regources

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.RFC3339))
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var rawTime string
	err := json.Unmarshal(data, &rawTime)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal time")
	}

	rawT, err := time.Parse(time.RFC3339, rawTime)
	*t = Time(rawT)

	return errors.Wrap(err, "can't parse time")
}

func (t Time) Time() time.Time {
	return time.Time(t)
}
