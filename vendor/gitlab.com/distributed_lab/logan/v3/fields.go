package logan

import "gitlab.com/distributed_lab/logan/v3/fields"

// F type is for fields, connected to `withFields` error.
type F map[string]interface{}

// WithField creates new `F` fields map and add provided key-value pair into it
// using Add method.
//
// DEPRECATED: Use F{key: value} directly instead. Fields expanding is now happening on adding to log Entry.
func Field(key string, value interface{}) F {
	result := make(F)
	return result.Add(key, value)
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
// Add doesn't change any of maps - only creates a new one.
//
// DEPRECATED: Use `fields[key] = value` instead. Fields expanding is now happening on adding to log Entry.
func (f F) Add(key string, value interface{}) F {
	return f.AddFields(fields.Obtain(key, value))
}

// AddFields returns `F` map, which contains key-values from both maps.
// If both maps has some key - the value from the `newF` will be used.
//
// AddFields doesn't change any of maps - only creates a new one.
//
// DEPRECATED: Use Merge method instead (it's same, but more obvious that it doesn't mutate the instance).
func (f F) AddFields(newF F) F {
	return F(fields.Merge(f, newF))
}

// Merge returns `F` map, which contains key-values from both maps.
// If both maps has some key - the value from the `newF` will be used.
//
// Merge doesn't change any of maps - only creates a new one.
func (f F) Merge(newF F) F {
	return F(fields.Merge(f, newF))
}
