package horizon

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/xeipuuv/gojsonschema"
	"gitlab.com/distributed_lab/json-api-connector/horizon/jsonschema"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var ErrNotFound = errors.New("resource not found")
var ErrDataEmpty = errors.New("no data on page")
var ErrLinksNull = errors.New("got null links") // FIXME: remove after new links in account signers would arrive

type Getter struct {
	Endpoint *url.URL
	Client   *Client

	shouldValidateResponses bool
	schema                  *gojsonschema.Schema
}

func NewGetter(endpoint *url.URL, client *Client) *Getter {
	return &Getter{
		Endpoint:                endpoint,
		Client:                  client,
		shouldValidateResponses: true,
	}
}

func (getter *Getter) retrieve(dst interface{}, postRetrievers ...func(dst interface{}, err error) error) error {
	resp, err := getter.Client.Get(getter.Endpoint)
	if err != nil {
		return errors.Wrap(err, "cannot get resources")
	}
	if resp == nil {
		return ErrNotFound
	}

	if getter.shouldValidateResponses {
		if err := jsonschema.EnsureValid(dst, resp); err != nil {
			return errors.Wrap(err, "invalid response")
		}
	}

	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&dst); err != nil {
		return errors.Wrap(err, "failed to unmarshal response")
	}

	return nil
}
