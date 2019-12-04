package helpers

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type Map map[string]interface{}

type RuleEntryFunc func(d Map) (*xdr.RuleResource, error)

var RuleEntries = map[string]RuleEntryFunc{
	"custom":             customRuleEntry,
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
	"asset":     ruleResourceAsset,
	"balance":   ruleResourceBalance,
	"role":      ruleResourceRole,
	"signer":    ruleResourceSigner,
	"key_value": ruleResourceKeyValue,
	"data":      ruleResourceData,
	"custom":    customRuleEntry,
}

func RuleEntry(d *schema.ResourceData) (*xdr.RuleResource, error) {
	tpe := d.Get("entry_type").(string)
	createEntry, ok := RuleEntries[tpe]
	if !ok {
		return nil, fmt.Errorf(`entry_type "%s" is not supported`, tpe)
	}
	var entry map[string]interface{}
	t, ok := d.GetOk(tpe)
	if ok {
		switch t.(type) {
		case []interface{}:
			entry = t.([]interface{})[0].(map[string]interface{})
		case map[string]interface{}:
			entry = t.(map[string]interface{})
		}
	}
	resource, err := createEntry(entry)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule resource")
	}

	return resource, nil
}

func reviewableRequestRuleEntry(d Map) (*xdr.RuleResource, error) {
	tpe := d["entry_type"].(string)
	createEntry, ok := reviewableRequestEntries[tpe]
	if !ok {
		return nil, fmt.Errorf(`resource_type "%s" is not supported`, tpe)
	}
	t := d[tpe]
	var entry map[string]interface{}
	switch t.(type) {
	case []interface{}:
		entry = t.([]interface{})[0].(map[string]interface{})
	case map[string]interface{}:
		entry = t.(map[string]interface{})
	}

	resource, err := createEntry(entry)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule resource")
	}

	return resource, nil
}

func customRuleEntry(d Map) (*xdr.RuleResource, error) {
	resourceType := d["custom.resource_type"].(string)
	resourcePayload := d["custom.resource_payload"].(string)
	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeCustom,
		CustomRuleResource: &xdr.CustomRuleResource{
			ResourceType:    xdr.Longstring(resourceType),
			ResourcePayload: xdr.Longstring(resourcePayload),
		},
	}, nil
}

func ruleResourceTransaction(_ Map) (*xdr.RuleResource, error) {
	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeTransaction,
		},
	}, nil
}

func ruleResourceSigner(d Map) (*xdr.RuleResource, error) {
	roleIDs := d["role_ids"].([]interface{})
	roles := make([]xdr.Uint64, 0, len(roleIDs))
	for _, r := range roleIDs {
		roleID, err := WildCardUint64FromRaw(r.(string))
		if err != nil {
			return nil, errors.Wrap(err, "failed to cast role_id")
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

func ruleResourceBalance(_ Map) (*xdr.RuleResource, error) {
	return ruleResourceBalanceRaw()
}

func ruleResourceBalanceRaw() (*xdr.RuleResource, error) {
	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeBalance,
		},
	}, nil
}

func ruleResourceKeyValue(d Map) (*xdr.RuleResource, error) {
	prefix := d["prefix"].(string)
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

func ruleResourceAsset(d Map) (*xdr.RuleResource, error) {
	var internalResource xdr.InternalRuleResource
	internalResource.Type = xdr.LedgerEntryTypeAsset
	assetCode := d["asset_code"].(string)
	secTypeRaw := d["security_type"].(string)
	secType, err := WildCardUint32FromRaw(secTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast security_type")
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

func ruleResourceData(d Map) (*xdr.RuleResource, error) {
	var internalResource xdr.InternalRuleResource
	internalResource.Type = xdr.LedgerEntryTypeData
	secTypeRaw := d["security_type"].(string)
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

func ruleResourceRole(d Map) (*xdr.RuleResource, error) {
	var internalResource xdr.InternalRuleResource
	internalResource.Type = xdr.LedgerEntryTypeAsset
	roleIDRaw := d["role_id"].(string)
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

func ruleResourceReviewableRequest(d Map) (*xdr.RuleResource, error) {
	securityTypeRaw := d["security_type"].(string)
	securityType, err := WildCardUint32FromRaw(securityTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast security_type")
	}

	opRulesRaw := d["op_rules"].([]interface{})

	opRules := make([]xdr.ReviewableRequestOperationRule, 0, len(opRulesRaw))

	for _, opr := range opRulesRaw {
		opResource, err := reviewableRequestRuleEntry(opr.(map[string]interface{}))
		if err != nil {
			return nil, errors.Wrap(err, "failed to cast entry")
		}

		action, err := ruleActionRaw(opr.(map[string]interface{}))
		if err != nil {
			return nil, errors.Wrap(err, "failed to cast action")
		}

		opRules = append(opRules, xdr.ReviewableRequestOperationRule{
			Resource: *opResource,
			Action:   *action,
		})
	}
	return &xdr.RuleResource{
		ResourceType: xdr.RuleResourceTypeLedgerEntry,
		InternalRuleResource: &xdr.InternalRuleResource{
			Type: xdr.LedgerEntryTypeReviewableRequest,
			ReviewableRequest: &xdr.InternalRuleResourceReviewableRequest{
				SecurityType: xdr.Uint32(securityType),
				OpRules:      opRules,
			},
		},
	}, nil
}

func RuleAction(d *schema.ResourceData) (*xdr.RuleAction, error) {
	actionTypeRaw := d.Get("action_type").(string)
	actionRaw := d.Get("action").([]interface{})
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

	var entry map[string]interface{}
	if len(actionRaw) != 0 {
		actionMap := actionRaw[0].(map[string]interface{})
		entry = GetMapFromInterface(actionMap[actionTypeRaw])
	}

	actionFunc, ok := RuleActions[actionTypeRaw]
	if !ok {
		return nil, fmt.Errorf(`action_type "%s" is not supported`, actionTypeRaw)
	}

	action, err := actionFunc(entry)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule resource")
	}

	action.Type = actionType
	return action, nil
}

func ruleActionRaw(d Map) (*xdr.RuleAction, error) {
	actionTypeRaw := d["action_type"].(string)
	actionRaw := d["action"].([]interface{})

	var actionType xdr.RuleActionType
	if actionTypeRaw == "*" {
		actionType = xdr.RuleActionTypeAny
	} else {
		for _, guess := range xdr.RuleActionTypeAll {
			if guess.ShortString() == actionTypeRaw {
				actionType = guess
			}
		}
		if actionType == 0 {
			return nil, fmt.Errorf("unknown rule action type: %s", actionTypeRaw)
		}
	}

	var entry map[string]interface{}
	if len(actionRaw) != 0 {
		actionMap := actionRaw[0].(map[string]interface{})
		entry = GetMapFromInterface(actionMap[actionTypeRaw])
	}
	actionFunc, ok := RuleActions[actionTypeRaw]
	if !ok {
		return nil, fmt.Errorf(`action_type "%s" is not supported`, actionType)
	}

	action, err := actionFunc(entry)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule resource")
	}

	action.Type = actionType
	return action, nil
}

type RuleActionFunc func(d Map) (*xdr.RuleAction, error)

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

func ruleActionCreate(d Map) (*xdr.RuleAction, error) {
	var forOther bool
	if b, ok := d["for_other"]; ok {
		forOther = cast.ToBool(b)
	}
	return &xdr.RuleAction{
		Create: &xdr.RuleActionCreate{
			ForOther: forOther,
		},
	}, nil
}

func ruleActionUpdate(d Map) (*xdr.RuleAction, error) {
	var forOther bool
	if b, ok := d["for_other"]; ok {
		forOther = cast.ToBool(b)
	}
	return &xdr.RuleAction{
		Update: &xdr.RuleActionUpdate{
			ForOther: forOther,
		},
	}, nil
}

func ruleActionDestroy(d Map) (*xdr.RuleAction, error) {
	securityTypeRaw := d["security_type"].(string)
	securityType, err := WildCardUint32FromRaw(securityTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast security_type")
	}

	var forOther bool
	if b, ok := d["for_other"]; ok {
		forOther = cast.ToBool(b)
	}

	return &xdr.RuleAction{
		Destroy: &xdr.RuleActionDestroy{
			ForOther:     forOther,
			SecurityType: xdr.Uint32(securityType),
		},
	}, nil
}

func ruleActionSend(d Map) (*xdr.RuleAction, error) {
	securityTypeRaw := d["security_type"].(string)
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

func ruleActionReceive(d Map) (*xdr.RuleAction, error) {
	securityTypeRaw := d["security_type"].(string)
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

func ruleActionReceiveIssuance(d Map) (*xdr.RuleAction, error) {
	securityTypeRaw := d["security_type"].(string)
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

func ruleActionInitiateRecovery(d Map) (*xdr.RuleAction, error) {
	roleIDs := d["action_role_ids"].([]interface{})
	roles := make([]xdr.Uint64, 0, len(roleIDs))
	for _, r := range roleIDs {
		roleID, err := WildCardUint64FromRaw(r.(string))
		if err != nil {
			return nil, errors.Wrap(err, "failed to cast roleID to uint64")
		}

		roles = append(roles, xdr.Uint64(roleID))
	}

	return &xdr.RuleAction{
		InitiateRecovery: &xdr.RuleActionInitiateRecovery{
			RoleIDs: roles,
		},
	}, nil
}

func ruleActionReview(d Map) (*xdr.RuleAction, error) {
	tasksToAddRaw := d["tasks_to_add"].(string)
	tasksToRemoveRaw := d["tasks_to_remove"].(string)

	tasksToAdd, err := WildCardUint64FromRaw(tasksToAddRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast roleID to uint64")
	}
	tasksToRemove, err := WildCardUint64FromRaw(tasksToRemoveRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast roleID to uint64")
	}

	return &xdr.RuleAction{
		Review: &xdr.RuleActionReview{
			TasksToAdd:    xdr.Uint64(tasksToAdd),
			TasksToRemove: xdr.Uint64(tasksToRemove),
		},
	}, nil
}

func ruleActionCustom(d Map) (*xdr.RuleAction, error) {
	customType := d["custom_type"].(string)
	customPayload := d["custom_payload"].(string)

	return &xdr.RuleAction{
		CustomRuleAction: &xdr.CustomRuleAction{
			ActionType:    xdr.Longstring(customType),
			ActionPayload: xdr.Longstring(customPayload),
		},
	}, nil
}
