/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type LicenseOp struct {
	Key
	Attributes LicenseOpAttributes `json:"attributes"`
}
type LicenseOpResponse struct {
	Data     LicenseOp `json:"data"`
	Included Included  `json:"included"`
}

type LicenseOpListResponse struct {
	Data     []LicenseOp     `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
	Meta     json.RawMessage `json:"meta,omitempty"`
}

func (r *LicenseOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *LicenseOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustLicenseOp - returns LicenseOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustLicenseOp(key Key) *LicenseOp {
	var licenseOp LicenseOp
	if c.tryFindEntry(key, &licenseOp) {
		return &licenseOp
	}
	return nil
}
