package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// DEPRECATED
// Anyway, nobody checks the Path anywhere
type Path string
// DEPRECATED
// Anyway, nobody checks the Kind anywhere
type Kind int32
type Status int
type Response []byte

const (
	// zero for default value if not provided
	Other Kind = 1 << iota
	Unauthorized
	Network
	Runtime
)

var (
	New  = errors.New
	Wrap = errors.Wrap
)

type Error struct {
	msg    string
	cause  error
	// DEPRECATED
	// Anyway, nobody checks the path anywhere
	path   Path
	// DEPRECATED
	// Anyway, nobody checks the kind anywhere
	kind   Kind
	status Status
	body   Response
}

func (e Error) Error() string {
	if e.cause == nil {
		return e.msg
	}

	return fmt.Sprintf("%s: %s", e.msg, e.cause.Error())
}

func E(msg string, args ...interface{}) error {
	e := Error{
		msg: msg,
	}

	for _, arg := range args {
		switch arg := arg.(type) {
		case Path:
			e.path = arg
		case Kind:
			e.kind = arg
		case Status:
			e.status = arg
		case error:
			e.cause = arg
		case Response:
			e.body = arg
		}

	}

	return e
}
