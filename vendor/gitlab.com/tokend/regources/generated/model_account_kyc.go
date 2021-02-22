/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type AccountKyc struct {
	Key
	Attributes AccountKycAttributes `json:"attributes"`
}
type AccountKycResponse struct {
	Data     AccountKyc `json:"data"`
	Included Included   `json:"included"`
}

type AccountKycListResponse struct {
	Data     []AccountKyc    `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *AccountKycListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *AccountKycListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustAccountKyc - returns AccountKyc from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAccountKyc(key Key) *AccountKyc {
	var accountKYC AccountKyc
	if c.tryFindEntry(key, &accountKYC) {
		return &accountKYC
	}
	return nil
}
