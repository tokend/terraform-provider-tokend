package horizon

import (
	"fmt"
	"gitlab.com/tokend/regources"
	"path"

	lru "github.com/hashicorp/golang-lru"
	connector "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdrbuild"
)

type Connector struct {
	client            *Client
	receivers         *lru.Cache
	useCache          bool
	validateResponses bool
}

func NewConnector(client *Client, useCache bool) *Connector {
	cache, _ := lru.New(2 << 10)

	connector := &Connector{
		client:            client,
		receivers:         cache,
		useCache:          useCache,
		validateResponses: true,
	}

	return connector
}

type InfoPathParams struct{}

func (p *InfoPathParams) Path() string {
	return ""
}

func (c *Connector) TXBuilder() (*xdrbuild.Builder, error) {
	var info regources.IngesterStateResponse
	err := c.One("/info", &InfoPathParams{}).ValidateResponses(false).Get(&info)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get info")
	}
	return xdrbuild.NewBuilder(info.Data.Attributes.NetworkPassphrase, info.Data.Attributes.TxExpirationPeriod), nil
}

func (c *Connector) CacheReceivers(cache bool) connector.Connector {
	c.useCache = cache
	return c
}

func (c *Connector) One(endpoint string, pathParams types.PathParamer) connector.Singler {
	if !c.useCache {
		return NewSingler(endpoint, c.client, pathParams)
	}

	key := fmt.Sprintf("%s:%s", path.Join(endpoint, pathParams.Path()), "one")
	if r, ok := c.receivers.Get(key); ok {
		receiver, ok := r.(connector.Singler)
		if !ok {
			panic("cast failed")
		}
		return receiver
	}

	singler := NewSingler(endpoint, c.client, pathParams)
	c.receivers.Add(key, singler)

	return singler
}

func (c *Connector) List(endpoint string) connector.Nexter {
	if !c.useCache {
		return NewNexter(endpoint, c.client)
	}

	key := fmt.Sprintf("%s:%s", endpoint, "mul")
	if r, ok := c.receivers.Get(key); ok {
		receiver, ok := r.(connector.Nexter)
		if !ok {
			panic("cast failed")
		}
		return receiver
	}

	nexter := NewNexter(endpoint, c.client)
	c.receivers.Add(key, nexter)

	return nexter
}
