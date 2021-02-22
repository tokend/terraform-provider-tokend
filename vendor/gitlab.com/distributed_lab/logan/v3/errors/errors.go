package errors

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3/fields"
	"fmt"
)

const (
	UnknownStack = "unknown"
)

// FromPanic extracts the err from the result of a recover() call.
// If rec is not actually an error - a new error will be created, formatting the `rec` as "%s".
func FromPanic(rec interface{}) error {
	err, ok := rec.(error)
	if !ok {
		err = errors.Errorf("%s", rec)
	}

	if stack := GetStack(err); stack == UnknownStack {
		// No stack is connected to the error from recover
		err = errors.WithStack(err)
	}

	return err
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(msg string) error {
	return errors.New(msg)
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	return errors.WithStack(err)
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
//
// Fields can optionally be added. If provided multiple - fields will be merged.
//
// If err is nil, Wrap returns nil.
func Wrap(err error, msg string, errorFields... map[string]interface{}) error {
	wrapped := errors.Wrap(err, msg)
	if wrapped == nil {
		return nil
	}

	var mergedFields map[string]interface{}
	for _, f := range errorFields {
		mergedFields = fields.Merge(mergedFields, f)
	}

	return &withFields{
		wrapped,
		mergedFields,
	}
}

// From returns an error annotating err with a stack trace
// at the point From is called, and the provided fields.
//
// If err is nil, From returns nil.
func From(err error, fields map[string]interface{}) error {
	withStack := errors.WithStack(err)

	if withStack == nil {
		return nil
	}

	return &withFields{
		withStack,
		fields,
	}
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	return errors.Cause(err)
}

// GetFields returns the underlying fields of the error and its nested cause-errors, if possible.
// An error value has fields if it (or any of its nested cause) implements the following interface:
//
//     type fieldsProvider interface {
//            GetFields() F
//     }
//
// If the error and all of its nested causes do not implement GetFields, empty fields map will
// be returned.
func GetFields(err error) map[string]interface{} {
	type fieldsProvider interface {
		GetFields() eFields
	}

	type causer interface {
		Cause() error
	}

	mergedResult := eFields{}
	for err != nil {
		fError, ok := err.(fieldsProvider)
		if ok {
			mergedResult = fields.Merge(mergedResult, fError.GetFields())
		}

		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return mergedResult
}

// GetStack returns the string representation of stacktrace of the
// provided error (see getErrorStack func for stack retrieving details).
//
// If the provided error does not provide stack,
// GetStack will try to retrieve stack from its causer,
// then from causer of its cause, and so one.
//
// If no stack was provided by any of the causers,
// the value of `UnknownStack` const will be returned.
func GetStack(err error) string {
	type causer interface {
		Cause() error
	}

	for err != nil {
		stack := getErrorStack(err)
		if stack != UnknownStack {
			return stack
		}

		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return UnknownStack
}

// GetErrorStack returns the stack, as a string, if one can be extracted from `err`.
// Currently 2 interfaces of stack providing are supported:
//
//		type stackTraceProvider interface {
//			StackTrace() errors.StackTrace
//		}
//
//		and
//
//		type stackProvider interface {
//			Stack() []byte
//		}
//
// The first one is implemented by errors from pkg/errors and
// the second one - from go-errors.
func getErrorStack(err error) string {
	// pkg/errors
	type stackTraceProvider interface {
		StackTrace() errors.StackTrace
	}
	if s, ok := err.(stackTraceProvider); ok {
		return fmt.Sprintf("%+v", s.StackTrace())
	}

	// go-errors
	type stackProvider interface {
		Stack() []byte
	}
	if s, ok := err.(stackProvider); ok {
		return string(s.Stack())
	}

	return UnknownStack
}

type withFields struct {
	error
	eFields
}

func (w withFields) Error() string {
	return w.error.Error()
}

func (w withFields) GetFields() eFields {
	return w.eFields
}

func (w withFields) Cause() error {
	return w.error
}
