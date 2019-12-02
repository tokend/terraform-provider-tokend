/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type RuleAttributes struct {
	// defines an action to be performed over the specified resource
	Action  xdr.RuleAction `json:"action"`
	Details Details        `json:"details"`
	// defines whether or not the specified action is forbidden
	Forbids bool `json:"forbids"`
	// defines resource to which the rule is applied. TODO: add link to XDR
	Resource xdr.RuleResource `json:"resource"`
}
