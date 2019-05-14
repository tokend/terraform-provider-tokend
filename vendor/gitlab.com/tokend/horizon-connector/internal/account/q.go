package account

import (
	"encoding/json"
	"fmt"
	"gitlab.com/tokend/horizon-connector/internal/resources/operations"
	"gitlab.com/tokend/horizon-connector/types"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	goresources "gitlab.com/tokend/go/resources"
	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/horizon-connector/internal/resources"
	"gitlab.com/tokend/horizon-connector/internal/responses"
	"gitlab.com/tokend/regources"
)

var (
	ErrNoSigner      = errors.New("No such signer")
	ErrNotEnoughType = errors.New("Not enough types")
	ErrNoBalance     = errors.New("no such balance")
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

//func (q *Q) IsSigner(master string, signer string, signerType ...xdr.SignerType) error {
//	signers, err := q.Signers(master)
//	if err != nil {
//		return errors.Wrap(err, "Failed to load master signers")
//	}
//
//	isAllowedSigner := false
//	var notEnoughTypes []xdr.SignerType
//	for _, s := range signers {
//		if signer == s.AccountID {
//			isAllowedSigner = true
//		}
//
//		for _, t := range signerType {
//			if s.SignerType&int(t) == 0 {
//				notEnoughTypes = append(notEnoughTypes, t)
//			}
//		}
//	}
//
//	if !isAllowedSigner {
//		return errors.Wrap(ErrNoSigner, "Unknown signer", logan.F{"address": signer})
//	}
//
//	if len(notEnoughTypes) != 0 {
//		return errors.Wrap(ErrNotEnoughType, "Signer type not valid", logan.F{"types": notEnoughTypes})
//	}
//
//	return nil
//}

func (q *Q) Signers(address string) ([]goresources.Signer, error) {
	endpoint := fmt.Sprintf("/accounts/%s/signers", address)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var result responses.AccountSigners
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	return result.Signers, nil
}

// ByBalance return account address by balance
// DEPRECATED: use AccountID() from Balances Q instead
func (q *Q) ByBalance(balanceID string) (*string, error) {
	endpoint := fmt.Sprintf("/balances/%s/account", balanceID)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	// actually it's different struct (HistoryAccount) but it works since we only need account_id
	var account resources.Account
	if err := json.Unmarshal(response, &account); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	return &account.AccountID, nil
}

func (q *Q) ByAddress(address string) (*resources.Account, error) {
	endpoint := fmt.Sprintf("/accounts/%s", address)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var account resources.Account
	if err := json.Unmarshal(response, &account); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	return &account, nil
}

// CurrentBalanceIn return account's balance in provided asset
// ErrNoBalance if balance does not exist
func (q *Q) CurrentBalanceIn(address, asset string) (result resources.Balance, err error) {
	account, err := q.ByAddress(address)
	if err != nil {
		return result, errors.Wrap(err, "failed to get account")
	}
	for _, balance := range account.Balances {
		if balance.Asset == asset {
			return balance, nil
		}
	}
	return result, ErrNoBalance
}

// CurrentExternalBindingData will return (nil, nil) if account external binding does not exist
func (q *Q) CurrentExternalBindingData(address string, externalSystem int32) (*string, error) {
	account, err := q.ByAddress(address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account")
	}

	for _, system := range account.ExternalSystemAccounts {
		if system.Type.Value == externalSystem {
			return &system.Data, nil
		}
	}

	return nil, nil
}

func (q *Q) Balances(address string) ([]resources.Balance, error) {
	endpoint := fmt.Sprintf("/accounts/%s/balances", address)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var result responses.AccountBalances
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	return result, nil
}

// Offers requests Offers by the Account.
// All parameters except `address` are optional - use default values for parameters you don't need to provide.
func (q *Q) Offers(address, baseAsset, quoteAsset string, isBuy *bool, offerID string, orderBookID *uint64) ([]regources.Offer, error) {
	if (baseAsset == "") != (quoteAsset == "") {
		return nil, errors.New("base and quote assets must be both set or both not set")
	}

	endpoint := fmt.Sprintf("/accounts/%s/offers?base_asset=%s&quote_asset=%s&offer_id=%s",
		address, baseAsset, quoteAsset, offerID)
	if isBuy != nil {
		endpoint = fmt.Sprintf("%s&is_buy=%v", endpoint, *isBuy)
	}
	if orderBookID != nil {
		endpoint = fmt.Sprintf("%s&order_book_id=%d", endpoint, *orderBookID)
	}

	respBB, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if respBB == nil {
		return nil, nil
	}

	var result responses.OfferIndex
	if err := json.Unmarshal(respBB, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	return result.Embedded.Records, nil
}

func (q *Q) References(address string) ([]resources.Reference, error) {
	respBB, err := q.client.Get(fmt.Sprintf("/accounts/%s/references", address))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to send GET request")
	}

	if respBB == nil {
		// No References
		return nil, nil
	}

	var response struct {
		Data []resources.Reference `json:"data"`
	}
	if err := json.Unmarshal(respBB, &response); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal response bytes", logan.F{
			"raw_response": string(respBB),
		})
	}

	return response.Data, nil
}

func (q *Q) ReferenceExist(address, reference string) (bool, error) {
	refs, err := q.References(address)
	if err != nil {
		return false, errors.Wrap(err, "failed to get references")
	}
	for _, ref := range refs {
		if ref.Reference == reference {
			return true, nil
		}
	}
	return false, nil
}

func (q *Q) Payments(accountID string, opts types.PaymentsOpts) ([]operations.PaymentV2, error) {
	endpoint := fmt.Sprintf("/accounts/%s/payments?%s", accountID, opts.Encode())

	raw, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if raw == nil {
		return nil, nil
	}

	var result responses.PaymentV2OperationIndex
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	return result.Embedded.Records, nil
}
