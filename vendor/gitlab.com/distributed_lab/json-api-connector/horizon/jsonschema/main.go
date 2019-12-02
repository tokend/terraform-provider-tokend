package jsonschema

import (
	"reflect"

	"github.com/alecthomas/jsonschema"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gobuffalo/packr/v2"
	xdr "github.com/nullstyle/go-xdr/xdr3"
	. "github.com/xeipuuv/gojsonschema"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/regources"
)

//go:generate packr2
func New(endpoint string) (*Schema, error) {
	schemasBox := packr.New("schemas", "./schemas")

	raw, err := schemasBox.Find(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find strigng for schema")
	}

	return NewSchema(NewBytesLoader(raw))
}

func NewFromRegources(resource interface{}) (*Schema, error) {
	schema := newReflector().Reflect(resource)
	return NewSchema(NewGoLoader(schema))
}

func newReflector() *jsonschema.Reflector {
	return &jsonschema.Reflector{
		TypeMapper: func(i reflect.Type) *jsonschema.Type {
			switch {
			case i.Implements(reflect.TypeOf(new(regources.Flagger)).Elem()), i.Implements(reflect.TypeOf(new(xdr.Enum)).Elem()):
				return &jsonschema.Type{
					Type: "object",
					Properties: map[string]*jsonschema.Type{
						"name": {
							Type: "string",
						},
						"value": {
							Type: "integer",
						},
					},
				}
			case i == reflect.TypeOf(new(regources.Amount)).Elem():
				return &jsonschema.Type{
					Type: "string",
				}
			case i == reflect.TypeOf(new(regources.Details)).Elem():
				return &jsonschema.Type{
					Type: "object",
				}
			case i == reflect.TypeOf(regources.Included{}):
				return &jsonschema.Type{
					Type: "array",
				}
			}
			return nil
		},
	}
}

func EnsureValid(target interface{}, response []byte) error {
	schema, err := NewFromRegources(target)
	if err != nil {
		return errors.Wrap(err, "failed to create schema from regource")
	}

	res, err := schema.Validate(NewBytesLoader(response))
	if err != nil {
		return errors.Wrap(err, "failed to validate")
	}

	if !res.Valid() {
		validationErr := validation.Errors{}
		for _, err := range res.Errors() {
			validationErr[err.Field()] = errors.From(errors.New(err.Description()), logan.F(err.Details()))
		}
		return validationErr
	}

	return nil
}
