package connector

import (
	"github.com/tokend/terraform-provider-tokend/tokend/connector/keyvalues"
	"github.com/tokend/terraform-provider-tokend/tokend/connector/tx"
	"github.com/tokend/terraform-provider-tokend/tokend/data"
	"gitlab.com/distributed_lab/json-api-connector/client"
	"gitlab.com/tokend/connectors/submit"
)

type connector struct {
	client    client.Client
	submitter *submit.Submitter
}

func NewConnector(client client.Client, submitter *submit.Submitter) data.Connector {
	return &connector{
		client:    client,
		submitter: submitter,
	}
}

func (c *connector) KeyValues() keyvalues.KeyValues {
	return keyvalues.NewKeyValues(c.client)
}

func (c *connector) Submitter() tx.HorizonSubmitter {
	return tx.NewSubmitter(c.submitter)
}
