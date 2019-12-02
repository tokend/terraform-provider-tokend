package xdr

import (
	"unicode"

	"gitlab.com/tokend/go/amount"
)

func (e AssetCode) IsValid() bool {
	if len(e) == 0 {
		return false
	}

	for _, let := range e {
		if !unicode.IsLetter(let) {

		}
	}

	return true
}

func (e AssetEntry) WillExceedMaxIssuanceAmount(toIssue Uint64) bool {
	newAmount, ok := amount.SafePositiveSum(int64(e.Issued), int64(toIssue))

	if !ok {
		return false
	}

	return Uint64(newAmount) > e.MaxIssuanceAmount
}

func (e *AssetEntry) AmountFits(am Uint64, maximumDigits Uint32) bool {
	return AmountFits(am, e.TrailingDigitsCount, maximumDigits)
}

func AmountFits(amount Uint64, actualDigits, maxDigits Uint32) bool {
	minAmount := 1
	exponent := maxDigits - actualDigits
	for exponent > 0 {
		minAmount *= 10
		exponent--
	}

	return amount%Uint64(minAmount) == 0
}

func (e *AssetEntry) TryIssue(toIssue Uint64) bool {
	if !e.CanAddAmount(toIssue) {
		return false
	}

	e.Issued += toIssue
	return true
}

func (e *AssetEntry) TryLock(toLock Uint64) bool {
	if !e.CanAddAmount(toLock) {
		return false
	}
	e.PendingIssuance += toLock
	return true
}

func (e *AssetEntry) TryUnIssue(toUnissue Uint64) bool {
	if toUnissue > e.Issued {
		return false
	}

	e.Issued -= toUnissue
	return true
}

func (e *AssetEntry) TryUnlock(toUnlock Uint64) bool {
	if toUnlock > e.PendingIssuance {
		return false
	}

	e.PendingIssuance -= toUnlock
	return true
}

func (e *AssetEntry) CanAddAmount(toAdd Uint64) bool {
	if int64(toAdd) <= 0 {
		return false
	}

	total, ok := amount.SafePositiveSumForSlice(int64(toAdd), int64(e.Issued), int64(e.PendingIssuance))
	if !ok {
		return false
	}

	return Uint64(total) <= e.MaxIssuanceAmount
}
