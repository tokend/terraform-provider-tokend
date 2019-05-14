// Package amount provides utilities for converting numbers to/from
// the format used internally to stellar-core.
//
// stellar-core represents asset "amounts" as 64-bit integers, but to enable
// fractional units of an asset, horizon, the client-libraries and other built
// on top of stellar-core use a convention, encoding amounts as a string of
// decimal digits with up to seven digits of precision in the fractional
// portion. For example, an amount shown as "101.001" in horizon would be
// represented in stellar-core as 1010010.
package amount

import (
	"fmt"
	"math/big"
	"strconv"
)

// One is the value of one whole unit of currency.  The system uses 6 fixed digits
// for fractional values, thus One is 1 million (10^6)
const One = 1000000

// MustParse is the panicking version of Parse
func MustParse(v string) int64 {
	ret, err := Parse(v)
	if err != nil {
		panic(err)
	}
	return ret
}

// MustParseU is the panicking version of ParseU
func MustParseU(v string) uint64 {
	ret, err := ParseU(v)
	if err != nil {
		panic(err)
	}
	return ret
}

// Parse parses the provided as a stellar "amount", i.e. A 64-bit signed integer
// that represents a decimal number with 7 digits of significance in the
// fractional portion of the number.
func Parse(v string) (int64, error) {
	is, err := stringToIntString(v)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(is, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// String returns an "amount string" from the provided raw value `v`.
func String(v int64) string {
	var f, o, r big.Rat

	f.SetInt64(v)
	o.SetInt64(One)
	r.Quo(&f, &o)

	return r.FloatString(6)
}

func stringToIntString(v string) (string, error) {
	var f, o, r big.Rat

	_, ok := f.SetString(v)
	if !ok {
		return "", fmt.Errorf("cannot parse amount: %s", v)
	}

	o.SetInt64(One)
	r.Mul(&f, &o)

	is := r.FloatString(0)
	return is, nil
}

// Parse parses the provided as a stellar "amount", i.e. A 64-bit unsigned integer
// that represents a decimal number with 7 digits of significance in the
// fractional portion of the number.
func ParseU(v string) (uint64, error) {
	is, err := stringToIntString(v)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseUint(is, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func StringU(v uint64) string {
	var f, o, r big.Rat

	var fInt big.Int
	fInt.SetUint64(v)
	f.SetInt(&fInt)
	o.SetInt64(One)
	r.Quo(&f, &o)

	return r.FloatString(6)
}
