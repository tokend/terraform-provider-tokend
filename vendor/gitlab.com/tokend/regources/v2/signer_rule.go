package regources

import "gitlab.com/tokend/go/xdr"

type SignerRule struct {
	Key
	Attributes    SignerRuleAttr     `json:"attributes"`
	Relationships SignerRuleRelation `json:"relationships"`
}

type SignerRuleRelation struct {
	Owner *Relation `json:"owner"`
}

type SignerRuleAttr struct {
	Resource  xdr.SignerRuleResource `json:"resource"`
	Action    xdr.SignerRuleAction   `json:"action"`
	Forbids   bool                   `json:"forbids"`
	IsDefault bool                   `json:"is_default"`
	Details   Details                `json:"details"`
}

type SignerRuleResponse struct {
	Data SignerRule `json:"data"`
}

type SignerRulesResponse struct {
	Links *Links       `json:"links"`
	Data  []SignerRule `json:"data"`
}
