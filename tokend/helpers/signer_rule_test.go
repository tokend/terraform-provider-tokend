package helpers

import (
	"math"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
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
			RequestType:   xdr.ReviewableRequestTypeAny,
		},
	}

	resource := schema.TestResourceDataRaw(t, accountRuleSchema, c)
	got, err := SignerRuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}
