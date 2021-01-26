package regources

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo_Unmarshal(t *testing.T) {
	body := `{
	  "_links": {
		"account": {
		  "href": "http://localhost:8000/accounts/{account_id}",
		  "templated": true
		},
		"account_transactions": {
		  "href": "http://localhost:8000/accounts/{account_id}/transactions{?cursor,limit,order}",
		  "templated": true
		},
		"metrics": {
		  "href": "http://localhost:8000/metrics"
		},
		"self": {
		  "href": "http://localhost:8000/"
		},
		"transaction": {
		  "href": "http://localhost:8000/transactions/{hash}",
		  "templated": true
		},
		"transactions": {
		  "href": "http://localhost:8000/transactions{?cursor,limit,order}",
		  "templated": true
		}
	  },
	  "history_latest_ledger": 3587,
	  "history_elder_ledger": 1,
	  "core_latest_ledger": 3587,
	  "core_elder_ledger": 1,
	  "network_passphrase": "Test Network",
	  "commission_account_id": "GAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHV4",
	  "operational_account_id": "GABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVCX",
	  "storage_fee_account_id": "",
	  "admin_account_id": "GDHK26UFBGC63UBQCVQLHJD6RAQXLAS7RKJAR5FZQAWMCUBFHRNKFSKC",
	  "master_exchange_name": "MasterExchange",
	  "tx_expiration_period": 601200
	}`
	expected := Info{
		MasterAccountID:    "GDHK26UFBGC63UBQCVQLHJD6RAQXLAS7RKJAR5FZQAWMCUBFHRNKFSKC",
		TXExpirationPeriod: 601200,
		Passphrase:         "Test Network",
	}
	var got Info
	assert.NoError(t, json.Unmarshal([]byte(body), &got))
	assert.Equal(t, expected, got)
}
