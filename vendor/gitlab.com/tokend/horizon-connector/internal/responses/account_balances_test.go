package responses

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountBalancesUnmarshal(t *testing.T) {
	data := []byte(`[
    {
        "account_id": "GD7AHJHCDSQI6LVMEJEE2FTNCA2LJQZ4R64GUI3PWANSVEO4GEOWB636",
        "asset": "SUN",
        "balance_id": "BBXG7L5SE6SHBG2UEOUUCMB3WK5PLQ2NFMDUVCLJNSXUN3FKVMKJCXRU"
    },
    {
        "account_id": "GD7AHJHCDSQI6LVMEJEE2FTNCA2LJQZ4R64GUI3PWANSVEO4GEOWB636",
        "asset": "USD",
        "balance_id": "BBKXV4BIZDGDW2QPAXQNJVKIVW7NJN6674LQ6DEQ4PGRQP7CGTQDKLHO"
    }]`)

	var got AccountBalances
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	assert.Len(t, got, 2)
	balance := got[0]
	assert.EqualValues(t, "SUN", balance.Asset)
	assert.EqualValues(t, "GD7AHJHCDSQI6LVMEJEE2FTNCA2LJQZ4R64GUI3PWANSVEO4GEOWB636", balance.AccountID)
	assert.EqualValues(t, "BBXG7L5SE6SHBG2UEOUUCMB3WK5PLQ2NFMDUVCLJNSXUN3FKVMKJCXRU", balance.BalanceID)
}
