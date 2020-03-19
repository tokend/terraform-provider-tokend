package xdr

import (
	"errors"
	"fmt"

	"gitlab.com/tokend/go/strkey"
)

func (bid *BalanceId) AsString() string {
	if bid == nil {
		return ""
	}

	switch bid.Type {
	case CryptoKeyTypeKeyTypeEd25519:
		ed := bid.MustEd25519()
		raw := make([]byte, 32)
		copy(raw, ed[:])
		return strkey.MustEncode(strkey.VersionByteBalanceID, raw)
	default:
		panic(fmt.Errorf("Unknown account id type: %v", bid.Type))
	}
}

func (bid *BalanceId) SetString(addr string) error {
	if bid == nil {
		return nil
	}

	raw, err := strkey.Decode(strkey.VersionByteBalanceID, addr)
	if err != nil {
		return err
	}

	if len(raw) != 32 {
		return errors.New("invalid address")
	}

	var ui Uint256
	copy(ui[:], raw)

	*bid, err = NewBalanceId(CryptoKeyTypeKeyTypeEd25519, ui)

	return err
}

func (bid *BalanceId) MarshalJSON() ([]byte, error) {
	return []byte("\"" + bid.AsString() + "\""), nil
}

func (bid *BalanceId) UnmarshalJSON(raw []byte) error {
	return bid.SetString(string(raw[1 : len(raw)-1]))
}

// Equals returns true if `other` is equivalent to `aid`
func (bid *BalanceId) Equals(other BalanceId) bool {
	if bid.Type != other.Type {
		return false
	}

	switch bid.Type {
	case CryptoKeyTypeKeyTypeEd25519:
		l := bid.MustEd25519()
		r := other.MustEd25519()
		return l == r
	default:
		panic(fmt.Errorf("Unknown account id type: %v", bid.Type))
	}
}
