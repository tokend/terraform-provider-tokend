package balance

import (
	"encoding/json"
	"fmt"

	"gitlab.com/tokend/horizon-connector/internal/responses"

	"github.com/pkg/errors"
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

func (q *Q) AccountID(balanceID string) (*string, error) {
	endpoint := fmt.Sprintf("/balances/%s/account", balanceID)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP GET request failed")
	}

	if response == nil {
		return nil, nil
	}

	// actually it's different struct (HistoryAccount) but it works since we only need account_id
	var account resources.Account
	if err := json.Unmarshal(response, &account); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response bytes into Account struct")
	}
	return &account.AccountID, nil
}

func (q *Q) BalancesByAsset(asset, cursor string) ([]resources.ChoppedBalance, error) {
	endpoint := fmt.Sprintf("/balances?asset=%s&cursor=%s", asset, cursor)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP GET request failed")
	}

	if response == nil {
		panic(errors.New("unexpected nil response"))
	}

	var balances responses.Balances
	if err := json.Unmarshal(response, &balances); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response bytes into Balances struct")
	}
	return balances.Embedded.Records, nil
}

func (q *Q) AssetByBalanceID(balanceID string) (*string, error) {
	endpoint := fmt.Sprintf("/balances/%s/asset", balanceID)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP GET request failed")
	}

	if response == nil {
		return nil, nil
	}

	var asset struct {
		Code string `json:"code"`
	}
	if err := json.Unmarshal(response, &asset); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response bytes into asset struct")
	}
	return &asset.Code, nil
}
