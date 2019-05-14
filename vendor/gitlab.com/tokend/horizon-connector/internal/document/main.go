package document

import (
	"encoding/json"
	"fmt"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/horizon-connector/internal/resources"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

// Document obtains a single Document by its ID.
// If Document doesn't exist the link will still be returned - existence of the document is not checked.
func (q *Q) Document(docID string) (*resources.Document, error) {
	respBB, err := q.client.Get(fmt.Sprintf("/documents/%s", docID))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to send GET request")
	}

	if respBB == nil {
		// something must be returned, so there is some problem with request
		return nil, errors.New("failed to get response")
	}

	document := resources.Document{}
	if err := json.Unmarshal(respBB, &document); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response bytes", logan.F{
			"raw_response": string(respBB),
		})
	}

	return &document, nil
}
