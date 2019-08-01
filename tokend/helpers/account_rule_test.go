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
	"entry_type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"entry": {
		Type:     schema.TypeMap,
		Optional: true,
	},
}

func TestAccountRules_Transaction(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "transaction",
	}
	expected := &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeTransaction,
		Ext:  &xdr.EmptyExt{},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := AccountRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestAccountRules_Signer(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "signer",
	}
	expected := &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeSigner,
		Ext:  &xdr.EmptyExt{},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := AccountRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestAccountRules_ReviewableRequest(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "reviewable_request",
		"entry": map[string]interface{} {
			"request_type": "*",
		},
	}
	expected := &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeReviewableRequest,
		ReviewableRequest: &xdr.AccountRuleResourceReviewableRequest{
			RequestType: xdr.ReviewableRequestTypeAny,
			Ext: xdr.EmptyExt{},
		},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := AccountRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

