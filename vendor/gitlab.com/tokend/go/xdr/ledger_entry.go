package xdr

import "fmt"

// LedgerKey implements the `Keyer` interface
// Deprecated: LedgerKey is not fully // TODO implementation for all types, used in core
func (entry *LedgerEntry) LedgerKey() LedgerKey {
	var body interface{}

	switch entry.Data.Type {
	case LedgerEntryTypeAccount:
		account := entry.Data.MustAccount()
		body = LedgerKeyAccount{
			AccountId: account.AccountId,
		}
	case LedgerEntryTypeKeyValue:
		keyValue := entry.Data.MustKeyValue()
		body = LedgerKeyKeyValue{
			Key: keyValue.Key,
		}
	case LedgerEntryTypeRule:
		rule := entry.Data.MustRule()
		body = LedgerKeyRule{
			Id: rule.Id,
		}
	case LedgerEntryTypeRole:
		role := entry.Data.MustRole()
		body = LedgerKeyRole{
			Id: role.Id,
		}
	case LedgerEntryTypeSigner:
		signer := entry.Data.MustSigner()
		body = LedgerKeySigner{
			PubKey:    signer.PubKey,
			AccountId: signer.AccountId,
		}
	case LedgerEntryTypeReviewableRequest:
		reviewableRequest := entry.Data.MustReviewableRequest()
		body = LedgerKeyReviewableRequest{
			RequestId: reviewableRequest.RequestId,
		}
	case LedgerEntryTypeAsset:
		asset := entry.Data.MustAsset()
		body = LedgerKeyAsset{
			Code: asset.Code,
			Ext:  LedgerKeyAssetExt{},
		}
	case LedgerEntryTypeData:
		data := entry.Data.MustData()
		body = LedgerKeyData{
			Id:  data.Id,
			Ext: EmptyExt{},
		}
	case LedgerEntryTypeBalance:
		balance := entry.Data.MustBalance()
		body = LedgerKeyBalance{
			BalanceId: balance.BalanceId,
			Ext:       LedgerKeyBalanceExt{},
		}
	default:
		panic(fmt.Errorf("unknown entry type: %v", entry.Data.Type))
	}

	ret, err := NewLedgerKey(entry.Data.Type, body)
	if err != nil {
		panic(err)
	}

	return ret
}
