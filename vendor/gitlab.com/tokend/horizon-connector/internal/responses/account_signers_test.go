package responses

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountSignersUnmarshal(t *testing.T) {
	data := []byte(`{
  "signers": [
    {
      "public_key": "GD7AHJHCDSQI6LVMEJEE2FTNCA2LJQZ4R64GUI3PWANSVEO4GEOWB636",
      "weight": 1,
      "signer_type_i": 16777215,
      "signer_types": [
        {
          "name": "SignerTypeReader",
          "value": 1
        },
        {
          "name": "SignerTypeNotVerifiedAccManager",
          "value": 2
        },
        {
          "name": "SignerTypeGeneralAccManager",
          "value": 4
        },
        {
          "name": "SignerTypeDirectDebitOperator",
          "value": 8
        },
        {
          "name": "SignerTypeAssetManager",
          "value": 16
        },
        {
          "name": "SignerTypeAssetRateManager",
          "value": 32
        },
        {
          "name": "SignerTypeBalanceManager",
          "value": 64
        },
        {
          "name": "SignerTypeIssuanceManager",
          "value": 128
        },
        {
          "name": "SignerTypeInvoiceManager",
          "value": 256
        },
        {
          "name": "SignerTypePaymentOperator",
          "value": 512
        },
        {
          "name": "SignerTypeLimitsManager",
          "value": 1024
        },
        {
          "name": "SignerTypeAccountManager",
          "value": 2048
        },
        {
          "name": "SignerTypeCommissionBalanceManager",
          "value": 4096
        },
        {
          "name": "SignerTypeOperationalBalanceManager",
          "value": 8192
        },
        {
          "name": "SignerTypeEventsChecker",
          "value": 16384
        },
        {
          "name": "SignerTypeExchangeAccManager",
          "value": 32768
        },
        {
          "name": "SignerTypeSyndicateAccManager",
          "value": 65536
        },
        {
          "name": "SignerTypeUserAssetManager",
          "value": 131072
        },
        {
          "name": "SignerTypeUserIssuanceManager",
          "value": 262144
        },
        {
          "name": "SignerTypeWithdrawManager",
          "value": 524288
        },
        {
          "name": "SignerTypeFeesManager",
          "value": 1048576
        },
        {
          "name": "SignerTypeTxSender",
          "value": 2097152
        },
        {
          "name": "SignerTypeAmlAlertManager",
          "value": 4194304
        },
        {
          "name": "SignerTypeAmlAlertReviewer",
          "value": 8388608
        }
      ],
      "signer_identity": 1,
      "signer_name": "foobar"
    }
  ]
}`)

	var got AccountSigners
	if err := json.Unmarshal(data, &got); err != nil {
		assert.NoError(t, err)
	}

	assert.Len(t, got.Signers, 1)
	signer := got.Signers[0]
	assert.EqualValues(t, 1, signer.Weight)
	assert.EqualValues(t, 16777215, signer.SignerType)
	assert.EqualValues(t, "GD7AHJHCDSQI6LVMEJEE2FTNCA2LJQZ4R64GUI3PWANSVEO4GEOWB636", signer.AccountID)
	assert.EqualValues(t, 1, signer.Identity)
	assert.EqualValues(t, "foobar", signer.Name)
}
