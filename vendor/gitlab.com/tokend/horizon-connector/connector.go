package horizon

import (
	"encoding/json"
	"gitlab.com/tokend/horizon-connector/internal/poll"
	"net/http"
	"net/url"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/horizon-connector/internal/account"
	"gitlab.com/tokend/horizon-connector/internal/asset"
	"gitlab.com/tokend/horizon-connector/internal/balance"
	"gitlab.com/tokend/horizon-connector/internal/blob"
	"gitlab.com/tokend/horizon-connector/internal/document"
	"gitlab.com/tokend/horizon-connector/internal/keyvalue"
	"gitlab.com/tokend/horizon-connector/internal/listener"
	"gitlab.com/tokend/horizon-connector/internal/operation"
	"gitlab.com/tokend/horizon-connector/internal/orders"
	"gitlab.com/tokend/horizon-connector/internal/sale"
	"gitlab.com/tokend/horizon-connector/internal/system"
	"gitlab.com/tokend/horizon-connector/internal/templates"
	"gitlab.com/tokend/horizon-connector/internal/transaction"
	"gitlab.com/tokend/horizon-connector/internal/transactionv2"
	"gitlab.com/tokend/horizon-connector/internal/user"
	"gitlab.com/tokend/horizon-connector/internal/wallets"
	"gitlab.com/tokend/keypair"
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

func (c *Connector) Clone() *Connector {
	return &Connector{
		client: c.client.clone(),
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
	info, err := c.System().Info()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get horizon info")
	}

	return xdrbuild.NewBuilder(info.Passphrase, info.TXExpirationPeriod), nil
}

// DEPRECATED: use .System().Info() instead
func (c *Connector) Info() (info *Info, err error) {
	response, err := c.client.Get("/")
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	if err := json.Unmarshal(response, &info); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal info")
	}
	return info, nil
}

func (c *Connector) KeyValue() *keyvalue.Q {
	return keyvalue.NewQ(c.client)
}

func (c *Connector) System() *system.Q {
	return system.NewQ(c.client)
}

func (c *Connector) Submitter() *Submitter {
	return &Submitter{
		client: c.client,
	}
}

func (c *Connector) Assets() *asset.Q {
	return asset.NewQ(c.client)
}

func (c *Connector) Orders() *orders.Q {
	return orders.NewQ(c.client)
}

func (c *Connector) Accounts() *account.Q {
	return account.NewQ(c.client)
}

func (c *Connector) Transactions() *transaction.Q {
	return transaction.NewQ(c.client)
}

func (c *Connector) TransactionsV2() *transactionv2.Q {
	return transactionv2.NewQ(c.client)
}

func (c *Connector) Sales() *sale.Q {
	return sale.NewQ(c.client)
}

func (c *Connector) Polls() *poll.Q {
	return poll.NewQ(c.client)
}

func (c *Connector) Users() *user.Q {
	return user.NewQ(c.client)
}

func (c *Connector) Balances() *balance.Q {
	return balance.NewQ(c.client)
}

func (c *Connector) Listener() *listener.Q {
	// TODO Rename Operations to Requests? it does actually manages Requests only.
	return listener.NewQ(c.Transactions(), c.TransactionsV2(), c.Operations(), c.Balances())
}

// TODO Rename to Requests? it does actually manages Requests only.
func (c *Connector) Operations() *operation.Q {
	return operation.NewQ(c.client)
}

func (c *Connector) Blobs() *blob.Q {
	return blob.NewQ(c.client)
}

func (c *Connector) Documents() *document.Q {
	return document.NewQ(c.client)
}

func (c *Connector) Templates() *templates.Q {
	return templates.NewQ(c.client)
}

func (c *Connector) Wallets() *wallets.Q {
	return wallets.NewQ(c.client)
}
