package horizon_test

import (
	"context"
	"net/url"
	"testing"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/horizon-connector"
	"gitlab.com/tokend/keypair"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"fmt"
)

func connector() *horizon.Connector {
	base, _ := url.Parse("http://localhost:8000") // SAMJKTZVW5UOHCDK5INYJNORF2HRKYI72M5XSZCBYAHQHR34FFR4Z6G4
	return horizon.NewConnector(base)
}

func builder(horizon *horizon.Connector) *xdrbuild.Builder {
	info, err := connector().System().Info()
	if err != nil {
		panic(errors.Wrap(err, "failed to get info"))
	}
	return xdrbuild.NewBuilder(info.Passphrase, info.TXExpirationPeriod)
}

func mustSuccessSubmit(tx *xdrbuild.Transaction) xdr.TransactionResult {
	envelope, err := tx.Marshal()
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal tx"))
	}

	response := connector().Submitter().Submit(context.Background(), envelope)
	if response.Err != nil {
		logan.New().WithField("response", response).WithError(response.Err).Panic("failed to submit tx")
	}

	var result xdr.TransactionResult
	err = xdr.SafeUnmarshalBase64(response.ResultXDR, &result)
	if err != nil {
		panic(errors.Wrap(err, "failed to unmarshal result"))
	}

	if result.Result.Code != xdr.TransactionResultCodeTxSuccess {
		resultAsJSON, err := json.Marshal(result)
		if err != nil {
			panic(errors.Wrap(err, "failed to marshal tx result into json"))
		}
		panic(errors.From(errors.New("transaction failed"), logan.F{
			"tx_result": resultAsJSON,
		}))
	}

	return result
}

type Details map[string]interface{}

func (d Details) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}(d))
}

func createAccountRule(source keypair.Address, kp keypair.Full) uint64 {
	tx := builder(connector()).Transaction(source).Op(&xdrbuild.CreateAccountRule{
		Resource: xdr.AccountRuleResource{
			Type: xdr.LedgerEntryTypeAny,
		},
		Action:  xdr.AccountRuleActionAny,
		Details: Details{
			"name": "test_name",
			"other_details": 1231123,
		},
	}).Sign(kp)

	result := mustSuccessSubmit(tx)
	ruleID := result.Result.MustResults()[0].MustTr().MustManageAccountRuleResult().MustSuccess().RuleId
	return uint64(ruleID)
}

func createAccountRoleWithNewRule(source keypair.Address, kp keypair.Full) uint64 {
	ruleID := createAccountRule(source, kp)

	tx := builder(connector()).Transaction(source).Op(&xdrbuild.CreateAccountRole{
		Details: Details{
			"name": "testing_awesome_roles",
		},
		Rules:   []uint64{ruleID},
	}).Sign(kp)

	result := mustSuccessSubmit(tx)
	roleID:= result.Result.MustResults()[0].MustTr().MustManageAccountRoleResult().MustSuccess().RoleId
	return uint64(roleID)
}

func createAccount(source keypair.Address, kp keypair.Full) (keypair.Full) {
	destination, _ := keypair.Random()
	tx := builder(connector()).Transaction(source).Op(&xdrbuild.CreateAccount{
		Destination: destination.Address(),
		RoleID:      1,
		Signers: []xdrbuild.SignerData{
			{
				PublicKey: destination.Address(),
				RoleID:    1,
				Weight:    1000,
				Identity:  1,
				Details:   Details{},
			},
		},
	}).Sign(kp)

	_ = mustSuccessSubmit(tx)
	return destination
}

type kvOperation xdr.ManageKeyValueOp

func (op kvOperation) XDR() (*xdr.Operation, error) {
	rawOp := xdr.ManageKeyValueOp(op)
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageKeyValue,
			ManageKeyValueOp: &rawOp,
		},
	}, nil
}

func mustSetChangeRoleTasks(source keypair.Address, kp keypair.Full, fromRole, toRole uint64, tasks uint32) {
	rawTasks := xdr.Uint32(tasks)
	op := kvOperation(xdr.ManageKeyValueOp{
		Key: xdr.Longstring(fmt.Sprintf("change_role_tasks:%d:%d", fromRole, toRole)),
		Action:xdr.ManageKeyValueOpAction{
			Action: xdr.ManageKvActionPut,
			Value: &xdr.KeyValueEntryValue{
				Type: xdr.KeyValueEntryTypeUint32,
				Ui32Value:&rawTasks,
			},
		},
	})

	tx := builder(connector()).Transaction(source).Op(op).Sign(kp)
	_ = mustSuccessSubmit(tx)
}

func mustUpdateRole(source keypair.Address, kp keypair.Full, account string, oldRole, newRole uint64) {
	mustSetChangeRoleTasks(source, kp, oldRole, newRole, 12)

	tasks := uint32(0)
	tx := builder(connector()).Transaction(source).Op(&xdrbuild.CreateChangeRoleRequest{
		RequestID: 0,
		DestinationAccount: account,
		RoleToSet: newRole,
		KYCData: map[string]interface{}{
			"name": "Yoba",
			"last_name": "vihodi",
		},
		AllTasks: &tasks,
	}).Sign(kp)

	result := mustSuccessSubmit(tx)
	isFulfilled := result.Result.MustResults()[0].MustTr().MustCreateChangeRoleRequestResult().MustSuccess().Fulfilled
	if !isFulfilled {
		panic("expected request to change role to be auto fulfilled")
	}
}

func TestCreateAccountWithRoleUpdate(t *testing.T) {
	source := keypair.MustParseAddress("GBA4EX43M25UPV4WIE6RRMQOFTWXZZRIPFAI5VPY6Z2ZVVXVWZ6NEOOB")
	kp := keypair.MustParseSeed("SAMJKTZVW5UOHCDK5INYJNORF2HRKYI72M5XSZCBYAHQHR34FFR4Z6G4")

	newRole := createAccountRoleWithNewRule(source, kp)
	newAccount := createAccount(source, kp)
	mustUpdateRole(source, kp, newAccount.Address(), 1, newRole)
}
