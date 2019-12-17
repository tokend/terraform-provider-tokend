/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/amount"
)

type Amount uint64

func (a Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(amount.StringU(uint64(a)))
}

func (a *Amount) UnmarshalJSON(data []byte) error {
	var rawAmount string
	err := json.Unmarshal(data, &rawAmount)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal amount")
	}

	rawA, err := amount.ParseU(rawAmount)
	*a = Amount(rawA)

	return errors.Wrap(err, "can't parse amount")
}

func (a Amount) String() string {
	return amount.StringU(uint64(a))
}
