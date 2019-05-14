package listener

import (
	"context"
	"gitlab.com/tokend/horizon-connector/internal/resources/operations"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CheckSaleStateResponse struct {
	body *operations.CheckSaleState
	err  error
}

func (r *CheckSaleStateResponse) Unwrap() (*operations.CheckSaleState, error) {
	return r.body, r.err
}

type checkSaleStateStreamer struct {
	q       *Q
	results chan CheckSaleStateResponse
}

func streamCheckSaleState(q *Q, ctx context.Context, buffer int) <-chan CheckSaleStateResponse {
	streamer := checkSaleStateStreamer{
		q:       q,
		results: make(chan CheckSaleStateResponse, buffer),
	}

	go streamer.startStream(ctx)
	return streamer.results
}

func (s *checkSaleStateStreamer) startStream(ctx context.Context) {
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

func (s *checkSaleStateStreamer) stream(ctx context.Context, cursor string) (newCursor string) {
	newCursor = cursor

	defer func() {
		if recovered := recover(); recovered != nil {
			err := errors.FromPanic(recovered)
			s.send(ctx, nil, err)
		}
	}()

	checkSaleStateOperations, err := s.q.opQ.CheckSaleStateOperations(newCursor)
	if err != nil {
		s.send(ctx, nil, errors.Wrap(err, "failed to get check sale state operations"))
		return
	}

	for i := range checkSaleStateOperations {
		if !s.send(ctx, &checkSaleStateOperations[i], nil) {
			return
		}
		newCursor = checkSaleStateOperations[i].PT
	}

	return
}

func (s *checkSaleStateStreamer) send(ctx context.Context, body *operations.CheckSaleState, err error) bool {
	response := CheckSaleStateResponse{
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
