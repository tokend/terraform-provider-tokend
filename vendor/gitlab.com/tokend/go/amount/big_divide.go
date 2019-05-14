package amount

import (
	"math"
	"math/big"
)

type Rounding int32

const (
	ROUND_UP Rounding = iota
	ROUND_DOWN
)

// calculates A*B/C when A*B overflows 64bits
// panics if one of the arguments is negative or C = 0
// returns false, if overflow
func BigDivide(rawA, rawB, rawC int64, rounding Rounding, minimalAmount int64) (int64, bool) {
	if rawA < 0 || rawB < 0 || rawC <= 0 || minimalAmount <= 0 {
		panic("one of the arguments is negative or C = 0 or minimalAmount = 0")
	}

	a := big.NewInt(rawA)
	b := big.NewInt(rawB)
	c := big.NewInt(rawC)

	ab := big.NewInt(0).Mul(a, b)

	if rounding == ROUND_UP {
		ab = ab.Add(ab, big.NewInt(rawC-1))
	}

	result := big.NewInt(0).Div(ab, c)

	min := big.NewInt(minimalAmount)
	modRes := ab.Mod(result, min) // ab will not be used, but we need any *Int var
	switch rounding {
	case ROUND_UP:
		if modRes.Cmp(big.NewInt(0)) != 0 {
			result.Add(result, min.Sub(min, modRes))
		}
	case ROUND_DOWN:
		result.Sub(result, modRes)
	}

	isOverflow := result.Cmp(big.NewInt(math.MaxInt64)) == 1
	return result.Int64(), isOverflow
}
