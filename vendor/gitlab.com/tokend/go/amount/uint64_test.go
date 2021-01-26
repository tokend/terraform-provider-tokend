package amount_test

import (
	"gitlab.com/tokend/go/amount"
	"math"
	"testing"
)

var testsDataU = []struct {
	S string
	I uint64
}{
	{"100.000000", 100000000},
	{"100.000100", 100000100},
	{"123.000100", 123000100},
	{"18446744073709.551615", math.MaxUint64},
}

func TestUint64Parse(t *testing.T) {
	for _, v := range testsDataU {
		o, err := amount.ParseU(v.S)
		if err != nil {
			t.Errorf("Couldn't parse %s: %v+", v.S, err)
			continue
		}

		if o != v.I {
			t.Errorf("%s parsed to %d, not %d", v.S, o, v.I)
		}
	}
}

func TestUString(t *testing.T) {
	for _, v := range testsDataU {
		o := amount.StringU(v.I)

		if o != v.S {
			t.Errorf("%d stringified to %s, not %s", v.I, o, v.S)
		}
	}
}
