/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type SaleAccessDefinitionType int

const (
	SaleAccessDefinitionTypeNone SaleAccessDefinitionType = iota + 1
	SaleAccessDefinitionTypeWhitelist
	SaleAccessDefinitionTypeBlacklist
)

var saleAccessDefinitionTypeMap = map[SaleAccessDefinitionType]string{
	SaleAccessDefinitionTypeNone:      "none",
	SaleAccessDefinitionTypeWhitelist: "whitelist",
	SaleAccessDefinitionTypeBlacklist: "blacklist",
}

func (s SaleAccessDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(Flag{
		Name:  saleAccessDefinitionTypeMap[s],
		Value: int32(s),
	})
}

func (s *SaleAccessDefinitionType) UnmarshalJSON(data []byte) error {
	var flag Flag
	if err := json.Unmarshal(data, &flag); err != nil {
		return errors.Wrap(err, "failed to unmarshal sale definition type to byte")
	}

	*s = SaleAccessDefinitionType(flag.Value)

	return nil
}

// String - converts int enum to string
func (s SaleAccessDefinitionType) String() string {
	return saleAccessDefinitionTypeMap[s]
}
