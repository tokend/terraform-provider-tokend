package keyvalues

import (
	"encoding/json"
	"fmt"
	"github.com/tokend/terraform-provider-tokend/tokend/horizon"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/regources/generated"
)

//go:generate mockery -case underscore -name KeyValues
type KeyValues interface {
	Value(key string) (*regources.KeyValueEntryValue, error)
}

type keyValues struct {
	client *horizon.Client
}

func NewKeyValues(client *horizon.Client) KeyValues {
	return &keyValues{
		client: client,
	}
}

func (q *keyValues) Value(key string) (*regources.KeyValueEntryValue, error) {
	resp, err := q.client.Get(fmt.Sprintf("/v3/key_values/%s", key))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get key value")
	}

	if resp == nil {
		return nil, nil
	}

	var result regources.KeyValueEntryResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal key value")
	}

	return &result.Data.Attributes.Value, nil
}
