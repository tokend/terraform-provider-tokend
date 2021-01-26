package internal

import (
	"testing"
)

func TestNewAddress(t *testing.T) {
	cases := []struct {
		name  string
		input []byte
		panic bool
	}{
		{"valid address", rawAddress, false},
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
			_ = NewAddress(tc.input)
		})
	}
}

func TestAddress_Address(t *testing.T) {
	got := NewAddress(rawAddress).Address()
	if got != address {
		t.Fatalf("got %s expected %s", got, address)
	}
}

func TestAddress_Hint(t *testing.T) {
	got := NewAddress(rawAddress).Hint()
	if got != hint {
		t.Fatalf("got %s expected %s", got, hint)
	}
}

func TestAddress_Verify(t *testing.T) {
	if err := NewAddress(rawAddress).Verify(message, signature); err != nil {
		t.Fatal("expected nil error")
	}
}
