package regources

//CreateChangeRoleRequest - details of corresponding op
type CreateChangeRoleRequest struct {
	Key
	Attributes    CreateChangeRoleRequestAttrs       `json:"attributes"`
	Relationships CreateChangeRoleRequestOpRelations `json:"relationships"`
}

//CreateChangeRoleRequestAttrs - details of corresponding op
type CreateChangeRoleRequestAttrs struct {
	CreatorDetails Details `json:"creator_details"`
	AllTasks       *uint32 `json:"all_tasks"`
}

// CreateChangeRoleRequestOpRelations - relationships of the operation
type CreateChangeRoleRequestOpRelations struct {
	AccountToUpdateRole *Relation `json:"account_to_update_role"`
	Request             *Relation `json:"request"`
	RoleToSet           *Relation `json:"role_to_set"`
}
