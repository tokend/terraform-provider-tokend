package system

import (
	"encoding/json"
	"fmt"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/regources"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

// Info returns basic information about the network
func (q *Q) Info() (info *regources.Info, err error) {
	response, err := q.client.Get("/")
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if err := json.Unmarshal(response, &info); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal info")
	}
	return info, nil
}

func (q *Q) Statistics() (stats *regources.SystemStatistics, err error) {
	response, err := q.client.Get("/statistics")
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if err := json.Unmarshal(response, &stats); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal system stats")
	}
	return stats, nil
}

func (q *Q) Balances(assetType string, threshold int64) (stats *regources.BalancesReport, err error) {
	url := fmt.Sprintf("/statistics/balances?asset_code=%s&threshold=%d", assetType, threshold)
	response, err := q.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "request failed", logan.F{
			"url": url,
		})
	}

	if err := json.Unmarshal(response, &stats); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal balances", logan.F{
			"url": url,
		})
	}
	return stats, nil
}
