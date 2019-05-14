package listener

import (
	"context"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal/resources"
	"gitlab.com/tokend/regources"
)

// StreamTransactionV2Builder idea is to provide graceful migration to transactionV2 streaming for old clients
func (q *Q) StreamTransactionsV2Builder(effects, entryTypes []int) func(ctx context.Context) (<-chan resources.TransactionEvent, <-chan error) {
	return func(ctx context.Context) (<-chan resources.TransactionEvent, <-chan error) {
		return q.StreamTransactionsV2(ctx, effects, entryTypes)
	}
}

// StreamTransactionsV2 streams transactions fetched for specified filters.
// If there is no new transactions, but ledger has been closed, `TransactionV2Event` with nil tx will be returned.
// Consumer should not rely on closing of any of this channels.
func (q *Q) StreamTransactionsV2(ctx context.Context, effects, entryTypes []int) (<-chan resources.TransactionEvent, <-chan error) {
	txStream := make(chan resources.TransactionEvent)
	errChan := make(chan error)

	go func() {
		cursor := ""
		for {
			select {
			case <-ctx.Done():
				return
			default:
				break
			}

			transactionsV2, meta, err := q.txV2Q.ByEffectsAndEntryTypes(cursor, effects, entryTypes)
			if err != nil {
				errChan <- errors.Wrap(err, "Failed to obtain Transactions", logan.F{
					"cursor":      cursor,
					"effects":     effects,
					"entry_types": entryTypes,
				})
				continue
			}

			for _, tx := range transactionsV2 {
				ohaigo := tx

				txEvent := resources.TransactionEvent{
					Transaction: &ohaigo,
					// emulating discrete transactions stream by spoofing meta
					// to not let bump cursor too much before actually consuming all transactions
					Meta: regources.PageMeta{
						LatestLedger: regources.LedgerMeta{
							ClosedAt: tx.LedgerCloseTime,
						},
					},
				}
				ok := streamTxEvent(ctx, txEvent, txStream)
				if !ok {
					// Ctx was canceled
					return
				}

				cursor = tx.PT
			}

			// letting consumer know about current ledger cursor
			ok := streamTxEvent(ctx, resources.TransactionEvent{
				Transaction: nil,
				Meta:        *meta,
			}, txStream)
			if !ok {
				// Ctx was canceled
				return
			}
		}
	}()

	return txStream, errChan
}
