package regources

import (
	"encoding/json"
	"gitlab.com/tokend/go/xdr"
	"time"
)

type ReviewableRequestResponse struct {
	Data     ReviewableRequest `json:"data"`
	Included Included          `json:"included"`
}

type ReviewableRequestsResponse struct {
	Links    *Links              `json:"links"`
	Data     []ReviewableRequest `json:"data"`
	Included Included            `json:"included"`
}

func (r ReviewableRequestsResponse) MarshalJSON() ([]byte, error) {
	if r.Data == nil {
		r.Data = []ReviewableRequest{}
	}

	type temp ReviewableRequestsResponse
	return json.Marshal(temp(r))
}

type ReviewableRequest struct {
	Key
	Attributes    ReviewableRequestAttrs     `json:"attributes"`
	Relationships ReviewableRequestRelations `json:"relationships"`
}

type ReviewableRequestAttrs struct {
	Reference       *string                   `json:"reference,omitempty"`
	RejectReason    string                    `json:"reject_reason"`
	Hash            string                    `json:"hash"`
	AllTasks        uint32                    `json:"all_tasks"`
	PendingTasks    uint32                    `json:"pending_tasks"`
	ExternalDetails map[string]interface{}    `json:"external_details,omitempty"`
	CreatedAt       time.Time                 `json:"created_at"`
	UpdatedAt       time.Time                 `json:"updated_at"`
	State           string                    `json:"state"`
	StateI          int32                     `json:"state_i"`
	XDRType         xdr.ReviewableRequestType `json:"xdr_type"`
}

type ReviewableRequestRelations struct {
	Requestor      *Relation `json:"requestor"`
	Reviewer       *Relation `json:"reviewer"`
	RequestDetails *Relation `json:"request_details"`
}
