/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type SaleState int

const (
	SaleStateOpen SaleState = iota + 1
	SaleStateClosed
	SaleStateCanceled
)

var saleStateMap = map[SaleState]string{
	SaleStateOpen:     "open",
	SaleStateClosed:   "closed",
	SaleStateCanceled: "canceled",
}

func (s SaleState) MarshalJSON() ([]byte, error) {
	return json.Marshal(Flag{
		Name:  saleStateMap[s],
		Value: int32(s),
	})
}

//String - converts int enum to string
func (s SaleState) String() string {
	return saleStateMap[s]
}

func (s *SaleState) UnmarshalJSON(b []byte) error {
	var res Flag
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}

	*s = SaleState(res.Value)
	return nil
}

func (s SaleState) IsFlag() bool {
	return true
}
