package xdr

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/strkey"
)

func (u *PublicKey) FromString(str string) error {
	if u == nil {
		return nil
	}

	raw, err := strkey.Decode(strkey.VersionByteAccountID, str)
	if err != nil {
		return errors.Wrap(err, "failed to decode string address")
	}

	if len(raw) != 32 {
		return errors.New("invalid string address")
	}

	var ui Uint256
	copy(ui[:], raw)

	u.Type = CryptoKeyTypeKeyTypeEd25519
	u.Ed25519 = &ui

	return nil
}

func (u *PublicKey) ToString() string {
	if u == nil {
		return ""
	}

	switch u.Type {
	case CryptoKeyTypeKeyTypeEd25519:
		ed := u.MustEd25519()
		return strkey.MustEncode(strkey.VersionByteAccountID, ed[:])
	default:
		panic(fmt.Errorf("unknown public key type: %v", u.Type))
	}
}

// Equals returns true if `other` is equivalent to `aid`
func (u *PublicKey) Equals(other PublicKey) bool {
	if u.Type != other.Type {
		return false
	}

	switch u.Type {
	case CryptoKeyTypeKeyTypeEd25519:
		l := u.MustEd25519()
		r := other.MustEd25519()
		return l == r
	default:
		panic(fmt.Errorf("unknown account id type: %v", u.Type))
	}
}
