package responses

import "gitlab.com/tokend/horizon-connector/internal/resources/operations"

type CreateKYCRequestOperationIndex struct {
	Embedded struct {
		Records []operations.CreateKYCRequest `json:"records"`
	} `json:"_embedded"`
}
