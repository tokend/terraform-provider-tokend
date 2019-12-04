package tokend

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tokend/terraform-provider-tokend/tokend/connector"

	"github.com/tokend/terraform-provider-tokend/tokend/helpers"

	"gitlab.com/tokend/go/xdr"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdrbuild"
)

func resourceRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRuleCreate,
		Update: resourceRuleUpdate,
		Read:   resourceRuleRead,
		Delete: resourceRuleDelete,
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
		},
	}
}

type VoidDetails struct{}

func (VoidDetails) MarshalJSON() ([]byte, error) {
	return []byte(`{}`), nil
}

func resourceRuleCreate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	resource, err := helpers.RuleEntry(d)
	if err != nil {
		return errors.Wrap(err, "failed to cast entry")
	}

	action, err := helpers.RuleAction(d)
	if err != nil {
		return errors.Wrap(err, "failed to cast action")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.CreateRule{
		Resource: *resource,
		Action:   *action,
		Forbid:   d.Get("forbids").(bool),
		Details:  VoidDetails{},
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	resp, err := m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(connector.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
	}
	var txResult xdr.TransactionResult
	if err := xdr.SafeUnmarshalBase64(resp.Data.Attributes.ResultXdr, &txResult); err != nil {
		return errors.Wrap(err, "failed to decode result")
	}
	txCodes := *(txResult.Result.Results)
	ruleID := txCodes[0].Tr.CreateRuleResult.Success.RuleId
	d.SetId(fmt.Sprintf("%d", ruleID))
	return nil
}

func resourceRuleUpdate(d *schema.ResourceData, _m interface{}) (err error) {
	m := _m.(Meta)

	id, err := cast.ToUint64E(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to cast account role id")
	}

	resource, err := helpers.RuleEntry(d)
	if err != nil {
		return errors.Wrap(err, "failed to cast entry")
	}

	action, err := helpers.RuleAction(d)
	if err != nil {
		return errors.Wrap(err, "failed to cast action")
	}

	env, err := m.Builder.Transaction(m.Source).Op(&xdrbuild.UpdateRule{
		ID:       id,
		Resource: *resource,
		Action:   *action,
		Forbid:   d.Get("forbids").(bool),
		Details:  VoidDetails{},
	}).Sign(m.Signer).Marshal()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx")
	}
	_, err = m.Submitter.Submit(context.TODO(), env, true)
	if err != nil {
		if txErr, ok := err.(connector.TxFailure); ok {
			return errors.Wrapf(err, "failed to submit tx: %s %q", txErr.TransactionResultCode, txErr.OperationResultCodes)
		}
		return errors.Wrap(err, "unknown error occurred")
	}
	return nil
}

func resourceRuleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRuleDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
	//return errors.New("tokend_account_rule delete is not implemented")
}
