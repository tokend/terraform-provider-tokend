package keypair

import (
	"crypto/rand"
	"github.com/pkg/errors"
	"gitlab.com/tokend/keypair/internal"
	"golang.org/x/crypto/ed25519"
)

// Random creates a random full keypair
func Random() (Full, error) {
	seed := make([]byte, ed25519.SeedSize)
	// rand.Read guaranties that we'll read full seed, if there is no error
	_, err := rand.Read(seed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read random seed")
	}

	return internal.NewFull(seed), nil
}
