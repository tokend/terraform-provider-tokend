package balance

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/tokend/horizon-connector/internal/mocks"
	"fmt"
)

func TestQ_Balances(t *testing.T) {
	client := mocks.Client{}
	q := NewQ(&client)

	t.Run("Balance not found", func(t *testing.T) {
		client.On("Get", mock.Anything).Return(nil, nil).Once()
		defer client.AssertExpectations(t)

		got, err := q.AccountID("testBalanceID")
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("Balance exists", func(t *testing.T) {
		accID := "GD7AHJHCDSQI6LVMEJEE2FTNCA2LJQZ4R64GUI3PWANSVEO4GEOWB636"

		data := []byte(fmt.Sprintf(`{
        	"account_id": "%s"
    	}`, accID))

		client.On("Get", mock.Anything).Return(data, nil).Once()
		defer client.AssertExpectations(t)

		got, err := q.AccountID("testBalanceID")
		assert.NoError(t, err)
		assert.Equal(t, accID, *got)
	})
}
