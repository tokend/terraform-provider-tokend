/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type SaleWhitelist struct {
	Key
	Relationships SaleWhitelistRelationships `json:"relationships"`
}
type SaleWhitelistResponse struct {
	Data     SaleWhitelist `json:"data"`
	Included Included      `json:"included"`
}

type SaleWhitelistListResponse struct {
	Data     []SaleWhitelist `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *SaleWhitelistListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *SaleWhitelistListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustSaleWhitelist - returns SaleWhitelist from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSaleWhitelist(key Key) *SaleWhitelist {
	var saleWhitelist SaleWhitelist
	if c.tryFindEntry(key, &saleWhitelist) {
		return &saleWhitelist
	}
	return nil
}
