package amount

import (
	"math"
	"testing"
)

var Tests = []struct {
	A             int64
	B             int64
	C             int64
	Rounding      Rounding
	Result        int64
	Overflow      bool
	MinimalAmount int64
}{
	{3, 200, 6, ROUND_UP, 100, false, 100},
	{3, 200, 6, ROUND_DOWN, 100, false, 100},
	{1, 1, 2, ROUND_UP, 1, false, 1},
	{1, 1, 2, ROUND_DOWN, 0, false, 1},
	{1, 1, 2 * One, ROUND_UP, 1, false, 1},
	{1, 1, 2 * One, ROUND_DOWN, 0, false, 1},

	// A*B overflows
	{math.MaxInt64, 2, math.MaxInt64, ROUND_UP, 2, false, 2},
	{math.MaxInt64, 2, math.MaxInt64, ROUND_DOWN, 2, false, 2},
	{math.MaxInt64, math.MaxInt64, math.MaxInt64, ROUND_DOWN, math.MaxInt64, false, math.MaxInt64},
	{math.MaxInt64, math.MaxInt64, math.MaxInt64, ROUND_UP, math.MaxInt64, false, math.MaxInt64},

	// overflow
	{math.MaxInt64, 4, 3, ROUND_UP, 0, true, 1},
	{math.MaxInt64, 4, 3, ROUND_DOWN, 0, true, 1},
	{math.MaxInt64, 3, 2, ROUND_UP, 0, true, 1},
	{math.MaxInt64, 3, 2, ROUND_DOWN, 0, true, 1},

	{3 * One, One, 2 * One, ROUND_UP, 2 * One, false, One},
	{3 * One, One, 2 * One, ROUND_DOWN, One, false, One},
	{math.MaxInt64, math.MaxInt64, math.MaxInt64, ROUND_UP, math.MaxInt64, true, 10},
	{math.MaxInt64, math.MaxInt64, math.MaxInt64, ROUND_DOWN, math.MaxInt64 - (math.MaxInt64 % 10), false, 10},
}

func TestParse(t *testing.T) {
	for _, v := range Tests {
		result, overflow := BigDivide(v.A, v.B, v.C, v.Rounding, v.MinimalAmount)
		if overflow != v.Overflow {
			t.Fatalf("Unexpected overflow result. Result: %d, overflow: %t. Data: %+v.", result, overflow, v)
		}

		if overflow {
			continue
		}

		if result != v.Result {
			t.Fatalf("Unexpected result. Result: %d, overflow: %t. Data: %+v.", result, overflow, v)
		}
	}
}
