package helpers

import (
	"testing"

	"gitlab.com/tokend/go/xdr"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/assert"
)

var accountRuleSchema = map[string]*schema.Schema{
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
	"resource_type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"resource_json": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

func TestAccountRules_Transaction(t *testing.T) {
	c := map[string]interface{}{
		"resource_type": "ledger_entry",
		"resource_json": `{"entry_type": "transaction"}`,
	}
	expected := &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeTransaction,
		},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := RuleEntry(resource)

	assert.NoError(t, err)
	assert.EqualValues(t, expected, got)
}

func TestAccountRules_Signer(t *testing.T) {
	c := map[string]interface{}{
		"resource_type": "ledger_entry",
		"resource_json": `{"entry_type": "signer","role_ids":["1", "2", "3"]}`,
	}
	expected := &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeSigner,
			Signer: &xdr.InternalRuleResourceSigner{
				RoleIDs: []xdr.Uint64{1, 2, 3},
				Ext:     xdr.EmptyExt{},
			},
		},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := RuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

//
//func TestAccountRules_Balance(t *testing.T) {
//	c := map[string]interface{}{
//		"entry_type": "balance",
//	}
//	expected := &xdr.AccountRuleResource{
//		Type: xdr.LedgerEntryTypeBalance,
//		Ext:  &xdr.EmptyExt{},
//	}
//
//	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
//	got, err := AccountRuleEntry(resource)
//
//	assert.NoError(t, err)
//	assert.Equal(t, expected, got)
//}
//
//func TestAccountRules_ReviewableRequest(t *testing.T) {
//	c := map[string]interface{}{
//		"entry_type": "reviewable_request",
//		"entry": map[string]interface{}{
//			"request_type": "*",
//		},
//	}
//	expected := &xdr.AccountRuleResource{
//		Type: xdr.LedgerEntryTypeReviewableRequest,
//		ReviewableRequest: &xdr.AccountRuleResourceReviewableRequest{
//			Details: xdr.ReviewableRequestResource{
//				RequestType: xdr.ReviewableRequestTypeAny,
//				Ext:         &xdr.EmptyExt{},
//			},
//		},
//	}
//
//	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
//	got, err := AccountRuleEntry(resource)
//
//	assert.NoError(t, err)
//	assert.Equal(t, expected, got)
//}
//
//func TestAccountRules_Asset(t *testing.T) {
//	t.Run("maxUint64", func(t *testing.T) {
//		c := map[string]interface{}{
//			"entry_type": "asset",
//			"entry": map[string]interface{}{
//				"asset_code": "*",
//				"asset_type": "*"},
//		}
//		expected := &xdr.AccountRuleResource{
//			Type: xdr.LedgerEntryTypeAsset,
//			Asset: &xdr.AccountRuleResourceAsset{
//				AssetCode: xdr.AssetCode("*"),
//				AssetType: math.MaxUint64,
//			},
//		}
//
//		resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
//		got, err := AccountRuleEntry(resource)
//
//		assert.NoError(t, err)
//		assert.Equal(t, expected, got)
//	})
//
//	t.Run("with value", func(t *testing.T) {
//		c := map[string]interface{}{
//			"entry_type": "asset",
//			"entry": map[string]interface{}{
//				"asset_code": "*",
//				"asset_type": "1"},
//		}
//		expected := &xdr.AccountRuleResource{
//			Type: xdr.LedgerEntryTypeAsset,
//			Asset: &xdr.AccountRuleResourceAsset{
//				AssetCode: xdr.AssetCode("*"),
//				AssetType: xdr.Uint64(1),
//			},
//		}
//
//		resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
//		got, err := AccountRuleEntry(resource)
//
//		assert.NoError(t, err)
//		assert.Equal(t, expected, got)
//	})
//}
//
//func TestAccountRules_Sale(t *testing.T) {
//	t.Run("maxUint64", func(t *testing.T) {
//		c := map[string]interface{}{
//			"entry_type": "sale",
//			"entry": map[string]interface{}{
//				"sale_id":   "*",
//				"sale_type": "*"},
//		}
//		expected := &xdr.AccountRuleResource{
//			Type: xdr.LedgerEntryTypeSale,
//			Sale: &xdr.AccountRuleResourceSale{
//				SaleId:   math.MaxUint64,
//				SaleType: math.MaxUint64,
//			},
//		}
//
//		resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
//		got, err := AccountRuleEntry(resource)
//
//		assert.NoError(t, err)
//		assert.Equal(t, expected, got)
//	})
//
//	t.Run("with value", func(t *testing.T) {
//		c := map[string]interface{}{
//			"entry_type": "sale",
//			"entry": map[string]interface{}{
//				"sale_id":   "1",
//				"sale_type": "1"},
//		}
//		expected := &xdr.AccountRuleResource{
//			Type: xdr.LedgerEntryTypeSale,
//			Sale: &xdr.AccountRuleResourceSale{
//				SaleId:   xdr.Uint64(1),
//				SaleType: xdr.Uint64(1),
//			},
//		}
//
//		resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
//		got, err := AccountRuleEntry(resource)
//
//		assert.NoError(t, err)
//		assert.Equal(t, expected, got)
//	})
//}
//
//func TestAccountRules_Vote(t *testing.T) {
//	t.Run("maxUint64", func(t *testing.T) {
//		c := map[string]interface{}{
//			"entry_type": "vote",
//			"entry": map[string]interface{}{
//				"poll_id":         "*",
//				"permission_type": "*"},
//		}
//		expected := &xdr.AccountRuleResource{
//			Type: xdr.LedgerEntryTypeVote,
//			Vote: &xdr.AccountRuleResourceVote{
//				PollId:         math.MaxUint64,
//				PermissionType: math.MaxUint32,
//			},
//		}
//
//		resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
//		got, err := AccountRuleEntry(resource)
//
//		assert.NoError(t, err)
//		assert.Equal(t, expected, got)
//	})
//}
