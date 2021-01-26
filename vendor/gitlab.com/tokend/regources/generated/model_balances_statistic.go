/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type BalancesStatistic struct {
	Key
	Attributes BalancesStatisticAttributes `json:"attributes"`
}
type BalancesStatisticResponse struct {
	Data     BalancesStatistic `json:"data"`
	Included Included          `json:"included"`
}

type BalancesStatisticListResponse struct {
	Data     []BalancesStatistic `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *BalancesStatisticListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *BalancesStatisticListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustBalancesStatistic - returns BalancesStatistic from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBalancesStatistic(key Key) *BalancesStatistic {
	var balancesStatistic BalancesStatistic
	if c.tryFindEntry(key, &balancesStatistic) {
		return &balancesStatistic
	}
	return nil
}
