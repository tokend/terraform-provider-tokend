package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type RuleEntryFunc func(d *schema.ResourceData) (*xdr.RuleResource, error)

var RuleEntries = map[string]RuleEntryFunc{
	"ledger_entry": internalRuleEntry,
	"custom":       customRuleEntry,
}

var RuleLedgerEntries = map[string]RuleEntryFunc{
	"reviewable_request": ruleResourceReviewableRequest,
	"asset":              ruleResourceAsset,
	"balance":            ruleResourceBalance,
	"transaction":        ruleResourceTransaction,
	"role":               ruleResourceRole,
	"signer":             ruleResourceSigner,
	"key_value":          ruleResourceKeyValue,
	"data":               ruleResourceData,
}

var reviewableRequestEntries = map[string]RuleEntryFunc{
	"asset":       ruleResourceAsset,
	"balance":     ruleResourceBalance,
	"transaction": ruleResourceTransaction,
	"role":        ruleResourceRole,
	"signer":      ruleResourceSigner,
	"key_value":   ruleResourceKeyValue,
	"data":        ruleResourceData,
}

func RuleEntry(d *schema.ResourceData) (*xdr.RuleResource, error) {
	tpe := d.Get("resource_type").(string)
	createEntry, ok := RuleEntries[tpe]
	if !ok {
		return nil, fmt.Errorf(`resource_type "%s" is not supported`, tpe)
	}

	resource, err := createEntry(d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule resource")
	}

	return resource, nil
}

func reviewableRequestRuleEntry(d *schema.ResourceData) (*xdr.RuleResource, error) {
	tpe := d.Get("resource_type").(string)
	createEntry, ok := reviewableRequestEntries[tpe]
	if !ok {
		return nil, fmt.Errorf(`resource_type "%s" is not supported`, tpe)
	}

	resource, err := createEntry(d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule resource")
	}

	return resource, nil
}

func internalRuleEntry(d *schema.ResourceData) (*xdr.RuleResource, error) {
	res := d.Get("resource_json").(string)
	resMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(res), &resMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal resource_json")
	}
	tpe := resMap["entry_type"].(string)
	createEntry, ok := RuleLedgerEntries[tpe]
	if !ok {
		return nil, fmt.Errorf(`entry_type "%s" is not supported`, tpe)
	}
	resource, err := createEntry(d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule resource")
	}

	return resource, nil
}

func customRuleEntry(d *schema.ResourceData) (*xdr.RuleResource, error) {
	resourceType := d.Get("resource.custom_type").(string)
	resourcePayload := d.Get("resource.custom_payload").(string)
	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeCustom,
		CustomRuleResource: &xdr.CustomRuleResource{
			ResourceType:    xdr.Longstring(resourceType),
			ResourcePayload: xdr.Longstring(resourcePayload),
		},
	}, nil
}

func ruleResourceTransaction(_ *schema.ResourceData) (*xdr.RuleResource, error) {
	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeTransaction,
		},
	}, nil
}

func ruleResourceSigner(d *schema.ResourceData) (*xdr.RuleResource, error) {
	res := d.Get("resource_json").(string)
	resMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(res), &resMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal resource_json")
	}

	rawRoleIDs := resMap["role_ids"].([]interface{})
	roles := make([]xdr.Uint64, 0, len(rawRoleIDs))
	for _, r := range rawRoleIDs {
		roleID, err := WildCardUint64FromRaw(r.(string))
		if err != nil {
			return nil, errors.Wrap(err, "failed to cast roleID to uint64")
		}
		roles = append(roles, xdr.Uint64(roleID))
	}

	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeSigner,
			Signer: &xdr.InternalRuleResourceSigner{
				RoleIDs: roles,
				Ext:     xdr.EmptyExt{},
			},
		},
	}, nil
}

func ruleResourceBalance(d *schema.ResourceData) (*xdr.RuleResource, error) {

	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeBalance,
		},
	}, nil
}

func ruleResourceKeyValue(d *schema.ResourceData) (*xdr.RuleResource, error) {
	prefix := d.Get("resource.prefix").(string)
	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeKeyValue,
			KeyValue: &xdr.InternalRuleResourceKeyValue{
				KeyPrefix: xdr.Longstring(prefix),
			},
		},
	}, nil
}

func ruleResourceAsset(d *schema.ResourceData) (*xdr.RuleResource, error) {
	var internalResource xdr.InternalRuleResource
	internalResource.Type = xdr.LedgerEntryTypeAsset
	entry := d.Get("resource").(map[string]interface{})
	assetCode := entry["asset_code"].(string)
	secTypeRaw := entry["security_type"].(string)
	secType, err := WildCardUint32FromRaw(secTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast asset_type")
	}

	internalResource.Asset = &xdr.InternalRuleResourceAsset{
		AssetCode:    xdr.AssetCode(assetCode),
		SecurityType: xdr.Uint32(secType),
	}
	return &xdr.RuleResource{
		ResourceType:         xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &internalResource,
	}, nil
}

func ruleResourceData(d *schema.ResourceData) (*xdr.RuleResource, error) {
	var internalResource xdr.InternalRuleResource
	internalResource.Type = xdr.LedgerEntryTypeData
	entry := d.Get("resource").(map[string]interface{})
	secTypeRaw := entry["security_type"].(string)
	secType, err := WildCardUint32FromRaw(secTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast asset_type")
	}

	internalResource.Data = &xdr.InternalRuleResourceData{
		SecurityType: xdr.Uint32(secType),
	}
	return &xdr.RuleResource{
		ResourceType:         xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &internalResource,
	}, nil
}

func ruleResourceRole(d *schema.ResourceData) (*xdr.RuleResource, error) {
	var internalResource xdr.InternalRuleResource
	internalResource.Type = xdr.LedgerEntryTypeAsset
	entry := d.Get("resource").(map[string]interface{})
	roleIDRaw := entry["role_id"].(string)
	roleID, err := WildCardUint64FromRaw(roleIDRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast role_id")
	}
	internalResource.Role = &xdr.InternalRuleResourceRole{
		RoleId: xdr.Uint64(roleID),
	}
	return &xdr.RuleResource{
		ResourceType:         xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &internalResource,
	}, nil
}

func ruleResourceReviewableRequest(d *schema.ResourceData) (*xdr.RuleResource, error) {
	securityTypeRaw := d.Get("resource.security_type").(string)
	securityType, err := WildCardUint32FromRaw(securityTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast security_type")
	}

	operations := d.Get("resource.op_rules").([]interface{})
	opRules := make([]xdr.ReviewableRequestOperationRule, 0, len(operations))

	for _, opr := range operations {
		opRule := opr.(*schema.ResourceData)

		opResource, err := reviewableRequestRuleEntry(opRule)
		if err != nil {
			return nil, errors.Wrap(err, "failed to cast entry")
		}

		//todo probably should validate
		opRules = append(opRules, xdr.ReviewableRequestOperationRule{
			Resource: *opResource.InternalRuleResource,
		})
	}
	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			ReviewableRequest: &xdr.InternalRuleResourceReviewableRequest{
				SecurityType: xdr.Uint32(securityType),
				OpRules:      opRules,
			},
		},
	}, nil
}

func RuleAction(d *schema.ResourceData) (*xdr.RuleAction, error) {
	actionTypeRaw := d.Get("action").(string)
	var actionType xdr.RuleActionType
	if actionTypeRaw == "*" {
		actionType = xdr.RuleActionTypeAny
	} else {
		for _, guess := range xdr.RuleActionTypeAll {
			fmt.Println(guess.ShortString(), actionTypeRaw)
			if guess.ShortString() == actionTypeRaw {
				actionType = guess
			}
		}
		if actionType == 0 {
			return nil, fmt.Errorf("unknown rule action type: %s", actionTypeRaw)
		}
	}

	tpe := d.Get("action_type").(string)
	actionFunc, ok := RuleActions[tpe]
	if !ok {
		return nil, fmt.Errorf(`resource_type "%s" is not supported`, tpe)
	}

	action, err := actionFunc(d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule resource")
	}

	action.Type = actionType
	return action, nil
}

type RuleActionFunc func(d *schema.ResourceData) (*xdr.RuleAction, error)

var RuleActions = map[string]RuleActionFunc{
	"create":            ruleActionCreate,
	"update":            ruleActionUpdate,
	"destroy":           ruleActionDestroy,
	"send":              ruleActionSend,
	"receive":           ruleActionReceive,
	"receive_issuance":  ruleActionReceiveIssuance,
	"initiate_recovery": ruleActionInitiateRecovery,
	"review":            ruleActionReview,
	"custom":            ruleActionCustom,
}

func ruleActionCreate(d *schema.ResourceData) (*xdr.RuleAction, error) {
	return &xdr.RuleAction{
		Create: &xdr.RuleActionCreate{
			ForOther: d.Get("action.for_other").(bool),
		},
	}, nil
}

func ruleActionUpdate(d *schema.ResourceData) (*xdr.RuleAction, error) {
	return &xdr.RuleAction{
		Update: &xdr.RuleActionUpdate{
			ForOther: d.Get("action.for_other").(bool),
		},
	}, nil
}

func ruleActionDestroy(d *schema.ResourceData) (*xdr.RuleAction, error) {
	securityTypeRaw := d.Get("action.security_type").(string)
	securityType, err := WildCardUint32FromRaw(securityTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast security_type")
	}

	return &xdr.RuleAction{
		Destroy: &xdr.RuleActionDestroy{
			ForOther:     d.Get("action.for_other").(bool),
			SecurityType: xdr.Uint32(securityType),
		},
	}, nil
}

func ruleActionSend(d *schema.ResourceData) (*xdr.RuleAction, error) {
	securityTypeRaw := d.Get("action.security_type").(string)
	securityType, err := WildCardUint32FromRaw(securityTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast security_type")
	}

	return &xdr.RuleAction{
		Send: &xdr.RuleActionSend{
			SecurityType: xdr.Uint32(securityType),
		},
	}, nil
}

func ruleActionReceive(d *schema.ResourceData) (*xdr.RuleAction, error) {
	securityTypeRaw := d.Get("action.security_type").(string)
	securityType, err := WildCardUint32FromRaw(securityTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast security_type")
	}

	return &xdr.RuleAction{
		Receive: &xdr.RuleActionReceive{
			SecurityType: xdr.Uint32(securityType),
		},
	}, nil
}

func ruleActionReceiveIssuance(d *schema.ResourceData) (*xdr.RuleAction, error) {
	securityTypeRaw := d.Get("action.security_type").(string)
	securityType, err := WildCardUint32FromRaw(securityTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast security_type")
	}

	return &xdr.RuleAction{
		ReceiveIssuance: &xdr.RuleActionReceiveIssuance{
			SecurityType: xdr.Uint32(securityType),
		},
	}, nil
}

func ruleActionInitiateRecovery(d *schema.ResourceData) (*xdr.RuleAction, error) {
	roleIDs := d.Get("action.role_ids").([]uint64)
	roles := make([]xdr.Uint64, 0, len(roleIDs))
	for _, r := range roleIDs {
		roles = append(roles, xdr.Uint64(r))
	}

	return &xdr.RuleAction{
		InitiateRecovery: &xdr.RuleActionInitiateRecovery{
			RoleIDs: roles,
		},
	}, nil
}

func ruleActionReview(d *schema.ResourceData) (*xdr.RuleAction, error) {
	tasksToAdd := d.Get("action.tasks_to_add").(uint64)
	tasksToRemove := d.Get("action.tasks_to_remove").(uint64)

	return &xdr.RuleAction{
		Review: &xdr.RuleActionReview{
			TasksToAdd:    xdr.Uint64(tasksToAdd),
			TasksToRemove: xdr.Uint64(tasksToRemove),
		},
	}, nil
}

func ruleActionCustom(d *schema.ResourceData) (*xdr.RuleAction, error) {
	customType := d.Get("action.custom_type").(string)
	customPayload := d.Get("action.custom_payload").(string)

	return &xdr.RuleAction{
		CustomRuleAction: &xdr.CustomRuleAction{
			ActionType:    xdr.Longstring(customType),
			ActionPayload: xdr.Longstring(customPayload),
		},
	}, nil
}
