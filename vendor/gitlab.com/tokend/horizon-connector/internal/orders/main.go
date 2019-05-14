package orders

import (
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/regources/v2"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

// DEPRECATED
func (q Q) ForAssetPair(orderBookID int64, baseAsset, quoteAsset string, isBuy bool) (*regources.OrderBookEntriesResponse, error) {
	order := "asc"
	if isBuy {
		order = "desc"
	}

	endpoint := fmt.Sprintf("/v3/order_book/%d?filter[base_asset]=%s&filter[quote_asset]=%s&filter[is_buy]=%v&page[order]=%s", orderBookID,
		baseAsset, quoteAsset, isBuy, order)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var orderBook regources.OrderBookEntriesResponse
	if err := json.Unmarshal(response, &orderBook); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response", logan.F{
			"raw_response": string(response),
		})
	}
	return &orderBook, nil
}

