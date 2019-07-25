package derive

import (
	"encoding/hex"
	"math"

	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type ETHDeriver struct {
	key *hdkeychain.ExtendedKey
}

func NewETHDeriver(src string) (*ETHDeriver, error) {
	key, err := hdkeychain.NewKeyFromString(src)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse key")
	}
	return &ETHDeriver{key}, nil
}

func (s *ETHDeriver) ChildAddress(i uint64) (string, error) {
	if i >= math.MaxUint32 {
		panic("child overflow")
	}
	child, err := s.key.Child(uint32(i))
	if err != nil {
		return "", err
	}

	public, err := child.ECPubKey()
	if err != nil {
		return "", errors.Wrap(err, "failed to get public key")
	}
	address := crypto.PubkeyToAddress(*public.ToECDSA())

	return address.Hex(), nil
}

func (s *ETHDeriver) ChildPrivate(i uint32) (string, error) {
	child, err := s.key.Child(i)
	if err != nil {
		return "", err
	}

	private, err := child.ECPrivKey()
	if err != nil {
		return "", errors.Wrap(err, "failed to get public key")
	}

	return hex.EncodeToString(private.Serialize()), nil
}
