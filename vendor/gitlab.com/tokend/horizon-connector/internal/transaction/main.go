package transaction

import (
	"encoding/json"

	"fmt"

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

func (q *Q) Transactions(cursor string) ([]regources.Transaction, *regources.PageMeta, error) {
	response, err := q.client.Get(fmt.Sprintf("/transactions?limit=2000000000&cursor=%s", cursor))
	if err != nil {
		return nil, nil, errors.Wrap(err, "request failed")
	}

	var result responses.TransactionIndex
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshal")
	}
	return result.Embedded.Records, &result.Embedded.Meta, nil
}

func (q *Q) TransactionByID(txID string) (*regources.Transaction, error) {
	response, err := q.client.Get(fmt.Sprintf("/transactions/%s", txID))
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		// No such Transaction
		return nil, nil
	}

	var result regources.Transaction
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	return &result, nil
}
