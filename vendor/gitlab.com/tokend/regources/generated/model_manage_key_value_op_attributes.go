/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "gitlab.com/tokend/go/xdr"

type ManageKeyValueOpAttributes struct {
	// * 1: \"put\", * 2: \"remove\"
	Action xdr.ManageKvAction `json:"action"`
	// Key of key-value entry to manage
	Key   string              `json:"key"`
	Value *KeyValueEntryValue `json:"value,omitempty"`
}
