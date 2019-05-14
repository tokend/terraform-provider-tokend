package regources

import (
	"time"

	"gitlab.com/tokend/go/xdr"
)

//ManageCreatePollRequestOp - stores details of manage create poll request poll operation
type ManageCreatePollRequestOp struct {
	Key
	Attributes    ManageCreatePollRequestOpAttrs     `json:"attributes"`
	Relationships ManageCreatePollRequestOpRelations `json:"relationships"`
}

//ManageCreatePollRequestOpAttrs - details of ManageCreatePollRequestOp
type ManageCreatePollRequestOpAttrs struct {
	Action xdr.ManageCreatePollRequestAction `json:"action"`
	Create *CreatePollRequestOp              `json:"create,omitempty"`
}

type CreatePollRequestOp struct {
	PermissionType           uint32    `json:"permission_type"`
	NumberOfChoices          uint32    `json:"number_of_choices"`
	CreatorDetails           Details   `json:"creator_details"`
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
	ResultProviderID         string    `json:"result_provider_id"`
	VoteConfirmationRequired bool      `json:"vote_confirmation_required"`
	PollData                 PollData  `json:"poll_data"`
	AllTasks                 *uint32   `json:"all_tasks,omitempty"`
}

//ManageCreatePollRequestOpRelations - relationships of ManageCreatePollRequestOp
type ManageCreatePollRequestOpRelations struct {
	Request        *Relation `json:"request"`
	ResultProvider *Relation `json:"result_provider"`
}
