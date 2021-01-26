/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type QuoteAsset struct {
	Key
	Attributes QuoteAssetAttributes `json:"attributes"`
}
type QuoteAssetResponse struct {
	Data     QuoteAsset `json:"data"`
	Included Included   `json:"included"`
}

type QuoteAssetListResponse struct {
	Data     []QuoteAsset    `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *QuoteAssetListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *QuoteAssetListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustQuoteAsset - returns QuoteAsset from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustQuoteAsset(key Key) *QuoteAsset {
	var quoteAsset QuoteAsset
	if c.tryFindEntry(key, &quoteAsset) {
		return &quoteAsset
	}
	return nil
}
