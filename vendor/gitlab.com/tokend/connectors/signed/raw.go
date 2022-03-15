package signed

import (
	"net/http"
	"net/url"
	"path"

	"gitlab.com/distributed_lab/logan/v3/errors"
	depkeypair "gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/signcontrol"
	"gitlab.com/tokend/keypair"
)

type RawClient struct {
	base   *url.URL
	signer keypair.Full
	source keypair.Address
	// Client must only be called from Client.Do() methods only
	client *http.Client
}

func (c *RawClient) Do(request *http.Request) (*http.Response, error) {
	// ensure content-type just in case
	request.Header.Set("content-type", "application/json")

	if c.signer != nil {
		seed := depkeypair.MustParse(c.signer.Seed())

		request.Header.Set("account-id", c.signer.Address())
		if c.source != nil {
			request.Header.Set("account-id", c.source.Address())
		}

		err := signcontrol.SignRequest(request, seed)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to sign request")
		}
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to perform http request")
	}

	return response, nil
}

func NewRawClient(client *http.Client, base *url.URL) *RawClient {
	return &RawClient{
		base:   base,
		signer: nil,
		source: nil,
		client: client,
	}
}

func (c *RawClient) WithSigner(kp keypair.Full) *RawClient {
	return &RawClient{
		base:   c.base,
		signer: kp,
		source: c.source,
		client: c.client,
	}
}

func (c *RawClient) WithSource(source keypair.Address) *RawClient {
	return &RawClient{
		base:   c.base,
		signer: c.signer,
		source: source,
		client: c.client,
	}
}

func (c *RawClient) Resolve(endpoint *url.URL) (string, error) {
	u := *c.base
	basePath := u.Path
	prevPath := endpoint.Path

	if basePath != "" {
		endpoint.Path = path.Join(basePath, endpoint.Path)
		u.Path = ""
	}

	resolved := u.ResolveReference(endpoint)
	endpoint.Path = prevPath
	return resolved.String(), nil
}

func (c *RawClient) URL() *url.URL {
	return c.base
}
