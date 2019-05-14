package regources

import "time"

type PageMeta struct {
	LatestLedger LedgerMeta `json:"latest_ledger"`
}

type LedgerMeta struct {
	ClosedAt time.Time `json:"closed_at"`
}
