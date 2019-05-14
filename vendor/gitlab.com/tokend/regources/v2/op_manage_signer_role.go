package regources

//ManageSignerRole - details of corresponding op
// NOTE key type will be different for all actions
type ManageSignerRole struct {
	Key
	Attributes    *ManageSignerRoleAttrs   `json:"attributes,omitempty"`
	Relationships ManageSignerRoleRelation `json:"relationships"`
}

//ManageSignerRoleRelation - defines op relations
type ManageSignerRoleRelation struct {
	Role  *Relation           `json:"role,omitempty"`
	Rules *RelationCollection `json:"rules,omitempty"`
}

// ManageSignerRoleAttrs - details of new or updated rule
type ManageSignerRoleAttrs struct {
	Details    Details `json:"details"`
	IsReadOnly bool    `json:"is_read_only"`
}
