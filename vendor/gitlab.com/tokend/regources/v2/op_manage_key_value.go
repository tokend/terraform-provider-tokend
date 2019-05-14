package regources

import "gitlab.com/tokend/go/xdr"

//ManageKeyValueOp - stores details of create account operation
type ManageKeyValueOp struct {
	Key
	Attributes ManageKeyValueOpAttrs `json:"attributes"`
}

//ManageKeyValueAttrsOp - details of ManageKeyValueOp
type ManageKeyValueOpAttrs struct {
	Key    string              `json:"key"`
	Action xdr.ManageKvAction  `json:"action"`
	Value  *KeyValueEntryValue `json:"value,omitempty"`
}
