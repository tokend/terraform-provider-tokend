package responses

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionBadRequest_Unmarshal(t *testing.T) {
	cases := []struct {
		name     string
		body     string
		expected TransactionBadRequest
	}{
		{
			"rejected",
			`{
			  "type": "transaction_failed",
			  "title": "Transaction Failed",
			  "status": 400,
			  "instance": "oregon/dhMfCdDFR6-032252",
			  "extras": {
				"envelope_xdr": "AAAAAAAAA==",
				"result_codes": {
				  "transaction": "tx_bad_auth_extra",
				  "operations": ["op_yoba"]
				},
				"result_xdr": "AAAAAAAAAAD////5AAAAAA=="
			  }
			}`,
			TransactionBadRequest{
				Type: "transaction_failed",
				Extras: TransactionBadRequestExtras{
					EnvelopeXDR: "AAAAAAAAA==",
					ResultXDR:   "AAAAAAAAAAD////5AAAAAA==",
					ResultCodes: TransactionResultCodes{
						Transaction: "tx_bad_auth_extra",
						Operations:    []string{"op_yoba"},
					},
				},
			},
		},
		{
			"malformed",
			`{
			  "type": "transaction_malformed",
			  "title": "Transaction Malformed",
			  "status": 400,
			  "instance": "oregon/dhMfCdDFR6-032394",
			  "extras": {
				"envelope_xdr": "foobar"
			  }
			}`,
			TransactionBadRequest{
				Type: "transaction_malformed",
				Extras: TransactionBadRequestExtras{
					EnvelopeXDR: "foobar",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got TransactionBadRequest
			if err := json.Unmarshal([]byte(tc.body), &got); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tc.expected, got)
		})
	}
}
