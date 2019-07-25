package derive

import (
	"math"

	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/pkg/errors"
)

type BTCFamilyDeriver struct {
	network NetworkType
	key     *hdkeychain.ExtendedKey
}

func NewBTCFamilyDeriver(network NetworkType, src string) (*BTCFamilyDeriver, error) {
	key, err := hdkeychain.NewKeyFromString(src)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse key")
	}
	return &BTCFamilyDeriver{network, key}, nil
}

func (s *BTCFamilyDeriver) ChildAddress(i uint64) (string, error) {
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

	pubKeyHash := btcutil.Hash160(public.SerializeCompressed())
	addr, err := btcutil.NewAddressPubKeyHash(pubKeyHash, NetworkParams(s.network))
	if err != nil {
		return "", errors.Wrap(err, "failed to hash public key")
	}

	return addr.EncodeAddress(), nil
}

func (s *BTCFamilyDeriver) ChildPrivate(i uint64) (string, error) {
	if i >= math.MaxUint32 {
		panic("child overflow")
	}
	child, err := s.key.Child(uint32(i))
	if err != nil {
		return "", errors.Wrap(err, "failed to derive child")
	}

	private, err := child.ECPrivKey()
	if err != nil {
		return "", errors.Wrap(err, "failed to get public key")
	}

	wif, err := btcutil.NewWIF(private, NetworkParams(s.network), true)
	if err != nil {
		return "", errors.Wrap(err, "failed to get WIF")
	}

	return wif.String(), nil
}
