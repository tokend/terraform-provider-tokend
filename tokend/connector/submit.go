package connector

import (
	"context"

	"gitlab.com/tokend/connectors/submit"
)

type HorizonSubmitter struct {
	submitter *submit.Submitter
}

func NewSubmitter(submitter *submit.Submitter) *HorizonSubmitter {
	return &HorizonSubmitter{
		submitter: submitter,
	}
}

func (s *HorizonSubmitter) Submitter() *HorizonSubmitter {
	return s
}

func (s *HorizonSubmitter) Submit(ctx context.Context, envelope string) Result {
	response, err := s.submitter.Submit(ctx, envelope, true, false)
	if err != nil {
		if txFailed, ok := err.(submit.TxFailure); ok {
			return Result{
				Err:       txFailed,
				ResultXDR: txFailed.ResultXDR,
				TXCode:    txFailed.TransactionResultCode,
				OpCodes:   txFailed.OperationResultCodes,
			}
		}

		return Result{
			Err: err,
		}
	}

	return Result{
		ResultXDR: response.Data.Attributes.ResultXdr,
	}
}

type Result struct {
	Err       error
	ResultXDR string
	TXCode    string
	OpCodes   []string
}
