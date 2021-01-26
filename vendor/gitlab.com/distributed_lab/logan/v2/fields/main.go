package fields

// Provider if passed as a field value (see `logan.Entity.WithField()` and `errors.F.Add()`),
// will be transformed into map with multiple key-values:
// each key provided in the return map of implementation of `GetLoganFields()`
// will be prefixed with the key of the whole entity, using `_` delimiter.
// See example in tests.
//
// DEPRECATED: Use logan/v3 instead
type Provider interface {
	GetLoganFields() map[string]interface{}
}

// Fields is just to simplify the code readability
//
// DEPRECATED
type fields map[string]interface{}

// Obtain tries to extract fields from `value`, if `value` implements Provider.
//
//		type Provider interface {
//			GetLoganFields() map[string]interface{}
//		}
//
// If `value` does not implement Provider - a map with 1 key and plain value will be returned.
//
// DEPRECATED: Use logan/v3 instead
func Obtain(entityName string, value interface{}) map[string]interface{} {
	fieldedEntity, ok := value.(Provider)

	if ok {
		return obtain(entityName, fieldedEntity)
	} else {
		return map[string]interface{}{
			entityName: value,
		}
	}
}

// DEPRECATED
func obtain(entityName string, fieldedEntity Provider) (result fields) {
	result = make(fields)

	defer func() {
		// This defer is added to simplify catching the situation of `fieldedEntity` being nil of some specific type.
		if r := recover(); r != nil {
			// `fieldedEntity` is nil
			result = make(fields)
			result[entityName] = nil
		}
	}()

	for key, value := range fieldedEntity.GetLoganFields() {
		compositeKey := entityName + "_" + key

		result = Merge(result, Obtain(compositeKey, value))
	}

	return result
}

// Merge merges two instances of `map[string]interface{}`.
// You can pass here as arguments any types, which are in fact `map[string]interface{}`.
// If both maps has some key - the value from the `f2` will be used.
//
// DEPRECATED: Use logan/v3 instead
func Merge(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	result := make(fields, len(m1)+len(m2))

	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}

	return result
}
