package listener

import (
	"context"
	"gitlab.com/tokend/horizon-connector/internal/balance"
	"gitlab.com/tokend/horizon-connector/internal/operation"
	"gitlab.com/tokend/horizon-connector/internal/transaction"
	"gitlab.com/tokend/horizon-connector/internal/transactionv2"
	"gitlab.com/tokend/regources"
)

// Q wraps queues to use their methods
type Q struct {
	txQ   *transaction.Q
	txV2Q *transactionv2.Q
	// TODO Rename - it'a actually RequestQ
	opQ      *operation.Q
	balanceQ *balance.Q
}

func NewQ(tx *transaction.Q, txV2Q *transactionv2.Q, op *operation.Q, balancesQ *balance.Q) *Q {
	return &Q{
		tx,
		txV2Q,
		op,
		balancesQ,
	}
}

// DEPRECATED: use StreamAllReviewableRequests instead
func (q *Q) Requests(result chan<- regources.ReviewableRequest) <-chan error {
	errs := make(chan error)
	go func() {
		defer func() {
			close(errs)
		}()
		cursor := ""
		for {
			requests, err := q.opQ.AllRequests(cursor)
			if err != nil {
				errs <- err
				continue
			}
			for _, request := range requests {
				result <- request
				cursor = request.PagingToken()
			}
		}
	}()
	return errs
}

// DEPRECATED: Use StreamWithdrawalRequests instead
//func (q *Q) WithdrawalRequests(result chan<- regources.ReviewableRequest) <-chan error {

func (q *Q) StreamAllCheckSaleStateOps(ctx context.Context, buffer int) <-chan CheckSaleStateResponse {
	return streamCheckSaleState(q, ctx, buffer)
}

func (q *Q) StreamAllCreateKYCRequestOps(ctx context.Context, buffer int) <-chan CreateKYCRequestOpResponse {
	return streamCreateKYCRequestOp(q, ctx, buffer)
}

func (q *Q) StreamAllPaymentV2Ops(ctx context.Context, buffer int) <-chan PaymentOpV2Response {
	return streamPaymentV2(q, ctx, buffer)
}

func (q *Q) StreamAllReviewRequestOps(ctx context.Context, buffer int) <-chan ReviewRequestOpResponse {
	return streamReviewRequestOp(q, ctx, buffer)
}
