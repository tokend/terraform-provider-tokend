/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type AccountAttributes struct {
	// Indicates current kyc recovery status of account * 0 - none * 1 - initiated * 2 - pending * 3 - rejected * 4 - permanently_rejected
	KycRecoveryStatus *KYCRecoveryStatus `json:"kyc_recovery_status,omitempty"`
}
