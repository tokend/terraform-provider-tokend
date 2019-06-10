package httpsignatures

import (
	"net/http"
	"strings"
)

// Signer is used to create a signature for a given request.
type Signer struct {
	algorithm Algorithm
	headers   HeaderList
}

func NewSigner(algorithm Algorithm, headers ...string) *Signer {
	hl := HeaderList{}

	for _, header := range headers {
		hl = append(hl, strings.ToLower(header))
	}

	return &Signer{
		algorithm: algorithm,
		headers:   hl,
	}
}

// SignRequest adds a http signature to the Signature: HTTP Header
func (s Signer) SignRequest(keyId string, key interface{}, r *http.Request) error {
	sig, err := s.buildSignature(keyId, key, r)
	if err != nil {
		return err
	}

	r.Header.Add(headerSignature, sig.String())

	return nil
}

func (s Signer) buildSignature(keyId string, key interface{}, r *http.Request) (*Signature, error) {
	sig := &Signature{
		KeyID:     keyId,
		Algorithm: s.algorithm,
		Headers:   s.headers,
	}

	err := sig.sign(key, r)
	if err != nil {
		return nil, err
	}

	return sig, nil
}
