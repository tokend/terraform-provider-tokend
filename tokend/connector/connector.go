package connector

import (
	"github.com/tokend/terraform-provider-tokend/tokend/connector/keyvalues"
	"github.com/tokend/terraform-provider-tokend/tokend/data"
	"gitlab.com/distributed_lab/json-api-connector/client"
)

type connector struct {
	client client.Client
}

func NewConnector(client client.Client) data.Connector {
	return &connector{
		client: client,
	}
}

func (c *connector) KeyValues() keyvalues.KeyValues {
	return keyvalues.NewKeyValues(c.client)
}
