package logan

import (
	"testing"
)

func TestFields(t *testing.T) {
	cases := []struct {
		actual   F
		expected F
	}{
		// 1
		{F{"key": "value"}, map[string]interface{}{
			"key": "value",
		}},
		{F{"key": "value"}, map[string]interface{}{
			"key": "value",
		}},
		// 2
		{F{"key": "value"}.Add("key_2", "value_2"), map[string]interface{}{
			"key": "value",
			"key_2": "value_2",
		}},
		{F{"key": "value", "key_2": "value_2"}, map[string]interface{}{
			"key": "value",
			"key_2": "value_2",
		}},
		// 3
		{F{"key": "value"}.AddFields(F{"key_2": "value_2", "key_3": "value_3"}), map[string]interface{}{
			"key": "value",
			"key_2": "value_2",
			"key_3": "value_3",
		}},
		{F{"key": "value"}.AddFields(F{"key_2": "value_2"}.Add("key_3", "value_3")), map[string]interface{}{
			"key": "value",
			"key_2": "value_2",
			"key_3": "value_3",
		}},
		// Overwrite
		{F{"key": "value"}.AddFields(F{"key_2": "value_2", "key": "value_new"}), map[string]interface{}{
			"key": "value_new",
			"key_2": "value_2",
		}},
		{F{"key": "value"}.AddFields(F{"key_2": "value_2"}.Add("key", "value_new")), map[string]interface{}{
			"key": "value_new",
			"key_2": "value_2",
		}},
	}

	for _, tc := range cases {
		if len(tc.actual) != len(tc.expected) {
			t.Errorf("Wrong fields size: expected %d got %d", tc.expected, len(tc.actual))
		}

		for expectedK, expectedV := range tc.expected {
			value, ok := tc.actual[expectedK]
			if !ok {
				t.Errorf("Missing key: %s", expectedK)
				continue
			}

			if value != expectedV {
				t.Errorf("Wrong value for %s: expected %s got %s", expectedK, expectedV, value)
			}
		}
	}
}
