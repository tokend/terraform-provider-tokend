package xdrbuild

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type ClosePoll struct {
	PollID  uint64
	Details json.Marshaler
	Result  xdr.PollResult
}

func (op *ClosePoll) XDR() (*xdr.Operation, error) {
	details, err := op.Details.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal details")
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManagePoll,
			ManagePollOp: &xdr.ManagePollOp{
				PollId: xdr.Uint64(op.PollID),
				Data: xdr.ManagePollOpData{
					Action: xdr.ManagePollActionClose,
					ClosePollData: &xdr.ClosePollData{
						Details: xdr.Longstring(details),
						Result:  op.Result,
					},
				},
			},
		},
	}, nil
}
