package signcontrol

import "errors"

var (
	// ErrNilRequest internal error which will be panicked
	ErrNilRequest = &Error{"nil request passed"}
	// ErrDateMalformed unable to parse date header value according to RFC
	ErrDateMalformed = &Error{"date header is malformed"}
	// ErrExpired provided date header value is below acceptable threshold
	ErrExpired = &Error{"expired signature"}
	// ErrSignature signature fails public key verification
	ErrSignature = &Error{"signature is not valid"}

	// errors from legacy signature flow:
	ErrNotSigned  = &Error{"request is not signed"}
	ErrValidUntil = &Error{"valid until is not valid"}
	ErrSignerKey  = &Error{"signer key is not valid"}
	// DEPRECATED: for no obvious reason it was defined here and used by doorman
	ErrNotAllowed = errors.New("not allowed")
)

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) BadRequest() bool {
	return true
}
