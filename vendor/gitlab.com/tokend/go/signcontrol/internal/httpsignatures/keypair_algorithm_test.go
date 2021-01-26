package httpsignatures_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/signcontrol/internal/httpsignatures"
	"gitlab.com/tokend/keypair"
)

func TestKeypairAlgorithm(t *testing.T) {
	algorithm := httpsignatures.KeypairAlgorithm{}
	referenceSignature := []byte{
		0x54, 0x6e, 0xf4, 0x7b, 0xb1, 0x3e, 0x75, 0x97,
		0x7c, 0x63, 0xba, 0xe8, 0xd7, 0xe5, 0x56, 0xa8,
		0x8, 0xf7, 0x1e, 0x7b, 0xcf, 0xad, 0xd0, 0x4f,
		0xf7, 0x86, 0x62, 0xce, 0x24, 0xaf, 0x79, 0x4a,
		0x7a, 0x8d, 0x41, 0xfd, 0x18, 0x88, 0xc6, 0xd0,
		0xa1, 0xa0, 0x6e, 0x76, 0x94, 0x2a, 0x9f, 0xe6,
		0xd7, 0x28, 0x9e, 0xd3, 0x7f, 0xda, 0xb6, 0xf5,
		0x77, 0x6e, 0x9d, 0xdc, 0x75, 0x43, 0x5c, 0xb}

	kp, _ := keypair.ParseSeed("SDIJIVREVHDO2P5WSOP557WL7EBNS4LSRLHQ47B3SP4NQTY5IDPTQO3E")
	msg := []byte("hello")

	t.Run("name", func(t *testing.T) {
		assert.Equal(t, "ed25519-sha256", algorithm.Name())
	})

	t.Run("sign", func(t *testing.T) {
		sig, err := algorithm.Sign(kp, msg)
		assert.NoError(t, err)
		assert.Equal(
			t,
			referenceSignature,
			sig,
		)
	})

	t.Run("key mismatch", func(t *testing.T) {
		_, err := algorithm.Sign(1, msg)
		assert.Error(t, err)
		assert.Equal(t, err, httpsignatures.ErrUnexpectedKey)
	})

	t.Run("verify", func(t *testing.T) {
		signature := httpsignatures.Signature{
			KeyID:     kp.Address(),
			Algorithm: algorithm,
			Signature: referenceSignature,
		}
		ok := algorithm.Verify(signature, msg)
		assert.True(t, ok)
	})

	t.Run("invalid signature", func(t *testing.T) {
		signature := httpsignatures.Signature{
			KeyID:     "GC3HECZFQYBURRXWFXEKW5ULQAOYHTO5K3IFM2YJHKPI6BT3KVUEHCWB",
			Algorithm: algorithm,
			Signature: referenceSignature,
		}
		ok := algorithm.Verify(signature, msg)
		assert.False(t, ok)
	})
}
