package cerrors

import (
	"net/http"
)

func NotFound(err error) bool {
	cerr, ok := err.(Error)

	if !ok {
		return false
	}

	return cerr.status == http.StatusNotFound
}

func NoContent(err error) bool {
	cerr, ok := err.(Error)

	if !ok {
		return false
	}

	return cerr.status == http.StatusNoContent
}
