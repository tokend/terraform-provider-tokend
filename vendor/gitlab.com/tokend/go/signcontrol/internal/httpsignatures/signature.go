// httpsignatures is a golang implementation of the http-signatures spec
// found at https://tools.ietf.org/html/draft-cavage-http-signatures
package httpsignatures

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

const (
	headerSignature     = "Signature"
	headerAuthorization = "Authorization"
	RequestTarget       = "(request-target)"
	authScheme          = "Signature "
)

var (
	ErrorNoSignatureHeader = errors.New("No Signature header found in request")

	signatureRegex = regexp.MustCompile(`(\w+)="([^"]*)"`)
)

// Signature is the hashed key + headers, either from a request or a signer
type Signature struct {
	KeyID     string
	Algorithm Algorithm
	Headers   HeaderList
	Signature []byte
}

// FromRequest creates a new Signature from the Request
// both Signature and Authorization http headers are supported.
func FromRequest(r *http.Request) (*Signature, error) {
	if s, ok := r.Header[headerSignature]; ok {
		return FromString(s[0])
	}
	if a, ok := r.Header[headerAuthorization]; ok {
		return FromString(strings.TrimPrefix(a[0], authScheme))
	}
	return nil, ErrorNoSignatureHeader
}

// FromString creates a new Signature from its encoded form,
// eg `keyId="a",algorithm="b",headers="c",signature="d"`
func FromString(in string) (*Signature, error) {
	var res Signature = Signature{}
	var key string
	var value string

	for _, m := range signatureRegex.FindAllStringSubmatch(in, -1) {
		key = m[1]
		value = m[2]

		if key == "keyId" {
			res.KeyID = value
		} else if key == "algorithm" {
			alg, err := GetAlgorithm(value)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get algorithm")
			}
			res.Algorithm = alg
		} else if key == "headers" {
			res.Headers = headerListFromString(value)
		} else if key == "signature" {
			signature, err := base64.StdEncoding.DecodeString(value)
			if err != nil {
				return nil, errors.Wrap(err, "failed to decode signature")
			}
			res.Signature = signature
		} else {
			return nil, errors.New(fmt.Sprintf("Unexpected key in signature '%s'", key))
		}
	}

	if len(res.Signature) == 0 {
		return nil, errors.New("Missing signature")
	}

	if len(res.KeyID) == 0 {
		return nil, errors.New("Missing keyId")
	}

	if res.Algorithm == nil {
		return nil, errors.New("Missing algorithm")
	}

	return &res, nil
}

// String returns the encoded form of the Signature
func (s Signature) String() string {
	str := fmt.Sprintf(
		`keyId="%s",algorithm="%s",signature="%s"`,
		s.KeyID,
		s.Algorithm.Name(),
		base64.StdEncoding.EncodeToString(s.Signature),
	)

	if len(s.Headers) > 0 {
		str += fmt.Sprintf(`,headers="%s"`, s.Headers.String())
	}

	return str
}

func (s Signature) calculateSignature(key interface{}, r *http.Request) ([]byte, error) {
	signingString, err := s.Headers.signingString(r)
	if err != nil {
		return nil, err
	}

	b, err := s.Algorithm.Sign(key, []byte(signingString))
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Sign this signature using the given key
func (s *Signature) sign(key interface{}, r *http.Request) error {
	sig, err := s.calculateSignature(key, r)
	if err != nil {
		return err
	}

	s.Signature = sig
	return nil
}

// IsValid validates this signature for the given key
func (s Signature) IsValid(r *http.Request) bool {
	signingString, err := s.Headers.signingString(r)
	if err != nil {
		return false
	}

	return s.Algorithm.Verify(s, []byte(signingString))
}

type HeaderList []string

func headerListFromString(list string) HeaderList {
	return strings.Split(strings.ToLower(string(list)), " ")
}

func (h HeaderList) String() string {
	return strings.ToLower(strings.Join(h, " "))
}

func (h HeaderList) hasDate() bool {
	for _, header := range h {
		if header == "date" {
			return true
		}
	}

	return false
}

func (h HeaderList) signingString(req *http.Request) (string, error) {
	lines := []string{}

	for _, header := range h {
		if header == RequestTarget {
			lines = append(lines, requestTargetLine(req))
		} else {
			line, err := headerLine(req, header)
			if err != nil {
				return "", err
			}
			lines = append(lines, line)
		}
	}

	return strings.Join(lines, "\n"), nil
}

func requestTargetLine(req *http.Request) string {
	var url string = ""
	if req.URL != nil {
		url = req.URL.RequestURI()
	}

	return fmt.Sprintf("%s: %s", RequestTarget, RequestTargetValue(strings.ToLower(req.Method), url))
}

func RequestTargetValue(method, url string) string {
	return fmt.Sprintf("%s %s", strings.ToLower(method), url)
}

func headerLine(req *http.Request, header string) (string, error) {

	if value := req.Header.Get(header); value != "" {
		return fmt.Sprintf("%s: %s", header, value), nil
	}

	return "", errors.New(fmt.Sprintf("Missing required header '%s'", header))
}
