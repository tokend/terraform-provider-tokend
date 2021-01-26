/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type SwapState int

const (
	SwapStateOpen SwapState = iota + 1
	SwapStateClosed
	SwapStateCanceled
)

var swapStateMap = map[SwapState]string{
	SwapStateOpen:     "open",
	SwapStateClosed:   "closed",
	SwapStateCanceled: "canceled",
}

func (s SwapState) MarshalJSON() ([]byte, error) {
	return json.Marshal(Flag{
		Name:  swapStateMap[s],
		Value: int32(s),
	})
}

//String - converts int enum to string
func (s SwapState) String() string {
	return swapStateMap[s]
}

func (s *SwapState) UnmarshalJSON(b []byte) error {
	var res Flag
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}

	*s = SwapState(res.Value)
	return nil
}

func (s SwapState) IsFlag() bool {
	return true
}
