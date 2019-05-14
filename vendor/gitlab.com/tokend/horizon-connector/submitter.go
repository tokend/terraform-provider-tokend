package horizon

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal/resources"
	"gitlab.com/tokend/horizon-connector/internal/responses"
)

var (
	ErrSubmitTimeout              = errors.New("submit timed out")
	ErrSubmitInternal             = errors.New("internal submit error")
	ErrSubmitRejected             = errors.New("transaction rejected")
	ErrSubmitMalformed            = errors.New("transaction malformed")
	ErrSubmitFailed               = errors.New("Transaction failed.")
	ErrSubmitUnexpectedStatusCode = errors.New("Unexpected unsuccessful status code.")
)

type Submitter struct {
	client *Client
}

type SubmitResult struct {
	Err         error
	RawResponse []byte
	TXCode      string
	OpCodes     []string
	ResultXDR   string
}

func (r SubmitResult) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"err":        r.Err.Error(),
		"raw":        string(r.RawResponse),
		"tx_code":    r.TXCode,
		"op_codes":   r.OpCodes,
		"result_xdr": r.ResultXDR,
	}
}

func (s *Submitter) Submit(ctx context.Context, envelope string) SubmitResult {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&resources.TransactionSubmit{
		Transaction: envelope,
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal request"))
	}

	response, err := s.client.Post("/transactions", &buf)
	if err == nil {
		var success responses.TransactionSuccess
		if err := json.Unmarshal(response, &success); err != nil {
			// oops, tx was successful but we failed to unmarshal response.
			// let's ignore error and hope nothing will break
			// TODO debug log
		}
		// successful submission
		return SubmitResult{
			RawResponse: response,
			ResultXDR:   success.Result,
		}
	}

	cerr := errors.Cause(err).(Error)
	result := SubmitResult{
		RawResponse: cerr.Body(),
	}

	// go through known response codes and try to build meaningful result
	switch cerr.Status() {
	case http.StatusGatewayTimeout: // timeout
		result.Err = ErrSubmitTimeout
	case http.StatusBadRequest: // rejected or malformed
		// check which error it was exactly, might be useful for consumer
		var response responses.TransactionBadRequest
		if err := json.Unmarshal(result.RawResponse, &response); err != nil {
			panic(errors.Wrap(err, "failed to unmarshal horizon response"))
		}
		switch response.Type {
		case "transaction_malformed":
			result.Err = ErrSubmitMalformed
		case "transaction_failed":
			result.Err = ErrSubmitRejected
			result.TXCode = response.Extras.ResultCodes.Transaction
			result.OpCodes = response.Extras.ResultCodes.Operations
			result.ResultXDR = response.Extras.ResultXDR
		default:
			panic("unknown reject type")
		}
	case http.StatusInternalServerError: // internal error
		result.Err = ErrSubmitInternal
	default:
		// TODO poke someone who touched horizon
		panic("unexpected submission result")
	}
	return result
}

type SubmitResponseDetails struct {
	StatusCode  int
	RawResponse []byte
	TXCode      string
	OpCodes     []string
}

func (d SubmitResponseDetails) GetLoganFields() map[string]interface{} {
	result := map[string]interface{}{
		"status_code":  d.StatusCode,
		"raw_response": string(d.RawResponse),
	}

	if d.TXCode != "" {
		result["tx_code"] = d.TXCode
		result["op_codes"] = d.OpCodes
	}
	if len(d.OpCodes) > 0 {
		result["op_codes"] = d.OpCodes
	}

	return result
}

// SubmitE submits txEnvelope checking common unsuccessful status codes and forming
// appropriate errors for status codes.
//
// Returned SubmitResponseDetails always contain info (even with a non-nil error),
// however all the details (status-code, raw-response) are duplicated in the returned error fields
// for convenient error logging.
//
// For transaction_failed response in BadRequest(400) response, SubmitE also parses
// TXCode and OpCodes into SubmitResponseDetails structure.
func (s *Submitter) SubmitE(txEnvelope string) (SubmitResponseDetails, error) {
	req := resources.TransactionSubmit{
		Transaction: txEnvelope,
	}

	statusCode, respBB, err := s.client.PostJSON("/transactions", req)
	if err != nil {
		return SubmitResponseDetails{}, errors.Wrap(err, "Failed to send POST request via Client")
	}
	details := SubmitResponseDetails{
		StatusCode:  statusCode,
		RawResponse: respBB,
	}

	if isStatusCodeSuccessful(statusCode) {
		return details, nil
	}

	// go through known unsuccessful response status codes and try to build meaningful result
	switch statusCode {
	case http.StatusGatewayTimeout:
		return details, errors.From(ErrSubmitTimeout, details.GetLoganFields())
	case http.StatusBadRequest: // rejected or malformed
		// check which error it was exactly, might be useful for consumer
		var response responses.TransactionBadRequest
		if err := json.Unmarshal(respBB, &response); err != nil {
			return details, errors.Wrap(err, "Failed to unmarshal BadRequest Horizon response bytes into TransactionBadRequest struct", details.GetLoganFields())
		}
		details.TXCode = response.Extras.ResultCodes.Transaction
		details.OpCodes = response.Extras.ResultCodes.Operations

		switch response.Type {
		case "transaction_malformed":
			return details, errors.From(ErrSubmitMalformed, details.GetLoganFields())
		case "transaction_failed":
			return details, errors.From(ErrSubmitFailed, details.GetLoganFields())
		default:
			return details, errors.From(errors.New("Unknown reject type."), logan.F{"reject_type": response.Type}.Merge(details.GetLoganFields()))
		}
	case http.StatusInternalServerError: // internal error
		return details, errors.From(ErrSubmitInternal, details.GetLoganFields())
	default:
		// Normally must never happen. Looks like somebody changed Horizon.
		return details, errors.From(ErrSubmitUnexpectedStatusCode, details.GetLoganFields())
	}

	return details, nil
}
