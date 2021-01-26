package keypair

import (
	"testing"
)

var (
	addressCases = []struct {
		name  string
		input string
		err   bool
	}{
		{"valid address", "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H", false},
		{"corrupted address", "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7O32H", true},
		{"valid seed", "SDHOAMBNLGCE2MV5ZKIVZAQD3VCLGP53P3OBSBI6UN5L5XZI5TKHFQL4", true},
		{"corrupted seed", "SDHOAMBNLGCE2MV5ZKIVZAQD3VCLGP53P3OBSBI6UN5L5XZI5TKHFQL3", true},
		{"blank", "", true},
	}
)

func TestParseAddress(t *testing.T) {
	for _, tc := range addressCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseAddress(tc.input)
			if err != nil && !tc.err {
				t.Fatalf("got %s error expected nil", err)
			}
			if err == nil && tc.err {
				t.Fatal("expected error")
			}
			if err == nil && got.Address() != tc.input {
				t.Fatalf("expected %s got %s", tc.input, got.Address())
			}
		})
	}
}

func TestMustParseAddress(t *testing.T) {
	for _, tc := range addressCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				rvr := recover()
				if rvr != nil && !tc.err {
					t.Fatalf("got %s panic expected nil", rvr)
				}
				if rvr == nil && tc.err {
					t.Fatal("expected panic")
				}
			}()
			got := MustParseAddress(tc.input)
			if got.Address() != tc.input {
				t.Fatalf("expected %s got %s", tc.input, got.Address())
			}
		})
	}
}
