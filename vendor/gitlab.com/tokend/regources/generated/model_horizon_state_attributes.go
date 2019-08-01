/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import (
	"time"
)

type HorizonStateAttributes struct {
	Core LedgerInfo `json:"core"`
	// version of the TokenD core to which this horizon is connected
	CoreVersion string `json:"core_version"`
	// current horizon time
	CurrentTime time.Time `json:"current_time"`
	// current horizon time in unix timestamp
	CurrentTimeUnix int64 `json:"current_time_unix"`
	// Defines user friendly name of the network
	EnvironmentName string     `json:"environment_name"`
	History         LedgerInfo `json:"history"`
	HistoryV2       LedgerInfo `json:"history_v2"`
	// Defines public account id of the master account for this network. All admins are signers of this account
	MasterAccountId string `json:"master_account_id"`
	// TokenD network identificator. Shows in which network horizon is working.
	NetworkPassphrase string `json:"network_passphrase"`
	// default asset precision in system
	Precision int64 `json:"precision"`
	// Defines max allowed duration between timeBounds.maxTime & close time
	TxExpirationPeriod int64 `json:"tx_expiration_period"`
	// revision of xdr used by Horizon
	XdrRevision string `json:"xdr_revision"`
}
