package derive

import (
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/pkg/errors"
)

type HDMaster struct {
	master *hdkeychain.ExtendedKey
}

func NewHDMaster(network NetworkType) (*HDMaster, error) {
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate seed")
	}
	master, err := hdkeychain.NewMaster(seed, NetworkParams(network))
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate master")
	}
	return &HDMaster{master}, nil
}

func (s *HDMaster) ExtendedPrivate() (string, error) {
	return s.master.String(), nil
}

func (s *HDMaster) ExtendedPublic() (string, error) {
	public, err := s.master.Neuter()
	if err != nil {
		return "", err
	}
	return public.String(), nil
}
