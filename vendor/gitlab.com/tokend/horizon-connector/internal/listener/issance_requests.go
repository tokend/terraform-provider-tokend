package listener

import (
	"context"
	"fmt"
	"gitlab.com/tokend/horizon-connector/internal/operation"
)

// IssuanceRequestStreamingOpts is used to provide settings of ReviewableRequests streaming.
// Empty values are valid for strings and ints.
type IssuanceRequestStreamingOpts struct {
	StopOnEmptyPage bool
	StartCursor     string
	ReverseOrder    bool
	// Filters
	AssetCode    string
	RequestState int32
}

// StreamIssuanceRequests streams all ReviewableRequests of type CreateIssuance
func (q *Q) StreamIssuanceRequests(ctx context.Context, opts IssuanceRequestStreamingOpts) <-chan ReviewableRequestEvent {
	var getParams string

	if opts.ReverseOrder {
		getParams += "&order=desc"
	}
	if opts.AssetCode != "" {
		getParams += fmt.Sprintf("&asset=%s", opts.AssetCode)
	}
	if opts.RequestState > 0 {
		getParams += fmt.Sprintf("&state=%d", opts.RequestState)
	}

	if len(getParams) > 0 {
		// Remove leading ampersand
		getParams = getParams[1:]
	}

	return q.getAndStreamReviewableRequests(ctx, getParams, opts.StartCursor, operation.IssuanceReviewableRequestType, opts.StopOnEmptyPage)
}
