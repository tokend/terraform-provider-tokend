package transactionv2

import (
	"encoding/json"
	"net/url"
	"strconv"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/horizon-connector/internal/responses"
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

// ByEffectsAndEntryTypes do request to horizon to get transactions v2
// by specific entry type and effects, returns transactionsV2 and page meta
func (q *Q) ByEffectsAndEntryTypes(
	cursor string, effects, entryTypes []int,
) ([]regources.Transaction, *regources.PageMeta, error) {
	u := url.Values{}
	u.Add("limit", "1000")
	u.Add("cursor", cursor)
	addQuerySlice(u, "effect", effects)
	addQuerySlice(u, "entry_type", entryTypes)
	response, err := q.client.Get("/v2/transactions?" + u.Encode())
	if err != nil {
		return nil, nil, errors.Wrap(err, "transactions_v2 request failed")
	}

	var result responses.TransactionIndex
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshal transactions_v2")
	}

	return result.Embedded.Records, &result.Embedded.Meta, nil
}

func addQuerySlice(u url.Values, fieldName string, input []int) {
	for _, value := range input {
		u.Add(fieldName, strconv.Itoa(value))
	}
}
