/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type Limits struct {
	Key
	Attributes    LimitsAttributes    `json:"attributes"`
	Relationships LimitsRelationships `json:"relationships"`
}
type LimitsResponse struct {
	Data     Limits   `json:"data"`
	Included Included `json:"included"`
}

type LimitsListResponse struct {
	Data     []Limits        `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *LimitsListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *LimitsListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustLimits - returns Limits from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustLimits(key Key) *Limits {
	var limits Limits
	if c.tryFindEntry(key, &limits) {
		return &limits
	}
	return nil
}
