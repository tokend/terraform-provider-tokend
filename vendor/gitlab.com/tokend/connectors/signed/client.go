package signed

import (
	"net/http"
	"net/url"
	"time"

	"gitlab.com/tokend/keypair"
)

func throttle() chan time.Time {
	burst := 2 << 10
	ch := make(chan time.Time, burst)

	go func() {
		tick := time.Tick(1 * time.Second)
		// prefill buffer
		for i := 0; i < burst; i++ {
			ch <- time.Now()
		}
		for {
			select {
			case ch <- <-tick:
			default:
			}
		}
	}()
	return ch
}

type Client struct {
	rc       *RawClient
	throttle chan time.Time
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	<-c.throttle
	return c.rc.Do(request)
}

func NewClient(client *http.Client, base *url.URL) *Client {
	return &Client{
		rc: &RawClient{
			base:   base,
			signer: nil,
			source: nil,
			client: client,
		},
		throttle: throttle(),
	}
}

func (c *Client) WithSigner(kp keypair.Full) *Client {
	return &Client{
		rc: &RawClient{
			base:   c.rc.base,
			signer: kp,
			source: c.rc.source,
			client: c.rc.client,
		},
		throttle: c.throttle,
	}
}

func (c *Client) WithSource(source keypair.Address) *Client {
	return &Client{
		rc: &RawClient{
			base:   c.rc.base,
			signer: c.rc.signer,
			source: source,
			client: c.rc.client,
		},
		throttle: c.throttle,
	}
}

func (c *Client) Resolve(endpoint *url.URL) (string, error) {
	return c.rc.Resolve(endpoint)
}

func (c *Client) URL() *url.URL {
	return c.rc.base
}
