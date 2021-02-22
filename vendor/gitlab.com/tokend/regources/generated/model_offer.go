/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Offer struct {
	Key
	Attributes    OfferAttributes    `json:"attributes"`
	Relationships OfferRelationships `json:"relationships"`
}
type OfferResponse struct {
	Data     Offer    `json:"data"`
	Included Included `json:"included"`
}

type OfferListResponse struct {
	Data     []Offer         `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *OfferListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *OfferListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustOffer - returns Offer from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustOffer(key Key) *Offer {
	var offer Offer
	if c.tryFindEntry(key, &offer) {
		return &offer
	}
	return nil
}
