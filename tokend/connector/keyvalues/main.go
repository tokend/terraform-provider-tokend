package keyvalues

import (
	"fmt"
	connector "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/cerrors"
	url2 "net/url"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/regources"
)

//go:generate mockery -case underscore -name KeyValues
type KeyValues interface {
	Value(key string) (*regources.KeyValueEntryValue, error)
}

type keyValues struct {
	conn *connector.Connector
}

func NewKeyValues(conn *connector.Connector) KeyValues {
	return &keyValues{
		conn: conn,
	}
}

func (q *keyValues) Value(key string) (*regources.KeyValueEntryValue, error) {
	url, err := url2.Parse(fmt.Sprintf("/horizon/key_values/%s", key))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}

	var result regources.KeyValueEntryResponse
	err = q.conn.Get(url, result)
	if err != nil {
		if cerrors.NotFound(err) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to get key value")
	}

	return &result.Data.Attributes.Value, nil
}
