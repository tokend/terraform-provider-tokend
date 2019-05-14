package responses

import "gitlab.com/tokend/horizon-connector/internal/resources/operations"

type PaymentV2OperationIndex struct {
	Embedded struct {
		Records []operations.PaymentV2 `json:"records"`
	} `json:"_embedded"`
}
