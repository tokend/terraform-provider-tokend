/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type AssetState int

const (
	AssetStateActive AssetState = iota
	AssetStateDeleted
)

var assetStateStr = map[AssetState]string{
	AssetStateActive:  "active",
	AssetStateDeleted: "deleted",
}

func (s AssetState) String() string {
	return assetStateStr[s]
}

func (s AssetState) MarshalJSON() ([]byte, error) {
	return json.Marshal(Flag{
		Name:  assetStateStr[s],
		Value: int32(s),
	})
}

func (s *AssetState) UnmarshalJSON(b []byte) error {
	var res Flag
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}

	*s = AssetState(res.Value)
	return nil
}

func (s AssetState) IsFlag() bool {
	return true
}
