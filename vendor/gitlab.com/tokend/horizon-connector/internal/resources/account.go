package resources

import "gitlab.com/tokend/regources"

type Account struct {
	AccountID              string                              `json:"account_id"`
	IsBlocked              bool                                `json:"is_blocked"`
	AccountTypeI           int32                               `json:"account_type_i"`
	AccountType            string                              `json:"account_type"`
	ExternalSystemAccounts []regources.ExternalSystemAccountID `json:"external_system_accounts"`
	KYC                    AccountKYC                          `json:"account_kyc"`
	Referrer               string                              `json:"referrer"`
	Balances               []Balance                           `json:"balances"`
}

func (a Account) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"account_id":               a.AccountID,
		"is_blocked":               a.IsBlocked,
		"account_type_i":           a.AccountTypeI,
		"account_type":             a.AccountType,
		"external_system_accounts": a.ExternalSystemAccounts,
		"kyc":      a.KYC,
		"referrer": a.Referrer,
	}
}

type AccountKYC struct {
	Data *regources.KYCData `json:"KYCData"`
}

func (k AccountKYC) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"data": k.Data,
	}
}
