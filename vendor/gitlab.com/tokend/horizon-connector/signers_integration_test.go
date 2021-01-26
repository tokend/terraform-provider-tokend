package horizon_test

import (
	"testing"

	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
	"gitlab.com/distributed_lab/logan/v3"
)

func createSignerRule(source keypair.Address, kp keypair.Full) uint64 {
	tx := builder(connector()).Transaction(source).Op(&xdrbuild.CreateSignerRule{
		Resource: xdr.SignerRuleResource{
			Type: xdr.LedgerEntryTypeAny,
		},
		Action:  xdr.SignerRuleActionAny,
		Details: Details{
			"name": "signer_rule",
			"other_details": 1231123,
		},
	}).Sign(kp)

	result := mustSuccessSubmit(tx)
	ruleID := result.Result.MustResults()[0].MustTr().MustManageSignerRuleResult().MustSuccess().RuleId
	return uint64(ruleID)
}

func createSignerRoleWithNewRule(source keypair.Address, kp keypair.Full) uint64 {
	ruleID := createSignerRule(source, kp)

	tx := builder(connector()).Transaction(source).Op(&xdrbuild.CreateSignerRole{
		Details: Details{
			"name": "testing_awesome_roles",
		},
		Rules:   []uint64{ruleID},
	}).Sign(kp)

	result := mustSuccessSubmit(tx)
	roleID:= result.Result.MustResults()[0].MustTr().MustManageSignerRoleResult().MustSuccess().RoleId
	return uint64(roleID)
}

func createSigner(source keypair.Address, kp keypair.Full, roleID uint64) (keypair.Full) {
	destination, _ := keypair.Random()
	tx := builder(connector()).Transaction(source).Op(&xdrbuild.CreateSigner{
		SignerData: xdrbuild.SignerData{
			PublicKey: destination.Address(),
			RoleID: roleID,
			Weight: 1000,
			Identity: 123,
			Details:Details{
				"name": "signer_name",
			},
		},
	}).Sign(kp)

	_ = mustSuccessSubmit(tx)
	return destination
}

func TestCreateSignerWithRoleUpdate(t *testing.T) {
	source := keypair.MustParseAddress("GBA4EX43M25UPV4WIE6RRMQOFTWXZZRIPFAI5VPY6Z2ZVVXVWZ6NEOOB")
	kp := keypair.MustParseSeed("SAMJKTZVW5UOHCDK5INYJNORF2HRKYI72M5XSZCBYAHQHR34FFR4Z6G4")

	newAccount := createAccount(source, kp)
	logan.New().WithField("account", newAccount.Address()).Info("create new account to test it's signers")
	roleID := createSignerRoleWithNewRule(newAccount, newAccount)
	createSigner(newAccount, newAccount, roleID)
}
