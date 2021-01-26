package __old

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

func TestSetOptions_XDR(t *testing.T) {
	t.Run("successfully created operation and pass validation", func(t *testing.T) {
		var lowThreshold, medThreshold, highThreshold uint32 = 10, 100, 250

		signer := &Signer{
			PublicKey:  "GC2252WQCUA2TUS7KBL7CVVBW2LJANJS7HC3J4VYP7JMXZHDR5MRZ757",
			Weight:     4,
			SignerType: 4,
			Identity:   123,
			Name:       "signer_name",
		}

		setOps := SetOptions{
			Signer:        signer,
			LowThreshold:  &lowThreshold,
			MedThreshold:  &medThreshold,
			HighThreshold: &highThreshold,
		}

		assert.NoError(t, setOps.Validate())

		got, err := setOps.XDR()
		assert.NoError(t, err)

		body := got.Body.SetOptionsOp
		assert.Equal(t, signer.PublicKey, body.Signer.PubKey.Address())
		assert.Equal(t, signer.SignerType, uint32(body.Signer.SignerType))
		assert.Equal(t, signer.Weight, uint32(body.Signer.Weight))
		assert.Equal(t, signer.Identity, uint32(body.Signer.Identity))
		assert.Equal(t, signer.Name, string(body.Signer.Name))
		assert.Equal(t, lowThreshold, uint32(*body.LowThreshold))
		assert.Equal(t, medThreshold, uint32(*body.MedThreshold))
		assert.Equal(t, highThreshold, uint32(*body.HighThreshold))
	})

	t.Run("not valid thresholds", func(t *testing.T) {
		var lowThreshold, medThreshold, highThreshold uint32 = 256, 257, 258

		setOps := SetOptions{
			LowThreshold:  &lowThreshold,
			MedThreshold:  &medThreshold,
			HighThreshold: &highThreshold,
		}

		assert.Error(t, setOps.Validate())
	})

	t.Run("not valid signer", func(t *testing.T) {
		setOps := SetOptions{
			Signer: &Signer{},
		}

		assert.Error(t, setOps.Validate())
	})

	t.Run("AddSigner", func(t *testing.T) {
		setOpts := AddSigner(
			"GC2252WQCUA2TUS7KBL7CVVBW2LJANJS7HC3J4VYP7JMXZHDR5MRZ757",
			"Vasya Petrovich",
			228,
			uint32(xdr.SignerTypeAccountManager),
			123,
		)

		assert.NoError(t, setOpts.Validate())
		xdrOp, err := setOpts.XDR()
		assert.NoError(t, err)

		assert.Equal(t, setOpts.Signer.PublicKey, xdrOp.Body.SetOptionsOp.Signer.PubKey.Address())
		assert.Equal(t, setOpts.Signer.Weight, uint32(xdrOp.Body.SetOptionsOp.Signer.Weight))
		assert.Equal(t, setOpts.Signer.Name, string(xdrOp.Body.SetOptionsOp.Signer.Name))
		assert.Equal(t, setOpts.Signer.Identity, uint32(xdrOp.Body.SetOptionsOp.Signer.Identity))
		assert.Equal(t, setOpts.Signer.SignerType, uint32(xdrOp.Body.SetOptionsOp.Signer.SignerType))
	})

	t.Run("DeleteSigner", func(t *testing.T) {
		setOpts := DeleteSigner("GC2252WQCUA2TUS7KBL7CVVBW2LJANJS7HC3J4VYP7JMXZHDR5MRZ757")

		assert.NoError(t, setOpts.Validate())

		xdrOp, err := setOpts.XDR()
		assert.NoError(t, err)

		assert.Equal(t, setOpts.Signer.PublicKey, xdrOp.Body.SetOptionsOp.Signer.PubKey.Address())
		assert.Equal(t, uint32(0), uint32(xdrOp.Body.SetOptionsOp.Signer.Weight))
	})

	t.Run("SetThresholds", func(t *testing.T) {
		setOpts := SetThresholds(255, 10, 125, 200)

		assert.NoError(t, setOpts.Validate())

		xdrOp, err := setOpts.XDR()
		assert.NoError(t, err)
		assert.Equal(t, *setOpts.MasterWeight, uint32(*xdrOp.Body.SetOptionsOp.MasterWeight))
		assert.Equal(t, *setOpts.LowThreshold, uint32(*xdrOp.Body.SetOptionsOp.LowThreshold))
		assert.Equal(t, *setOpts.MedThreshold, uint32(*xdrOp.Body.SetOptionsOp.MedThreshold))
		assert.Equal(t, *setOpts.HighThreshold, uint32(*xdrOp.Body.SetOptionsOp.HighThreshold))
	})
}
