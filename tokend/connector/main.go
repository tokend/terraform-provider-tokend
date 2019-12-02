package connector

import (
	"context"
	"encoding/json"
	"gitlab.com/distributed_lab/json-api-connector/horizon"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/regources"
	"net/http"
	"net/url"
)

type Submitter struct {
	connector *horizon.Connector
}

func NewSubmitter(connector *horizon.Connector) Submitter {
	return Submitter{connector: connector}
}

var (
	//ErrSubmitTimeout indicates that transaction submission has timed out
	ErrSubmitTimeout = errors.New("submit timed out")
	//ErrSubmitInternal indicates that transaction submission has failed with internal error
	ErrSubmitInternal = errors.New("internal submit error")
	//ErrSubmitUnexpectedStatusCode indicates that transaction submission has failed with unexpected status code
	ErrSubmitUnexpectedStatusCode = errors.New("unexpected unsuccessful status code")
)

//TxFailure is a helper structure to represent transction submission failure details
type TxFailure struct {
	error
	ResultXDR             string
	TransactionResultCode string
	OperationResultCodes  []string
}

type txFailureResponse struct {
	Errors []struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Status string `json:"status"`
		Meta   *struct {
			Envelope     string                `json:"envelope"`
			ResultXDR    string                `json:"result_xdr"`
			ParsedResult xdr.TransactionResult `json:"parsed_result"`
			ResultCodes  struct {
				TransactionCode string   `json:"transaction"`
				OperationCodes  []string `json:"operations,omitempty"`
				Messages        []string `json:"messages"`
			} `json:"result_codes"`
		} `json:"meta,omitempty"`
	} `json:"errors"`
}

func (s *Submitter) Submit(ctx context.Context, envelope string, waitIngest bool) (*regources.TransactionResponse, error) {
	u, _ := url.Parse("/v3/transactions")
	status, respBB, err := s.connector.Submit(ctx, u, envelope, waitIngest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit tx to horizon")
	}

	if isStatusCodeSuccessful(status) && err == nil {
		var success regources.TransactionResponse
		if err := json.Unmarshal(respBB, &success); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal transaction response")
		}
		return &success, nil
	}

	// go through known response codes and try to build meaningful result
	switch status {
	case http.StatusGatewayTimeout: // timeout
		return nil, ErrSubmitTimeout
	case http.StatusBadRequest: // rejected or malformed
		// check which error it was exactly, might be useful for consumer
		var failureResp txFailureResponse
		if err := json.Unmarshal(respBB, &failureResp); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal horizon response")
		}
		return nil, newTxFailure(failureResp)
	case http.StatusInternalServerError: // internal error
		return nil, ErrSubmitInternal
	default:
		return nil, ErrSubmitUnexpectedStatusCode
	}
}

func isStatusCodeSuccessful(code int) bool {
	return code >= 200 && code < 300
}

func newTxFailure(response txFailureResponse) TxFailure {
	failure := TxFailure{
		error: errors.New(response.Errors[0].Detail),
	}

	if response.Errors[0].Meta != nil {
		failure.ResultXDR = response.Errors[0].Meta.ResultXDR
		failure.OperationResultCodes = response.Errors[0].Meta.ResultCodes.OperationCodes
		failure.TransactionResultCode = response.Errors[0].Meta.ResultCodes.TransactionCode
	}

	return failure
}
