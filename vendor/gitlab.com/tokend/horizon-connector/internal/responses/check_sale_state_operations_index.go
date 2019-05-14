package responses

import "gitlab.com/tokend/horizon-connector/internal/resources/operations"

type CheckSaleStateOperationsIndex struct {
	Embedded struct {
		Records []operations.CheckSaleState `json:"records"`
	} `json:"_embedded"`
}
