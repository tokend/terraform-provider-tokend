package xdr

import "gitlab.com/distributed_lab/logan/v3/errors"

func (u *ReviewableRequestOperation) ToOperationBody() OperationBody {
	result := OperationBody{
		Type: u.Type,
	}

	switch u.Type {
	case OperationTypeCreateAccount:
		result.CreateAccountOp = u.CreateAccountOp
	case OperationTypeCreateSigner:
		result.CreateSignerOp = u.CreateSignerOp
	case OperationTypeUpdateSigner:
		result.UpdateSignerOp = u.UpdateSignerOp
	case OperationTypeRemoveSigner:
		result.RemoveSignerOp = u.RemoveSignerOp
	case OperationTypeCreateRule:
		result.CreateRuleOp = u.CreateRuleOp
	case OperationTypeUpdateRule:
		result.UpdateRuleOp = u.UpdateRuleOp
	case OperationTypeRemoveRule:
		result.RemoveRuleOp = u.RemoveRuleOp
	case OperationTypeCreateRole:
		result.CreateRoleOp = u.CreateRoleOp
	case OperationTypeUpdateRole:
		result.UpdateRoleOp = u.UpdateRoleOp
	case OperationTypeRemoveRole:
		result.RemoveRoleOp = u.RemoveRoleOp
	case OperationTypeCreateBalance:
		result.CreateBalanceOp = u.CreateBalanceOp
	case OperationTypePayment:
		result.PaymentOp = u.PaymentOp
	case OperationTypeIssuance:
		result.IssuanceOp = u.IssuanceOp
	case OperationTypeDestruction:
		result.DestructionOp = u.DestructionOp
	case OperationTypeChangeAccountRoles:
		result.ChangeAccountRolesOp = u.ChangeAccountRolesOp
	case OperationTypeCreateAsset:
		result.CreateAssetOp = u.CreateAssetOp
	case OperationTypeUpdateAsset:
		result.UpdateAssetOp = u.UpdateAssetOp
	case OperationTypeCreateData:
		result.CreateDataOp = u.CreateDataOp
	case OperationTypeUpdateData:
		result.UpdateAssetOp = u.UpdateAssetOp
	case OperationTypeRemoveData:
		result.RemoveDataOp = u.RemoveDataOp
	case OperationTypeInitiateKycRecovery:
		result.InitiateKycRecoveryOp = u.InitiateKycRecoveryOp
	case OperationTypeKycRecovery:
		result.KycRecoveryOp = u.KycRecoveryOp
	default:
		panic("unexpected reviewable request operation type")
	}

	return result
}

func (u *OperationBody) ToReviewableRequestOp() (ReviewableRequestOperation, error) {
	result := ReviewableRequestOperation{
		Type: u.Type,
	}

	switch u.Type {
	case OperationTypeCreateAccount:
		result.CreateAccountOp = u.CreateAccountOp
	case OperationTypeCreateSigner:
		result.CreateSignerOp = u.CreateSignerOp
	case OperationTypeUpdateSigner:
		result.UpdateSignerOp = u.UpdateSignerOp
	case OperationTypeRemoveSigner:
		result.RemoveSignerOp = u.RemoveSignerOp
	case OperationTypeCreateRule:
		result.CreateRuleOp = u.CreateRuleOp
	case OperationTypeUpdateRule:
		result.UpdateRuleOp = u.UpdateRuleOp
	case OperationTypeRemoveRule:
		result.RemoveRuleOp = u.RemoveRuleOp
	case OperationTypeCreateRole:
		result.CreateRoleOp = u.CreateRoleOp
	case OperationTypeUpdateRole:
		result.UpdateRoleOp = u.UpdateRoleOp
	case OperationTypeRemoveRole:
		result.RemoveRoleOp = u.RemoveRoleOp
	case OperationTypeCreateBalance:
		result.CreateBalanceOp = u.CreateBalanceOp
	case OperationTypePayment:
		result.PaymentOp = u.PaymentOp
	case OperationTypeIssuance:
		result.IssuanceOp = u.IssuanceOp
	case OperationTypeDestruction:
		result.DestructionOp = u.DestructionOp
	case OperationTypeChangeAccountRoles:
		result.ChangeAccountRolesOp = u.ChangeAccountRolesOp
	case OperationTypeCreateAsset:
		result.CreateAssetOp = u.CreateAssetOp
	case OperationTypeUpdateAsset:
		result.UpdateAssetOp = u.UpdateAssetOp
	case OperationTypeCreateData:
		result.CreateDataOp = u.CreateDataOp
	case OperationTypeUpdateData:
		result.UpdateAssetOp = u.UpdateAssetOp
	case OperationTypeRemoveData:
		result.RemoveDataOp = u.RemoveDataOp
	case OperationTypeInitiateKycRecovery:
		result.InitiateKycRecoveryOp = u.InitiateKycRecoveryOp
	case OperationTypeKycRecovery:
		result.KycRecoveryOp = u.KycRecoveryOp
	default:
		return result, errors.From(
			errors.New("unexpected operation type for reviewable request operation"),
			map[string]interface{}{
				"operation_type": u.Type.String(),
			},
		)
	}

	return result, nil
}
