package templates

import (
	"fmt"
	"io"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

func (q *Q) Get(id string) ([]byte, error) {
	endpoint := fmt.Sprintf("/templates/%s", id)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	return response, nil
}

func (q *Q) Put(id string, body io.Reader) error {
	endpoint := fmt.Sprintf("/templates/%s", id)

	_, err := q.client.Put(endpoint, body)
	if err != nil {
		return errors.Wrap(err, "Failed to send request")
	}

	return nil
}
