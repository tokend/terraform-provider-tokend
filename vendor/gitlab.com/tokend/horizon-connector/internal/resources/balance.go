package resources

import (
	"gitlab.com/tokend/regources"
)

type Balance struct {
	Asset     string           `json:"asset"`
	BalanceID string           `json:"balance_id"`
	AccountID string           `json:"account_id"`
	Balance   regources.Amount `json:"balance"`
	Locked    regources.Amount `json:"locked"`
}

type ChoppedBalance struct {
	ID        string `json:"id"`
	Asset     string `json:"asset"`
	BalanceID string `json:"balance_id"`
	AccountID string `json:"account_id"`
}
