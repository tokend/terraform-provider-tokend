package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountIDFromRaw(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		invalidAccountID := "invalid"
		accountID, err := AccountIDFromRaw(invalidAccountID)
		assert.Error(t, err)
		assert.Nil(t, accountID)
	})

	t.Run("valid", func(t *testing.T) {
		invalidAccountID := "GBA4EX43M25UPV4WIE6RRMQOFTWXZZRIPFAI5VPY6Z2ZVVXVWZ6NEOOB"
		accountID, err := AccountIDFromRaw(invalidAccountID)
		assert.NoError(t, err)
		//assert.NNil(t, accountID)
		assert.Equal(t, invalidAccountID, accountID.Address())
	})
}

func TestPolicyFromRaw(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		policiesNames := []interface{}{"withdrawable", "transferable"}
		policyCode, err := PolicyFromRaw(policiesNames)
		assert.NoError(t, err)
		assert.NotNil(t, policyCode)
		assert.Equal(t, uint32(9), policyCode)
	})

	t.Run("invalid", func(t *testing.T) {
		policiesNames := []interface{}{"invalid_data", "invalid_data_too"}
		policyCode, err := PolicyFromRaw(policiesNames)
		assert.Error(t, err)
		assert.Equal(t, uint32(0), policyCode)

	})

}

func TestTypeFromRaw(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		typeName := "withdraw"
		typeCode, err := StatsOpTypeFromRaw(typeName)
		assert.NoError(t, err)
		assert.NotNil(t, typeCode)
		assert.Equal(t, int32(1), typeCode)
	})

	t.Run("invalid", func(t *testing.T) {
		typeName := "aaaabbbccd"
		typeCode, err := StatsOpTypeFromRaw(typeName)
		assert.Error(t, err)
		assert.Equal(t, uint32(0), typeCode)

	})
}
