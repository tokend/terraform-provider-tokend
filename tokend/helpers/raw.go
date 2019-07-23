package helpers

import (
	"encoding/json"
	"math"

	"github.com/mitchellh/mapstructure"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func WildCardUintFromRaw(raw string) (uint64, error) {
	var result uint64
	var err error
	if raw == "*" {
		result = math.MaxUint64
	} else {
		result, err = cast.ToUint64E(raw)
		if err != nil {
			return 0, errors.Wrap(err, "failed to cast to uint")
		}
	}
	return result, nil
}

type MapDetails map[string]interface{}

func (d MapDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}(d))
}

func DetailsFromRaw(raw interface{}) (MapDetails, error) {
	details := MapDetails{}
	if raw != nil {
		rawDetails, err := cast.ToStringMapE(raw)
		if err != nil {
			return nil, errors.Wrap(err, "failed to cast to map")
		}
		details = MapDetails(rawDetails)
	}
	return details, nil
}

//TODO move to regources
type StellarData struct {
	Address string
}

func StellarDataFromRaw(raw interface{}) (*StellarData, error) {
	rawData, err := cast.ToStringMapE(raw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast to map")
	}
	var data StellarData
	err = mapstructure.Decode(rawData, &data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode data")
	}
	return &data, nil
}
