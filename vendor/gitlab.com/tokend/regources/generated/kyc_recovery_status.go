/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

import "encoding/json"

type KYCRecoveryStatus int

const (
	KYCRecoveryStatusNone KYCRecoveryStatus = iota
	KYCRecoveryStatusInitiated
	KYCRecoveryStatusPending
	KYCRecoveryStatusRejected
	KYCRecoveryStatusPermanentlyRejected
)

var kycRecoveryStatusStr = map[KYCRecoveryStatus]string{
	KYCRecoveryStatusNone:                "none",
	KYCRecoveryStatusInitiated:           "initiated",
	KYCRecoveryStatusPending:             "pending",
	KYCRecoveryStatusRejected:            "rejected",
	KYCRecoveryStatusPermanentlyRejected: "permanently_rejected",
}

func (s KYCRecoveryStatus) String() string {
	return kycRecoveryStatusStr[s]
}

func (s KYCRecoveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(Flag{
		Name:  kycRecoveryStatusStr[s],
		Value: int32(s),
	})
}

func (s *KYCRecoveryStatus) UnmarshalJSON(b []byte) error {
	var res Flag
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}

	*s = KYCRecoveryStatus(res.Value)
	return nil
}

func (s KYCRecoveryStatus) IsFlag() bool {
	return true
}
