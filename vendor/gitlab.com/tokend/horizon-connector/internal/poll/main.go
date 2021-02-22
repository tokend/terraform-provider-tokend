package poll

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/horizon-connector/types"
	regources "gitlab.com/tokend/regources/generated"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

type Filters struct {
	Owner            *string
	PermissionType   *uint32
	VoteConfirmation *bool
	MinStartTime     *time.Time
	MinEndTime       *time.Time
	MaxStartTime     *time.Time
	MaxEndTime       *time.Time
	State            *uint32
}

func (f *Filters) Encode() string {
	if f == nil {
		return ""
	}

	u := url.Values{}

	if f.State != nil {
		u.Add("filter[state]", cast.ToString(*f.State))
	}
	if f.Owner != nil {
		u.Add("filter[owner]", *f.Owner)
	}
	if f.PermissionType != nil {
		u.Add("filter[permission_type]", cast.ToString(*f.PermissionType))
	}
	if f.MinStartTime != nil {
		u.Add("filter[min_start_time]", f.MinStartTime.Format(time.RFC3339))
	}
	if f.MaxStartTime != nil {
		u.Add("filter[max_start_time]", f.MaxStartTime.Format(time.RFC3339))
	}
	if f.MinEndTime != nil {
		u.Add("filter[min_end_time]", f.MinEndTime.Format(time.RFC3339))
	}
	if f.MaxEndTime != nil {
		u.Add("filter[max_end_time]", f.MaxEndTime.Format(time.RFC3339))
	}
	if f.State != nil {
		u.Add("filter[pending_tasks_not_set]", cast.ToString(*f.State))
	}

	return u.Encode()
}

func (q *Q) CorePolls(urlParams ...types.URLParamer) ([]regources.Poll, error) {
	params := make([]string, 0, len(urlParams))
	for _, p := range urlParams {
		params = append(params, p.Encode())
	}
	endpoint := fmt.Sprintf("/v3/polls?%s", strings.Join(params, "&"))

	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var polls regources.PollListResponse
	if err := json.Unmarshal(response, &polls); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response", logan.F{
			"raw_response": string(response),
		})
	}
	return polls.Data, nil
}

type Include struct {
	Participation      bool
	ParticipationVotes bool
}

func (i *Include) Encode() string {
	if i == nil {
		return ""
	}

	u := url.Values{}

	if i.Participation {
		u.Add("include", "participation")
	}
	if i.ParticipationVotes {
		u.Add("include", "participation.votes")
	}

	return u.Encode()
}

func (q *Q) ByID(id uint64, urlParams ...types.URLParamer) (*regources.PollResponse, error) {
	params := make([]string, 0, len(urlParams))
	for _, p := range urlParams {
		params = append(params, p.Encode())
	}
	endpoint := fmt.Sprintf("/v3/polls/%d?%s", id, strings.Join(params, ","))

	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var poll regources.PollResponse
	if err := json.Unmarshal(response, &poll); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response", logan.F{
			"raw": string(response),
		})
	}

	return &poll, nil
}
