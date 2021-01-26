package core

import "fmt"

type infoResponse struct {
	Info Info `json:"info"`
}

// Info response for /info request
type Info struct {
	// CoreVersion version of the core
	CoreVersion string `json:"build"`
	// NetworkPassphrase passphrase of the network
	NetworkPassphrase string `json:"network"`
	// MasterExchangeName name of the exchange managed by master key
	MasterExchangeName string `json:"base_exchange_name"`
	// TxExpirationPeriod max allowed period for tx time bounds max
	TxExpirationPeriod int64 `json:"tx_expiration_period"`
	// WithdrawalDetailsMaxLength max length of details field for withdrawal operation
	WithdrawalDetailsMaxLength int64 `json:"withdrawal_details_max_length"`
	// DemurragePeriod frequency of demurrage been charged
	DemurragePeriod int64 `json:"demurrage_period"`
	// Array of the base assets
	BaseAssets []string `json:"base_assets"`

	// MasterAccountID account ID of master
	MasterAccountID string `json:"master_account_id"`
	// CommissionAccountID account ID of commission account
	CommissionAccountID string `json:"commission_account_id"`
	// OperationalAccountID account ID of operational account
	OperationalAccountID string `json:"operational_account_id"`
	// StorageFeeManageAccountID account ID of account which stores fees charged by demurrage op
	StorageFeeManageAccountID string `json:"storage_fee_manager_account_id"`
}

func (i *Info) validate() error {
	errorProvider := func(name string) error {
		return fmt.Errorf("%s must not be empty. Please check connection with stellar-core", name)
	}
	if i.NetworkPassphrase == "" {
		return errorProvider("NetworkPassphrase")
	}

	if i.MasterAccountID == "" {
		return errorProvider("MasterAccountID")
	}

	if i.CommissionAccountID == "" {
		return errorProvider("CommissionAccountID")
	}

	if i.TxExpirationPeriod <= 0 {
		return errorProvider("TxExpirationPeriod")
	}

	if i.DemurragePeriod <= 0 {
		return errorProvider("DemurragePeriod")
	}

	if i.WithdrawalDetailsMaxLength <= 0 {
		return errorProvider("WithdrawalDetailsMaxLength")
	}
	return nil
}

