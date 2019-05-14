package listener

import (
	"gitlab.com/tokend/horizon-connector/internal/resources/operations"
	"context"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateKYCRequestOpResponse struct {
	body *operations.CreateKYCRequest
	err  error
}

func (r *CreateKYCRequestOpResponse) Unwrap() (*operations.CreateKYCRequest, error) {
	return r.body, r.err
}

type createKYCRequestOpStreamer struct {
	q       *Q
	results chan CreateKYCRequestOpResponse
}

func (s *createKYCRequestOpStreamer) send(ctx context.Context, body *operations.CreateKYCRequest, err error) bool {
	response := CreateKYCRequestOpResponse{
		body: body,
		err:  err,
	}
	select {
	case <-ctx.Done():
		return false
	case s.results <- response:
		return true;
	}
}

func (s *createKYCRequestOpStreamer) stream(ctx context.Context, cursor string) (newCursor string) {
	newCursor = cursor

	defer func() {
		if recovered := recover(); recovered != nil {
			err := errors.FromPanic(recovered)
			s.send(ctx, nil, err)
		}
	}()

	createKYCRequestOperations, err := s.q.opQ.CreateKYCRequestOperations(newCursor)
	if err != nil {
		s.send(ctx, nil, errors.Wrap(err, "failed to get create KYC request operations"))
		return
	}

	for i := range createKYCRequestOperations {
		if !s.send(ctx, &createKYCRequestOperations[i], nil) {
			return
		}
		newCursor = createKYCRequestOperations[i].PT
	}

	return
}

func (s *createKYCRequestOpStreamer) startStream(ctx context.Context) {
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

func streamCreateKYCRequestOp(q *Q, ctx context.Context, buffer int) <-chan CreateKYCRequestOpResponse {
	streamer := createKYCRequestOpStreamer{
		q:       q,
		results: make(chan CreateKYCRequestOpResponse, buffer),
	}

	go streamer.startStream(ctx)
	return streamer.results
}
