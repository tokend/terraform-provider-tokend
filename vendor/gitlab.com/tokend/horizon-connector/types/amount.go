package types

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/amount"
)

// DEPRECATED: use regources.Amount (gitlab.com/tokend/regources)
type Amount int64

// DEPRECATED: use regources.Amount (gitlab.com/tokend/regources)
func (a Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(amount.String(int64(a)))
}

// DEPRECATED: use regources.Amount (gitlab.com/tokend/regources)
func (a *Amount) UnmarshalJSON(data []byte) error {
	var rawAmount string
	err := json.Unmarshal(data, &rawAmount)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal amount")
	}

	rawA, err := amount.Parse(rawAmount)
	*a = Amount(rawA)

	return errors.Wrap(err, "can't parse amount")
}

// DEPRECATED: use regources.Amount (gitlab.com/tokend/regources)
func (a Amount) String() string {
	return amount.String(int64(a))
}
