/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type ConvertedBalanceStateAttributes struct {
	ConvertedAmounts BalanceStateAttributeAmounts `json:"converted_amounts"`
	InitialAmounts   BalanceStateAttributeAmounts `json:"initial_amounts"`
	// if `false` - the price doesn't exist and conversion is not possible
	IsConverted bool `json:"is_converted"`
	// Price that was used to convert amounts
	Price Amount `json:"price"`
}
