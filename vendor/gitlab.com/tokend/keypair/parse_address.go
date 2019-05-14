package keypair

import (
	"gitlab.com/tokend/keypair/internal"
	"gitlab.com/tokend/keypair/internal/strkey"
)

func ParseAddress(address string) (Address, error) {
	public, err := strkey.Decode(strkey.VersionByteAccountID, address)
	if err != nil {
		return nil, err
	}

	return internal.NewAddress(public), nil
}

func MustParseAddress(address string) Address {
	kp, err := ParseAddress(address)
	if err != nil {
		panic(err)
	}
	return kp
}
