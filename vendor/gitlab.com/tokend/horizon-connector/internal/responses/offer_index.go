package responses

import "gitlab.com/tokend/regources"

type OfferIndex struct {
	Embedded struct {
		Meta    regources.PageMeta      `json:"meta"`
		Records []regources.Offer `json:"records"`
	} `json:"_embedded"`
}
