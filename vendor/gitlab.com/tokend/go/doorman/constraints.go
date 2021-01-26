package doorman

import (
	"gitlab.com/tokend/go/resources"
	"net/http"

	"gitlab.com/tokend/go/signcontrol"
)

type SignerOfExt interface {
	Check(signer resources.Signer) bool
}

func ClearSignerOf(address string) SignerConstraint {
	return func(r *http.Request, doorman Doorman) error {
		return signerOf(address, r, doorman)
	}
}

// If ext passed - SignerOf overrides default doorman constraints
func SignerOf(address string, ext ...SignerOfExt) SignerConstraint {
	return func(r *http.Request, doorman Doorman) error {
		if len(ext) == 0 {
			ext = doorman.DefaultSignerOfConstraints()
		}

		return signerOf(address, r, doorman, ext...)
	}
}

func signerOf(address string, r *http.Request, doorman Doorman, ext ...SignerOfExt) error {
	signer, err := signcontrol.CheckSignature(r)
	if err != nil {
		return err
	}

	signers, err := doorman.AccountSigners(address)
	if err != nil {
		return err
	}

	for _, accountSigner := range signers {
		if accountSigner.AccountID == signer && accountSigner.Weight > 0 && checkConstraints(accountSigner, ext) {
			return nil
		}
	}
	return ErrNotAllowed
}

func SignatureOf(address string) SignerConstraint {
	return SignerOf(address)
}

func checkConstraints(accountSigner resources.Signer, constraints []SignerOfExt) bool {
	for _, c := range constraints {
		if !c.Check(accountSigner) {
			return false
		}
	}
	return true
}
