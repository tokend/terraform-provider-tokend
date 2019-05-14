package listener

import (
	"context"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal/resources"
	"gitlab.com/tokend/regources"
)

// DEPRECATED
// Use StreamTXs instead
func (q *Q) Transactions(result chan<- resources.TransactionEvent) <-chan error {
	errs := make(chan error)
	go func() {
		defer func() {
			close(errs)
		}()
		cursor := ""
		for {
			transactions, meta, err := q.txQ.Transactions(cursor)
			if err != nil {
				errs <- err
				continue
			}
			for _, tx := range transactions {
				ohaigo := tx
				result <- resources.TransactionEvent{
					Transaction: &ohaigo,
					// emulating discrete transactions stream by spoofing meta
					// to not let bump cursor too much before actually consuming all transactions
					Meta: regources.PageMeta{
						LatestLedger: regources.LedgerMeta{
							ClosedAt: tx.LedgerCloseTime,
						},
					},
				}
				cursor = tx.PagingToken()
			}
			// letting consumer know about current ledger cursor
			result <- resources.TransactionEvent{
				Transaction: nil,
				Meta:        *meta,
			}
		}
	}()
	return errs
}

// DEPRECATED
// Use StreamTXs instead
func (q *Q) StreamTransactions(ctx context.Context) (<-chan resources.TransactionEvent, <-chan error) {
	txStream := make(chan resources.TransactionEvent)
	errChan := make(chan error)

	go func() {
		defer func() {
			close(txStream)
			close(errChan)
		}()

		cursor := ""
		for {
			select {
			case <-ctx.Done():
				return
			default:
				break
			}

			transactions, meta, err := q.txQ.Transactions(cursor)
			if err != nil {
				errChan <- errors.Wrap(err, "Failed to obtain Transactions", logan.F{"cursor": cursor})
				continue
			}

			for _, tx := range transactions {
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

				cursor = tx.PagingToken()
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

// StreamTXs streams all Transactions from Horizon from the very beginning of TX history
// into returned channel (TransactionEvent + error).
//
// See StreamTXsFromCursor method for more details.
func (q *Q) StreamTXs(ctx context.Context, stopOnEmptyPage bool) <-chan TXPacket {
	return q.StreamTXsFromCursor(ctx, "", stopOnEmptyPage)
}

// StreamTXsFromCursor obtains Transactions from Horizon in pages (using txQ),
// starting from the provided cursor.
// And streams TXPackets(TransactionEvent + error) into returned channel.
//
// Cursor is PagingToken, use empty string to stream all Transactions from the very beginning.
//
// StreamTXsFromCursor starts goroutine inside and returns immediately.
//
// Errors happening during TX streaming will be sent via returned channel in a TXPacket.
// If TXPacket contains non-nil error, the TransactionEvent is nil.
//
// Returned TransactionEvent can have nil Transaction with non-empty Meta,
// such a TransactionEvent is sent at the end of every page of Transactions retrieved from Horizon.
func (q *Q) StreamTXsFromCursor(ctx context.Context, cursor string, stopOnEmptyPage bool) <-chan TXPacket {
	txStream := make(chan TXPacket)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				txStream <- TXPacket{
					body: nil,
					err:  errors.Wrap(errors.FromPanic(r), "panic happened, stopping work, leaving channel unclosed"),
				}
			} else {
				close(txStream)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				break
			}

			transactions, meta, err := q.txQ.Transactions(cursor)
			if err != nil {
				streamTxPacket(ctx, TXPacket{
					body: nil,
					err:  err,
				}, txStream)
				continue
			}

			if stopOnEmptyPage && len(transactions) == 0 {
				// The stream channel is closed in defer.
				return
			}

			for _, tx := range transactions {
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
				ok := streamTxPacket(ctx, TXPacket{
					body: &txEvent,
				}, txStream)
				if !ok {
					// Ctx was canceled
					return
				}
				cursor = tx.PagingToken()
			}

			// letting consumer know about current ledger cursor
			ok := streamTxPacket(ctx, TXPacket{
				body: &resources.TransactionEvent{
					Transaction: nil,
					Meta:        *meta,
				},
			}, txStream)
			if !ok {
				// Ctx was canceled
				return
			}
		}
	}()

	return txStream
}

func streamTxEvent(ctx context.Context, txEvent resources.TransactionEvent, txStream chan<- resources.TransactionEvent) bool {
	select {
	case <-ctx.Done():
		return false
	case txStream <- txEvent:
		return true
	}
}

func streamTxPacket(ctx context.Context, txPacket TXPacket, txStream chan<- TXPacket) bool {
	select {
	case <-ctx.Done():
		return false
	case txStream <- txPacket:
		return true
	}
}
