/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type AccountRelationships struct {
	Balances          *RelationCollection `json:"balances,omitempty"`
	ExternalSystemIds *RelationCollection `json:"external_system_ids,omitempty"`
	Fees              *RelationCollection `json:"fees,omitempty"`
	Limits            *RelationCollection `json:"limits,omitempty"`
	Referrer          *Relation           `json:"referrer,omitempty"`
	Role              *Relation           `json:"role,omitempty"`
}
