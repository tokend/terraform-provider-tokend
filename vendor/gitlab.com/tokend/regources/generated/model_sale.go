/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Sale struct {
	Key
	Attributes    SaleAttributes    `json:"attributes"`
	Relationships SaleRelationships `json:"relationships"`
}
type SaleResponse struct {
	Data     Sale     `json:"data"`
	Included Included `json:"included"`
}

type SaleListResponse struct {
	Data     []Sale          `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *SaleListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *SaleListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustSale - returns Sale from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSale(key Key) *Sale {
	var sale Sale
	if c.tryFindEntry(key, &sale) {
		return &sale
	}
	return nil
}
