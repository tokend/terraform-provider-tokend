/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type SignerRuleAttributes struct {
	// defines an action to be performed over the specified resource * 1:  \"any\" * 2:  \"create\" * 3:  \"create_for_other\" * 4:  \"update\" * 5:  \"manage\" * 6:  \"send\" * 7:  \"remove\" * 8:  \"cancel\" * 9:  \"review\" * 10: \"receive_atomic_swap\" * 11: \"participate\" * 12: \"bind\" * 13: \"update_max_issuance\" * 14: \"check\"
	Action  xdr.SignerRuleAction `json:"action"`
	Details Details              `json:"details"`
	// defines whether the specified action is forbidden
	Forbids bool `json:"forbids"`
	// defines whether this rule should be included into all new roles
	IsDefault bool `json:"is_default"`
	// defines a resource to which the rule is applied. TODO: add link to XDR
	Resource xdr.SignerRuleResource `json:"resource"`
}
