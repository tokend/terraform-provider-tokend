package xdr

import (
	"gitlab.com/tokend/go/amount"
)

func (b *BalanceEntry) CanLock(toLock Uint64) BalanceResult {
	if toLock > b.Amount {
		return BalanceResultUnderfunded
	}

	_, ok := amount.SafePositiveSum(int64(b.Locked), int64(toLock))
	if !ok {
		return BalanceResultLineFull
	}
	return BalanceResultSuccess
}
func (b *BalanceEntry) Lock(toLock Uint64) BalanceResult {
	if toLock > b.Amount {
		return BalanceResultUnderfunded
	}

	newLocked, ok := amount.SafePositiveSum(int64(b.Locked), int64(toLock))
	if !ok {
		return BalanceResultLineFull
	}

	b.Locked = Uint64(newLocked)
	b.Amount -= toLock
	return BalanceResultSuccess
}

func (b *BalanceEntry) CanUnlock(toUnlock Uint64) BalanceResult {
	if toUnlock > b.Locked {
		return BalanceResultUnderfunded
	}

	_, ok := amount.SafePositiveSum(int64(b.Amount), int64(toUnlock))
	if !ok {
		return BalanceResultLineFull
	}

	return BalanceResultSuccess
}

func (b *BalanceEntry) Unlock(toUnlock Uint64) BalanceResult {
	if toUnlock > b.Locked {
		return BalanceResultUnderfunded
	}

	newAmount, ok := amount.SafePositiveSum(int64(b.Amount), int64(toUnlock))
	if !ok {
		return BalanceResultLineFull
	}
	b.Locked -= toUnlock
	b.Amount = Uint64(newAmount)

	return BalanceResultSuccess
}

func (b *BalanceEntry) CanCharge(toCharge Uint64) BalanceResult {
	if toCharge > b.Amount {
		return BalanceResultUnderfunded
	}

	return BalanceResultSuccess
}

func (b *BalanceEntry) Charge(toCharge Uint64) BalanceResult {
	if toCharge > b.Amount {
		return BalanceResultUnderfunded
	}

	b.Amount -= toCharge
	return BalanceResultSuccess
}

func (b *BalanceEntry) CanChargeFromLocked(toCharge Uint64) BalanceResult {
	if toCharge > b.Locked {
		return BalanceResultUnderfunded
	}

	return BalanceResultSuccess
}

func (b *BalanceEntry) ChargeFromLocked(toCharge Uint64) BalanceResult {
	if toCharge > b.Locked {
		return BalanceResultUnderfunded
	}

	b.Locked -= toCharge
	return BalanceResultSuccess
}

func (b *BalanceEntry) CanFund(toFund Uint64) BalanceResult {
	_, ok := amount.SafePositiveSum(int64(b.Amount), int64(toFund))
	if !ok {
		return BalanceResultLineFull
	}

	return BalanceResultSuccess
}

func (b *BalanceEntry) Fund(toFund Uint64) BalanceResult {
	newAmount, ok := amount.SafePositiveSum(int64(b.Amount), int64(toFund))
	if !ok {
		return BalanceResultLineFull
	}

	b.Amount = Uint64(newAmount)
	return BalanceResultSuccess
}

type BalanceResult int32

const (
	BalanceResultSuccess BalanceResult = iota
	BalanceResultLineFull
	BalanceResultUnderfunded
)
