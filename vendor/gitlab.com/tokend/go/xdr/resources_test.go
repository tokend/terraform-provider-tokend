package xdr_test

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

func TestFlag(t *testing.T) {
	expected := `{"value":10,"flags":[{"name":"base_asset","value":2},{"name":"withdrawable","value":8}]}`
	value := xdr.AssetPolicyBaseAsset ^ xdr.AssetPolicyWithdrawable

	t.Run("marshal", func(t *testing.T) {
		got, err := json.Marshal(&value)
		if err != nil {
			t.Fatal(err)
		}
		assert.JSONEq(t, string(got), expected)
	})

	t.Run("unmarshal", func(t *testing.T) {
		var got xdr.AssetPolicy
		assert.NoError(t, json.Unmarshal([]byte(expected), &got))
		assert.Equal(t, got, value)
	})
}

func TestEnum(t *testing.T) {
	expected := `{"value":26,"name":"manage_sale"}`
	value := xdr.OperationTypeManageSale

	t.Run("marshal", func(t *testing.T) {
		got, err := json.Marshal(&value)
		if err != nil {
			t.Fatal(err)
		}
		assert.JSONEq(t, string(got), expected)
	})

	t.Run("unmarshal", func(t *testing.T) {
		var got xdr.OperationType
		assert.NoError(t, json.Unmarshal([]byte(expected), &got))
		assert.Equal(t, got, value)
	})
}
