package regources

type AccountRole struct {
	Key
	Attributes    AccountRoleAttrs `json:"attributes"`
	Relationships RoleRelation     `json:"relationships"`
}

type AccountRoleAttrs struct {
	Details Details `json:"details"`
}

type RoleRelation struct {
	Rules RelationCollection `json:"rules"`
}

type AccountRoleResponse struct {
	Data     AccountRole `json:"data"`
	Included Included    `json:"included"`
}

type AccountRolesResponse struct {
	Links    *Links        `json:"links"`
	Data     []AccountRole `json:"data"`
	Included Included      `json:"included"`
}
