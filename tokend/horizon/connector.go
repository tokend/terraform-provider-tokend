package horizon

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
	regources "gitlab.com/tokend/regources/generated"
	"net/http"
	"net/url"
)

type Connector struct {
	client *Client
}

func NewConnector(base *url.URL) *Connector {
	client := NewClient(http.DefaultClient, base)
	return &Connector{
		client,
	}
}

func (c *Connector) Base() *url.URL {
	return c.client.base
}

func (c *Connector) WithSigner(kp keypair.Full) *Connector {
	return &Connector{
		c.client.WithSigner(kp),
	}
}

func (c *Connector) Client() *Client {
	return c.client
}

func (c *Connector) TXBuilder() (*xdrbuild.Builder, error) {
	info, err := c.Info()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get horizon info")
	}

	return xdrbuild.NewBuilder(info.Attributes.NetworkPassphrase, info.Attributes.TxExpirationPeriod), nil
}

func (c *Connector) Info() (info *regources.HorizonState, err error) {
	response, err := c.client.Get("/v3/info")
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	if err := json.Unmarshal(response, &info); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal info")
	}
	return info, nil
}


func (c *Connector) Submitter() *Submitter {
	return &Submitter{
		client: c.client,
	}
}
