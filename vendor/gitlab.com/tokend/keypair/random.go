package keypair

import (
	"crypto/rand"

	"gitlab.com/tokend/keypair/internal"
	"golang.org/x/crypto/ed25519"
)

// Random creates a random full keypair
func Random() (Full, error) {
	_, private, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return internal.NewFull(private[32:]), nil
}
