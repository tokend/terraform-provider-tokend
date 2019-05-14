package regources

type Info struct {
	Passphrase         string `json:"network_passphrase"`
	MasterAccountID    string `json:"admin_account_id"`
	TXExpirationPeriod int64  `json:"tx_expiration_period"`
}

func (c *Info) GetMasterAccountID() string {
	return c.MasterAccountID
}

func (c *Info) GetPassphrase() string {
	return c.Passphrase
}

func (c *Info) GetTXExpire() int64 {
	return c.TXExpirationPeriod
}
