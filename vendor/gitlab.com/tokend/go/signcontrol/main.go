package signcontrol

import (
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/hash"
	"gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/signcontrol/internal/httpsignatures"
	"gitlab.com/tokend/go/xdr"
)

const (
	SignatureHeader  = "x-authsignature"
	PublicKeyHeader  = "x-authpublickey"
	ValidUntilHeader = "x-authvaliduntilltimestamp"

	// SignatureAlgorithm default algorithm used to sign requests
	SignatureAlgorithm = "ed25519-sha256"

	signatureExpireAfter = 1 * time.Hour

	realRequestTargetHeader = "Real-Request-Target"
)

var (
	dateHeaders = map[string]struct{}{
		"Date":   {},
		"X-Date": {}, // https://fetch.spec.whatwg.org/#forbidden-header-name
	}
)

type Signature struct {
	Address       string
	Signer        string
	RawValidUntil string
	Signature     string
}

// IsSignedV1 - checks if v1 authorization headers is presented
// DEPRECATED
func IsSignedV1(request *http.Request) (*Signature, bool) {
	if request == nil {
		panic(ErrNilRequest)
	}

	s := Signature{
		Signer:        request.Header.Get(PublicKeyHeader),
		Signature:     request.Header.Get(SignatureHeader),
		RawValidUntil: request.Header.Get(ValidUntilHeader),
	}
	return &s, s.Signature != "" && s.Signer != "" && s.RawValidUntil != ""
}

func SignRequest(request *http.Request, kp keypair.KP) error {
	algorithm, err := httpsignatures.GetAlgorithm(SignatureAlgorithm)
	if err != nil {
		return errors.Wrap(err, "failed to get signature algorithm")
	}

	if request.Header.Get(realRequestTargetHeader) == "" {
		request.Header.Set(
			realRequestTargetHeader,
			httpsignatures.RequestTargetValue(request.Method, request.URL.RequestURI()),
		)
	}

	if request.Header.Get("date") == "" {
		request.Header.Set("date", time.Now().UTC().Format(http.TimeFormat))
	}

	signer := httpsignatures.NewSigner(algorithm, "date", realRequestTargetHeader)
	if err = signer.SignRequest(kp.Address(), kp, request); err != nil {
		return err
	}
	return nil
}

// IsSigned - checks if authorization headers is presented
func IsSigned(request *http.Request) bool {
	return request.Header.Get("signature") != "" || request.Header.Get("authorization") != ""
}

func CheckSignature(request *http.Request) (string, error) {
	// check if it v2 signature
	if IsSigned(request) {
		return checkV2(request)
	}

	sig, ok := IsSignedV1(request)
	if !ok {
		return "", ErrNotSigned
	}

	validUntil, err := strconv.ParseInt(sig.RawValidUntil, 10, 64)
	if err != nil {
		return "", ErrValidUntil
	}

	if time.Unix(validUntil, 0).Before(time.Now()) {
		return "", ErrExpired
	}

	signatureBase := "{ uri: '" + request.URL.RequestURI() + "', valid_untill: '" + sig.RawValidUntil + "'}"
	hashBase := hash.Hash([]byte(signatureBase))
	pubKey, err := keypair.Parse(sig.Signer)
	if err != nil {
		return "", ErrSignerKey
	}

	var decoratedSign xdr.DecoratedSignature
	err = xdr.SafeUnmarshalBase64(sig.Signature, &decoratedSign)
	if err != nil {
		return "", ErrSignature
	}

	if pubKey.Verify(hashBase[:], decoratedSign.Signature) != nil {
		return "", ErrSignature
	}

	return sig.Signer, nil
}

func checkV2(request *http.Request) (string, error) {
	signature, err := httpsignatures.FromRequest(request)
	if err != nil {
		// TODO check for no such algorithm
		// making error castable to signcontrol w/o exposing httpsignatures impl
		return "", &Error{err.Error()}
	}
	if ok := signature.IsValid(request); !ok {
		return "", ErrSignature
	}

	// check signature freshness if one of `dateHeaders` is signed
	for _, header := range signature.Headers {
		if _, ok := dateHeaders[http.CanonicalHeaderKey(header)]; !ok {
			continue
		}
		date := request.Header.Get(header)
		if err := checkDate(date); err != nil {
			return "", err
		}
		break
	}

	return signature.KeyID, nil
}

func checkDate(raw string) error {
	date, err := http.ParseTime(raw)
	if err != nil {
		return ErrDateMalformed
	}

	expireAt := date.Add(signatureExpireAfter)
	if expireAt.Before(time.Now()) {
		return ErrExpired
	}

	return nil
}
