package xdrbuild

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/keypair"
)

func TestTransaction_Marshal(t *testing.T) {
	passphrase := "passphrase"
	builder := NewBuilder(passphrase, 100)
	source := keypair.MustParseSeed("SDW646VAMVV2R5H3WO22N4XLT3ZXMT46DCAHW2PQAYL4WMQKD3RYQIK7")

	t.Run("source", func(t *testing.T) {
		envelope, err := builder.Transaction(source).Marshal()
		if err != nil {
			t.Fatal(err)
		}
		var got xdr.TransactionEnvelope
		if err := xdr.SafeUnmarshalBase64(envelope, &got); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, source.Address(), got.Tx.SourceAccount.Address())
	})

	t.Run("no signature", func(t *testing.T) {
		envelope, err := builder.Transaction(source).Marshal()
		if err != nil {
			t.Fatal(err)
		}
		var got xdr.TransactionEnvelope
		if err := xdr.SafeUnmarshalBase64(envelope, &got); err != nil {
			t.Fatal(err)
		}
		assert.Len(t, got.Signatures, 0)
	})

	t.Run("signatures exists", func(t *testing.T) {
		envelope, err := builder.Transaction(source).Sign(source).Sign(source).Marshal()
		if err != nil {
			t.Fatal(err)
		}
		var got xdr.TransactionEnvelope
		if err := xdr.SafeUnmarshalBase64(envelope, &got); err != nil {
			t.Fatal(err)
		}
		assert.Len(t, got.Signatures, 2)
	})

	t.Run("implicit salt", func(t *testing.T) {
		envelope, err := builder.Transaction(source).Marshal()
		if err != nil {
			t.Fatal(err)
		}
		var got xdr.TransactionEnvelope
		if err := xdr.SafeUnmarshalBase64(envelope, &got); err != nil {
			t.Fatal(err)
		}
		assert.NotZero(t, got.Tx.Salt)
	})

	t.Run("explicit salt", func(t *testing.T) {
		envelope, err := builder.Transaction(source).Salt(42).Marshal()
		if err != nil {
			t.Fatal(err)
		}
		var got xdr.TransactionEnvelope
		if err := xdr.SafeUnmarshalBase64(envelope, &got); err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, 42, got.Tx.Salt)
	})

	t.Run("implicit time bounds", func(t *testing.T) {
		envelope, err := builder.Transaction(source).Marshal()
		if err != nil {
			t.Fatal(err)
		}
		var got xdr.TransactionEnvelope
		if err := xdr.SafeUnmarshalBase64(envelope, &got); err != nil {
			t.Fatal(err)
		}
		assert.NotZero(t, got.Tx.TimeBounds.MaxTime)
	})

	t.Run("explicit time bounds", func(t *testing.T) {
		envelope, err := builder.Transaction(source).TimeBounds(10, 20).Marshal()
		if err != nil {
			t.Fatal(err)
		}
		var got xdr.TransactionEnvelope
		if err := xdr.SafeUnmarshalBase64(envelope, &got); err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, 10, got.Tx.TimeBounds.MinTime)
		assert.EqualValues(t, 20, got.Tx.TimeBounds.MaxTime)
	})
}
