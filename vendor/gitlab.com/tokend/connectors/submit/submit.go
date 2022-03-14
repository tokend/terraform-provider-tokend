package submit

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/json-api-connector/cerrors"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
	regources "gitlab.com/tokend/regources/generated"
)

var (
	//ErrSubmitTimeout indicates that transaction submission has timed out
	ErrSubmitTimeout = errors.New("submit timed out")
	//ErrSubmitInternal indicates that transaction submission has failed with internal error
	ErrSubmitInternal = errors.New("internal submit error")
	//ErrSubmitUnexpectedStatusCode indicates that transaction submission has failed with unexpected status code
	ErrSubmitUnexpectedStatusCode = errors.New("unexpected unsuccessful status code")
)

// CreateTransaction creates new transaction with specified operations and source.
// Returns signed transaction envelope and its hash
// If signer argument is nil returns unsigned transaction and its hash
func (t *Submitter) CreateTransaction(source keypair.Address, signer keypair.Full, ops ...xdrbuild.Operation) (string, string, error) {
	builder, err := t.TXBuilder()
	if err != nil {
		return "", "", errors.Wrap(err, "failed to init builder")
	}

	tx := builder.Transaction(source)
	for _, op := range ops {
		tx.Op(op)
	}

	if signer != nil {
		tx.Sign(signer)
	}

	env, err := tx.Marshal()
	if err != nil {
		return "", "", errors.Wrap(err, "failed marshal tx")
	}

	hash, err := builder.TXHashHex(env)
	if err != nil {
		return "", "", errors.Wrap(err, "failed hashing tx")
	}

	return env, hash, nil
}

func (t *Submitter) Submit(ctx context.Context, envelope string, waitResult, waitIngest bool) (*regources.TransactionResponse, error) {
	if !waitResult && waitIngest {
		return nil, errors.New("not allowed to wait for ingest without waiting for result")
	}

	body := regources.SubmitTransactionBody{
		Tx:            envelope,
		WaitForIngest: &waitIngest,
		WaitForResult: &waitResult,
	}

	var success regources.TransactionResponse

	err := t.base.PostJSON(t.submissionUrl, body, ctx, &success)
	if err == nil {
		return &success, nil
	}

	cerr, ok := err.(cerrors.Error)
	if !ok {
		return nil, errors.Wrap(err, "failed to submit tx to horizon")
	}

	// go through known response codes and try to build meaningful result
	switch cerr.Status() {
	case http.StatusGatewayTimeout: // timeout
		return nil, ErrSubmitTimeout
	case http.StatusBadRequest: // rejected or malformed
		// check which error it was exactly, might be useful for consumer
		var failureResp txFailureResponse
		if err := json.Unmarshal(cerr.Body(), &failureResp); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal horizon response")
		}
		return nil, newTxFailure(failureResp)
	case http.StatusInternalServerError: // internal error
		return nil, ErrSubmitInternal
	default:
		return nil, errors.Wrap(err, ErrSubmitUnexpectedStatusCode.Error())
	}
}
