package regources

import (
	"gitlab.com/tokend/go/xdr"
)

//ManageVoteOp - stores details of manage vote operation
type ManageVoteOp struct {
	Key
	Attributes    ManageVoteOpAttributes `json:"attributes"`
	Relationships ManageVoteOpRelations  `json:"relationships"`
}

//ManageCreateVoteRequestOpAttributes - details of ManageCreateVoteRequestOp
type ManageVoteOpAttributes struct {
	Action xdr.ManageVoteAction `json:"action"`
	Create *CreateVoteOp        `json:"create,omitempty"`
	Remove *RemoveVoteOp        `json:"remove,omitempty"`
}

type CreateVoteOp struct {
	PollID   int64    `json:"poll_id"`
	VoteData VoteData `json:"vote_data"`
}

type RemoveVoteOp struct {
	PollID int64 `json:"poll_id"`
}

//ManageVoteOpRelations- relationships of ManageVoteOp
type ManageVoteOpRelations struct {
	Poll           *Relation `json:"poll"`
	Voter          *Relation `json:"voter"`
	ResultProvider *Relation `json:"result_provider"`
}
