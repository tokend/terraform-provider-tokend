package horizon

import (
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"io/ioutil"

	"io"

	"bytes"
	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	depkeypair "gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/signcontrol"
	internalErrs "gitlab.com/tokend/horizon-connector/internal/errors"
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
	base     *url.URL
	signer   keypair.Full
	throttle chan time.Time
	// Client must only be called from Client.Do() methods only
	client *http.Client
}

func NewClient(client *http.Client, base *url.URL) *Client {
	return &Client{
		base, nil, throttle(), client,
	}
}

func (c *Client) Get(endpoint string) ([]byte, error) {
	endpoint, err := c.prepareURL(endpoint)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, internalErrs.E(
			"failed to build request",
			err,
			internalErrs.Runtime,
		)
	}

	return c.Do(request)
}

func (c *Client) PostJSON(endpoint string, req interface{}) (statusCode int, response []byte, err error) {
	reqBB, err := json.Marshal(req)
	if err != nil {
		return 0, nil, errors.Wrap(err, "Failed to marshal request into JSON bytes")
	}
	fields := logan.F{
		"raw_request": string(reqBB),
	}

	endpoint, err = c.resolveURL(endpoint)
	if err != nil {
		return 0, nil, errors.Wrap(err, "Failed to resolve url")
	}
	fields["resolved_url"] = endpoint

	request, err := http.NewRequest("POST", endpoint, bytes.NewReader(reqBB))
	if err != nil {
		return 0, nil, errors.Wrap(err, "Failed to create POST http.Request", fields)
	}

	statusCode, responseBB, err := c.do(request)
	if err != nil {
		return 0, nil, errors.Wrap(err, "Failed to do the request", fields)
	}

	return statusCode, responseBB, nil
}

// DEPRECATED
// Use PostJSON, it doesn't use strange errors, like this method does.
func (c *Client) Post(endpoint string, body io.Reader) ([]byte, error) {
	endpoint, err := c.prepareURL(endpoint)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, internalErrs.E(
			"failed to build request",
			err,
			internalErrs.Runtime,
		)
	}

	return c.Do(request)
}

func (c *Client) Put(endpoint string, body io.Reader) ([]byte, error) {
	endpoint, err := c.prepareURL(endpoint)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PUT", endpoint, body)
	if err != nil {
		return nil, internalErrs.E(
			"failed to build request",
			err,
			internalErrs.Runtime,
		)
	}

	return c.Do(request)
}

func (c *Client) Delete(endpoint string) ([]byte, error) {
	endpoint, err := c.prepareURL(endpoint)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return nil, internalErrs.E(
			"failed to build request",
			err,
			internalErrs.Runtime,
		)
	}

	return c.Do(request)
}

func (c *Client) WithSigner(kp keypair.Full) *Client {
	return &Client{
		c.base,
		kp,
		c.throttle,
		http.DefaultClient,
	}
}

// Do only returns non-nil error, if request hasn't happened successfully or response could not be read.
//
// There's no common handling of erroneous status codes, all codes will be returned without error;
// consumers of do() should handle status codes themselves,
// for example, Get() method would like to return nil-error in case 404,
// however Post() method should behave in other way.
func (c *Client) do(request *http.Request) (int, []byte, error) {
	<-c.throttle

	// ensure content-type just in case
	request.Header.Set("content-type", "application/json")

	if c.signer != nil {
		// TODO move to proper keypair
		err := signcontrol.SignRequest(request, depkeypair.MustParse(c.signer.Seed()))
		if err != nil {
			return 0, nil, errors.Wrap(err, "Failed to sign request")
		}
	}

	response, err := c.client.Do(request)
	if err != nil {
		return 0, nil, errors.Wrap(err, "Failed to perform http request")
	}

	defer response.Body.Close()

	respBB, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, nil, errors.Wrap(err, "Failed to read response body", logan.F{
			"status_code": response.StatusCode,
		})
	}

	return response.StatusCode, respBB, nil
}

func (c *Client) resolveURL(endpoint string) (string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", errors.Wrap(err, "Failed to parse endpoint into URL")
	}

	return c.base.ResolveReference(u).String(), nil
}

func isStatusCodeSuccessful(code int) bool {
	return code >= 200 && code < 300
}

// DEPRECATED
// Don't use Do() method in new places, use do() instead,
// new do() doesn't use deprecated errors and doesn't parse status code, but returns it instead.
func (c *Client) Do(request *http.Request) ([]byte, error) {
	<-c.throttle

	// ensure content-type just in case
	request.Header.Set("content-type", "application/json")

	if c.signer != nil {
		// TODO move to proper keypair
		signcontrol.SignRequest(request, depkeypair.MustParse(c.signer.Seed()))
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, internalErrs.E(
			"failed to perform request",
			err,
			internalErrs.Network,
			internalErrs.Path(request.URL.String()),
		)
	}
	defer response.Body.Close()

	bodyBB, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, internalErrs.E(
			"failed to read response body",
			err,
			internalErrs.Runtime,
			internalErrs.Path(request.URL.String()),
		)
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return bodyBB, nil
	case http.StatusNotFound, http.StatusNoContent:
		return nil, nil
	case http.StatusTooManyRequests:
		// TODO look at x-rate-limit headers and slow down
		panic("not implemented")
	case http.StatusBadRequest:
		return nil, internalErrs.E(
			"request was invalid in some way",
			internalErrs.Runtime,
			internalErrs.Response(bodyBB),
			internalErrs.Status(response.StatusCode),
			internalErrs.Path(request.URL.String()),
		)
	case http.StatusUnauthorized:
		return nil, internalErrs.E(
			"signer is not allowed to access resource",
			internalErrs.Runtime,
			internalErrs.Response(bodyBB),
			internalErrs.Status(response.StatusCode),
			internalErrs.Path(request.URL.String()),
		)
	default:
		return nil, internalErrs.E(
			"something bad happened",
			internalErrs.Runtime,
			internalErrs.Response(bodyBB),
			internalErrs.Status(response.StatusCode),
			internalErrs.Path(request.URL.String()),
		)
	}
}

// DEPRECATED
// Use resolveURL, it doesn't use deprecated errors, like this method does.
func (c *Client) prepareURL(endpoint string) (string, error) {
	if !strings.HasPrefix(endpoint, "/") {
		return "", errors.New("endpoint should start with /")
	}

	cpy, _ := url.Parse(c.base.String())
	cpy.Path = path.Join(c.base.Path, endpoint)

	// FIXME skipped client test
	result, err := url.PathUnescape(cpy.String())
	if err != nil {
		return "", errors.Wrap(err, "failed to unescape query")
	}
	return result, nil
}
