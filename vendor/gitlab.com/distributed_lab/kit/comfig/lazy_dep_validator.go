package comfig

import (
	"reflect"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//ValidateLazyDep - goes trough all methods of the dep and tries to call them. Returns error if one of the method calls
// panicked or if methods accepts any arguments.
func ValidateLazyDep(dep interface{}) error {
	value := reflect.ValueOf(dep)
	for i := 0; i < value.NumMethod(); i++ {
		method := value.Method(i)
		if method.Type().NumIn() > 0 {
			return errors.From(errors.New("unexpected number of params for method"), logan.F{
				"num_of_params": method.Type().NumIn(),
				"name":          method.Type().Name(),
			})
		}

		err := safeCall(method)
		if err != nil {
			return errors.Wrap(err, "failed to call", logan.F{
				"name": method.Type().Name(),
			})
		}
	}

	return nil
}

func safeCall(v reflect.Value) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.Wrap(errors.FromPanic(rec), "method call panicked")
		}
	}()

	// TODO: consider checking if it returned errors
	_ = v.Call([]reflect.Value{})
	return nil
}
