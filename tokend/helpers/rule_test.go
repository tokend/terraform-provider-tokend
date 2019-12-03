package helpers

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
	"math"
	"testing"
)

var ruleschema = map[string]*schema.Schema{
	"action": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"create": {
					Type:     schema.TypeMap,
					Optional: true,
				},
				"update": {
					Type:     schema.TypeMap,
					Optional: true,
				},
				"issue": {
					Type:     schema.TypeMap,
					Optional: true,
				},
				"send": {
					Type:     schema.TypeMap,
					Optional: true,
				},
				"receive": {
					Type:     schema.TypeMap,
					Optional: true,
				},
				"receive_issuance": {
					Type:     schema.TypeMap,
					Optional: true,
				},
				"remove": {
					Type:     schema.TypeMap,
					Optional: true,
				},
				"destroy": {
					Type:     schema.TypeMap,
					Optional: true,
				},
				"review": {
					Type:     schema.TypeMap,
					Optional: true,
				},
				"change_roles": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"role_ids": {
								Type:     schema.TypeList,
								Required: true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},
					},
				},
				"initiate_recovery": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"role_ids": {
								Type:     schema.TypeList,
								Required: true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},
					},
				},
			},
		},
	},
	"action_type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"forbids": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"details": {
		Type:     schema.TypeMap,
		Optional: true,
	},
	"entry_type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"reviewable_request": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"op_rules": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"action": {
								Type:     schema.TypeList,
								Optional: true,
								MaxItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"create": {
											Type:     schema.TypeMap,
											Optional: true,
										},
										"update": {
											Type:     schema.TypeMap,
											Optional: true,
										},
										"issue": {
											Type:     schema.TypeMap,
											Optional: true,
										},
										"send": {
											Type:     schema.TypeMap,
											Optional: true,
										},
										"receive": {
											Type:     schema.TypeMap,
											Optional: true,
										},
										"receive_issuance": {
											Type:     schema.TypeMap,
											Optional: true,
										},
										"remove": {
											Type:     schema.TypeMap,
											Optional: true,
										},
										"destroy": {
											Type:     schema.TypeMap,
											Optional: true,
										},
										"review": {
											Type:     schema.TypeMap,
											Optional: true,
										},
										"change_roles": {
											Type:     schema.TypeList,
											Optional: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"role_ids": {
														Type:     schema.TypeList,
														Required: true,
														Elem: &schema.Schema{
															Type: schema.TypeString,
														},
													},
												},
											},
										},
										"initiate_recovery": {
											Type:     schema.TypeList,
											Optional: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"role_ids": {
														Type:     schema.TypeList,
														Required: true,
														Elem: &schema.Schema{
															Type: schema.TypeString,
														},
													},
												},
											},
										},
									},
								},
							},
							"action_type": {
								Type:     schema.TypeString,
								Required: true,
							},
							"forbids": {
								Type:     schema.TypeBool,
								Optional: true,
							},
							"details": {
								Type:     schema.TypeMap,
								Optional: true,
							},
							"entry_type": {
								Type:     schema.TypeString,
								Required: true,
							},
							"asset": {
								Type:     schema.TypeMap,
								Optional: true,
							},
							"role": {
								Type:     schema.TypeMap,
								Optional: true,
							},
							"signer": {
								Type:     schema.TypeList,
								Optional: true,
								MaxItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"role_ids": {
											Type: schema.TypeList,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
									},
								},
							},
							"key_value": {
								Type:     schema.TypeMap,
								Optional: true,
							},
							"data": {
								Type:     schema.TypeMap,
								Optional: true,
							},
						},
					},
				},
				"security_type": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
	"asset": {
		Type:     schema.TypeMap,
		Optional: true,
	},
	"role": {
		Type:     schema.TypeMap,
		Optional: true,
	},
	"signer": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"role_ids": {
					Type: schema.TypeList,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
	"key_value": {
		Type:     schema.TypeMap,
		Optional: true,
	},
	"data": {
		Type:     schema.TypeMap,
		Optional: true,
	},
	"custom": {
		Type:     schema.TypeMap,
		Optional: true,
	},
}

func TestRules_Transaction(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "transaction",
	}
	expected := &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeTransaction,
		},
	}

	resource := schema.TestResourceDataRaw(t, ruleschema, c)
	got, err := RuleEntry(resource)

	assert.NoError(t, err)
	assert.EqualValues(t, expected, got)
}

func TestRules_Signer(t *testing.T) {
	c := map[string]interface{}{
		"resource_type": "ledger_entry",
		"entry_type":    "signer",
		"signer": []interface{}{
			map[string]interface{}{
				"role_ids": []string{
					"1", "2", "3",
				},
			},
		},
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

	resource := schema.TestResourceDataRaw(t, ruleschema, c)
	got, err := RuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestRules_Balance(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "balance",
	}
	expected := &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeBalance,
		},
	}

	resource := schema.TestResourceDataRaw(t, ruleschema, c)
	got, err := RuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestRules_KeyValue(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "key_value",
		"key_value": map[string]interface{}{
			"prefix": "a",
		},
	}
	expected := &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeKeyValue,
			KeyValue: &xdr.InternalRuleResourceKeyValue{
				KeyPrefix: xdr.Longstring("a"),
				Ext:       xdr.EmptyExt{},
			},
		},
	}

	resource := schema.TestResourceDataRaw(t, ruleschema, c)
	got, err := RuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestRules_ReviewableRequest(t *testing.T) {
	c := map[string]interface{}{
		"entry_type": "reviewable_request",
		"reviewable_request": []interface{}{
			map[string]interface{}{
				"security_type": "*",
				"op_rules": []interface{}{
					map[string]interface{}{
						"entry_type":  "balance",
						"action_type": "create",
					},
					map[string]interface{}{
						"entry_type":  "signer",
						"action_type": "create",
						"action": []interface{}{
							map[string]interface{}{
								"create": map[string]interface{}{
									"for_other": true,
								},
							},
						},
						"signer": []interface{}{
							map[string]interface{}{
								"role_ids": []string{
									"1", "2", "3",
								},
							},
						},
					},
				},
			},
		},
	}
	expected := &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeReviewableRequest,
			ReviewableRequest: &xdr.InternalRuleResourceReviewableRequest{
				OpRules: []xdr.ReviewableRequestOperationRule{
					{
						Resource: xdr.InternalRuleResource{
							Type: xdr.LedgerEntryTypeBalance,
						},
						Action: xdr.RuleAction{
							Type: xdr.RuleActionTypeCreate,
							Create: &xdr.RuleActionCreate{
								ForOther: false,
								Ext:      xdr.EmptyExt{},
							},
						},
					},

					{
						Resource: xdr.InternalRuleResource{
							Type: xdr.LedgerEntryTypeSigner,
							Signer: &xdr.InternalRuleResourceSigner{
								RoleIDs: []xdr.Uint64{1, 2, 3},
								Ext:     xdr.EmptyExt{},
							},
						},
						Action: xdr.RuleAction{
							Type: xdr.RuleActionTypeCreate,
							Create: &xdr.RuleActionCreate{
								ForOther: true,
								Ext:      xdr.EmptyExt{},
							},
						},
					},
				},
				SecurityType: xdr.Uint32(math.MaxUint32),
				Ext:          xdr.EmptyExt{},
			},
		},
	}

	resource := schema.TestResourceDataRaw(t, ruleschema, c)
	got, err := RuleEntry(resource)

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestRules_Asset(t *testing.T) {
	t.Run("maxUint32", func(t *testing.T) {
		c := map[string]interface{}{
			"entry_type": "asset",
			"asset": map[string]interface{}{
				"asset_code":    "*",
				"security_type": "*"},
		}

		expected := &xdr.RuleResource{
			ResourceType: xdr.RuleResourceTypeLedgerEntry,
			InternalRuleResource: &xdr.InternalRuleResource{
				Type: xdr.LedgerEntryTypeAsset,
				Asset: &xdr.InternalRuleResourceAsset{
					AssetCode:    xdr.AssetCode("*"),
					SecurityType: math.MaxUint32,
				},
			},
		}

		resource := schema.TestResourceDataRaw(t, ruleschema, c)
		got, err := RuleEntry(resource)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("with value", func(t *testing.T) {
		c := map[string]interface{}{
			"entry_type": "asset",
			"asset": map[string]interface{}{
				"asset_code":    "*",
				"security_type": "1"},
		}
		expected := &xdr.RuleResource{
			ResourceType: xdr.RuleResourceTypeLedgerEntry,
			InternalRuleResource: &xdr.InternalRuleResource{
				Type: xdr.LedgerEntryTypeAsset,
				Asset: &xdr.InternalRuleResourceAsset{
					AssetCode:    xdr.AssetCode("*"),
					SecurityType: xdr.Uint32(1),
				},
			},
		}

		resource := schema.TestResourceDataRaw(t, ruleschema, c)
		got, err := RuleEntry(resource)

		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}
