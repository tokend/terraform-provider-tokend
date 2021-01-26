/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type ManageAssetPairOp struct {
	Key
	Attributes    ManageAssetPairOpAttributes    `json:"attributes"`
	Relationships ManageAssetPairOpRelationships `json:"relationships"`
}
type ManageAssetPairOpResponse struct {
	Data     ManageAssetPairOp `json:"data"`
	Included Included          `json:"included"`
}

type ManageAssetPairOpListResponse struct {
	Data     []ManageAssetPairOp `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
	Meta     json.RawMessage     `json:"meta,omitempty"`
}

func (r *ManageAssetPairOpListResponse) PutMeta(v interface{}) (err error) {
	r.Meta, err = json.Marshal(v)
	return err
}

func (r *ManageAssetPairOpListResponse) GetMeta(out interface{}) error {
	return json.Unmarshal(r.Meta, out)
}

// MustManageAssetPairOp - returns ManageAssetPairOp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustManageAssetPairOp(key Key) *ManageAssetPairOp {
	var manageAssetPairOp ManageAssetPairOp
	if c.tryFindEntry(key, &manageAssetPairOp) {
		return &manageAssetPairOp
	}
	return nil
}
