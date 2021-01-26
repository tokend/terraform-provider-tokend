package resources

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestSaleUnmarshal(t *testing.T) {
	cases := []struct {
		name     string
		data     string
		expected Sale
	}{
		{
			name: "sale",
			data: `{
				  "paging_token": "1",
				  "id": "1",
				  "owner_id": "GACIKE6WPDNCLNHZYGP5G2GO3VDIXB6AJQCC54IMFEBOEKNAOU4JH4ZR",
				  "base_asset": "BTC890",
				  "default_quote_asset": "USD242",
				  "start_time": "2018-03-16T12:13:34Z",
				  "end_time": "2018-03-16T12:23:34Z",
				  "soft_cap": "2250.000000",
				  "hard_cap": "4500.000000",
				  "details": {
					"description": "Token sale description",
					"logo": {
					  "type": "logo_type",
					  "url": "logo_url"
					},
					"name": "sale name",
					"short_description": "short description"
				  },
				  "state": {
					"name": "closed",
					"value": 2
				  },
				  "statistics": {
					"investors": 1
				  },
				  "quote_assets": {
					"quote_assets": [
					  {
						"asset": "USD242",
						"price": "4.500000",
						"quote_balance_id": "BA2CLLLFLGN6WRRM34XHLDQU4UNHH6CWOX53HN3QCRQO5GH7PJ6HKUIM",
						"current_cap": "4500.000000",
						"total_current_cap": "4500.000000",
						"hard_cap": "4500.000000"
					  }
					]
				  },
				  "base_hard_cap": "1000.000000",
				  "base_current_cap": "1000.000000",
				  "current_cap": "4500.000000",
				  "sale_type": {
					"name": "basic_sale",
					"value": 1
				  }
				}`,
			expected: Sale{
				ID: "1",
				Details: SaleDetails{
					Name: "sale name",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got Sale
			err := json.Unmarshal([]byte(tc.data), &got)
			if err != nil {
				t.Fatal(err)
			}
			assert.EqualValues(t, tc.expected, got)
		})
	}
}
