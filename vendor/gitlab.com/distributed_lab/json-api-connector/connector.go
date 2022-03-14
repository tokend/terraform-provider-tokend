package connector

import (
	"context"
	"encoding/json"
	"io"
	"net/url"

	"gitlab.com/distributed_lab/json-api-connector/base"

	"gitlab.com/distributed_lab/json-api-connector/client"
)

type Connector struct {
	base *base.Connector
}

func NewConnector(client client.Client) *Connector {
	return &Connector{base: base.NewConnector(client)}
}

func (c *Connector) Get(endpoint *url.URL, dst interface{}) (err error) {
	response, err := c.base.Get(endpoint)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) PostJSON(endpoint *url.URL, req interface{}, ctx context.Context, dst interface{},
) (err error) {
	_, response, err := c.base.PostJSON(endpoint, req, ctx)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) Post(endpoint *url.URL, body io.Reader, ctx context.Context, dst interface{},
) (err error) {
	_, response, err := c.base.Post(endpoint, body, ctx)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) Put(endpoint *url.URL, body io.Reader, dst interface{},
) (err error) {
	response, err := c.base.Put(endpoint, body)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) PatchJSON(endpoint *url.URL, req interface{}, ctx context.Context, dst interface{},
) (err error) {
	_, response, err := c.base.PatchJSON(endpoint, req, ctx)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)

}

func (c *Connector) Patch(endpoint *url.URL, body io.Reader, ctx context.Context, dst interface{},
) (err error) {
	_, response, err := c.base.Patch(endpoint, body, ctx)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) PutJSON(endpoint *url.URL, req interface{}, dst interface{},
) (err error) {
	response, err := c.base.PutJSON(endpoint, req)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) Delete(endpoint *url.URL, dst interface{}) (err error) {
	response, err := c.base.Delete(endpoint)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}
