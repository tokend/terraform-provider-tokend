package internal

import (
	"bytes"

	"gitlab.com/tokend/keypair/internal/strkey"
	"golang.org/x/crypto/ed25519"
)

type Full struct {
	raw     []byte
	private ed25519.PrivateKey
	public  ed25519.PublicKey
}

func NewFull(raw []byte) *Full {
	public, private, err := ed25519.GenerateKey(bytes.NewReader(raw))
	if err != nil {
		panic(err)
	}

	return &Full{
		private: private,
		public:  public,
		raw:     raw,
	}
}

func (kp *Full) Address() string {
	return strkey.MustEncode(strkey.VersionByteAccountID, kp.public)
}

func (kp *Full) Hint() (r [4]byte) {
	copy(r[:], kp.public[28:])
	return
}

func (kp *Full) Seed() string {
	return strkey.MustEncode(strkey.VersionByteSeed, kp.raw)
}

func (kp *Full) Verify(input []byte, sig []byte) error {
	if len(sig) != 64 {
		return ErrInvalidSignature
	}

	if !ed25519.Verify(kp.public, input, sig) {
		return ErrInvalidSignature
	}
	return nil
}

func (kp *Full) Sign(input []byte) ([]byte, error) {
	return ed25519.Sign(kp.private, input)[:], nil
}