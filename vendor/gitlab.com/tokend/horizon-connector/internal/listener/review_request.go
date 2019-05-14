package listener

import (
	"gitlab.com/tokend/horizon-connector/internal/resources/operations"
	"context"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type ReviewRequestOpResponse struct {
	body *operations.ReviewRequest
	err  error
}

func (r *ReviewRequestOpResponse) Unwrap() (*operations.ReviewRequest, error) {
	return r.body, r.err
}

type reviewRequestOpStreamer struct {
	q       *Q
	results chan ReviewRequestOpResponse
}

func streamReviewRequestOp(q *Q, ctx context.Context, buffer int) <-chan ReviewRequestOpResponse {
	streamer := reviewRequestOpStreamer{
		q:       q,
		results: make(chan ReviewRequestOpResponse, buffer),
	}

	go streamer.startStream(ctx)
	return streamer.results
}

func (s *reviewRequestOpStreamer) startStream(ctx context.Context) {
	defer close(s.results)
	cursor := ""
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		cursor = s.stream(ctx, cursor)
	}
}

func (s *reviewRequestOpStreamer) stream(ctx context.Context, cursor string) (newCursor string) {
	newCursor = cursor

	defer func() {
		if recovered := recover(); recovered != nil {
			err := errors.FromPanic(recovered)
			s.send(ctx, nil, err)
		}
	}()

	reviewRequestOperations, err := s.q.opQ.ReviewRequestOperations(newCursor)
	if err != nil {
		s.send(ctx, nil, errors.Wrap(err, "failed to get review request operations"))
		return
	}

	for i := range reviewRequestOperations {
		if !s.send(ctx, &reviewRequestOperations[i], nil) {
			return
		}
		newCursor = reviewRequestOperations[i].PT
	}

	return
}

func (s *reviewRequestOpStreamer) send(ctx context.Context, body *operations.ReviewRequest, err error) bool {
	response := ReviewRequestOpResponse{
		body: body,
		err:  err,
	}

	select {
	case <-ctx.Done():
		return false
	case s.results <- response:
		return true
	}
}
