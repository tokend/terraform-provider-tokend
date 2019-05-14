package regources

type SignerRoleRelation struct {
	Rules *RelationCollection `json:"rules"`
	Owner *Relation           `json:"owner"`
}

type SignerRole struct {
	Key
	Attributes    SignerRoleAttrs    `json:"attributes"`
	Relationships SignerRoleRelation `json:"relationships"`
}

type SignerRoleAttrs struct {
	Details Details `json:"details"`
}

type SignerRoleResponse struct {
	Data     SignerRole `json:"data"`
	Included Included   `json:"included"`
}

type SignerRolesResponse struct {
	Links    *Links       `json:"links"`
	Data     []SignerRole `json:"data"`
	Included Included     `json:"included"`
}
