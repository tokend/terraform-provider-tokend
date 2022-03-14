package submit

import (
	connector "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/client"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/connectors/lazyinfo"
	"gitlab.com/tokend/go/xdrbuild"
	"net/url"
)

type Submitter struct {
	base          *connector.Connector
	info          *lazyinfo.LazyInfoer
	submissionUrl *url.URL
}

func New(client client.Client) *Submitter {
	submission, _ := url.Parse("/v3/transactions")

	return &Submitter{
		base:          connector.NewConnector(client),
		info:          lazyinfo.New(client),
		submissionUrl: submission,
	}
}

func (t *Submitter) TXBuilder() (*xdrbuild.Builder, error) {
	info, err := t.info.Info()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get horizon info")
	}

	return xdrbuild.NewBuilder(info.Attributes.NetworkPassphrase, info.Attributes.TxExpirationPeriod), nil
}
