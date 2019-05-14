package listener

import (
	"gitlab.com/tokend/regources"
)

type ReviewableRequestEvent struct {
	body *regources.ReviewableRequest
	err  error
}

func (e *ReviewableRequestEvent) Unwrap() (*regources.ReviewableRequest, error) {
	return e.body, e.err
}

func (e *ReviewableRequestEvent) GetLoganFields() map[string]interface{} {
	return map[string]interface{} {
		"body": e.body,
		"err":  e.err,
	}
}
