package connector

import (
	"github.com/tokend/terraform-provider-tokend/tokend/connector/keyvalues"
	"github.com/tokend/terraform-provider-tokend/tokend/data"
	connector2 "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/client"
)

type connector struct {
	connector *connector2.Connector
}

func NewConnector(client client.Client) data.Connector {
	return &connector{
		connector: connector2.NewConnector(client),
	}
}

func (c *connector) KeyValues() keyvalues.KeyValues {
	return keyvalues.NewKeyValues(c.connector)
}
