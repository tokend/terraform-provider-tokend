package resources

import "time"

type Wallet struct {
	ID         string           `json:"id"`
	Attributes WalletAttributes `json:"attributes"`
}

type WalletAttributes struct {
	LastSentAt *time.Time `json:"last_sent_at,omitempty"`
	Email      string     `json:"email,omitempty"`
	AccountID  string     `json:"account_id"`
}
