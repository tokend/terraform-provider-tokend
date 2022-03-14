package submit

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

func newTxFailure(response txFailureResponse) TxFailure {
	failure := TxFailure{
		error:    errors.New(response.Errors[0].Detail),
		Response: response,
	}

	if response.Errors[0].Meta != nil {
		failure.ResultXDR = response.Errors[0].Meta.ResultXDR
		failure.OperationResultCodes = response.Errors[0].Meta.ResultCodes.OperationCodes
		failure.TransactionResultCode = response.Errors[0].Meta.ResultCodes.TransactionCode
	}

	return failure
}

//TxFailure is a helper structure to represent transction submission failure details
type TxFailure struct {
	error
	Response              interface{}
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

func (f TxFailure) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"response":                f.Response,
		"result_xdr":              f.ResultXDR,
		"transaction_result_code": f.TransactionResultCode,
		"operation_result_codes":  f.OperationResultCodes,
	}
}
