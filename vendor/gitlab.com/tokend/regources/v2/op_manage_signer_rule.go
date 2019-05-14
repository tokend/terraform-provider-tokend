package regources

import "gitlab.com/tokend/go/xdr"

//ManageSignerRule - details of corresponding op
// NOTE key type will be different for all actions
type ManageSignerRule struct {
	Key
	Attributes    *ManageSignerRuleAttrs    `json:"attributes,omitempty"`
	Relationships *ManageSignerRuleRelation `json:"relationships,omitempty"`
}

type ManageSignerRuleRelation struct {
	Rule *Relation `json:"rule"`
}

// ManageSignerRuleAttrs - details of new or updated rule
type ManageSignerRuleAttrs struct {
	Resource   xdr.SignerRuleResource `json:"resource"`
	Action     xdr.SignerRuleAction   `json:"action"`
	Forbids    bool                   `json:"forbids"`
	IsDefault  bool                   `json:"is_default"`
	IsReadOnly bool                   `json:"is_read_only"`
	Details    Details                `json:"details"`
}
