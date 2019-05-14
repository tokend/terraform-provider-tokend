package xdrbuild

import (
	"math/rand"

	"time"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/network"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild/internal"
	"gitlab.com/tokend/keypair"
)

type Operation = internal.Operation

type Transaction struct {
	builder     *Builder
	source      keypair.Address
	operations  []internal.Operation
	signers     []keypair.Full
	salt        *int64
	timebounds  *internal.TimeBounds
}

func (tx *Transaction) Salt(v int64) *Transaction {
	tx.salt = &v
	return tx
}

func (tx *Transaction) TimeBounds(min, max int64) *Transaction {
	tx.timebounds = &internal.TimeBounds{
		Min: min, Max: max,
	}
	return tx
}

func (tx *Transaction) Marshal() (string, error) {
	envelope := xdr.TransactionEnvelope{}

	// source
	if err := envelope.Tx.SourceAccount.SetAddress(tx.source.Address()); err != nil {
		return "", errors.Wrap(err, "failed to set source account")
	}

	// salt
	if tx.salt == nil {
		envelope.Tx.Salt = xdr.Salt(rand.Int63())
	} else {
		envelope.Tx.Salt = xdr.Salt(*tx.salt)
	}

	// time bounds
	if tx.timebounds == nil {
		envelope.Tx.TimeBounds = xdr.TimeBounds{
			MaxTime: xdr.Uint64(time.Now().Unix() + tx.builder.txExpirationPeriod),
		}
	} else {
		envelope.Tx.TimeBounds = xdr.TimeBounds{
			MinTime: xdr.Uint64(tx.timebounds.Min),
			MaxTime: xdr.Uint64(tx.timebounds.Max),
		}
	}

	// marshal operations
	for _, op := range tx.operations {
		if validator, ok := op.(internal.Validatable); ok {
			if err := validator.Validate(); err != nil {
				return "", errors.Wrap(err, "failed to validate op")
			}
		}
		xdrop, err := op.XDR()
		if err != nil {
			return "", errors.Wrap(err, "failed to marshal op")
		}
		envelope.Tx.Operations = append(envelope.Tx.Operations, *xdrop)
	}

	// hash
	// TODO check passphrase is set
	hash, err := network.HashTransaction(&envelope.Tx, tx.builder.passphrase)
	if err != nil {
		return "", errors.Wrap(err, "failed to hash transaction")
	}

	// sign
	for _, kp := range tx.signers {
		sig, err := kp.Sign(hash[:])
		if err != nil {
			return "", errors.Wrap(err, "failed to sign hash")
		}

		envelope.Signatures = append(envelope.Signatures, xdr.DecoratedSignature{
			Hint:      xdr.SignatureHint(kp.Hint()),
			Signature: xdr.Signature(sig),
		})
	}

	encoded, err := xdr.MarshalBase64(&envelope)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal tx envelope")
	}
	return encoded, nil
}

func (tx *Transaction) Sign(kp keypair.Full) *Transaction {
	tx.signers = append(tx.signers, kp)
	return tx
}

func (tx *Transaction) Op(op Operation) *Transaction {
	tx.operations = append(tx.operations, op)
	return tx
}
