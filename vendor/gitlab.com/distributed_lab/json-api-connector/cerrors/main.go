package cerrors

import (
	"fmt"

	"github.com/pkg/errors"
)

type Path string
type Status int
type Response []byte

var (
	New  = errors.New
	Wrap = errors.Wrap
)

type Error struct {
	msg    string
	cause  error
	path   Path
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
