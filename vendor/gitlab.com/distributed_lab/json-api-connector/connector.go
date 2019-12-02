package connector

import (
	"context"
	"net/url"

	"gitlab.com/distributed_lab/json-api-connector/types"
	"gitlab.com/tokend/go/xdrbuild"
)

// Interfaces are here only in purposes of implementing other custom connectors

type Connector interface {
	One(endpoint string, pathParams types.PathParamer) Singler
	List(endpoint string) Nexter
	// Submit *must* submit transaction to specified endpoint and return:
	// - nil if transaction was successfully submitted
	// - error if smth went wrong (specify what exactly)
	Submit(ctx context.Context, endpoint *url.URL, envelope string, waitForIngest bool) (int, []byte, error)
	// TXBuilder() must return xdr builder for building transactions
	// or error if cannot get horizon info from `/v3/info` endpoint
	TXBuilder() (*xdrbuild.Builder, error)
	// CacheReceivers should be used to tell connector to cache created receivers. By default connector should use cache
	CacheReceivers(cache bool) Connector
}

type Singler interface {
	Get(dst interface{}, query ...types.QueryParamer) error
	// ValidateResponses should be used to tell singler (not) to validate responses against jsonschemas. By default singler should validate responses
	ValidateResponses(shouldValidate bool) Singler
}

type Nexter interface {
	Next(dst interface{}, query ...types.QueryParamer) error
	// WithPathParams sets additional path params for endpoint where we expect to receive
	// ResourceListResponse but also need to pass id or other non-query parameter in URL.
	// E.g. `/accounts/{id}/signers` which returns SignerListResponse and gets AccountID as path parameter
	WithPathParams(pathParams types.PathParamer) Nexter
	// ValidateResponses should be used to tell nexter (not) to validate responses against jsonschemas. By default nexter should validate responses
	ValidateResponses(shouldValidate bool) Nexter
}
