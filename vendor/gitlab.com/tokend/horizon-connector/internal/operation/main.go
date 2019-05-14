package operation

import (
	"fmt"

	"encoding/json"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/horizon-connector/internal/resources/operations"
	"gitlab.com/tokend/horizon-connector/internal/responses"
	"gitlab.com/tokend/horizon-connector/types"
	"gitlab.com/tokend/regources"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

func (q *Q) AllRequests(cursor string) ([]regources.ReviewableRequest, error) {
	url := fmt.Sprintf("/requests?limit=200&cursor=%s", cursor)
	return q.getRequests(url)
}

// DEPRECATED: Instead use Requests() method providing specific ReviewableRequestType
//func (q *Q) WithdrawalRequests(cursor string) ([]regources.ReviewableRequest, error) {

// Requests obtains batch of Requests of the provided type from the provided cursor
// It differs from the AllRequests method, as the latter uses `/requests` path to obtain Requests.
func (q *Q) Requests(getParams, cursor string, reqType ReviewableRequestType) ([]regources.ReviewableRequest, error) {
	url := fmt.Sprintf("/request/%s?limit=200", string(reqType))

	if getParams != "" {
		url = fmt.Sprintf("%s&%s", url, getParams)
	}

	url = fmt.Sprintf("%s&cursor=%s", url, cursor)

	return q.getRequests(url)
}

func (q *Q) UpdateKycRequests(opts types.UpdateKYCRequestsOpts) ([]regources.ReviewableRequest, error) {
	url := fmt.Sprintf("request/%s?limit=200", string(KYCReviewableRequestType))
	url = fmt.Sprintf("%s&%s", url, opts.Encode())

	return q.getRequests(url)
}

func (q *Q) getRequests(url string) ([]regources.ReviewableRequest, error) {
	response, err := q.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "Request failed", logan.F{
			"request_url": url,
		})
	}

	var result responses.RequestsIndex
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response", logan.F{
			"raw_response": string(response),
			"request_url":  url,
		})
	}

	return result.Embedded.Records, nil
}

func (q *Q) GetRequestByID(requestID uint64) (*regources.ReviewableRequest, error) {
	response, err := q.client.Get(fmt.Sprintf("/requests/%d", requestID))
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		// No such Request
		return nil, nil
	}

	var result regources.ReviewableRequest
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	return &result, nil
}

func (q *Q) Operations(cursor string, operationType *xdr.OperationType) ([]byte, error) {
	var url string
	if operationType != nil {
		url = fmt.Sprintf("/operations?cursor=%s&operation_type=%d", cursor, *operationType)
	} else {
		url = fmt.Sprintf("/operations?cursor=%s", cursor)
	}

	response, err := q.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "request failed", logan.F{"request_url": url})
	}

	return response, err
}

func (q *Q) CheckSaleStateOperations(cursor string) ([]operations.CheckSaleState, error) {
	operationType := xdr.OperationTypeCheckSaleState
	response, err := q.Operations(cursor, &operationType)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get operations", logan.F{
			"operation_type": xdr.OperationTypeCheckSaleState.String(),
		})
	}

	var result responses.CheckSaleStateOperationsIndex
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response", logan.F{"response": string(response)})
	}

	return result.Embedded.Records, nil
}

func (q *Q) CreateKYCRequestOperations(cursor string) ([]operations.CreateKYCRequest, error) {
	//operationType := xdr.OperationTypeCreateKycRequest
	//response, err := q.Operations(cursor, &operationType)
	//if err != nil {
	//	return nil, errors.Wrap(err, "Failed to get operations", logan.F{
	//		"operation_type": xdr.OperationTypeCreateKycRequest.String(),
	//	})
	//}
	//
	//var result responses.CreateKYCRequestOperationIndex
	//if err := json.Unmarshal(response, &result); err != nil {
	//	return nil, errors.Wrap(err, "Failed to unmarshal response", logan.F{"response": string(response)})
	//}
	//
	//return result.Embedded.Records, nil
	panic("not implemented")
}

func (q *Q) PaymentV2Operations(cursor string) ([]operations.PaymentV2, error) {
	operationType := xdr.OperationTypePayment
	response, err := q.Operations(cursor, &operationType)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get operations", logan.F{
			"operation_type": xdr.OperationTypePayment.String(),
		})
	}

	var result responses.PaymentV2OperationIndex
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response", logan.F{"response": string(response)})
	}

	return result.Embedded.Records, nil
}

func (q *Q) ReviewRequestOperations(cursor string) ([]operations.ReviewRequest, error) {
	operationType := xdr.OperationTypeReviewRequest
	response, err := q.Operations(cursor, &operationType)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get operations", logan.F{
			"operation_type": xdr.OperationTypeReviewRequest.String(),
		})
	}

	var result responses.ReviewRequestOperationIndex
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response", logan.F{
			"response": string(response),
		})
	}

	return result.Embedded.Records, nil
}
