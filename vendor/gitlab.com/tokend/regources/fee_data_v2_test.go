package regources

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/amount"
	"testing"
)

func TestFeeDataV2_Unmarshal(t *testing.T) {
	body := `{
		"actual_payment_fee": "50.000000",
		"actual_payment_fee_asset_code": "ETH169",
		"fixed_fee": "10.000000"
	}`
	expected := FeeDataV2{
		FixedFee:                  10 * amount.One,
		ActualPaymentFee:          50 * amount.One,
		ActualPaymentFeeAssetCode: "ETH169",
	}
	var got FeeDataV2
	assert.NoError(t, json.Unmarshal([]byte(body), &got))
	assert.Equal(t, expected, got)
}
