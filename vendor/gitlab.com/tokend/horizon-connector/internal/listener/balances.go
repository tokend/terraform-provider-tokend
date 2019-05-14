package listener

import (
	"context"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal/resources"
)

type BalanceEvent struct {
	Balance *resources.ChoppedBalance
	Err     error
}

// StreamBalancesByAsset stream all balances in given asset once
func (q *Q) StreamBalancesByAsset(ctx context.Context, asset string) <-chan BalanceEvent {
	balancesStream := make(chan BalanceEvent)
	cursor := ""
	go func() {
		defer func() {
			close(balancesStream)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				break
			}

			balances, err := q.balanceQ.BalancesByAsset(asset, cursor)
			if err != nil {
				balancesStream <- BalanceEvent{Err: errors.Wrap(err, "Failed to obtain balances", logan.F{"cursor": cursor})}
				continue
			}

			if len(balances) == 0 {
				// The stream channel is closed in defer.
				return
			}

			for _, balance := range balances {
				ok := streamBalances(ctx, balance, balancesStream)
				if !ok {
					// Ctx was canceled
					return
				}
				cursor = balance.ID
			}
		}
	}()
	return balancesStream
}

func streamBalances(ctx context.Context, balance resources.ChoppedBalance, balanceStream chan<- BalanceEvent) bool {
	select {
	case <-ctx.Done():
		return false
	case balanceStream <- BalanceEvent{Balance: &balance}:
		return true
	}
}
