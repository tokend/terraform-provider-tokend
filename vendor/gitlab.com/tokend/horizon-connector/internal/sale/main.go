package sale

import (
	"encoding/json"
	"fmt"

	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/horizon-connector/internal/errors"
	"gitlab.com/tokend/horizon-connector/internal/resources"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

func (q *Q) CoreSales() ([]resources.CoreSale, error) {
	endpoint := fmt.Sprintf("/core_sales")
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var result []resources.CoreSale
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	return result, nil
}

func (q *Q) SaleByID(saleID uint64) (*resources.Sale, error) {
	response, err := q.client.Get(fmt.Sprintf("/sales/%d", saleID))
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		// No such sale
		return nil, nil
	}

	var result resources.Sale
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	return &result, nil
}
