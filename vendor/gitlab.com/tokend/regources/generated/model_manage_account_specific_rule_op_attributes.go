/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManageAccountSpecificRuleOpAttributes struct {
	// * 0: create * 1: remove
	Action     xdr.ManageAccountSpecificRuleAction `json:"action"`
	CreateData *CreateAccountSpecificRuleData      `json:"create_data,omitempty"`
}
