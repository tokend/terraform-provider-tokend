// revision: e98bc93279bf4e9fbb9107380be2fb8c799236b5
// branch:   keys
// Package xdr is generated from:
//
//  xdr/SCP.x
//  xdr/ledger-entries-account-KYC.x
//  xdr/ledger-entries-account-role.x
//  xdr/ledger-entries-account-rule.x
//  xdr/ledger-entries-account.x
//  xdr/ledger-entries-key-value.x
//  xdr/ledger-entries-reference.x
//  xdr/ledger-entries-reviewable-request.x
//  xdr/ledger-entries-signer-role.x
//  xdr/ledger-entries-signer-rule.x
//  xdr/ledger-entries-signer.x
//  xdr/ledger-entries.x
//  xdr/ledger-keys.x
//  xdr/ledger.x
//  xdr/operation-cancel-change-role-request.x
//  xdr/operation-create-account.x
//  xdr/operation-create-change-role-request.x
//  xdr/operation-create-kyc-recovery-request.x
//  xdr/operation-initiate-kyc-recovery.x
//  xdr/operation-manage-account-role.x
//  xdr/operation-manage-account-rule.x
//  xdr/operation-manage-key-value.x
//  xdr/operation-manage-signer-role.x
//  xdr/operation-manage-signer-rule.x
//  xdr/operation-manage-signer.x
//  xdr/operation-review-request.x
//  xdr/overlay.x
//  xdr/resource-account-rule.x
//  xdr/resource-signer-rule.x
//  xdr/reviewable-request-change-role.x
//  xdr/reviewable-request-kyc-recovery.x
//  xdr/transaction.x
//  xdr/types.x
//
// DO NOT EDIT or your changes may be overwritten
package xdr

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/nullstyle/go-xdr/xdr3"
)

// Unmarshal reads an xdr element from `r` into `v`.
func Unmarshal(r io.Reader, v interface{}) (int, error) {
	// delegate to xdr package's Unmarshal
	return xdr.Unmarshal(r, v)
}

// Marshal writes an xdr element `v` into `w`.
func Marshal(w io.Writer, v interface{}) (int, error) {
	// delegate to xdr package's Marshal
	return xdr.Marshal(w, v)
}

// Value is an XDR Typedef defines as:
//
//   typedef opaque Value<>;
//
type Value []byte

// ScpBallot is an XDR Struct defines as:
//
//   struct SCPBallot
//    {
//        uint32 counter; // n
//        Value value;    // x
//    };
//
type ScpBallot struct {
	Counter Uint32 `json:"counter,omitempty"`
	Value   Value  `json:"value,omitempty"`
}

// ScpStatementType is an XDR Enum defines as:
//
//   enum SCPStatementType
//    {
//        PREPARE = 0,
//        CONFIRM = 1,
//        EXTERNALIZE = 2,
//        NOMINATE = 3
//    };
//
type ScpStatementType int32

const (
	ScpStatementTypePrepare     ScpStatementType = 0
	ScpStatementTypeConfirm     ScpStatementType = 1
	ScpStatementTypeExternalize ScpStatementType = 2
	ScpStatementTypeNominate    ScpStatementType = 3
)

var ScpStatementTypeAll = []ScpStatementType{
	ScpStatementTypePrepare,
	ScpStatementTypeConfirm,
	ScpStatementTypeExternalize,
	ScpStatementTypeNominate,
}

var scpStatementTypeMap = map[int32]string{
	0: "ScpStatementTypePrepare",
	1: "ScpStatementTypeConfirm",
	2: "ScpStatementTypeExternalize",
	3: "ScpStatementTypeNominate",
}

var scpStatementTypeShortMap = map[int32]string{
	0: "prepare",
	1: "confirm",
	2: "externalize",
	3: "nominate",
}

var scpStatementTypeRevMap = map[string]int32{
	"ScpStatementTypePrepare":     0,
	"ScpStatementTypeConfirm":     1,
	"ScpStatementTypeExternalize": 2,
	"ScpStatementTypeNominate":    3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ScpStatementType
func (e ScpStatementType) ValidEnum(v int32) bool {
	_, ok := scpStatementTypeMap[v]
	return ok
}
func (e ScpStatementType) isFlag() bool {
	for i := len(ScpStatementTypeAll) - 1; i >= 0; i-- {
		expected := ScpStatementType(2) << uint64(len(ScpStatementTypeAll)-1) >> uint64(len(ScpStatementTypeAll)-i)
		if expected != ScpStatementTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ScpStatementType) String() string {
	name, _ := scpStatementTypeMap[int32(e)]
	return name
}

func (e ScpStatementType) ShortString() string {
	name, _ := scpStatementTypeShortMap[int32(e)]
	return name
}

func (e ScpStatementType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ScpStatementTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ScpStatementType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ScpStatementType(t.Value)
	return nil
}

// ScpNomination is an XDR Struct defines as:
//
//   struct SCPNomination
//    {
//        Hash quorumSetHash; // D
//        Value votes<>;      // X
//        Value accepted<>;   // Y
//    };
//
type ScpNomination struct {
	QuorumSetHash Hash    `json:"quorumSetHash,omitempty"`
	Votes         []Value `json:"votes,omitempty"`
	Accepted      []Value `json:"accepted,omitempty"`
}

// ScpStatementPrepare is an XDR NestedStruct defines as:
//
//   struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            }
//
type ScpStatementPrepare struct {
	QuorumSetHash Hash       `json:"quorumSetHash,omitempty"`
	Ballot        ScpBallot  `json:"ballot,omitempty"`
	Prepared      *ScpBallot `json:"prepared,omitempty"`
	PreparedPrime *ScpBallot `json:"preparedPrime,omitempty"`
	NC            Uint32     `json:"nC,omitempty"`
	NH            Uint32     `json:"nH,omitempty"`
}

// ScpStatementConfirm is an XDR NestedStruct defines as:
//
//   struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            }
//
type ScpStatementConfirm struct {
	Ballot        ScpBallot `json:"ballot,omitempty"`
	NPrepared     Uint32    `json:"nPrepared,omitempty"`
	NCommit       Uint32    `json:"nCommit,omitempty"`
	NH            Uint32    `json:"nH,omitempty"`
	QuorumSetHash Hash      `json:"quorumSetHash,omitempty"`
}

// ScpStatementExternalize is an XDR NestedStruct defines as:
//
//   struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            }
//
type ScpStatementExternalize struct {
	Commit              ScpBallot `json:"commit,omitempty"`
	NH                  Uint32    `json:"nH,omitempty"`
	CommitQuorumSetHash Hash      `json:"commitQuorumSetHash,omitempty"`
}

// ScpStatementPledges is an XDR NestedUnion defines as:
//
//   union switch (SCPStatementType type)
//        {
//        case PREPARE:
//            struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            } prepare;
//        case CONFIRM:
//            struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            } confirm;
//        case EXTERNALIZE:
//            struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            } externalize;
//        case NOMINATE:
//            SCPNomination nominate;
//        }
//
type ScpStatementPledges struct {
	Type        ScpStatementType         `json:"type,omitempty"`
	Prepare     *ScpStatementPrepare     `json:"prepare,omitempty"`
	Confirm     *ScpStatementConfirm     `json:"confirm,omitempty"`
	Externalize *ScpStatementExternalize `json:"externalize,omitempty"`
	Nominate    *ScpNomination           `json:"nominate,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ScpStatementPledges) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ScpStatementPledges
func (u ScpStatementPledges) ArmForSwitch(sw int32) (string, bool) {
	switch ScpStatementType(sw) {
	case ScpStatementTypePrepare:
		return "Prepare", true
	case ScpStatementTypeConfirm:
		return "Confirm", true
	case ScpStatementTypeExternalize:
		return "Externalize", true
	case ScpStatementTypeNominate:
		return "Nominate", true
	}
	return "-", false
}

// NewScpStatementPledges creates a new  ScpStatementPledges.
func NewScpStatementPledges(aType ScpStatementType, value interface{}) (result ScpStatementPledges, err error) {
	result.Type = aType
	switch ScpStatementType(aType) {
	case ScpStatementTypePrepare:
		tv, ok := value.(ScpStatementPrepare)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementPrepare")
			return
		}
		result.Prepare = &tv
	case ScpStatementTypeConfirm:
		tv, ok := value.(ScpStatementConfirm)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementConfirm")
			return
		}
		result.Confirm = &tv
	case ScpStatementTypeExternalize:
		tv, ok := value.(ScpStatementExternalize)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementExternalize")
			return
		}
		result.Externalize = &tv
	case ScpStatementTypeNominate:
		tv, ok := value.(ScpNomination)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpNomination")
			return
		}
		result.Nominate = &tv
	}
	return
}

// MustPrepare retrieves the Prepare value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustPrepare() ScpStatementPrepare {
	val, ok := u.GetPrepare()

	if !ok {
		panic("arm Prepare is not set")
	}

	return val
}

// GetPrepare retrieves the Prepare value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetPrepare() (result ScpStatementPrepare, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Prepare" {
		result = *u.Prepare
		ok = true
	}

	return
}

// MustConfirm retrieves the Confirm value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustConfirm() ScpStatementConfirm {
	val, ok := u.GetConfirm()

	if !ok {
		panic("arm Confirm is not set")
	}

	return val
}

// GetConfirm retrieves the Confirm value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetConfirm() (result ScpStatementConfirm, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Confirm" {
		result = *u.Confirm
		ok = true
	}

	return
}

// MustExternalize retrieves the Externalize value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustExternalize() ScpStatementExternalize {
	val, ok := u.GetExternalize()

	if !ok {
		panic("arm Externalize is not set")
	}

	return val
}

// GetExternalize retrieves the Externalize value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetExternalize() (result ScpStatementExternalize, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Externalize" {
		result = *u.Externalize
		ok = true
	}

	return
}

// MustNominate retrieves the Nominate value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustNominate() ScpNomination {
	val, ok := u.GetNominate()

	if !ok {
		panic("arm Nominate is not set")
	}

	return val
}

// GetNominate retrieves the Nominate value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetNominate() (result ScpNomination, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Nominate" {
		result = *u.Nominate
		ok = true
	}

	return
}

// ScpStatement is an XDR Struct defines as:
//
//   struct SCPStatement
//    {
//        NodeID nodeID;    // v
//        uint64 slotIndex; // i
//
//        union switch (SCPStatementType type)
//        {
//        case PREPARE:
//            struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            } prepare;
//        case CONFIRM:
//            struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            } confirm;
//        case EXTERNALIZE:
//            struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            } externalize;
//        case NOMINATE:
//            SCPNomination nominate;
//        }
//        pledges;
//    };
//
type ScpStatement struct {
	NodeId    NodeId              `json:"nodeID,omitempty"`
	SlotIndex Uint64              `json:"slotIndex,omitempty"`
	Pledges   ScpStatementPledges `json:"pledges,omitempty"`
}

// ScpEnvelope is an XDR Struct defines as:
//
//   struct SCPEnvelope
//    {
//        SCPStatement statement;
//        Signature signature;
//    };
//
type ScpEnvelope struct {
	Statement ScpStatement `json:"statement,omitempty"`
	Signature Signature    `json:"signature,omitempty"`
}

// ScpQuorumSet is an XDR Struct defines as:
//
//   struct SCPQuorumSet
//    {
//        uint32 threshold;
//        PublicKey validators<>;
//        SCPQuorumSet innerSets<>;
//    };
//
type ScpQuorumSet struct {
	Threshold  Uint32         `json:"threshold,omitempty"`
	Validators []PublicKey    `json:"validators,omitempty"`
	InnerSets  []ScpQuorumSet `json:"innerSets,omitempty"`
}

// AccountKycEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AccountKycEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountKycEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountKycEntryExt
func (u AccountKycEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAccountKycEntryExt creates a new  AccountKycEntryExt.
func NewAccountKycEntryExt(v LedgerVersion, value interface{}) (result AccountKycEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AccountKycEntry is an XDR Struct defines as:
//
//   struct AccountKYCEntry
//    {
//        AccountID accountID;
//        longstring KYCData;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AccountKycEntry struct {
	AccountId AccountId          `json:"accountID,omitempty"`
	KycData   Longstring         `json:"KYCData,omitempty"`
	Ext       AccountKycEntryExt `json:"ext,omitempty"`
}

// AccountRoleEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AccountRoleEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountRoleEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountRoleEntryExt
func (u AccountRoleEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAccountRoleEntryExt creates a new  AccountRoleEntryExt.
func NewAccountRoleEntryExt(v LedgerVersion, value interface{}) (result AccountRoleEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AccountRoleEntry is an XDR Struct defines as:
//
//   struct AccountRoleEntry
//    {
//        uint64 id;
//
//        uint64 ruleIDs<>;
//
//        longstring details;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AccountRoleEntry struct {
	Id      Uint64              `json:"id,omitempty"`
	RuleIDs []Uint64            `json:"ruleIDs,omitempty"`
	Details Longstring          `json:"details,omitempty"`
	Ext     AccountRoleEntryExt `json:"ext,omitempty"`
}

// AccountRuleEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AccountRuleEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountRuleEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountRuleEntryExt
func (u AccountRuleEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAccountRuleEntryExt creates a new  AccountRuleEntryExt.
func NewAccountRuleEntryExt(v LedgerVersion, value interface{}) (result AccountRuleEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AccountRuleEntry is an XDR Struct defines as:
//
//   struct AccountRuleEntry
//    {
//        uint64 id;
//
//        AccountRuleResource resource;
//        AccountRuleAction action;
//
//        bool forbids;
//
//        longstring details;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AccountRuleEntry struct {
	Id       Uint64              `json:"id,omitempty"`
	Resource AccountRuleResource `json:"resource,omitempty"`
	Action   AccountRuleAction   `json:"action,omitempty"`
	Forbids  bool                `json:"forbids,omitempty"`
	Details  Longstring          `json:"details,omitempty"`
	Ext      AccountRuleEntryExt `json:"ext,omitempty"`
}

// AccountEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AccountEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountEntryExt
func (u AccountEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAccountEntryExt creates a new  AccountEntryExt.
func NewAccountEntryExt(v LedgerVersion, value interface{}) (result AccountEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AccountEntry is an XDR Struct defines as:
//
//   struct AccountEntry
//    {
//        AccountID accountID;      // master public key for this account
//
//        // Referral marketing
//        AccountID* referrer; // parent account
//
//        // sequential ID - unique identifier of the account, used by ingesting applications to
//        // identify account, while keeping size of index small
//        uint64 sequentialID;
//
//    	uint64 roleID;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AccountEntry struct {
	AccountId    AccountId       `json:"accountID,omitempty"`
	Referrer     *AccountId      `json:"referrer,omitempty"`
	SequentialId Uint64          `json:"sequentialID,omitempty"`
	RoleId       Uint64          `json:"roleID,omitempty"`
	Ext          AccountEntryExt `json:"ext,omitempty"`
}

// KeyValueEntryType is an XDR Enum defines as:
//
//   //: `KeyValueEntryType` defines the type of value in the key-value entry
//        enum KeyValueEntryType
//        {
//            UINT32 = 1,
//            STRING = 2,
//            UINT64 = 3
//        };
//
type KeyValueEntryType int32

const (
	KeyValueEntryTypeUint32 KeyValueEntryType = 1
	KeyValueEntryTypeString KeyValueEntryType = 2
	KeyValueEntryTypeUint64 KeyValueEntryType = 3
)

var KeyValueEntryTypeAll = []KeyValueEntryType{
	KeyValueEntryTypeUint32,
	KeyValueEntryTypeString,
	KeyValueEntryTypeUint64,
}

var keyValueEntryTypeMap = map[int32]string{
	1: "KeyValueEntryTypeUint32",
	2: "KeyValueEntryTypeString",
	3: "KeyValueEntryTypeUint64",
}

var keyValueEntryTypeShortMap = map[int32]string{
	1: "uint32",
	2: "string",
	3: "uint64",
}

var keyValueEntryTypeRevMap = map[string]int32{
	"KeyValueEntryTypeUint32": 1,
	"KeyValueEntryTypeString": 2,
	"KeyValueEntryTypeUint64": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for KeyValueEntryType
func (e KeyValueEntryType) ValidEnum(v int32) bool {
	_, ok := keyValueEntryTypeMap[v]
	return ok
}
func (e KeyValueEntryType) isFlag() bool {
	for i := len(KeyValueEntryTypeAll) - 1; i >= 0; i-- {
		expected := KeyValueEntryType(2) << uint64(len(KeyValueEntryTypeAll)-1) >> uint64(len(KeyValueEntryTypeAll)-i)
		if expected != KeyValueEntryTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e KeyValueEntryType) String() string {
	name, _ := keyValueEntryTypeMap[int32(e)]
	return name
}

func (e KeyValueEntryType) ShortString() string {
	name, _ := keyValueEntryTypeShortMap[int32(e)]
	return name
}

func (e KeyValueEntryType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range KeyValueEntryTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *KeyValueEntryType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = KeyValueEntryType(t.Value)
	return nil
}

// KeyValueEntryValue is an XDR Union defines as:
//
//   //: `KeyValueEntryValue` represents the value based on given `KeyValueEntryType`
//        union KeyValueEntryValue switch (KeyValueEntryType type)
//        {
//            case UINT32:
//                uint32 ui32Value;
//            case STRING:
//                string stringValue<>;
//            case UINT64:
//                uint64 ui64Value;
//        };
//
type KeyValueEntryValue struct {
	Type        KeyValueEntryType `json:"type,omitempty"`
	Ui32Value   *Uint32           `json:"ui32Value,omitempty"`
	StringValue *string           `json:"stringValue,omitempty"`
	Ui64Value   *Uint64           `json:"ui64Value,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u KeyValueEntryValue) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of KeyValueEntryValue
func (u KeyValueEntryValue) ArmForSwitch(sw int32) (string, bool) {
	switch KeyValueEntryType(sw) {
	case KeyValueEntryTypeUint32:
		return "Ui32Value", true
	case KeyValueEntryTypeString:
		return "StringValue", true
	case KeyValueEntryTypeUint64:
		return "Ui64Value", true
	}
	return "-", false
}

// NewKeyValueEntryValue creates a new  KeyValueEntryValue.
func NewKeyValueEntryValue(aType KeyValueEntryType, value interface{}) (result KeyValueEntryValue, err error) {
	result.Type = aType
	switch KeyValueEntryType(aType) {
	case KeyValueEntryTypeUint32:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.Ui32Value = &tv
	case KeyValueEntryTypeString:
		tv, ok := value.(string)
		if !ok {
			err = fmt.Errorf("invalid value, must be string")
			return
		}
		result.StringValue = &tv
	case KeyValueEntryTypeUint64:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.Ui64Value = &tv
	}
	return
}

// MustUi32Value retrieves the Ui32Value value from the union,
// panicing if the value is not set.
func (u KeyValueEntryValue) MustUi32Value() Uint32 {
	val, ok := u.GetUi32Value()

	if !ok {
		panic("arm Ui32Value is not set")
	}

	return val
}

// GetUi32Value retrieves the Ui32Value value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u KeyValueEntryValue) GetUi32Value() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ui32Value" {
		result = *u.Ui32Value
		ok = true
	}

	return
}

// MustStringValue retrieves the StringValue value from the union,
// panicing if the value is not set.
func (u KeyValueEntryValue) MustStringValue() string {
	val, ok := u.GetStringValue()

	if !ok {
		panic("arm StringValue is not set")
	}

	return val
}

// GetStringValue retrieves the StringValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u KeyValueEntryValue) GetStringValue() (result string, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "StringValue" {
		result = *u.StringValue
		ok = true
	}

	return
}

// MustUi64Value retrieves the Ui64Value value from the union,
// panicing if the value is not set.
func (u KeyValueEntryValue) MustUi64Value() Uint64 {
	val, ok := u.GetUi64Value()

	if !ok {
		panic("arm Ui64Value is not set")
	}

	return val
}

// GetUi64Value retrieves the Ui64Value value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u KeyValueEntryValue) GetUi64Value() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ui64Value" {
		result = *u.Ui64Value
		ok = true
	}

	return
}

// KeyValueEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type KeyValueEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u KeyValueEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of KeyValueEntryExt
func (u KeyValueEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewKeyValueEntryExt creates a new  KeyValueEntryExt.
func NewKeyValueEntryExt(v LedgerVersion, value interface{}) (result KeyValueEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// KeyValueEntry is an XDR Struct defines as:
//
//   //: `KeyValueEntry` is an entry used to store key mapped values
//        struct KeyValueEntry
//        {
//            //: String value that must be unique among other keys for kev-value pairs
//            longstring key;
//
//            //: Value that corresponds to particular key (depending on `KeyValueEntryType`,
//            //: the value can be either uint32, or uint64, or string)
//            KeyValueEntryValue value;
//
//            //: reserved for future use
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        };
//
type KeyValueEntry struct {
	Key   Longstring         `json:"key,omitempty"`
	Value KeyValueEntryValue `json:"value,omitempty"`
	Ext   KeyValueEntryExt   `json:"ext,omitempty"`
}

// ReferenceEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ReferenceEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReferenceEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReferenceEntryExt
func (u ReferenceEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReferenceEntryExt creates a new  ReferenceEntryExt.
func NewReferenceEntryExt(v LedgerVersion, value interface{}) (result ReferenceEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReferenceEntry is an XDR Struct defines as:
//
//   struct ReferenceEntry
//    {
//    	AccountID sender;
//        string64 reference;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ReferenceEntry struct {
	Sender    AccountId         `json:"sender,omitempty"`
	Reference String64          `json:"reference,omitempty"`
	Ext       ReferenceEntryExt `json:"ext,omitempty"`
}

// ReviewableRequestType is an XDR Enum defines as:
//
//   enum ReviewableRequestType
//    {
//    	NONE = 0, // use this request type in ReviewRequestOp extended result if additional info is not required
//    	ANY = 1,
//    	CHANGE_ROLE = 2,
//    	KYC_RECOVERY = 3
//    };
//
type ReviewableRequestType int32

const (
	ReviewableRequestTypeNone        ReviewableRequestType = 0
	ReviewableRequestTypeAny         ReviewableRequestType = 1
	ReviewableRequestTypeChangeRole  ReviewableRequestType = 2
	ReviewableRequestTypeKycRecovery ReviewableRequestType = 3
)

var ReviewableRequestTypeAll = []ReviewableRequestType{
	ReviewableRequestTypeNone,
	ReviewableRequestTypeAny,
	ReviewableRequestTypeChangeRole,
	ReviewableRequestTypeKycRecovery,
}

var reviewableRequestTypeMap = map[int32]string{
	0: "ReviewableRequestTypeNone",
	1: "ReviewableRequestTypeAny",
	2: "ReviewableRequestTypeChangeRole",
	3: "ReviewableRequestTypeKycRecovery",
}

var reviewableRequestTypeShortMap = map[int32]string{
	0: "none",
	1: "any",
	2: "change_role",
	3: "kyc_recovery",
}

var reviewableRequestTypeRevMap = map[string]int32{
	"ReviewableRequestTypeNone":        0,
	"ReviewableRequestTypeAny":         1,
	"ReviewableRequestTypeChangeRole":  2,
	"ReviewableRequestTypeKycRecovery": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewableRequestType
func (e ReviewableRequestType) ValidEnum(v int32) bool {
	_, ok := reviewableRequestTypeMap[v]
	return ok
}
func (e ReviewableRequestType) isFlag() bool {
	for i := len(ReviewableRequestTypeAll) - 1; i >= 0; i-- {
		expected := ReviewableRequestType(2) << uint64(len(ReviewableRequestTypeAll)-1) >> uint64(len(ReviewableRequestTypeAll)-i)
		if expected != ReviewableRequestTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewableRequestType) String() string {
	name, _ := reviewableRequestTypeMap[int32(e)]
	return name
}

func (e ReviewableRequestType) ShortString() string {
	name, _ := reviewableRequestTypeShortMap[int32(e)]
	return name
}

func (e ReviewableRequestType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ReviewableRequestTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ReviewableRequestType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewableRequestType(t.Value)
	return nil
}

// TasksExtExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TasksExtExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TasksExtExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TasksExtExt
func (u TasksExtExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTasksExtExt creates a new  TasksExtExt.
func NewTasksExtExt(v LedgerVersion, value interface{}) (result TasksExtExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TasksExt is an XDR Struct defines as:
//
//   struct TasksExt {
//        // Tasks are represented by a bitmask
//        uint32 allTasks;
//        uint32 pendingTasks;
//
//        // External details vector consists of comments written by request reviewers
//        longstring externalDetails<>;
//
//        // Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TasksExt struct {
	AllTasks        Uint32       `json:"allTasks,omitempty"`
	PendingTasks    Uint32       `json:"pendingTasks,omitempty"`
	ExternalDetails []Longstring `json:"externalDetails,omitempty"`
	Ext             TasksExtExt  `json:"ext,omitempty"`
}

// ReviewableRequestEntryBody is an XDR NestedUnion defines as:
//
//   union switch (ReviewableRequestType type) {
//            case CHANGE_ROLE:
//                ChangeRoleRequest changeRoleRequest;
//            case KYC_RECOVERY:
//                KYCRecoveryRequest kycRecoveryRequest;
//    	}
//
type ReviewableRequestEntryBody struct {
	Type               ReviewableRequestType `json:"type,omitempty"`
	ChangeRoleRequest  *ChangeRoleRequest    `json:"changeRoleRequest,omitempty"`
	KycRecoveryRequest *KycRecoveryRequest   `json:"kycRecoveryRequest,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewableRequestEntryBody) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewableRequestEntryBody
func (u ReviewableRequestEntryBody) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewableRequestType(sw) {
	case ReviewableRequestTypeChangeRole:
		return "ChangeRoleRequest", true
	case ReviewableRequestTypeKycRecovery:
		return "KycRecoveryRequest", true
	}
	return "-", false
}

// NewReviewableRequestEntryBody creates a new  ReviewableRequestEntryBody.
func NewReviewableRequestEntryBody(aType ReviewableRequestType, value interface{}) (result ReviewableRequestEntryBody, err error) {
	result.Type = aType
	switch ReviewableRequestType(aType) {
	case ReviewableRequestTypeChangeRole:
		tv, ok := value.(ChangeRoleRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be ChangeRoleRequest")
			return
		}
		result.ChangeRoleRequest = &tv
	case ReviewableRequestTypeKycRecovery:
		tv, ok := value.(KycRecoveryRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be KycRecoveryRequest")
			return
		}
		result.KycRecoveryRequest = &tv
	}
	return
}

// MustChangeRoleRequest retrieves the ChangeRoleRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustChangeRoleRequest() ChangeRoleRequest {
	val, ok := u.GetChangeRoleRequest()

	if !ok {
		panic("arm ChangeRoleRequest is not set")
	}

	return val
}

// GetChangeRoleRequest retrieves the ChangeRoleRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetChangeRoleRequest() (result ChangeRoleRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChangeRoleRequest" {
		result = *u.ChangeRoleRequest
		ok = true
	}

	return
}

// MustKycRecoveryRequest retrieves the KycRecoveryRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustKycRecoveryRequest() KycRecoveryRequest {
	val, ok := u.GetKycRecoveryRequest()

	if !ok {
		panic("arm KycRecoveryRequest is not set")
	}

	return val
}

// GetKycRecoveryRequest retrieves the KycRecoveryRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetKycRecoveryRequest() (result KycRecoveryRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KycRecoveryRequest" {
		result = *u.KycRecoveryRequest
		ok = true
	}

	return
}

// ReviewableRequestEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	}
//
type ReviewableRequestEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewableRequestEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewableRequestEntryExt
func (u ReviewableRequestEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewableRequestEntryExt creates a new  ReviewableRequestEntryExt.
func NewReviewableRequestEntryExt(v LedgerVersion, value interface{}) (result ReviewableRequestEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewableRequestEntry is an XDR Struct defines as:
//
//   struct ReviewableRequestEntry {
//    	uint64 requestID;
//    	Hash hash; // hash of the request body
//    	AccountID requestor;
//    	longstring rejectReason;
//    	AccountID reviewer;
//    	string64* reference; // reference for request which will act as an unique key for the request (will reject request with the same reference from same requestor)
//    	int64 createdAt; // when request was created
//
//    	union switch (ReviewableRequestType type) {
//            case CHANGE_ROLE:
//                ChangeRoleRequest changeRoleRequest;
//            case KYC_RECOVERY:
//                KYCRecoveryRequest kycRecoveryRequest;
//    	} body;
//
//    	TasksExt tasks;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	}
//        ext;
//    };
//
type ReviewableRequestEntry struct {
	RequestId    Uint64                     `json:"requestID,omitempty"`
	Hash         Hash                       `json:"hash,omitempty"`
	Requestor    AccountId                  `json:"requestor,omitempty"`
	RejectReason Longstring                 `json:"rejectReason,omitempty"`
	Reviewer     AccountId                  `json:"reviewer,omitempty"`
	Reference    *String64                  `json:"reference,omitempty"`
	CreatedAt    Int64                      `json:"createdAt,omitempty"`
	Body         ReviewableRequestEntryBody `json:"body,omitempty"`
	Tasks        TasksExt                   `json:"tasks,omitempty"`
	Ext          ReviewableRequestEntryExt  `json:"ext,omitempty"`
}

// SignerRoleEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SignerRoleEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SignerRoleEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SignerRoleEntryExt
func (u SignerRoleEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSignerRoleEntryExt creates a new  SignerRoleEntryExt.
func NewSignerRoleEntryExt(v LedgerVersion, value interface{}) (result SignerRoleEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SignerRoleEntry is an XDR Struct defines as:
//
//   struct SignerRoleEntry
//    {
//        uint64 id;
//        uint64 ruleIDs<>;
//
//        AccountID ownerID;
//
//        longstring details;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SignerRoleEntry struct {
	Id      Uint64             `json:"id,omitempty"`
	RuleIDs []Uint64           `json:"ruleIDs,omitempty"`
	OwnerId AccountId          `json:"ownerID,omitempty"`
	Details Longstring         `json:"details,omitempty"`
	Ext     SignerRoleEntryExt `json:"ext,omitempty"`
}

// SignerRuleEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SignerRuleEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SignerRuleEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SignerRuleEntryExt
func (u SignerRuleEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSignerRuleEntryExt creates a new  SignerRuleEntryExt.
func NewSignerRuleEntryExt(v LedgerVersion, value interface{}) (result SignerRuleEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SignerRuleEntry is an XDR Struct defines as:
//
//   struct SignerRuleEntry
//    {
//        uint64 id;
//
//        SignerRuleResource resource;
//        SignerRuleAction action;
//
//        bool forbids;
//        bool isDefault; // default rules will be in each role
//
//        longstring details;
//
//        AccountID ownerID;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SignerRuleEntry struct {
	Id        Uint64             `json:"id,omitempty"`
	Resource  SignerRuleResource `json:"resource,omitempty"`
	Action    SignerRuleAction   `json:"action,omitempty"`
	Forbids   bool               `json:"forbids,omitempty"`
	IsDefault bool               `json:"isDefault,omitempty"`
	Details   Longstring         `json:"details,omitempty"`
	OwnerId   AccountId          `json:"ownerID,omitempty"`
	Ext       SignerRuleEntryExt `json:"ext,omitempty"`
}

// SignerEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SignerEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SignerEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SignerEntryExt
func (u SignerEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSignerEntryExt creates a new  SignerEntryExt.
func NewSignerEntryExt(v LedgerVersion, value interface{}) (result SignerEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SignerEntry is an XDR Struct defines as:
//
//   struct SignerEntry
//    {
//        PublicKey pubKey;
//        AccountID accountID; // account to which signer had attached
//
//        uint32 weight; // threshold for all SignerRules equals 1000
//    	uint32 identity;
//
//    	longstring details;
//
//    	uint64 roleID;
//
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SignerEntry struct {
	PubKey    PublicKey      `json:"pubKey,omitempty"`
	AccountId AccountId      `json:"accountID,omitempty"`
	Weight    Uint32         `json:"weight,omitempty"`
	Identity  Uint32         `json:"identity,omitempty"`
	Details   Longstring     `json:"details,omitempty"`
	RoleId    Uint64         `json:"roleID,omitempty"`
	Ext       SignerEntryExt `json:"ext,omitempty"`
}

// LedgerEntryData is an XDR NestedUnion defines as:
//
//   union switch (LedgerEntryType type)
//        {
//        case ACCOUNT:
//            AccountEntry account;
//        case SIGNER:
//            SignerEntry signer;
//        case REFERENCE_ENTRY:
//            ReferenceEntry reference;
//    	case REVIEWABLE_REQUEST:
//    		ReviewableRequestEntry reviewableRequest;
//    	case KEY_VALUE:
//    	    KeyValueEntry keyValue;
//    	case ACCOUNT_KYC:
//            AccountKYCEntry accountKYC;
//        case ACCOUNT_ROLE:
//            AccountRoleEntry accountRole;
//        case ACCOUNT_RULE:
//            AccountRuleEntry accountRule;
//        case SIGNER_RULE:
//            SignerRuleEntry signerRule;
//        case SIGNER_ROLE:
//            SignerRoleEntry signerRole;
//        }
//
type LedgerEntryData struct {
	Type              LedgerEntryType         `json:"type,omitempty"`
	Account           *AccountEntry           `json:"account,omitempty"`
	Signer            *SignerEntry            `json:"signer,omitempty"`
	Reference         *ReferenceEntry         `json:"reference,omitempty"`
	ReviewableRequest *ReviewableRequestEntry `json:"reviewableRequest,omitempty"`
	KeyValue          *KeyValueEntry          `json:"keyValue,omitempty"`
	AccountKyc        *AccountKycEntry        `json:"accountKYC,omitempty"`
	AccountRole       *AccountRoleEntry       `json:"accountRole,omitempty"`
	AccountRule       *AccountRuleEntry       `json:"accountRule,omitempty"`
	SignerRule        *SignerRuleEntry        `json:"signerRule,omitempty"`
	SignerRole        *SignerRoleEntry        `json:"signerRole,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryData) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryData
func (u LedgerEntryData) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeAccount:
		return "Account", true
	case LedgerEntryTypeSigner:
		return "Signer", true
	case LedgerEntryTypeReferenceEntry:
		return "Reference", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeKeyValue:
		return "KeyValue", true
	case LedgerEntryTypeAccountKyc:
		return "AccountKyc", true
	case LedgerEntryTypeAccountRole:
		return "AccountRole", true
	case LedgerEntryTypeAccountRule:
		return "AccountRule", true
	case LedgerEntryTypeSignerRule:
		return "SignerRule", true
	case LedgerEntryTypeSignerRole:
		return "SignerRole", true
	}
	return "-", false
}

// NewLedgerEntryData creates a new  LedgerEntryData.
func NewLedgerEntryData(aType LedgerEntryType, value interface{}) (result LedgerEntryData, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeAccount:
		tv, ok := value.(AccountEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountEntry")
			return
		}
		result.Account = &tv
	case LedgerEntryTypeSigner:
		tv, ok := value.(SignerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerEntry")
			return
		}
		result.Signer = &tv
	case LedgerEntryTypeReferenceEntry:
		tv, ok := value.(ReferenceEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReferenceEntry")
			return
		}
		result.Reference = &tv
	case LedgerEntryTypeReviewableRequest:
		tv, ok := value.(ReviewableRequestEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewableRequestEntry")
			return
		}
		result.ReviewableRequest = &tv
	case LedgerEntryTypeKeyValue:
		tv, ok := value.(KeyValueEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be KeyValueEntry")
			return
		}
		result.KeyValue = &tv
	case LedgerEntryTypeAccountKyc:
		tv, ok := value.(AccountKycEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountKycEntry")
			return
		}
		result.AccountKyc = &tv
	case LedgerEntryTypeAccountRole:
		tv, ok := value.(AccountRoleEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountRoleEntry")
			return
		}
		result.AccountRole = &tv
	case LedgerEntryTypeAccountRule:
		tv, ok := value.(AccountRuleEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountRuleEntry")
			return
		}
		result.AccountRule = &tv
	case LedgerEntryTypeSignerRule:
		tv, ok := value.(SignerRuleEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerRuleEntry")
			return
		}
		result.SignerRule = &tv
	case LedgerEntryTypeSignerRole:
		tv, ok := value.(SignerRoleEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerRoleEntry")
			return
		}
		result.SignerRole = &tv
	}
	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccount() AccountEntry {
	val, ok := u.GetAccount()

	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccount() (result AccountEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MustSigner retrieves the Signer value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustSigner() SignerEntry {
	val, ok := u.GetSigner()

	if !ok {
		panic("arm Signer is not set")
	}

	return val
}

// GetSigner retrieves the Signer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetSigner() (result SignerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Signer" {
		result = *u.Signer
		ok = true
	}

	return
}

// MustReference retrieves the Reference value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustReference() ReferenceEntry {
	val, ok := u.GetReference()

	if !ok {
		panic("arm Reference is not set")
	}

	return val
}

// GetReference retrieves the Reference value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetReference() (result ReferenceEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Reference" {
		result = *u.Reference
		ok = true
	}

	return
}

// MustReviewableRequest retrieves the ReviewableRequest value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustReviewableRequest() ReviewableRequestEntry {
	val, ok := u.GetReviewableRequest()

	if !ok {
		panic("arm ReviewableRequest is not set")
	}

	return val
}

// GetReviewableRequest retrieves the ReviewableRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetReviewableRequest() (result ReviewableRequestEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewableRequest" {
		result = *u.ReviewableRequest
		ok = true
	}

	return
}

// MustKeyValue retrieves the KeyValue value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustKeyValue() KeyValueEntry {
	val, ok := u.GetKeyValue()

	if !ok {
		panic("arm KeyValue is not set")
	}

	return val
}

// GetKeyValue retrieves the KeyValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetKeyValue() (result KeyValueEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KeyValue" {
		result = *u.KeyValue
		ok = true
	}

	return
}

// MustAccountKyc retrieves the AccountKyc value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccountKyc() AccountKycEntry {
	val, ok := u.GetAccountKyc()

	if !ok {
		panic("arm AccountKyc is not set")
	}

	return val
}

// GetAccountKyc retrieves the AccountKyc value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccountKyc() (result AccountKycEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountKyc" {
		result = *u.AccountKyc
		ok = true
	}

	return
}

// MustAccountRole retrieves the AccountRole value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccountRole() AccountRoleEntry {
	val, ok := u.GetAccountRole()

	if !ok {
		panic("arm AccountRole is not set")
	}

	return val
}

// GetAccountRole retrieves the AccountRole value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccountRole() (result AccountRoleEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountRole" {
		result = *u.AccountRole
		ok = true
	}

	return
}

// MustAccountRule retrieves the AccountRule value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccountRule() AccountRuleEntry {
	val, ok := u.GetAccountRule()

	if !ok {
		panic("arm AccountRule is not set")
	}

	return val
}

// GetAccountRule retrieves the AccountRule value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccountRule() (result AccountRuleEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountRule" {
		result = *u.AccountRule
		ok = true
	}

	return
}

// MustSignerRule retrieves the SignerRule value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustSignerRule() SignerRuleEntry {
	val, ok := u.GetSignerRule()

	if !ok {
		panic("arm SignerRule is not set")
	}

	return val
}

// GetSignerRule retrieves the SignerRule value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetSignerRule() (result SignerRuleEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SignerRule" {
		result = *u.SignerRule
		ok = true
	}

	return
}

// MustSignerRole retrieves the SignerRole value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustSignerRole() SignerRoleEntry {
	val, ok := u.GetSignerRole()

	if !ok {
		panic("arm SignerRole is not set")
	}

	return val
}

// GetSignerRole retrieves the SignerRole value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetSignerRole() (result SignerRoleEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SignerRole" {
		result = *u.SignerRole
		ok = true
	}

	return
}

// LedgerEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LedgerEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryExt
func (u LedgerEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerEntryExt creates a new  LedgerEntryExt.
func NewLedgerEntryExt(v LedgerVersion, value interface{}) (result LedgerEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerEntry is an XDR Struct defines as:
//
//   struct LedgerEntry
//    {
//        uint32 lastModifiedLedgerSeq; // ledger the LedgerEntry was last changed
//
//        union switch (LedgerEntryType type)
//        {
//        case ACCOUNT:
//            AccountEntry account;
//        case SIGNER:
//            SignerEntry signer;
//        case REFERENCE_ENTRY:
//            ReferenceEntry reference;
//    	case REVIEWABLE_REQUEST:
//    		ReviewableRequestEntry reviewableRequest;
//    	case KEY_VALUE:
//    	    KeyValueEntry keyValue;
//    	case ACCOUNT_KYC:
//            AccountKYCEntry accountKYC;
//        case ACCOUNT_ROLE:
//            AccountRoleEntry accountRole;
//        case ACCOUNT_RULE:
//            AccountRuleEntry accountRule;
//        case SIGNER_RULE:
//            SignerRuleEntry signerRule;
//        case SIGNER_ROLE:
//            SignerRoleEntry signerRole;
//        }
//        data;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LedgerEntry struct {
	LastModifiedLedgerSeq Uint32          `json:"lastModifiedLedgerSeq,omitempty"`
	Data                  LedgerEntryData `json:"data,omitempty"`
	Ext                   LedgerEntryExt  `json:"ext,omitempty"`
}

// EnvelopeType is an XDR Enum defines as:
//
//   enum EnvelopeType
//    {
//        SCP = 1,
//        TX = 2,
//        AUTH = 3
//    };
//
type EnvelopeType int32

const (
	EnvelopeTypeScp  EnvelopeType = 1
	EnvelopeTypeTx   EnvelopeType = 2
	EnvelopeTypeAuth EnvelopeType = 3
)

var EnvelopeTypeAll = []EnvelopeType{
	EnvelopeTypeScp,
	EnvelopeTypeTx,
	EnvelopeTypeAuth,
}

var envelopeTypeMap = map[int32]string{
	1: "EnvelopeTypeScp",
	2: "EnvelopeTypeTx",
	3: "EnvelopeTypeAuth",
}

var envelopeTypeShortMap = map[int32]string{
	1: "scp",
	2: "tx",
	3: "auth",
}

var envelopeTypeRevMap = map[string]int32{
	"EnvelopeTypeScp":  1,
	"EnvelopeTypeTx":   2,
	"EnvelopeTypeAuth": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for EnvelopeType
func (e EnvelopeType) ValidEnum(v int32) bool {
	_, ok := envelopeTypeMap[v]
	return ok
}
func (e EnvelopeType) isFlag() bool {
	for i := len(EnvelopeTypeAll) - 1; i >= 0; i-- {
		expected := EnvelopeType(2) << uint64(len(EnvelopeTypeAll)-1) >> uint64(len(EnvelopeTypeAll)-i)
		if expected != EnvelopeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e EnvelopeType) String() string {
	name, _ := envelopeTypeMap[int32(e)]
	return name
}

func (e EnvelopeType) ShortString() string {
	name, _ := envelopeTypeShortMap[int32(e)]
	return name
}

func (e EnvelopeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range EnvelopeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *EnvelopeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = EnvelopeType(t.Value)
	return nil
}

// LedgerKeyAccountExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//           {
//           case EMPTY_VERSION:
//              void;
//           }
//
type LedgerKeyAccountExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAccountExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAccountExt
func (u LedgerKeyAccountExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAccountExt creates a new  LedgerKeyAccountExt.
func NewLedgerKeyAccountExt(v LedgerVersion, value interface{}) (result LedgerKeyAccountExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAccount is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID accountID;
//            union switch (LedgerVersion v)
//           {
//           case EMPTY_VERSION:
//              void;
//           }
//           ext;
//        }
//
type LedgerKeyAccount struct {
	AccountId AccountId           `json:"accountID,omitempty"`
	Ext       LedgerKeyAccountExt `json:"ext,omitempty"`
}

// LedgerKeySignerExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeySignerExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeySignerExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeySignerExt
func (u LedgerKeySignerExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeySignerExt creates a new  LedgerKeySignerExt.
func NewLedgerKeySignerExt(v LedgerVersion, value interface{}) (result LedgerKeySignerExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeySigner is an XDR NestedStruct defines as:
//
//   struct
//        {
//            PublicKey pubKey;
//            AccountID accountID;
//
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeySigner struct {
	PubKey    PublicKey          `json:"pubKey,omitempty"`
	AccountId AccountId          `json:"accountID,omitempty"`
	Ext       LedgerKeySignerExt `json:"ext,omitempty"`
}

// LedgerKeyReferenceExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyReferenceExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyReferenceExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyReferenceExt
func (u LedgerKeyReferenceExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyReferenceExt creates a new  LedgerKeyReferenceExt.
func NewLedgerKeyReferenceExt(v LedgerVersion, value interface{}) (result LedgerKeyReferenceExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyReference is an XDR NestedStruct defines as:
//
//   struct
//        {
//    		AccountID sender;
//    		string64 reference;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyReference struct {
	Sender    AccountId             `json:"sender,omitempty"`
	Reference String64              `json:"reference,omitempty"`
	Ext       LedgerKeyReferenceExt `json:"ext,omitempty"`
}

// LedgerKeyReviewableRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyReviewableRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyReviewableRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyReviewableRequestExt
func (u LedgerKeyReviewableRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyReviewableRequestExt creates a new  LedgerKeyReviewableRequestExt.
func NewLedgerKeyReviewableRequestExt(v LedgerVersion, value interface{}) (result LedgerKeyReviewableRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyReviewableRequest is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 requestID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyReviewableRequest struct {
	RequestId Uint64                        `json:"requestID,omitempty"`
	Ext       LedgerKeyReviewableRequestExt `json:"ext,omitempty"`
}

// LedgerKeyKeyValueExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyKeyValueExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyKeyValueExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyKeyValueExt
func (u LedgerKeyKeyValueExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyKeyValueExt creates a new  LedgerKeyKeyValueExt.
func NewLedgerKeyKeyValueExt(v LedgerVersion, value interface{}) (result LedgerKeyKeyValueExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyKeyValue is an XDR NestedStruct defines as:
//
//   struct {
//            longstring key;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyKeyValue struct {
	Key Longstring           `json:"key,omitempty"`
	Ext LedgerKeyKeyValueExt `json:"ext,omitempty"`
}

// LedgerKeyAccountKycExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyAccountKycExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAccountKycExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAccountKycExt
func (u LedgerKeyAccountKycExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAccountKycExt creates a new  LedgerKeyAccountKycExt.
func NewLedgerKeyAccountKycExt(v LedgerVersion, value interface{}) (result LedgerKeyAccountKycExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAccountKyc is an XDR NestedStruct defines as:
//
//   struct {
//            AccountID accountID;
//            union switch(LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyAccountKyc struct {
	AccountId AccountId              `json:"accountID,omitempty"`
	Ext       LedgerKeyAccountKycExt `json:"ext,omitempty"`
}

// LedgerKeyAccountRoleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyAccountRoleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAccountRoleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAccountRoleExt
func (u LedgerKeyAccountRoleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAccountRoleExt creates a new  LedgerKeyAccountRoleExt.
func NewLedgerKeyAccountRoleExt(v LedgerVersion, value interface{}) (result LedgerKeyAccountRoleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAccountRole is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyAccountRole struct {
	Id  Uint64                  `json:"id,omitempty"`
	Ext LedgerKeyAccountRoleExt `json:"ext,omitempty"`
}

// LedgerKeyAccountRuleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyAccountRuleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAccountRuleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAccountRuleExt
func (u LedgerKeyAccountRuleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAccountRuleExt creates a new  LedgerKeyAccountRuleExt.
func NewLedgerKeyAccountRuleExt(v LedgerVersion, value interface{}) (result LedgerKeyAccountRuleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAccountRule is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyAccountRule struct {
	Id  Uint64                  `json:"id,omitempty"`
	Ext LedgerKeyAccountRuleExt `json:"ext,omitempty"`
}

// LedgerKeySignerRoleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeySignerRoleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeySignerRoleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeySignerRoleExt
func (u LedgerKeySignerRoleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeySignerRoleExt creates a new  LedgerKeySignerRoleExt.
func NewLedgerKeySignerRoleExt(v LedgerVersion, value interface{}) (result LedgerKeySignerRoleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeySignerRole is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeySignerRole struct {
	Id  Uint64                 `json:"id,omitempty"`
	Ext LedgerKeySignerRoleExt `json:"ext,omitempty"`
}

// LedgerKeySignerRuleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeySignerRuleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeySignerRuleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeySignerRuleExt
func (u LedgerKeySignerRuleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeySignerRuleExt creates a new  LedgerKeySignerRuleExt.
func NewLedgerKeySignerRuleExt(v LedgerVersion, value interface{}) (result LedgerKeySignerRuleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeySignerRule is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeySignerRule struct {
	Id  Uint64                 `json:"id,omitempty"`
	Ext LedgerKeySignerRuleExt `json:"ext,omitempty"`
}

// LedgerKey is an XDR Union defines as:
//
//   union LedgerKey switch (LedgerEntryType type)
//    {
//    case ACCOUNT:
//        struct
//        {
//            AccountID accountID;
//            union switch (LedgerVersion v)
//           {
//           case EMPTY_VERSION:
//              void;
//           }
//           ext;
//        } account;
//    case SIGNER:
//        struct
//        {
//            PublicKey pubKey;
//            AccountID accountID;
//
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } signer;
//    case REFERENCE_ENTRY:
//        struct
//        {
//    		AccountID sender;
//    		string64 reference;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } reference;
//    case REVIEWABLE_REQUEST:
//        struct {
//            uint64 requestID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } reviewableRequest;
//    case KEY_VALUE:
//        struct {
//            longstring key;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } keyValue;
//    case ACCOUNT_KYC:
//        struct {
//            AccountID accountID;
//            union switch(LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } accountKYC;
//    case ACCOUNT_ROLE:
//        struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } accountRole;
//    case ACCOUNT_RULE:
//        struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } accountRule;
//    case SIGNER_ROLE:
//        struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } signerRole;
//    case SIGNER_RULE:
//        struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } signerRule;
//    };
//
type LedgerKey struct {
	Type              LedgerEntryType             `json:"type,omitempty"`
	Account           *LedgerKeyAccount           `json:"account,omitempty"`
	Signer            *LedgerKeySigner            `json:"signer,omitempty"`
	Reference         *LedgerKeyReference         `json:"reference,omitempty"`
	ReviewableRequest *LedgerKeyReviewableRequest `json:"reviewableRequest,omitempty"`
	KeyValue          *LedgerKeyKeyValue          `json:"keyValue,omitempty"`
	AccountKyc        *LedgerKeyAccountKyc        `json:"accountKYC,omitempty"`
	AccountRole       *LedgerKeyAccountRole       `json:"accountRole,omitempty"`
	AccountRule       *LedgerKeyAccountRule       `json:"accountRule,omitempty"`
	SignerRole        *LedgerKeySignerRole        `json:"signerRole,omitempty"`
	SignerRule        *LedgerKeySignerRule        `json:"signerRule,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKey
func (u LedgerKey) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeAccount:
		return "Account", true
	case LedgerEntryTypeSigner:
		return "Signer", true
	case LedgerEntryTypeReferenceEntry:
		return "Reference", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeKeyValue:
		return "KeyValue", true
	case LedgerEntryTypeAccountKyc:
		return "AccountKyc", true
	case LedgerEntryTypeAccountRole:
		return "AccountRole", true
	case LedgerEntryTypeAccountRule:
		return "AccountRule", true
	case LedgerEntryTypeSignerRole:
		return "SignerRole", true
	case LedgerEntryTypeSignerRule:
		return "SignerRule", true
	}
	return "-", false
}

// NewLedgerKey creates a new  LedgerKey.
func NewLedgerKey(aType LedgerEntryType, value interface{}) (result LedgerKey, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeAccount:
		tv, ok := value.(LedgerKeyAccount)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccount")
			return
		}
		result.Account = &tv
	case LedgerEntryTypeSigner:
		tv, ok := value.(LedgerKeySigner)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeySigner")
			return
		}
		result.Signer = &tv
	case LedgerEntryTypeReferenceEntry:
		tv, ok := value.(LedgerKeyReference)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyReference")
			return
		}
		result.Reference = &tv
	case LedgerEntryTypeReviewableRequest:
		tv, ok := value.(LedgerKeyReviewableRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyReviewableRequest")
			return
		}
		result.ReviewableRequest = &tv
	case LedgerEntryTypeKeyValue:
		tv, ok := value.(LedgerKeyKeyValue)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyKeyValue")
			return
		}
		result.KeyValue = &tv
	case LedgerEntryTypeAccountKyc:
		tv, ok := value.(LedgerKeyAccountKyc)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccountKyc")
			return
		}
		result.AccountKyc = &tv
	case LedgerEntryTypeAccountRole:
		tv, ok := value.(LedgerKeyAccountRole)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccountRole")
			return
		}
		result.AccountRole = &tv
	case LedgerEntryTypeAccountRule:
		tv, ok := value.(LedgerKeyAccountRule)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccountRule")
			return
		}
		result.AccountRule = &tv
	case LedgerEntryTypeSignerRole:
		tv, ok := value.(LedgerKeySignerRole)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeySignerRole")
			return
		}
		result.SignerRole = &tv
	case LedgerEntryTypeSignerRule:
		tv, ok := value.(LedgerKeySignerRule)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeySignerRule")
			return
		}
		result.SignerRule = &tv
	}
	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccount() LedgerKeyAccount {
	val, ok := u.GetAccount()

	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccount() (result LedgerKeyAccount, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MustSigner retrieves the Signer value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustSigner() LedgerKeySigner {
	val, ok := u.GetSigner()

	if !ok {
		panic("arm Signer is not set")
	}

	return val
}

// GetSigner retrieves the Signer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetSigner() (result LedgerKeySigner, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Signer" {
		result = *u.Signer
		ok = true
	}

	return
}

// MustReference retrieves the Reference value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustReference() LedgerKeyReference {
	val, ok := u.GetReference()

	if !ok {
		panic("arm Reference is not set")
	}

	return val
}

// GetReference retrieves the Reference value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetReference() (result LedgerKeyReference, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Reference" {
		result = *u.Reference
		ok = true
	}

	return
}

// MustReviewableRequest retrieves the ReviewableRequest value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustReviewableRequest() LedgerKeyReviewableRequest {
	val, ok := u.GetReviewableRequest()

	if !ok {
		panic("arm ReviewableRequest is not set")
	}

	return val
}

// GetReviewableRequest retrieves the ReviewableRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetReviewableRequest() (result LedgerKeyReviewableRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewableRequest" {
		result = *u.ReviewableRequest
		ok = true
	}

	return
}

// MustKeyValue retrieves the KeyValue value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustKeyValue() LedgerKeyKeyValue {
	val, ok := u.GetKeyValue()

	if !ok {
		panic("arm KeyValue is not set")
	}

	return val
}

// GetKeyValue retrieves the KeyValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetKeyValue() (result LedgerKeyKeyValue, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KeyValue" {
		result = *u.KeyValue
		ok = true
	}

	return
}

// MustAccountKyc retrieves the AccountKyc value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccountKyc() LedgerKeyAccountKyc {
	val, ok := u.GetAccountKyc()

	if !ok {
		panic("arm AccountKyc is not set")
	}

	return val
}

// GetAccountKyc retrieves the AccountKyc value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccountKyc() (result LedgerKeyAccountKyc, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountKyc" {
		result = *u.AccountKyc
		ok = true
	}

	return
}

// MustAccountRole retrieves the AccountRole value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccountRole() LedgerKeyAccountRole {
	val, ok := u.GetAccountRole()

	if !ok {
		panic("arm AccountRole is not set")
	}

	return val
}

// GetAccountRole retrieves the AccountRole value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccountRole() (result LedgerKeyAccountRole, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountRole" {
		result = *u.AccountRole
		ok = true
	}

	return
}

// MustAccountRule retrieves the AccountRule value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccountRule() LedgerKeyAccountRule {
	val, ok := u.GetAccountRule()

	if !ok {
		panic("arm AccountRule is not set")
	}

	return val
}

// GetAccountRule retrieves the AccountRule value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccountRule() (result LedgerKeyAccountRule, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountRule" {
		result = *u.AccountRule
		ok = true
	}

	return
}

// MustSignerRole retrieves the SignerRole value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustSignerRole() LedgerKeySignerRole {
	val, ok := u.GetSignerRole()

	if !ok {
		panic("arm SignerRole is not set")
	}

	return val
}

// GetSignerRole retrieves the SignerRole value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetSignerRole() (result LedgerKeySignerRole, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SignerRole" {
		result = *u.SignerRole
		ok = true
	}

	return
}

// MustSignerRule retrieves the SignerRule value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustSignerRule() LedgerKeySignerRule {
	val, ok := u.GetSignerRule()

	if !ok {
		panic("arm SignerRule is not set")
	}

	return val
}

// GetSignerRule retrieves the SignerRule value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetSignerRule() (result LedgerKeySignerRule, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SignerRule" {
		result = *u.SignerRule
		ok = true
	}

	return
}

// UpgradeType is an XDR Typedef defines as:
//
//   typedef opaque UpgradeType<128>;
//
type UpgradeType []byte

// StellarValueExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type StellarValueExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StellarValueExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StellarValueExt
func (u StellarValueExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewStellarValueExt creates a new  StellarValueExt.
func NewStellarValueExt(v LedgerVersion, value interface{}) (result StellarValueExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// StellarValue is an XDR Struct defines as:
//
//   struct StellarValue
//    {
//        Hash txSetHash;   // transaction set to apply to previous ledger
//        uint64 closeTime; // network close time
//
//        // upgrades to apply to the previous ledger (usually empty)
//        // this is a vector of encoded 'LedgerUpgrade' so that nodes can drop
//        // unknown steps during consensus if needed.
//        // see notes below on 'LedgerUpgrade' for more detail
//        // max size is dictated by number of upgrade types (+ room for future)
//        UpgradeType upgrades<6>;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type StellarValue struct {
	TxSetHash Hash            `json:"txSetHash,omitempty"`
	CloseTime Uint64          `json:"closeTime,omitempty"`
	Upgrades  []UpgradeType   `json:"upgrades,omitempty" xdrmaxsize:"6"`
	Ext       StellarValueExt `json:"ext,omitempty"`
}

// IdGenerator is an XDR Struct defines as:
//
//   struct IdGenerator {
//    	LedgerEntryType entryType; // type of the entry, for which ids will be generated
//    	uint64 idPool; // last used entry specific ID, used for generating entry of specified type
//    };
//
type IdGenerator struct {
	EntryType LedgerEntryType `json:"entryType,omitempty"`
	IdPool    Uint64          `json:"idPool,omitempty"`
}

// LedgerHeaderExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LedgerHeaderExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerHeaderExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerHeaderExt
func (u LedgerHeaderExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerHeaderExt creates a new  LedgerHeaderExt.
func NewLedgerHeaderExt(v LedgerVersion, value interface{}) (result LedgerHeaderExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerHeader is an XDR Struct defines as:
//
//   struct LedgerHeader
//    {
//        uint32 ledgerVersion;    // the protocol version of the ledger
//        Hash previousLedgerHash; // hash of the previous ledger header
//        StellarValue scpValue;   // what consensus agreed to
//        Hash txSetResultHash;    // the TransactionResultSet that led to this ledger
//        Hash bucketListHash;     // hash of the ledger state
//
//        uint32 ledgerSeq; // sequence number of this ledger
//
//        IdGenerator idGenerators<>; // generators of ids
//
//        uint32 maxTxSetSize; // maximum size a transaction set can be
//
//        int64 txExpirationPeriod;
//
//        Hash skipList[4]; // hashes of ledgers in the past. allows you to jump back
//                          // in time without walking the chain back ledger by ledger
//                          // each slot contains the oldest ledger that is mod of
//                          // either 50  5000  50000 or 500000 depending on index
//                          // skipList[0] mod(50), skipList[1] mod(5000), etc
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LedgerHeader struct {
	LedgerVersion      Uint32          `json:"ledgerVersion,omitempty"`
	PreviousLedgerHash Hash            `json:"previousLedgerHash,omitempty"`
	ScpValue           StellarValue    `json:"scpValue,omitempty"`
	TxSetResultHash    Hash            `json:"txSetResultHash,omitempty"`
	BucketListHash     Hash            `json:"bucketListHash,omitempty"`
	LedgerSeq          Uint32          `json:"ledgerSeq,omitempty"`
	IdGenerators       []IdGenerator   `json:"idGenerators,omitempty"`
	MaxTxSetSize       Uint32          `json:"maxTxSetSize,omitempty"`
	TxExpirationPeriod Int64           `json:"txExpirationPeriod,omitempty"`
	SkipList           [4]Hash         `json:"skipList,omitempty"`
	Ext                LedgerHeaderExt `json:"ext,omitempty"`
}

// LedgerUpgradeType is an XDR Enum defines as:
//
//   enum LedgerUpgradeType
//    {
//        VERSION = 1,
//        MAX_TX_SET_SIZE = 2,
//        TX_EXPIRATION_PERIOD = 3
//    };
//
type LedgerUpgradeType int32

const (
	LedgerUpgradeTypeVersion            LedgerUpgradeType = 1
	LedgerUpgradeTypeMaxTxSetSize       LedgerUpgradeType = 2
	LedgerUpgradeTypeTxExpirationPeriod LedgerUpgradeType = 3
)

var LedgerUpgradeTypeAll = []LedgerUpgradeType{
	LedgerUpgradeTypeVersion,
	LedgerUpgradeTypeMaxTxSetSize,
	LedgerUpgradeTypeTxExpirationPeriod,
}

var ledgerUpgradeTypeMap = map[int32]string{
	1: "LedgerUpgradeTypeVersion",
	2: "LedgerUpgradeTypeMaxTxSetSize",
	3: "LedgerUpgradeTypeTxExpirationPeriod",
}

var ledgerUpgradeTypeShortMap = map[int32]string{
	1: "version",
	2: "max_tx_set_size",
	3: "tx_expiration_period",
}

var ledgerUpgradeTypeRevMap = map[string]int32{
	"LedgerUpgradeTypeVersion":            1,
	"LedgerUpgradeTypeMaxTxSetSize":       2,
	"LedgerUpgradeTypeTxExpirationPeriod": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerUpgradeType
func (e LedgerUpgradeType) ValidEnum(v int32) bool {
	_, ok := ledgerUpgradeTypeMap[v]
	return ok
}
func (e LedgerUpgradeType) isFlag() bool {
	for i := len(LedgerUpgradeTypeAll) - 1; i >= 0; i-- {
		expected := LedgerUpgradeType(2) << uint64(len(LedgerUpgradeTypeAll)-1) >> uint64(len(LedgerUpgradeTypeAll)-i)
		if expected != LedgerUpgradeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerUpgradeType) String() string {
	name, _ := ledgerUpgradeTypeMap[int32(e)]
	return name
}

func (e LedgerUpgradeType) ShortString() string {
	name, _ := ledgerUpgradeTypeShortMap[int32(e)]
	return name
}

func (e LedgerUpgradeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range LedgerUpgradeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerUpgradeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerUpgradeType(t.Value)
	return nil
}

// LedgerUpgrade is an XDR Union defines as:
//
//   union LedgerUpgrade switch (LedgerUpgradeType type)
//    {
//    case VERSION:
//        uint32 newLedgerVersion; // update ledgerVersion
//    case MAX_TX_SET_SIZE:
//        uint32 newMaxTxSetSize; // update maxTxSetSize
//    case TX_EXPIRATION_PERIOD:
//        int64 newTxExpirationPeriod;
//    };
//
type LedgerUpgrade struct {
	Type                  LedgerUpgradeType `json:"type,omitempty"`
	NewLedgerVersion      *Uint32           `json:"newLedgerVersion,omitempty"`
	NewMaxTxSetSize       *Uint32           `json:"newMaxTxSetSize,omitempty"`
	NewTxExpirationPeriod *Int64            `json:"newTxExpirationPeriod,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerUpgrade) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerUpgrade
func (u LedgerUpgrade) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerUpgradeType(sw) {
	case LedgerUpgradeTypeVersion:
		return "NewLedgerVersion", true
	case LedgerUpgradeTypeMaxTxSetSize:
		return "NewMaxTxSetSize", true
	case LedgerUpgradeTypeTxExpirationPeriod:
		return "NewTxExpirationPeriod", true
	}
	return "-", false
}

// NewLedgerUpgrade creates a new  LedgerUpgrade.
func NewLedgerUpgrade(aType LedgerUpgradeType, value interface{}) (result LedgerUpgrade, err error) {
	result.Type = aType
	switch LedgerUpgradeType(aType) {
	case LedgerUpgradeTypeVersion:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewLedgerVersion = &tv
	case LedgerUpgradeTypeMaxTxSetSize:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewMaxTxSetSize = &tv
	case LedgerUpgradeTypeTxExpirationPeriod:
		tv, ok := value.(Int64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Int64")
			return
		}
		result.NewTxExpirationPeriod = &tv
	}
	return
}

// MustNewLedgerVersion retrieves the NewLedgerVersion value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewLedgerVersion() Uint32 {
	val, ok := u.GetNewLedgerVersion()

	if !ok {
		panic("arm NewLedgerVersion is not set")
	}

	return val
}

// GetNewLedgerVersion retrieves the NewLedgerVersion value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewLedgerVersion() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewLedgerVersion" {
		result = *u.NewLedgerVersion
		ok = true
	}

	return
}

// MustNewMaxTxSetSize retrieves the NewMaxTxSetSize value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewMaxTxSetSize() Uint32 {
	val, ok := u.GetNewMaxTxSetSize()

	if !ok {
		panic("arm NewMaxTxSetSize is not set")
	}

	return val
}

// GetNewMaxTxSetSize retrieves the NewMaxTxSetSize value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewMaxTxSetSize() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewMaxTxSetSize" {
		result = *u.NewMaxTxSetSize
		ok = true
	}

	return
}

// MustNewTxExpirationPeriod retrieves the NewTxExpirationPeriod value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewTxExpirationPeriod() Int64 {
	val, ok := u.GetNewTxExpirationPeriod()

	if !ok {
		panic("arm NewTxExpirationPeriod is not set")
	}

	return val
}

// GetNewTxExpirationPeriod retrieves the NewTxExpirationPeriod value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewTxExpirationPeriod() (result Int64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewTxExpirationPeriod" {
		result = *u.NewTxExpirationPeriod
		ok = true
	}

	return
}

// BucketEntryType is an XDR Enum defines as:
//
//   enum BucketEntryType
//    {
//        LIVEENTRY = 0,
//        DEADENTRY = 1
//    };
//
type BucketEntryType int32

const (
	BucketEntryTypeLiveentry BucketEntryType = 0
	BucketEntryTypeDeadentry BucketEntryType = 1
)

var BucketEntryTypeAll = []BucketEntryType{
	BucketEntryTypeLiveentry,
	BucketEntryTypeDeadentry,
}

var bucketEntryTypeMap = map[int32]string{
	0: "BucketEntryTypeLiveentry",
	1: "BucketEntryTypeDeadentry",
}

var bucketEntryTypeShortMap = map[int32]string{
	0: "liveentry",
	1: "deadentry",
}

var bucketEntryTypeRevMap = map[string]int32{
	"BucketEntryTypeLiveentry": 0,
	"BucketEntryTypeDeadentry": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BucketEntryType
func (e BucketEntryType) ValidEnum(v int32) bool {
	_, ok := bucketEntryTypeMap[v]
	return ok
}
func (e BucketEntryType) isFlag() bool {
	for i := len(BucketEntryTypeAll) - 1; i >= 0; i-- {
		expected := BucketEntryType(2) << uint64(len(BucketEntryTypeAll)-1) >> uint64(len(BucketEntryTypeAll)-i)
		if expected != BucketEntryTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e BucketEntryType) String() string {
	name, _ := bucketEntryTypeMap[int32(e)]
	return name
}

func (e BucketEntryType) ShortString() string {
	name, _ := bucketEntryTypeShortMap[int32(e)]
	return name
}

func (e BucketEntryType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range BucketEntryTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *BucketEntryType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = BucketEntryType(t.Value)
	return nil
}

// BucketEntry is an XDR Union defines as:
//
//   union BucketEntry switch (BucketEntryType type)
//    {
//    case LIVEENTRY:
//        LedgerEntry liveEntry;
//
//    case DEADENTRY:
//        LedgerKey deadEntry;
//    };
//
type BucketEntry struct {
	Type      BucketEntryType `json:"type,omitempty"`
	LiveEntry *LedgerEntry    `json:"liveEntry,omitempty"`
	DeadEntry *LedgerKey      `json:"deadEntry,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BucketEntry) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BucketEntry
func (u BucketEntry) ArmForSwitch(sw int32) (string, bool) {
	switch BucketEntryType(sw) {
	case BucketEntryTypeLiveentry:
		return "LiveEntry", true
	case BucketEntryTypeDeadentry:
		return "DeadEntry", true
	}
	return "-", false
}

// NewBucketEntry creates a new  BucketEntry.
func NewBucketEntry(aType BucketEntryType, value interface{}) (result BucketEntry, err error) {
	result.Type = aType
	switch BucketEntryType(aType) {
	case BucketEntryTypeLiveentry:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.LiveEntry = &tv
	case BucketEntryTypeDeadentry:
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.DeadEntry = &tv
	}
	return
}

// MustLiveEntry retrieves the LiveEntry value from the union,
// panicing if the value is not set.
func (u BucketEntry) MustLiveEntry() LedgerEntry {
	val, ok := u.GetLiveEntry()

	if !ok {
		panic("arm LiveEntry is not set")
	}

	return val
}

// GetLiveEntry retrieves the LiveEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BucketEntry) GetLiveEntry() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LiveEntry" {
		result = *u.LiveEntry
		ok = true
	}

	return
}

// MustDeadEntry retrieves the DeadEntry value from the union,
// panicing if the value is not set.
func (u BucketEntry) MustDeadEntry() LedgerKey {
	val, ok := u.GetDeadEntry()

	if !ok {
		panic("arm DeadEntry is not set")
	}

	return val
}

// GetDeadEntry retrieves the DeadEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BucketEntry) GetDeadEntry() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DeadEntry" {
		result = *u.DeadEntry
		ok = true
	}

	return
}

// TransactionSet is an XDR Struct defines as:
//
//   struct TransactionSet
//    {
//        Hash previousLedgerHash;
//        TransactionEnvelope txs<>;
//    };
//
type TransactionSet struct {
	PreviousLedgerHash Hash                  `json:"previousLedgerHash,omitempty"`
	Txs                []TransactionEnvelope `json:"txs,omitempty"`
}

// TransactionResultPair is an XDR Struct defines as:
//
//   struct TransactionResultPair
//    {
//        Hash transactionHash;
//        TransactionResult result; // result for the transaction
//    };
//
type TransactionResultPair struct {
	TransactionHash Hash              `json:"transactionHash,omitempty"`
	Result          TransactionResult `json:"result,omitempty"`
}

// TransactionResultSet is an XDR Struct defines as:
//
//   struct TransactionResultSet
//    {
//        TransactionResultPair results<>;
//    };
//
type TransactionResultSet struct {
	Results []TransactionResultPair `json:"results,omitempty"`
}

// TransactionHistoryEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionHistoryEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionHistoryEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionHistoryEntryExt
func (u TransactionHistoryEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionHistoryEntryExt creates a new  TransactionHistoryEntryExt.
func NewTransactionHistoryEntryExt(v LedgerVersion, value interface{}) (result TransactionHistoryEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionHistoryEntry is an XDR Struct defines as:
//
//   struct TransactionHistoryEntry
//    {
//        uint32 ledgerSeq;
//        TransactionSet txSet;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TransactionHistoryEntry struct {
	LedgerSeq Uint32                     `json:"ledgerSeq,omitempty"`
	TxSet     TransactionSet             `json:"txSet,omitempty"`
	Ext       TransactionHistoryEntryExt `json:"ext,omitempty"`
}

// TransactionHistoryResultEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionHistoryResultEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionHistoryResultEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionHistoryResultEntryExt
func (u TransactionHistoryResultEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionHistoryResultEntryExt creates a new  TransactionHistoryResultEntryExt.
func NewTransactionHistoryResultEntryExt(v LedgerVersion, value interface{}) (result TransactionHistoryResultEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionHistoryResultEntry is an XDR Struct defines as:
//
//   struct TransactionHistoryResultEntry
//    {
//        uint32 ledgerSeq;
//        TransactionResultSet txResultSet;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TransactionHistoryResultEntry struct {
	LedgerSeq   Uint32                           `json:"ledgerSeq,omitempty"`
	TxResultSet TransactionResultSet             `json:"txResultSet,omitempty"`
	Ext         TransactionHistoryResultEntryExt `json:"ext,omitempty"`
}

// LedgerHeaderHistoryEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LedgerHeaderHistoryEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerHeaderHistoryEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerHeaderHistoryEntryExt
func (u LedgerHeaderHistoryEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerHeaderHistoryEntryExt creates a new  LedgerHeaderHistoryEntryExt.
func NewLedgerHeaderHistoryEntryExt(v LedgerVersion, value interface{}) (result LedgerHeaderHistoryEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerHeaderHistoryEntry is an XDR Struct defines as:
//
//   struct LedgerHeaderHistoryEntry
//    {
//        Hash hash;
//        LedgerHeader header;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LedgerHeaderHistoryEntry struct {
	Hash   Hash                        `json:"hash,omitempty"`
	Header LedgerHeader                `json:"header,omitempty"`
	Ext    LedgerHeaderHistoryEntryExt `json:"ext,omitempty"`
}

// LedgerScpMessages is an XDR Struct defines as:
//
//   struct LedgerSCPMessages
//    {
//        uint32 ledgerSeq;
//        SCPEnvelope messages<>;
//    };
//
type LedgerScpMessages struct {
	LedgerSeq Uint32        `json:"ledgerSeq,omitempty"`
	Messages  []ScpEnvelope `json:"messages,omitempty"`
}

// ScpHistoryEntryV0 is an XDR Struct defines as:
//
//   struct SCPHistoryEntryV0
//    {
//        SCPQuorumSet quorumSets<>; // additional quorum sets used by ledgerMessages
//        LedgerSCPMessages ledgerMessages;
//    };
//
type ScpHistoryEntryV0 struct {
	QuorumSets     []ScpQuorumSet    `json:"quorumSets,omitempty"`
	LedgerMessages LedgerScpMessages `json:"ledgerMessages,omitempty"`
}

// ScpHistoryEntry is an XDR Union defines as:
//
//   union SCPHistoryEntry switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        SCPHistoryEntryV0 v0;
//    };
//
type ScpHistoryEntry struct {
	V  LedgerVersion      `json:"v,omitempty"`
	V0 *ScpHistoryEntryV0 `json:"v0,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ScpHistoryEntry) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ScpHistoryEntry
func (u ScpHistoryEntry) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "V0", true
	}
	return "-", false
}

// NewScpHistoryEntry creates a new  ScpHistoryEntry.
func NewScpHistoryEntry(v LedgerVersion, value interface{}) (result ScpHistoryEntry, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		tv, ok := value.(ScpHistoryEntryV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpHistoryEntryV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u ScpHistoryEntry) MustV0() ScpHistoryEntryV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpHistoryEntry) GetV0() (result ScpHistoryEntryV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// LedgerEntryChangeType is an XDR Enum defines as:
//
//   enum LedgerEntryChangeType
//    {
//        CREATED = 0, // entry was added to the ledger
//        UPDATED = 1, // entry was modified in the ledger
//        REMOVED = 2, // entry was removed from the ledger
//        STATE = 3    // value of the entry
//    };
//
type LedgerEntryChangeType int32

const (
	LedgerEntryChangeTypeCreated LedgerEntryChangeType = 0
	LedgerEntryChangeTypeUpdated LedgerEntryChangeType = 1
	LedgerEntryChangeTypeRemoved LedgerEntryChangeType = 2
	LedgerEntryChangeTypeState   LedgerEntryChangeType = 3
)

var LedgerEntryChangeTypeAll = []LedgerEntryChangeType{
	LedgerEntryChangeTypeCreated,
	LedgerEntryChangeTypeUpdated,
	LedgerEntryChangeTypeRemoved,
	LedgerEntryChangeTypeState,
}

var ledgerEntryChangeTypeMap = map[int32]string{
	0: "LedgerEntryChangeTypeCreated",
	1: "LedgerEntryChangeTypeUpdated",
	2: "LedgerEntryChangeTypeRemoved",
	3: "LedgerEntryChangeTypeState",
}

var ledgerEntryChangeTypeShortMap = map[int32]string{
	0: "created",
	1: "updated",
	2: "removed",
	3: "state",
}

var ledgerEntryChangeTypeRevMap = map[string]int32{
	"LedgerEntryChangeTypeCreated": 0,
	"LedgerEntryChangeTypeUpdated": 1,
	"LedgerEntryChangeTypeRemoved": 2,
	"LedgerEntryChangeTypeState":   3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerEntryChangeType
func (e LedgerEntryChangeType) ValidEnum(v int32) bool {
	_, ok := ledgerEntryChangeTypeMap[v]
	return ok
}
func (e LedgerEntryChangeType) isFlag() bool {
	for i := len(LedgerEntryChangeTypeAll) - 1; i >= 0; i-- {
		expected := LedgerEntryChangeType(2) << uint64(len(LedgerEntryChangeTypeAll)-1) >> uint64(len(LedgerEntryChangeTypeAll)-i)
		if expected != LedgerEntryChangeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerEntryChangeType) String() string {
	name, _ := ledgerEntryChangeTypeMap[int32(e)]
	return name
}

func (e LedgerEntryChangeType) ShortString() string {
	name, _ := ledgerEntryChangeTypeShortMap[int32(e)]
	return name
}

func (e LedgerEntryChangeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range LedgerEntryChangeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerEntryChangeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerEntryChangeType(t.Value)
	return nil
}

// LedgerEntryChange is an XDR Union defines as:
//
//   union LedgerEntryChange switch (LedgerEntryChangeType type)
//    {
//    case CREATED:
//        LedgerEntry created;
//    case UPDATED:
//        LedgerEntry updated;
//    case REMOVED:
//        LedgerKey removed;
//    case STATE:
//        LedgerEntry state;
//    };
//
type LedgerEntryChange struct {
	Type    LedgerEntryChangeType `json:"type,omitempty"`
	Created *LedgerEntry          `json:"created,omitempty"`
	Updated *LedgerEntry          `json:"updated,omitempty"`
	Removed *LedgerKey            `json:"removed,omitempty"`
	State   *LedgerEntry          `json:"state,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryChange) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryChange
func (u LedgerEntryChange) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryChangeType(sw) {
	case LedgerEntryChangeTypeCreated:
		return "Created", true
	case LedgerEntryChangeTypeUpdated:
		return "Updated", true
	case LedgerEntryChangeTypeRemoved:
		return "Removed", true
	case LedgerEntryChangeTypeState:
		return "State", true
	}
	return "-", false
}

// NewLedgerEntryChange creates a new  LedgerEntryChange.
func NewLedgerEntryChange(aType LedgerEntryChangeType, value interface{}) (result LedgerEntryChange, err error) {
	result.Type = aType
	switch LedgerEntryChangeType(aType) {
	case LedgerEntryChangeTypeCreated:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.Created = &tv
	case LedgerEntryChangeTypeUpdated:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.Updated = &tv
	case LedgerEntryChangeTypeRemoved:
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.Removed = &tv
	case LedgerEntryChangeTypeState:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.State = &tv
	}
	return
}

// MustCreated retrieves the Created value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustCreated() LedgerEntry {
	val, ok := u.GetCreated()

	if !ok {
		panic("arm Created is not set")
	}

	return val
}

// GetCreated retrieves the Created value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetCreated() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Created" {
		result = *u.Created
		ok = true
	}

	return
}

// MustUpdated retrieves the Updated value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustUpdated() LedgerEntry {
	val, ok := u.GetUpdated()

	if !ok {
		panic("arm Updated is not set")
	}

	return val
}

// GetUpdated retrieves the Updated value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetUpdated() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Updated" {
		result = *u.Updated
		ok = true
	}

	return
}

// MustRemoved retrieves the Removed value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustRemoved() LedgerKey {
	val, ok := u.GetRemoved()

	if !ok {
		panic("arm Removed is not set")
	}

	return val
}

// GetRemoved retrieves the Removed value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetRemoved() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Removed" {
		result = *u.Removed
		ok = true
	}

	return
}

// MustState retrieves the State value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustState() LedgerEntry {
	val, ok := u.GetState()

	if !ok {
		panic("arm State is not set")
	}

	return val
}

// GetState retrieves the State value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetState() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "State" {
		result = *u.State
		ok = true
	}

	return
}

// LedgerEntryChanges is an XDR Typedef defines as:
//
//   typedef LedgerEntryChange LedgerEntryChanges<>;
//
type LedgerEntryChanges []LedgerEntryChange

// OperationMeta is an XDR Struct defines as:
//
//   struct OperationMeta
//    {
//        LedgerEntryChanges changes;
//    };
//
type OperationMeta struct {
	Changes LedgerEntryChanges `json:"changes,omitempty"`
}

// TransactionMeta is an XDR Union defines as:
//
//   union TransactionMeta switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        OperationMeta operations<>;
//    };
//
type TransactionMeta struct {
	V          LedgerVersion    `json:"v,omitempty"`
	Operations *[]OperationMeta `json:"operations,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionMeta) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionMeta
func (u TransactionMeta) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "Operations", true
	}
	return "-", false
}

// NewTransactionMeta creates a new  TransactionMeta.
func NewTransactionMeta(v LedgerVersion, value interface{}) (result TransactionMeta, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		tv, ok := value.([]OperationMeta)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationMeta")
			return
		}
		result.Operations = &tv
	}
	return
}

// MustOperations retrieves the Operations value from the union,
// panicing if the value is not set.
func (u TransactionMeta) MustOperations() []OperationMeta {
	val, ok := u.GetOperations()

	if !ok {
		panic("arm Operations is not set")
	}

	return val
}

// GetOperations retrieves the Operations value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionMeta) GetOperations() (result []OperationMeta, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "Operations" {
		result = *u.Operations
		ok = true
	}

	return
}

// CancelChangeRoleRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CancelChangeRoleRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CancelChangeRoleRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CancelChangeRoleRequestOpExt
func (u CancelChangeRoleRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCancelChangeRoleRequestOpExt creates a new  CancelChangeRoleRequestOpExt.
func NewCancelChangeRoleRequestOpExt(v LedgerVersion, value interface{}) (result CancelChangeRoleRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CancelChangeRoleRequestOp is an XDR Struct defines as:
//
//   //: CancelChangeRoleRequestOp is used to cancel reviwable request for changing role.
//    //: If successful, request with the corresponding ID will be deleted
//    struct CancelChangeRoleRequestOp
//    {
//        //: ID of the ChangeRoleRequest request to be canceled
//        uint64 requestID;
//
//        //: Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//
//    };
//
type CancelChangeRoleRequestOp struct {
	RequestId Uint64                       `json:"requestID,omitempty"`
	Ext       CancelChangeRoleRequestOpExt `json:"ext,omitempty"`
}

// CancelChangeRoleRequestResultCode is an XDR Enum defines as:
//
//   //: Result codes for CancelChangeRoleRequest operation
//    enum CancelChangeRoleRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        //: Operation is successfully applied
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: ID of a request cannot be 0
//        REQUEST_ID_INVALID = -1, // request id can not be equal zero
//        //: ChangeRole request with provided ID is not found
//        REQUEST_NOT_FOUND = -2 // trying to cancel not existing reviewable request
//    };
//
type CancelChangeRoleRequestResultCode int32

const (
	CancelChangeRoleRequestResultCodeSuccess          CancelChangeRoleRequestResultCode = 0
	CancelChangeRoleRequestResultCodeRequestIdInvalid CancelChangeRoleRequestResultCode = -1
	CancelChangeRoleRequestResultCodeRequestNotFound  CancelChangeRoleRequestResultCode = -2
)

var CancelChangeRoleRequestResultCodeAll = []CancelChangeRoleRequestResultCode{
	CancelChangeRoleRequestResultCodeSuccess,
	CancelChangeRoleRequestResultCodeRequestIdInvalid,
	CancelChangeRoleRequestResultCodeRequestNotFound,
}

var cancelChangeRoleRequestResultCodeMap = map[int32]string{
	0:  "CancelChangeRoleRequestResultCodeSuccess",
	-1: "CancelChangeRoleRequestResultCodeRequestIdInvalid",
	-2: "CancelChangeRoleRequestResultCodeRequestNotFound",
}

var cancelChangeRoleRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "request_id_invalid",
	-2: "request_not_found",
}

var cancelChangeRoleRequestResultCodeRevMap = map[string]int32{
	"CancelChangeRoleRequestResultCodeSuccess":          0,
	"CancelChangeRoleRequestResultCodeRequestIdInvalid": -1,
	"CancelChangeRoleRequestResultCodeRequestNotFound":  -2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CancelChangeRoleRequestResultCode
func (e CancelChangeRoleRequestResultCode) ValidEnum(v int32) bool {
	_, ok := cancelChangeRoleRequestResultCodeMap[v]
	return ok
}
func (e CancelChangeRoleRequestResultCode) isFlag() bool {
	for i := len(CancelChangeRoleRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CancelChangeRoleRequestResultCode(2) << uint64(len(CancelChangeRoleRequestResultCodeAll)-1) >> uint64(len(CancelChangeRoleRequestResultCodeAll)-i)
		if expected != CancelChangeRoleRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CancelChangeRoleRequestResultCode) String() string {
	name, _ := cancelChangeRoleRequestResultCodeMap[int32(e)]
	return name
}

func (e CancelChangeRoleRequestResultCode) ShortString() string {
	name, _ := cancelChangeRoleRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CancelChangeRoleRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CancelChangeRoleRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CancelChangeRoleRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CancelChangeRoleRequestResultCode(t.Value)
	return nil
}

// CancelChangeRoleSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CancelChangeRoleSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CancelChangeRoleSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CancelChangeRoleSuccessExt
func (u CancelChangeRoleSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCancelChangeRoleSuccessExt creates a new  CancelChangeRoleSuccessExt.
func NewCancelChangeRoleSuccessExt(v LedgerVersion, value interface{}) (result CancelChangeRoleSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CancelChangeRoleSuccess is an XDR Struct defines as:
//
//   //: Result of successful `CancelChangeRoleRequestOp` application
//    struct CancelChangeRoleSuccess {
//
//        //: Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CancelChangeRoleSuccess struct {
	Ext CancelChangeRoleSuccessExt `json:"ext,omitempty"`
}

// CancelChangeRoleRequestResult is an XDR Union defines as:
//
//   //: Result of CancelChangeRoleRequest operation application along with the result code
//    union CancelChangeRoleRequestResult switch (CancelChangeRoleRequestResultCode code)
//    {
//        case SUCCESS:
//            CancelChangeRoleSuccess success;
//        default:
//            void;
//    };
//
type CancelChangeRoleRequestResult struct {
	Code    CancelChangeRoleRequestResultCode `json:"code,omitempty"`
	Success *CancelChangeRoleSuccess          `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CancelChangeRoleRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CancelChangeRoleRequestResult
func (u CancelChangeRoleRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CancelChangeRoleRequestResultCode(sw) {
	case CancelChangeRoleRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCancelChangeRoleRequestResult creates a new  CancelChangeRoleRequestResult.
func NewCancelChangeRoleRequestResult(code CancelChangeRoleRequestResultCode, value interface{}) (result CancelChangeRoleRequestResult, err error) {
	result.Code = code
	switch CancelChangeRoleRequestResultCode(code) {
	case CancelChangeRoleRequestResultCodeSuccess:
		tv, ok := value.(CancelChangeRoleSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CancelChangeRoleSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CancelChangeRoleRequestResult) MustSuccess() CancelChangeRoleSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CancelChangeRoleRequestResult) GetSuccess() (result CancelChangeRoleSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateAccountOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateAccountOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountOpExt
func (u CreateAccountOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAccountOpExt creates a new  CreateAccountOpExt.
func NewCreateAccountOpExt(v LedgerVersion, value interface{}) (result CreateAccountOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAccountOp is an XDR Struct defines as:
//
//   //: CreateAccountOp is used to create new account
//    struct CreateAccountOp
//    {
//        //: ID of account to be created
//        AccountID destination;
//        //: ID of an another account that introduced this account into the system.
//        //: If account with such ID does not exist or it's Admin Account. Referrer won't be set.
//        AccountID* referrer;
//        //: ID of the role that will be attached to an account
//        uint64 roleID;
//
//        //: Array of data about 'destination' account signers to be created
//        SignerData signersData<>;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateAccountOp struct {
	Destination AccountId          `json:"destination,omitempty"`
	Referrer    *AccountId         `json:"referrer,omitempty"`
	RoleId      Uint64             `json:"roleID,omitempty"`
	SignersData []SignerData       `json:"signersData,omitempty"`
	Ext         CreateAccountOpExt `json:"ext,omitempty"`
}

// CreateAccountResultCode is an XDR Enum defines as:
//
//   //: Result codes of CreateAccountOp
//    enum CreateAccountResultCode
//    {
//        //: Means that `destination` account has been successfully created with signers specified in `signersData`
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: Source account cannot be the same as the destination account
//        INVALID_DESTINATION = -1,
//        //: Account with such an ID already exists
//        ALREADY_EXISTS = -2, // account already exist
//        //: Sum of weights of signers with different identities must exceed the threshold (for now, 1000)
//        INVALID_WEIGHT = -3,
//        //: There is no role with such an ID
//        NO_SUCH_ROLE = -4,
//        //: Creation of a signer for an account is failed because `signersData` is invalid.
//        //: See `createSignerErrorCode`
//        INVALID_SIGNER_DATA = -5,
//        //: It is not allowed to create accounts without signers
//        NO_SIGNER_DATA = -6 // empty signer data array not allowed
//    };
//
type CreateAccountResultCode int32

const (
	CreateAccountResultCodeSuccess            CreateAccountResultCode = 0
	CreateAccountResultCodeInvalidDestination CreateAccountResultCode = -1
	CreateAccountResultCodeAlreadyExists      CreateAccountResultCode = -2
	CreateAccountResultCodeInvalidWeight      CreateAccountResultCode = -3
	CreateAccountResultCodeNoSuchRole         CreateAccountResultCode = -4
	CreateAccountResultCodeInvalidSignerData  CreateAccountResultCode = -5
	CreateAccountResultCodeNoSignerData       CreateAccountResultCode = -6
)

var CreateAccountResultCodeAll = []CreateAccountResultCode{
	CreateAccountResultCodeSuccess,
	CreateAccountResultCodeInvalidDestination,
	CreateAccountResultCodeAlreadyExists,
	CreateAccountResultCodeInvalidWeight,
	CreateAccountResultCodeNoSuchRole,
	CreateAccountResultCodeInvalidSignerData,
	CreateAccountResultCodeNoSignerData,
}

var createAccountResultCodeMap = map[int32]string{
	0:  "CreateAccountResultCodeSuccess",
	-1: "CreateAccountResultCodeInvalidDestination",
	-2: "CreateAccountResultCodeAlreadyExists",
	-3: "CreateAccountResultCodeInvalidWeight",
	-4: "CreateAccountResultCodeNoSuchRole",
	-5: "CreateAccountResultCodeInvalidSignerData",
	-6: "CreateAccountResultCodeNoSignerData",
}

var createAccountResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_destination",
	-2: "already_exists",
	-3: "invalid_weight",
	-4: "no_such_role",
	-5: "invalid_signer_data",
	-6: "no_signer_data",
}

var createAccountResultCodeRevMap = map[string]int32{
	"CreateAccountResultCodeSuccess":            0,
	"CreateAccountResultCodeInvalidDestination": -1,
	"CreateAccountResultCodeAlreadyExists":      -2,
	"CreateAccountResultCodeInvalidWeight":      -3,
	"CreateAccountResultCodeNoSuchRole":         -4,
	"CreateAccountResultCodeInvalidSignerData":  -5,
	"CreateAccountResultCodeNoSignerData":       -6,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateAccountResultCode
func (e CreateAccountResultCode) ValidEnum(v int32) bool {
	_, ok := createAccountResultCodeMap[v]
	return ok
}
func (e CreateAccountResultCode) isFlag() bool {
	for i := len(CreateAccountResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateAccountResultCode(2) << uint64(len(CreateAccountResultCodeAll)-1) >> uint64(len(CreateAccountResultCodeAll)-i)
		if expected != CreateAccountResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateAccountResultCode) String() string {
	name, _ := createAccountResultCodeMap[int32(e)]
	return name
}

func (e CreateAccountResultCode) ShortString() string {
	name, _ := createAccountResultCodeShortMap[int32(e)]
	return name
}

func (e CreateAccountResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateAccountResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateAccountResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateAccountResultCode(t.Value)
	return nil
}

// CreateAccountSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateAccountSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountSuccessExt
func (u CreateAccountSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAccountSuccessExt creates a new  CreateAccountSuccessExt.
func NewCreateAccountSuccessExt(v LedgerVersion, value interface{}) (result CreateAccountSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAccountSuccess is an XDR Struct defines as:
//
//   //: Result of successful application of `CreateAccount` operation
//    struct CreateAccountSuccess
//    {
//        //: Unique unsigned integer identifier of the new account
//        uint64 sequentialID;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateAccountSuccess struct {
	SequentialId Uint64                  `json:"sequentialID,omitempty"`
	Ext          CreateAccountSuccessExt `json:"ext,omitempty"`
}

// CreateAccountResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union CreateAccountResult switch (CreateAccountResultCode code)
//    {
//    case SUCCESS:
//        CreateAccountSuccess success;
//    case INVALID_SIGNER_DATA:
//        //: `createSignerErrorCode` is used to determine the reason of signer creation failure
//        ManageSignerResultCode createSignerErrorCode;
//    default:
//        void;
//    };
//
type CreateAccountResult struct {
	Code                  CreateAccountResultCode `json:"code,omitempty"`
	Success               *CreateAccountSuccess   `json:"success,omitempty"`
	CreateSignerErrorCode *ManageSignerResultCode `json:"createSignerErrorCode,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountResult
func (u CreateAccountResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateAccountResultCode(sw) {
	case CreateAccountResultCodeSuccess:
		return "Success", true
	case CreateAccountResultCodeInvalidSignerData:
		return "CreateSignerErrorCode", true
	default:
		return "", true
	}
}

// NewCreateAccountResult creates a new  CreateAccountResult.
func NewCreateAccountResult(code CreateAccountResultCode, value interface{}) (result CreateAccountResult, err error) {
	result.Code = code
	switch CreateAccountResultCode(code) {
	case CreateAccountResultCodeSuccess:
		tv, ok := value.(CreateAccountSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountSuccess")
			return
		}
		result.Success = &tv
	case CreateAccountResultCodeInvalidSignerData:
		tv, ok := value.(ManageSignerResultCode)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSignerResultCode")
			return
		}
		result.CreateSignerErrorCode = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateAccountResult) MustSuccess() CreateAccountSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAccountResult) GetSuccess() (result CreateAccountSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustCreateSignerErrorCode retrieves the CreateSignerErrorCode value from the union,
// panicing if the value is not set.
func (u CreateAccountResult) MustCreateSignerErrorCode() ManageSignerResultCode {
	val, ok := u.GetCreateSignerErrorCode()

	if !ok {
		panic("arm CreateSignerErrorCode is not set")
	}

	return val
}

// GetCreateSignerErrorCode retrieves the CreateSignerErrorCode value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAccountResult) GetCreateSignerErrorCode() (result ManageSignerResultCode, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "CreateSignerErrorCode" {
		result = *u.CreateSignerErrorCode
		ok = true
	}

	return
}

// CreateChangeRoleRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateChangeRoleRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateChangeRoleRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateChangeRoleRequestOpExt
func (u CreateChangeRoleRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateChangeRoleRequestOpExt creates a new  CreateChangeRoleRequestOpExt.
func NewCreateChangeRoleRequestOpExt(v LedgerVersion, value interface{}) (result CreateChangeRoleRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateChangeRoleRequestOp is an XDR Struct defines as:
//
//   //: `CreateChangeRoleRequestOp` is used to create reviewable requests
//    //: that, with admin's approval, will change the role of `destinationAccount`
//    //: from current role to `accountRoleToSet`
//    struct CreateChangeRoleRequestOp
//    {
//        //: Set zero to create new request, set non zero to update existing request
//        uint64 requestID;
//
//        //: AccountID of an account whose role will be changed
//        AccountID destinationAccount;
//        //: ID of account role that will be attached to `destinationAccount`
//        uint64 accountRoleToSet;
//        //: Arbitrary stringified json object that can be used to attach data to be reviewed by an admin
//        longstring creatorDetails;
//
//        //: Bit mask that will be used instead of the value from key-value entry by
//        //: `change_role_tasks:<currentRoleID>:<accountRoleToSet>` key
//        uint32* allTasks;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateChangeRoleRequestOp struct {
	RequestId          Uint64                       `json:"requestID,omitempty"`
	DestinationAccount AccountId                    `json:"destinationAccount,omitempty"`
	AccountRoleToSet   Uint64                       `json:"accountRoleToSet,omitempty"`
	CreatorDetails     Longstring                   `json:"creatorDetails,omitempty"`
	AllTasks           *Uint32                      `json:"allTasks,omitempty"`
	Ext                CreateChangeRoleRequestOpExt `json:"ext,omitempty"`
}

// CreateChangeRoleRequestResultCode is an XDR Enum defines as:
//
//   //: Result codes of CreateChangeRoleRequestOp
//    enum CreateChangeRoleRequestResultCode
//    {
//        //: Change role request has either been successfully created
//        //: or auto approved
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: There is no destination account with such accountID
//        ACC_TO_UPDATE_DOES_NOT_EXIST = -1,
//        //: There is another change role request for such destination account
//        REQUEST_ALREADY_EXISTS = -2,
//        //: There is no request with such `requestID`
//        REQUEST_DOES_NOT_EXIST = -4,
//        //: Only `destinationAccount` can update change role request
//        //: `destinationAccount` must be equal source Account
//        NOT_ALLOWED_TO_UPDATE_REQUEST = -6,
//        //: It is not allowed to change `destinationAccount`, `accountRoleToSet`
//        //: or set `allTasks` on update change role request
//        INVALID_CHANGE_ROLE_REQUEST_DATA = -7,
//        //: `creatorDetails` must be in a valid JSON format
//        INVALID_CREATOR_DETAILS = -8,
//        //: There is no key-value entry by `change_role_tasks` key in the system;
//        //: configuration does not allow changing the role from current to `accountRoleToSet`
//        CHANGE_ROLE_TASKS_NOT_FOUND = -9,
//        //: There is no account role with provided id
//        ACCOUNT_ROLE_TO_SET_DOES_NOT_EXIST = -10
//    };
//
type CreateChangeRoleRequestResultCode int32

const (
	CreateChangeRoleRequestResultCodeSuccess                      CreateChangeRoleRequestResultCode = 0
	CreateChangeRoleRequestResultCodeAccToUpdateDoesNotExist      CreateChangeRoleRequestResultCode = -1
	CreateChangeRoleRequestResultCodeRequestAlreadyExists         CreateChangeRoleRequestResultCode = -2
	CreateChangeRoleRequestResultCodeRequestDoesNotExist          CreateChangeRoleRequestResultCode = -4
	CreateChangeRoleRequestResultCodeNotAllowedToUpdateRequest    CreateChangeRoleRequestResultCode = -6
	CreateChangeRoleRequestResultCodeInvalidChangeRoleRequestData CreateChangeRoleRequestResultCode = -7
	CreateChangeRoleRequestResultCodeInvalidCreatorDetails        CreateChangeRoleRequestResultCode = -8
	CreateChangeRoleRequestResultCodeChangeRoleTasksNotFound      CreateChangeRoleRequestResultCode = -9
	CreateChangeRoleRequestResultCodeAccountRoleToSetDoesNotExist CreateChangeRoleRequestResultCode = -10
)

var CreateChangeRoleRequestResultCodeAll = []CreateChangeRoleRequestResultCode{
	CreateChangeRoleRequestResultCodeSuccess,
	CreateChangeRoleRequestResultCodeAccToUpdateDoesNotExist,
	CreateChangeRoleRequestResultCodeRequestAlreadyExists,
	CreateChangeRoleRequestResultCodeRequestDoesNotExist,
	CreateChangeRoleRequestResultCodeNotAllowedToUpdateRequest,
	CreateChangeRoleRequestResultCodeInvalidChangeRoleRequestData,
	CreateChangeRoleRequestResultCodeInvalidCreatorDetails,
	CreateChangeRoleRequestResultCodeChangeRoleTasksNotFound,
	CreateChangeRoleRequestResultCodeAccountRoleToSetDoesNotExist,
}

var createChangeRoleRequestResultCodeMap = map[int32]string{
	0:   "CreateChangeRoleRequestResultCodeSuccess",
	-1:  "CreateChangeRoleRequestResultCodeAccToUpdateDoesNotExist",
	-2:  "CreateChangeRoleRequestResultCodeRequestAlreadyExists",
	-4:  "CreateChangeRoleRequestResultCodeRequestDoesNotExist",
	-6:  "CreateChangeRoleRequestResultCodeNotAllowedToUpdateRequest",
	-7:  "CreateChangeRoleRequestResultCodeInvalidChangeRoleRequestData",
	-8:  "CreateChangeRoleRequestResultCodeInvalidCreatorDetails",
	-9:  "CreateChangeRoleRequestResultCodeChangeRoleTasksNotFound",
	-10: "CreateChangeRoleRequestResultCodeAccountRoleToSetDoesNotExist",
}

var createChangeRoleRequestResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "acc_to_update_does_not_exist",
	-2:  "request_already_exists",
	-4:  "request_does_not_exist",
	-6:  "not_allowed_to_update_request",
	-7:  "invalid_change_role_request_data",
	-8:  "invalid_creator_details",
	-9:  "change_role_tasks_not_found",
	-10: "account_role_to_set_does_not_exist",
}

var createChangeRoleRequestResultCodeRevMap = map[string]int32{
	"CreateChangeRoleRequestResultCodeSuccess":                      0,
	"CreateChangeRoleRequestResultCodeAccToUpdateDoesNotExist":      -1,
	"CreateChangeRoleRequestResultCodeRequestAlreadyExists":         -2,
	"CreateChangeRoleRequestResultCodeRequestDoesNotExist":          -4,
	"CreateChangeRoleRequestResultCodeNotAllowedToUpdateRequest":    -6,
	"CreateChangeRoleRequestResultCodeInvalidChangeRoleRequestData": -7,
	"CreateChangeRoleRequestResultCodeInvalidCreatorDetails":        -8,
	"CreateChangeRoleRequestResultCodeChangeRoleTasksNotFound":      -9,
	"CreateChangeRoleRequestResultCodeAccountRoleToSetDoesNotExist": -10,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateChangeRoleRequestResultCode
func (e CreateChangeRoleRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createChangeRoleRequestResultCodeMap[v]
	return ok
}
func (e CreateChangeRoleRequestResultCode) isFlag() bool {
	for i := len(CreateChangeRoleRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateChangeRoleRequestResultCode(2) << uint64(len(CreateChangeRoleRequestResultCodeAll)-1) >> uint64(len(CreateChangeRoleRequestResultCodeAll)-i)
		if expected != CreateChangeRoleRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateChangeRoleRequestResultCode) String() string {
	name, _ := createChangeRoleRequestResultCodeMap[int32(e)]
	return name
}

func (e CreateChangeRoleRequestResultCode) ShortString() string {
	name, _ := createChangeRoleRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreateChangeRoleRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateChangeRoleRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateChangeRoleRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateChangeRoleRequestResultCode(t.Value)
	return nil
}

// CreateChangeRoleRequestResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type CreateChangeRoleRequestResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateChangeRoleRequestResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateChangeRoleRequestResultSuccessExt
func (u CreateChangeRoleRequestResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateChangeRoleRequestResultSuccessExt creates a new  CreateChangeRoleRequestResultSuccessExt.
func NewCreateChangeRoleRequestResultSuccessExt(v LedgerVersion, value interface{}) (result CreateChangeRoleRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateChangeRoleRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//            //: ID of a created or updated request
//            uint64 requestID;
//            //: True if request was auto approved (pending tasks == 0),
//            //: `destinationAccount` must have new account role
//            bool fulfilled;
//            // Reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    	}
//
type CreateChangeRoleRequestResultSuccess struct {
	RequestId Uint64                                  `json:"requestID,omitempty"`
	Fulfilled bool                                    `json:"fulfilled,omitempty"`
	Ext       CreateChangeRoleRequestResultSuccessExt `json:"ext,omitempty"`
}

// CreateChangeRoleRequestResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union CreateChangeRoleRequestResult switch (CreateChangeRoleRequestResultCode code)
//    {
//    case SUCCESS:
//        struct {
//            //: ID of a created or updated request
//            uint64 requestID;
//            //: True if request was auto approved (pending tasks == 0),
//            //: `destinationAccount` must have new account role
//            bool fulfilled;
//            // Reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//    	} success;
//    default:
//        void;
//    };
//
type CreateChangeRoleRequestResult struct {
	Code    CreateChangeRoleRequestResultCode     `json:"code,omitempty"`
	Success *CreateChangeRoleRequestResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateChangeRoleRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateChangeRoleRequestResult
func (u CreateChangeRoleRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateChangeRoleRequestResultCode(sw) {
	case CreateChangeRoleRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateChangeRoleRequestResult creates a new  CreateChangeRoleRequestResult.
func NewCreateChangeRoleRequestResult(code CreateChangeRoleRequestResultCode, value interface{}) (result CreateChangeRoleRequestResult, err error) {
	result.Code = code
	switch CreateChangeRoleRequestResultCode(code) {
	case CreateChangeRoleRequestResultCodeSuccess:
		tv, ok := value.(CreateChangeRoleRequestResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateChangeRoleRequestResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateChangeRoleRequestResult) MustSuccess() CreateChangeRoleRequestResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateChangeRoleRequestResult) GetSuccess() (result CreateChangeRoleRequestResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateKycRecoveryRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateKycRecoveryRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateKycRecoveryRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateKycRecoveryRequestOpExt
func (u CreateKycRecoveryRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateKycRecoveryRequestOpExt creates a new  CreateKycRecoveryRequestOpExt.
func NewCreateKycRecoveryRequestOpExt(v LedgerVersion, value interface{}) (result CreateKycRecoveryRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateKycRecoveryRequestOp is an XDR Struct defines as:
//
//   //: CreateKYCRecoveryRequestOp to create KYC recovery request and set new signers for account
//    struct CreateKYCRecoveryRequestOp
//    {
//        //: ID of a reviewable request. If set 0, request is created, else - request is updated
//        uint64 requestID;
//        //: Account for which signers will be set
//        AccountID targetAccount;
//        //: New signers to set
//        SignerData signersData<>;
//
//         //: Arbitrary stringified json object that can be used to attach data to be reviewed by an admin
//        longstring creatorDetails; // details set by requester
//
//        //: (optional) Bit mask whose flags must be cleared in order for KYC recovery request to be approved, which will be used by key `create_kyc_recovery_tasks`
//        //: instead of key-value
//        uint32* allTasks;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type CreateKycRecoveryRequestOp struct {
	RequestId      Uint64                        `json:"requestID,omitempty"`
	TargetAccount  AccountId                     `json:"targetAccount,omitempty"`
	SignersData    []SignerData                  `json:"signersData,omitempty"`
	CreatorDetails Longstring                    `json:"creatorDetails,omitempty"`
	AllTasks       *Uint32                       `json:"allTasks,omitempty"`
	Ext            CreateKycRecoveryRequestOpExt `json:"ext,omitempty"`
}

// CreateKycRecoveryRequestResultCode is an XDR Enum defines as:
//
//   //: Result codes of CreateKYCRecoveryRequestOp
//    enum CreateKYCRecoveryRequestResultCode
//    {
//        //: KYC Recovery request was successfully created
//        SUCCESS = 0,
//
//        //: Creator details are not in a valid JSON format
//        INVALID_CREATOR_DETAILS = -1,
//        //: KYC recovery tasks are not set in the system
//        KYC_RECOVERY_TASKS_NOT_FOUND = -2,
//        //: Not allowed to provide empty slice of signers
//        NO_SIGNER_DATA = -3,
//        //: SignerData contains duplicates
//        SIGNER_DUPLICATION = -4,
//        //: Signer has weight > threshold
//        INVALID_WEIGHT = -5,
//        //: Signer has invalid details
//        INVALID_DETAILS = -6,
//        //: Request with provided parameters already exists
//        REQUEST_ALREADY_EXISTS = -7,
//        //: Account with provided account address does not exist
//        TARGET_ACCOUNT_NOT_FOUND = -8,
//        //: System configuration forbids KYC recovery
//        RECOVERY_NOT_ALLOWED = -10,
//        //: Only target account can update request
//        NOT_ALLOWED_TO_UPDATE_REQUEST = -11,
//        //: There is no request with such ID
//        REQUEST_NOT_FOUND = -12,
//        //: It is forbidden to change target account on update
//        INVALID_UPDATE_DATA = -13,
//        //: It is forbidden to set `allTasks` on update
//        NOT_ALLOWED_TO_SET_TASKS_ON_UPDATE = -14
//    };
//
type CreateKycRecoveryRequestResultCode int32

const (
	CreateKycRecoveryRequestResultCodeSuccess                      CreateKycRecoveryRequestResultCode = 0
	CreateKycRecoveryRequestResultCodeInvalidCreatorDetails        CreateKycRecoveryRequestResultCode = -1
	CreateKycRecoveryRequestResultCodeKycRecoveryTasksNotFound     CreateKycRecoveryRequestResultCode = -2
	CreateKycRecoveryRequestResultCodeNoSignerData                 CreateKycRecoveryRequestResultCode = -3
	CreateKycRecoveryRequestResultCodeSignerDuplication            CreateKycRecoveryRequestResultCode = -4
	CreateKycRecoveryRequestResultCodeInvalidWeight                CreateKycRecoveryRequestResultCode = -5
	CreateKycRecoveryRequestResultCodeInvalidDetails               CreateKycRecoveryRequestResultCode = -6
	CreateKycRecoveryRequestResultCodeRequestAlreadyExists         CreateKycRecoveryRequestResultCode = -7
	CreateKycRecoveryRequestResultCodeTargetAccountNotFound        CreateKycRecoveryRequestResultCode = -8
	CreateKycRecoveryRequestResultCodeRecoveryNotAllowed           CreateKycRecoveryRequestResultCode = -10
	CreateKycRecoveryRequestResultCodeNotAllowedToUpdateRequest    CreateKycRecoveryRequestResultCode = -11
	CreateKycRecoveryRequestResultCodeRequestNotFound              CreateKycRecoveryRequestResultCode = -12
	CreateKycRecoveryRequestResultCodeInvalidUpdateData            CreateKycRecoveryRequestResultCode = -13
	CreateKycRecoveryRequestResultCodeNotAllowedToSetTasksOnUpdate CreateKycRecoveryRequestResultCode = -14
)

var CreateKycRecoveryRequestResultCodeAll = []CreateKycRecoveryRequestResultCode{
	CreateKycRecoveryRequestResultCodeSuccess,
	CreateKycRecoveryRequestResultCodeInvalidCreatorDetails,
	CreateKycRecoveryRequestResultCodeKycRecoveryTasksNotFound,
	CreateKycRecoveryRequestResultCodeNoSignerData,
	CreateKycRecoveryRequestResultCodeSignerDuplication,
	CreateKycRecoveryRequestResultCodeInvalidWeight,
	CreateKycRecoveryRequestResultCodeInvalidDetails,
	CreateKycRecoveryRequestResultCodeRequestAlreadyExists,
	CreateKycRecoveryRequestResultCodeTargetAccountNotFound,
	CreateKycRecoveryRequestResultCodeRecoveryNotAllowed,
	CreateKycRecoveryRequestResultCodeNotAllowedToUpdateRequest,
	CreateKycRecoveryRequestResultCodeRequestNotFound,
	CreateKycRecoveryRequestResultCodeInvalidUpdateData,
	CreateKycRecoveryRequestResultCodeNotAllowedToSetTasksOnUpdate,
}

var createKycRecoveryRequestResultCodeMap = map[int32]string{
	0:   "CreateKycRecoveryRequestResultCodeSuccess",
	-1:  "CreateKycRecoveryRequestResultCodeInvalidCreatorDetails",
	-2:  "CreateKycRecoveryRequestResultCodeKycRecoveryTasksNotFound",
	-3:  "CreateKycRecoveryRequestResultCodeNoSignerData",
	-4:  "CreateKycRecoveryRequestResultCodeSignerDuplication",
	-5:  "CreateKycRecoveryRequestResultCodeInvalidWeight",
	-6:  "CreateKycRecoveryRequestResultCodeInvalidDetails",
	-7:  "CreateKycRecoveryRequestResultCodeRequestAlreadyExists",
	-8:  "CreateKycRecoveryRequestResultCodeTargetAccountNotFound",
	-10: "CreateKycRecoveryRequestResultCodeRecoveryNotAllowed",
	-11: "CreateKycRecoveryRequestResultCodeNotAllowedToUpdateRequest",
	-12: "CreateKycRecoveryRequestResultCodeRequestNotFound",
	-13: "CreateKycRecoveryRequestResultCodeInvalidUpdateData",
	-14: "CreateKycRecoveryRequestResultCodeNotAllowedToSetTasksOnUpdate",
}

var createKycRecoveryRequestResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "invalid_creator_details",
	-2:  "kyc_recovery_tasks_not_found",
	-3:  "no_signer_data",
	-4:  "signer_duplication",
	-5:  "invalid_weight",
	-6:  "invalid_details",
	-7:  "request_already_exists",
	-8:  "target_account_not_found",
	-10: "recovery_not_allowed",
	-11: "not_allowed_to_update_request",
	-12: "request_not_found",
	-13: "invalid_update_data",
	-14: "not_allowed_to_set_tasks_on_update",
}

var createKycRecoveryRequestResultCodeRevMap = map[string]int32{
	"CreateKycRecoveryRequestResultCodeSuccess":                      0,
	"CreateKycRecoveryRequestResultCodeInvalidCreatorDetails":        -1,
	"CreateKycRecoveryRequestResultCodeKycRecoveryTasksNotFound":     -2,
	"CreateKycRecoveryRequestResultCodeNoSignerData":                 -3,
	"CreateKycRecoveryRequestResultCodeSignerDuplication":            -4,
	"CreateKycRecoveryRequestResultCodeInvalidWeight":                -5,
	"CreateKycRecoveryRequestResultCodeInvalidDetails":               -6,
	"CreateKycRecoveryRequestResultCodeRequestAlreadyExists":         -7,
	"CreateKycRecoveryRequestResultCodeTargetAccountNotFound":        -8,
	"CreateKycRecoveryRequestResultCodeRecoveryNotAllowed":           -10,
	"CreateKycRecoveryRequestResultCodeNotAllowedToUpdateRequest":    -11,
	"CreateKycRecoveryRequestResultCodeRequestNotFound":              -12,
	"CreateKycRecoveryRequestResultCodeInvalidUpdateData":            -13,
	"CreateKycRecoveryRequestResultCodeNotAllowedToSetTasksOnUpdate": -14,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateKycRecoveryRequestResultCode
func (e CreateKycRecoveryRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createKycRecoveryRequestResultCodeMap[v]
	return ok
}
func (e CreateKycRecoveryRequestResultCode) isFlag() bool {
	for i := len(CreateKycRecoveryRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateKycRecoveryRequestResultCode(2) << uint64(len(CreateKycRecoveryRequestResultCodeAll)-1) >> uint64(len(CreateKycRecoveryRequestResultCodeAll)-i)
		if expected != CreateKycRecoveryRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateKycRecoveryRequestResultCode) String() string {
	name, _ := createKycRecoveryRequestResultCodeMap[int32(e)]
	return name
}

func (e CreateKycRecoveryRequestResultCode) ShortString() string {
	name, _ := createKycRecoveryRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreateKycRecoveryRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateKycRecoveryRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateKycRecoveryRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateKycRecoveryRequestResultCode(t.Value)
	return nil
}

// CreateKycRecoveryRequestResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type CreateKycRecoveryRequestResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateKycRecoveryRequestResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateKycRecoveryRequestResultSuccessExt
func (u CreateKycRecoveryRequestResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateKycRecoveryRequestResultSuccessExt creates a new  CreateKycRecoveryRequestResultSuccessExt.
func NewCreateKycRecoveryRequestResultSuccessExt(v LedgerVersion, value interface{}) (result CreateKycRecoveryRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateKycRecoveryRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//            //: id of the created request
//            uint64 requestID;
//
//            //: Indicates whether or not the KYC Recovery request was auto approved and fulfilled
//            bool fulfilled;
//
//            //: reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type CreateKycRecoveryRequestResultSuccess struct {
	RequestId Uint64                                   `json:"requestID,omitempty"`
	Fulfilled bool                                     `json:"fulfilled,omitempty"`
	Ext       CreateKycRecoveryRequestResultSuccessExt `json:"ext,omitempty"`
}

// CreateKycRecoveryRequestResult is an XDR Union defines as:
//
//   //: Result of operation applying
//    union CreateKYCRecoveryRequestResult switch (CreateKYCRecoveryRequestResultCode code)
//    {
//    case SUCCESS:
//        //: Is used to pass useful params if operation is success
//        struct {
//            //: id of the created request
//            uint64 requestID;
//
//            //: Indicates whether or not the KYC Recovery request was auto approved and fulfilled
//            bool fulfilled;
//
//            //: reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } success;
//    default:
//        void;
//    };
//
type CreateKycRecoveryRequestResult struct {
	Code    CreateKycRecoveryRequestResultCode     `json:"code,omitempty"`
	Success *CreateKycRecoveryRequestResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateKycRecoveryRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateKycRecoveryRequestResult
func (u CreateKycRecoveryRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateKycRecoveryRequestResultCode(sw) {
	case CreateKycRecoveryRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateKycRecoveryRequestResult creates a new  CreateKycRecoveryRequestResult.
func NewCreateKycRecoveryRequestResult(code CreateKycRecoveryRequestResultCode, value interface{}) (result CreateKycRecoveryRequestResult, err error) {
	result.Code = code
	switch CreateKycRecoveryRequestResultCode(code) {
	case CreateKycRecoveryRequestResultCodeSuccess:
		tv, ok := value.(CreateKycRecoveryRequestResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateKycRecoveryRequestResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateKycRecoveryRequestResult) MustSuccess() CreateKycRecoveryRequestResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateKycRecoveryRequestResult) GetSuccess() (result CreateKycRecoveryRequestResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// InitiateKycRecoveryOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type InitiateKycRecoveryOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InitiateKycRecoveryOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InitiateKycRecoveryOpExt
func (u InitiateKycRecoveryOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewInitiateKycRecoveryOpExt creates a new  InitiateKycRecoveryOpExt.
func NewInitiateKycRecoveryOpExt(v LedgerVersion, value interface{}) (result InitiateKycRecoveryOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// InitiateKycRecoveryOp is an XDR Struct defines as:
//
//   //: InitiateKYCRecoveryOp is used to start KYC recovery process
//    struct InitiateKYCRecoveryOp
//    {
//        //: Address of account to be recovered
//        AccountID account;
//        //: New signer to set
//        PublicKey signer;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type InitiateKycRecoveryOp struct {
	Account AccountId                `json:"account,omitempty"`
	Signer  PublicKey                `json:"signer,omitempty"`
	Ext     InitiateKycRecoveryOpExt `json:"ext,omitempty"`
}

// InitiateKycRecoveryResultCode is an XDR Enum defines as:
//
//   //: Result codes of InitiateKYCRecoveryOp
//    enum InitiateKYCRecoveryResultCode
//    {
//        //: Means that KYC recovery was successfully initiated
//        SUCCESS = 0,
//
//        //: System configuration forbids KYC recovery
//        RECOVERY_NOT_ALLOWED = -1,
//        //: Either, there is no entry by key `kyc_recovery_signer_role`, or such role does not exists
//        RECOVERY_SIGNER_ROLE_NOT_FOUND = -2
//    };
//
type InitiateKycRecoveryResultCode int32

const (
	InitiateKycRecoveryResultCodeSuccess                    InitiateKycRecoveryResultCode = 0
	InitiateKycRecoveryResultCodeRecoveryNotAllowed         InitiateKycRecoveryResultCode = -1
	InitiateKycRecoveryResultCodeRecoverySignerRoleNotFound InitiateKycRecoveryResultCode = -2
)

var InitiateKycRecoveryResultCodeAll = []InitiateKycRecoveryResultCode{
	InitiateKycRecoveryResultCodeSuccess,
	InitiateKycRecoveryResultCodeRecoveryNotAllowed,
	InitiateKycRecoveryResultCodeRecoverySignerRoleNotFound,
}

var initiateKycRecoveryResultCodeMap = map[int32]string{
	0:  "InitiateKycRecoveryResultCodeSuccess",
	-1: "InitiateKycRecoveryResultCodeRecoveryNotAllowed",
	-2: "InitiateKycRecoveryResultCodeRecoverySignerRoleNotFound",
}

var initiateKycRecoveryResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "recovery_not_allowed",
	-2: "recovery_signer_role_not_found",
}

var initiateKycRecoveryResultCodeRevMap = map[string]int32{
	"InitiateKycRecoveryResultCodeSuccess":                    0,
	"InitiateKycRecoveryResultCodeRecoveryNotAllowed":         -1,
	"InitiateKycRecoveryResultCodeRecoverySignerRoleNotFound": -2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for InitiateKycRecoveryResultCode
func (e InitiateKycRecoveryResultCode) ValidEnum(v int32) bool {
	_, ok := initiateKycRecoveryResultCodeMap[v]
	return ok
}
func (e InitiateKycRecoveryResultCode) isFlag() bool {
	for i := len(InitiateKycRecoveryResultCodeAll) - 1; i >= 0; i-- {
		expected := InitiateKycRecoveryResultCode(2) << uint64(len(InitiateKycRecoveryResultCodeAll)-1) >> uint64(len(InitiateKycRecoveryResultCodeAll)-i)
		if expected != InitiateKycRecoveryResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e InitiateKycRecoveryResultCode) String() string {
	name, _ := initiateKycRecoveryResultCodeMap[int32(e)]
	return name
}

func (e InitiateKycRecoveryResultCode) ShortString() string {
	name, _ := initiateKycRecoveryResultCodeShortMap[int32(e)]
	return name
}

func (e InitiateKycRecoveryResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range InitiateKycRecoveryResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *InitiateKycRecoveryResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = InitiateKycRecoveryResultCode(t.Value)
	return nil
}

// InitiateKycRecoveryResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//             {
//             case EMPTY_VERSION:
//                 void;
//             }
//
type InitiateKycRecoveryResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InitiateKycRecoveryResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InitiateKycRecoveryResultSuccessExt
func (u InitiateKycRecoveryResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewInitiateKycRecoveryResultSuccessExt creates a new  InitiateKycRecoveryResultSuccessExt.
func NewInitiateKycRecoveryResultSuccessExt(v LedgerVersion, value interface{}) (result InitiateKycRecoveryResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// InitiateKycRecoveryResultSuccess is an XDR NestedStruct defines as:
//
//   struct
//        {
//             //: reserved for future use
//             union switch (LedgerVersion v)
//             {
//             case EMPTY_VERSION:
//                 void;
//             } ext;
//        }
//
type InitiateKycRecoveryResultSuccess struct {
	Ext InitiateKycRecoveryResultSuccessExt `json:"ext,omitempty"`
}

// InitiateKycRecoveryResult is an XDR Union defines as:
//
//   //: Result of operation applying
//    union InitiateKYCRecoveryResult switch (InitiateKYCRecoveryResultCode code)
//    {
//    case SUCCESS:
//        struct
//        {
//             //: reserved for future use
//             union switch (LedgerVersion v)
//             {
//             case EMPTY_VERSION:
//                 void;
//             } ext;
//        } success;
//    default:
//        void;
//    };
//
type InitiateKycRecoveryResult struct {
	Code    InitiateKycRecoveryResultCode     `json:"code,omitempty"`
	Success *InitiateKycRecoveryResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InitiateKycRecoveryResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InitiateKycRecoveryResult
func (u InitiateKycRecoveryResult) ArmForSwitch(sw int32) (string, bool) {
	switch InitiateKycRecoveryResultCode(sw) {
	case InitiateKycRecoveryResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewInitiateKycRecoveryResult creates a new  InitiateKycRecoveryResult.
func NewInitiateKycRecoveryResult(code InitiateKycRecoveryResultCode, value interface{}) (result InitiateKycRecoveryResult, err error) {
	result.Code = code
	switch InitiateKycRecoveryResultCode(code) {
	case InitiateKycRecoveryResultCodeSuccess:
		tv, ok := value.(InitiateKycRecoveryResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be InitiateKycRecoveryResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u InitiateKycRecoveryResult) MustSuccess() InitiateKycRecoveryResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InitiateKycRecoveryResult) GetSuccess() (result InitiateKycRecoveryResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageAccountRoleAction is an XDR Enum defines as:
//
//   //: Actions that can be performed with the account role
//    enum ManageAccountRoleAction
//    {
//        CREATE = 0,
//        UPDATE = 1,
//        REMOVE = 2
//    };
//
type ManageAccountRoleAction int32

const (
	ManageAccountRoleActionCreate ManageAccountRoleAction = 0
	ManageAccountRoleActionUpdate ManageAccountRoleAction = 1
	ManageAccountRoleActionRemove ManageAccountRoleAction = 2
)

var ManageAccountRoleActionAll = []ManageAccountRoleAction{
	ManageAccountRoleActionCreate,
	ManageAccountRoleActionUpdate,
	ManageAccountRoleActionRemove,
}

var manageAccountRoleActionMap = map[int32]string{
	0: "ManageAccountRoleActionCreate",
	1: "ManageAccountRoleActionUpdate",
	2: "ManageAccountRoleActionRemove",
}

var manageAccountRoleActionShortMap = map[int32]string{
	0: "create",
	1: "update",
	2: "remove",
}

var manageAccountRoleActionRevMap = map[string]int32{
	"ManageAccountRoleActionCreate": 0,
	"ManageAccountRoleActionUpdate": 1,
	"ManageAccountRoleActionRemove": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageAccountRoleAction
func (e ManageAccountRoleAction) ValidEnum(v int32) bool {
	_, ok := manageAccountRoleActionMap[v]
	return ok
}
func (e ManageAccountRoleAction) isFlag() bool {
	for i := len(ManageAccountRoleActionAll) - 1; i >= 0; i-- {
		expected := ManageAccountRoleAction(2) << uint64(len(ManageAccountRoleActionAll)-1) >> uint64(len(ManageAccountRoleActionAll)-i)
		if expected != ManageAccountRoleActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageAccountRoleAction) String() string {
	name, _ := manageAccountRoleActionMap[int32(e)]
	return name
}

func (e ManageAccountRoleAction) ShortString() string {
	name, _ := manageAccountRoleActionShortMap[int32(e)]
	return name
}

func (e ManageAccountRoleAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageAccountRoleActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageAccountRoleAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageAccountRoleAction(t.Value)
	return nil
}

// CreateAccountRoleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateAccountRoleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountRoleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountRoleDataExt
func (u CreateAccountRoleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAccountRoleDataExt creates a new  CreateAccountRoleDataExt.
func NewCreateAccountRoleDataExt(v LedgerVersion, value interface{}) (result CreateAccountRoleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAccountRoleData is an XDR Struct defines as:
//
//   //: CreateAccountRoleData is used to pass necessary params to create a new account role
//    struct CreateAccountRoleData
//    {
//        //: Arbitrary stringified json object that will be attached to the role
//        longstring details;
//        //: Array of ids of existing unique rules
//        uint64 ruleIDs<>;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type CreateAccountRoleData struct {
	Details Longstring               `json:"details,omitempty"`
	RuleIDs []Uint64                 `json:"ruleIDs,omitempty"`
	Ext     CreateAccountRoleDataExt `json:"ext,omitempty"`
}

// UpdateAccountRoleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateAccountRoleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateAccountRoleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateAccountRoleDataExt
func (u UpdateAccountRoleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateAccountRoleDataExt creates a new  UpdateAccountRoleDataExt.
func NewUpdateAccountRoleDataExt(v LedgerVersion, value interface{}) (result UpdateAccountRoleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateAccountRoleData is an XDR Struct defines as:
//
//   //: UpdateAccountRoleData is used to pass necessary params to update existing account role
//    struct UpdateAccountRoleData
//    {
//        //: Identifier of existing signer role
//        uint64 roleID;
//        //: Arbitrary stringified json object that will be attached to the role
//        longstring details;
//        //: Array of ids of existing unique rules
//        uint64 ruleIDs<>;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type UpdateAccountRoleData struct {
	RoleId  Uint64                   `json:"roleID,omitempty"`
	Details Longstring               `json:"details,omitempty"`
	RuleIDs []Uint64                 `json:"ruleIDs,omitempty"`
	Ext     UpdateAccountRoleDataExt `json:"ext,omitempty"`
}

// RemoveAccountRoleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type RemoveAccountRoleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveAccountRoleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveAccountRoleDataExt
func (u RemoveAccountRoleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRemoveAccountRoleDataExt creates a new  RemoveAccountRoleDataExt.
func NewRemoveAccountRoleDataExt(v LedgerVersion, value interface{}) (result RemoveAccountRoleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RemoveAccountRoleData is an XDR Struct defines as:
//
//   //: RemoveAccountRoleData is used to pass necessary params to remove an existing account role
//    struct RemoveAccountRoleData
//    {
//        //: Identifier of an existing account role
//        uint64 roleID;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type RemoveAccountRoleData struct {
	RoleId Uint64                   `json:"roleID,omitempty"`
	Ext    RemoveAccountRoleDataExt `json:"ext,omitempty"`
}

// ManageAccountRoleOpData is an XDR NestedUnion defines as:
//
//   union switch (ManageAccountRoleAction action)
//        {
//        case CREATE:
//            CreateAccountRoleData createData;
//        case UPDATE:
//            UpdateAccountRoleData updateData;
//        case REMOVE:
//            RemoveAccountRoleData removeData;
//        }
//
type ManageAccountRoleOpData struct {
	Action     ManageAccountRoleAction `json:"action,omitempty"`
	CreateData *CreateAccountRoleData  `json:"createData,omitempty"`
	UpdateData *UpdateAccountRoleData  `json:"updateData,omitempty"`
	RemoveData *RemoveAccountRoleData  `json:"removeData,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountRoleOpData) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountRoleOpData
func (u ManageAccountRoleOpData) ArmForSwitch(sw int32) (string, bool) {
	switch ManageAccountRoleAction(sw) {
	case ManageAccountRoleActionCreate:
		return "CreateData", true
	case ManageAccountRoleActionUpdate:
		return "UpdateData", true
	case ManageAccountRoleActionRemove:
		return "RemoveData", true
	}
	return "-", false
}

// NewManageAccountRoleOpData creates a new  ManageAccountRoleOpData.
func NewManageAccountRoleOpData(action ManageAccountRoleAction, value interface{}) (result ManageAccountRoleOpData, err error) {
	result.Action = action
	switch ManageAccountRoleAction(action) {
	case ManageAccountRoleActionCreate:
		tv, ok := value.(CreateAccountRoleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountRoleData")
			return
		}
		result.CreateData = &tv
	case ManageAccountRoleActionUpdate:
		tv, ok := value.(UpdateAccountRoleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateAccountRoleData")
			return
		}
		result.UpdateData = &tv
	case ManageAccountRoleActionRemove:
		tv, ok := value.(RemoveAccountRoleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveAccountRoleData")
			return
		}
		result.RemoveData = &tv
	}
	return
}

// MustCreateData retrieves the CreateData value from the union,
// panicing if the value is not set.
func (u ManageAccountRoleOpData) MustCreateData() CreateAccountRoleData {
	val, ok := u.GetCreateData()

	if !ok {
		panic("arm CreateData is not set")
	}

	return val
}

// GetCreateData retrieves the CreateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRoleOpData) GetCreateData() (result CreateAccountRoleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "CreateData" {
		result = *u.CreateData
		ok = true
	}

	return
}

// MustUpdateData retrieves the UpdateData value from the union,
// panicing if the value is not set.
func (u ManageAccountRoleOpData) MustUpdateData() UpdateAccountRoleData {
	val, ok := u.GetUpdateData()

	if !ok {
		panic("arm UpdateData is not set")
	}

	return val
}

// GetUpdateData retrieves the UpdateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRoleOpData) GetUpdateData() (result UpdateAccountRoleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateData" {
		result = *u.UpdateData
		ok = true
	}

	return
}

// MustRemoveData retrieves the RemoveData value from the union,
// panicing if the value is not set.
func (u ManageAccountRoleOpData) MustRemoveData() RemoveAccountRoleData {
	val, ok := u.GetRemoveData()

	if !ok {
		panic("arm RemoveData is not set")
	}

	return val
}

// GetRemoveData retrieves the RemoveData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRoleOpData) GetRemoveData() (result RemoveAccountRoleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "RemoveData" {
		result = *u.RemoveData
		ok = true
	}

	return
}

// ManageAccountRoleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageAccountRoleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountRoleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountRoleOpExt
func (u ManageAccountRoleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAccountRoleOpExt creates a new  ManageAccountRoleOpExt.
func NewManageAccountRoleOpExt(v LedgerVersion, value interface{}) (result ManageAccountRoleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAccountRoleOp is an XDR Struct defines as:
//
//   //: ManageAccountRoleOp is used to create, update or remove account role
//    struct ManageAccountRoleOp
//    {
//        //: data is used to pass one of `ManageAccountRoleAction` with required params
//        union switch (ManageAccountRoleAction action)
//        {
//        case CREATE:
//            CreateAccountRoleData createData;
//        case UPDATE:
//            UpdateAccountRoleData updateData;
//        case REMOVE:
//            RemoveAccountRoleData removeData;
//        } data;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageAccountRoleOp struct {
	Data ManageAccountRoleOpData `json:"data,omitempty"`
	Ext  ManageAccountRoleOpExt  `json:"ext,omitempty"`
}

// ManageAccountRoleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageAccountRoleResultCode
//    enum ManageAccountRoleResultCode
//    {
//        //: This means that the specified action in `data` of ManageAccountRoleOp was successfully performed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: There is no account role with such id
//        NOT_FOUND = -1,
//        //: THe role cannot be removed if it is attached to at least one account
//        ROLE_IS_USED = -2,
//        //: Passed details has an invalid json structure
//        INVALID_DETAILS = -3,
//        //: There is no rule with id passed through `ruleIDs`
//        NO_SUCH_RULE = -4,
//        //: It is not allowed to duplicate ids in `ruleIDs` array
//        RULE_ID_DUPLICATION = -5
//    };
//
type ManageAccountRoleResultCode int32

const (
	ManageAccountRoleResultCodeSuccess           ManageAccountRoleResultCode = 0
	ManageAccountRoleResultCodeNotFound          ManageAccountRoleResultCode = -1
	ManageAccountRoleResultCodeRoleIsUsed        ManageAccountRoleResultCode = -2
	ManageAccountRoleResultCodeInvalidDetails    ManageAccountRoleResultCode = -3
	ManageAccountRoleResultCodeNoSuchRule        ManageAccountRoleResultCode = -4
	ManageAccountRoleResultCodeRuleIdDuplication ManageAccountRoleResultCode = -5
)

var ManageAccountRoleResultCodeAll = []ManageAccountRoleResultCode{
	ManageAccountRoleResultCodeSuccess,
	ManageAccountRoleResultCodeNotFound,
	ManageAccountRoleResultCodeRoleIsUsed,
	ManageAccountRoleResultCodeInvalidDetails,
	ManageAccountRoleResultCodeNoSuchRule,
	ManageAccountRoleResultCodeRuleIdDuplication,
}

var manageAccountRoleResultCodeMap = map[int32]string{
	0:  "ManageAccountRoleResultCodeSuccess",
	-1: "ManageAccountRoleResultCodeNotFound",
	-2: "ManageAccountRoleResultCodeRoleIsUsed",
	-3: "ManageAccountRoleResultCodeInvalidDetails",
	-4: "ManageAccountRoleResultCodeNoSuchRule",
	-5: "ManageAccountRoleResultCodeRuleIdDuplication",
}

var manageAccountRoleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "role_is_used",
	-3: "invalid_details",
	-4: "no_such_rule",
	-5: "rule_id_duplication",
}

var manageAccountRoleResultCodeRevMap = map[string]int32{
	"ManageAccountRoleResultCodeSuccess":           0,
	"ManageAccountRoleResultCodeNotFound":          -1,
	"ManageAccountRoleResultCodeRoleIsUsed":        -2,
	"ManageAccountRoleResultCodeInvalidDetails":    -3,
	"ManageAccountRoleResultCodeNoSuchRule":        -4,
	"ManageAccountRoleResultCodeRuleIdDuplication": -5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageAccountRoleResultCode
func (e ManageAccountRoleResultCode) ValidEnum(v int32) bool {
	_, ok := manageAccountRoleResultCodeMap[v]
	return ok
}
func (e ManageAccountRoleResultCode) isFlag() bool {
	for i := len(ManageAccountRoleResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageAccountRoleResultCode(2) << uint64(len(ManageAccountRoleResultCodeAll)-1) >> uint64(len(ManageAccountRoleResultCodeAll)-i)
		if expected != ManageAccountRoleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageAccountRoleResultCode) String() string {
	name, _ := manageAccountRoleResultCodeMap[int32(e)]
	return name
}

func (e ManageAccountRoleResultCode) ShortString() string {
	name, _ := manageAccountRoleResultCodeShortMap[int32(e)]
	return name
}

func (e ManageAccountRoleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageAccountRoleResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageAccountRoleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageAccountRoleResultCode(t.Value)
	return nil
}

// ManageAccountRoleResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type ManageAccountRoleResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountRoleResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountRoleResultSuccessExt
func (u ManageAccountRoleResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAccountRoleResultSuccessExt creates a new  ManageAccountRoleResultSuccessExt.
func NewManageAccountRoleResultSuccessExt(v LedgerVersion, value interface{}) (result ManageAccountRoleResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAccountRoleResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                //: id of the role that was managed
//                uint64 roleID;
//
//                //: reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//            }
//
type ManageAccountRoleResultSuccess struct {
	RoleId Uint64                            `json:"roleID,omitempty"`
	Ext    ManageAccountRoleResultSuccessExt `json:"ext,omitempty"`
}

// ManageAccountRoleResult is an XDR Union defines as:
//
//   //: Result of the operation performed
//    union ManageAccountRoleResult switch (ManageAccountRoleResultCode code)
//    {
//        case SUCCESS:
//            //: Is used to pass useful params if the operation is successful
//            struct {
//                //: id of the role that was managed
//                uint64 roleID;
//
//                //: reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//            } success;
//        case RULE_ID_DUPLICATION:
//        case NO_SUCH_RULE:
//            //: ID of a rule that was either duplicated or does not exist
//            uint64 ruleID;
//        default:
//            void;
//    };
//
type ManageAccountRoleResult struct {
	Code    ManageAccountRoleResultCode     `json:"code,omitempty"`
	Success *ManageAccountRoleResultSuccess `json:"success,omitempty"`
	RuleId  *Uint64                         `json:"ruleID,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountRoleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountRoleResult
func (u ManageAccountRoleResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageAccountRoleResultCode(sw) {
	case ManageAccountRoleResultCodeSuccess:
		return "Success", true
	case ManageAccountRoleResultCodeRuleIdDuplication:
		return "RuleId", true
	case ManageAccountRoleResultCodeNoSuchRule:
		return "RuleId", true
	default:
		return "", true
	}
}

// NewManageAccountRoleResult creates a new  ManageAccountRoleResult.
func NewManageAccountRoleResult(code ManageAccountRoleResultCode, value interface{}) (result ManageAccountRoleResult, err error) {
	result.Code = code
	switch ManageAccountRoleResultCode(code) {
	case ManageAccountRoleResultCodeSuccess:
		tv, ok := value.(ManageAccountRoleResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAccountRoleResultSuccess")
			return
		}
		result.Success = &tv
	case ManageAccountRoleResultCodeRuleIdDuplication:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RuleId = &tv
	case ManageAccountRoleResultCodeNoSuchRule:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RuleId = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageAccountRoleResult) MustSuccess() ManageAccountRoleResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRoleResult) GetSuccess() (result ManageAccountRoleResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustRuleId retrieves the RuleId value from the union,
// panicing if the value is not set.
func (u ManageAccountRoleResult) MustRuleId() Uint64 {
	val, ok := u.GetRuleId()

	if !ok {
		panic("arm RuleId is not set")
	}

	return val
}

// GetRuleId retrieves the RuleId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRoleResult) GetRuleId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RuleId" {
		result = *u.RuleId
		ok = true
	}

	return
}

// ManageAccountRuleAction is an XDR Enum defines as:
//
//   //: Actions that can be performed with account rule
//    enum ManageAccountRuleAction
//    {
//        CREATE = 0,
//        UPDATE = 1,
//        REMOVE = 2
//    };
//
type ManageAccountRuleAction int32

const (
	ManageAccountRuleActionCreate ManageAccountRuleAction = 0
	ManageAccountRuleActionUpdate ManageAccountRuleAction = 1
	ManageAccountRuleActionRemove ManageAccountRuleAction = 2
)

var ManageAccountRuleActionAll = []ManageAccountRuleAction{
	ManageAccountRuleActionCreate,
	ManageAccountRuleActionUpdate,
	ManageAccountRuleActionRemove,
}

var manageAccountRuleActionMap = map[int32]string{
	0: "ManageAccountRuleActionCreate",
	1: "ManageAccountRuleActionUpdate",
	2: "ManageAccountRuleActionRemove",
}

var manageAccountRuleActionShortMap = map[int32]string{
	0: "create",
	1: "update",
	2: "remove",
}

var manageAccountRuleActionRevMap = map[string]int32{
	"ManageAccountRuleActionCreate": 0,
	"ManageAccountRuleActionUpdate": 1,
	"ManageAccountRuleActionRemove": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageAccountRuleAction
func (e ManageAccountRuleAction) ValidEnum(v int32) bool {
	_, ok := manageAccountRuleActionMap[v]
	return ok
}
func (e ManageAccountRuleAction) isFlag() bool {
	for i := len(ManageAccountRuleActionAll) - 1; i >= 0; i-- {
		expected := ManageAccountRuleAction(2) << uint64(len(ManageAccountRuleActionAll)-1) >> uint64(len(ManageAccountRuleActionAll)-i)
		if expected != ManageAccountRuleActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageAccountRuleAction) String() string {
	name, _ := manageAccountRuleActionMap[int32(e)]
	return name
}

func (e ManageAccountRuleAction) ShortString() string {
	name, _ := manageAccountRuleActionShortMap[int32(e)]
	return name
}

func (e ManageAccountRuleAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageAccountRuleActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageAccountRuleAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageAccountRuleAction(t.Value)
	return nil
}

// CreateAccountRuleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateAccountRuleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountRuleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountRuleDataExt
func (u CreateAccountRuleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAccountRuleDataExt creates a new  CreateAccountRuleDataExt.
func NewCreateAccountRuleDataExt(v LedgerVersion, value interface{}) (result CreateAccountRuleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAccountRuleData is an XDR Struct defines as:
//
//   //: CreateAccountRuleData is used to pass necessary params to create a new account rule
//    struct CreateAccountRuleData
//    {
//        //: Resource is used to specify an entity (for some - with properties) that can be managed through operations
//        AccountRuleResource resource;
//        //: Value from enum that can be applied to `resource`
//        AccountRuleAction action;
//        //: True if such `action` on such `resource` is prohibited, otherwise allows
//        bool forbids;
//        //: Arbitrary stringified json object that will be attached to rule
//        longstring details;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type CreateAccountRuleData struct {
	Resource AccountRuleResource      `json:"resource,omitempty"`
	Action   AccountRuleAction        `json:"action,omitempty"`
	Forbids  bool                     `json:"forbids,omitempty"`
	Details  Longstring               `json:"details,omitempty"`
	Ext      CreateAccountRuleDataExt `json:"ext,omitempty"`
}

// UpdateAccountRuleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateAccountRuleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateAccountRuleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateAccountRuleDataExt
func (u UpdateAccountRuleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateAccountRuleDataExt creates a new  UpdateAccountRuleDataExt.
func NewUpdateAccountRuleDataExt(v LedgerVersion, value interface{}) (result UpdateAccountRuleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateAccountRuleData is an XDR Struct defines as:
//
//   //: UpdateAccountRuleData is used to pass necessary params to update existing account rule
//    struct UpdateAccountRuleData
//    {
//        //: Identifier of existing signer rule
//        uint64 ruleID;
//        //: Resource is used to specify entity (for some - with properties) that can be managed through operations
//        AccountRuleResource resource;
//        //: Value from enum that can be applied to `resource`
//        AccountRuleAction action;
//        //: True if such `action` on such `resource` is prohibited, otherwise allows
//        bool forbids;
//        //: Arbitrary stringified json object that will be attached to rule
//        longstring details;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type UpdateAccountRuleData struct {
	RuleId   Uint64                   `json:"ruleID,omitempty"`
	Resource AccountRuleResource      `json:"resource,omitempty"`
	Action   AccountRuleAction        `json:"action,omitempty"`
	Forbids  bool                     `json:"forbids,omitempty"`
	Details  Longstring               `json:"details,omitempty"`
	Ext      UpdateAccountRuleDataExt `json:"ext,omitempty"`
}

// RemoveAccountRuleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type RemoveAccountRuleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveAccountRuleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveAccountRuleDataExt
func (u RemoveAccountRuleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRemoveAccountRuleDataExt creates a new  RemoveAccountRuleDataExt.
func NewRemoveAccountRuleDataExt(v LedgerVersion, value interface{}) (result RemoveAccountRuleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RemoveAccountRuleData is an XDR Struct defines as:
//
//   //: RemoveAccountRuleData is used to pass necessary params to remove existing account rule
//    struct RemoveAccountRuleData
//    {
//        //: Identifier of existing account rule
//        uint64 ruleID;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type RemoveAccountRuleData struct {
	RuleId Uint64                   `json:"ruleID,omitempty"`
	Ext    RemoveAccountRuleDataExt `json:"ext,omitempty"`
}

// ManageAccountRuleOpData is an XDR NestedUnion defines as:
//
//   union switch (ManageAccountRuleAction action)
//        {
//        case CREATE:
//            CreateAccountRuleData createData;
//        case UPDATE:
//            UpdateAccountRuleData updateData;
//        case REMOVE:
//            RemoveAccountRuleData removeData;
//        }
//
type ManageAccountRuleOpData struct {
	Action     ManageAccountRuleAction `json:"action,omitempty"`
	CreateData *CreateAccountRuleData  `json:"createData,omitempty"`
	UpdateData *UpdateAccountRuleData  `json:"updateData,omitempty"`
	RemoveData *RemoveAccountRuleData  `json:"removeData,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountRuleOpData) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountRuleOpData
func (u ManageAccountRuleOpData) ArmForSwitch(sw int32) (string, bool) {
	switch ManageAccountRuleAction(sw) {
	case ManageAccountRuleActionCreate:
		return "CreateData", true
	case ManageAccountRuleActionUpdate:
		return "UpdateData", true
	case ManageAccountRuleActionRemove:
		return "RemoveData", true
	}
	return "-", false
}

// NewManageAccountRuleOpData creates a new  ManageAccountRuleOpData.
func NewManageAccountRuleOpData(action ManageAccountRuleAction, value interface{}) (result ManageAccountRuleOpData, err error) {
	result.Action = action
	switch ManageAccountRuleAction(action) {
	case ManageAccountRuleActionCreate:
		tv, ok := value.(CreateAccountRuleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountRuleData")
			return
		}
		result.CreateData = &tv
	case ManageAccountRuleActionUpdate:
		tv, ok := value.(UpdateAccountRuleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateAccountRuleData")
			return
		}
		result.UpdateData = &tv
	case ManageAccountRuleActionRemove:
		tv, ok := value.(RemoveAccountRuleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveAccountRuleData")
			return
		}
		result.RemoveData = &tv
	}
	return
}

// MustCreateData retrieves the CreateData value from the union,
// panicing if the value is not set.
func (u ManageAccountRuleOpData) MustCreateData() CreateAccountRuleData {
	val, ok := u.GetCreateData()

	if !ok {
		panic("arm CreateData is not set")
	}

	return val
}

// GetCreateData retrieves the CreateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRuleOpData) GetCreateData() (result CreateAccountRuleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "CreateData" {
		result = *u.CreateData
		ok = true
	}

	return
}

// MustUpdateData retrieves the UpdateData value from the union,
// panicing if the value is not set.
func (u ManageAccountRuleOpData) MustUpdateData() UpdateAccountRuleData {
	val, ok := u.GetUpdateData()

	if !ok {
		panic("arm UpdateData is not set")
	}

	return val
}

// GetUpdateData retrieves the UpdateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRuleOpData) GetUpdateData() (result UpdateAccountRuleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateData" {
		result = *u.UpdateData
		ok = true
	}

	return
}

// MustRemoveData retrieves the RemoveData value from the union,
// panicing if the value is not set.
func (u ManageAccountRuleOpData) MustRemoveData() RemoveAccountRuleData {
	val, ok := u.GetRemoveData()

	if !ok {
		panic("arm RemoveData is not set")
	}

	return val
}

// GetRemoveData retrieves the RemoveData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRuleOpData) GetRemoveData() (result RemoveAccountRuleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "RemoveData" {
		result = *u.RemoveData
		ok = true
	}

	return
}

// ManageAccountRuleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageAccountRuleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountRuleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountRuleOpExt
func (u ManageAccountRuleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAccountRuleOpExt creates a new  ManageAccountRuleOpExt.
func NewManageAccountRuleOpExt(v LedgerVersion, value interface{}) (result ManageAccountRuleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAccountRuleOp is an XDR Struct defines as:
//
//   //: ManageAccountRuleOp is used to create, update or remove account rule
//    struct ManageAccountRuleOp
//    {
//        //: data is used to pass one of `ManageAccountRuleAction` with required params
//        union switch (ManageAccountRuleAction action)
//        {
//        case CREATE:
//            CreateAccountRuleData createData;
//        case UPDATE:
//            UpdateAccountRuleData updateData;
//        case REMOVE:
//            RemoveAccountRuleData removeData;
//        } data;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageAccountRuleOp struct {
	Data ManageAccountRuleOpData `json:"data,omitempty"`
	Ext  ManageAccountRuleOpExt  `json:"ext,omitempty"`
}

// ManageAccountRuleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageAccountRuleResultCode
//    enum ManageAccountRuleResultCode
//    {
//        //: Means that specified action in `data` of ManageAccountRuleOp was successfully performed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: There is no account rule with such id
//        NOT_FOUND = -1,
//        //: It is not allowed to remove the rule if it is used at least in one role
//        RULE_IS_USED = -2,
//        //: Passed details has invalid json structure
//        INVALID_DETAILS = -3
//    };
//
type ManageAccountRuleResultCode int32

const (
	ManageAccountRuleResultCodeSuccess        ManageAccountRuleResultCode = 0
	ManageAccountRuleResultCodeNotFound       ManageAccountRuleResultCode = -1
	ManageAccountRuleResultCodeRuleIsUsed     ManageAccountRuleResultCode = -2
	ManageAccountRuleResultCodeInvalidDetails ManageAccountRuleResultCode = -3
)

var ManageAccountRuleResultCodeAll = []ManageAccountRuleResultCode{
	ManageAccountRuleResultCodeSuccess,
	ManageAccountRuleResultCodeNotFound,
	ManageAccountRuleResultCodeRuleIsUsed,
	ManageAccountRuleResultCodeInvalidDetails,
}

var manageAccountRuleResultCodeMap = map[int32]string{
	0:  "ManageAccountRuleResultCodeSuccess",
	-1: "ManageAccountRuleResultCodeNotFound",
	-2: "ManageAccountRuleResultCodeRuleIsUsed",
	-3: "ManageAccountRuleResultCodeInvalidDetails",
}

var manageAccountRuleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "rule_is_used",
	-3: "invalid_details",
}

var manageAccountRuleResultCodeRevMap = map[string]int32{
	"ManageAccountRuleResultCodeSuccess":        0,
	"ManageAccountRuleResultCodeNotFound":       -1,
	"ManageAccountRuleResultCodeRuleIsUsed":     -2,
	"ManageAccountRuleResultCodeInvalidDetails": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageAccountRuleResultCode
func (e ManageAccountRuleResultCode) ValidEnum(v int32) bool {
	_, ok := manageAccountRuleResultCodeMap[v]
	return ok
}
func (e ManageAccountRuleResultCode) isFlag() bool {
	for i := len(ManageAccountRuleResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageAccountRuleResultCode(2) << uint64(len(ManageAccountRuleResultCodeAll)-1) >> uint64(len(ManageAccountRuleResultCodeAll)-i)
		if expected != ManageAccountRuleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageAccountRuleResultCode) String() string {
	name, _ := manageAccountRuleResultCodeMap[int32(e)]
	return name
}

func (e ManageAccountRuleResultCode) ShortString() string {
	name, _ := manageAccountRuleResultCodeShortMap[int32(e)]
	return name
}

func (e ManageAccountRuleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageAccountRuleResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageAccountRuleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageAccountRuleResultCode(t.Value)
	return nil
}

// ManageAccountRuleResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type ManageAccountRuleResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountRuleResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountRuleResultSuccessExt
func (u ManageAccountRuleResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAccountRuleResultSuccessExt creates a new  ManageAccountRuleResultSuccessExt.
func NewManageAccountRuleResultSuccessExt(v LedgerVersion, value interface{}) (result ManageAccountRuleResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAccountRuleResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                //: id of the rule that was managed
//                uint64 ruleID;
//
//                //: reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//            }
//
type ManageAccountRuleResultSuccess struct {
	RuleId Uint64                            `json:"ruleID,omitempty"`
	Ext    ManageAccountRuleResultSuccessExt `json:"ext,omitempty"`
}

// ManageAccountRuleResult is an XDR Union defines as:
//
//   //: Result of operation applying
//    union ManageAccountRuleResult switch (ManageAccountRuleResultCode code)
//    {
//        case SUCCESS:
//            //: Is used to pass useful params if operation is success
//            struct {
//                //: id of the rule that was managed
//                uint64 ruleID;
//
//                //: reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//            } success;
//        case RULE_IS_USED:
//            //: ids of roles that use the rule that cannot be removed
//            uint64 roleIDs<>;
//        default:
//            void;
//    };
//
type ManageAccountRuleResult struct {
	Code    ManageAccountRuleResultCode     `json:"code,omitempty"`
	Success *ManageAccountRuleResultSuccess `json:"success,omitempty"`
	RoleIDs *[]Uint64                       `json:"roleIDs,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountRuleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountRuleResult
func (u ManageAccountRuleResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageAccountRuleResultCode(sw) {
	case ManageAccountRuleResultCodeSuccess:
		return "Success", true
	case ManageAccountRuleResultCodeRuleIsUsed:
		return "RoleIDs", true
	default:
		return "", true
	}
}

// NewManageAccountRuleResult creates a new  ManageAccountRuleResult.
func NewManageAccountRuleResult(code ManageAccountRuleResultCode, value interface{}) (result ManageAccountRuleResult, err error) {
	result.Code = code
	switch ManageAccountRuleResultCode(code) {
	case ManageAccountRuleResultCodeSuccess:
		tv, ok := value.(ManageAccountRuleResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAccountRuleResultSuccess")
			return
		}
		result.Success = &tv
	case ManageAccountRuleResultCodeRuleIsUsed:
		tv, ok := value.([]Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be []Uint64")
			return
		}
		result.RoleIDs = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageAccountRuleResult) MustSuccess() ManageAccountRuleResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRuleResult) GetSuccess() (result ManageAccountRuleResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustRoleIDs retrieves the RoleIDs value from the union,
// panicing if the value is not set.
func (u ManageAccountRuleResult) MustRoleIDs() []Uint64 {
	val, ok := u.GetRoleIDs()

	if !ok {
		panic("arm RoleIDs is not set")
	}

	return val
}

// GetRoleIDs retrieves the RoleIDs value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountRuleResult) GetRoleIDs() (result []Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RoleIDs" {
		result = *u.RoleIDs
		ok = true
	}

	return
}

// ManageKvAction is an XDR Enum defines as:
//
//   //: Actions that can be performed on `KeyValueEntry`
//        enum ManageKVAction
//        {
//            PUT = 1,
//            REMOVE = 2
//        };
//
type ManageKvAction int32

const (
	ManageKvActionPut    ManageKvAction = 1
	ManageKvActionRemove ManageKvAction = 2
)

var ManageKvActionAll = []ManageKvAction{
	ManageKvActionPut,
	ManageKvActionRemove,
}

var manageKvActionMap = map[int32]string{
	1: "ManageKvActionPut",
	2: "ManageKvActionRemove",
}

var manageKvActionShortMap = map[int32]string{
	1: "put",
	2: "remove",
}

var manageKvActionRevMap = map[string]int32{
	"ManageKvActionPut":    1,
	"ManageKvActionRemove": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageKvAction
func (e ManageKvAction) ValidEnum(v int32) bool {
	_, ok := manageKvActionMap[v]
	return ok
}
func (e ManageKvAction) isFlag() bool {
	for i := len(ManageKvActionAll) - 1; i >= 0; i-- {
		expected := ManageKvAction(2) << uint64(len(ManageKvActionAll)-1) >> uint64(len(ManageKvActionAll)-i)
		if expected != ManageKvActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageKvAction) String() string {
	name, _ := manageKvActionMap[int32(e)]
	return name
}

func (e ManageKvAction) ShortString() string {
	name, _ := manageKvActionShortMap[int32(e)]
	return name
}

func (e ManageKvAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageKvActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageKvAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageKvAction(t.Value)
	return nil
}

// ManageKeyValueOpAction is an XDR NestedUnion defines as:
//
//   union switch(ManageKVAction action)
//            {
//                case PUT:
//                     KeyValueEntryValue value;
//                case REMOVE:
//                    void;
//            }
//
type ManageKeyValueOpAction struct {
	Action ManageKvAction      `json:"action,omitempty"`
	Value  *KeyValueEntryValue `json:"value,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageKeyValueOpAction) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageKeyValueOpAction
func (u ManageKeyValueOpAction) ArmForSwitch(sw int32) (string, bool) {
	switch ManageKvAction(sw) {
	case ManageKvActionPut:
		return "Value", true
	case ManageKvActionRemove:
		return "", true
	}
	return "-", false
}

// NewManageKeyValueOpAction creates a new  ManageKeyValueOpAction.
func NewManageKeyValueOpAction(action ManageKvAction, value interface{}) (result ManageKeyValueOpAction, err error) {
	result.Action = action
	switch ManageKvAction(action) {
	case ManageKvActionPut:
		tv, ok := value.(KeyValueEntryValue)
		if !ok {
			err = fmt.Errorf("invalid value, must be KeyValueEntryValue")
			return
		}
		result.Value = &tv
	case ManageKvActionRemove:
		// void
	}
	return
}

// MustValue retrieves the Value value from the union,
// panicing if the value is not set.
func (u ManageKeyValueOpAction) MustValue() KeyValueEntryValue {
	val, ok := u.GetValue()

	if !ok {
		panic("arm Value is not set")
	}

	return val
}

// GetValue retrieves the Value value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageKeyValueOpAction) GetValue() (result KeyValueEntryValue, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "Value" {
		result = *u.Value
		ok = true
	}

	return
}

// ManageKeyValueOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type ManageKeyValueOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageKeyValueOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageKeyValueOpExt
func (u ManageKeyValueOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageKeyValueOpExt creates a new  ManageKeyValueOpExt.
func NewManageKeyValueOpExt(v LedgerVersion, value interface{}) (result ManageKeyValueOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageKeyValueOp is an XDR Struct defines as:
//
//   //: `ManageKeyValueOp` is used to create the manage key-value operation which, if applied successfully, will update the key-value entry present in the system
//        struct ManageKeyValueOp
//        {
//            //: `key` is the key for KeyValueEntry
//            longstring key;
//            //: `action` defines an action applied to the KeyValue based on given ManageKVAction
//            //: * Action `PUT` stores new value for a particular key
//            //: * Action `REMOVE` removes the value by a particular key
//            union switch(ManageKVAction action)
//            {
//                case PUT:
//                     KeyValueEntryValue value;
//                case REMOVE:
//                    void;
//            }
//            action;
//
//            //: reserved for future use
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        };
//
type ManageKeyValueOp struct {
	Key    Longstring             `json:"key,omitempty"`
	Action ManageKeyValueOpAction `json:"action,omitempty"`
	Ext    ManageKeyValueOpExt    `json:"ext,omitempty"`
}

// ManageKeyValueSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type ManageKeyValueSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageKeyValueSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageKeyValueSuccessExt
func (u ManageKeyValueSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageKeyValueSuccessExt creates a new  ManageKeyValueSuccessExt.
func NewManageKeyValueSuccessExt(v LedgerVersion, value interface{}) (result ManageKeyValueSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageKeyValueSuccess is an XDR Struct defines as:
//
//   //: `ManageKeyValueSuccess` represents details returned after the successful application of `ManageKeyValueOp`
//        struct ManageKeyValueSuccess
//        {
//            //: reserved for future use
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        };
//
type ManageKeyValueSuccess struct {
	Ext ManageKeyValueSuccessExt `json:"ext,omitempty"`
}

// ManageKeyValueResultCode is an XDR Enum defines as:
//
//   //: Result codes for `ManageKeyValueOp`
//        enum ManageKeyValueResultCode
//        {
//            //: `ManageKeyValueOp` is applied successfully
//            SUCCESS = 0,
//            //: There is no key value with such key
//            NOT_FOUND = -1,
//            //: Value type of the key-value entry is forbidden for the provided key
//            INVALID_TYPE = -2,
//            //: zero value is forbidden for the provided key
//            ZERO_VALUE_NOT_ALLOWED = -3
//        };
//
type ManageKeyValueResultCode int32

const (
	ManageKeyValueResultCodeSuccess             ManageKeyValueResultCode = 0
	ManageKeyValueResultCodeNotFound            ManageKeyValueResultCode = -1
	ManageKeyValueResultCodeInvalidType         ManageKeyValueResultCode = -2
	ManageKeyValueResultCodeZeroValueNotAllowed ManageKeyValueResultCode = -3
)

var ManageKeyValueResultCodeAll = []ManageKeyValueResultCode{
	ManageKeyValueResultCodeSuccess,
	ManageKeyValueResultCodeNotFound,
	ManageKeyValueResultCodeInvalidType,
	ManageKeyValueResultCodeZeroValueNotAllowed,
}

var manageKeyValueResultCodeMap = map[int32]string{
	0:  "ManageKeyValueResultCodeSuccess",
	-1: "ManageKeyValueResultCodeNotFound",
	-2: "ManageKeyValueResultCodeInvalidType",
	-3: "ManageKeyValueResultCodeZeroValueNotAllowed",
}

var manageKeyValueResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "invalid_type",
	-3: "zero_value_not_allowed",
}

var manageKeyValueResultCodeRevMap = map[string]int32{
	"ManageKeyValueResultCodeSuccess":             0,
	"ManageKeyValueResultCodeNotFound":            -1,
	"ManageKeyValueResultCodeInvalidType":         -2,
	"ManageKeyValueResultCodeZeroValueNotAllowed": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageKeyValueResultCode
func (e ManageKeyValueResultCode) ValidEnum(v int32) bool {
	_, ok := manageKeyValueResultCodeMap[v]
	return ok
}
func (e ManageKeyValueResultCode) isFlag() bool {
	for i := len(ManageKeyValueResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageKeyValueResultCode(2) << uint64(len(ManageKeyValueResultCodeAll)-1) >> uint64(len(ManageKeyValueResultCodeAll)-i)
		if expected != ManageKeyValueResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageKeyValueResultCode) String() string {
	name, _ := manageKeyValueResultCodeMap[int32(e)]
	return name
}

func (e ManageKeyValueResultCode) ShortString() string {
	name, _ := manageKeyValueResultCodeShortMap[int32(e)]
	return name
}

func (e ManageKeyValueResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageKeyValueResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageKeyValueResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageKeyValueResultCode(t.Value)
	return nil
}

// ManageKeyValueResult is an XDR Union defines as:
//
//   //: `ManageKeyValueResult` represents the result of ManageKeyValueOp
//        union ManageKeyValueResult switch (ManageKeyValueResultCode code)
//        {
//            case SUCCESS:
//                ManageKeyValueSuccess success;
//            default:
//                void;
//        };
//
type ManageKeyValueResult struct {
	Code    ManageKeyValueResultCode `json:"code,omitempty"`
	Success *ManageKeyValueSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageKeyValueResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageKeyValueResult
func (u ManageKeyValueResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageKeyValueResultCode(sw) {
	case ManageKeyValueResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageKeyValueResult creates a new  ManageKeyValueResult.
func NewManageKeyValueResult(code ManageKeyValueResultCode, value interface{}) (result ManageKeyValueResult, err error) {
	result.Code = code
	switch ManageKeyValueResultCode(code) {
	case ManageKeyValueResultCodeSuccess:
		tv, ok := value.(ManageKeyValueSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageKeyValueSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageKeyValueResult) MustSuccess() ManageKeyValueSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageKeyValueResult) GetSuccess() (result ManageKeyValueSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageSignerRoleAction is an XDR Enum defines as:
//
//   //: Actions that can be performed on a signer role
//    enum ManageSignerRoleAction
//    {
//        CREATE = 0,
//        UPDATE = 1,
//        REMOVE = 2
//    };
//
type ManageSignerRoleAction int32

const (
	ManageSignerRoleActionCreate ManageSignerRoleAction = 0
	ManageSignerRoleActionUpdate ManageSignerRoleAction = 1
	ManageSignerRoleActionRemove ManageSignerRoleAction = 2
)

var ManageSignerRoleActionAll = []ManageSignerRoleAction{
	ManageSignerRoleActionCreate,
	ManageSignerRoleActionUpdate,
	ManageSignerRoleActionRemove,
}

var manageSignerRoleActionMap = map[int32]string{
	0: "ManageSignerRoleActionCreate",
	1: "ManageSignerRoleActionUpdate",
	2: "ManageSignerRoleActionRemove",
}

var manageSignerRoleActionShortMap = map[int32]string{
	0: "create",
	1: "update",
	2: "remove",
}

var manageSignerRoleActionRevMap = map[string]int32{
	"ManageSignerRoleActionCreate": 0,
	"ManageSignerRoleActionUpdate": 1,
	"ManageSignerRoleActionRemove": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageSignerRoleAction
func (e ManageSignerRoleAction) ValidEnum(v int32) bool {
	_, ok := manageSignerRoleActionMap[v]
	return ok
}
func (e ManageSignerRoleAction) isFlag() bool {
	for i := len(ManageSignerRoleActionAll) - 1; i >= 0; i-- {
		expected := ManageSignerRoleAction(2) << uint64(len(ManageSignerRoleActionAll)-1) >> uint64(len(ManageSignerRoleActionAll)-i)
		if expected != ManageSignerRoleActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageSignerRoleAction) String() string {
	name, _ := manageSignerRoleActionMap[int32(e)]
	return name
}

func (e ManageSignerRoleAction) ShortString() string {
	name, _ := manageSignerRoleActionShortMap[int32(e)]
	return name
}

func (e ManageSignerRoleAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageSignerRoleActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageSignerRoleAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageSignerRoleAction(t.Value)
	return nil
}

// CreateSignerRoleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateSignerRoleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateSignerRoleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateSignerRoleDataExt
func (u CreateSignerRoleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateSignerRoleDataExt creates a new  CreateSignerRoleDataExt.
func NewCreateSignerRoleDataExt(v LedgerVersion, value interface{}) (result CreateSignerRoleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateSignerRoleData is an XDR Struct defines as:
//
//   //: CreateSignerRoleData is used to pass necessary params to create a new signer role
//    struct CreateSignerRoleData
//    {
//        //: Array of ids of existing, unique and not default rules
//        uint64 ruleIDs<>;
//        //: Indicates whether or not a rule can be modified in the future
//        bool isReadOnly;
//        //: Arbitrary stringified json object with details to attach to the role
//        longstring details;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type CreateSignerRoleData struct {
	RuleIDs    []Uint64                `json:"ruleIDs,omitempty"`
	IsReadOnly bool                    `json:"isReadOnly,omitempty"`
	Details    Longstring              `json:"details,omitempty"`
	Ext        CreateSignerRoleDataExt `json:"ext,omitempty"`
}

// UpdateSignerRoleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateSignerRoleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateSignerRoleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateSignerRoleDataExt
func (u UpdateSignerRoleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateSignerRoleDataExt creates a new  UpdateSignerRoleDataExt.
func NewUpdateSignerRoleDataExt(v LedgerVersion, value interface{}) (result UpdateSignerRoleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateSignerRoleData is an XDR Struct defines as:
//
//   //: UpdateSignerRoleData is used to pass necessary params to update an existing signer role
//    struct UpdateSignerRoleData
//    {
//        //: ID of an existing signer role
//        uint64 roleID;
//        //: Array of ids of existing, unique and not default rules
//        uint64 ruleIDs<>;
//
//        //: Arbitrary stringified json object with details to attach to the role
//        longstring details;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type UpdateSignerRoleData struct {
	RoleId  Uint64                  `json:"roleID,omitempty"`
	RuleIDs []Uint64                `json:"ruleIDs,omitempty"`
	Details Longstring              `json:"details,omitempty"`
	Ext     UpdateSignerRoleDataExt `json:"ext,omitempty"`
}

// RemoveSignerRoleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type RemoveSignerRoleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveSignerRoleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveSignerRoleDataExt
func (u RemoveSignerRoleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRemoveSignerRoleDataExt creates a new  RemoveSignerRoleDataExt.
func NewRemoveSignerRoleDataExt(v LedgerVersion, value interface{}) (result RemoveSignerRoleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RemoveSignerRoleData is an XDR Struct defines as:
//
//   //: RemoveSignerRoleData is used to pass necessary params to remove existing signer role
//    struct RemoveSignerRoleData
//    {
//        //: Identifier of an existing signer role
//        uint64 roleID;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type RemoveSignerRoleData struct {
	RoleId Uint64                  `json:"roleID,omitempty"`
	Ext    RemoveSignerRoleDataExt `json:"ext,omitempty"`
}

// ManageSignerRoleOpData is an XDR NestedUnion defines as:
//
//   union switch (ManageSignerRoleAction action)
//        {
//        case CREATE:
//            CreateSignerRoleData createData;
//        case UPDATE:
//            UpdateSignerRoleData updateData;
//        case REMOVE:
//            RemoveSignerRoleData removeData;
//        }
//
type ManageSignerRoleOpData struct {
	Action     ManageSignerRoleAction `json:"action,omitempty"`
	CreateData *CreateSignerRoleData  `json:"createData,omitempty"`
	UpdateData *UpdateSignerRoleData  `json:"updateData,omitempty"`
	RemoveData *RemoveSignerRoleData  `json:"removeData,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerRoleOpData) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerRoleOpData
func (u ManageSignerRoleOpData) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSignerRoleAction(sw) {
	case ManageSignerRoleActionCreate:
		return "CreateData", true
	case ManageSignerRoleActionUpdate:
		return "UpdateData", true
	case ManageSignerRoleActionRemove:
		return "RemoveData", true
	}
	return "-", false
}

// NewManageSignerRoleOpData creates a new  ManageSignerRoleOpData.
func NewManageSignerRoleOpData(action ManageSignerRoleAction, value interface{}) (result ManageSignerRoleOpData, err error) {
	result.Action = action
	switch ManageSignerRoleAction(action) {
	case ManageSignerRoleActionCreate:
		tv, ok := value.(CreateSignerRoleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateSignerRoleData")
			return
		}
		result.CreateData = &tv
	case ManageSignerRoleActionUpdate:
		tv, ok := value.(UpdateSignerRoleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateSignerRoleData")
			return
		}
		result.UpdateData = &tv
	case ManageSignerRoleActionRemove:
		tv, ok := value.(RemoveSignerRoleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveSignerRoleData")
			return
		}
		result.RemoveData = &tv
	}
	return
}

// MustCreateData retrieves the CreateData value from the union,
// panicing if the value is not set.
func (u ManageSignerRoleOpData) MustCreateData() CreateSignerRoleData {
	val, ok := u.GetCreateData()

	if !ok {
		panic("arm CreateData is not set")
	}

	return val
}

// GetCreateData retrieves the CreateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRoleOpData) GetCreateData() (result CreateSignerRoleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "CreateData" {
		result = *u.CreateData
		ok = true
	}

	return
}

// MustUpdateData retrieves the UpdateData value from the union,
// panicing if the value is not set.
func (u ManageSignerRoleOpData) MustUpdateData() UpdateSignerRoleData {
	val, ok := u.GetUpdateData()

	if !ok {
		panic("arm UpdateData is not set")
	}

	return val
}

// GetUpdateData retrieves the UpdateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRoleOpData) GetUpdateData() (result UpdateSignerRoleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateData" {
		result = *u.UpdateData
		ok = true
	}

	return
}

// MustRemoveData retrieves the RemoveData value from the union,
// panicing if the value is not set.
func (u ManageSignerRoleOpData) MustRemoveData() RemoveSignerRoleData {
	val, ok := u.GetRemoveData()

	if !ok {
		panic("arm RemoveData is not set")
	}

	return val
}

// GetRemoveData retrieves the RemoveData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRoleOpData) GetRemoveData() (result RemoveSignerRoleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "RemoveData" {
		result = *u.RemoveData
		ok = true
	}

	return
}

// ManageSignerRoleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageSignerRoleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerRoleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerRoleOpExt
func (u ManageSignerRoleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageSignerRoleOpExt creates a new  ManageSignerRoleOpExt.
func NewManageSignerRoleOpExt(v LedgerVersion, value interface{}) (result ManageSignerRoleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageSignerRoleOp is an XDR Struct defines as:
//
//   //: ManageSignerRoleOp is used to create, update or remove a signer role
//    struct ManageSignerRoleOp
//    {
//        //: data is used to pass one of `ManageSignerRoleAction` with required params
//        union switch (ManageSignerRoleAction action)
//        {
//        case CREATE:
//            CreateSignerRoleData createData;
//        case UPDATE:
//            UpdateSignerRoleData updateData;
//        case REMOVE:
//            RemoveSignerRoleData removeData;
//        } data;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageSignerRoleOp struct {
	Data ManageSignerRoleOpData `json:"data,omitempty"`
	Ext  ManageSignerRoleOpExt  `json:"ext,omitempty"`
}

// ManageSignerRoleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerRoleResultCode
//    enum ManageSignerRoleResultCode
//    {
//        //: Means that the specified action in `data` of ManageSignerRoleOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: There is no signer role with such id or the source cannot manage a role
//        NOT_FOUND = -1, // does not exist or owner mismatched
//        //: It is not allowed to remove role if it is attached to at least one singer
//        ROLE_IS_USED = -2,
//        //: Passed details have invalid json structure
//        INVALID_DETAILS = -3,
//        //: There is no rule with id passed through `ruleIDs`
//        NO_SUCH_RULE = -4,
//        //: It is not allowed to duplicate ids in `ruleIDs` array
//        RULE_ID_DUPLICATION = -5,
//        //: It is not allowed to pass ids of default rules on `ruleIDs` array
//        DEFAULT_RULE_ID_DUPLICATION = -6,
//        //: It is not allowed to pass ruleIDs that are more than maxSignerRuleCount (by default, 128)
//        TOO_MANY_RULE_IDS = -7
//    };
//
type ManageSignerRoleResultCode int32

const (
	ManageSignerRoleResultCodeSuccess                  ManageSignerRoleResultCode = 0
	ManageSignerRoleResultCodeNotFound                 ManageSignerRoleResultCode = -1
	ManageSignerRoleResultCodeRoleIsUsed               ManageSignerRoleResultCode = -2
	ManageSignerRoleResultCodeInvalidDetails           ManageSignerRoleResultCode = -3
	ManageSignerRoleResultCodeNoSuchRule               ManageSignerRoleResultCode = -4
	ManageSignerRoleResultCodeRuleIdDuplication        ManageSignerRoleResultCode = -5
	ManageSignerRoleResultCodeDefaultRuleIdDuplication ManageSignerRoleResultCode = -6
	ManageSignerRoleResultCodeTooManyRuleIds           ManageSignerRoleResultCode = -7
)

var ManageSignerRoleResultCodeAll = []ManageSignerRoleResultCode{
	ManageSignerRoleResultCodeSuccess,
	ManageSignerRoleResultCodeNotFound,
	ManageSignerRoleResultCodeRoleIsUsed,
	ManageSignerRoleResultCodeInvalidDetails,
	ManageSignerRoleResultCodeNoSuchRule,
	ManageSignerRoleResultCodeRuleIdDuplication,
	ManageSignerRoleResultCodeDefaultRuleIdDuplication,
	ManageSignerRoleResultCodeTooManyRuleIds,
}

var manageSignerRoleResultCodeMap = map[int32]string{
	0:  "ManageSignerRoleResultCodeSuccess",
	-1: "ManageSignerRoleResultCodeNotFound",
	-2: "ManageSignerRoleResultCodeRoleIsUsed",
	-3: "ManageSignerRoleResultCodeInvalidDetails",
	-4: "ManageSignerRoleResultCodeNoSuchRule",
	-5: "ManageSignerRoleResultCodeRuleIdDuplication",
	-6: "ManageSignerRoleResultCodeDefaultRuleIdDuplication",
	-7: "ManageSignerRoleResultCodeTooManyRuleIds",
}

var manageSignerRoleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "role_is_used",
	-3: "invalid_details",
	-4: "no_such_rule",
	-5: "rule_id_duplication",
	-6: "default_rule_id_duplication",
	-7: "too_many_rule_ids",
}

var manageSignerRoleResultCodeRevMap = map[string]int32{
	"ManageSignerRoleResultCodeSuccess":                  0,
	"ManageSignerRoleResultCodeNotFound":                 -1,
	"ManageSignerRoleResultCodeRoleIsUsed":               -2,
	"ManageSignerRoleResultCodeInvalidDetails":           -3,
	"ManageSignerRoleResultCodeNoSuchRule":               -4,
	"ManageSignerRoleResultCodeRuleIdDuplication":        -5,
	"ManageSignerRoleResultCodeDefaultRuleIdDuplication": -6,
	"ManageSignerRoleResultCodeTooManyRuleIds":           -7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageSignerRoleResultCode
func (e ManageSignerRoleResultCode) ValidEnum(v int32) bool {
	_, ok := manageSignerRoleResultCodeMap[v]
	return ok
}
func (e ManageSignerRoleResultCode) isFlag() bool {
	for i := len(ManageSignerRoleResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageSignerRoleResultCode(2) << uint64(len(ManageSignerRoleResultCodeAll)-1) >> uint64(len(ManageSignerRoleResultCodeAll)-i)
		if expected != ManageSignerRoleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageSignerRoleResultCode) String() string {
	name, _ := manageSignerRoleResultCodeMap[int32(e)]
	return name
}

func (e ManageSignerRoleResultCode) ShortString() string {
	name, _ := manageSignerRoleResultCodeShortMap[int32(e)]
	return name
}

func (e ManageSignerRoleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageSignerRoleResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageSignerRoleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageSignerRoleResultCode(t.Value)
	return nil
}

// ManageSignerRoleResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type ManageSignerRoleResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerRoleResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerRoleResultSuccessExt
func (u ManageSignerRoleResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageSignerRoleResultSuccessExt creates a new  ManageSignerRoleResultSuccessExt.
func NewManageSignerRoleResultSuccessExt(v LedgerVersion, value interface{}) (result ManageSignerRoleResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageSignerRoleResultSuccess is an XDR NestedStruct defines as:
//
//   struct
//            {
//                //: id of a role that was managed
//                uint64 roleID;
//
//                //: reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//            }
//
type ManageSignerRoleResultSuccess struct {
	RoleId Uint64                           `json:"roleID,omitempty"`
	Ext    ManageSignerRoleResultSuccessExt `json:"ext,omitempty"`
}

// ManageSignerRoleResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union ManageSignerRoleResult switch (ManageSignerRoleResultCode code)
//    {
//        case SUCCESS:
//            struct
//            {
//                //: id of a role that was managed
//                uint64 roleID;
//
//                //: reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//            } success;
//        case RULE_ID_DUPLICATION:
//        case DEFAULT_RULE_ID_DUPLICATION:
//        case NO_SUCH_RULE:
//            //: ID of a rule that was either duplicated or is default or does not exist
//            uint64 ruleID;
//        case TOO_MANY_RULE_IDS:
//            //: max count of rule ids that can be passed in `ruleIDs` array
//            uint64 maxRuleIDsCount;
//        default:
//            void;
//    };
//
type ManageSignerRoleResult struct {
	Code            ManageSignerRoleResultCode     `json:"code,omitempty"`
	Success         *ManageSignerRoleResultSuccess `json:"success,omitempty"`
	RuleId          *Uint64                        `json:"ruleID,omitempty"`
	MaxRuleIDsCount *Uint64                        `json:"maxRuleIDsCount,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerRoleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerRoleResult
func (u ManageSignerRoleResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSignerRoleResultCode(sw) {
	case ManageSignerRoleResultCodeSuccess:
		return "Success", true
	case ManageSignerRoleResultCodeRuleIdDuplication:
		return "RuleId", true
	case ManageSignerRoleResultCodeDefaultRuleIdDuplication:
		return "RuleId", true
	case ManageSignerRoleResultCodeNoSuchRule:
		return "RuleId", true
	case ManageSignerRoleResultCodeTooManyRuleIds:
		return "MaxRuleIDsCount", true
	default:
		return "", true
	}
}

// NewManageSignerRoleResult creates a new  ManageSignerRoleResult.
func NewManageSignerRoleResult(code ManageSignerRoleResultCode, value interface{}) (result ManageSignerRoleResult, err error) {
	result.Code = code
	switch ManageSignerRoleResultCode(code) {
	case ManageSignerRoleResultCodeSuccess:
		tv, ok := value.(ManageSignerRoleResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSignerRoleResultSuccess")
			return
		}
		result.Success = &tv
	case ManageSignerRoleResultCodeRuleIdDuplication:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RuleId = &tv
	case ManageSignerRoleResultCodeDefaultRuleIdDuplication:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RuleId = &tv
	case ManageSignerRoleResultCodeNoSuchRule:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RuleId = &tv
	case ManageSignerRoleResultCodeTooManyRuleIds:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.MaxRuleIDsCount = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageSignerRoleResult) MustSuccess() ManageSignerRoleResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRoleResult) GetSuccess() (result ManageSignerRoleResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustRuleId retrieves the RuleId value from the union,
// panicing if the value is not set.
func (u ManageSignerRoleResult) MustRuleId() Uint64 {
	val, ok := u.GetRuleId()

	if !ok {
		panic("arm RuleId is not set")
	}

	return val
}

// GetRuleId retrieves the RuleId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRoleResult) GetRuleId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RuleId" {
		result = *u.RuleId
		ok = true
	}

	return
}

// MustMaxRuleIDsCount retrieves the MaxRuleIDsCount value from the union,
// panicing if the value is not set.
func (u ManageSignerRoleResult) MustMaxRuleIDsCount() Uint64 {
	val, ok := u.GetMaxRuleIDsCount()

	if !ok {
		panic("arm MaxRuleIDsCount is not set")
	}

	return val
}

// GetMaxRuleIDsCount retrieves the MaxRuleIDsCount value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRoleResult) GetMaxRuleIDsCount() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "MaxRuleIDsCount" {
		result = *u.MaxRuleIDsCount
		ok = true
	}

	return
}

// ManageSignerRuleAction is an XDR Enum defines as:
//
//   //: Actions that can be performed with a signer rule
//    enum ManageSignerRuleAction
//    {
//        CREATE = 0,
//        UPDATE = 1,
//        REMOVE = 2
//    };
//
type ManageSignerRuleAction int32

const (
	ManageSignerRuleActionCreate ManageSignerRuleAction = 0
	ManageSignerRuleActionUpdate ManageSignerRuleAction = 1
	ManageSignerRuleActionRemove ManageSignerRuleAction = 2
)

var ManageSignerRuleActionAll = []ManageSignerRuleAction{
	ManageSignerRuleActionCreate,
	ManageSignerRuleActionUpdate,
	ManageSignerRuleActionRemove,
}

var manageSignerRuleActionMap = map[int32]string{
	0: "ManageSignerRuleActionCreate",
	1: "ManageSignerRuleActionUpdate",
	2: "ManageSignerRuleActionRemove",
}

var manageSignerRuleActionShortMap = map[int32]string{
	0: "create",
	1: "update",
	2: "remove",
}

var manageSignerRuleActionRevMap = map[string]int32{
	"ManageSignerRuleActionCreate": 0,
	"ManageSignerRuleActionUpdate": 1,
	"ManageSignerRuleActionRemove": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageSignerRuleAction
func (e ManageSignerRuleAction) ValidEnum(v int32) bool {
	_, ok := manageSignerRuleActionMap[v]
	return ok
}
func (e ManageSignerRuleAction) isFlag() bool {
	for i := len(ManageSignerRuleActionAll) - 1; i >= 0; i-- {
		expected := ManageSignerRuleAction(2) << uint64(len(ManageSignerRuleActionAll)-1) >> uint64(len(ManageSignerRuleActionAll)-i)
		if expected != ManageSignerRuleActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageSignerRuleAction) String() string {
	name, _ := manageSignerRuleActionMap[int32(e)]
	return name
}

func (e ManageSignerRuleAction) ShortString() string {
	name, _ := manageSignerRuleActionShortMap[int32(e)]
	return name
}

func (e ManageSignerRuleAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageSignerRuleActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageSignerRuleAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageSignerRuleAction(t.Value)
	return nil
}

// CreateSignerRuleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateSignerRuleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateSignerRuleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateSignerRuleDataExt
func (u CreateSignerRuleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateSignerRuleDataExt creates a new  CreateSignerRuleDataExt.
func NewCreateSignerRuleDataExt(v LedgerVersion, value interface{}) (result CreateSignerRuleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateSignerRuleData is an XDR Struct defines as:
//
//   //: CreateSignerRuleData is used to pass necessary params to create a new signer rule
//    struct CreateSignerRuleData
//    {
//        //: Resource is used to specify an entity (for some, with properties) that can be managed through operations
//        SignerRuleResource resource;
//        //: Value from enum that can be applied to `resource`
//        SignerRuleAction action;
//        //: Indicate whether or not an `action` on the provided `resource` is prohibited
//        bool forbids;
//        //: True means that such rule will be automatically added to each new or updated signer role
//        bool isDefault;
//        //: Indicates whether or not a rule can be modified in the future
//        bool isReadOnly;
//        //: Arbitrary stringified json object with details that will be attached to a rule
//        longstring details;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type CreateSignerRuleData struct {
	Resource   SignerRuleResource      `json:"resource,omitempty"`
	Action     SignerRuleAction        `json:"action,omitempty"`
	Forbids    bool                    `json:"forbids,omitempty"`
	IsDefault  bool                    `json:"isDefault,omitempty"`
	IsReadOnly bool                    `json:"isReadOnly,omitempty"`
	Details    Longstring              `json:"details,omitempty"`
	Ext        CreateSignerRuleDataExt `json:"ext,omitempty"`
}

// UpdateSignerRuleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateSignerRuleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateSignerRuleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateSignerRuleDataExt
func (u UpdateSignerRuleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateSignerRuleDataExt creates a new  UpdateSignerRuleDataExt.
func NewUpdateSignerRuleDataExt(v LedgerVersion, value interface{}) (result UpdateSignerRuleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateSignerRuleData is an XDR Struct defines as:
//
//   //: UpdateSignerRuleData is used to pass necessary params to update an existing signer rule
//    struct UpdateSignerRuleData
//    {
//        //: Identifier of an existing signer rule
//        uint64 ruleID;
//        //: Resource is used to specify entity (for some, with properties) that can be managed through operations
//        SignerRuleResource resource;
//        //: Value from enum that can be applied to `resource`
//        SignerRuleAction action;
//        //: True means that such rule will be automatically added to each new or updated signer role
//        bool forbids;
//        //: True means that no one can manage such rule after creating
//        bool isDefault;
//        //: Arbitrary stringified json object with details that will be attached to a rule
//        longstring details;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type UpdateSignerRuleData struct {
	RuleId    Uint64                  `json:"ruleID,omitempty"`
	Resource  SignerRuleResource      `json:"resource,omitempty"`
	Action    SignerRuleAction        `json:"action,omitempty"`
	Forbids   bool                    `json:"forbids,omitempty"`
	IsDefault bool                    `json:"isDefault,omitempty"`
	Details   Longstring              `json:"details,omitempty"`
	Ext       UpdateSignerRuleDataExt `json:"ext,omitempty"`
}

// RemoveSignerRuleDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type RemoveSignerRuleDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveSignerRuleDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveSignerRuleDataExt
func (u RemoveSignerRuleDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRemoveSignerRuleDataExt creates a new  RemoveSignerRuleDataExt.
func NewRemoveSignerRuleDataExt(v LedgerVersion, value interface{}) (result RemoveSignerRuleDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RemoveSignerRuleData is an XDR Struct defines as:
//
//   //: RemoveSignerRuleData is used to pass necessary params to remove existing signer rule
//    struct RemoveSignerRuleData
//    {
//        //: Identifier of an existing signer rule
//        uint64 ruleID;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type RemoveSignerRuleData struct {
	RuleId Uint64                  `json:"ruleID,omitempty"`
	Ext    RemoveSignerRuleDataExt `json:"ext,omitempty"`
}

// ManageSignerRuleOpData is an XDR NestedUnion defines as:
//
//   union switch (ManageSignerRuleAction action)
//        {
//        case CREATE:
//            CreateSignerRuleData createData;
//        case UPDATE:
//            UpdateSignerRuleData updateData;
//        case REMOVE:
//            RemoveSignerRuleData removeData;
//        }
//
type ManageSignerRuleOpData struct {
	Action     ManageSignerRuleAction `json:"action,omitempty"`
	CreateData *CreateSignerRuleData  `json:"createData,omitempty"`
	UpdateData *UpdateSignerRuleData  `json:"updateData,omitempty"`
	RemoveData *RemoveSignerRuleData  `json:"removeData,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerRuleOpData) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerRuleOpData
func (u ManageSignerRuleOpData) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSignerRuleAction(sw) {
	case ManageSignerRuleActionCreate:
		return "CreateData", true
	case ManageSignerRuleActionUpdate:
		return "UpdateData", true
	case ManageSignerRuleActionRemove:
		return "RemoveData", true
	}
	return "-", false
}

// NewManageSignerRuleOpData creates a new  ManageSignerRuleOpData.
func NewManageSignerRuleOpData(action ManageSignerRuleAction, value interface{}) (result ManageSignerRuleOpData, err error) {
	result.Action = action
	switch ManageSignerRuleAction(action) {
	case ManageSignerRuleActionCreate:
		tv, ok := value.(CreateSignerRuleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateSignerRuleData")
			return
		}
		result.CreateData = &tv
	case ManageSignerRuleActionUpdate:
		tv, ok := value.(UpdateSignerRuleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateSignerRuleData")
			return
		}
		result.UpdateData = &tv
	case ManageSignerRuleActionRemove:
		tv, ok := value.(RemoveSignerRuleData)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveSignerRuleData")
			return
		}
		result.RemoveData = &tv
	}
	return
}

// MustCreateData retrieves the CreateData value from the union,
// panicing if the value is not set.
func (u ManageSignerRuleOpData) MustCreateData() CreateSignerRuleData {
	val, ok := u.GetCreateData()

	if !ok {
		panic("arm CreateData is not set")
	}

	return val
}

// GetCreateData retrieves the CreateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRuleOpData) GetCreateData() (result CreateSignerRuleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "CreateData" {
		result = *u.CreateData
		ok = true
	}

	return
}

// MustUpdateData retrieves the UpdateData value from the union,
// panicing if the value is not set.
func (u ManageSignerRuleOpData) MustUpdateData() UpdateSignerRuleData {
	val, ok := u.GetUpdateData()

	if !ok {
		panic("arm UpdateData is not set")
	}

	return val
}

// GetUpdateData retrieves the UpdateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRuleOpData) GetUpdateData() (result UpdateSignerRuleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateData" {
		result = *u.UpdateData
		ok = true
	}

	return
}

// MustRemoveData retrieves the RemoveData value from the union,
// panicing if the value is not set.
func (u ManageSignerRuleOpData) MustRemoveData() RemoveSignerRuleData {
	val, ok := u.GetRemoveData()

	if !ok {
		panic("arm RemoveData is not set")
	}

	return val
}

// GetRemoveData retrieves the RemoveData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRuleOpData) GetRemoveData() (result RemoveSignerRuleData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "RemoveData" {
		result = *u.RemoveData
		ok = true
	}

	return
}

// ManageSignerRuleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageSignerRuleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerRuleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerRuleOpExt
func (u ManageSignerRuleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageSignerRuleOpExt creates a new  ManageSignerRuleOpExt.
func NewManageSignerRuleOpExt(v LedgerVersion, value interface{}) (result ManageSignerRuleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageSignerRuleOp is an XDR Struct defines as:
//
//   //: ManageSignerRuleOp is used to create, update or remove signer rule
//    struct ManageSignerRuleOp
//    {
//        //: data is used to pass one of `ManageSignerRuleAction` with required params
//        union switch (ManageSignerRuleAction action)
//        {
//        case CREATE:
//            CreateSignerRuleData createData;
//        case UPDATE:
//            UpdateSignerRuleData updateData;
//        case REMOVE:
//            RemoveSignerRuleData removeData;
//        } data;
//
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageSignerRuleOp struct {
	Data ManageSignerRuleOpData `json:"data,omitempty"`
	Ext  ManageSignerRuleOpExt  `json:"ext,omitempty"`
}

// ManageSignerRuleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerRuleOp
//    enum ManageSignerRuleResultCode
//    {
//        //: Specified action in `data` of ManageSignerRuleOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: There is no signer rule with such id or source cannot manage the rule
//        NOT_FOUND = -1, // does not exists or owner mismatched
//        //: It is not allowed to remove the rule if it is attached to at least one role
//        RULE_IS_USED = -2,
//        //: Passed details have invalid json structure
//        INVALID_DETAILS = -3
//    };
//
type ManageSignerRuleResultCode int32

const (
	ManageSignerRuleResultCodeSuccess        ManageSignerRuleResultCode = 0
	ManageSignerRuleResultCodeNotFound       ManageSignerRuleResultCode = -1
	ManageSignerRuleResultCodeRuleIsUsed     ManageSignerRuleResultCode = -2
	ManageSignerRuleResultCodeInvalidDetails ManageSignerRuleResultCode = -3
)

var ManageSignerRuleResultCodeAll = []ManageSignerRuleResultCode{
	ManageSignerRuleResultCodeSuccess,
	ManageSignerRuleResultCodeNotFound,
	ManageSignerRuleResultCodeRuleIsUsed,
	ManageSignerRuleResultCodeInvalidDetails,
}

var manageSignerRuleResultCodeMap = map[int32]string{
	0:  "ManageSignerRuleResultCodeSuccess",
	-1: "ManageSignerRuleResultCodeNotFound",
	-2: "ManageSignerRuleResultCodeRuleIsUsed",
	-3: "ManageSignerRuleResultCodeInvalidDetails",
}

var manageSignerRuleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "rule_is_used",
	-3: "invalid_details",
}

var manageSignerRuleResultCodeRevMap = map[string]int32{
	"ManageSignerRuleResultCodeSuccess":        0,
	"ManageSignerRuleResultCodeNotFound":       -1,
	"ManageSignerRuleResultCodeRuleIsUsed":     -2,
	"ManageSignerRuleResultCodeInvalidDetails": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageSignerRuleResultCode
func (e ManageSignerRuleResultCode) ValidEnum(v int32) bool {
	_, ok := manageSignerRuleResultCodeMap[v]
	return ok
}
func (e ManageSignerRuleResultCode) isFlag() bool {
	for i := len(ManageSignerRuleResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageSignerRuleResultCode(2) << uint64(len(ManageSignerRuleResultCodeAll)-1) >> uint64(len(ManageSignerRuleResultCodeAll)-i)
		if expected != ManageSignerRuleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageSignerRuleResultCode) String() string {
	name, _ := manageSignerRuleResultCodeMap[int32(e)]
	return name
}

func (e ManageSignerRuleResultCode) ShortString() string {
	name, _ := manageSignerRuleResultCodeShortMap[int32(e)]
	return name
}

func (e ManageSignerRuleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageSignerRuleResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageSignerRuleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageSignerRuleResultCode(t.Value)
	return nil
}

// ManageSignerRuleResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//
type ManageSignerRuleResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerRuleResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerRuleResultSuccessExt
func (u ManageSignerRuleResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageSignerRuleResultSuccessExt creates a new  ManageSignerRuleResultSuccessExt.
func NewManageSignerRuleResultSuccessExt(v LedgerVersion, value interface{}) (result ManageSignerRuleResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageSignerRuleResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//                //: id of the rule that was managed
//                uint64 ruleID;
//
//                //: reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//            }
//
type ManageSignerRuleResultSuccess struct {
	RuleId Uint64                           `json:"ruleID,omitempty"`
	Ext    ManageSignerRuleResultSuccessExt `json:"ext,omitempty"`
}

// ManageSignerRuleResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union ManageSignerRuleResult switch (ManageSignerRuleResultCode code)
//    {
//        case SUCCESS:
//            struct {
//                //: id of the rule that was managed
//                uint64 ruleID;
//
//                //: reserved for future use
//                union switch (LedgerVersion v)
//                {
//                case EMPTY_VERSION:
//                    void;
//                }
//                ext;
//            } success;
//        case RULE_IS_USED:
//            //: ids of roles which use a rule that cannot be removed
//            uint64 roleIDs<>;
//        default:
//            void;
//    };
//
type ManageSignerRuleResult struct {
	Code    ManageSignerRuleResultCode     `json:"code,omitempty"`
	Success *ManageSignerRuleResultSuccess `json:"success,omitempty"`
	RoleIDs *[]Uint64                      `json:"roleIDs,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerRuleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerRuleResult
func (u ManageSignerRuleResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSignerRuleResultCode(sw) {
	case ManageSignerRuleResultCodeSuccess:
		return "Success", true
	case ManageSignerRuleResultCodeRuleIsUsed:
		return "RoleIDs", true
	default:
		return "", true
	}
}

// NewManageSignerRuleResult creates a new  ManageSignerRuleResult.
func NewManageSignerRuleResult(code ManageSignerRuleResultCode, value interface{}) (result ManageSignerRuleResult, err error) {
	result.Code = code
	switch ManageSignerRuleResultCode(code) {
	case ManageSignerRuleResultCodeSuccess:
		tv, ok := value.(ManageSignerRuleResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSignerRuleResultSuccess")
			return
		}
		result.Success = &tv
	case ManageSignerRuleResultCodeRuleIsUsed:
		tv, ok := value.([]Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be []Uint64")
			return
		}
		result.RoleIDs = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageSignerRuleResult) MustSuccess() ManageSignerRuleResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRuleResult) GetSuccess() (result ManageSignerRuleResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustRoleIDs retrieves the RoleIDs value from the union,
// panicing if the value is not set.
func (u ManageSignerRuleResult) MustRoleIDs() []Uint64 {
	val, ok := u.GetRoleIDs()

	if !ok {
		panic("arm RoleIDs is not set")
	}

	return val
}

// GetRoleIDs retrieves the RoleIDs value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerRuleResult) GetRoleIDs() (result []Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RoleIDs" {
		result = *u.RoleIDs
		ok = true
	}

	return
}

// ManageSignerAction is an XDR Enum defines as:
//
//   //: Actions that can be applied to a signer
//    enum ManageSignerAction
//    {
//        CREATE = 0,
//        UPDATE = 1,
//        REMOVE = 2
//    };
//
type ManageSignerAction int32

const (
	ManageSignerActionCreate ManageSignerAction = 0
	ManageSignerActionUpdate ManageSignerAction = 1
	ManageSignerActionRemove ManageSignerAction = 2
)

var ManageSignerActionAll = []ManageSignerAction{
	ManageSignerActionCreate,
	ManageSignerActionUpdate,
	ManageSignerActionRemove,
}

var manageSignerActionMap = map[int32]string{
	0: "ManageSignerActionCreate",
	1: "ManageSignerActionUpdate",
	2: "ManageSignerActionRemove",
}

var manageSignerActionShortMap = map[int32]string{
	0: "create",
	1: "update",
	2: "remove",
}

var manageSignerActionRevMap = map[string]int32{
	"ManageSignerActionCreate": 0,
	"ManageSignerActionUpdate": 1,
	"ManageSignerActionRemove": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageSignerAction
func (e ManageSignerAction) ValidEnum(v int32) bool {
	_, ok := manageSignerActionMap[v]
	return ok
}
func (e ManageSignerAction) isFlag() bool {
	for i := len(ManageSignerActionAll) - 1; i >= 0; i-- {
		expected := ManageSignerAction(2) << uint64(len(ManageSignerActionAll)-1) >> uint64(len(ManageSignerActionAll)-i)
		if expected != ManageSignerActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageSignerAction) String() string {
	name, _ := manageSignerActionMap[int32(e)]
	return name
}

func (e ManageSignerAction) ShortString() string {
	name, _ := manageSignerActionShortMap[int32(e)]
	return name
}

func (e ManageSignerAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageSignerActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageSignerAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageSignerAction(t.Value)
	return nil
}

// SignerData is an XDR Struct defines as:
//
//   //: SignerData is used to pass necessary data to create or update the signer
//    struct SignerData
//    {
//        //: Public key of a signer
//        PublicKey publicKey;
//        //: id of the role that will be attached to a signer
//        uint64 roleID;
//
//        //: weight that signer will have, threshold for all SignerRequirements equals 1000
//        uint32 weight;
//        //: If there are some signers with equal identity, only one signer will be chosen
//        //: (either the one with the biggest weight or the one who was the first to satisfy a threshold)
//        uint32 identity;
//
//        //: Arbitrary stringified json object with details that will be attached to signer
//        longstring details;
//
//        //: reserved for future extension
//        EmptyExt ext;
//    };
//
type SignerData struct {
	PublicKey PublicKey  `json:"publicKey,omitempty"`
	RoleId    Uint64     `json:"roleID,omitempty"`
	Weight    Uint32     `json:"weight,omitempty"`
	Identity  Uint32     `json:"identity,omitempty"`
	Details   Longstring `json:"details,omitempty"`
	Ext       EmptyExt   `json:"ext,omitempty"`
}

// RemoveSignerData is an XDR Struct defines as:
//
//   //: RemoveSignerData is used to pass necessary data to remove a signer
//    struct RemoveSignerData
//    {
//        //: Public key of an existing signer
//        PublicKey publicKey;
//
//        //: reserved for future extension
//        EmptyExt ext;
//    };
//
type RemoveSignerData struct {
	PublicKey PublicKey `json:"publicKey,omitempty"`
	Ext       EmptyExt  `json:"ext,omitempty"`
}

// ManageSignerOpData is an XDR NestedUnion defines as:
//
//   union switch (ManageSignerAction action)
//        {
//        case CREATE:
//            SignerData createData;
//        case UPDATE:
//            SignerData updateData;
//        case REMOVE:
//            RemoveSignerData removeData;
//        }
//
type ManageSignerOpData struct {
	Action     ManageSignerAction `json:"action,omitempty"`
	CreateData *SignerData        `json:"createData,omitempty"`
	UpdateData *SignerData        `json:"updateData,omitempty"`
	RemoveData *RemoveSignerData  `json:"removeData,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerOpData) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerOpData
func (u ManageSignerOpData) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSignerAction(sw) {
	case ManageSignerActionCreate:
		return "CreateData", true
	case ManageSignerActionUpdate:
		return "UpdateData", true
	case ManageSignerActionRemove:
		return "RemoveData", true
	}
	return "-", false
}

// NewManageSignerOpData creates a new  ManageSignerOpData.
func NewManageSignerOpData(action ManageSignerAction, value interface{}) (result ManageSignerOpData, err error) {
	result.Action = action
	switch ManageSignerAction(action) {
	case ManageSignerActionCreate:
		tv, ok := value.(SignerData)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerData")
			return
		}
		result.CreateData = &tv
	case ManageSignerActionUpdate:
		tv, ok := value.(SignerData)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerData")
			return
		}
		result.UpdateData = &tv
	case ManageSignerActionRemove:
		tv, ok := value.(RemoveSignerData)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveSignerData")
			return
		}
		result.RemoveData = &tv
	}
	return
}

// MustCreateData retrieves the CreateData value from the union,
// panicing if the value is not set.
func (u ManageSignerOpData) MustCreateData() SignerData {
	val, ok := u.GetCreateData()

	if !ok {
		panic("arm CreateData is not set")
	}

	return val
}

// GetCreateData retrieves the CreateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerOpData) GetCreateData() (result SignerData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "CreateData" {
		result = *u.CreateData
		ok = true
	}

	return
}

// MustUpdateData retrieves the UpdateData value from the union,
// panicing if the value is not set.
func (u ManageSignerOpData) MustUpdateData() SignerData {
	val, ok := u.GetUpdateData()

	if !ok {
		panic("arm UpdateData is not set")
	}

	return val
}

// GetUpdateData retrieves the UpdateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerOpData) GetUpdateData() (result SignerData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateData" {
		result = *u.UpdateData
		ok = true
	}

	return
}

// MustRemoveData retrieves the RemoveData value from the union,
// panicing if the value is not set.
func (u ManageSignerOpData) MustRemoveData() RemoveSignerData {
	val, ok := u.GetRemoveData()

	if !ok {
		panic("arm RemoveData is not set")
	}

	return val
}

// GetRemoveData retrieves the RemoveData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerOpData) GetRemoveData() (result RemoveSignerData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "RemoveData" {
		result = *u.RemoveData
		ok = true
	}

	return
}

// ManageSignerOp is an XDR Struct defines as:
//
//   //: ManageSignerOp is used to create, update or remove a signer
//    struct ManageSignerOp
//    {
//        //: data is used to pass one of `ManageSignerAction` with required params
//        union switch (ManageSignerAction action)
//        {
//        case CREATE:
//            SignerData createData;
//        case UPDATE:
//            SignerData updateData;
//        case REMOVE:
//            RemoveSignerData removeData;
//        }
//        data;
//
//        //: reserved for future extension
//        EmptyExt ext;
//    };
//
type ManageSignerOp struct {
	Data ManageSignerOpData `json:"data,omitempty"`
	Ext  EmptyExt           `json:"ext,omitempty"`
}

// ManageSignerResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerOp
//    enum ManageSignerResultCode
//    {
//        //: Specified action in `data` of ManageSignerOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: Passed details have invalid json structure
//        INVALID_DETAILS = -1, // invalid json details
//        //: Signer with such public key is already attached to the source account
//        ALREADY_EXISTS = -2, // signer already exist
//        //: There is no role with such id
//        NO_SUCH_ROLE = -3,
//        //: It is not allowed to set weight more than 1000
//        INVALID_WEIGHT = -4, // more than 1000
//        //: Source account does not have a signer with the provided public key
//        NOT_FOUND = -5 // there is no signer with such public key
//    };
//
type ManageSignerResultCode int32

const (
	ManageSignerResultCodeSuccess        ManageSignerResultCode = 0
	ManageSignerResultCodeInvalidDetails ManageSignerResultCode = -1
	ManageSignerResultCodeAlreadyExists  ManageSignerResultCode = -2
	ManageSignerResultCodeNoSuchRole     ManageSignerResultCode = -3
	ManageSignerResultCodeInvalidWeight  ManageSignerResultCode = -4
	ManageSignerResultCodeNotFound       ManageSignerResultCode = -5
)

var ManageSignerResultCodeAll = []ManageSignerResultCode{
	ManageSignerResultCodeSuccess,
	ManageSignerResultCodeInvalidDetails,
	ManageSignerResultCodeAlreadyExists,
	ManageSignerResultCodeNoSuchRole,
	ManageSignerResultCodeInvalidWeight,
	ManageSignerResultCodeNotFound,
}

var manageSignerResultCodeMap = map[int32]string{
	0:  "ManageSignerResultCodeSuccess",
	-1: "ManageSignerResultCodeInvalidDetails",
	-2: "ManageSignerResultCodeAlreadyExists",
	-3: "ManageSignerResultCodeNoSuchRole",
	-4: "ManageSignerResultCodeInvalidWeight",
	-5: "ManageSignerResultCodeNotFound",
}

var manageSignerResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_details",
	-2: "already_exists",
	-3: "no_such_role",
	-4: "invalid_weight",
	-5: "not_found",
}

var manageSignerResultCodeRevMap = map[string]int32{
	"ManageSignerResultCodeSuccess":        0,
	"ManageSignerResultCodeInvalidDetails": -1,
	"ManageSignerResultCodeAlreadyExists":  -2,
	"ManageSignerResultCodeNoSuchRole":     -3,
	"ManageSignerResultCodeInvalidWeight":  -4,
	"ManageSignerResultCodeNotFound":       -5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageSignerResultCode
func (e ManageSignerResultCode) ValidEnum(v int32) bool {
	_, ok := manageSignerResultCodeMap[v]
	return ok
}
func (e ManageSignerResultCode) isFlag() bool {
	for i := len(ManageSignerResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageSignerResultCode(2) << uint64(len(ManageSignerResultCodeAll)-1) >> uint64(len(ManageSignerResultCodeAll)-i)
		if expected != ManageSignerResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageSignerResultCode) String() string {
	name, _ := manageSignerResultCodeMap[int32(e)]
	return name
}

func (e ManageSignerResultCode) ShortString() string {
	name, _ := manageSignerResultCodeShortMap[int32(e)]
	return name
}

func (e ManageSignerResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ManageSignerResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageSignerResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageSignerResultCode(t.Value)
	return nil
}

// ManageSignerResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union ManageSignerResult switch (ManageSignerResultCode code)
//    {
//    case SUCCESS:
//        //: reserved for future extension
//        EmptyExt ext;
//    default:
//        void;
//    };
//
type ManageSignerResult struct {
	Code ManageSignerResultCode `json:"code,omitempty"`
	Ext  *EmptyExt              `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSignerResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSignerResult
func (u ManageSignerResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSignerResultCode(sw) {
	case ManageSignerResultCodeSuccess:
		return "Ext", true
	default:
		return "", true
	}
}

// NewManageSignerResult creates a new  ManageSignerResult.
func NewManageSignerResult(code ManageSignerResultCode, value interface{}) (result ManageSignerResult, err error) {
	result.Code = code
	switch ManageSignerResultCode(code) {
	case ManageSignerResultCodeSuccess:
		tv, ok := value.(EmptyExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be EmptyExt")
			return
		}
		result.Ext = &tv
	default:
		// void
	}
	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u ManageSignerResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSignerResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// ReviewRequestOpAction is an XDR Enum defines as:
//
//   //: Actions that can be performed on request that is being reviewed
//    enum ReviewRequestOpAction {
//        //: Approve request
//        APPROVE = 1,
//        //: Reject request
//        REJECT = 2,
//        //: Permanently reject request
//        PERMANENT_REJECT = 3
//    };
//
type ReviewRequestOpAction int32

const (
	ReviewRequestOpActionApprove         ReviewRequestOpAction = 1
	ReviewRequestOpActionReject          ReviewRequestOpAction = 2
	ReviewRequestOpActionPermanentReject ReviewRequestOpAction = 3
)

var ReviewRequestOpActionAll = []ReviewRequestOpAction{
	ReviewRequestOpActionApprove,
	ReviewRequestOpActionReject,
	ReviewRequestOpActionPermanentReject,
}

var reviewRequestOpActionMap = map[int32]string{
	1: "ReviewRequestOpActionApprove",
	2: "ReviewRequestOpActionReject",
	3: "ReviewRequestOpActionPermanentReject",
}

var reviewRequestOpActionShortMap = map[int32]string{
	1: "approve",
	2: "reject",
	3: "permanent_reject",
}

var reviewRequestOpActionRevMap = map[string]int32{
	"ReviewRequestOpActionApprove":         1,
	"ReviewRequestOpActionReject":          2,
	"ReviewRequestOpActionPermanentReject": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewRequestOpAction
func (e ReviewRequestOpAction) ValidEnum(v int32) bool {
	_, ok := reviewRequestOpActionMap[v]
	return ok
}
func (e ReviewRequestOpAction) isFlag() bool {
	for i := len(ReviewRequestOpActionAll) - 1; i >= 0; i-- {
		expected := ReviewRequestOpAction(2) << uint64(len(ReviewRequestOpActionAll)-1) >> uint64(len(ReviewRequestOpActionAll)-i)
		if expected != ReviewRequestOpActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewRequestOpAction) String() string {
	name, _ := reviewRequestOpActionMap[int32(e)]
	return name
}

func (e ReviewRequestOpAction) ShortString() string {
	name, _ := reviewRequestOpActionShortMap[int32(e)]
	return name
}

func (e ReviewRequestOpAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ReviewRequestOpActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ReviewRequestOpAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewRequestOpAction(t.Value)
	return nil
}

// ReviewDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ReviewDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewDetailsExt
func (u ReviewDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewDetailsExt creates a new  ReviewDetailsExt.
func NewReviewDetailsExt(v LedgerVersion, value interface{}) (result ReviewDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewDetails is an XDR Struct defines as:
//
//   //: Details of a request review
//    struct ReviewDetails {
//        //: Tasks to add to pending
//        uint32 tasksToAdd;
//        //: Tasks to remove from pending
//        uint32 tasksToRemove;
//        //: Details of the current review
//        string externalDetails<>;
//        //: Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ReviewDetails struct {
	TasksToAdd      Uint32           `json:"tasksToAdd,omitempty"`
	TasksToRemove   Uint32           `json:"tasksToRemove,omitempty"`
	ExternalDetails string           `json:"externalDetails,omitempty"`
	Ext             ReviewDetailsExt `json:"ext,omitempty"`
}

// ExtendedResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ExtendedResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ExtendedResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ExtendedResultExt
func (u ExtendedResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewExtendedResultExt creates a new  ExtendedResultExt.
func NewExtendedResultExt(v LedgerVersion, value interface{}) (result ExtendedResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ExtendedResult is an XDR Struct defines as:
//
//   //: Extended result of a Review Request operation containing details specific to certain request types
//    struct ExtendedResult {
//        //: Indicates whether or not the request that is being reviewed was applied
//        bool fulfilled;
//        //: Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ExtendedResult struct {
	Fulfilled bool              `json:"fulfilled,omitempty"`
	Ext       ExtendedResultExt `json:"ext,omitempty"`
}

// ReviewRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ReviewRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestOpExt
func (u ReviewRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewRequestOpExt creates a new  ReviewRequestOpExt.
func NewReviewRequestOpExt(v LedgerVersion, value interface{}) (result ReviewRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewRequestOp is an XDR Struct defines as:
//
//   //: Review Request operation
//    struct ReviewRequestOp
//    {
//        //: ID of a request that is being reviewed
//        uint64 requestID;
//        //: Hash of a request that is being reviewed
//        Hash requestHash;
//        //: Type of request to review
//        ReviewableRequestType requestType;
//        //: Review action defines an action performed on the pending ReviewableRequest
//        ReviewRequestOpAction action;
//        //: Contains reject reason
//        longstring reason;
//        //: Details of the ReviewRequest operation
//        ReviewDetails reviewDetails;
//
//        //: Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ReviewRequestOp struct {
	RequestId     Uint64                `json:"requestID,omitempty"`
	RequestHash   Hash                  `json:"requestHash,omitempty"`
	RequestType   ReviewableRequestType `json:"requestType,omitempty"`
	Action        ReviewRequestOpAction `json:"action,omitempty"`
	Reason        Longstring            `json:"reason,omitempty"`
	ReviewDetails ReviewDetails         `json:"reviewDetails,omitempty"`
	Ext           ReviewRequestOpExt    `json:"ext,omitempty"`
}

// ReviewRequestResultCode is an XDR Enum defines as:
//
//   //: Result code of the ReviewRequest operation
//    enum ReviewRequestResultCode
//    {
//        //: Codes considered as "success" for an operation
//        //: Operation is applied successfuly
//        SUCCESS = 0,
//
//        //: Codes considered as "failure" for an operation
//        //: Reject reason must be empty on approve and not empty on reject/permanent
//        INVALID_REASON = -1,
//        //: Unknown action to perform on ReviewableRequest
//        INVALID_ACTION = -2,
//        //: Actual hash of the request and provided hash are mismatched
//        HASH_MISMATCHED = -3,
//        //: ReviewableRequest is not found
//        NOT_FOUND = -4,
//        //: Actual type of a reviewable request and provided type are mismatched
//        TYPE_MISMATCHED = -5,
//        //: Reject is not allowed. Only permanent reject should be used
//        REJECT_NOT_ALLOWED = -6,
//        //: External details must be a valid JSON
//        INVALID_EXTERNAL_DETAILS = -7,
//        //: Source of ReviewableRequest is blocked
//        REQUESTOR_IS_BLOCKED = -8,
//        //: Permanent reject is not allowed. Only reject should be used
//        PERMANENT_REJECT_NOT_ALLOWED = -9,
//        //: Trying to remove tasks which are not set
//        REMOVING_NOT_SET_TASKS = -100,// cannot remove tasks which are not set
//
//        //: Change role
//        //: Trying to remove zero tasks
//        NON_ZERO_TASKS_TO_REMOVE_NOT_ALLOWED = -600,
//        //: There is no account role with provided id
//        ACCOUNT_ROLE_TO_SET_DOES_NOT_EXIST = -610,
//
//        //KYC
//        //:Signer data is invalid - either weight is wrong or details are invalid
//        INVALID_SIGNER_DATA = -1600
//
//    };
//
type ReviewRequestResultCode int32

const (
	ReviewRequestResultCodeSuccess                        ReviewRequestResultCode = 0
	ReviewRequestResultCodeInvalidReason                  ReviewRequestResultCode = -1
	ReviewRequestResultCodeInvalidAction                  ReviewRequestResultCode = -2
	ReviewRequestResultCodeHashMismatched                 ReviewRequestResultCode = -3
	ReviewRequestResultCodeNotFound                       ReviewRequestResultCode = -4
	ReviewRequestResultCodeTypeMismatched                 ReviewRequestResultCode = -5
	ReviewRequestResultCodeRejectNotAllowed               ReviewRequestResultCode = -6
	ReviewRequestResultCodeInvalidExternalDetails         ReviewRequestResultCode = -7
	ReviewRequestResultCodeRequestorIsBlocked             ReviewRequestResultCode = -8
	ReviewRequestResultCodePermanentRejectNotAllowed      ReviewRequestResultCode = -9
	ReviewRequestResultCodeRemovingNotSetTasks            ReviewRequestResultCode = -100
	ReviewRequestResultCodeNonZeroTasksToRemoveNotAllowed ReviewRequestResultCode = -600
	ReviewRequestResultCodeAccountRoleToSetDoesNotExist   ReviewRequestResultCode = -610
	ReviewRequestResultCodeInvalidSignerData              ReviewRequestResultCode = -1600
)

var ReviewRequestResultCodeAll = []ReviewRequestResultCode{
	ReviewRequestResultCodeSuccess,
	ReviewRequestResultCodeInvalidReason,
	ReviewRequestResultCodeInvalidAction,
	ReviewRequestResultCodeHashMismatched,
	ReviewRequestResultCodeNotFound,
	ReviewRequestResultCodeTypeMismatched,
	ReviewRequestResultCodeRejectNotAllowed,
	ReviewRequestResultCodeInvalidExternalDetails,
	ReviewRequestResultCodeRequestorIsBlocked,
	ReviewRequestResultCodePermanentRejectNotAllowed,
	ReviewRequestResultCodeRemovingNotSetTasks,
	ReviewRequestResultCodeNonZeroTasksToRemoveNotAllowed,
	ReviewRequestResultCodeAccountRoleToSetDoesNotExist,
	ReviewRequestResultCodeInvalidSignerData,
}

var reviewRequestResultCodeMap = map[int32]string{
	0:     "ReviewRequestResultCodeSuccess",
	-1:    "ReviewRequestResultCodeInvalidReason",
	-2:    "ReviewRequestResultCodeInvalidAction",
	-3:    "ReviewRequestResultCodeHashMismatched",
	-4:    "ReviewRequestResultCodeNotFound",
	-5:    "ReviewRequestResultCodeTypeMismatched",
	-6:    "ReviewRequestResultCodeRejectNotAllowed",
	-7:    "ReviewRequestResultCodeInvalidExternalDetails",
	-8:    "ReviewRequestResultCodeRequestorIsBlocked",
	-9:    "ReviewRequestResultCodePermanentRejectNotAllowed",
	-100:  "ReviewRequestResultCodeRemovingNotSetTasks",
	-600:  "ReviewRequestResultCodeNonZeroTasksToRemoveNotAllowed",
	-610:  "ReviewRequestResultCodeAccountRoleToSetDoesNotExist",
	-1600: "ReviewRequestResultCodeInvalidSignerData",
}

var reviewRequestResultCodeShortMap = map[int32]string{
	0:     "success",
	-1:    "invalid_reason",
	-2:    "invalid_action",
	-3:    "hash_mismatched",
	-4:    "not_found",
	-5:    "type_mismatched",
	-6:    "reject_not_allowed",
	-7:    "invalid_external_details",
	-8:    "requestor_is_blocked",
	-9:    "permanent_reject_not_allowed",
	-100:  "removing_not_set_tasks",
	-600:  "non_zero_tasks_to_remove_not_allowed",
	-610:  "account_role_to_set_does_not_exist",
	-1600: "invalid_signer_data",
}

var reviewRequestResultCodeRevMap = map[string]int32{
	"ReviewRequestResultCodeSuccess":                        0,
	"ReviewRequestResultCodeInvalidReason":                  -1,
	"ReviewRequestResultCodeInvalidAction":                  -2,
	"ReviewRequestResultCodeHashMismatched":                 -3,
	"ReviewRequestResultCodeNotFound":                       -4,
	"ReviewRequestResultCodeTypeMismatched":                 -5,
	"ReviewRequestResultCodeRejectNotAllowed":               -6,
	"ReviewRequestResultCodeInvalidExternalDetails":         -7,
	"ReviewRequestResultCodeRequestorIsBlocked":             -8,
	"ReviewRequestResultCodePermanentRejectNotAllowed":      -9,
	"ReviewRequestResultCodeRemovingNotSetTasks":            -100,
	"ReviewRequestResultCodeNonZeroTasksToRemoveNotAllowed": -600,
	"ReviewRequestResultCodeAccountRoleToSetDoesNotExist":   -610,
	"ReviewRequestResultCodeInvalidSignerData":              -1600,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewRequestResultCode
func (e ReviewRequestResultCode) ValidEnum(v int32) bool {
	_, ok := reviewRequestResultCodeMap[v]
	return ok
}
func (e ReviewRequestResultCode) isFlag() bool {
	for i := len(ReviewRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := ReviewRequestResultCode(2) << uint64(len(ReviewRequestResultCodeAll)-1) >> uint64(len(ReviewRequestResultCodeAll)-i)
		if expected != ReviewRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewRequestResultCode) String() string {
	name, _ := reviewRequestResultCodeMap[int32(e)]
	return name
}

func (e ReviewRequestResultCode) ShortString() string {
	name, _ := reviewRequestResultCodeShortMap[int32(e)]
	return name
}

func (e ReviewRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ReviewRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ReviewRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewRequestResultCode(t.Value)
	return nil
}

// ReviewRequestResult is an XDR Union defines as:
//
//   //: Result of applying the review request with result code
//    union ReviewRequestResult switch (ReviewRequestResultCode code)
//    {
//    case SUCCESS:
//        ExtendedResult success;
//    default:
//        void;
//    };
//
type ReviewRequestResult struct {
	Code    ReviewRequestResultCode `json:"code,omitempty"`
	Success *ExtendedResult         `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestResult
func (u ReviewRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewRequestResultCode(sw) {
	case ReviewRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewReviewRequestResult creates a new  ReviewRequestResult.
func NewReviewRequestResult(code ReviewRequestResultCode, value interface{}) (result ReviewRequestResult, err error) {
	result.Code = code
	switch ReviewRequestResultCode(code) {
	case ReviewRequestResultCodeSuccess:
		tv, ok := value.(ExtendedResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ExtendedResult")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ReviewRequestResult) MustSuccess() ExtendedResult {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestResult) GetSuccess() (result ExtendedResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ErrorCode is an XDR Enum defines as:
//
//   enum ErrorCode
//    {
//        MISC = 0, // Unspecific error
//        DATA = 1, // Malformed data
//        CONF = 2, // Misconfiguration error
//        AUTH = 3, // Authentication failure
//        LOAD = 4  // System overloaded
//    };
//
type ErrorCode int32

const (
	ErrorCodeMisc ErrorCode = 0
	ErrorCodeData ErrorCode = 1
	ErrorCodeConf ErrorCode = 2
	ErrorCodeAuth ErrorCode = 3
	ErrorCodeLoad ErrorCode = 4
)

var ErrorCodeAll = []ErrorCode{
	ErrorCodeMisc,
	ErrorCodeData,
	ErrorCodeConf,
	ErrorCodeAuth,
	ErrorCodeLoad,
}

var errorCodeMap = map[int32]string{
	0: "ErrorCodeMisc",
	1: "ErrorCodeData",
	2: "ErrorCodeConf",
	3: "ErrorCodeAuth",
	4: "ErrorCodeLoad",
}

var errorCodeShortMap = map[int32]string{
	0: "misc",
	1: "data",
	2: "conf",
	3: "auth",
	4: "load",
}

var errorCodeRevMap = map[string]int32{
	"ErrorCodeMisc": 0,
	"ErrorCodeData": 1,
	"ErrorCodeConf": 2,
	"ErrorCodeAuth": 3,
	"ErrorCodeLoad": 4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ErrorCode
func (e ErrorCode) ValidEnum(v int32) bool {
	_, ok := errorCodeMap[v]
	return ok
}
func (e ErrorCode) isFlag() bool {
	for i := len(ErrorCodeAll) - 1; i >= 0; i-- {
		expected := ErrorCode(2) << uint64(len(ErrorCodeAll)-1) >> uint64(len(ErrorCodeAll)-i)
		if expected != ErrorCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ErrorCode) String() string {
	name, _ := errorCodeMap[int32(e)]
	return name
}

func (e ErrorCode) ShortString() string {
	name, _ := errorCodeShortMap[int32(e)]
	return name
}

func (e ErrorCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ErrorCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ErrorCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ErrorCode(t.Value)
	return nil
}

// Error is an XDR Struct defines as:
//
//   struct Error
//    {
//        ErrorCode code;
//        string msg<100>;
//    };
//
type Error struct {
	Code ErrorCode `json:"code,omitempty"`
	Msg  string    `json:"msg,omitempty" xdrmaxsize:"100"`
}

// AuthCert is an XDR Struct defines as:
//
//   struct AuthCert
//    {
//        Curve25519Public pubkey;
//        uint64 expiration;
//        Signature sig;
//    };
//
type AuthCert struct {
	Pubkey     Curve25519Public `json:"pubkey,omitempty"`
	Expiration Uint64           `json:"expiration,omitempty"`
	Sig        Signature        `json:"sig,omitempty"`
}

// Hello is an XDR Struct defines as:
//
//   struct Hello
//    {
//        uint32 ledgerVersion;
//        uint32 overlayVersion;
//        uint32 overlayMinVersion;
//        Hash networkID;
//        string versionStr<100>;
//        int listeningPort;
//        NodeID peerID;
//        AuthCert cert;
//        uint256 nonce;
//    };
//
type Hello struct {
	LedgerVersion     Uint32   `json:"ledgerVersion,omitempty"`
	OverlayVersion    Uint32   `json:"overlayVersion,omitempty"`
	OverlayMinVersion Uint32   `json:"overlayMinVersion,omitempty"`
	NetworkId         Hash     `json:"networkID,omitempty"`
	VersionStr        string   `json:"versionStr,omitempty" xdrmaxsize:"100"`
	ListeningPort     int32    `json:"listeningPort,omitempty"`
	PeerId            NodeId   `json:"peerID,omitempty"`
	Cert              AuthCert `json:"cert,omitempty"`
	Nonce             Uint256  `json:"nonce,omitempty"`
}

// Auth is an XDR Struct defines as:
//
//   struct Auth
//    {
//        // Empty message, just to confirm
//        // establishment of MAC keys.
//        int unused;
//    };
//
type Auth struct {
	Unused int32 `json:"unused,omitempty"`
}

// IpAddrType is an XDR Enum defines as:
//
//   enum IPAddrType
//    {
//        IPv4 = 0,
//        IPv6 = 1
//    };
//
type IpAddrType int32

const (
	IpAddrTypeIPv4 IpAddrType = 0
	IpAddrTypeIPv6 IpAddrType = 1
)

var IpAddrTypeAll = []IpAddrType{
	IpAddrTypeIPv4,
	IpAddrTypeIPv6,
}

var ipAddrTypeMap = map[int32]string{
	0: "IpAddrTypeIPv4",
	1: "IpAddrTypeIPv6",
}

var ipAddrTypeShortMap = map[int32]string{
	0: "i_pv4",
	1: "i_pv6",
}

var ipAddrTypeRevMap = map[string]int32{
	"IpAddrTypeIPv4": 0,
	"IpAddrTypeIPv6": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for IpAddrType
func (e IpAddrType) ValidEnum(v int32) bool {
	_, ok := ipAddrTypeMap[v]
	return ok
}
func (e IpAddrType) isFlag() bool {
	for i := len(IpAddrTypeAll) - 1; i >= 0; i-- {
		expected := IpAddrType(2) << uint64(len(IpAddrTypeAll)-1) >> uint64(len(IpAddrTypeAll)-i)
		if expected != IpAddrTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e IpAddrType) String() string {
	name, _ := ipAddrTypeMap[int32(e)]
	return name
}

func (e IpAddrType) ShortString() string {
	name, _ := ipAddrTypeShortMap[int32(e)]
	return name
}

func (e IpAddrType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range IpAddrTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *IpAddrType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = IpAddrType(t.Value)
	return nil
}

// PeerAddressIp is an XDR NestedUnion defines as:
//
//   union switch (IPAddrType type)
//        {
//        case IPv4:
//            opaque ipv4[4];
//        case IPv6:
//            opaque ipv6[16];
//        }
//
type PeerAddressIp struct {
	Type IpAddrType `json:"type,omitempty"`
	Ipv4 *[4]byte   `json:"ipv4,omitempty"`
	Ipv6 *[16]byte  `json:"ipv6,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PeerAddressIp) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PeerAddressIp
func (u PeerAddressIp) ArmForSwitch(sw int32) (string, bool) {
	switch IpAddrType(sw) {
	case IpAddrTypeIPv4:
		return "Ipv4", true
	case IpAddrTypeIPv6:
		return "Ipv6", true
	}
	return "-", false
}

// NewPeerAddressIp creates a new  PeerAddressIp.
func NewPeerAddressIp(aType IpAddrType, value interface{}) (result PeerAddressIp, err error) {
	result.Type = aType
	switch IpAddrType(aType) {
	case IpAddrTypeIPv4:
		tv, ok := value.([4]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [4]byte")
			return
		}
		result.Ipv4 = &tv
	case IpAddrTypeIPv6:
		tv, ok := value.([16]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [16]byte")
			return
		}
		result.Ipv6 = &tv
	}
	return
}

// MustIpv4 retrieves the Ipv4 value from the union,
// panicing if the value is not set.
func (u PeerAddressIp) MustIpv4() [4]byte {
	val, ok := u.GetIpv4()

	if !ok {
		panic("arm Ipv4 is not set")
	}

	return val
}

// GetIpv4 retrieves the Ipv4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PeerAddressIp) GetIpv4() (result [4]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ipv4" {
		result = *u.Ipv4
		ok = true
	}

	return
}

// MustIpv6 retrieves the Ipv6 value from the union,
// panicing if the value is not set.
func (u PeerAddressIp) MustIpv6() [16]byte {
	val, ok := u.GetIpv6()

	if !ok {
		panic("arm Ipv6 is not set")
	}

	return val
}

// GetIpv6 retrieves the Ipv6 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PeerAddressIp) GetIpv6() (result [16]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ipv6" {
		result = *u.Ipv6
		ok = true
	}

	return
}

// PeerAddress is an XDR Struct defines as:
//
//   struct PeerAddress
//    {
//        union switch (IPAddrType type)
//        {
//        case IPv4:
//            opaque ipv4[4];
//        case IPv6:
//            opaque ipv6[16];
//        }
//        ip;
//        uint32 port;
//        uint32 numFailures;
//    };
//
type PeerAddress struct {
	Ip          PeerAddressIp `json:"ip,omitempty"`
	Port        Uint32        `json:"port,omitempty"`
	NumFailures Uint32        `json:"numFailures,omitempty"`
}

// MessageType is an XDR Enum defines as:
//
//   enum MessageType
//    {
//        ERROR_MSG = 0,
//        AUTH = 2,
//        DONT_HAVE = 3,
//
//        GET_PEERS = 4, // gets a list of peers this guy knows about
//        PEERS = 5,
//
//        GET_TX_SET = 6, // gets a particular txset by hash
//        TX_SET = 7,
//
//        TRANSACTION = 8, // pass on a tx you have heard about
//
//        // SCP
//        GET_SCP_QUORUMSET = 9,
//        SCP_QUORUMSET = 10,
//        SCP_MESSAGE = 11,
//        GET_SCP_STATE = 12,
//
//        // new messages
//        HELLO = 13
//    };
//
type MessageType int32

const (
	MessageTypeErrorMsg        MessageType = 0
	MessageTypeAuth            MessageType = 2
	MessageTypeDontHave        MessageType = 3
	MessageTypeGetPeers        MessageType = 4
	MessageTypePeers           MessageType = 5
	MessageTypeGetTxSet        MessageType = 6
	MessageTypeTxSet           MessageType = 7
	MessageTypeTransaction     MessageType = 8
	MessageTypeGetScpQuorumset MessageType = 9
	MessageTypeScpQuorumset    MessageType = 10
	MessageTypeScpMessage      MessageType = 11
	MessageTypeGetScpState     MessageType = 12
	MessageTypeHello           MessageType = 13
)

var MessageTypeAll = []MessageType{
	MessageTypeErrorMsg,
	MessageTypeAuth,
	MessageTypeDontHave,
	MessageTypeGetPeers,
	MessageTypePeers,
	MessageTypeGetTxSet,
	MessageTypeTxSet,
	MessageTypeTransaction,
	MessageTypeGetScpQuorumset,
	MessageTypeScpQuorumset,
	MessageTypeScpMessage,
	MessageTypeGetScpState,
	MessageTypeHello,
}

var messageTypeMap = map[int32]string{
	0:  "MessageTypeErrorMsg",
	2:  "MessageTypeAuth",
	3:  "MessageTypeDontHave",
	4:  "MessageTypeGetPeers",
	5:  "MessageTypePeers",
	6:  "MessageTypeGetTxSet",
	7:  "MessageTypeTxSet",
	8:  "MessageTypeTransaction",
	9:  "MessageTypeGetScpQuorumset",
	10: "MessageTypeScpQuorumset",
	11: "MessageTypeScpMessage",
	12: "MessageTypeGetScpState",
	13: "MessageTypeHello",
}

var messageTypeShortMap = map[int32]string{
	0:  "error_msg",
	2:  "auth",
	3:  "dont_have",
	4:  "get_peers",
	5:  "peers",
	6:  "get_tx_set",
	7:  "tx_set",
	8:  "transaction",
	9:  "get_scp_quorumset",
	10: "scp_quorumset",
	11: "scp_message",
	12: "get_scp_state",
	13: "hello",
}

var messageTypeRevMap = map[string]int32{
	"MessageTypeErrorMsg":        0,
	"MessageTypeAuth":            2,
	"MessageTypeDontHave":        3,
	"MessageTypeGetPeers":        4,
	"MessageTypePeers":           5,
	"MessageTypeGetTxSet":        6,
	"MessageTypeTxSet":           7,
	"MessageTypeTransaction":     8,
	"MessageTypeGetScpQuorumset": 9,
	"MessageTypeScpQuorumset":    10,
	"MessageTypeScpMessage":      11,
	"MessageTypeGetScpState":     12,
	"MessageTypeHello":           13,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for MessageType
func (e MessageType) ValidEnum(v int32) bool {
	_, ok := messageTypeMap[v]
	return ok
}
func (e MessageType) isFlag() bool {
	for i := len(MessageTypeAll) - 1; i >= 0; i-- {
		expected := MessageType(2) << uint64(len(MessageTypeAll)-1) >> uint64(len(MessageTypeAll)-i)
		if expected != MessageTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e MessageType) String() string {
	name, _ := messageTypeMap[int32(e)]
	return name
}

func (e MessageType) ShortString() string {
	name, _ := messageTypeShortMap[int32(e)]
	return name
}

func (e MessageType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range MessageTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *MessageType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = MessageType(t.Value)
	return nil
}

// DontHave is an XDR Struct defines as:
//
//   struct DontHave
//    {
//        MessageType type;
//        uint256 reqHash;
//    };
//
type DontHave struct {
	Type    MessageType `json:"type,omitempty"`
	ReqHash Uint256     `json:"reqHash,omitempty"`
}

// StellarMessage is an XDR Union defines as:
//
//   union StellarMessage switch (MessageType type)
//    {
//    case ERROR_MSG:
//        Error error;
//    case HELLO:
//        Hello hello;
//    case AUTH:
//        Auth auth;
//    case DONT_HAVE:
//        DontHave dontHave;
//    case GET_PEERS:
//        void;
//    case PEERS:
//        PeerAddress peers<>;
//
//    case GET_TX_SET:
//        uint256 txSetHash;
//    case TX_SET:
//        TransactionSet txSet;
//
//    case TRANSACTION:
//        TransactionEnvelope transaction;
//
//    // SCP
//    case GET_SCP_QUORUMSET:
//        uint256 qSetHash;
//    case SCP_QUORUMSET:
//        SCPQuorumSet qSet;
//    case SCP_MESSAGE:
//        SCPEnvelope envelope;
//    case GET_SCP_STATE:
//        uint32 getSCPLedgerSeq; // ledger seq requested ; if 0, requests the latest
//    };
//
type StellarMessage struct {
	Type            MessageType          `json:"type,omitempty"`
	Error           *Error               `json:"error,omitempty"`
	Hello           *Hello               `json:"hello,omitempty"`
	Auth            *Auth                `json:"auth,omitempty"`
	DontHave        *DontHave            `json:"dontHave,omitempty"`
	Peers           *[]PeerAddress       `json:"peers,omitempty"`
	TxSetHash       *Uint256             `json:"txSetHash,omitempty"`
	TxSet           *TransactionSet      `json:"txSet,omitempty"`
	Transaction     *TransactionEnvelope `json:"transaction,omitempty"`
	QSetHash        *Uint256             `json:"qSetHash,omitempty"`
	QSet            *ScpQuorumSet        `json:"qSet,omitempty"`
	Envelope        *ScpEnvelope         `json:"envelope,omitempty"`
	GetScpLedgerSeq *Uint32              `json:"getSCPLedgerSeq,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StellarMessage) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StellarMessage
func (u StellarMessage) ArmForSwitch(sw int32) (string, bool) {
	switch MessageType(sw) {
	case MessageTypeErrorMsg:
		return "Error", true
	case MessageTypeHello:
		return "Hello", true
	case MessageTypeAuth:
		return "Auth", true
	case MessageTypeDontHave:
		return "DontHave", true
	case MessageTypeGetPeers:
		return "", true
	case MessageTypePeers:
		return "Peers", true
	case MessageTypeGetTxSet:
		return "TxSetHash", true
	case MessageTypeTxSet:
		return "TxSet", true
	case MessageTypeTransaction:
		return "Transaction", true
	case MessageTypeGetScpQuorumset:
		return "QSetHash", true
	case MessageTypeScpQuorumset:
		return "QSet", true
	case MessageTypeScpMessage:
		return "Envelope", true
	case MessageTypeGetScpState:
		return "GetScpLedgerSeq", true
	}
	return "-", false
}

// NewStellarMessage creates a new  StellarMessage.
func NewStellarMessage(aType MessageType, value interface{}) (result StellarMessage, err error) {
	result.Type = aType
	switch MessageType(aType) {
	case MessageTypeErrorMsg:
		tv, ok := value.(Error)
		if !ok {
			err = fmt.Errorf("invalid value, must be Error")
			return
		}
		result.Error = &tv
	case MessageTypeHello:
		tv, ok := value.(Hello)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hello")
			return
		}
		result.Hello = &tv
	case MessageTypeAuth:
		tv, ok := value.(Auth)
		if !ok {
			err = fmt.Errorf("invalid value, must be Auth")
			return
		}
		result.Auth = &tv
	case MessageTypeDontHave:
		tv, ok := value.(DontHave)
		if !ok {
			err = fmt.Errorf("invalid value, must be DontHave")
			return
		}
		result.DontHave = &tv
	case MessageTypeGetPeers:
		// void
	case MessageTypePeers:
		tv, ok := value.([]PeerAddress)
		if !ok {
			err = fmt.Errorf("invalid value, must be []PeerAddress")
			return
		}
		result.Peers = &tv
	case MessageTypeGetTxSet:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.TxSetHash = &tv
	case MessageTypeTxSet:
		tv, ok := value.(TransactionSet)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionSet")
			return
		}
		result.TxSet = &tv
	case MessageTypeTransaction:
		tv, ok := value.(TransactionEnvelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionEnvelope")
			return
		}
		result.Transaction = &tv
	case MessageTypeGetScpQuorumset:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.QSetHash = &tv
	case MessageTypeScpQuorumset:
		tv, ok := value.(ScpQuorumSet)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpQuorumSet")
			return
		}
		result.QSet = &tv
	case MessageTypeScpMessage:
		tv, ok := value.(ScpEnvelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpEnvelope")
			return
		}
		result.Envelope = &tv
	case MessageTypeGetScpState:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.GetScpLedgerSeq = &tv
	}
	return
}

// MustError retrieves the Error value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustError() Error {
	val, ok := u.GetError()

	if !ok {
		panic("arm Error is not set")
	}

	return val
}

// GetError retrieves the Error value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetError() (result Error, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Error" {
		result = *u.Error
		ok = true
	}

	return
}

// MustHello retrieves the Hello value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustHello() Hello {
	val, ok := u.GetHello()

	if !ok {
		panic("arm Hello is not set")
	}

	return val
}

// GetHello retrieves the Hello value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetHello() (result Hello, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hello" {
		result = *u.Hello
		ok = true
	}

	return
}

// MustAuth retrieves the Auth value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustAuth() Auth {
	val, ok := u.GetAuth()

	if !ok {
		panic("arm Auth is not set")
	}

	return val
}

// GetAuth retrieves the Auth value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetAuth() (result Auth, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Auth" {
		result = *u.Auth
		ok = true
	}

	return
}

// MustDontHave retrieves the DontHave value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustDontHave() DontHave {
	val, ok := u.GetDontHave()

	if !ok {
		panic("arm DontHave is not set")
	}

	return val
}

// GetDontHave retrieves the DontHave value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetDontHave() (result DontHave, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DontHave" {
		result = *u.DontHave
		ok = true
	}

	return
}

// MustPeers retrieves the Peers value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustPeers() []PeerAddress {
	val, ok := u.GetPeers()

	if !ok {
		panic("arm Peers is not set")
	}

	return val
}

// GetPeers retrieves the Peers value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetPeers() (result []PeerAddress, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Peers" {
		result = *u.Peers
		ok = true
	}

	return
}

// MustTxSetHash retrieves the TxSetHash value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTxSetHash() Uint256 {
	val, ok := u.GetTxSetHash()

	if !ok {
		panic("arm TxSetHash is not set")
	}

	return val
}

// GetTxSetHash retrieves the TxSetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTxSetHash() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TxSetHash" {
		result = *u.TxSetHash
		ok = true
	}

	return
}

// MustTxSet retrieves the TxSet value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTxSet() TransactionSet {
	val, ok := u.GetTxSet()

	if !ok {
		panic("arm TxSet is not set")
	}

	return val
}

// GetTxSet retrieves the TxSet value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTxSet() (result TransactionSet, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TxSet" {
		result = *u.TxSet
		ok = true
	}

	return
}

// MustTransaction retrieves the Transaction value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTransaction() TransactionEnvelope {
	val, ok := u.GetTransaction()

	if !ok {
		panic("arm Transaction is not set")
	}

	return val
}

// GetTransaction retrieves the Transaction value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTransaction() (result TransactionEnvelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Transaction" {
		result = *u.Transaction
		ok = true
	}

	return
}

// MustQSetHash retrieves the QSetHash value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustQSetHash() Uint256 {
	val, ok := u.GetQSetHash()

	if !ok {
		panic("arm QSetHash is not set")
	}

	return val
}

// GetQSetHash retrieves the QSetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetQSetHash() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "QSetHash" {
		result = *u.QSetHash
		ok = true
	}

	return
}

// MustQSet retrieves the QSet value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustQSet() ScpQuorumSet {
	val, ok := u.GetQSet()

	if !ok {
		panic("arm QSet is not set")
	}

	return val
}

// GetQSet retrieves the QSet value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetQSet() (result ScpQuorumSet, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "QSet" {
		result = *u.QSet
		ok = true
	}

	return
}

// MustEnvelope retrieves the Envelope value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustEnvelope() ScpEnvelope {
	val, ok := u.GetEnvelope()

	if !ok {
		panic("arm Envelope is not set")
	}

	return val
}

// GetEnvelope retrieves the Envelope value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetEnvelope() (result ScpEnvelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Envelope" {
		result = *u.Envelope
		ok = true
	}

	return
}

// MustGetScpLedgerSeq retrieves the GetScpLedgerSeq value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustGetScpLedgerSeq() Uint32 {
	val, ok := u.GetGetScpLedgerSeq()

	if !ok {
		panic("arm GetScpLedgerSeq is not set")
	}

	return val
}

// GetGetScpLedgerSeq retrieves the GetScpLedgerSeq value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetGetScpLedgerSeq() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "GetScpLedgerSeq" {
		result = *u.GetScpLedgerSeq
		ok = true
	}

	return
}

// AuthenticatedMessageV0 is an XDR NestedStruct defines as:
//
//   struct
//    {
//       uint64 sequence;
//       StellarMessage message;
//       HmacSha256Mac mac;
//        }
//
type AuthenticatedMessageV0 struct {
	Sequence Uint64         `json:"sequence,omitempty"`
	Message  StellarMessage `json:"message,omitempty"`
	Mac      HmacSha256Mac  `json:"mac,omitempty"`
}

// AuthenticatedMessage is an XDR Union defines as:
//
//   union AuthenticatedMessage switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        struct
//    {
//       uint64 sequence;
//       StellarMessage message;
//       HmacSha256Mac mac;
//        } v0;
//    };
//
type AuthenticatedMessage struct {
	V  LedgerVersion           `json:"v,omitempty"`
	V0 *AuthenticatedMessageV0 `json:"v0,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AuthenticatedMessage) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AuthenticatedMessage
func (u AuthenticatedMessage) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "V0", true
	}
	return "-", false
}

// NewAuthenticatedMessage creates a new  AuthenticatedMessage.
func NewAuthenticatedMessage(v LedgerVersion, value interface{}) (result AuthenticatedMessage, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		tv, ok := value.(AuthenticatedMessageV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be AuthenticatedMessageV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u AuthenticatedMessage) MustV0() AuthenticatedMessageV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AuthenticatedMessage) GetV0() (result AuthenticatedMessageV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// AccountRuleResourceReviewableRequest is an XDR NestedStruct defines as:
//
//   struct {
//            //: type of request
//            ReviewableRequestType requestType;
//            //: reserved for future extension
//            EmptyExt ext;
//        }
//
type AccountRuleResourceReviewableRequest struct {
	RequestType ReviewableRequestType `json:"requestType,omitempty"`
	Ext         EmptyExt              `json:"ext,omitempty"`
}

// AccountRuleResourceKeyValue is an XDR NestedStruct defines as:
//
//   struct
//        {
//            //: prefix of key
//            longstring keyPrefix;
//
//            //: reserved for future extension
//            EmptyExt ext;
//        }
//
type AccountRuleResourceKeyValue struct {
	KeyPrefix Longstring `json:"keyPrefix,omitempty"`
	Ext       EmptyExt   `json:"ext,omitempty"`
}

// AccountRuleResourceInitiateKycRecovery is an XDR NestedStruct defines as:
//
//   struct
//        {
//            //: Role id
//            uint64 roleID;
//
//            //: reserved for future extension
//            EmptyExt ext;
//        }
//
type AccountRuleResourceInitiateKycRecovery struct {
	RoleId Uint64   `json:"roleID,omitempty"`
	Ext    EmptyExt `json:"ext,omitempty"`
}

// AccountRuleResource is an XDR Union defines as:
//
//   //: Describes properties of some entries that can be used to restrict the usage of entries
//    union AccountRuleResource switch (LedgerEntryType type)
//    {
//    case REVIEWABLE_REQUEST:
//        struct {
//            //: type of request
//            ReviewableRequestType requestType;
//            //: reserved for future extension
//            EmptyExt ext;
//        } reviewableRequest;
//    case ANY:
//        void;
//    case KEY_VALUE:
//        struct
//        {
//            //: prefix of key
//            longstring keyPrefix;
//
//            //: reserved for future extension
//            EmptyExt ext;
//        } keyValue;
//    case INITIATE_KYC_RECOVERY:
//        struct
//        {
//            //: Role id
//            uint64 roleID;
//
//            //: reserved for future extension
//            EmptyExt ext;
//        } initiateKYCRecovery;
//    default:
//        //: reserved for future extension
//        EmptyExt ext;
//    };
//
type AccountRuleResource struct {
	Type                LedgerEntryType                         `json:"type,omitempty"`
	ReviewableRequest   *AccountRuleResourceReviewableRequest   `json:"reviewableRequest,omitempty"`
	KeyValue            *AccountRuleResourceKeyValue            `json:"keyValue,omitempty"`
	InitiateKycRecovery *AccountRuleResourceInitiateKycRecovery `json:"initiateKYCRecovery,omitempty"`
	Ext                 *EmptyExt                               `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountRuleResource) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountRuleResource
func (u AccountRuleResource) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeAny:
		return "", true
	case LedgerEntryTypeKeyValue:
		return "KeyValue", true
	case LedgerEntryTypeInitiateKycRecovery:
		return "InitiateKycRecovery", true
	default:
		return "Ext", true
	}
}

// NewAccountRuleResource creates a new  AccountRuleResource.
func NewAccountRuleResource(aType LedgerEntryType, value interface{}) (result AccountRuleResource, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeReviewableRequest:
		tv, ok := value.(AccountRuleResourceReviewableRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountRuleResourceReviewableRequest")
			return
		}
		result.ReviewableRequest = &tv
	case LedgerEntryTypeAny:
		// void
	case LedgerEntryTypeKeyValue:
		tv, ok := value.(AccountRuleResourceKeyValue)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountRuleResourceKeyValue")
			return
		}
		result.KeyValue = &tv
	case LedgerEntryTypeInitiateKycRecovery:
		tv, ok := value.(AccountRuleResourceInitiateKycRecovery)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountRuleResourceInitiateKycRecovery")
			return
		}
		result.InitiateKycRecovery = &tv
	default:
		tv, ok := value.(EmptyExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be EmptyExt")
			return
		}
		result.Ext = &tv
	}
	return
}

// MustReviewableRequest retrieves the ReviewableRequest value from the union,
// panicing if the value is not set.
func (u AccountRuleResource) MustReviewableRequest() AccountRuleResourceReviewableRequest {
	val, ok := u.GetReviewableRequest()

	if !ok {
		panic("arm ReviewableRequest is not set")
	}

	return val
}

// GetReviewableRequest retrieves the ReviewableRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountRuleResource) GetReviewableRequest() (result AccountRuleResourceReviewableRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewableRequest" {
		result = *u.ReviewableRequest
		ok = true
	}

	return
}

// MustKeyValue retrieves the KeyValue value from the union,
// panicing if the value is not set.
func (u AccountRuleResource) MustKeyValue() AccountRuleResourceKeyValue {
	val, ok := u.GetKeyValue()

	if !ok {
		panic("arm KeyValue is not set")
	}

	return val
}

// GetKeyValue retrieves the KeyValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountRuleResource) GetKeyValue() (result AccountRuleResourceKeyValue, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KeyValue" {
		result = *u.KeyValue
		ok = true
	}

	return
}

// MustInitiateKycRecovery retrieves the InitiateKycRecovery value from the union,
// panicing if the value is not set.
func (u AccountRuleResource) MustInitiateKycRecovery() AccountRuleResourceInitiateKycRecovery {
	val, ok := u.GetInitiateKycRecovery()

	if !ok {
		panic("arm InitiateKycRecovery is not set")
	}

	return val
}

// GetInitiateKycRecovery retrieves the InitiateKycRecovery value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountRuleResource) GetInitiateKycRecovery() (result AccountRuleResourceInitiateKycRecovery, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "InitiateKycRecovery" {
		result = *u.InitiateKycRecovery
		ok = true
	}

	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u AccountRuleResource) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountRuleResource) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// AccountRuleAction is an XDR Enum defines as:
//
//   //: Actions that can be applied to account rule resource
//    enum AccountRuleAction
//    {
//        ANY = 1,
//        CREATE = 2,
//        CREATE_FOR_OTHER = 3,
//        CREATE_WITH_TASKS = 4,
//        MANAGE = 5,
//        SEND = 6,
//        WITHDRAW = 7,
//        RECEIVE_ISSUANCE = 8,
//        RECEIVE_PAYMENT = 9,
//        RECEIVE_ATOMIC_SWAP = 10,
//        PARTICIPATE = 11,
//        BIND = 12,
//        UPDATE_MAX_ISSUANCE = 13,
//        CHECK = 14,
//        CANCEL = 15,
//        CLOSE = 16,
//        REMOVE = 17,
//        UPDATE_END_TIME = 18,
//        CREATE_FOR_OTHER_WITH_TASKS = 19
//    };
//
type AccountRuleAction int32

const (
	AccountRuleActionAny                     AccountRuleAction = 1
	AccountRuleActionCreate                  AccountRuleAction = 2
	AccountRuleActionCreateForOther          AccountRuleAction = 3
	AccountRuleActionCreateWithTasks         AccountRuleAction = 4
	AccountRuleActionManage                  AccountRuleAction = 5
	AccountRuleActionSend                    AccountRuleAction = 6
	AccountRuleActionWithdraw                AccountRuleAction = 7
	AccountRuleActionReceiveIssuance         AccountRuleAction = 8
	AccountRuleActionReceivePayment          AccountRuleAction = 9
	AccountRuleActionReceiveAtomicSwap       AccountRuleAction = 10
	AccountRuleActionParticipate             AccountRuleAction = 11
	AccountRuleActionBind                    AccountRuleAction = 12
	AccountRuleActionUpdateMaxIssuance       AccountRuleAction = 13
	AccountRuleActionCheck                   AccountRuleAction = 14
	AccountRuleActionCancel                  AccountRuleAction = 15
	AccountRuleActionClose                   AccountRuleAction = 16
	AccountRuleActionRemove                  AccountRuleAction = 17
	AccountRuleActionUpdateEndTime           AccountRuleAction = 18
	AccountRuleActionCreateForOtherWithTasks AccountRuleAction = 19
)

var AccountRuleActionAll = []AccountRuleAction{
	AccountRuleActionAny,
	AccountRuleActionCreate,
	AccountRuleActionCreateForOther,
	AccountRuleActionCreateWithTasks,
	AccountRuleActionManage,
	AccountRuleActionSend,
	AccountRuleActionWithdraw,
	AccountRuleActionReceiveIssuance,
	AccountRuleActionReceivePayment,
	AccountRuleActionReceiveAtomicSwap,
	AccountRuleActionParticipate,
	AccountRuleActionBind,
	AccountRuleActionUpdateMaxIssuance,
	AccountRuleActionCheck,
	AccountRuleActionCancel,
	AccountRuleActionClose,
	AccountRuleActionRemove,
	AccountRuleActionUpdateEndTime,
	AccountRuleActionCreateForOtherWithTasks,
}

var accountRuleActionMap = map[int32]string{
	1:  "AccountRuleActionAny",
	2:  "AccountRuleActionCreate",
	3:  "AccountRuleActionCreateForOther",
	4:  "AccountRuleActionCreateWithTasks",
	5:  "AccountRuleActionManage",
	6:  "AccountRuleActionSend",
	7:  "AccountRuleActionWithdraw",
	8:  "AccountRuleActionReceiveIssuance",
	9:  "AccountRuleActionReceivePayment",
	10: "AccountRuleActionReceiveAtomicSwap",
	11: "AccountRuleActionParticipate",
	12: "AccountRuleActionBind",
	13: "AccountRuleActionUpdateMaxIssuance",
	14: "AccountRuleActionCheck",
	15: "AccountRuleActionCancel",
	16: "AccountRuleActionClose",
	17: "AccountRuleActionRemove",
	18: "AccountRuleActionUpdateEndTime",
	19: "AccountRuleActionCreateForOtherWithTasks",
}

var accountRuleActionShortMap = map[int32]string{
	1:  "any",
	2:  "create",
	3:  "create_for_other",
	4:  "create_with_tasks",
	5:  "manage",
	6:  "send",
	7:  "withdraw",
	8:  "receive_issuance",
	9:  "receive_payment",
	10: "receive_atomic_swap",
	11: "participate",
	12: "bind",
	13: "update_max_issuance",
	14: "check",
	15: "cancel",
	16: "close",
	17: "remove",
	18: "update_end_time",
	19: "create_for_other_with_tasks",
}

var accountRuleActionRevMap = map[string]int32{
	"AccountRuleActionAny":                     1,
	"AccountRuleActionCreate":                  2,
	"AccountRuleActionCreateForOther":          3,
	"AccountRuleActionCreateWithTasks":         4,
	"AccountRuleActionManage":                  5,
	"AccountRuleActionSend":                    6,
	"AccountRuleActionWithdraw":                7,
	"AccountRuleActionReceiveIssuance":         8,
	"AccountRuleActionReceivePayment":          9,
	"AccountRuleActionReceiveAtomicSwap":       10,
	"AccountRuleActionParticipate":             11,
	"AccountRuleActionBind":                    12,
	"AccountRuleActionUpdateMaxIssuance":       13,
	"AccountRuleActionCheck":                   14,
	"AccountRuleActionCancel":                  15,
	"AccountRuleActionClose":                   16,
	"AccountRuleActionRemove":                  17,
	"AccountRuleActionUpdateEndTime":           18,
	"AccountRuleActionCreateForOtherWithTasks": 19,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AccountRuleAction
func (e AccountRuleAction) ValidEnum(v int32) bool {
	_, ok := accountRuleActionMap[v]
	return ok
}
func (e AccountRuleAction) isFlag() bool {
	for i := len(AccountRuleActionAll) - 1; i >= 0; i-- {
		expected := AccountRuleAction(2) << uint64(len(AccountRuleActionAll)-1) >> uint64(len(AccountRuleActionAll)-i)
		if expected != AccountRuleActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AccountRuleAction) String() string {
	name, _ := accountRuleActionMap[int32(e)]
	return name
}

func (e AccountRuleAction) ShortString() string {
	name, _ := accountRuleActionShortMap[int32(e)]
	return name
}

func (e AccountRuleAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range AccountRuleActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AccountRuleAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AccountRuleAction(t.Value)
	return nil
}

// SignerRuleResourceReviewableRequest is an XDR NestedStruct defines as:
//
//   struct
//        {
//            //: Type of request
//            ReviewableRequestType requestType;
//            //: Bit mask of tasks that is allowed to add to reviewable request pending tasks
//            uint64 tasksToAdd;
//            //: Bit mask of tasks that is allowed to remove from reviewable request pending tasks
//            uint64 tasksToRemove;
//            //: Bit mask of tasks that is allowed to use as reviewable request pending tasks
//            uint64 allTasks;
//
//            EmptyExt ext;
//        }
//
type SignerRuleResourceReviewableRequest struct {
	RequestType   ReviewableRequestType `json:"requestType,omitempty"`
	TasksToAdd    Uint64                `json:"tasksToAdd,omitempty"`
	TasksToRemove Uint64                `json:"tasksToRemove,omitempty"`
	AllTasks      Uint64                `json:"allTasks,omitempty"`
	Ext           EmptyExt              `json:"ext,omitempty"`
}

// SignerRuleResourceSignerRule is an XDR NestedStruct defines as:
//
//   struct
//        {
//            bool isDefault;
//
//            EmptyExt ext;
//        }
//
type SignerRuleResourceSignerRule struct {
	IsDefault bool     `json:"isDefault,omitempty"`
	Ext       EmptyExt `json:"ext,omitempty"`
}

// SignerRuleResourceSignerRole is an XDR NestedStruct defines as:
//
//   struct
//        {
//            //: For signer role creating resource will be triggered if `roleID` equals `0`
//            uint64 roleID;
//
//            EmptyExt ext;
//        }
//
type SignerRuleResourceSignerRole struct {
	RoleId Uint64   `json:"roleID,omitempty"`
	Ext    EmptyExt `json:"ext,omitempty"`
}

// SignerRuleResourceSigner is an XDR NestedStruct defines as:
//
//   struct
//        {
//            uint64 roleID;
//
//            EmptyExt ext;
//        }
//
type SignerRuleResourceSigner struct {
	RoleId Uint64   `json:"roleID,omitempty"`
	Ext    EmptyExt `json:"ext,omitempty"`
}

// SignerRuleResourceKeyValue is an XDR NestedStruct defines as:
//
//   struct
//        {
//            //: prefix of key
//            longstring keyPrefix;
//
//            //: reserved for future extension
//            EmptyExt ext;
//        }
//
type SignerRuleResourceKeyValue struct {
	KeyPrefix Longstring `json:"keyPrefix,omitempty"`
	Ext       EmptyExt   `json:"ext,omitempty"`
}

// SignerRuleResourceInitiateKycRecovery is an XDR NestedStruct defines as:
//
//   struct
//        {
//            //: Role id
//            uint64 roleID;
//
//            //: reserved for future extension
//            EmptyExt ext;
//        }
//
type SignerRuleResourceInitiateKycRecovery struct {
	RoleId Uint64   `json:"roleID,omitempty"`
	Ext    EmptyExt `json:"ext,omitempty"`
}

// SignerRuleResource is an XDR Union defines as:
//
//   //: Describes properties of some entries that can be used to restrict the usage of entries
//    union SignerRuleResource switch (LedgerEntryType type)
//    {
//    case REVIEWABLE_REQUEST:
//        //: Describes properties that are equal to managed reviewable request entry fields
//        struct
//        {
//            //: Type of request
//            ReviewableRequestType requestType;
//            //: Bit mask of tasks that is allowed to add to reviewable request pending tasks
//            uint64 tasksToAdd;
//            //: Bit mask of tasks that is allowed to remove from reviewable request pending tasks
//            uint64 tasksToRemove;
//            //: Bit mask of tasks that is allowed to use as reviewable request pending tasks
//            uint64 allTasks;
//
//            EmptyExt ext;
//        } reviewableRequest;
//    case ANY:
//        void;
//    case SIGNER_RULE:
//        //: Describes properties that are equal to managed signer rule entry fields
//        struct
//        {
//            bool isDefault;
//
//            EmptyExt ext;
//        } signerRule;
//    case SIGNER_ROLE:
//        //: Describes properties that are equal to managed signer role entry fields
//        struct
//        {
//            //: For signer role creating resource will be triggered if `roleID` equals `0`
//            uint64 roleID;
//
//            EmptyExt ext;
//        } signerRole;
//    case SIGNER:
//        //: Describes properties that are equal to managed signer entry fields
//        struct
//        {
//            uint64 roleID;
//
//            EmptyExt ext;
//        } signer;
//    case KEY_VALUE:
//        //: Describes properties that are equal to managed key value entry fields
//        struct
//        {
//            //: prefix of key
//            longstring keyPrefix;
//
//            //: reserved for future extension
//            EmptyExt ext;
//        } keyValue;
//    case INITIATE_KYC_RECOVERY:
//        struct
//        {
//            //: Role id
//            uint64 roleID;
//
//            //: reserved for future extension
//            EmptyExt ext;
//        } initiateKYCRecovery;
//    default:
//        //: reserved for future extension
//        EmptyExt ext;
//    };
//
type SignerRuleResource struct {
	Type                LedgerEntryType                        `json:"type,omitempty"`
	ReviewableRequest   *SignerRuleResourceReviewableRequest   `json:"reviewableRequest,omitempty"`
	SignerRule          *SignerRuleResourceSignerRule          `json:"signerRule,omitempty"`
	SignerRole          *SignerRuleResourceSignerRole          `json:"signerRole,omitempty"`
	Signer              *SignerRuleResourceSigner              `json:"signer,omitempty"`
	KeyValue            *SignerRuleResourceKeyValue            `json:"keyValue,omitempty"`
	InitiateKycRecovery *SignerRuleResourceInitiateKycRecovery `json:"initiateKYCRecovery,omitempty"`
	Ext                 *EmptyExt                              `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SignerRuleResource) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SignerRuleResource
func (u SignerRuleResource) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeAny:
		return "", true
	case LedgerEntryTypeSignerRule:
		return "SignerRule", true
	case LedgerEntryTypeSignerRole:
		return "SignerRole", true
	case LedgerEntryTypeSigner:
		return "Signer", true
	case LedgerEntryTypeKeyValue:
		return "KeyValue", true
	case LedgerEntryTypeInitiateKycRecovery:
		return "InitiateKycRecovery", true
	default:
		return "Ext", true
	}
}

// NewSignerRuleResource creates a new  SignerRuleResource.
func NewSignerRuleResource(aType LedgerEntryType, value interface{}) (result SignerRuleResource, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeReviewableRequest:
		tv, ok := value.(SignerRuleResourceReviewableRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerRuleResourceReviewableRequest")
			return
		}
		result.ReviewableRequest = &tv
	case LedgerEntryTypeAny:
		// void
	case LedgerEntryTypeSignerRule:
		tv, ok := value.(SignerRuleResourceSignerRule)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerRuleResourceSignerRule")
			return
		}
		result.SignerRule = &tv
	case LedgerEntryTypeSignerRole:
		tv, ok := value.(SignerRuleResourceSignerRole)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerRuleResourceSignerRole")
			return
		}
		result.SignerRole = &tv
	case LedgerEntryTypeSigner:
		tv, ok := value.(SignerRuleResourceSigner)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerRuleResourceSigner")
			return
		}
		result.Signer = &tv
	case LedgerEntryTypeKeyValue:
		tv, ok := value.(SignerRuleResourceKeyValue)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerRuleResourceKeyValue")
			return
		}
		result.KeyValue = &tv
	case LedgerEntryTypeInitiateKycRecovery:
		tv, ok := value.(SignerRuleResourceInitiateKycRecovery)
		if !ok {
			err = fmt.Errorf("invalid value, must be SignerRuleResourceInitiateKycRecovery")
			return
		}
		result.InitiateKycRecovery = &tv
	default:
		tv, ok := value.(EmptyExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be EmptyExt")
			return
		}
		result.Ext = &tv
	}
	return
}

// MustReviewableRequest retrieves the ReviewableRequest value from the union,
// panicing if the value is not set.
func (u SignerRuleResource) MustReviewableRequest() SignerRuleResourceReviewableRequest {
	val, ok := u.GetReviewableRequest()

	if !ok {
		panic("arm ReviewableRequest is not set")
	}

	return val
}

// GetReviewableRequest retrieves the ReviewableRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerRuleResource) GetReviewableRequest() (result SignerRuleResourceReviewableRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewableRequest" {
		result = *u.ReviewableRequest
		ok = true
	}

	return
}

// MustSignerRule retrieves the SignerRule value from the union,
// panicing if the value is not set.
func (u SignerRuleResource) MustSignerRule() SignerRuleResourceSignerRule {
	val, ok := u.GetSignerRule()

	if !ok {
		panic("arm SignerRule is not set")
	}

	return val
}

// GetSignerRule retrieves the SignerRule value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerRuleResource) GetSignerRule() (result SignerRuleResourceSignerRule, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SignerRule" {
		result = *u.SignerRule
		ok = true
	}

	return
}

// MustSignerRole retrieves the SignerRole value from the union,
// panicing if the value is not set.
func (u SignerRuleResource) MustSignerRole() SignerRuleResourceSignerRole {
	val, ok := u.GetSignerRole()

	if !ok {
		panic("arm SignerRole is not set")
	}

	return val
}

// GetSignerRole retrieves the SignerRole value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerRuleResource) GetSignerRole() (result SignerRuleResourceSignerRole, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SignerRole" {
		result = *u.SignerRole
		ok = true
	}

	return
}

// MustSigner retrieves the Signer value from the union,
// panicing if the value is not set.
func (u SignerRuleResource) MustSigner() SignerRuleResourceSigner {
	val, ok := u.GetSigner()

	if !ok {
		panic("arm Signer is not set")
	}

	return val
}

// GetSigner retrieves the Signer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerRuleResource) GetSigner() (result SignerRuleResourceSigner, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Signer" {
		result = *u.Signer
		ok = true
	}

	return
}

// MustKeyValue retrieves the KeyValue value from the union,
// panicing if the value is not set.
func (u SignerRuleResource) MustKeyValue() SignerRuleResourceKeyValue {
	val, ok := u.GetKeyValue()

	if !ok {
		panic("arm KeyValue is not set")
	}

	return val
}

// GetKeyValue retrieves the KeyValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerRuleResource) GetKeyValue() (result SignerRuleResourceKeyValue, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KeyValue" {
		result = *u.KeyValue
		ok = true
	}

	return
}

// MustInitiateKycRecovery retrieves the InitiateKycRecovery value from the union,
// panicing if the value is not set.
func (u SignerRuleResource) MustInitiateKycRecovery() SignerRuleResourceInitiateKycRecovery {
	val, ok := u.GetInitiateKycRecovery()

	if !ok {
		panic("arm InitiateKycRecovery is not set")
	}

	return val
}

// GetInitiateKycRecovery retrieves the InitiateKycRecovery value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerRuleResource) GetInitiateKycRecovery() (result SignerRuleResourceInitiateKycRecovery, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "InitiateKycRecovery" {
		result = *u.InitiateKycRecovery
		ok = true
	}

	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u SignerRuleResource) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SignerRuleResource) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// SignerRuleAction is an XDR Enum defines as:
//
//   //: Actions that can be applied to a signer rule resource
//    enum SignerRuleAction
//    {
//        ANY = 1,
//        CREATE = 2,
//        CREATE_FOR_OTHER = 3,
//        UPDATE = 4,
//        MANAGE = 5,
//        SEND = 6,
//        REMOVE = 7,
//        CANCEL = 8,
//        REVIEW = 9,
//        RECEIVE_ATOMIC_SWAP = 10,
//        PARTICIPATE = 11,
//        BIND = 12,
//        UPDATE_MAX_ISSUANCE = 13,
//        CHECK = 14,
//        CLOSE = 15,
//        UPDATE_END_TIME = 16,
//        CREATE_WITH_TASKS = 17,
//        CREATE_FOR_OTHER_WITH_TASKS = 18
//    };
//
type SignerRuleAction int32

const (
	SignerRuleActionAny                     SignerRuleAction = 1
	SignerRuleActionCreate                  SignerRuleAction = 2
	SignerRuleActionCreateForOther          SignerRuleAction = 3
	SignerRuleActionUpdate                  SignerRuleAction = 4
	SignerRuleActionManage                  SignerRuleAction = 5
	SignerRuleActionSend                    SignerRuleAction = 6
	SignerRuleActionRemove                  SignerRuleAction = 7
	SignerRuleActionCancel                  SignerRuleAction = 8
	SignerRuleActionReview                  SignerRuleAction = 9
	SignerRuleActionReceiveAtomicSwap       SignerRuleAction = 10
	SignerRuleActionParticipate             SignerRuleAction = 11
	SignerRuleActionBind                    SignerRuleAction = 12
	SignerRuleActionUpdateMaxIssuance       SignerRuleAction = 13
	SignerRuleActionCheck                   SignerRuleAction = 14
	SignerRuleActionClose                   SignerRuleAction = 15
	SignerRuleActionUpdateEndTime           SignerRuleAction = 16
	SignerRuleActionCreateWithTasks         SignerRuleAction = 17
	SignerRuleActionCreateForOtherWithTasks SignerRuleAction = 18
)

var SignerRuleActionAll = []SignerRuleAction{
	SignerRuleActionAny,
	SignerRuleActionCreate,
	SignerRuleActionCreateForOther,
	SignerRuleActionUpdate,
	SignerRuleActionManage,
	SignerRuleActionSend,
	SignerRuleActionRemove,
	SignerRuleActionCancel,
	SignerRuleActionReview,
	SignerRuleActionReceiveAtomicSwap,
	SignerRuleActionParticipate,
	SignerRuleActionBind,
	SignerRuleActionUpdateMaxIssuance,
	SignerRuleActionCheck,
	SignerRuleActionClose,
	SignerRuleActionUpdateEndTime,
	SignerRuleActionCreateWithTasks,
	SignerRuleActionCreateForOtherWithTasks,
}

var signerRuleActionMap = map[int32]string{
	1:  "SignerRuleActionAny",
	2:  "SignerRuleActionCreate",
	3:  "SignerRuleActionCreateForOther",
	4:  "SignerRuleActionUpdate",
	5:  "SignerRuleActionManage",
	6:  "SignerRuleActionSend",
	7:  "SignerRuleActionRemove",
	8:  "SignerRuleActionCancel",
	9:  "SignerRuleActionReview",
	10: "SignerRuleActionReceiveAtomicSwap",
	11: "SignerRuleActionParticipate",
	12: "SignerRuleActionBind",
	13: "SignerRuleActionUpdateMaxIssuance",
	14: "SignerRuleActionCheck",
	15: "SignerRuleActionClose",
	16: "SignerRuleActionUpdateEndTime",
	17: "SignerRuleActionCreateWithTasks",
	18: "SignerRuleActionCreateForOtherWithTasks",
}

var signerRuleActionShortMap = map[int32]string{
	1:  "any",
	2:  "create",
	3:  "create_for_other",
	4:  "update",
	5:  "manage",
	6:  "send",
	7:  "remove",
	8:  "cancel",
	9:  "review",
	10: "receive_atomic_swap",
	11: "participate",
	12: "bind",
	13: "update_max_issuance",
	14: "check",
	15: "close",
	16: "update_end_time",
	17: "create_with_tasks",
	18: "create_for_other_with_tasks",
}

var signerRuleActionRevMap = map[string]int32{
	"SignerRuleActionAny":                     1,
	"SignerRuleActionCreate":                  2,
	"SignerRuleActionCreateForOther":          3,
	"SignerRuleActionUpdate":                  4,
	"SignerRuleActionManage":                  5,
	"SignerRuleActionSend":                    6,
	"SignerRuleActionRemove":                  7,
	"SignerRuleActionCancel":                  8,
	"SignerRuleActionReview":                  9,
	"SignerRuleActionReceiveAtomicSwap":       10,
	"SignerRuleActionParticipate":             11,
	"SignerRuleActionBind":                    12,
	"SignerRuleActionUpdateMaxIssuance":       13,
	"SignerRuleActionCheck":                   14,
	"SignerRuleActionClose":                   15,
	"SignerRuleActionUpdateEndTime":           16,
	"SignerRuleActionCreateWithTasks":         17,
	"SignerRuleActionCreateForOtherWithTasks": 18,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SignerRuleAction
func (e SignerRuleAction) ValidEnum(v int32) bool {
	_, ok := signerRuleActionMap[v]
	return ok
}
func (e SignerRuleAction) isFlag() bool {
	for i := len(SignerRuleActionAll) - 1; i >= 0; i-- {
		expected := SignerRuleAction(2) << uint64(len(SignerRuleActionAll)-1) >> uint64(len(SignerRuleActionAll)-i)
		if expected != SignerRuleActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SignerRuleAction) String() string {
	name, _ := signerRuleActionMap[int32(e)]
	return name
}

func (e SignerRuleAction) ShortString() string {
	name, _ := signerRuleActionShortMap[int32(e)]
	return name
}

func (e SignerRuleAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range SignerRuleActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *SignerRuleAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SignerRuleAction(t.Value)
	return nil
}

// ChangeRoleRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ChangeRoleRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ChangeRoleRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ChangeRoleRequestExt
func (u ChangeRoleRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewChangeRoleRequestExt creates a new  ChangeRoleRequestExt.
func NewChangeRoleRequestExt(v LedgerVersion, value interface{}) (result ChangeRoleRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ChangeRoleRequest is an XDR Struct defines as:
//
//   struct ChangeRoleRequest
//    {
//    	AccountID destinationAccount;
//    	uint64 accountRoleToSet;
//
//    	// Sequence number increases when request is rejected
//    	uint32 sequenceNumber;
//
//        longstring creatorDetails; // details set by requester
//
//        // Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ChangeRoleRequest struct {
	DestinationAccount AccountId            `json:"destinationAccount,omitempty"`
	AccountRoleToSet   Uint64               `json:"accountRoleToSet,omitempty"`
	SequenceNumber     Uint32               `json:"sequenceNumber,omitempty"`
	CreatorDetails     Longstring           `json:"creatorDetails,omitempty"`
	Ext                ChangeRoleRequestExt `json:"ext,omitempty"`
}

// KycRecoveryRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type KycRecoveryRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u KycRecoveryRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of KycRecoveryRequestExt
func (u KycRecoveryRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewKycRecoveryRequestExt creates a new  KycRecoveryRequestExt.
func NewKycRecoveryRequestExt(v LedgerVersion, value interface{}) (result KycRecoveryRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// KycRecoveryRequest is an XDR Struct defines as:
//
//   //: KYCRecoveryRequest is used to change signers of target account
//    struct KYCRecoveryRequest {
//        //: Account to be recovered
//        AccountID targetAccount;
//        //: New signers for the target account
//        SignerData signersData<>;
//
//        //: Arbitrary stringified json object that can be used to attach data to be reviewed by an admin
//        longstring creatorDetails; // details set by requester
//        //: Sequence number increases when request is rejected
//        uint32 sequenceNumber;
//
//        //: Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type KycRecoveryRequest struct {
	TargetAccount  AccountId             `json:"targetAccount,omitempty"`
	SignersData    []SignerData          `json:"signersData,omitempty"`
	CreatorDetails Longstring            `json:"creatorDetails,omitempty"`
	SequenceNumber Uint32                `json:"sequenceNumber,omitempty"`
	Ext            KycRecoveryRequestExt `json:"ext,omitempty"`
}

// OperationBody is an XDR NestedUnion defines as:
//
//   union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountOp createAccountOp;
//    	case REVIEW_REQUEST:
//    		ReviewRequestOp reviewRequestOp;
//    	case MANAGE_KEY_VALUE:
//    	    ManageKeyValueOp manageKeyValueOp;
//    	case CREATE_CHANGE_ROLE_REQUEST:
//    		CreateChangeRoleRequestOp createChangeRoleRequestOp;
//        case MANAGE_ACCOUNT_ROLE:
//            ManageAccountRoleOp manageAccountRoleOp;
//        case MANAGE_ACCOUNT_RULE:
//            ManageAccountRuleOp manageAccountRuleOp;
//        case MANAGE_SIGNER:
//            ManageSignerOp manageSignerOp;
//        case MANAGE_SIGNER_ROLE:
//            ManageSignerRoleOp manageSignerRoleOp;
//        case MANAGE_SIGNER_RULE:
//            ManageSignerRuleOp manageSignerRuleOp;
//        case CANCEL_CHANGE_ROLE_REQUEST:
//            CancelChangeRoleRequestOp cancelChangeRoleRequestOp;
//        case INITIATE_KYC_RECOVERY:
//            InitiateKYCRecoveryOp initiateKYCRecoveryOp;
//        case CREATE_KYC_RECOVERY_REQUEST:
//            CreateKYCRecoveryRequestOp createKYCRecoveryRequestOp;
//        }
//
type OperationBody struct {
	Type                       OperationType               `json:"type,omitempty"`
	CreateAccountOp            *CreateAccountOp            `json:"createAccountOp,omitempty"`
	ReviewRequestOp            *ReviewRequestOp            `json:"reviewRequestOp,omitempty"`
	ManageKeyValueOp           *ManageKeyValueOp           `json:"manageKeyValueOp,omitempty"`
	CreateChangeRoleRequestOp  *CreateChangeRoleRequestOp  `json:"createChangeRoleRequestOp,omitempty"`
	ManageAccountRoleOp        *ManageAccountRoleOp        `json:"manageAccountRoleOp,omitempty"`
	ManageAccountRuleOp        *ManageAccountRuleOp        `json:"manageAccountRuleOp,omitempty"`
	ManageSignerOp             *ManageSignerOp             `json:"manageSignerOp,omitempty"`
	ManageSignerRoleOp         *ManageSignerRoleOp         `json:"manageSignerRoleOp,omitempty"`
	ManageSignerRuleOp         *ManageSignerRuleOp         `json:"manageSignerRuleOp,omitempty"`
	CancelChangeRoleRequestOp  *CancelChangeRoleRequestOp  `json:"cancelChangeRoleRequestOp,omitempty"`
	InitiateKycRecoveryOp      *InitiateKycRecoveryOp      `json:"initiateKYCRecoveryOp,omitempty"`
	CreateKycRecoveryRequestOp *CreateKycRecoveryRequestOp `json:"createKYCRecoveryRequestOp,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationBody) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationBody
func (u OperationBody) ArmForSwitch(sw int32) (string, bool) {
	switch OperationType(sw) {
	case OperationTypeCreateAccount:
		return "CreateAccountOp", true
	case OperationTypeReviewRequest:
		return "ReviewRequestOp", true
	case OperationTypeManageKeyValue:
		return "ManageKeyValueOp", true
	case OperationTypeCreateChangeRoleRequest:
		return "CreateChangeRoleRequestOp", true
	case OperationTypeManageAccountRole:
		return "ManageAccountRoleOp", true
	case OperationTypeManageAccountRule:
		return "ManageAccountRuleOp", true
	case OperationTypeManageSigner:
		return "ManageSignerOp", true
	case OperationTypeManageSignerRole:
		return "ManageSignerRoleOp", true
	case OperationTypeManageSignerRule:
		return "ManageSignerRuleOp", true
	case OperationTypeCancelChangeRoleRequest:
		return "CancelChangeRoleRequestOp", true
	case OperationTypeInitiateKycRecovery:
		return "InitiateKycRecoveryOp", true
	case OperationTypeCreateKycRecoveryRequest:
		return "CreateKycRecoveryRequestOp", true
	}
	return "-", false
}

// NewOperationBody creates a new  OperationBody.
func NewOperationBody(aType OperationType, value interface{}) (result OperationBody, err error) {
	result.Type = aType
	switch OperationType(aType) {
	case OperationTypeCreateAccount:
		tv, ok := value.(CreateAccountOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountOp")
			return
		}
		result.CreateAccountOp = &tv
	case OperationTypeReviewRequest:
		tv, ok := value.(ReviewRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestOp")
			return
		}
		result.ReviewRequestOp = &tv
	case OperationTypeManageKeyValue:
		tv, ok := value.(ManageKeyValueOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageKeyValueOp")
			return
		}
		result.ManageKeyValueOp = &tv
	case OperationTypeCreateChangeRoleRequest:
		tv, ok := value.(CreateChangeRoleRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateChangeRoleRequestOp")
			return
		}
		result.CreateChangeRoleRequestOp = &tv
	case OperationTypeManageAccountRole:
		tv, ok := value.(ManageAccountRoleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAccountRoleOp")
			return
		}
		result.ManageAccountRoleOp = &tv
	case OperationTypeManageAccountRule:
		tv, ok := value.(ManageAccountRuleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAccountRuleOp")
			return
		}
		result.ManageAccountRuleOp = &tv
	case OperationTypeManageSigner:
		tv, ok := value.(ManageSignerOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSignerOp")
			return
		}
		result.ManageSignerOp = &tv
	case OperationTypeManageSignerRole:
		tv, ok := value.(ManageSignerRoleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSignerRoleOp")
			return
		}
		result.ManageSignerRoleOp = &tv
	case OperationTypeManageSignerRule:
		tv, ok := value.(ManageSignerRuleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSignerRuleOp")
			return
		}
		result.ManageSignerRuleOp = &tv
	case OperationTypeCancelChangeRoleRequest:
		tv, ok := value.(CancelChangeRoleRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CancelChangeRoleRequestOp")
			return
		}
		result.CancelChangeRoleRequestOp = &tv
	case OperationTypeInitiateKycRecovery:
		tv, ok := value.(InitiateKycRecoveryOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be InitiateKycRecoveryOp")
			return
		}
		result.InitiateKycRecoveryOp = &tv
	case OperationTypeCreateKycRecoveryRequest:
		tv, ok := value.(CreateKycRecoveryRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateKycRecoveryRequestOp")
			return
		}
		result.CreateKycRecoveryRequestOp = &tv
	}
	return
}

// MustCreateAccountOp retrieves the CreateAccountOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateAccountOp() CreateAccountOp {
	val, ok := u.GetCreateAccountOp()

	if !ok {
		panic("arm CreateAccountOp is not set")
	}

	return val
}

// GetCreateAccountOp retrieves the CreateAccountOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateAccountOp() (result CreateAccountOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAccountOp" {
		result = *u.CreateAccountOp
		ok = true
	}

	return
}

// MustReviewRequestOp retrieves the ReviewRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustReviewRequestOp() ReviewRequestOp {
	val, ok := u.GetReviewRequestOp()

	if !ok {
		panic("arm ReviewRequestOp is not set")
	}

	return val
}

// GetReviewRequestOp retrieves the ReviewRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetReviewRequestOp() (result ReviewRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewRequestOp" {
		result = *u.ReviewRequestOp
		ok = true
	}

	return
}

// MustManageKeyValueOp retrieves the ManageKeyValueOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageKeyValueOp() ManageKeyValueOp {
	val, ok := u.GetManageKeyValueOp()

	if !ok {
		panic("arm ManageKeyValueOp is not set")
	}

	return val
}

// GetManageKeyValueOp retrieves the ManageKeyValueOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageKeyValueOp() (result ManageKeyValueOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageKeyValueOp" {
		result = *u.ManageKeyValueOp
		ok = true
	}

	return
}

// MustCreateChangeRoleRequestOp retrieves the CreateChangeRoleRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateChangeRoleRequestOp() CreateChangeRoleRequestOp {
	val, ok := u.GetCreateChangeRoleRequestOp()

	if !ok {
		panic("arm CreateChangeRoleRequestOp is not set")
	}

	return val
}

// GetCreateChangeRoleRequestOp retrieves the CreateChangeRoleRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateChangeRoleRequestOp() (result CreateChangeRoleRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateChangeRoleRequestOp" {
		result = *u.CreateChangeRoleRequestOp
		ok = true
	}

	return
}

// MustManageAccountRoleOp retrieves the ManageAccountRoleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageAccountRoleOp() ManageAccountRoleOp {
	val, ok := u.GetManageAccountRoleOp()

	if !ok {
		panic("arm ManageAccountRoleOp is not set")
	}

	return val
}

// GetManageAccountRoleOp retrieves the ManageAccountRoleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageAccountRoleOp() (result ManageAccountRoleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAccountRoleOp" {
		result = *u.ManageAccountRoleOp
		ok = true
	}

	return
}

// MustManageAccountRuleOp retrieves the ManageAccountRuleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageAccountRuleOp() ManageAccountRuleOp {
	val, ok := u.GetManageAccountRuleOp()

	if !ok {
		panic("arm ManageAccountRuleOp is not set")
	}

	return val
}

// GetManageAccountRuleOp retrieves the ManageAccountRuleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageAccountRuleOp() (result ManageAccountRuleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAccountRuleOp" {
		result = *u.ManageAccountRuleOp
		ok = true
	}

	return
}

// MustManageSignerOp retrieves the ManageSignerOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageSignerOp() ManageSignerOp {
	val, ok := u.GetManageSignerOp()

	if !ok {
		panic("arm ManageSignerOp is not set")
	}

	return val
}

// GetManageSignerOp retrieves the ManageSignerOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageSignerOp() (result ManageSignerOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSignerOp" {
		result = *u.ManageSignerOp
		ok = true
	}

	return
}

// MustManageSignerRoleOp retrieves the ManageSignerRoleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageSignerRoleOp() ManageSignerRoleOp {
	val, ok := u.GetManageSignerRoleOp()

	if !ok {
		panic("arm ManageSignerRoleOp is not set")
	}

	return val
}

// GetManageSignerRoleOp retrieves the ManageSignerRoleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageSignerRoleOp() (result ManageSignerRoleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSignerRoleOp" {
		result = *u.ManageSignerRoleOp
		ok = true
	}

	return
}

// MustManageSignerRuleOp retrieves the ManageSignerRuleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageSignerRuleOp() ManageSignerRuleOp {
	val, ok := u.GetManageSignerRuleOp()

	if !ok {
		panic("arm ManageSignerRuleOp is not set")
	}

	return val
}

// GetManageSignerRuleOp retrieves the ManageSignerRuleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageSignerRuleOp() (result ManageSignerRuleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSignerRuleOp" {
		result = *u.ManageSignerRuleOp
		ok = true
	}

	return
}

// MustCancelChangeRoleRequestOp retrieves the CancelChangeRoleRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCancelChangeRoleRequestOp() CancelChangeRoleRequestOp {
	val, ok := u.GetCancelChangeRoleRequestOp()

	if !ok {
		panic("arm CancelChangeRoleRequestOp is not set")
	}

	return val
}

// GetCancelChangeRoleRequestOp retrieves the CancelChangeRoleRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCancelChangeRoleRequestOp() (result CancelChangeRoleRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CancelChangeRoleRequestOp" {
		result = *u.CancelChangeRoleRequestOp
		ok = true
	}

	return
}

// MustInitiateKycRecoveryOp retrieves the InitiateKycRecoveryOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustInitiateKycRecoveryOp() InitiateKycRecoveryOp {
	val, ok := u.GetInitiateKycRecoveryOp()

	if !ok {
		panic("arm InitiateKycRecoveryOp is not set")
	}

	return val
}

// GetInitiateKycRecoveryOp retrieves the InitiateKycRecoveryOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetInitiateKycRecoveryOp() (result InitiateKycRecoveryOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "InitiateKycRecoveryOp" {
		result = *u.InitiateKycRecoveryOp
		ok = true
	}

	return
}

// MustCreateKycRecoveryRequestOp retrieves the CreateKycRecoveryRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateKycRecoveryRequestOp() CreateKycRecoveryRequestOp {
	val, ok := u.GetCreateKycRecoveryRequestOp()

	if !ok {
		panic("arm CreateKycRecoveryRequestOp is not set")
	}

	return val
}

// GetCreateKycRecoveryRequestOp retrieves the CreateKycRecoveryRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateKycRecoveryRequestOp() (result CreateKycRecoveryRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateKycRecoveryRequestOp" {
		result = *u.CreateKycRecoveryRequestOp
		ok = true
	}

	return
}

// Operation is an XDR Struct defines as:
//
//   //: An operation is the lowest unit of work that a transaction does
//    struct Operation
//    {
//        //: sourceAccount is the account used to run the operation
//        //: if not set, the runtime defaults to "sourceAccount" specified at
//        //: the transaction level
//        AccountID* sourceAccount;
//
//        union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountOp createAccountOp;
//    	case REVIEW_REQUEST:
//    		ReviewRequestOp reviewRequestOp;
//    	case MANAGE_KEY_VALUE:
//    	    ManageKeyValueOp manageKeyValueOp;
//    	case CREATE_CHANGE_ROLE_REQUEST:
//    		CreateChangeRoleRequestOp createChangeRoleRequestOp;
//        case MANAGE_ACCOUNT_ROLE:
//            ManageAccountRoleOp manageAccountRoleOp;
//        case MANAGE_ACCOUNT_RULE:
//            ManageAccountRuleOp manageAccountRuleOp;
//        case MANAGE_SIGNER:
//            ManageSignerOp manageSignerOp;
//        case MANAGE_SIGNER_ROLE:
//            ManageSignerRoleOp manageSignerRoleOp;
//        case MANAGE_SIGNER_RULE:
//            ManageSignerRuleOp manageSignerRuleOp;
//        case CANCEL_CHANGE_ROLE_REQUEST:
//            CancelChangeRoleRequestOp cancelChangeRoleRequestOp;
//        case INITIATE_KYC_RECOVERY:
//            InitiateKYCRecoveryOp initiateKYCRecoveryOp;
//        case CREATE_KYC_RECOVERY_REQUEST:
//            CreateKYCRecoveryRequestOp createKYCRecoveryRequestOp;
//        }
//        body;
//    };
//
type Operation struct {
	SourceAccount *AccountId    `json:"sourceAccount,omitempty"`
	Body          OperationBody `json:"body,omitempty"`
}

// MemoType is an XDR Enum defines as:
//
//   enum MemoType
//    {
//        MEMO_NONE = 0,
//        MEMO_TEXT = 1,
//        MEMO_ID = 2,
//        MEMO_HASH = 3,
//        MEMO_RETURN = 4
//    };
//
type MemoType int32

const (
	MemoTypeMemoNone   MemoType = 0
	MemoTypeMemoText   MemoType = 1
	MemoTypeMemoId     MemoType = 2
	MemoTypeMemoHash   MemoType = 3
	MemoTypeMemoReturn MemoType = 4
)

var MemoTypeAll = []MemoType{
	MemoTypeMemoNone,
	MemoTypeMemoText,
	MemoTypeMemoId,
	MemoTypeMemoHash,
	MemoTypeMemoReturn,
}

var memoTypeMap = map[int32]string{
	0: "MemoTypeMemoNone",
	1: "MemoTypeMemoText",
	2: "MemoTypeMemoId",
	3: "MemoTypeMemoHash",
	4: "MemoTypeMemoReturn",
}

var memoTypeShortMap = map[int32]string{
	0: "memo_none",
	1: "memo_text",
	2: "memo_id",
	3: "memo_hash",
	4: "memo_return",
}

var memoTypeRevMap = map[string]int32{
	"MemoTypeMemoNone":   0,
	"MemoTypeMemoText":   1,
	"MemoTypeMemoId":     2,
	"MemoTypeMemoHash":   3,
	"MemoTypeMemoReturn": 4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for MemoType
func (e MemoType) ValidEnum(v int32) bool {
	_, ok := memoTypeMap[v]
	return ok
}
func (e MemoType) isFlag() bool {
	for i := len(MemoTypeAll) - 1; i >= 0; i-- {
		expected := MemoType(2) << uint64(len(MemoTypeAll)-1) >> uint64(len(MemoTypeAll)-i)
		if expected != MemoTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e MemoType) String() string {
	name, _ := memoTypeMap[int32(e)]
	return name
}

func (e MemoType) ShortString() string {
	name, _ := memoTypeShortMap[int32(e)]
	return name
}

func (e MemoType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range MemoTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *MemoType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = MemoType(t.Value)
	return nil
}

// Memo is an XDR Union defines as:
//
//   union Memo switch (MemoType type)
//    {
//    case MEMO_NONE:
//        void;
//    case MEMO_TEXT:
//        string text<28>;
//    case MEMO_ID:
//        uint64 id;
//    case MEMO_HASH:
//        Hash hash; // the hash of what to pull from the content server
//    case MEMO_RETURN:
//        Hash retHash; // the hash of the tx you are rejecting
//    };
//
type Memo struct {
	Type    MemoType `json:"type,omitempty"`
	Text    *string  `json:"text,omitempty" xdrmaxsize:"28"`
	Id      *Uint64  `json:"id,omitempty"`
	Hash    *Hash    `json:"hash,omitempty"`
	RetHash *Hash    `json:"retHash,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Memo) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Memo
func (u Memo) ArmForSwitch(sw int32) (string, bool) {
	switch MemoType(sw) {
	case MemoTypeMemoNone:
		return "", true
	case MemoTypeMemoText:
		return "Text", true
	case MemoTypeMemoId:
		return "Id", true
	case MemoTypeMemoHash:
		return "Hash", true
	case MemoTypeMemoReturn:
		return "RetHash", true
	}
	return "-", false
}

// NewMemo creates a new  Memo.
func NewMemo(aType MemoType, value interface{}) (result Memo, err error) {
	result.Type = aType
	switch MemoType(aType) {
	case MemoTypeMemoNone:
		// void
	case MemoTypeMemoText:
		tv, ok := value.(string)
		if !ok {
			err = fmt.Errorf("invalid value, must be string")
			return
		}
		result.Text = &tv
	case MemoTypeMemoId:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.Id = &tv
	case MemoTypeMemoHash:
		tv, ok := value.(Hash)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hash")
			return
		}
		result.Hash = &tv
	case MemoTypeMemoReturn:
		tv, ok := value.(Hash)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hash")
			return
		}
		result.RetHash = &tv
	}
	return
}

// MustText retrieves the Text value from the union,
// panicing if the value is not set.
func (u Memo) MustText() string {
	val, ok := u.GetText()

	if !ok {
		panic("arm Text is not set")
	}

	return val
}

// GetText retrieves the Text value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetText() (result string, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Text" {
		result = *u.Text
		ok = true
	}

	return
}

// MustId retrieves the Id value from the union,
// panicing if the value is not set.
func (u Memo) MustId() Uint64 {
	val, ok := u.GetId()

	if !ok {
		panic("arm Id is not set")
	}

	return val
}

// GetId retrieves the Id value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Id" {
		result = *u.Id
		ok = true
	}

	return
}

// MustHash retrieves the Hash value from the union,
// panicing if the value is not set.
func (u Memo) MustHash() Hash {
	val, ok := u.GetHash()

	if !ok {
		panic("arm Hash is not set")
	}

	return val
}

// GetHash retrieves the Hash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetHash() (result Hash, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hash" {
		result = *u.Hash
		ok = true
	}

	return
}

// MustRetHash retrieves the RetHash value from the union,
// panicing if the value is not set.
func (u Memo) MustRetHash() Hash {
	val, ok := u.GetRetHash()

	if !ok {
		panic("arm RetHash is not set")
	}

	return val
}

// GetRetHash retrieves the RetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetRetHash() (result Hash, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RetHash" {
		result = *u.RetHash
		ok = true
	}

	return
}

// TimeBounds is an XDR Struct defines as:
//
//   struct TimeBounds
//    {
//        //: specifies inclusive min ledger close time after which transaction is valid
//        uint64 minTime;
//        //: specifies inclusive max ledger close time before which transaction is valid.
//        //: note: transaction will be rejected if max time exceeds close time of current ledger on more then [`tx_expiration_period`](https://tokend.gitlab.io/horizon/#operation/info)
//        uint64 maxTime; // 0 here means no maxTime
//    };
//
type TimeBounds struct {
	MinTime Uint64 `json:"minTime,omitempty"`
	MaxTime Uint64 `json:"maxTime,omitempty"`
}

// TransactionExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionExt
func (u TransactionExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionExt creates a new  TransactionExt.
func NewTransactionExt(v LedgerVersion, value interface{}) (result TransactionExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Transaction is an XDR Struct defines as:
//
//   //: Transaction is a container for a set of operations
//    //:    - is executed by an account
//    //:    - operations are executed in order as one ACID transaction
//    //: (either all operations are applied or none are if any returns a failing code)
//    struct Transaction
//    {
//        //: account used to run the transaction
//        AccountID sourceAccount;
//
//        //: random number used to ensure there is no hash collisions
//        Salt salt;
//
//        //: validity range (inclusive) for the last ledger close time
//        TimeBounds timeBounds;
//
//        //: allows to attach additional data to the transactions
//        Memo memo;
//
//        //: list of operations to be applied. Max size is 100
//        Operation operations<100>;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type Transaction struct {
	SourceAccount AccountId      `json:"sourceAccount,omitempty"`
	Salt          Salt           `json:"salt,omitempty"`
	TimeBounds    TimeBounds     `json:"timeBounds,omitempty"`
	Memo          Memo           `json:"memo,omitempty"`
	Operations    []Operation    `json:"operations,omitempty" xdrmaxsize:"100"`
	Ext           TransactionExt `json:"ext,omitempty"`
}

// TransactionEnvelope is an XDR Struct defines as:
//
//   struct TransactionEnvelope
//    {
//        Transaction tx;
//        //: list of signatures used to authorize transaction
//        DecoratedSignature signatures<20>;
//    };
//
type TransactionEnvelope struct {
	Tx         Transaction          `json:"tx,omitempty"`
	Signatures []DecoratedSignature `json:"signatures,omitempty" xdrmaxsize:"20"`
}

// OperationResultCode is an XDR Enum defines as:
//
//   enum OperationResultCode
//    {
//        opINNER = 0, // inner object result is valid
//
//        opBAD_AUTH = -1,      // too few valid signatures / wrong network
//        opNO_ACCOUNT = -2,    // source account was not found
//    	opNOT_ALLOWED = -3,   // operation is not allowed for this type of source account
//    	opACCOUNT_BLOCKED = -4, // account is blocked
//        opNO_COUNTERPARTY = -5,
//        opCOUNTERPARTY_BLOCKED = -6,
//        opCOUNTERPARTY_WRONG_TYPE = -7,
//        opBAD_AUTH_EXTRA = -8,
//        opNO_ROLE_PERMISSION = -9, // not allowed for this role of source account
//        opNO_ENTRY = -10,
//        opNOT_SUPPORTED = -11,
//        //: operation was skipped cause of failure validation of previous operation
//        opSKIPPED = -12
//    };
//
type OperationResultCode int32

const (
	OperationResultCodeOpInner                 OperationResultCode = 0
	OperationResultCodeOpBadAuth               OperationResultCode = -1
	OperationResultCodeOpNoAccount             OperationResultCode = -2
	OperationResultCodeOpNotAllowed            OperationResultCode = -3
	OperationResultCodeOpAccountBlocked        OperationResultCode = -4
	OperationResultCodeOpNoCounterparty        OperationResultCode = -5
	OperationResultCodeOpCounterpartyBlocked   OperationResultCode = -6
	OperationResultCodeOpCounterpartyWrongType OperationResultCode = -7
	OperationResultCodeOpBadAuthExtra          OperationResultCode = -8
	OperationResultCodeOpNoRolePermission      OperationResultCode = -9
	OperationResultCodeOpNoEntry               OperationResultCode = -10
	OperationResultCodeOpNotSupported          OperationResultCode = -11
	OperationResultCodeOpSkipped               OperationResultCode = -12
)

var OperationResultCodeAll = []OperationResultCode{
	OperationResultCodeOpInner,
	OperationResultCodeOpBadAuth,
	OperationResultCodeOpNoAccount,
	OperationResultCodeOpNotAllowed,
	OperationResultCodeOpAccountBlocked,
	OperationResultCodeOpNoCounterparty,
	OperationResultCodeOpCounterpartyBlocked,
	OperationResultCodeOpCounterpartyWrongType,
	OperationResultCodeOpBadAuthExtra,
	OperationResultCodeOpNoRolePermission,
	OperationResultCodeOpNoEntry,
	OperationResultCodeOpNotSupported,
	OperationResultCodeOpSkipped,
}

var operationResultCodeMap = map[int32]string{
	0:   "OperationResultCodeOpInner",
	-1:  "OperationResultCodeOpBadAuth",
	-2:  "OperationResultCodeOpNoAccount",
	-3:  "OperationResultCodeOpNotAllowed",
	-4:  "OperationResultCodeOpAccountBlocked",
	-5:  "OperationResultCodeOpNoCounterparty",
	-6:  "OperationResultCodeOpCounterpartyBlocked",
	-7:  "OperationResultCodeOpCounterpartyWrongType",
	-8:  "OperationResultCodeOpBadAuthExtra",
	-9:  "OperationResultCodeOpNoRolePermission",
	-10: "OperationResultCodeOpNoEntry",
	-11: "OperationResultCodeOpNotSupported",
	-12: "OperationResultCodeOpSkipped",
}

var operationResultCodeShortMap = map[int32]string{
	0:   "op_inner",
	-1:  "op_bad_auth",
	-2:  "op_no_account",
	-3:  "op_not_allowed",
	-4:  "op_account_blocked",
	-5:  "op_no_counterparty",
	-6:  "op_counterparty_blocked",
	-7:  "op_counterparty_wrong_type",
	-8:  "op_bad_auth_extra",
	-9:  "op_no_role_permission",
	-10: "op_no_entry",
	-11: "op_not_supported",
	-12: "op_skipped",
}

var operationResultCodeRevMap = map[string]int32{
	"OperationResultCodeOpInner":                 0,
	"OperationResultCodeOpBadAuth":               -1,
	"OperationResultCodeOpNoAccount":             -2,
	"OperationResultCodeOpNotAllowed":            -3,
	"OperationResultCodeOpAccountBlocked":        -4,
	"OperationResultCodeOpNoCounterparty":        -5,
	"OperationResultCodeOpCounterpartyBlocked":   -6,
	"OperationResultCodeOpCounterpartyWrongType": -7,
	"OperationResultCodeOpBadAuthExtra":          -8,
	"OperationResultCodeOpNoRolePermission":      -9,
	"OperationResultCodeOpNoEntry":               -10,
	"OperationResultCodeOpNotSupported":          -11,
	"OperationResultCodeOpSkipped":               -12,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OperationResultCode
func (e OperationResultCode) ValidEnum(v int32) bool {
	_, ok := operationResultCodeMap[v]
	return ok
}
func (e OperationResultCode) isFlag() bool {
	for i := len(OperationResultCodeAll) - 1; i >= 0; i-- {
		expected := OperationResultCode(2) << uint64(len(OperationResultCodeAll)-1) >> uint64(len(OperationResultCodeAll)-i)
		if expected != OperationResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e OperationResultCode) String() string {
	name, _ := operationResultCodeMap[int32(e)]
	return name
}

func (e OperationResultCode) ShortString() string {
	name, _ := operationResultCodeShortMap[int32(e)]
	return name
}

func (e OperationResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range OperationResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *OperationResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = OperationResultCode(t.Value)
	return nil
}

// AccountRuleRequirement is an XDR Struct defines as:
//
//   //: Defines requirements for tx or operation which were not fulfilled
//    struct AccountRuleRequirement
//    {
//    	//: defines resources to which access was denied
//        AccountRuleResource resource;
//    	//: defines action which was denied
//        AccountRuleAction action;
//    	//: defines account for which requirements were not met
//    	AccountID account;
//
//    	//: reserved for future extension
//        EmptyExt ext;
//    };
//
type AccountRuleRequirement struct {
	Resource AccountRuleResource `json:"resource,omitempty"`
	Action   AccountRuleAction   `json:"action,omitempty"`
	Account  AccountId           `json:"account,omitempty"`
	Ext      EmptyExt            `json:"ext,omitempty"`
}

// OperationResultTr is an XDR NestedUnion defines as:
//
//   union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountResult createAccountResult;
//    	case REVIEW_REQUEST:
//    		ReviewRequestResult reviewRequestResult;
//    	case MANAGE_KEY_VALUE:
//    	    ManageKeyValueResult manageKeyValueResult;
//    	case CREATE_CHANGE_ROLE_REQUEST:
//    	    CreateChangeRoleRequestResult createChangeRoleRequestResult;
//        case MANAGE_ACCOUNT_ROLE:
//            ManageAccountRoleResult manageAccountRoleResult;
//        case MANAGE_ACCOUNT_RULE:
//            ManageAccountRuleResult manageAccountRuleResult;
//        case MANAGE_SIGNER:
//            ManageSignerResult manageSignerResult;
//        case MANAGE_SIGNER_ROLE:
//            ManageSignerRoleResult manageSignerRoleResult;
//        case MANAGE_SIGNER_RULE:
//            ManageSignerRuleResult manageSignerRuleResult;
//        case CANCEL_CHANGE_ROLE_REQUEST:
//            CancelChangeRoleRequestResult cancelChangeRoleRequestResult;
//        case CREATE_KYC_RECOVERY_REQUEST:
//            CreateKYCRecoveryRequestResult createKYCRecoveryRequestResult;
//        case INITIATE_KYC_RECOVERY:
//            InitiateKYCRecoveryResult initiateKYCRecoveryResult;
//        }
//
type OperationResultTr struct {
	Type                           OperationType                   `json:"type,omitempty"`
	CreateAccountResult            *CreateAccountResult            `json:"createAccountResult,omitempty"`
	ReviewRequestResult            *ReviewRequestResult            `json:"reviewRequestResult,omitempty"`
	ManageKeyValueResult           *ManageKeyValueResult           `json:"manageKeyValueResult,omitempty"`
	CreateChangeRoleRequestResult  *CreateChangeRoleRequestResult  `json:"createChangeRoleRequestResult,omitempty"`
	ManageAccountRoleResult        *ManageAccountRoleResult        `json:"manageAccountRoleResult,omitempty"`
	ManageAccountRuleResult        *ManageAccountRuleResult        `json:"manageAccountRuleResult,omitempty"`
	ManageSignerResult             *ManageSignerResult             `json:"manageSignerResult,omitempty"`
	ManageSignerRoleResult         *ManageSignerRoleResult         `json:"manageSignerRoleResult,omitempty"`
	ManageSignerRuleResult         *ManageSignerRuleResult         `json:"manageSignerRuleResult,omitempty"`
	CancelChangeRoleRequestResult  *CancelChangeRoleRequestResult  `json:"cancelChangeRoleRequestResult,omitempty"`
	CreateKycRecoveryRequestResult *CreateKycRecoveryRequestResult `json:"createKYCRecoveryRequestResult,omitempty"`
	InitiateKycRecoveryResult      *InitiateKycRecoveryResult      `json:"initiateKYCRecoveryResult,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationResultTr) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationResultTr
func (u OperationResultTr) ArmForSwitch(sw int32) (string, bool) {
	switch OperationType(sw) {
	case OperationTypeCreateAccount:
		return "CreateAccountResult", true
	case OperationTypeReviewRequest:
		return "ReviewRequestResult", true
	case OperationTypeManageKeyValue:
		return "ManageKeyValueResult", true
	case OperationTypeCreateChangeRoleRequest:
		return "CreateChangeRoleRequestResult", true
	case OperationTypeManageAccountRole:
		return "ManageAccountRoleResult", true
	case OperationTypeManageAccountRule:
		return "ManageAccountRuleResult", true
	case OperationTypeManageSigner:
		return "ManageSignerResult", true
	case OperationTypeManageSignerRole:
		return "ManageSignerRoleResult", true
	case OperationTypeManageSignerRule:
		return "ManageSignerRuleResult", true
	case OperationTypeCancelChangeRoleRequest:
		return "CancelChangeRoleRequestResult", true
	case OperationTypeCreateKycRecoveryRequest:
		return "CreateKycRecoveryRequestResult", true
	case OperationTypeInitiateKycRecovery:
		return "InitiateKycRecoveryResult", true
	}
	return "-", false
}

// NewOperationResultTr creates a new  OperationResultTr.
func NewOperationResultTr(aType OperationType, value interface{}) (result OperationResultTr, err error) {
	result.Type = aType
	switch OperationType(aType) {
	case OperationTypeCreateAccount:
		tv, ok := value.(CreateAccountResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountResult")
			return
		}
		result.CreateAccountResult = &tv
	case OperationTypeReviewRequest:
		tv, ok := value.(ReviewRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestResult")
			return
		}
		result.ReviewRequestResult = &tv
	case OperationTypeManageKeyValue:
		tv, ok := value.(ManageKeyValueResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageKeyValueResult")
			return
		}
		result.ManageKeyValueResult = &tv
	case OperationTypeCreateChangeRoleRequest:
		tv, ok := value.(CreateChangeRoleRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateChangeRoleRequestResult")
			return
		}
		result.CreateChangeRoleRequestResult = &tv
	case OperationTypeManageAccountRole:
		tv, ok := value.(ManageAccountRoleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAccountRoleResult")
			return
		}
		result.ManageAccountRoleResult = &tv
	case OperationTypeManageAccountRule:
		tv, ok := value.(ManageAccountRuleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAccountRuleResult")
			return
		}
		result.ManageAccountRuleResult = &tv
	case OperationTypeManageSigner:
		tv, ok := value.(ManageSignerResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSignerResult")
			return
		}
		result.ManageSignerResult = &tv
	case OperationTypeManageSignerRole:
		tv, ok := value.(ManageSignerRoleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSignerRoleResult")
			return
		}
		result.ManageSignerRoleResult = &tv
	case OperationTypeManageSignerRule:
		tv, ok := value.(ManageSignerRuleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSignerRuleResult")
			return
		}
		result.ManageSignerRuleResult = &tv
	case OperationTypeCancelChangeRoleRequest:
		tv, ok := value.(CancelChangeRoleRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CancelChangeRoleRequestResult")
			return
		}
		result.CancelChangeRoleRequestResult = &tv
	case OperationTypeCreateKycRecoveryRequest:
		tv, ok := value.(CreateKycRecoveryRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateKycRecoveryRequestResult")
			return
		}
		result.CreateKycRecoveryRequestResult = &tv
	case OperationTypeInitiateKycRecovery:
		tv, ok := value.(InitiateKycRecoveryResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be InitiateKycRecoveryResult")
			return
		}
		result.InitiateKycRecoveryResult = &tv
	}
	return
}

// MustCreateAccountResult retrieves the CreateAccountResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateAccountResult() CreateAccountResult {
	val, ok := u.GetCreateAccountResult()

	if !ok {
		panic("arm CreateAccountResult is not set")
	}

	return val
}

// GetCreateAccountResult retrieves the CreateAccountResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateAccountResult() (result CreateAccountResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAccountResult" {
		result = *u.CreateAccountResult
		ok = true
	}

	return
}

// MustReviewRequestResult retrieves the ReviewRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustReviewRequestResult() ReviewRequestResult {
	val, ok := u.GetReviewRequestResult()

	if !ok {
		panic("arm ReviewRequestResult is not set")
	}

	return val
}

// GetReviewRequestResult retrieves the ReviewRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetReviewRequestResult() (result ReviewRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewRequestResult" {
		result = *u.ReviewRequestResult
		ok = true
	}

	return
}

// MustManageKeyValueResult retrieves the ManageKeyValueResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageKeyValueResult() ManageKeyValueResult {
	val, ok := u.GetManageKeyValueResult()

	if !ok {
		panic("arm ManageKeyValueResult is not set")
	}

	return val
}

// GetManageKeyValueResult retrieves the ManageKeyValueResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageKeyValueResult() (result ManageKeyValueResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageKeyValueResult" {
		result = *u.ManageKeyValueResult
		ok = true
	}

	return
}

// MustCreateChangeRoleRequestResult retrieves the CreateChangeRoleRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateChangeRoleRequestResult() CreateChangeRoleRequestResult {
	val, ok := u.GetCreateChangeRoleRequestResult()

	if !ok {
		panic("arm CreateChangeRoleRequestResult is not set")
	}

	return val
}

// GetCreateChangeRoleRequestResult retrieves the CreateChangeRoleRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateChangeRoleRequestResult() (result CreateChangeRoleRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateChangeRoleRequestResult" {
		result = *u.CreateChangeRoleRequestResult
		ok = true
	}

	return
}

// MustManageAccountRoleResult retrieves the ManageAccountRoleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageAccountRoleResult() ManageAccountRoleResult {
	val, ok := u.GetManageAccountRoleResult()

	if !ok {
		panic("arm ManageAccountRoleResult is not set")
	}

	return val
}

// GetManageAccountRoleResult retrieves the ManageAccountRoleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageAccountRoleResult() (result ManageAccountRoleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAccountRoleResult" {
		result = *u.ManageAccountRoleResult
		ok = true
	}

	return
}

// MustManageAccountRuleResult retrieves the ManageAccountRuleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageAccountRuleResult() ManageAccountRuleResult {
	val, ok := u.GetManageAccountRuleResult()

	if !ok {
		panic("arm ManageAccountRuleResult is not set")
	}

	return val
}

// GetManageAccountRuleResult retrieves the ManageAccountRuleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageAccountRuleResult() (result ManageAccountRuleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAccountRuleResult" {
		result = *u.ManageAccountRuleResult
		ok = true
	}

	return
}

// MustManageSignerResult retrieves the ManageSignerResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageSignerResult() ManageSignerResult {
	val, ok := u.GetManageSignerResult()

	if !ok {
		panic("arm ManageSignerResult is not set")
	}

	return val
}

// GetManageSignerResult retrieves the ManageSignerResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageSignerResult() (result ManageSignerResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSignerResult" {
		result = *u.ManageSignerResult
		ok = true
	}

	return
}

// MustManageSignerRoleResult retrieves the ManageSignerRoleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageSignerRoleResult() ManageSignerRoleResult {
	val, ok := u.GetManageSignerRoleResult()

	if !ok {
		panic("arm ManageSignerRoleResult is not set")
	}

	return val
}

// GetManageSignerRoleResult retrieves the ManageSignerRoleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageSignerRoleResult() (result ManageSignerRoleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSignerRoleResult" {
		result = *u.ManageSignerRoleResult
		ok = true
	}

	return
}

// MustManageSignerRuleResult retrieves the ManageSignerRuleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageSignerRuleResult() ManageSignerRuleResult {
	val, ok := u.GetManageSignerRuleResult()

	if !ok {
		panic("arm ManageSignerRuleResult is not set")
	}

	return val
}

// GetManageSignerRuleResult retrieves the ManageSignerRuleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageSignerRuleResult() (result ManageSignerRuleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSignerRuleResult" {
		result = *u.ManageSignerRuleResult
		ok = true
	}

	return
}

// MustCancelChangeRoleRequestResult retrieves the CancelChangeRoleRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCancelChangeRoleRequestResult() CancelChangeRoleRequestResult {
	val, ok := u.GetCancelChangeRoleRequestResult()

	if !ok {
		panic("arm CancelChangeRoleRequestResult is not set")
	}

	return val
}

// GetCancelChangeRoleRequestResult retrieves the CancelChangeRoleRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCancelChangeRoleRequestResult() (result CancelChangeRoleRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CancelChangeRoleRequestResult" {
		result = *u.CancelChangeRoleRequestResult
		ok = true
	}

	return
}

// MustCreateKycRecoveryRequestResult retrieves the CreateKycRecoveryRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateKycRecoveryRequestResult() CreateKycRecoveryRequestResult {
	val, ok := u.GetCreateKycRecoveryRequestResult()

	if !ok {
		panic("arm CreateKycRecoveryRequestResult is not set")
	}

	return val
}

// GetCreateKycRecoveryRequestResult retrieves the CreateKycRecoveryRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateKycRecoveryRequestResult() (result CreateKycRecoveryRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateKycRecoveryRequestResult" {
		result = *u.CreateKycRecoveryRequestResult
		ok = true
	}

	return
}

// MustInitiateKycRecoveryResult retrieves the InitiateKycRecoveryResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustInitiateKycRecoveryResult() InitiateKycRecoveryResult {
	val, ok := u.GetInitiateKycRecoveryResult()

	if !ok {
		panic("arm InitiateKycRecoveryResult is not set")
	}

	return val
}

// GetInitiateKycRecoveryResult retrieves the InitiateKycRecoveryResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetInitiateKycRecoveryResult() (result InitiateKycRecoveryResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "InitiateKycRecoveryResult" {
		result = *u.InitiateKycRecoveryResult
		ok = true
	}

	return
}

// OperationResult is an XDR Union defines as:
//
//   union OperationResult switch (OperationResultCode code)
//    {
//    case opINNER:
//        union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountResult createAccountResult;
//    	case REVIEW_REQUEST:
//    		ReviewRequestResult reviewRequestResult;
//    	case MANAGE_KEY_VALUE:
//    	    ManageKeyValueResult manageKeyValueResult;
//    	case CREATE_CHANGE_ROLE_REQUEST:
//    	    CreateChangeRoleRequestResult createChangeRoleRequestResult;
//        case MANAGE_ACCOUNT_ROLE:
//            ManageAccountRoleResult manageAccountRoleResult;
//        case MANAGE_ACCOUNT_RULE:
//            ManageAccountRuleResult manageAccountRuleResult;
//        case MANAGE_SIGNER:
//            ManageSignerResult manageSignerResult;
//        case MANAGE_SIGNER_ROLE:
//            ManageSignerRoleResult manageSignerRoleResult;
//        case MANAGE_SIGNER_RULE:
//            ManageSignerRuleResult manageSignerRuleResult;
//        case CANCEL_CHANGE_ROLE_REQUEST:
//            CancelChangeRoleRequestResult cancelChangeRoleRequestResult;
//        case CREATE_KYC_RECOVERY_REQUEST:
//            CreateKYCRecoveryRequestResult createKYCRecoveryRequestResult;
//        case INITIATE_KYC_RECOVERY:
//            InitiateKYCRecoveryResult initiateKYCRecoveryResult;
//        }
//        tr;
//    case opNO_ENTRY:
//        LedgerEntryType entryType;
//    case opNO_ROLE_PERMISSION:
//        AccountRuleRequirement requirement;
//    default:
//        void;
//    };
//
type OperationResult struct {
	Code        OperationResultCode     `json:"code,omitempty"`
	Tr          *OperationResultTr      `json:"tr,omitempty"`
	EntryType   *LedgerEntryType        `json:"entryType,omitempty"`
	Requirement *AccountRuleRequirement `json:"requirement,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationResult
func (u OperationResult) ArmForSwitch(sw int32) (string, bool) {
	switch OperationResultCode(sw) {
	case OperationResultCodeOpInner:
		return "Tr", true
	case OperationResultCodeOpNoEntry:
		return "EntryType", true
	case OperationResultCodeOpNoRolePermission:
		return "Requirement", true
	default:
		return "", true
	}
}

// NewOperationResult creates a new  OperationResult.
func NewOperationResult(code OperationResultCode, value interface{}) (result OperationResult, err error) {
	result.Code = code
	switch OperationResultCode(code) {
	case OperationResultCodeOpInner:
		tv, ok := value.(OperationResultTr)
		if !ok {
			err = fmt.Errorf("invalid value, must be OperationResultTr")
			return
		}
		result.Tr = &tv
	case OperationResultCodeOpNoEntry:
		tv, ok := value.(LedgerEntryType)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntryType")
			return
		}
		result.EntryType = &tv
	case OperationResultCodeOpNoRolePermission:
		tv, ok := value.(AccountRuleRequirement)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountRuleRequirement")
			return
		}
		result.Requirement = &tv
	default:
		// void
	}
	return
}

// MustTr retrieves the Tr value from the union,
// panicing if the value is not set.
func (u OperationResult) MustTr() OperationResultTr {
	val, ok := u.GetTr()

	if !ok {
		panic("arm Tr is not set")
	}

	return val
}

// GetTr retrieves the Tr value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResult) GetTr() (result OperationResultTr, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Tr" {
		result = *u.Tr
		ok = true
	}

	return
}

// MustEntryType retrieves the EntryType value from the union,
// panicing if the value is not set.
func (u OperationResult) MustEntryType() LedgerEntryType {
	val, ok := u.GetEntryType()

	if !ok {
		panic("arm EntryType is not set")
	}

	return val
}

// GetEntryType retrieves the EntryType value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResult) GetEntryType() (result LedgerEntryType, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "EntryType" {
		result = *u.EntryType
		ok = true
	}

	return
}

// MustRequirement retrieves the Requirement value from the union,
// panicing if the value is not set.
func (u OperationResult) MustRequirement() AccountRuleRequirement {
	val, ok := u.GetRequirement()

	if !ok {
		panic("arm Requirement is not set")
	}

	return val
}

// GetRequirement retrieves the Requirement value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResult) GetRequirement() (result AccountRuleRequirement, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Requirement" {
		result = *u.Requirement
		ok = true
	}

	return
}

// TransactionResultCode is an XDR Enum defines as:
//
//   enum TransactionResultCode
//    {
//        txSUCCESS = 0, // all operations succeeded
//
//        txFAILED = -1, // one of the operations failed (none were applied)
//
//        txTOO_EARLY = -2,         // ledger closeTime before minTime
//        txTOO_LATE = -3,          // ledger closeTime after maxTime
//        txMISSING_OPERATION = -4, // no operation was specified
//
//        txBAD_AUTH = -5,                   // too few valid signatures / wrong network
//        txNO_ACCOUNT = -6,                 // source account not found
//        txBAD_AUTH_EXTRA = -7,             // unused signatures attached to transaction
//        txINTERNAL_ERROR = -8,             // an unknown error occurred
//        txACCOUNT_BLOCKED = -9,            // account is blocked and cannot be source of tx
//        txDUPLICATION = -10,               // if timing is stored
//        txNO_ROLE_PERMISSION = -11         // account role has not rule that allows send transaction
//    };
//
type TransactionResultCode int32

const (
	TransactionResultCodeTxSuccess          TransactionResultCode = 0
	TransactionResultCodeTxFailed           TransactionResultCode = -1
	TransactionResultCodeTxTooEarly         TransactionResultCode = -2
	TransactionResultCodeTxTooLate          TransactionResultCode = -3
	TransactionResultCodeTxMissingOperation TransactionResultCode = -4
	TransactionResultCodeTxBadAuth          TransactionResultCode = -5
	TransactionResultCodeTxNoAccount        TransactionResultCode = -6
	TransactionResultCodeTxBadAuthExtra     TransactionResultCode = -7
	TransactionResultCodeTxInternalError    TransactionResultCode = -8
	TransactionResultCodeTxAccountBlocked   TransactionResultCode = -9
	TransactionResultCodeTxDuplication      TransactionResultCode = -10
	TransactionResultCodeTxNoRolePermission TransactionResultCode = -11
)

var TransactionResultCodeAll = []TransactionResultCode{
	TransactionResultCodeTxSuccess,
	TransactionResultCodeTxFailed,
	TransactionResultCodeTxTooEarly,
	TransactionResultCodeTxTooLate,
	TransactionResultCodeTxMissingOperation,
	TransactionResultCodeTxBadAuth,
	TransactionResultCodeTxNoAccount,
	TransactionResultCodeTxBadAuthExtra,
	TransactionResultCodeTxInternalError,
	TransactionResultCodeTxAccountBlocked,
	TransactionResultCodeTxDuplication,
	TransactionResultCodeTxNoRolePermission,
}

var transactionResultCodeMap = map[int32]string{
	0:   "TransactionResultCodeTxSuccess",
	-1:  "TransactionResultCodeTxFailed",
	-2:  "TransactionResultCodeTxTooEarly",
	-3:  "TransactionResultCodeTxTooLate",
	-4:  "TransactionResultCodeTxMissingOperation",
	-5:  "TransactionResultCodeTxBadAuth",
	-6:  "TransactionResultCodeTxNoAccount",
	-7:  "TransactionResultCodeTxBadAuthExtra",
	-8:  "TransactionResultCodeTxInternalError",
	-9:  "TransactionResultCodeTxAccountBlocked",
	-10: "TransactionResultCodeTxDuplication",
	-11: "TransactionResultCodeTxNoRolePermission",
}

var transactionResultCodeShortMap = map[int32]string{
	0:   "tx_success",
	-1:  "tx_failed",
	-2:  "tx_too_early",
	-3:  "tx_too_late",
	-4:  "tx_missing_operation",
	-5:  "tx_bad_auth",
	-6:  "tx_no_account",
	-7:  "tx_bad_auth_extra",
	-8:  "tx_internal_error",
	-9:  "tx_account_blocked",
	-10: "tx_duplication",
	-11: "tx_no_role_permission",
}

var transactionResultCodeRevMap = map[string]int32{
	"TransactionResultCodeTxSuccess":          0,
	"TransactionResultCodeTxFailed":           -1,
	"TransactionResultCodeTxTooEarly":         -2,
	"TransactionResultCodeTxTooLate":          -3,
	"TransactionResultCodeTxMissingOperation": -4,
	"TransactionResultCodeTxBadAuth":          -5,
	"TransactionResultCodeTxNoAccount":        -6,
	"TransactionResultCodeTxBadAuthExtra":     -7,
	"TransactionResultCodeTxInternalError":    -8,
	"TransactionResultCodeTxAccountBlocked":   -9,
	"TransactionResultCodeTxDuplication":      -10,
	"TransactionResultCodeTxNoRolePermission": -11,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TransactionResultCode
func (e TransactionResultCode) ValidEnum(v int32) bool {
	_, ok := transactionResultCodeMap[v]
	return ok
}
func (e TransactionResultCode) isFlag() bool {
	for i := len(TransactionResultCodeAll) - 1; i >= 0; i-- {
		expected := TransactionResultCode(2) << uint64(len(TransactionResultCodeAll)-1) >> uint64(len(TransactionResultCodeAll)-i)
		if expected != TransactionResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e TransactionResultCode) String() string {
	name, _ := transactionResultCodeMap[int32(e)]
	return name
}

func (e TransactionResultCode) ShortString() string {
	name, _ := transactionResultCodeShortMap[int32(e)]
	return name
}

func (e TransactionResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range TransactionResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *TransactionResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = TransactionResultCode(t.Value)
	return nil
}

// TransactionResultResult is an XDR NestedUnion defines as:
//
//   union switch (TransactionResultCode code)
//        {
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        case txNO_ROLE_PERMISSION:
//            AccountRuleRequirement requirement;
//        default:
//            void;
//        }
//
type TransactionResultResult struct {
	Code        TransactionResultCode   `json:"code,omitempty"`
	Results     *[]OperationResult      `json:"results,omitempty"`
	Requirement *AccountRuleRequirement `json:"requirement,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultResult
func (u TransactionResultResult) ArmForSwitch(sw int32) (string, bool) {
	switch TransactionResultCode(sw) {
	case TransactionResultCodeTxSuccess:
		return "Results", true
	case TransactionResultCodeTxFailed:
		return "Results", true
	case TransactionResultCodeTxNoRolePermission:
		return "Requirement", true
	default:
		return "", true
	}
}

// NewTransactionResultResult creates a new  TransactionResultResult.
func NewTransactionResultResult(code TransactionResultCode, value interface{}) (result TransactionResultResult, err error) {
	result.Code = code
	switch TransactionResultCode(code) {
	case TransactionResultCodeTxSuccess:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	case TransactionResultCodeTxFailed:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	case TransactionResultCodeTxNoRolePermission:
		tv, ok := value.(AccountRuleRequirement)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountRuleRequirement")
			return
		}
		result.Requirement = &tv
	default:
		// void
	}
	return
}

// MustResults retrieves the Results value from the union,
// panicing if the value is not set.
func (u TransactionResultResult) MustResults() []OperationResult {
	val, ok := u.GetResults()

	if !ok {
		panic("arm Results is not set")
	}

	return val
}

// GetResults retrieves the Results value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionResultResult) GetResults() (result []OperationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Results" {
		result = *u.Results
		ok = true
	}

	return
}

// MustRequirement retrieves the Requirement value from the union,
// panicing if the value is not set.
func (u TransactionResultResult) MustRequirement() AccountRuleRequirement {
	val, ok := u.GetRequirement()

	if !ok {
		panic("arm Requirement is not set")
	}

	return val
}

// GetRequirement retrieves the Requirement value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionResultResult) GetRequirement() (result AccountRuleRequirement, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Requirement" {
		result = *u.Requirement
		ok = true
	}

	return
}

// TransactionResult is an XDR Struct defines as:
//
//   struct TransactionResult
//    {
//        union switch (TransactionResultCode code)
//        {
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        case txNO_ROLE_PERMISSION:
//            AccountRuleRequirement requirement;
//        default:
//            void;
//        }
//        result;
//
//        //: reserved for future use
//        EmptyExt ext;
//    };
//
type TransactionResult struct {
	Result TransactionResultResult `json:"result,omitempty"`
	Ext    EmptyExt                `json:"ext,omitempty"`
}

// LedgerVersion is an XDR Enum defines as:
//
//   enum LedgerVersion
//    {
//        EMPTY_VERSION = 0
//    };
//
type LedgerVersion int32

const (
	LedgerVersionEmptyVersion LedgerVersion = 0
)

var LedgerVersionAll = []LedgerVersion{
	LedgerVersionEmptyVersion,
}

var ledgerVersionMap = map[int32]string{
	0: "LedgerVersionEmptyVersion",
}

var ledgerVersionShortMap = map[int32]string{
	0: "empty_version",
}

var ledgerVersionRevMap = map[string]int32{
	"LedgerVersionEmptyVersion": 0,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerVersion
func (e LedgerVersion) ValidEnum(v int32) bool {
	_, ok := ledgerVersionMap[v]
	return ok
}
func (e LedgerVersion) isFlag() bool {
	for i := len(LedgerVersionAll) - 1; i >= 0; i-- {
		expected := LedgerVersion(2) << uint64(len(LedgerVersionAll)-1) >> uint64(len(LedgerVersionAll)-i)
		if expected != LedgerVersionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerVersion) String() string {
	name, _ := ledgerVersionMap[int32(e)]
	return name
}

func (e LedgerVersion) ShortString() string {
	name, _ := ledgerVersionShortMap[int32(e)]
	return name
}

func (e LedgerVersion) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range LedgerVersionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerVersion) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerVersion(t.Value)
	return nil
}

// EmptyExt is an XDR Union defines as:
//
//   union EmptyExt switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        void;
//    };
//
type EmptyExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u EmptyExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of EmptyExt
func (u EmptyExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewEmptyExt creates a new  EmptyExt.
func NewEmptyExt(v LedgerVersion, value interface{}) (result EmptyExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Hash is an XDR Typedef defines as:
//
//   typedef opaque Hash[32];
//
type Hash [32]byte

// Uint256 is an XDR Typedef defines as:
//
//   typedef opaque uint256[32];
//
type Uint256 [32]byte

// Uint32 is an XDR Typedef defines as:
//
//   typedef unsigned int uint32;
//
type Uint32 uint32

// Int32 is an XDR Typedef defines as:
//
//   typedef int int32;
//
type Int32 int32

// Uint64 is an XDR Typedef defines as:
//
//   typedef unsigned hyper uint64;
//
type Uint64 uint64

// Int64 is an XDR Typedef defines as:
//
//   typedef hyper int64;
//
type Int64 int64

// CryptoKeyType is an XDR Enum defines as:
//
//   enum CryptoKeyType
//    {
//        KEY_TYPE_ED25519 = 0
//    };
//
type CryptoKeyType int32

const (
	CryptoKeyTypeKeyTypeEd25519 CryptoKeyType = 0
)

var CryptoKeyTypeAll = []CryptoKeyType{
	CryptoKeyTypeKeyTypeEd25519,
}

var cryptoKeyTypeMap = map[int32]string{
	0: "CryptoKeyTypeKeyTypeEd25519",
}

var cryptoKeyTypeShortMap = map[int32]string{
	0: "key_type_ed25519",
}

var cryptoKeyTypeRevMap = map[string]int32{
	"CryptoKeyTypeKeyTypeEd25519": 0,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CryptoKeyType
func (e CryptoKeyType) ValidEnum(v int32) bool {
	_, ok := cryptoKeyTypeMap[v]
	return ok
}
func (e CryptoKeyType) isFlag() bool {
	for i := len(CryptoKeyTypeAll) - 1; i >= 0; i-- {
		expected := CryptoKeyType(2) << uint64(len(CryptoKeyTypeAll)-1) >> uint64(len(CryptoKeyTypeAll)-i)
		if expected != CryptoKeyTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CryptoKeyType) String() string {
	name, _ := cryptoKeyTypeMap[int32(e)]
	return name
}

func (e CryptoKeyType) ShortString() string {
	name, _ := cryptoKeyTypeShortMap[int32(e)]
	return name
}

func (e CryptoKeyType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CryptoKeyTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CryptoKeyType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CryptoKeyType(t.Value)
	return nil
}

// PublicKeyType is an XDR Enum defines as:
//
//   enum PublicKeyType
//    {
//    	PUBLIC_KEY_TYPE_ED25519 = 0
//    };
//
type PublicKeyType int32

const (
	PublicKeyTypePublicKeyTypeEd25519 PublicKeyType = 0
)

var PublicKeyTypeAll = []PublicKeyType{
	PublicKeyTypePublicKeyTypeEd25519,
}

var publicKeyTypeMap = map[int32]string{
	0: "PublicKeyTypePublicKeyTypeEd25519",
}

var publicKeyTypeShortMap = map[int32]string{
	0: "public_key_type_ed25519",
}

var publicKeyTypeRevMap = map[string]int32{
	"PublicKeyTypePublicKeyTypeEd25519": 0,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PublicKeyType
func (e PublicKeyType) ValidEnum(v int32) bool {
	_, ok := publicKeyTypeMap[v]
	return ok
}
func (e PublicKeyType) isFlag() bool {
	for i := len(PublicKeyTypeAll) - 1; i >= 0; i-- {
		expected := PublicKeyType(2) << uint64(len(PublicKeyTypeAll)-1) >> uint64(len(PublicKeyTypeAll)-i)
		if expected != PublicKeyTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PublicKeyType) String() string {
	name, _ := publicKeyTypeMap[int32(e)]
	return name
}

func (e PublicKeyType) ShortString() string {
	name, _ := publicKeyTypeShortMap[int32(e)]
	return name
}

func (e PublicKeyType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range PublicKeyTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *PublicKeyType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PublicKeyType(t.Value)
	return nil
}

// PublicKey is an XDR Union defines as:
//
//   union PublicKey switch (CryptoKeyType type)
//    {
//    case KEY_TYPE_ED25519:
//        uint256 ed25519;
//    };
//
type PublicKey struct {
	Type    CryptoKeyType `json:"type,omitempty"`
	Ed25519 *Uint256      `json:"ed25519,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PublicKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u PublicKey) ArmForSwitch(sw int32) (string, bool) {
	switch CryptoKeyType(sw) {
	case CryptoKeyTypeKeyTypeEd25519:
		return "Ed25519", true
	}
	return "-", false
}

// NewPublicKey creates a new  PublicKey.
func NewPublicKey(aType CryptoKeyType, value interface{}) (result PublicKey, err error) {
	result.Type = aType
	switch CryptoKeyType(aType) {
	case CryptoKeyTypeKeyTypeEd25519:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.Ed25519 = &tv
	}
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u PublicKey) MustEd25519() Uint256 {
	val, ok := u.GetEd25519()

	if !ok {
		panic("arm Ed25519 is not set")
	}

	return val
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PublicKey) GetEd25519() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ed25519" {
		result = *u.Ed25519
		ok = true
	}

	return
}

// LedgerEntryType is an XDR Enum defines as:
//
//   enum LedgerEntryType
//    {
//        ANY = 1,
//        ACCOUNT = 2,
//        SIGNER = 3,
//        REFERENCE_ENTRY = 4,
//    	REVIEWABLE_REQUEST = 5,
//    	ACCOUNT_KYC = 6,
//        KEY_VALUE = 7,
//        ACCOUNT_ROLE = 8,
//        ACCOUNT_RULE = 9,
//        TRANSACTION = 10, // is used for account rule resource
//        SIGNER_RULE = 11,
//        SIGNER_ROLE = 12,
//        INITIATE_KYC_RECOVERY = 13
//    };
//
type LedgerEntryType int32

const (
	LedgerEntryTypeAny                 LedgerEntryType = 1
	LedgerEntryTypeAccount             LedgerEntryType = 2
	LedgerEntryTypeSigner              LedgerEntryType = 3
	LedgerEntryTypeReferenceEntry      LedgerEntryType = 4
	LedgerEntryTypeReviewableRequest   LedgerEntryType = 5
	LedgerEntryTypeAccountKyc          LedgerEntryType = 6
	LedgerEntryTypeKeyValue            LedgerEntryType = 7
	LedgerEntryTypeAccountRole         LedgerEntryType = 8
	LedgerEntryTypeAccountRule         LedgerEntryType = 9
	LedgerEntryTypeTransaction         LedgerEntryType = 10
	LedgerEntryTypeSignerRule          LedgerEntryType = 11
	LedgerEntryTypeSignerRole          LedgerEntryType = 12
	LedgerEntryTypeInitiateKycRecovery LedgerEntryType = 13
)

var LedgerEntryTypeAll = []LedgerEntryType{
	LedgerEntryTypeAny,
	LedgerEntryTypeAccount,
	LedgerEntryTypeSigner,
	LedgerEntryTypeReferenceEntry,
	LedgerEntryTypeReviewableRequest,
	LedgerEntryTypeAccountKyc,
	LedgerEntryTypeKeyValue,
	LedgerEntryTypeAccountRole,
	LedgerEntryTypeAccountRule,
	LedgerEntryTypeTransaction,
	LedgerEntryTypeSignerRule,
	LedgerEntryTypeSignerRole,
	LedgerEntryTypeInitiateKycRecovery,
}

var ledgerEntryTypeMap = map[int32]string{
	1:  "LedgerEntryTypeAny",
	2:  "LedgerEntryTypeAccount",
	3:  "LedgerEntryTypeSigner",
	4:  "LedgerEntryTypeReferenceEntry",
	5:  "LedgerEntryTypeReviewableRequest",
	6:  "LedgerEntryTypeAccountKyc",
	7:  "LedgerEntryTypeKeyValue",
	8:  "LedgerEntryTypeAccountRole",
	9:  "LedgerEntryTypeAccountRule",
	10: "LedgerEntryTypeTransaction",
	11: "LedgerEntryTypeSignerRule",
	12: "LedgerEntryTypeSignerRole",
	13: "LedgerEntryTypeInitiateKycRecovery",
}

var ledgerEntryTypeShortMap = map[int32]string{
	1:  "any",
	2:  "account",
	3:  "signer",
	4:  "reference_entry",
	5:  "reviewable_request",
	6:  "account_kyc",
	7:  "key_value",
	8:  "account_role",
	9:  "account_rule",
	10: "transaction",
	11: "signer_rule",
	12: "signer_role",
	13: "initiate_kyc_recovery",
}

var ledgerEntryTypeRevMap = map[string]int32{
	"LedgerEntryTypeAny":                 1,
	"LedgerEntryTypeAccount":             2,
	"LedgerEntryTypeSigner":              3,
	"LedgerEntryTypeReferenceEntry":      4,
	"LedgerEntryTypeReviewableRequest":   5,
	"LedgerEntryTypeAccountKyc":          6,
	"LedgerEntryTypeKeyValue":            7,
	"LedgerEntryTypeAccountRole":         8,
	"LedgerEntryTypeAccountRule":         9,
	"LedgerEntryTypeTransaction":         10,
	"LedgerEntryTypeSignerRule":          11,
	"LedgerEntryTypeSignerRole":          12,
	"LedgerEntryTypeInitiateKycRecovery": 13,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerEntryType
func (e LedgerEntryType) ValidEnum(v int32) bool {
	_, ok := ledgerEntryTypeMap[v]
	return ok
}
func (e LedgerEntryType) isFlag() bool {
	for i := len(LedgerEntryTypeAll) - 1; i >= 0; i-- {
		expected := LedgerEntryType(2) << uint64(len(LedgerEntryTypeAll)-1) >> uint64(len(LedgerEntryTypeAll)-i)
		if expected != LedgerEntryTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerEntryType) String() string {
	name, _ := ledgerEntryTypeMap[int32(e)]
	return name
}

func (e LedgerEntryType) ShortString() string {
	name, _ := ledgerEntryTypeShortMap[int32(e)]
	return name
}

func (e LedgerEntryType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range LedgerEntryTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerEntryType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerEntryType(t.Value)
	return nil
}

// Signature is an XDR Typedef defines as:
//
//   typedef opaque Signature<64>;
//
type Signature []byte

// SignatureHint is an XDR Typedef defines as:
//
//   typedef opaque SignatureHint[4];
//
type SignatureHint [4]byte

// NodeId is an XDR Typedef defines as:
//
//   typedef PublicKey NodeID;
//
type NodeId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u NodeId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u NodeId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewNodeId creates a new  NodeId.
func NewNodeId(aType CryptoKeyType, value interface{}) (result NodeId, err error) {
	u, err := NewPublicKey(aType, value)
	result = NodeId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u NodeId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u NodeId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

// Curve25519Secret is an XDR Struct defines as:
//
//   struct Curve25519Secret
//    {
//            opaque key[32];
//    };
//
type Curve25519Secret struct {
	Key [32]byte `json:"key,omitempty"`
}

// Curve25519Public is an XDR Struct defines as:
//
//   struct Curve25519Public
//    {
//            opaque key[32];
//    };
//
type Curve25519Public struct {
	Key [32]byte `json:"key,omitempty"`
}

// HmacSha256Key is an XDR Struct defines as:
//
//   struct HmacSha256Key
//    {
//            opaque key[32];
//    };
//
type HmacSha256Key struct {
	Key [32]byte `json:"key,omitempty"`
}

// HmacSha256Mac is an XDR Struct defines as:
//
//   struct HmacSha256Mac
//    {
//            opaque mac[32];
//    };
//
type HmacSha256Mac struct {
	Mac [32]byte `json:"mac,omitempty"`
}

// AccountId is an XDR Typedef defines as:
//
//   typedef PublicKey AccountID;
//
type AccountId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u AccountId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewAccountId creates a new  AccountId.
func NewAccountId(aType CryptoKeyType, value interface{}) (result AccountId, err error) {
	u, err := NewPublicKey(aType, value)
	result = AccountId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u AccountId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

// Thresholds is an XDR Typedef defines as:
//
//   typedef opaque Thresholds[4];
//
type Thresholds [4]byte

// String32 is an XDR Typedef defines as:
//
//   typedef string string32<32>;
//
type String32 string

// XDRMaxSize implements the Sized interface for String32
func (e String32) XDRMaxSize() int {
	return 32
}

// String64 is an XDR Typedef defines as:
//
//   typedef string string64<64>;
//
type String64 string

// XDRMaxSize implements the Sized interface for String64
func (e String64) XDRMaxSize() int {
	return 64
}

// String256 is an XDR Typedef defines as:
//
//   typedef string string256<256>;
//
type String256 string

// XDRMaxSize implements the Sized interface for String256
func (e String256) XDRMaxSize() int {
	return 256
}

// Longstring is an XDR Typedef defines as:
//
//   typedef string longstring<>;
//
type Longstring string

// Salt is an XDR Typedef defines as:
//
//   typedef uint64 Salt;
//
type Salt Uint64

// DataValue is an XDR Typedef defines as:
//
//   typedef opaque DataValue<64>;
//
type DataValue []byte

// OperationType is an XDR Enum defines as:
//
//   enum OperationType
//    {
//        CREATE_ACCOUNT = 1,
//    	REVIEW_REQUEST = 2,
//        CREATE_CHANGE_ROLE_REQUEST = 3,
//        CANCEL_CHANGE_ROLE_REQUEST = 4,
//        MANAGE_KEY_VALUE = 5,
//        MANAGE_ACCOUNT_ROLE = 6,
//        MANAGE_ACCOUNT_RULE = 7,
//        MANAGE_SIGNER = 8,
//        MANAGE_SIGNER_ROLE = 9,
//        MANAGE_SIGNER_RULE = 10,
//        INITIATE_KYC_RECOVERY = 11,
//        CREATE_KYC_RECOVERY_REQUEST = 12
//    };
//
type OperationType int32

const (
	OperationTypeCreateAccount            OperationType = 1
	OperationTypeReviewRequest            OperationType = 2
	OperationTypeCreateChangeRoleRequest  OperationType = 3
	OperationTypeCancelChangeRoleRequest  OperationType = 4
	OperationTypeManageKeyValue           OperationType = 5
	OperationTypeManageAccountRole        OperationType = 6
	OperationTypeManageAccountRule        OperationType = 7
	OperationTypeManageSigner             OperationType = 8
	OperationTypeManageSignerRole         OperationType = 9
	OperationTypeManageSignerRule         OperationType = 10
	OperationTypeInitiateKycRecovery      OperationType = 11
	OperationTypeCreateKycRecoveryRequest OperationType = 12
)

var OperationTypeAll = []OperationType{
	OperationTypeCreateAccount,
	OperationTypeReviewRequest,
	OperationTypeCreateChangeRoleRequest,
	OperationTypeCancelChangeRoleRequest,
	OperationTypeManageKeyValue,
	OperationTypeManageAccountRole,
	OperationTypeManageAccountRule,
	OperationTypeManageSigner,
	OperationTypeManageSignerRole,
	OperationTypeManageSignerRule,
	OperationTypeInitiateKycRecovery,
	OperationTypeCreateKycRecoveryRequest,
}

var operationTypeMap = map[int32]string{
	1:  "OperationTypeCreateAccount",
	2:  "OperationTypeReviewRequest",
	3:  "OperationTypeCreateChangeRoleRequest",
	4:  "OperationTypeCancelChangeRoleRequest",
	5:  "OperationTypeManageKeyValue",
	6:  "OperationTypeManageAccountRole",
	7:  "OperationTypeManageAccountRule",
	8:  "OperationTypeManageSigner",
	9:  "OperationTypeManageSignerRole",
	10: "OperationTypeManageSignerRule",
	11: "OperationTypeInitiateKycRecovery",
	12: "OperationTypeCreateKycRecoveryRequest",
}

var operationTypeShortMap = map[int32]string{
	1:  "create_account",
	2:  "review_request",
	3:  "create_change_role_request",
	4:  "cancel_change_role_request",
	5:  "manage_key_value",
	6:  "manage_account_role",
	7:  "manage_account_rule",
	8:  "manage_signer",
	9:  "manage_signer_role",
	10: "manage_signer_rule",
	11: "initiate_kyc_recovery",
	12: "create_kyc_recovery_request",
}

var operationTypeRevMap = map[string]int32{
	"OperationTypeCreateAccount":            1,
	"OperationTypeReviewRequest":            2,
	"OperationTypeCreateChangeRoleRequest":  3,
	"OperationTypeCancelChangeRoleRequest":  4,
	"OperationTypeManageKeyValue":           5,
	"OperationTypeManageAccountRole":        6,
	"OperationTypeManageAccountRule":        7,
	"OperationTypeManageSigner":             8,
	"OperationTypeManageSignerRole":         9,
	"OperationTypeManageSignerRule":         10,
	"OperationTypeInitiateKycRecovery":      11,
	"OperationTypeCreateKycRecoveryRequest": 12,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OperationType
func (e OperationType) ValidEnum(v int32) bool {
	_, ok := operationTypeMap[v]
	return ok
}
func (e OperationType) isFlag() bool {
	for i := len(OperationTypeAll) - 1; i >= 0; i-- {
		expected := OperationType(2) << uint64(len(OperationTypeAll)-1) >> uint64(len(OperationTypeAll)-i)
		if expected != OperationTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e OperationType) String() string {
	name, _ := operationTypeMap[int32(e)]
	return name
}

func (e OperationType) ShortString() string {
	name, _ := operationTypeShortMap[int32(e)]
	return name
}

func (e OperationType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range OperationTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *OperationType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = OperationType(t.Value)
	return nil
}

// DecoratedSignature is an XDR Struct defines as:
//
//   struct DecoratedSignature
//    {
//        SignatureHint hint;  // last 4 bytes of the public key, used as a hint
//        Signature signature; // actual signature
//    };
//
type DecoratedSignature struct {
	Hint      SignatureHint `json:"hint,omitempty"`
	Signature Signature     `json:"signature,omitempty"`
}

var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
var Revision = "e98bc93279bf4e9fbb9107380be2fb8c799236b5"
