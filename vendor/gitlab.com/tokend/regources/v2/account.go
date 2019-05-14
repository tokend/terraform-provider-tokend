package regources

// AccountResponse - response on /account request
type AccountResponse struct {
	Data     Account  `json:"data"`
	Included Included `json:"included"`
}

// Account - Resource object representing AccountEntry
type Account struct {
	Key
	Relationships AccountRelationships `json:"relationships"`
}

type AccountRelationships struct {
	Role              *Relation           `json:"role,omitempty"`
	Balances          *RelationCollection `json:"balances,omitempty"`
	Fees              *RelationCollection `json:"fees,omitempty"`
	Referrer          *Relation           `json:"referrer,omitempty"`
	Limits            *RelationCollection `json:"limits,omitempty"`
	ExternalSystemIDs *RelationCollection `json:"external_system_ids,omitempty"`
}
