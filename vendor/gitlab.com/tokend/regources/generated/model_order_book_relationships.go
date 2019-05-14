/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type OrderBookRelationships struct {
	BaseAsset   *Relation           `json:"base_asset,omitempty"`
	BuyEntries  *RelationCollection `json:"buy_entries,omitempty"`
	QuoteAsset  *Relation           `json:"quote_asset,omitempty"`
	SellEntries *RelationCollection `json:"sell_entries,omitempty"`
}
