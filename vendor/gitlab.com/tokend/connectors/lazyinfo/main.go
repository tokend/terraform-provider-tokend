package lazyinfo

import (
	connector "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/client"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/regources/generated"
	"net/url"
	"sync"
)

type LazyInfoer struct {
	b    *connector.Connector
	once sync.Once
	val  regources.HorizonState
	err  error
}

func New(client client.Client) *LazyInfoer {
	return &LazyInfoer{
		b: connector.NewConnector(client),
	}
}

func (g *LazyInfoer) Info() (regources.HorizonState, error) {
	g.once.Do(func() {
		endpoint, _ := url.Parse("/v3/info")

		var response regources.HorizonStateResponse
		err := g.b.Get(endpoint, &response)
		if err != nil {
			g.err = errors.Wrap(err, "failed to get horizon info")
			return
		}

		g.val = response.Data
	})
	return g.val, g.err
}
