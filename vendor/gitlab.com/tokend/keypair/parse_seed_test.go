package keypair

import "testing"

var (
	seedCases = []struct {
		name  string
		input string
		err   bool
	}{
		{"valid address", "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H", true},
		{"corrupted address", "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7O32H", true},
		{"valid seed", "SDHOAMBNLGCE2MV5ZKIVZAQD3VCLGP53P3OBSBI6UN5L5XZI5TKHFQL4", false},
		{"corrupted seed", "SDHOAMBNLGCE2MV5ZKIVZAQD3VCLGP53P3OBSBI6UN5L5XZI5TKHFQL3", true},
		{"blank", "", true},
	}
)

func TestParseSeed(t *testing.T) {
	for _, tc := range seedCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseSeed(tc.input)
			if err != nil && !tc.err {
				t.Fatalf("got %s error expected nil", err)
			}
			if err == nil && tc.err {
				t.Fatal("expected error")
			}
			if err == nil && got.Seed() != tc.input {
				t.Fatalf("expected %s got %s", tc.input, got.Address())
			}
		})
	}
}

func TestMustParseSeed(t *testing.T) {
	for _, tc := range seedCases {
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
			got := MustParseSeed(tc.input)
			if got.Seed() != tc.input {
				t.Fatalf("expected %s got %s", tc.input, got.Address())
			}
		})
	}
}
