/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type BalanceStateAttributes struct {
	// Amount available to be spent by an account
	Available Amount `json:"available"`
	// Amount locked on the balance by some operation (ex: ManageOffer, CreateWithdrawalRequest, etc.) and controlled by the system.
	Locked Amount `json:"locked"`
}
