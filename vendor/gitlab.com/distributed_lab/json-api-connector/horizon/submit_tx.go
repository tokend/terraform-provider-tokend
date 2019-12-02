package horizon

import (
	"context"
	"net/http"
	"net/url"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	ErrRequestClosed = errors.New("request closed")
)

func (c *Connector) Submit(ctx context.Context, endpoint *url.URL, envelope string, waitForIngest bool) (int, []byte, error) {
	status, resp, err := c.client.PostJSONWithContext(endpoint, &TransactionSubmit{
		Tx:            envelope,
		WaitForIngest: waitForIngest,
	}, ctx)

	if isContextCanceled(ctx) {
		return http.StatusGatewayTimeout, nil, ErrRequestClosed
	}

	return status, resp, err
}

func isContextCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
