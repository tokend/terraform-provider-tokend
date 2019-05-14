package listener

import (
	"context"
	"fmt"
	"gitlab.com/tokend/horizon-connector/internal/operation"
)

// WithdrawalRequestStreamingOpts is used to provide settings of ReviewableRequests streaming.
// Empty values are valid for strings and ints.
type WithdrawalRequestStreamingOpts struct {
	StopOnEmptyPage bool
	StartCursor     string
	ReverseOrder    bool
	// Filters
	DestAssetCode string
	RequestState  int32
}

// StreamWithdrawalRequests streams all ReviewableRequests of type Withdraw and TwoStepWithdraw
func (q *Q) StreamWithdrawalRequests(ctx context.Context, opts WithdrawalRequestStreamingOpts) <-chan ReviewableRequestEvent {
	var getParams string

	if opts.ReverseOrder {
		getParams += "&order=desc"
	}
	if opts.DestAssetCode != "" {
		getParams += fmt.Sprintf("&dest_asset_code=%s", opts.DestAssetCode)
	}
	if opts.RequestState > 0 {
		getParams += fmt.Sprintf("&state=%d", opts.RequestState)
	}

	if len(getParams) > 0 {
		// Remove leading ampersand
		getParams = getParams[1:]
	}

	return q.getAndStreamReviewableRequests(ctx, getParams, opts.StartCursor, operation.WithdrawalsReviewableRequestType, opts.StopOnEmptyPage)
}

// StreamWithdrawalRequestsOfAsset streams all Withdraw and TwoStepWithdraw ReviewableRequests
// with filter by provided destAssetCode
// DEPRECATED: Use StreamWithdrawalRequests instead, its' now able to do everything
func (q *Q) StreamWithdrawalRequestsOfAsset(ctx context.Context, destAssetCode string, reverseOrder, endlessly bool) <-chan ReviewableRequestEvent {
	getParams := fmt.Sprintf("dest_asset_code=%s", destAssetCode)

	if reverseOrder {
		getParams += "&order=desc"
	}

	return q.getAndStreamReviewableRequests(ctx, getParams, "", operation.WithdrawalsReviewableRequestType, !endlessly)
}
