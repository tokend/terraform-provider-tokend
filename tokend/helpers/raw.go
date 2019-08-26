package helpers

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"gitlab.com/tokend/go/strkey"
	"gitlab.com/tokend/go/xdr"
	"math"

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

type EthereumData struct {
	XPub string `mapstructure:"x_pub"`
}

func EthereumDataFromRaw(raw interface{}) (*EthereumData, error) {
	rawData, err := cast.ToStringMapE(raw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast to map")
	}
	var data EthereumData
	err = mapstructure.Decode(rawData, &data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode data")
	}
	return &data, nil
}
func AccountIDFromRaw(raw interface{}) (*xdr.AccountId, error) {
	rawstr, errCast := cast.ToStringE(raw)
	if errCast != nil {
		return nil, errors.Wrap(errCast, "failed to cast to string")
	}
	var accountID xdr.AccountId
	err := ValidateAccountID(rawstr)
	if err == nil {
		errAdr := accountID.SetAddress(rawstr)
		if errAdr != nil {
			return nil, errors.Wrap(errCast, "failed to set address")
		}
	}
	return &accountID, nil
}

func ValidateAccountID(a string) error {
	_, err := strkey.Decode(strkey.VersionByteAccountID, string(a))
	if err != nil {
		return err
	}
	return nil
}

func ValidateLimits(dailyOut uint64, weeklyOut uint64, monthlyOut uint64, annualOut uint64) bool {
	return dailyOut <= weeklyOut && weeklyOut <= monthlyOut && monthlyOut <= annualOut
}
