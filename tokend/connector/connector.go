package connector

import (
	"gitlab.com/tokend/horizon-connector"
	"github.com/tokend/terraform-provider-tokend/tokend/connector/keyvalues"
	"github.com/tokend/terraform-provider-tokend/tokend/data"
)

type connector struct {
	Client *horizon.Client
}

func NewConnector(client *horizon.Client) data.Connector {
	return &connector{
		Client: client,
	}
}

func (c *connector) KeyValues() keyvalues.KeyValues {
	return keyvalues.NewKeyValues(c.Client)
}
