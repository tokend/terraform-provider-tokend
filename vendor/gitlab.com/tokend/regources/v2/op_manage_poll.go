package regources

import (
	"gitlab.com/tokend/go/xdr"
)

//ManagePollOp - stores details of manage poll operation
type ManagePollOp struct {
	Key
	Attributes    ManagePollOpAttributes `json:"attributes"`
	Relationships ManagePollOpRelations  `json:"relationships"`
}

//ManagePollOpAttributes - details of ManagePollOp
type ManagePollOpAttributes struct {
	Action xdr.ManagePollAction `json:"action"`
	Close  *ClosePollOp         `json:"close,omitempty"`
}

type ClosePollOp struct {
	PollID     int64          `json:"poll_id"`
	PollResult xdr.PollResult `json:"poll_result"`
	Details    Details        `json:"details"`
}

//ManagePollOpAttributes - relationships of ManageBalanceOp
type ManagePollOpRelations struct {
	Poll *Relation `json:"poll"`
}
