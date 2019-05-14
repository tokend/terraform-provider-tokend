/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type AccountRuleAttributes struct {
	// defines an action to be performed over the specified resource * 1:  \"any\" * 2:  \"create\" * 3:  \"create_for_other\" * 4:  \"create_with_tasks\" * 5:  \"manage\" * 6:  \"send\" * 7:  \"withdraw\" * 8:  \"receive_issuance\" * 9:  \"receive_payment\" * 10: \"receive_atomic_swap\" * 11: \"participate\" * 12: \"bind\" * 13: \"update_max_issuance\" * 14: \"check\" * 15: \"cancel\"
	Action  xdr.AccountRuleAction `json:"action"`
	Details Details               `json:"details"`
	// defines whether or not the specified action is forbidden
	Forbids bool `json:"forbids"`
	// defines resource to which the rule is applied. TODO: add link to XDR
	Resource xdr.AccountRuleResource `json:"resource"`
}
