package internal

import (
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/keypair"
)

func SignEnvelope(hash [32]byte, kp keypair.Full, envelope *xdr.TransactionEnvelope) error {
	sig, err := kp.Sign(hash[:])
	if err != nil {
		return err
	}

	envelope.Signatures = append(envelope.Signatures, xdr.DecoratedSignature{
		Hint:      xdr.SignatureHint(kp.Hint()),
		Signature: xdr.Signature(sig),
	})

	return nil
}
