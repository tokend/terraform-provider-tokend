package regources

import (
	"encoding/json"
	"testing"

	"gitlab.com/tokend/go/amount"

	"github.com/stretchr/testify/assert"
)

type TestAmount struct {
	TestValue Amount `json:"test_amount_value"`
}

func TestUnmarshalJSON(t *testing.T) {
	body := `{
		"test_amount_value": "50"
	}`
	expected := TestAmount{
		TestValue: 50 * amount.One,
	}
	var got TestAmount

	assert.NoError(t, json.Unmarshal([]byte(body), &got))
	assert.Equal(t, expected.TestValue, got.TestValue)
}

func TestMarshalUnmarshal(t *testing.T) {
	var expected Amount = 42
	bytes, err := json.Marshal(expected)
	assert.NoError(t, err)

	var got Amount
	err = json.Unmarshal(bytes, &got)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}
