package listener

import "gitlab.com/tokend/horizon-connector/internal/resources"

type TXPacket struct {
	body *resources.TransactionEvent
	err  error
}

func (e *TXPacket) Unwrap() (*resources.TransactionEvent, error) {
	return e.body, e.err
}

func (e *TXPacket) GetLoganFields() map[string]interface{} {
	return map[string]interface{} {
		"body": e.body,
		"err":  e.err,
	}
}
