package responses

import (
	"gitlab.com/tokend/regources"
)

type TransactionIndex struct {
	Embedded struct {
		Meta    regources.PageMeta      `json:"meta"`
		Records []regources.Transaction `json:"records"`
	} `json:"_embedded"`
}
