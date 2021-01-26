package core

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestInfoResponseMarshal(t *testing.T) {
	response := []byte(`
	{
		"info": {
			"UNSAFE_QUORUM": "UNSAFE QUORUM ALLOWED",
			"base_assets": [
				"XAAU",
				"XAAG",
				"USD"
			],
			"base_exchange_name": "BullionCoin System",
			"build": "v0.5.0-495-ge104283-dirty",
			"commission_account_id": "GAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHV4",
			"demurrage_period": 9223372036854775807,
			"ledger": {
				"age": 1,
				"closeTime": 1510162165,
				"hash": "3dd956b26a0c8de99833dd1d6d93635ab7d126886d7740c9c2a37758c55a2b8c",
				"num": 28166
			},
			"master_account_id": "GDRYPVZ63SR7V2G46GKRGABJD3XPDNWQ4B4PQPJBTTDUEAKH5ZECPTSN",
			"network": "Test SDF Network ; September 2015",
			"numPeers": 0,
			"operational_account_id": "GABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVCX",
			"protocol_version": 8,
			"quorum": {
				"28165": {
					"agree": 1,
					"disagree": 0,
					"fail_at": 0,
					"hash": "d05a3b",
					"missing": 0,
					"phase": "EXTERNALIZE"
				}
			},
			"state": "Synced!",
			"statistics_quote_asset": "USD",
			"storage_fee_manager_account_id": "GABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABEQO",
			"tx_expiration_period": 601200,
			"withdrawal_details_max_length": 1000
		}
	}`)
	r := infoResponse{}
	if err := json.Unmarshal(response, &r); err != nil {
		t.Fatalf("expected nil got %s", err)
	}
	reference := Info{
		CoreVersion:                "v0.5.0-495-ge104283-dirty",
		NetworkPassphrase:          "Test SDF Network ; September 2015",
		MasterExchangeName:         "BullionCoin System",
		TxExpirationPeriod:         601200,
		WithdrawalDetailsMaxLength: 1000,
		DemurragePeriod:            9223372036854775807,
		BaseAssets:                 []string{"XAAU", "XAAG", "USD"},
		MasterAccountID:            "GDRYPVZ63SR7V2G46GKRGABJD3XPDNWQ4B4PQPJBTTDUEAKH5ZECPTSN",
		CommissionAccountID:        "GAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHV4",
		OperationalAccountID:       "GABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVCX",
		StorageFeeManageAccountID:  "GABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABEQO",
	}
	if !reflect.DeepEqual(reference, r.Info) {
		t.Fatalf("expected %#v got %#v", reference, r.Info)
	}

}
