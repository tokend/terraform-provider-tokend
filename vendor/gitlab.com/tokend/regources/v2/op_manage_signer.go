package regources

//ManageSigner - details of corresponding op
// NOTE key type will be different for all actions
type ManageSigner struct {
	Key
	Attributes    *ManageSignerAttrs    `json:"attributes,omitempty"`
	Relationships *ManageSignerRelation `json:"relationships,omitempty"`
}

type ManageSignerRelation struct {
	Role   *Relation `json:"role"`
	Signer *Relation `json:"signer"`
}

// ManageSignerAttrs - details of new or updated rule
type ManageSignerAttrs struct {
	Weight   uint32  `json:"weight"`
	Identity uint32  `json:"identity"`
	Details  Details `json:"details"`
}
