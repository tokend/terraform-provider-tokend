package keyvalue

import (
	"fmt"
	"github.com/pkg/errors"
	connector "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/client"
	regources "gitlab.com/tokend/regources/generated"
	"net/url"
)

type KeyValuer struct {
	client *connector.Connector
}

func New(client client.Client) *KeyValuer {
	return &KeyValuer{
		client: connector.NewConnector(client),
	}
}

func (g *KeyValuer) KeyValue(key string) (regources.KeyValueEntry, error) {
	endpoint, err := url.Parse(fmt.Sprintf("/v3/key_values/%s", key))
	if err != nil {
		return regources.KeyValueEntry{}, errors.Wrap(err, "failed to parse url")
	}

	var response regources.KeyValueEntryResponse
	err = g.client.Get(endpoint, &response)
	if err != nil {
		return response.Data, errors.Wrap(err, "failed to get key value")
	}

	return response.Data, nil
}

func (g *KeyValuer) KeyValueUInt32(key string) (uint32, error) {
	entry, err := g.KeyValue(key)
	if err != nil {
		return 0, err
	}

	value := entry.Attributes.Value.U32
	if value == nil {
		return 0, errors.Errorf("value was found but has different type %s", entry.Attributes.Value.Type.String())
	}

	return *value, nil
}

func (g *KeyValuer) KeyValueUInt64(key string) (uint64, error) {
	entry, err := g.KeyValue(key)
	if err != nil {
		return 0, err
	}

	value := entry.Attributes.Value.U64
	if value == nil {
		return 0, errors.Errorf("value was found but has different type %s", entry.Attributes.Value.Type.String())
	}

	return *value, nil
}

func (g *KeyValuer) KeyValueString(key string) (string, error) {
	entry, err := g.KeyValue(key)
	if err != nil {
		return "", err
	}

	value := entry.Attributes.Value.Str
	if value == nil {
		return "", errors.Errorf("value was found but has different type %s", entry.Attributes.Value.Type.String())
	}

	return *value, nil
}
