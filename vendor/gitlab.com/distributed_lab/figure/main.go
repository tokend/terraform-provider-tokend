package figure

import (
	"reflect"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	keyTag   = "fig"
	ignore   = "-"        // Field is omitted
	required = "required" // Value must be presented in config
	nonZero  = "non_zero" // Field must have non zero value. Note: zero value for pointer is nil
)

var (
	ErrNonZeroValue  = errors.New("you must set non zero value to this field")
	ErrRequiredValue = errors.New("you must set the value in field")
	ErrNoHook        = errors.New("no such hook")
	ErrNotValid      = errors.New("not valid value")
)

type Validatable interface {
	// Validate validates the data and returns an error if validation fails.
	Validate() error
}

// Hook signature for custom hooks.
// Takes raw value expected to return target value
type Hook func(value interface{}) (reflect.Value, error)

// Hooks is mapping raw type -> `Hook` instance
type Hooks map[string]Hook

// With accepts hooks to be used for figuring out target from raw values.
// `BaseHooks` will be used implicitly if no hooks are provided
func (f *Figurator) With(hooks ...Hooks) *Figurator {
	merged := Hooks{}
	for _, partial := range hooks {
		for key, hook := range partial {
			merged[key] = hook
		}
	}
	f.hooks = merged
	return f
}

// Figurator holds state for chained call
type Figurator struct {
	values map[string]interface{}
	hooks  Hooks
	target interface{}
}

// Out is main entry point for package, used to start figure out chain
func Out(target interface{}) *Figurator {
	return &Figurator{
		target: target,
	}
}

// From takes raw config values to be used in figure out process
func (f *Figurator) From(values map[string]interface{}) *Figurator {
	f.values = values
	return f
}

// Please exit point for figure out chain.
// Will modify target partially in case of error
func (f *Figurator) Please() error {
	// if hooks were not explicitly set use default
	if len(f.hooks) == 0 {
		f.With(BaseHooks)
	}
	vle := reflect.Indirect(reflect.ValueOf(f.target))
	tpe := vle.Type()
	for fi := 0; fi < tpe.NumField(); fi++ {
		fieldType := tpe.Field(fi)
		fieldValue := vle.Field(fi)
		if err := f.SetField(fieldValue, fieldType, keyTag); err != nil {
			return errors.Wrap(err, "failed to set field", logan.F{"field": fieldType.Name})
		}
	}

	if data, ok := f.target.(Validatable); ok {
		return data.Validate()
	}

	return nil
}

func (f *Figurator) SetField(fieldValue reflect.Value, field reflect.StructField, keyTag string) error {
	tag, err := parseFieldTag(field, keyTag)
	if err != nil {
		return errors.Wrap(err, "failed to parse tag", logan.F{"tag": tag.Key})
	}

	if tag == nil {
		return nil
	}

	hook, hasHook := f.hooks[field.Type.String()]
	isSet := false
	raw, hasRaw := f.values[tag.Key]

	logFields := logan.F{
		"field": field.Name,
		"type":  field.Type.Name(),
		"raw":   raw,
	}

	if !hasHook && fieldValue.Kind() != reflect.Struct {
		return errors.Wrap(ErrNoHook, "failed to find hook", logFields)
	}
	if hasRaw && !hasHook && fieldValue.Kind() == reflect.Struct {
		rawValues, err := cast.ToStringMapE(raw)
		if err != nil {
			return errors.Wrap(ErrNotValid, "failed to cast to map", logFields)
		}
		err = please(rawValues, f.hooks, fieldValue)
		if err != nil {
			return errors.Wrap(err, "failed to parse internal struct", logFields)
		}
		isSet = true
	}

	if hasRaw && hasHook {
		value, err := hook(raw)
		if err != nil {
			return errors.Wrap(err, "failed to figure out", logFields)
		}
		fieldValue.Set(value)

		isSet = true
	}

	if tag.Required && !isSet {
		return errors.Wrap(ErrRequiredValue, "failed to get value for this field", logan.F{"field": field.Name})
	}

	if tag.NonZero && isZero(fieldValue) {
		return errors.Wrap(ErrNonZeroValue, "field has a zero value", logan.F{"field": field.Name})
	}

	return nil
}

func please(rawValues map[string]interface{}, hooks Hooks, fieldValue reflect.Value) error {

	vle := reflect.Indirect(fieldValue)
	if !vle.IsValid() {
		return errors.Wrap(ErrNotValid, "value is not valid", logan.F{"field": fieldValue.String()})
	}
	tpe := vle.Type()

	for fi := 0; fi < tpe.NumField(); fi++ {
		fieldType := tpe.Field(fi)
		fieldValue := vle.Field(fi)

		tag, err := parseFieldTag(fieldType, keyTag)
		if err != nil {
			return errors.Wrap(err, "failed to parse tag", logan.F{"tag": tag.Key})
		}

		if err := setField(fieldValue, fieldType, keyTag, rawValues, hooks); err != nil {
			return errors.Wrap(err, "failed to set field", logan.F{"field": fieldType.Name})
		}
	}

	return nil
}

func setField(fieldValue reflect.Value, field reflect.StructField, keyTag string, values map[string]interface{}, hooks Hooks) error {
	tag, err := parseFieldTag(field, keyTag)
	if err != nil {
		return errors.Wrap(err, "failed to parse tag", logan.F{"tag": tag.Key})
	}
	if tag == nil {
		return nil
	}

	hook, ok := hooks[field.Type.String()]
	isSet := false
	raw, hasRaw := values[tag.Key]
	if hasRaw && !ok {
		rawValues, err := cast.ToStringMapE(raw)
		if err != nil {
			return errors.Wrap(ErrRequiredValue, "failed to get value for this field", logan.F{"field": field.Name})
		}
		er := please(rawValues, hooks, fieldValue)
		isSet = true
		return er
	}

	if hasRaw {
		value, err := hook(raw)
		if err != nil {
			return errors.Wrap(err, "failed to figure out", logan.F{"hook": field.Type.String(), "value": raw})
		}
		fieldValue.Set(value)
		isSet = true
	}

	if tag.Required && !isSet {
		return errors.Wrap(ErrRequiredValue, "failed to get value for this field", logan.F{"field": field.Name})
	}

	if tag.NonZero && isZero(fieldValue) {
		return errors.Wrap(ErrNonZeroValue, "field has a zero value", logan.F{"field": field.Name})
	}

	return nil
}

// IsZeroer is used to check whether an object is zero to
// determine whether it should be omitted when marshaling
// with the omitempty flag. One notable implementation
// is time.Time.
type IsZeroer interface {
	IsZero() bool
}

func isZero(v reflect.Value) bool {
	kind := v.Kind()
	if z, ok := v.Interface().(IsZeroer); ok {
		if (kind == reflect.Ptr || kind == reflect.Interface) && v.IsNil() {
			return true
		}
		return z.IsZero()
	}
	switch kind {
	case reflect.String:
		return len(v.String()) == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Slice:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Struct:
		vt := v.Type()
		for i := v.NumField() - 1; i >= 0; i-- {
			if vt.Field(i).PkgPath != "" {
				continue // Private field
			}
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}
