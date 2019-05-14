package responses

import "gitlab.com/tokend/horizon-connector/internal/resources"

type Balances struct {
	Embedded struct {
		Records []resources.ChoppedBalance `json:"records"`
	} `json:"_embedded"`
}
