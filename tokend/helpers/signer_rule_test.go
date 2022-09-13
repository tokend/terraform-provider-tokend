package helpers

import (
	"math"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
)

var signerRuleSchema = map[string]*schema.Schema{
	"action": {
		Type:     schema.TypeString,
		Required: true,
	},
	"forbids": {
		Type:     schema.TypeBool,
		Optional: true,
		//Default:  false,
	},
	"details": {
		Type:     schema.TypeMap,
		Optional: true,
	},
	"entry_type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"entry": {
		Type:     schema.TypeMap,
		Optional: true,
	},
}

func TestSignerRules_Balance(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "transaction",
	}
	expected := &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeTransaction,
		Ext:  &xdr.EmptyExt{},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := SignerRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestSignerRules_Fee(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "fee",
	}
	expected := &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeFee,
		Ext:  &xdr.EmptyExt{},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := SignerRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestSignerRules_Limits(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "limits",
	}
	expected := &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeLimitsV2,
		Ext:  &xdr.EmptyExt{},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := SignerRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestSignerRules_Signer(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "signer",
	}
	expected := &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeSigner,
		Ext:  &xdr.EmptyExt{},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := SignerRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestSignerRules_Asset(t *testing.T) {
	t.Run("maxUint64", func(t *testing.T) {
		c := map[string]interface{}{
			"entry_type": "asset",
			"entry": map[string]interface{}{
				"asset_code": "*",
				"asset_type": "*"},
		}
		expected := &xdr.SignerRuleResource{
			Type: xdr.LedgerEntryTypeAsset,
			Asset: &xdr.SignerRuleResourceAsset{
				AssetCode: xdr.AssetCode("*"),
				AssetType: math.MaxUint64,
			},
		}

		resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
		got, err := SignerRuleEntry(resource)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("with value", func(t *testing.T) {
		c := map[string]interface{}{
			"entry_type": "asset",
			"entry": map[string]interface{}{
				"asset_code": "*",
				"asset_type": "1"},
		}
		expected := &xdr.SignerRuleResource{
			Type: xdr.LedgerEntryTypeAsset,
			Asset: &xdr.SignerRuleResourceAsset{
				AssetCode: xdr.AssetCode("*"),
				AssetType: xdr.Uint64(1),
			},
		}

		resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
		got, err := SignerRuleEntry(resource)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

}

func TestSignerRules_Sale(t *testing.T) {
	t.Run("maxUint64", func(t *testing.T) {
		c := map[string]interface{}{
			"entry_type": "sale",
			"entry": map[string]interface{}{
				"sale_id":   "*",
				"sale_type": "*"},
		}
		expected := &xdr.SignerRuleResource{
			Type: xdr.LedgerEntryTypeSale,
			Sale: &xdr.SignerRuleResourceSale{
				SaleId:   math.MaxUint64,
				SaleType: math.MaxUint64,
			},
		}

		resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
		got, err := SignerRuleEntry(resource)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("with value", func(t *testing.T) {
		c := map[string]interface{}{
			"entry_type": "sale",
			"entry": map[string]interface{}{
				"sale_id":   "1",
				"sale_type": "1"},
		}
		expected := &xdr.SignerRuleResource{
			Type: xdr.LedgerEntryTypeSale,
			Sale: &xdr.SignerRuleResourceSale{
				SaleId:   xdr.Uint64(1),
				SaleType: xdr.Uint64(1),
			},
		}

		resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
		got, err := SignerRuleEntry(resource)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

}

func TestSignerRules_KeyValue(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "key_value",
	}
	expected := &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeKeyValue,
		KeyValue: &xdr.SignerRuleResourceKeyValue{
			KeyPrefix: "",
		},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := SignerRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestSignerRules_Transaction(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "transaction",
	}
	expected := &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeTransaction,
		Ext:  &xdr.EmptyExt{},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := SignerRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestSignerRules_ReviewableRequest(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "reviewable_request",
		"entry": map[string]interface{}{
			"request_type": "*",
		},
	}
	expected := &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeReviewableRequest,
		ReviewableRequest: &xdr.SignerRuleResourceReviewableRequest{
			TasksToAdd:    xdr.Uint64(math.MaxUint64),
			TasksToRemove: xdr.Uint64(math.MaxUint64),
			AllTasks:      xdr.Uint64(math.MaxUint64),
			Details: xdr.ReviewableRequestResource{
				RequestType: xdr.ReviewableRequestTypeAny,
				Ext:         &xdr.EmptyExt{},
			},
		},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := SignerRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}
