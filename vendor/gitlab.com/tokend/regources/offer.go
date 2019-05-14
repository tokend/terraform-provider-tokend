package regources

import (
	"gitlab.com/distributed_lab/logan/v3/fields"
	"strconv"
)

type Offer struct {
	PT             string `json:"paging_token"`
	OwnerID        string `json:"owner_id"`
	OfferID        uint64 `json:"offer_id"`
	OrderBookID    uint64 `json:"order_book_id"`
	BaseBalanceID  string `json:"base_balance_id"`
	QuoteBalanceID string `json:"quote_balance_id"`
	Fee            Amount `json:"fee"`
	OfferData
}

// PagingToken implementation for hal.Pageable in Horizon
func (o Offer) PagingToken() string {
	return strconv.FormatUint(o.OfferID, 10)
}

func (o Offer) GetLoganFields() map[string]interface{} {
	return fields.Merge(map[string]interface{}{
		"paging_token":     o.PT,
		"owner_id":         o.OwnerID,
		"offer_id":         o.OfferID,
		"order_book_id":    o.OrderBookID,
		"base_balance_id":  o.BaseBalanceID,
		"quote_balance_id": o.QuoteBalanceID,
		"fee":              o.Fee,
	}, o.OfferData.GetLoganFields())
}

type OfferData struct {
	BaseAssetCode  string `json:"base_asset_code"`
	QuoteAssetCode string `json:"quote_asset_code"`
	IsBuy          bool   `json:"is_buy"`
	BaseAmount     Amount `json:"base_amount"`
	QuoteAmount    Amount `json:"quote_amount"`
	Price          Amount `json:"price"`
	CreatedAt      Time   `json:"created_at"`
}

func (d OfferData) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"base_asset_code":  d.BaseAssetCode,
		"quote_asset_code": d.QuoteAssetCode,
		"is_buy":           d.IsBuy,
		"base_amount":      d.BaseAmount,
		"quote_amount":     d.QuoteAmount,
		"price":            d.Price,
		"created_at":       d.CreatedAt,
	}
}
