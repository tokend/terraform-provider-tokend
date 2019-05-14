package regources

import (
	"time"
)

// AtomicSwapBid represents singe atomic swap entry
type AtomicSwapBid struct {
	ID              string                 `json:"id"`
	PT              string                 `json:"paging_token"`
	OwnerID         string                 `json:"owner_id"`
	BaseAsset       string                 `json:"base_asset"`
	BaseBalanceID   string                 `json:"base_balance_id"`
	AvailableAmount Amount                 `json:"available_amount"`
	LockedAmount    Amount                 `json:"locked_amount"`
	CreatedAt       time.Time              `json:"created_at"`
	IsCanceled      bool                   `json:"is_canceled"`
	Details         map[string]interface{} `json:"details"`
	QuoteAssets     []AssetPrice           `json:"quote_assets"`
}

func (b AtomicSwapBid) PagingToken() string {
	return b.PT
}

type AssetPrice struct {
	Asset string `json:"asset"`
	Price Amount `json:"price"`
}
