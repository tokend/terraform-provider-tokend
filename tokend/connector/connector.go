package connector

import (
	"github.com/tokend/terraform-provider-tokend/tokend/connector/keyvalues"
	"github.com/tokend/terraform-provider-tokend/tokend/data"
	"github.com/tokend/terraform-provider-tokend/tokend/horizon"
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
