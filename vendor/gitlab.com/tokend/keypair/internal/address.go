package internal

import (
	"gitlab.com/tokend/keypair/internal/strkey"
	"golang.org/x/crypto/ed25519"
)

type Address struct {
	public ed25519.PublicKey
}

func NewAddress(public []byte) *Address {
	if len(public) != ed25519.PublicKeySize {
		panic("unexpected key length")
	}
	return &Address{
		public: public,
	}
}

func (kp *Address) Address() string {
	return strkey.MustEncode(strkey.VersionByteAccountID, kp.public)
}

func (kp *Address) Hint() (r [4]byte) {
	copy(r[:], kp.public[28:])
	return
}

func (kp *Address) Verify(input []byte, sig []byte) error {
	if len(sig) != ed25519.SignatureSize {
		return ErrInvalidSignature
	}

	if !ed25519.Verify(kp.public, input, sig) {
		return ErrInvalidSignature
	}
	return nil
}
