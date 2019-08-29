package helpers

import (
	"encoding/json"
	"math"

	"github.com/mitchellh/mapstructure"
	"gitlab.com/tokend/go/strkey"
	"gitlab.com/tokend/go/xdr"

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

func PolicyFromRaw(raw []interface{}) (uint32, error) {
	var policiesCode uint32
	for _, policiesRaw := range raw {

		policyRaw := cast.ToString(policiesRaw)

		ok := false

		for _, guess := range xdr.AssetPolicyAll {
			if guess.ShortString() == policyRaw {
				ok = true
				policiesCode |= uint32(guess)
			}
		}

		if !ok {
			return 0, errors.New("invalid policy name")
		}
	}
	return policiesCode, nil
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

	if err != nil {
		return nil, errors.Wrap(err, "not valid account ID")
	}

	errAdr := accountID.SetAddress(rawstr)
	if errAdr != nil {
		return nil, errors.Wrap(errCast, "failed to set address")
	}

	return &accountID, nil
}
func StatsOpTypeFromRaw(raw interface{}) (int32, error) {
	var typesCode int32
	typeRaw := cast.ToString(raw)

	ok := false

	for index, guess := range xdr.StatsOpTypeAll {
		if guess.ShortString() == typeRaw {
			ok = true
			typesCode |= int32(index)
		}
	}
	if !ok {
		return 0, errors.New("invalid type code: %s")
	}

	return typesCode, nil
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
