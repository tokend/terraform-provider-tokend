package keyvalues

import (
	regources "gitlab.com/tokend/regources/generated"

	"gitlab.com/distributed_lab/json-api-connector/client"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/connectors/keyvalue"
)

//go:generate mockery -case underscore -name KeyValues
type KeyValues interface {
	Value(key string) (*regources.KeyValueEntryValue, error)
}

type keyValues struct {
	client *keyvalue.KeyValuer
}

func NewKeyValues(client client.Client) KeyValues {
	return &keyValues{
		client: keyvalue.New(client),
	}
}

func (q *keyValues) Value(key string) (*regources.KeyValueEntryValue, error) {
	entry, err := q.client.KeyValue(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get key value entry")
	}

	return &entry.Attributes.Value, nil
}
