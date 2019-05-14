/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManageBalanceOpAttributes struct {
	// * 0: \"create\" * 1: \"delete_balance\" * 2: \"create_unique\"
	Action xdr.ManageBalanceAction `json:"action"`
	// Address of balance to manage
	BalanceAddress string `json:"balance_address"`
}
