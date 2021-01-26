package amount_test

import (
	"testing"

	"gitlab.com/tokend/go/amount"
)

var testsData = []struct {
	S string
	I int64
}{
	{"100.000000", 100000000},
	{"100.000100", 100000100},
	{"123.000100", 123000100},
}

func TestParse(t *testing.T) {
	for _, v := range testsData {
		o, err := amount.Parse(v.S)
		if err != nil {
			t.Errorf("Couldn't parse %s: %v+", v.S, err)
			continue
		}

		if o != v.I {
			t.Errorf("%s parsed to %d, not %d", v.S, o, v.I)
		}
	}
}

func TestString(t *testing.T) {
	for _, v := range testsData {
		o := amount.String(v.I)

		if o != v.S {
			t.Errorf("%d stringified to %s, not %s", v.I, o, v.S)
		}
	}
}
