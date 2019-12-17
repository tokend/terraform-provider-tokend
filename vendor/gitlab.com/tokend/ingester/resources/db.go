/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"database/sql/driver"

	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

//driverValue - converts interface into db supported type
func driverValue(data interface{}) (driver.Value, error) {
	data, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("failed to marshal details")
	}

	return data, nil
}

//driveScan - converts jsonb into type struct
func driveScan(src, dest interface{}) error {
	data, err := convertJSONB(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal jsonb")
	}

	return nil
}

func convertJSONB(src interface{}) ([]byte, error) {
	var data []byte
	switch rawData := src.(type) {
	case []byte:
		data = rawData
	case string:
		data = []byte(rawData)
	default:
		return nil, errors.New("Unexpected type for jsonb")
	}

	return data, nil
}
