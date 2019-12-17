/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AccountRelationships struct {
	Balances *RelationCollection `json:"balances,omitempty"`
	KycData  *Relation           `json:"kyc_data,omitempty"`
	Referrer *Relation           `json:"referrer,omitempty"`
	Roles    *RelationCollection `json:"roles,omitempty"`
}
