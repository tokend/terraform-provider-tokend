package connector

import (
	"bytes"
	"context"
	"encoding/json"
	"gitlab.com/distributed_lab/json-api-connector/client"
	internalErrs "gitlab.com/distributed_lab/json-api-connector/cerrors"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Connector struct {
	client client.Client
}

func NewConnector(client client.Client) *Connector {
	return &Connector{client: client}
}

func (c *Connector) Get(endpoint *url.URL, dst interface{}) (err error) {
	fullEndpoint, err := c.client.Resolve(endpoint)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("GET", fullEndpoint, nil)
	if err != nil {
		return internalErrs.E(
			"failed to build request",
			err,
		)
	}
	_, response, err := c.perform(request)

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) PostJSON(endpoint *url.URL, req interface{}, ctx context.Context, dst interface{}) (err error) {
	reqBB, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal request into JSON bytes")
	}

	resolvedEndpoint, err := c.client.Resolve(endpoint)
	if err != nil {
		return errors.Wrap(err, "Failed to resolve url")
	}

	request, err := http.NewRequest("POST", resolvedEndpoint, bytes.NewReader(reqBB))
	if err != nil {
		return errors.Wrap(err, "Failed to create POST http.Request")
	}

	_, response, err := c.perform(request.WithContext(ctx))

	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) Post(endpoint *url.URL, body io.Reader, ctx context.Context, dst interface{}) (err error) {

	resolvedEndpoint, err := c.client.Resolve(endpoint)
	if err != nil {
		return errors.Wrap(err, "Failed to resolve url")
	}

	request, err := http.NewRequest("POST", resolvedEndpoint, body)
	if err != nil {
		return errors.Wrap(err, "Failed to create POST http.Request")
	}

	_, response, err := c.perform(request.WithContext(ctx))

	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) Put(endpoint *url.URL, body io.Reader, dst interface{}) (err error) {
	fullEndpoint, err := c.client.Resolve(endpoint)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("PUT", fullEndpoint, body)
	if err != nil {
		return internalErrs.E(
			"failed to build request",
			err,
		)
	}

	_, response, err := c.perform(request)

	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) PatchJSON(endpoint *url.URL, req interface{}, ctx context.Context, dst interface{}) (err error) {
	reqBB, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal request into JSON bytes")
	}

	resolvedEndpoint, err := c.client.Resolve(endpoint)
	if err != nil {
		return errors.Wrap(err, "Failed to resolve url")
	}

	request, err := http.NewRequest("PATCH", resolvedEndpoint, bytes.NewReader(reqBB))
	if err != nil {
		return errors.Wrap(err, "Failed to create POST http.Request")
	}

	_, response, err := c.perform(request.WithContext(ctx))

	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)

}

func (c *Connector) Patch(endpoint *url.URL, body io.Reader, ctx context.Context, dst interface{}) (err error) {
	resolvedEndpoint, err := c.client.Resolve(endpoint)
	if err != nil {
		return errors.Wrap(err, "Failed to resolve url")
	}

	request, err := http.NewRequest("PATCH", resolvedEndpoint, body)
	if err != nil {
		return errors.Wrap(err, "Failed to create POST http.Request")
	}

	_, response, err := c.perform(request.WithContext(ctx))

	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) PutJSON(endpoint *url.URL, body io.Reader, dst interface{}) (err error) {
	fullEndpoint, err := c.client.Resolve(endpoint)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("PUT", fullEndpoint, body)
	if err != nil {
		return internalErrs.E(
			"failed to build request",
			err,
		)
	}

	_, response, err := c.perform(request)

	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)

}

func (c *Connector) Delete(endpoint *url.URL, dst interface{}) (err error) {
	fullEndpoint, err := c.client.Resolve(endpoint)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("DELETE", fullEndpoint, nil)
	if err != nil {
		return internalErrs.E(
			"failed to build request",
			err,
		)
	}

	_, response, err := c.perform(request)

	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}

func (c *Connector) perform(request *http.Request) (statusCode int, response []byte, err error) {
	resp, err := c.client.Do(request)
	if err != nil {
		return 0, nil, internalErrs.E(
			"failed to perform request",
			err,
			internalErrs.Path(request.URL.String()),
		)
	}
	defer resp.Body.Close()

	bodyBB, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, internalErrs.E(
			"failed to read response body",
			err,
			internalErrs.Path(request.URL.String()),
		)
	}
	statusCode = resp.StatusCode
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return statusCode, bodyBB, nil
	case http.StatusNotFound, http.StatusNoContent, http.StatusAccepted:
		return statusCode, nil, nil
	case http.StatusTooManyRequests:
		panic("not implemented")
	case http.StatusBadRequest:
		return statusCode, nil, internalErrs.E(
			"request was invalid in some way",
			internalErrs.Response(bodyBB),
			internalErrs.Status(resp.StatusCode),
		)
	case http.StatusUnauthorized:
		return statusCode, nil, internalErrs.E(
			"signer is not allowed to access resource",
			internalErrs.Response(bodyBB),
			internalErrs.Status(resp.StatusCode),
			internalErrs.Path(request.URL.String()),
		)
	default:
		return statusCode, nil, internalErrs.E(
			"something bad happened",
			internalErrs.Response(bodyBB),
			internalErrs.Status(resp.StatusCode),
			internalErrs.Path(request.URL.String()),
		)
	}

	return
}
