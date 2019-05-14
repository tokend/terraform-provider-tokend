package httpsignatures

import (
	"crypto"

	"gitlab.com/tokend/keypair"
)

type KeypairAlgorithm struct{}

func (a KeypairAlgorithm) Sign(key interface{}, digest []byte) ([]byte, error) {
	kp, ok := key.(keypair.Full)
	if !ok {
		return nil, ErrUnexpectedKey
	}
	return kp.Sign(a.hash(digest))
}

func (a KeypairAlgorithm) hash(digest []byte) []byte {
	hash := crypto.SHA256.New()
	hash.Write(digest)
	return hash.Sum(nil)
}

func (a KeypairAlgorithm) Verify(signature Signature, digest []byte) bool {
	address, err := keypair.ParseAddress(signature.KeyID)
	if err != nil {
		return false
	}

	err = address.Verify(a.hash(digest), signature.Signature)
	if err != nil {
		return false
	}

	return true
}

func (a KeypairAlgorithm) Name() string {
	return "ed25519-sha256"
}

func init() {
	RegisterAlgorithm(KeypairAlgorithm{})
}
