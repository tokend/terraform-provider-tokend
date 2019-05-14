package xdrbuild

import (
	"encoding/hex"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/network"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild/internal"
	"gitlab.com/tokend/keypair"
)

type Builder struct {
	passphrase         string
	txExpirationPeriod int64
}

func NewBuilder(passphrase string, txExpire int64) *Builder {
	return &Builder{
		passphrase:         passphrase,
		txExpirationPeriod: txExpire,
	}
}

func (b *Builder) Transaction(source keypair.Address) *Transaction {
	return &Transaction{
		builder: b,
		source:  source,
	}
}

func (b *Builder) Sign(envelope *xdr.TransactionEnvelope, kp keypair.Full) (*xdr.TransactionEnvelope, error) {
	hash, err := internal.HashTX(&envelope.Tx, b.passphrase)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash tx")
	}

	if err := internal.SignEnvelope(hash, kp, envelope); err != nil {
		return nil, err
	}

	return envelope, nil
}

func (b *Builder) TXHashHex(envelope string) (string, error) {
	var tx xdr.TransactionEnvelope
	if err := xdr.SafeUnmarshalBase64(envelope, &tx); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal envelope")
	}
	hash, err := network.HashTransaction(&tx.Tx, b.passphrase)
	if err != nil {
		return "", errors.Wrap(err, "failed to hash tx")
	}
	return hex.EncodeToString(hash[:]), nil
}
