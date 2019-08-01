package horizon

import (
	"context"
	"encoding/json"
	"gitlab.com/tokend/go/xdr"
	regources "gitlab.com/tokend/regources/generated"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	ErrRequestClosed              = errors.New("request closed")
	ErrSubmitTimeout              = errors.New("submit timed out")
	ErrSubmitInternal             = errors.New("internal submit error")
)

type Submitter struct {
	client *Client
}

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
	_, response, err := s.client.PostJSONWithContext("/v3/transactions", &regources.SubmitTransactionBody{
		Tx: envelope,
	}, ctx)
	if err == nil {
		var success regources.TransactionResponse
		if err := json.Unmarshal(response, &success); err != nil {
			return SubmitResult{
				Err:         err,
				RawResponse: response,
			}
		}
		// successful submission
		return SubmitResult{
			RawResponse: response,
			ResultXDR:   success.Data.Attributes.ResultXdr,
		}
	}

	if isContextCanceled(ctx) {
		return SubmitResult{
			Err: ErrRequestClosed,
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
		var failureResp txFailureResponse
		if err := json.Unmarshal(response, &failureResp); err != nil {
			panic(errors.Wrap(err, "failed to unmarshal horizon response"))
		}
		failure := newTxFailure(failureResp)
		result.Err = failure.error
		result.OpCodes = failure.OperationResultCodes
		result.RawResponse = response
		result.ResultXDR = failure.ResultXDR
		result.TXCode = failure.TransactionResultCode
	case http.StatusInternalServerError: // internal error
		result.Err = ErrSubmitInternal
	default:
		// TODO poke someone who touched horizon
		panic("unexpected submission result")
	}
	return result
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

func isContextCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
