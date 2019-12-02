/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type UpdateRuleOpAttributes struct {
	// Union switch object with xdr.RuleActionType discriminant type
	Action  xdr.RuleAction `json:"action"`
	Details Details        `json:"details"`
	// false means rule allows to perform defined action with defined resource
	Forbids bool `json:"forbids"`
	// Union switch object with xdr.LedgerEntryType discriminant type
	Resource xdr.RuleResource `json:"resource"`
}
