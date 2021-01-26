/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type SaleParticipation struct {
	Key
	Attributes    SaleParticipationAttributes    `json:"attributes"`
	Relationships SaleParticipationRelationships `json:"relationships"`
}
type SaleParticipationResponse struct {
	Data     SaleParticipation `json:"data"`
	Included Included          `json:"included"`
}

type SaleParticipationListResponse struct {
	Data     []SaleParticipation `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *SaleParticipationListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *SaleParticipationListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustSaleParticipation - returns SaleParticipation from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSaleParticipation(key Key) *SaleParticipation {
	var saleParticipation SaleParticipation
	if c.tryFindEntry(key, &saleParticipation) {
		return &saleParticipation
	}
	return nil
}
