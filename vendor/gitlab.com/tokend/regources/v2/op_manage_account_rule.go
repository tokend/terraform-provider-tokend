package regources

import "gitlab.com/tokend/go/xdr"

//ManageAccountRule - details of corresponding op
// NOTE key type will be different for all actions
type ManageAccountRule struct {
	Key
	Attributes    *ManageAccountRuleAttrs    `json:"attributes,omitempty"`
	Relationships *ManageAccountRuleRelation `json:"relationships,omitempty"`
}

type ManageAccountRuleRelation struct {
	Rule *Relation `json:"rule"`
}

// ManageAccountRuleAttrs - details of new or updated rule
type ManageAccountRuleAttrs struct {
	Resource xdr.AccountRuleResource `json:"resource"`
	Action   xdr.AccountRuleAction   `json:"action"`
	Forbids  bool                    `json:"forbids"`
	Details  Details                 `json:"details"`
}
