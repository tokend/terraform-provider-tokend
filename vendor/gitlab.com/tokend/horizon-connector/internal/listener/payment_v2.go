package listener

import (
	"gitlab.com/tokend/horizon-connector/internal/resources/operations"
	"context"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type PaymentOpV2Response struct {
	body *operations.PaymentV2
	err  error
}

func (r *PaymentOpV2Response) Unwrap() (*operations.PaymentV2, error) {
	return r.body, r.err
}

type paymentOpV2Streamer struct {
	q       *Q
	results chan PaymentOpV2Response
}

func (s *paymentOpV2Streamer) send(ctx context.Context, body *operations.PaymentV2, err error) bool {
	response := PaymentOpV2Response{
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

func (s *paymentOpV2Streamer) stream(ctx context.Context, cursor string) (newCursor string) {
	newCursor = cursor

	defer func() {
		if recovered := recover(); recovered != nil {
			err := errors.FromPanic(recovered)
			s.send(ctx, nil, err)
		}
	}()

	paymentV2Operations, err := s.q.opQ.PaymentV2Operations(newCursor)
	if err != nil {
		s.send(ctx, nil, errors.Wrap(err, "failed to get PaymentV2 operations"))
		return
	}

	for _, op := range paymentV2Operations {
		if !s.send(ctx, &op, nil) {
			return
		}
		newCursor = op.PT
	}

	return
}

func (s *paymentOpV2Streamer) startStream(ctx context.Context) {
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

func streamPaymentV2(q *Q, ctx context.Context, buffer int) <-chan PaymentOpV2Response {
	streamer := paymentOpV2Streamer{
		q:       q,
		results: make(chan PaymentOpV2Response, buffer),
	}

	go streamer.startStream(ctx)
	return streamer.results
}
