package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/horizon-connector/internal/mocks"
	"gitlab.com/tokend/horizon-connector/internal/resources"
)

func TestQ_Balances(t *testing.T) {
	client := mocks.Client{}
	q := NewQ(&client)

	t.Run("account not found", func(t *testing.T) {
		client.On("Get", mock.Anything).Return(nil, nil).Once()
		defer client.AssertExpectations(t)

		got, err := q.Balances("foobar")
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("existing account", func(t *testing.T) {
		data := []byte(`[{
        	"account_id": "GD7AHJHCDSQI6LVMEJEE2FTNCA2LJQZ4R64GUI3PWANSVEO4GEOWB636",
        	"asset": "SUN",
        	"balance_id": "BBXG7L5SE6SHBG2UEOUUCMB3WK5PLQ2NFMDUVCLJNSXUN3FKVMKJCXRU"
    	}]`)
		expected := []resources.Balance{
			{
				Asset:     "SUN",
				AccountID: "GD7AHJHCDSQI6LVMEJEE2FTNCA2LJQZ4R64GUI3PWANSVEO4GEOWB636",
				BalanceID: "BBXG7L5SE6SHBG2UEOUUCMB3WK5PLQ2NFMDUVCLJNSXUN3FKVMKJCXRU",
			},
		}
		client.On("Get", mock.Anything).Return(data, nil).Once()
		defer client.AssertExpectations(t)

		got, err := q.Balances("foobar")
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("validate signers and types for account", func(t *testing.T) {
		t.Run("Get signers for master account", func(t *testing.T) {

			data := []byte(`{
  "signers": [
    {
      "public_key": "GAJHBDSVNMZVBJP24TMULBXK5VNJQEMCQKNWA2NRZ7HNA7OVJXLTKGZD",
      "weight": 51,
      "signer_type_i": 144,
      "signer_types": [
        {
          "name": "SignerTypeAssetManager",
          "value": 16
        },
        {
          "name": "SignerTypeIssuanceManager",
          "value": 128
        }
      ],
      "signer_identity": 12,
      "signer_name": "Withdraw signer 2"
    }
  ]
}`)

			client.On("Get", "/accounts/GDF6CDA63MU2IW6CQJPNOYEHQBHFF2XNHAPLR6ZUOJP3SBQRKROZFO7Z/signers").Return(data, nil).Once()
			defer client.AssertExpectations(t)

			err := q.IsSigner("GDF6CDA63MU2IW6CQJPNOYEHQBHFF2XNHAPLR6ZUOJP3SBQRKROZFO7Z", "GAJHBDSVNMZVBJP24TMULBXK5VNJQEMCQKNWA2NRZ7HNA7OVJXLTKGZD",
				xdr.SignerTypeAssetManager, xdr.SignerTypeIssuanceManager)
			assert.NoError(t, err)

		})

		t.Run("No such signer", func(t *testing.T) {

			client.On("Get", "/accounts/GDF6CDA63MU2IW6CQJPNOYEHQBHFF2XNHAPLR6ZUOJP3SBQRKROZFO7Z/signers").Return(nil, nil).Once()
			defer client.AssertExpectations(t)

			err := q.IsSigner("GDF6CDA63MU2IW6CQJPNOYEHQBHFF2XNHAPLR6ZUOJP3SBQRKROZFO7Z", "GAJHBDSVNMZVBJP24TMULBXK5VNJQEMCQKNWA2NRZ7HNA7OVJXLTKGZD")
			assert.Error(t, err)
			assert.EqualValues(t, ErrNoSigner, errors.Cause(err))

		})

		t.Run("No such type", func(t *testing.T) {
			data := []byte(`{
  "signers": [
    {
      "public_key": "GAJHBDSVNMZVBJP24TMULBXK5VNJQEMCQKNWA2NRZ7HNA7OVJXLTKGZD",
      "weight": 51,
      "signer_type_i": 16,
      "signer_types": [
        {
          "name": "SignerTypeNotVerifiedAccManager",
          "value": 2
        }
      ],
      "signer_identity": 12,
      "signer_name": "Withdraw signer 2"
    }
  ]
}`)
			client.On("Get", "/accounts/GDF6CDA63MU2IW6CQJPNOYEHQBHFF2XNHAPLR6ZUOJP3SBQRKROZFO7Z/signers").Return(data, nil).Once()
			defer client.AssertExpectations(t)

			err := q.IsSigner("GDF6CDA63MU2IW6CQJPNOYEHQBHFF2XNHAPLR6ZUOJP3SBQRKROZFO7Z", "GAJHBDSVNMZVBJP24TMULBXK5VNJQEMCQKNWA2NRZ7HNA7OVJXLTKGZD",
				xdr.SignerTypeAssetManager, xdr.SignerTypeIssuanceManager)
			assert.Error(t, err)
			assert.EqualValues(t, ErrNotEnoughType, errors.Cause(err))

		})
	})
}
