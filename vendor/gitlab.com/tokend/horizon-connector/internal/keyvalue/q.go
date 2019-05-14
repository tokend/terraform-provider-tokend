package keyvalue

import (
	"encoding/json"
	"fmt"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal"
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

func (q *Q) KeyValue() (keyvalue []resources.KeyValue, err error) {
	response, err := q.client.Get("/keyvalue/")
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if err := json.Unmarshal(response, &keyvalue); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal keyvalue")
	}
	return keyvalue, nil
}

func (q *Q) Value(key string) (result *resources.KeyValue, err error) {
	url := fmt.Sprint("/keyvalue/", key)
	response, err := q.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal key")
	}
	return result, nil
}
