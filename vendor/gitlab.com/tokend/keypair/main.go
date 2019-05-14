package keypair

import (
	"errors"
)

var (
	// ErrInvalidKey will be returned by operations when the keypair being used
	// could not be decoded.
	ErrInvalidKey = errors.New("invalid key")

	// ErrInvalidSignature is returned when the signature is invalid, either
	// through malformation or if it does not verify the message against the
	// provided public key
	ErrInvalidSignature = errors.New("signature verification failed")
)

type Address interface {
	Address() string
	Hint() [4]byte
	Verify(input []byte, signature []byte) error
}

type Full interface {
	Address
	Seed() string
	Sign(input []byte) ([]byte, error)
}
