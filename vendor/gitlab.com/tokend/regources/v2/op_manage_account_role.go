package regources

//ManageAccountRole - details of corresponding op
// NOTE key type will be different for all actions
type ManageAccountRole struct {
	Key
	Attributes    *ManageAccountRoleAttrs   `json:"attributes,omitempty"`
	Relationships ManageAccountRoleRelation `json:"relationships"`
}

//ManageAccountRoleRelation - defines op relations
type ManageAccountRoleRelation struct {
	Role  *Relation           `json:"role,omitempty"`
	Rules *RelationCollection `json:"rules,omitempty"`
}

// ManageAccountRoleAttrs - details of new or updated rule
type ManageAccountRoleAttrs struct {
	Details Details `json:"details"`
}
