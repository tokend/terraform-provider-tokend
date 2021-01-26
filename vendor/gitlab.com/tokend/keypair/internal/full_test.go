package internal

import (
	"bytes"
	"testing"

	"gitlab.com/tokend/keypair/internal/strkey"
)

var (
	address    = "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"
	rawAddress = strkey.MustDecode(strkey.VersionByteAccountID, address)
	seed       = "SDHOAMBNLGCE2MV5ZKIVZAQD3VCLGP53P3OBSBI6UN5L5XZI5TKHFQL4"
	rawSeed    = strkey.MustDecode(strkey.VersionByteSeed, seed)
	hint       = [4]byte{0x56, 0xfc, 0x05, 0xf7}
	message    = []byte("hello")
	signature  = []byte{
		0x2E, 0x75, 0xcc, 0x20, 0xd5, 0x19, 0x11, 0x1c, 0xaa, 0xaa, 0xdd, 0xdf,
		0x46, 0x4b, 0xb6, 0x50, 0xd2, 0xea, 0xf0, 0xa5, 0xd1, 0x8d, 0x74, 0x56,
		0x93, 0xa1, 0x61, 0x00, 0xf2, 0xa4, 0x93, 0x7b, 0xc1, 0xdf, 0xfa, 0x8b,
		0x0b, 0x1f, 0x61, 0xa2, 0x76, 0x99, 0x6d, 0x7e, 0xe8, 0xde, 0xb2, 0xd0,
		0xdd, 0x9e, 0xe5, 0x10, 0x55, 0x60, 0x77, 0xb0, 0x2d, 0xec, 0x16, 0x79,
		0x2e, 0x91, 0x5c, 0x0a,
	}
)

func TestNewFull(t *testing.T) {
	cases := []struct {
		name  string
		input []byte
		panic bool
	}{
		{"valid seed", rawSeed, false},
		{"invalid length", []byte{0, 1, 2, 3}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				rvr := recover()
				if rvr != nil && !tc.panic {
					t.Fatalf("got unexpected panic: %s", rvr)
				}
				if rvr == nil && tc.panic {
					t.Fatal("expected panic")
				}
			}()
			_ = NewFull(tc.input)
		})
	}
}
func TestFull_Address(t *testing.T) {
	got := NewFull(rawSeed).Address()
	if got != address {
		t.Fatalf("got %s expected %s", got, address)
	}
}

func TestFull_Hint(t *testing.T) {
	got := NewFull(rawSeed).Hint()
	if got != hint {
		t.Fatalf("got %s expected %s", got, hint)
	}
}

func TestFull_Seed(t *testing.T) {
	got := NewFull(rawSeed).Seed()
	if got != seed {
		t.Fatalf("got %s expected %s", got, seed)
	}
}

func TestFull_Sign(t *testing.T) {
	got, err := NewFull(rawSeed).Sign(message)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, signature) {
		t.Fatalf("got %s expected %s", got, signature)
	}
}

func TestFull_Verify(t *testing.T) {
	if err := NewFull(rawSeed).Verify(message, signature); err != nil {
		t.Fatal("expected nil error")
	}
}
