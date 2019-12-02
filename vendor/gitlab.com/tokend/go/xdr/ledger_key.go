package xdr

import (
	"crypto/sha256"
	"fmt"
)

// LedgerKey implements the `Keyer` interface
func (key *LedgerKey) LedgerKey() LedgerKey {
	return *key
}

// Equals returns true if `other` is equivalent to `key`
func (key *LedgerKey) Equals(other LedgerKey) bool {
	if key.Type != other.Type {
		return false
	}

	switch key.Type {
	case LedgerEntryTypeAccount:
		l := key.MustAccount()
		r := other.MustAccount()
		return l.AccountId.Equals(r.AccountId)
	default:
		panic(fmt.Errorf("unknown ledger key type: %v", key.Type))
	}
}

// SetAccount mutates `key` such that it represents the identity of `account`
func (key *LedgerKey) SetAccount(account AccountId) error {
	data := LedgerKeyAccount{AccountId: account}
	nkey, err := NewLedgerKey(LedgerEntryTypeAccount, data)
	if err != nil {
		return err
	}

	*key = nkey
	return nil
}

func LedgerKeyFromKeyValue(key Longstring) LedgerKey {
	return LedgerKey{
		Type: LedgerEntryTypeKeyValue,
		KeyValue: &LedgerKeyKeyValue{
			Key: key,
		},
	}
}

type CacheKey struct {
	Type LedgerEntryType
	// Value is a value of switch arm filed (Account, Asset), but without pointer
	Value interface{}
}

func (key LedgerKey) CacheKey() CacheKey {
	result := CacheKey{
		Type: key.Type,
	}

	switch key.Type {
	case LedgerEntryTypeAccount:
		result.Value = key.MustAccount()
	case LedgerEntryTypeSigner:
		result.Value = key.MustSigner()
	case LedgerEntryTypeBalance:
		result.Value = key.MustBalance()
	case LedgerEntryTypeData:
		result.Value = key.MustData()
	case LedgerEntryTypeAsset:
		result.Value = key.MustAsset()
	case LedgerEntryTypeReferenceEntry:
		result.Value = key.MustReference()
	case LedgerEntryTypeReviewableRequest:
		result.Value = key.MustReviewableRequest()
	case LedgerEntryTypeAccountKyc:
		result.Value = key.MustAccountKyc()
	case LedgerEntryTypeKeyValue:
		result.Value = key.MustKeyValue()
	case LedgerEntryTypeRule:
		result.Value = key.MustRule()
	case LedgerEntryTypeRole:
		result.Value = key.MustRole()
	default:
		panic("unexpected ledger key type in CacheKey method")
	}

	return result
}

func (key LedgerKey) Hash() [32]byte {
	h := sha256.New()
	_, err := Marshal(h, key)
	if err != nil {
		panic("failed to marshal during ledger key Hash method")
	}

	result := [32]byte{}
	copy(result[:], h.Sum(nil))

	return result
}
