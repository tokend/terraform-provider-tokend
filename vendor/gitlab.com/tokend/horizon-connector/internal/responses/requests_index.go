package responses

import (
	"gitlab.com/tokend/regources"
)

type RequestsIndex struct {
	Embedded struct {
		Records []regources.ReviewableRequest `json:"records"`
	} `json:"_embedded"`
}
