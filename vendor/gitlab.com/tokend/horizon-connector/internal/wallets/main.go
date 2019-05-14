package wallets

import (
	"encoding/json"
	"fmt"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/horizon-connector/internal/resources"
	"gitlab.com/tokend/horizon-connector/internal/responses"
	"gitlab.com/tokend/horizon-connector/types"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

//todo set urlval
func (q *Q) Filter(opts *types.GetOpts) ([]resources.Wallet, int32, error) {
	endpoint := opts.Encode()

	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Failed to get wallets")
	}

	var wallets responses.Wallets

	if err := json.Unmarshal(response, &wallets); err != nil {
		return nil, 0, errors.Wrap(err, "Failed to unmarshal wallets")
	}

	page, err := wallets.GetPage()
	if err != nil {
		return nil, 0, errors.Wrap(err, "Failed to get wallets page")
	}

	return wallets.Data, page, nil
}

func (q *Q) Delete(walletID string) error {
	endpoint := fmt.Sprintf("/wallets/%s", walletID)
	_, err := q.client.Delete(endpoint)
	if err != nil {
		return errors.Wrap(err, "Failed wallet delete request")
	}

	return nil
}

func (q *Q) WalletExist(walletID string) (bool, error) {
	endpoint := fmt.Sprintf("/wallets/%s", walletID)
	respBody, err := q.client.Get(endpoint)
	if err != nil {
		return false, errors.Wrap(err, "Failed to get wallet")
	}
	if respBody != nil {
		return true, nil
	}
	return false, nil
}

func (q *Q) WalletByEmail(email string) (*resources.Wallet, error) {
	endpoint := fmt.Sprintf("/wallets?email=%s", email)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil,  errors.Wrap(err, "Failed to get wallets")
	}
	var wallets responses.Wallets
	if err := json.Unmarshal(response, &wallets); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	if len(wallets.Data) == 0 {
		return nil, nil
	}
	return &wallets.Data[0], nil
}
