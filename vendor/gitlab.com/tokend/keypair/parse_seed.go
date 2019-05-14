package keypair

import (
	"gitlab.com/tokend/keypair/internal"
	"gitlab.com/tokend/keypair/internal/strkey"
)

func ParseSeed(seed string) (Full, error) {
	raw, err := strkey.Decode(strkey.VersionByteSeed, seed)
	if err != nil {
		return nil, err
	}

	return internal.NewFull(raw), nil
}

func MustParseSeed(seed string) Full {
	kp, err := ParseSeed(seed)
	if err != nil {
		panic(err)
	}
	return kp
}
