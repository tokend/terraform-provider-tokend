package logan

// DEPRECATED: Use logan/v3 instead
const NilValueToLog = "<nil>"

// DEPRECATED: Use logan/v3 instead
type FieldedEntityI interface {
	GetLogFields() F
}

// DEPRECATED
func obtainFields(entityName string, fieldedEntity FieldedEntityI) (result F) {
	result = make(F)

	defer func() {
		if r := recover(); r != nil {
			// `fieldedEntity` is nil
			result = make(F)
			result[entityName] = NilValueToLog
		}
	}()

	for fieldKey, fieldValue := range fieldedEntity.GetLogFields() {
		compositeKey := entityName + "_" + fieldKey
		result[compositeKey] = fieldValue
	}

	return result
}
