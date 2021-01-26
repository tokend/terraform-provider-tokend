package errors

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v2/fields"
)

// F type is for fields, connected to `withFields` error.
//
// DEPRECATED: Use logan/v3 instead
type F map[string]interface{}

// WithField creates new `F` fields map and add provided key-value pair into it
// using Add method.
//
// DEPRECATED: Use logan/v3 instead
func WithField(key string, value interface{}) F {
	result := make(F)
	result.Add(key, value)
	return result
}

// Add tries to extract fields from `value`, if `value` implements fields.Provider interface:
//
//		type Provider interface {
//			GetLoganFields() map[string]interface{}
//		}
//
// And adds these fields using AddFields.
// If `value` does not implement Provider - a single key-value pair is added.
//
// DEPRECATED: Use logan/v3 instead
func (f F) Add(key string, value interface{}) F {
	return f.AddFields(fields.Obtain(key, value))
}

// AddFields returns `F` map, which contains key-values from both maps.
// If both maps has some key - the value from the `newF` will be used.
//
// DEPRECATED: Use logan/v3 instead
func (f F) AddFields(newF F) F {
	return F(fields.Merge(f, newF))
}

// ToError returns new error with `message` and `f` fields.
//
// DEPRECATED: Use logan/v3 instead
func (f F) ToError(message string) error {
	return &withFields{
		errors.New(message),
		f,
	}
}

// Wrap wraps `base` error with `message` and adds `f` fields to the error.
// Returns nil if `base` is nil, which copies the `errors.Wrap` behaviour.
//
// DEPRECATED: Use logan/v3 instead
func (f F) Wrap(base error, message string) error {
	if base == nil {
		return nil
	}

	return &withFields{
		errors.Wrap(base, message),
		f,
	}
}
