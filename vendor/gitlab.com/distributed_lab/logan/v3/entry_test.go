package logan

import (
	"github.com/sirupsen/logrus"
	"testing"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type customError struct {
	Message string
	Field1  string
	Field2  int
}

func (e customError) Error() string {
	return e.Message
}

func (e customError) GetLoganFields() map[string]interface{} {
	return map[string]interface{} {
		"field_1": e.Field1,
		"field_2": e.Field2,
	}
}

func TestEntry(t *testing.T) {
	err := errors.From(errors.New("Error."), F{
		"key": "value",
	})

	customErr := customError{
		Message: "Awesome error message.",
		Field1:  "value_1",
		Field2:  17,
	}

	cases := []struct {
		actual   *Entry
		expected *Entry
	}{
		// Just 1 field and check for working with embedded logrus.Entry properly
		{New().WithField("key", "value"), &Entry{
			entry: &logrus.Entry{
				Data: map[string]interface{}{
					"key": "value",
				},
			},
		}},
		// Chaining `WithField()` calls
		{New().WithField("key", "value").WithField("key2", "value2"),
			New().WithFields(F{
				"key":  "value",
				"key2": "value2",
			})},
		// Extracting errorFields from general logging fields
		{New().WithField("my_err", err),
			New().WithFields(F{
				"my_err":  err,
				"key": "value",
			})},
		// Extracting GetLoganFields() fields from error
		{New().WithError(customErr),
			New().WithFields(F{
				"field_1": "value_1",
				"field_2": 17,
			})},
	}

	for _, tc := range cases {
		for expectedK, expectedV := range tc.expected.entry.Data {
			value, ok := tc.actual.entry.Data[expectedK]
			if !ok {
				t.Errorf("Missing key: %s", expectedK)
				continue
			}

			if value != expectedV {
				t.Errorf("Wrong value for %s: expected (%s) got (%s)", expectedK, expectedV, value)
			}
		}
	}

}
