// revision: 2efa51c931dc71565fcedc696e425d0bd065599a
// branch:   refactor/internal_work
// Package xdr is generated from:
//
//  xdr/ledger-entries-account-KYC.x
//  xdr/ledger-entries-account.x
//  xdr/ledger-entries-asset.x
//  xdr/ledger-entries-balance.x
//  xdr/ledger-entries-data.x
//  xdr/ledger-entries-key-value.x
//  xdr/ledger-entries-reference.x
//  xdr/ledger-entries-reviewable-request.x
//  xdr/ledger-entries-role.x
//  xdr/ledger-entries-rule.x
//  xdr/ledger-entries-signer.x
//  xdr/ledger-entries.x
//  xdr/ledger-keys.x
//  xdr/ledger.x
//  xdr/operation-change-account-roles.x
//  xdr/operation-create-account.x
//  xdr/operation-create-asset.x
//  xdr/operation-create-balance.x
//  xdr/operation-create-data.x
//  xdr/operation-create-reviewable-request.x
//  xdr/operation-create-role.x
//  xdr/operation-create-rule.x
//  xdr/operation-create-signer.x
//  xdr/operation-destruction.x
//  xdr/operation-initiate-kyc-recovery.x
//  xdr/operation-issuance.x
//  xdr/operation-kyc-recovery.x
//  xdr/operation-payment.x
//  xdr/operation-put-key-value.x
//  xdr/operation-remove-data.x
//  xdr/operation-remove-key-value.x
//  xdr/operation-remove-reviewable-request.x
//  xdr/operation-remove-role.x
//  xdr/operation-remove-rule.x
//  xdr/operation-remove-signer.x
//  xdr/operation-review-request.x
//  xdr/operation-update-asset.x
//  xdr/operation-update-data.x
//  xdr/operation-update-reviewable-request.x
//  xdr/operation-update-role.x
//  xdr/operation-update-rule.x
//  xdr/operation-update-signer.x
//  xdr/overlay.x
//  xdr/rule.x
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
//    	uint64 roleIDs<>;
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
	RoleIDs      []Uint64        `json:"roleIDs,omitempty"`
	Ext          AccountEntryExt `json:"ext,omitempty"`
}

// AssetEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AssetEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AssetEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AssetEntryExt
func (u AssetEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAssetEntryExt creates a new  AssetEntryExt.
func NewAssetEntryExt(v LedgerVersion, value interface{}) (result AssetEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AssetEntry is an XDR Struct defines as:
//
//   struct AssetEntry
//    {
//        AssetCode code;
//
//        uint32 securityType; // use instead policies that limit usage, use in account rules
//        uint32 state; // smth that can be used to disable asset
//
//    	uint64 maxIssuanceAmount; // max number of tokens to be issued
//    	uint64 issued; // number of issued tokens
//    	uint64 pendingIssuance; // number of tokens to be issued
//
//        uint32 trailingDigitsCount;
//
//        AccountID owner;
//    	longstring details;
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
type AssetEntry struct {
	Code                AssetCode     `json:"code,omitempty"`
	SecurityType        Uint32        `json:"securityType,omitempty"`
	State               Uint32        `json:"state,omitempty"`
	MaxIssuanceAmount   Uint64        `json:"maxIssuanceAmount,omitempty"`
	Issued              Uint64        `json:"issued,omitempty"`
	PendingIssuance     Uint64        `json:"pendingIssuance,omitempty"`
	TrailingDigitsCount Uint32        `json:"trailingDigitsCount,omitempty"`
	Owner               AccountId     `json:"owner,omitempty"`
	Details             Longstring    `json:"details,omitempty"`
	Ext                 AssetEntryExt `json:"ext,omitempty"`
}

// BalanceEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type BalanceEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BalanceEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BalanceEntryExt
func (u BalanceEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewBalanceEntryExt creates a new  BalanceEntryExt.
func NewBalanceEntryExt(v LedgerVersion, value interface{}) (result BalanceEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// BalanceEntry is an XDR Struct defines as:
//
//   struct BalanceEntry
//    {
//        BalanceID balanceID;
//    	// sequential ID - unique identifier of the balance, used by ingesting applications to
//    	// identify account, while keeping size of index small
//        uint64 sequentialID;
//        AssetCode asset;
//        AccountID accountID;
//        uint64 amount;
//        uint64 locked;
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
type BalanceEntry struct {
	BalanceId    BalanceId       `json:"balanceID,omitempty"`
	SequentialId Uint64          `json:"sequentialID,omitempty"`
	Asset        AssetCode       `json:"asset,omitempty"`
	AccountId    AccountId       `json:"accountID,omitempty"`
	Amount       Uint64          `json:"amount,omitempty"`
	Locked       Uint64          `json:"locked,omitempty"`
	Ext          BalanceEntryExt `json:"ext,omitempty"`
}

// DataEntry is an XDR Struct defines as:
//
//   struct DataEntry
//    {
//        uint64 id;
//        uint32 securityType;
//        longstring value;
//
//        AccountID owner;
//
//        EmptyExt ext;
//    };
//
type DataEntry struct {
	Id           Uint64     `json:"id,omitempty"`
	SecurityType Uint32     `json:"securityType,omitempty"`
	Value        Longstring `json:"value,omitempty"`
	Owner        AccountId  `json:"owner,omitempty"`
	Ext          EmptyExt   `json:"ext,omitempty"`
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

// ReviewableRequestOperation is an XDR Union defines as:
//
//   union ReviewableRequestOperation switch (OperationType type)
//    {
//    case CREATE_ACCOUNT:
//        CreateAccountOp createAccountOp;
//    case PAYMENT:
//        PaymentOp paymentOp;
//    case CREATE_SIGNER:
//        CreateSignerOp createSignerOp;
//    case UPDATE_SIGNER:
//        UpdateSignerOp updateSignerOp;
//    case REMOVE_SIGNER:
//        RemoveSignerOp removeSignerOp;
//    case CREATE_ROLE:
//        CreateRoleOp createRoleOp;
//    case UPDATE_ROLE:
//        UpdateRoleOp updateRoleOp;
//    case REMOVE_ROLE:
//        RemoveRoleOp removeRoleOp;
//    case CREATE_RULE:
//        CreateRuleOp createRuleOp;
//    case UPDATE_RULE:
//        UpdateRuleOp updateRuleOp;
//    case REMOVE_RULE:
//        RemoveRuleOp removeRuleOp;
//    case ISSUANCE:
//        IssuanceOp issuanceOp;
//    case DESTRUCTION:
//        DestructionOp destructionOp;
//    case CHANGE_ACCOUNT_ROLES:
//        ChangeAccountRolesOp changeAccountRolesOp;
//    case CREATE_ASSET:
//        CreateAssetOp createAssetOp;
//    case UPDATE_ASSET:
//        UpdateAssetOp updateAssetOp;
//    case PUT_KEY_VALUE:
//        PutKeyValueOp putKeyValueOp;
//    case REMOVE_KEY_VALUE:
//        RemoveKeyValueOp removeKeyValueOp;
//    case CREATE_DATA:
//        CreateDataOp createDataOp;
//    case UPDATE_DATA:
//        UpdateDataOp updateDataOp;
//    case REMOVE_DATA:
//        RemoveDataOp removeDataOp;
//    case CREATE_BALANCE:
//        CreateBalanceOp createBalanceOp;
//    case INITIATE_KYC_RECOVERY:
//        InitiateKYCRecoveryOp initiateKYCRecoveryOp;
//    case KYC_RECOVERY:
//        KYCRecoveryOp kycRecoveryOp;
//
//    };
//
type ReviewableRequestOperation struct {
	Type                  OperationType          `json:"type,omitempty"`
	CreateAccountOp       *CreateAccountOp       `json:"createAccountOp,omitempty"`
	PaymentOp             *PaymentOp             `json:"paymentOp,omitempty"`
	CreateSignerOp        *CreateSignerOp        `json:"createSignerOp,omitempty"`
	UpdateSignerOp        *UpdateSignerOp        `json:"updateSignerOp,omitempty"`
	RemoveSignerOp        *RemoveSignerOp        `json:"removeSignerOp,omitempty"`
	CreateRoleOp          *CreateRoleOp          `json:"createRoleOp,omitempty"`
	UpdateRoleOp          *UpdateRoleOp          `json:"updateRoleOp,omitempty"`
	RemoveRoleOp          *RemoveRoleOp          `json:"removeRoleOp,omitempty"`
	CreateRuleOp          *CreateRuleOp          `json:"createRuleOp,omitempty"`
	UpdateRuleOp          *UpdateRuleOp          `json:"updateRuleOp,omitempty"`
	RemoveRuleOp          *RemoveRuleOp          `json:"removeRuleOp,omitempty"`
	IssuanceOp            *IssuanceOp            `json:"issuanceOp,omitempty"`
	DestructionOp         *DestructionOp         `json:"destructionOp,omitempty"`
	ChangeAccountRolesOp  *ChangeAccountRolesOp  `json:"changeAccountRolesOp,omitempty"`
	CreateAssetOp         *CreateAssetOp         `json:"createAssetOp,omitempty"`
	UpdateAssetOp         *UpdateAssetOp         `json:"updateAssetOp,omitempty"`
	PutKeyValueOp         *PutKeyValueOp         `json:"putKeyValueOp,omitempty"`
	RemoveKeyValueOp      *RemoveKeyValueOp      `json:"removeKeyValueOp,omitempty"`
	CreateDataOp          *CreateDataOp          `json:"createDataOp,omitempty"`
	UpdateDataOp          *UpdateDataOp          `json:"updateDataOp,omitempty"`
	RemoveDataOp          *RemoveDataOp          `json:"removeDataOp,omitempty"`
	CreateBalanceOp       *CreateBalanceOp       `json:"createBalanceOp,omitempty"`
	InitiateKycRecoveryOp *InitiateKycRecoveryOp `json:"initiateKYCRecoveryOp,omitempty"`
	KycRecoveryOp         *KycRecoveryOp         `json:"kycRecoveryOp,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewableRequestOperation) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewableRequestOperation
func (u ReviewableRequestOperation) ArmForSwitch(sw int32) (string, bool) {
	switch OperationType(sw) {
	case OperationTypeCreateAccount:
		return "CreateAccountOp", true
	case OperationTypePayment:
		return "PaymentOp", true
	case OperationTypeCreateSigner:
		return "CreateSignerOp", true
	case OperationTypeUpdateSigner:
		return "UpdateSignerOp", true
	case OperationTypeRemoveSigner:
		return "RemoveSignerOp", true
	case OperationTypeCreateRole:
		return "CreateRoleOp", true
	case OperationTypeUpdateRole:
		return "UpdateRoleOp", true
	case OperationTypeRemoveRole:
		return "RemoveRoleOp", true
	case OperationTypeCreateRule:
		return "CreateRuleOp", true
	case OperationTypeUpdateRule:
		return "UpdateRuleOp", true
	case OperationTypeRemoveRule:
		return "RemoveRuleOp", true
	case OperationTypeIssuance:
		return "IssuanceOp", true
	case OperationTypeDestruction:
		return "DestructionOp", true
	case OperationTypeChangeAccountRoles:
		return "ChangeAccountRolesOp", true
	case OperationTypeCreateAsset:
		return "CreateAssetOp", true
	case OperationTypeUpdateAsset:
		return "UpdateAssetOp", true
	case OperationTypePutKeyValue:
		return "PutKeyValueOp", true
	case OperationTypeRemoveKeyValue:
		return "RemoveKeyValueOp", true
	case OperationTypeCreateData:
		return "CreateDataOp", true
	case OperationTypeUpdateData:
		return "UpdateDataOp", true
	case OperationTypeRemoveData:
		return "RemoveDataOp", true
	case OperationTypeCreateBalance:
		return "CreateBalanceOp", true
	case OperationTypeInitiateKycRecovery:
		return "InitiateKycRecoveryOp", true
	case OperationTypeKycRecovery:
		return "KycRecoveryOp", true
	}
	return "-", false
}

// NewReviewableRequestOperation creates a new  ReviewableRequestOperation.
func NewReviewableRequestOperation(aType OperationType, value interface{}) (result ReviewableRequestOperation, err error) {
	result.Type = aType
	switch OperationType(aType) {
	case OperationTypeCreateAccount:
		tv, ok := value.(CreateAccountOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountOp")
			return
		}
		result.CreateAccountOp = &tv
	case OperationTypePayment:
		tv, ok := value.(PaymentOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentOp")
			return
		}
		result.PaymentOp = &tv
	case OperationTypeCreateSigner:
		tv, ok := value.(CreateSignerOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateSignerOp")
			return
		}
		result.CreateSignerOp = &tv
	case OperationTypeUpdateSigner:
		tv, ok := value.(UpdateSignerOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateSignerOp")
			return
		}
		result.UpdateSignerOp = &tv
	case OperationTypeRemoveSigner:
		tv, ok := value.(RemoveSignerOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveSignerOp")
			return
		}
		result.RemoveSignerOp = &tv
	case OperationTypeCreateRole:
		tv, ok := value.(CreateRoleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateRoleOp")
			return
		}
		result.CreateRoleOp = &tv
	case OperationTypeUpdateRole:
		tv, ok := value.(UpdateRoleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateRoleOp")
			return
		}
		result.UpdateRoleOp = &tv
	case OperationTypeRemoveRole:
		tv, ok := value.(RemoveRoleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveRoleOp")
			return
		}
		result.RemoveRoleOp = &tv
	case OperationTypeCreateRule:
		tv, ok := value.(CreateRuleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateRuleOp")
			return
		}
		result.CreateRuleOp = &tv
	case OperationTypeUpdateRule:
		tv, ok := value.(UpdateRuleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateRuleOp")
			return
		}
		result.UpdateRuleOp = &tv
	case OperationTypeRemoveRule:
		tv, ok := value.(RemoveRuleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveRuleOp")
			return
		}
		result.RemoveRuleOp = &tv
	case OperationTypeIssuance:
		tv, ok := value.(IssuanceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be IssuanceOp")
			return
		}
		result.IssuanceOp = &tv
	case OperationTypeDestruction:
		tv, ok := value.(DestructionOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be DestructionOp")
			return
		}
		result.DestructionOp = &tv
	case OperationTypeChangeAccountRoles:
		tv, ok := value.(ChangeAccountRolesOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ChangeAccountRolesOp")
			return
		}
		result.ChangeAccountRolesOp = &tv
	case OperationTypeCreateAsset:
		tv, ok := value.(CreateAssetOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAssetOp")
			return
		}
		result.CreateAssetOp = &tv
	case OperationTypeUpdateAsset:
		tv, ok := value.(UpdateAssetOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateAssetOp")
			return
		}
		result.UpdateAssetOp = &tv
	case OperationTypePutKeyValue:
		tv, ok := value.(PutKeyValueOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PutKeyValueOp")
			return
		}
		result.PutKeyValueOp = &tv
	case OperationTypeRemoveKeyValue:
		tv, ok := value.(RemoveKeyValueOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveKeyValueOp")
			return
		}
		result.RemoveKeyValueOp = &tv
	case OperationTypeCreateData:
		tv, ok := value.(CreateDataOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateDataOp")
			return
		}
		result.CreateDataOp = &tv
	case OperationTypeUpdateData:
		tv, ok := value.(UpdateDataOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateDataOp")
			return
		}
		result.UpdateDataOp = &tv
	case OperationTypeRemoveData:
		tv, ok := value.(RemoveDataOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveDataOp")
			return
		}
		result.RemoveDataOp = &tv
	case OperationTypeCreateBalance:
		tv, ok := value.(CreateBalanceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateBalanceOp")
			return
		}
		result.CreateBalanceOp = &tv
	case OperationTypeInitiateKycRecovery:
		tv, ok := value.(InitiateKycRecoveryOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be InitiateKycRecoveryOp")
			return
		}
		result.InitiateKycRecoveryOp = &tv
	case OperationTypeKycRecovery:
		tv, ok := value.(KycRecoveryOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be KycRecoveryOp")
			return
		}
		result.KycRecoveryOp = &tv
	}
	return
}

// MustCreateAccountOp retrieves the CreateAccountOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustCreateAccountOp() CreateAccountOp {
	val, ok := u.GetCreateAccountOp()

	if !ok {
		panic("arm CreateAccountOp is not set")
	}

	return val
}

// GetCreateAccountOp retrieves the CreateAccountOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetCreateAccountOp() (result CreateAccountOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAccountOp" {
		result = *u.CreateAccountOp
		ok = true
	}

	return
}

// MustPaymentOp retrieves the PaymentOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustPaymentOp() PaymentOp {
	val, ok := u.GetPaymentOp()

	if !ok {
		panic("arm PaymentOp is not set")
	}

	return val
}

// GetPaymentOp retrieves the PaymentOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetPaymentOp() (result PaymentOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentOp" {
		result = *u.PaymentOp
		ok = true
	}

	return
}

// MustCreateSignerOp retrieves the CreateSignerOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustCreateSignerOp() CreateSignerOp {
	val, ok := u.GetCreateSignerOp()

	if !ok {
		panic("arm CreateSignerOp is not set")
	}

	return val
}

// GetCreateSignerOp retrieves the CreateSignerOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetCreateSignerOp() (result CreateSignerOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateSignerOp" {
		result = *u.CreateSignerOp
		ok = true
	}

	return
}

// MustUpdateSignerOp retrieves the UpdateSignerOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustUpdateSignerOp() UpdateSignerOp {
	val, ok := u.GetUpdateSignerOp()

	if !ok {
		panic("arm UpdateSignerOp is not set")
	}

	return val
}

// GetUpdateSignerOp retrieves the UpdateSignerOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetUpdateSignerOp() (result UpdateSignerOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateSignerOp" {
		result = *u.UpdateSignerOp
		ok = true
	}

	return
}

// MustRemoveSignerOp retrieves the RemoveSignerOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustRemoveSignerOp() RemoveSignerOp {
	val, ok := u.GetRemoveSignerOp()

	if !ok {
		panic("arm RemoveSignerOp is not set")
	}

	return val
}

// GetRemoveSignerOp retrieves the RemoveSignerOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetRemoveSignerOp() (result RemoveSignerOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveSignerOp" {
		result = *u.RemoveSignerOp
		ok = true
	}

	return
}

// MustCreateRoleOp retrieves the CreateRoleOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustCreateRoleOp() CreateRoleOp {
	val, ok := u.GetCreateRoleOp()

	if !ok {
		panic("arm CreateRoleOp is not set")
	}

	return val
}

// GetCreateRoleOp retrieves the CreateRoleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetCreateRoleOp() (result CreateRoleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateRoleOp" {
		result = *u.CreateRoleOp
		ok = true
	}

	return
}

// MustUpdateRoleOp retrieves the UpdateRoleOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustUpdateRoleOp() UpdateRoleOp {
	val, ok := u.GetUpdateRoleOp()

	if !ok {
		panic("arm UpdateRoleOp is not set")
	}

	return val
}

// GetUpdateRoleOp retrieves the UpdateRoleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetUpdateRoleOp() (result UpdateRoleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateRoleOp" {
		result = *u.UpdateRoleOp
		ok = true
	}

	return
}

// MustRemoveRoleOp retrieves the RemoveRoleOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustRemoveRoleOp() RemoveRoleOp {
	val, ok := u.GetRemoveRoleOp()

	if !ok {
		panic("arm RemoveRoleOp is not set")
	}

	return val
}

// GetRemoveRoleOp retrieves the RemoveRoleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetRemoveRoleOp() (result RemoveRoleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveRoleOp" {
		result = *u.RemoveRoleOp
		ok = true
	}

	return
}

// MustCreateRuleOp retrieves the CreateRuleOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustCreateRuleOp() CreateRuleOp {
	val, ok := u.GetCreateRuleOp()

	if !ok {
		panic("arm CreateRuleOp is not set")
	}

	return val
}

// GetCreateRuleOp retrieves the CreateRuleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetCreateRuleOp() (result CreateRuleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateRuleOp" {
		result = *u.CreateRuleOp
		ok = true
	}

	return
}

// MustUpdateRuleOp retrieves the UpdateRuleOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustUpdateRuleOp() UpdateRuleOp {
	val, ok := u.GetUpdateRuleOp()

	if !ok {
		panic("arm UpdateRuleOp is not set")
	}

	return val
}

// GetUpdateRuleOp retrieves the UpdateRuleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetUpdateRuleOp() (result UpdateRuleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateRuleOp" {
		result = *u.UpdateRuleOp
		ok = true
	}

	return
}

// MustRemoveRuleOp retrieves the RemoveRuleOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustRemoveRuleOp() RemoveRuleOp {
	val, ok := u.GetRemoveRuleOp()

	if !ok {
		panic("arm RemoveRuleOp is not set")
	}

	return val
}

// GetRemoveRuleOp retrieves the RemoveRuleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetRemoveRuleOp() (result RemoveRuleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveRuleOp" {
		result = *u.RemoveRuleOp
		ok = true
	}

	return
}

// MustIssuanceOp retrieves the IssuanceOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustIssuanceOp() IssuanceOp {
	val, ok := u.GetIssuanceOp()

	if !ok {
		panic("arm IssuanceOp is not set")
	}

	return val
}

// GetIssuanceOp retrieves the IssuanceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetIssuanceOp() (result IssuanceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "IssuanceOp" {
		result = *u.IssuanceOp
		ok = true
	}

	return
}

// MustDestructionOp retrieves the DestructionOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustDestructionOp() DestructionOp {
	val, ok := u.GetDestructionOp()

	if !ok {
		panic("arm DestructionOp is not set")
	}

	return val
}

// GetDestructionOp retrieves the DestructionOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetDestructionOp() (result DestructionOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DestructionOp" {
		result = *u.DestructionOp
		ok = true
	}

	return
}

// MustChangeAccountRolesOp retrieves the ChangeAccountRolesOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustChangeAccountRolesOp() ChangeAccountRolesOp {
	val, ok := u.GetChangeAccountRolesOp()

	if !ok {
		panic("arm ChangeAccountRolesOp is not set")
	}

	return val
}

// GetChangeAccountRolesOp retrieves the ChangeAccountRolesOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetChangeAccountRolesOp() (result ChangeAccountRolesOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChangeAccountRolesOp" {
		result = *u.ChangeAccountRolesOp
		ok = true
	}

	return
}

// MustCreateAssetOp retrieves the CreateAssetOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustCreateAssetOp() CreateAssetOp {
	val, ok := u.GetCreateAssetOp()

	if !ok {
		panic("arm CreateAssetOp is not set")
	}

	return val
}

// GetCreateAssetOp retrieves the CreateAssetOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetCreateAssetOp() (result CreateAssetOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAssetOp" {
		result = *u.CreateAssetOp
		ok = true
	}

	return
}

// MustUpdateAssetOp retrieves the UpdateAssetOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustUpdateAssetOp() UpdateAssetOp {
	val, ok := u.GetUpdateAssetOp()

	if !ok {
		panic("arm UpdateAssetOp is not set")
	}

	return val
}

// GetUpdateAssetOp retrieves the UpdateAssetOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetUpdateAssetOp() (result UpdateAssetOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateAssetOp" {
		result = *u.UpdateAssetOp
		ok = true
	}

	return
}

// MustPutKeyValueOp retrieves the PutKeyValueOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustPutKeyValueOp() PutKeyValueOp {
	val, ok := u.GetPutKeyValueOp()

	if !ok {
		panic("arm PutKeyValueOp is not set")
	}

	return val
}

// GetPutKeyValueOp retrieves the PutKeyValueOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetPutKeyValueOp() (result PutKeyValueOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PutKeyValueOp" {
		result = *u.PutKeyValueOp
		ok = true
	}

	return
}

// MustRemoveKeyValueOp retrieves the RemoveKeyValueOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustRemoveKeyValueOp() RemoveKeyValueOp {
	val, ok := u.GetRemoveKeyValueOp()

	if !ok {
		panic("arm RemoveKeyValueOp is not set")
	}

	return val
}

// GetRemoveKeyValueOp retrieves the RemoveKeyValueOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetRemoveKeyValueOp() (result RemoveKeyValueOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveKeyValueOp" {
		result = *u.RemoveKeyValueOp
		ok = true
	}

	return
}

// MustCreateDataOp retrieves the CreateDataOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustCreateDataOp() CreateDataOp {
	val, ok := u.GetCreateDataOp()

	if !ok {
		panic("arm CreateDataOp is not set")
	}

	return val
}

// GetCreateDataOp retrieves the CreateDataOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetCreateDataOp() (result CreateDataOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateDataOp" {
		result = *u.CreateDataOp
		ok = true
	}

	return
}

// MustUpdateDataOp retrieves the UpdateDataOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustUpdateDataOp() UpdateDataOp {
	val, ok := u.GetUpdateDataOp()

	if !ok {
		panic("arm UpdateDataOp is not set")
	}

	return val
}

// GetUpdateDataOp retrieves the UpdateDataOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetUpdateDataOp() (result UpdateDataOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateDataOp" {
		result = *u.UpdateDataOp
		ok = true
	}

	return
}

// MustRemoveDataOp retrieves the RemoveDataOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustRemoveDataOp() RemoveDataOp {
	val, ok := u.GetRemoveDataOp()

	if !ok {
		panic("arm RemoveDataOp is not set")
	}

	return val
}

// GetRemoveDataOp retrieves the RemoveDataOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetRemoveDataOp() (result RemoveDataOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveDataOp" {
		result = *u.RemoveDataOp
		ok = true
	}

	return
}

// MustCreateBalanceOp retrieves the CreateBalanceOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustCreateBalanceOp() CreateBalanceOp {
	val, ok := u.GetCreateBalanceOp()

	if !ok {
		panic("arm CreateBalanceOp is not set")
	}

	return val
}

// GetCreateBalanceOp retrieves the CreateBalanceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetCreateBalanceOp() (result CreateBalanceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateBalanceOp" {
		result = *u.CreateBalanceOp
		ok = true
	}

	return
}

// MustInitiateKycRecoveryOp retrieves the InitiateKycRecoveryOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustInitiateKycRecoveryOp() InitiateKycRecoveryOp {
	val, ok := u.GetInitiateKycRecoveryOp()

	if !ok {
		panic("arm InitiateKycRecoveryOp is not set")
	}

	return val
}

// GetInitiateKycRecoveryOp retrieves the InitiateKycRecoveryOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetInitiateKycRecoveryOp() (result InitiateKycRecoveryOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "InitiateKycRecoveryOp" {
		result = *u.InitiateKycRecoveryOp
		ok = true
	}

	return
}

// MustKycRecoveryOp retrieves the KycRecoveryOp value from the union,
// panicing if the value is not set.
func (u ReviewableRequestOperation) MustKycRecoveryOp() KycRecoveryOp {
	val, ok := u.GetKycRecoveryOp()

	if !ok {
		panic("arm KycRecoveryOp is not set")
	}

	return val
}

// GetKycRecoveryOp retrieves the KycRecoveryOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestOperation) GetKycRecoveryOp() (result KycRecoveryOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KycRecoveryOp" {
		result = *u.KycRecoveryOp
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
//   struct ReviewableRequestEntry
//    {
//    	uint64 requestID;
//    	Hash hash; // hash of the request body
//
//        uint32 securityType; // responsible for operations (types, count)
//
//    	AccountID requestor;
//        longstring rejectReason;
//    	int64 createdAt; // when request was created
//
//    	ReviewableRequestOperation operations<>;
//
//    	uint64 allTasks;
//        uint64 pendingTasks;
//        // maybe add sequenceNumber and creator details
//
//        // External details vector consists of comments written by request reviewers
//        longstring externalDetails<>;
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
	RequestId       Uint64                       `json:"requestID,omitempty"`
	Hash            Hash                         `json:"hash,omitempty"`
	SecurityType    Uint32                       `json:"securityType,omitempty"`
	Requestor       AccountId                    `json:"requestor,omitempty"`
	RejectReason    Longstring                   `json:"rejectReason,omitempty"`
	CreatedAt       Int64                        `json:"createdAt,omitempty"`
	Operations      []ReviewableRequestOperation `json:"operations,omitempty"`
	AllTasks        Uint64                       `json:"allTasks,omitempty"`
	PendingTasks    Uint64                       `json:"pendingTasks,omitempty"`
	ExternalDetails []Longstring                 `json:"externalDetails,omitempty"`
	Ext             ReviewableRequestEntryExt    `json:"ext,omitempty"`
}

// RoleEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type RoleEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RoleEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RoleEntryExt
func (u RoleEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRoleEntryExt creates a new  RoleEntryExt.
func NewRoleEntryExt(v LedgerVersion, value interface{}) (result RoleEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RoleEntry is an XDR Struct defines as:
//
//   struct RoleEntry
//    {
//        uint64 id;
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
type RoleEntry struct {
	Id      Uint64       `json:"id,omitempty"`
	RuleIDs []Uint64     `json:"ruleIDs,omitempty"`
	Details Longstring   `json:"details,omitempty"`
	Ext     RoleEntryExt `json:"ext,omitempty"`
}

// RuleEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type RuleEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RuleEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RuleEntryExt
func (u RuleEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRuleEntryExt creates a new  RuleEntryExt.
func NewRuleEntryExt(v LedgerVersion, value interface{}) (result RuleEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RuleEntry is an XDR Struct defines as:
//
//   struct RuleEntry
//    {
//        uint64 id;
//
//        RuleResource resource;
//        RuleAction action;
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
type RuleEntry struct {
	Id       Uint64       `json:"id,omitempty"`
	Resource RuleResource `json:"resource,omitempty"`
	Action   RuleAction   `json:"action,omitempty"`
	Forbids  bool         `json:"forbids,omitempty"`
	Details  Longstring   `json:"details,omitempty"`
	Ext      RuleEntryExt `json:"ext,omitempty"`
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
//    	uint64 roleIDs<>;
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
	RoleIDs   []Uint64       `json:"roleIDs,omitempty"`
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
//        case BALANCE:
//            BalanceEntry balance;
//        case ASSET:
//            AssetEntry asset;
//        case DATA:
//            DataEntry data;
//        case REFERENCE_ENTRY:
//            ReferenceEntry reference;
//        case REVIEWABLE_REQUEST:
//    		ReviewableRequestEntry reviewableRequest;
//    	case KEY_VALUE:
//    	    KeyValueEntry keyValue;
//    	case ACCOUNT_KYC:
//            AccountKYCEntry accountKYC;
//        case RULE:
//            RuleEntry rule;
//        case ROLE:
//            RoleEntry role;
//        }
//
type LedgerEntryData struct {
	Type              LedgerEntryType         `json:"type,omitempty"`
	Account           *AccountEntry           `json:"account,omitempty"`
	Signer            *SignerEntry            `json:"signer,omitempty"`
	Balance           *BalanceEntry           `json:"balance,omitempty"`
	Asset             *AssetEntry             `json:"asset,omitempty"`
	Data              *DataEntry              `json:"data,omitempty"`
	Reference         *ReferenceEntry         `json:"reference,omitempty"`
	ReviewableRequest *ReviewableRequestEntry `json:"reviewableRequest,omitempty"`
	KeyValue          *KeyValueEntry          `json:"keyValue,omitempty"`
	AccountKyc        *AccountKycEntry        `json:"accountKYC,omitempty"`
	Rule              *RuleEntry              `json:"rule,omitempty"`
	Role              *RoleEntry              `json:"role,omitempty"`
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
	case LedgerEntryTypeBalance:
		return "Balance", true
	case LedgerEntryTypeAsset:
		return "Asset", true
	case LedgerEntryTypeData:
		return "Data", true
	case LedgerEntryTypeReferenceEntry:
		return "Reference", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeKeyValue:
		return "KeyValue", true
	case LedgerEntryTypeAccountKyc:
		return "AccountKyc", true
	case LedgerEntryTypeRule:
		return "Rule", true
	case LedgerEntryTypeRole:
		return "Role", true
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
	case LedgerEntryTypeBalance:
		tv, ok := value.(BalanceEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be BalanceEntry")
			return
		}
		result.Balance = &tv
	case LedgerEntryTypeAsset:
		tv, ok := value.(AssetEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetEntry")
			return
		}
		result.Asset = &tv
	case LedgerEntryTypeData:
		tv, ok := value.(DataEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be DataEntry")
			return
		}
		result.Data = &tv
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
	case LedgerEntryTypeRule:
		tv, ok := value.(RuleEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleEntry")
			return
		}
		result.Rule = &tv
	case LedgerEntryTypeRole:
		tv, ok := value.(RoleEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be RoleEntry")
			return
		}
		result.Role = &tv
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

// MustBalance retrieves the Balance value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustBalance() BalanceEntry {
	val, ok := u.GetBalance()

	if !ok {
		panic("arm Balance is not set")
	}

	return val
}

// GetBalance retrieves the Balance value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetBalance() (result BalanceEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Balance" {
		result = *u.Balance
		ok = true
	}

	return
}

// MustAsset retrieves the Asset value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAsset() AssetEntry {
	val, ok := u.GetAsset()

	if !ok {
		panic("arm Asset is not set")
	}

	return val
}

// GetAsset retrieves the Asset value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAsset() (result AssetEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Asset" {
		result = *u.Asset
		ok = true
	}

	return
}

// MustData retrieves the Data value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustData() DataEntry {
	val, ok := u.GetData()

	if !ok {
		panic("arm Data is not set")
	}

	return val
}

// GetData retrieves the Data value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetData() (result DataEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Data" {
		result = *u.Data
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

// MustRule retrieves the Rule value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustRule() RuleEntry {
	val, ok := u.GetRule()

	if !ok {
		panic("arm Rule is not set")
	}

	return val
}

// GetRule retrieves the Rule value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetRule() (result RuleEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Rule" {
		result = *u.Rule
		ok = true
	}

	return
}

// MustRole retrieves the Role value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustRole() RoleEntry {
	val, ok := u.GetRole()

	if !ok {
		panic("arm Role is not set")
	}

	return val
}

// GetRole retrieves the Role value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetRole() (result RoleEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Role" {
		result = *u.Role
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
//        case BALANCE:
//            BalanceEntry balance;
//        case ASSET:
//            AssetEntry asset;
//        case DATA:
//            DataEntry data;
//        case REFERENCE_ENTRY:
//            ReferenceEntry reference;
//        case REVIEWABLE_REQUEST:
//    		ReviewableRequestEntry reviewableRequest;
//    	case KEY_VALUE:
//    	    KeyValueEntry keyValue;
//    	case ACCOUNT_KYC:
//            AccountKYCEntry accountKYC;
//        case RULE:
//            RuleEntry rule;
//        case ROLE:
//            RoleEntry role;
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

// LedgerKeyBalanceExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyBalanceExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyBalanceExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyBalanceExt
func (u LedgerKeyBalanceExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyBalanceExt creates a new  LedgerKeyBalanceExt.
func NewLedgerKeyBalanceExt(v LedgerVersion, value interface{}) (result LedgerKeyBalanceExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyBalance is an XDR NestedStruct defines as:
//
//   struct
//        {
//            BalanceID balanceID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyBalance struct {
	BalanceId BalanceId           `json:"balanceID,omitempty"`
	Ext       LedgerKeyBalanceExt `json:"ext,omitempty"`
}

// LedgerKeyAssetExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyAssetExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAssetExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAssetExt
func (u LedgerKeyAssetExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAssetExt creates a new  LedgerKeyAssetExt.
func NewLedgerKeyAssetExt(v LedgerVersion, value interface{}) (result LedgerKeyAssetExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAsset is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AssetCode code;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyAsset struct {
	Code AssetCode         `json:"code,omitempty"`
	Ext  LedgerKeyAssetExt `json:"ext,omitempty"`
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

// LedgerKeyRoleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyRoleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyRoleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyRoleExt
func (u LedgerKeyRoleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyRoleExt creates a new  LedgerKeyRoleExt.
func NewLedgerKeyRoleExt(v LedgerVersion, value interface{}) (result LedgerKeyRoleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyRole is an XDR NestedStruct defines as:
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
type LedgerKeyRole struct {
	Id  Uint64           `json:"id,omitempty"`
	Ext LedgerKeyRoleExt `json:"ext,omitempty"`
}

// LedgerKeyRuleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyRuleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyRuleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyRuleExt
func (u LedgerKeyRuleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyRuleExt creates a new  LedgerKeyRuleExt.
func NewLedgerKeyRuleExt(v LedgerVersion, value interface{}) (result LedgerKeyRuleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyRule is an XDR NestedStruct defines as:
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
type LedgerKeyRule struct {
	Id  Uint64           `json:"id,omitempty"`
	Ext LedgerKeyRuleExt `json:"ext,omitempty"`
}

// LedgerKeyData is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 id;
//
//            EmptyExt ext;
//        }
//
type LedgerKeyData struct {
	Id  Uint64   `json:"id,omitempty"`
	Ext EmptyExt `json:"ext,omitempty"`
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
//    case BALANCE:
//        struct
//        {
//            BalanceID balanceID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } balance;
//    case ASSET:
//        struct
//        {
//            AssetCode code;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } asset;
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
//    case ROLE:
//        struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } role;
//    case RULE:
//        struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } rule;
//    case DATA:
//        struct {
//            uint64 id;
//
//            EmptyExt ext;
//        } data;
//
//    };
//
type LedgerKey struct {
	Type              LedgerEntryType             `json:"type,omitempty"`
	Account           *LedgerKeyAccount           `json:"account,omitempty"`
	Signer            *LedgerKeySigner            `json:"signer,omitempty"`
	Balance           *LedgerKeyBalance           `json:"balance,omitempty"`
	Asset             *LedgerKeyAsset             `json:"asset,omitempty"`
	Reference         *LedgerKeyReference         `json:"reference,omitempty"`
	ReviewableRequest *LedgerKeyReviewableRequest `json:"reviewableRequest,omitempty"`
	KeyValue          *LedgerKeyKeyValue          `json:"keyValue,omitempty"`
	AccountKyc        *LedgerKeyAccountKyc        `json:"accountKYC,omitempty"`
	Role              *LedgerKeyRole              `json:"role,omitempty"`
	Rule              *LedgerKeyRule              `json:"rule,omitempty"`
	Data              *LedgerKeyData              `json:"data,omitempty"`
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
	case LedgerEntryTypeBalance:
		return "Balance", true
	case LedgerEntryTypeAsset:
		return "Asset", true
	case LedgerEntryTypeReferenceEntry:
		return "Reference", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeKeyValue:
		return "KeyValue", true
	case LedgerEntryTypeAccountKyc:
		return "AccountKyc", true
	case LedgerEntryTypeRole:
		return "Role", true
	case LedgerEntryTypeRule:
		return "Rule", true
	case LedgerEntryTypeData:
		return "Data", true
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
	case LedgerEntryTypeBalance:
		tv, ok := value.(LedgerKeyBalance)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyBalance")
			return
		}
		result.Balance = &tv
	case LedgerEntryTypeAsset:
		tv, ok := value.(LedgerKeyAsset)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAsset")
			return
		}
		result.Asset = &tv
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
	case LedgerEntryTypeRole:
		tv, ok := value.(LedgerKeyRole)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyRole")
			return
		}
		result.Role = &tv
	case LedgerEntryTypeRule:
		tv, ok := value.(LedgerKeyRule)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyRule")
			return
		}
		result.Rule = &tv
	case LedgerEntryTypeData:
		tv, ok := value.(LedgerKeyData)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyData")
			return
		}
		result.Data = &tv
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

// MustBalance retrieves the Balance value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustBalance() LedgerKeyBalance {
	val, ok := u.GetBalance()

	if !ok {
		panic("arm Balance is not set")
	}

	return val
}

// GetBalance retrieves the Balance value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetBalance() (result LedgerKeyBalance, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Balance" {
		result = *u.Balance
		ok = true
	}

	return
}

// MustAsset retrieves the Asset value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAsset() LedgerKeyAsset {
	val, ok := u.GetAsset()

	if !ok {
		panic("arm Asset is not set")
	}

	return val
}

// GetAsset retrieves the Asset value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAsset() (result LedgerKeyAsset, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Asset" {
		result = *u.Asset
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

// MustRole retrieves the Role value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustRole() LedgerKeyRole {
	val, ok := u.GetRole()

	if !ok {
		panic("arm Role is not set")
	}

	return val
}

// GetRole retrieves the Role value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetRole() (result LedgerKeyRole, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Role" {
		result = *u.Role
		ok = true
	}

	return
}

// MustRule retrieves the Rule value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustRule() LedgerKeyRule {
	val, ok := u.GetRule()

	if !ok {
		panic("arm Rule is not set")
	}

	return val
}

// GetRule retrieves the Rule value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetRule() (result LedgerKeyRule, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Rule" {
		result = *u.Rule
		ok = true
	}

	return
}

// MustData retrieves the Data value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustData() LedgerKeyData {
	val, ok := u.GetData()

	if !ok {
		panic("arm Data is not set")
	}

	return val
}

// GetData retrieves the Data value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetData() (result LedgerKeyData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Data" {
		result = *u.Data
		ok = true
	}

	return
}

// LedgerUpgradeType is an XDR Enum defines as:
//
//   enum LedgerUpgradeType
//    {
//        NONE = 0,
//        VERSION = 1,
//        MAX_TX_SET_SIZE = 2,
//        TX_EXPIRATION_PERIOD = 3
//    };
//
type LedgerUpgradeType int32

const (
	LedgerUpgradeTypeNone               LedgerUpgradeType = 0
	LedgerUpgradeTypeVersion            LedgerUpgradeType = 1
	LedgerUpgradeTypeMaxTxSetSize       LedgerUpgradeType = 2
	LedgerUpgradeTypeTxExpirationPeriod LedgerUpgradeType = 3
)

var LedgerUpgradeTypeAll = []LedgerUpgradeType{
	LedgerUpgradeTypeNone,
	LedgerUpgradeTypeVersion,
	LedgerUpgradeTypeMaxTxSetSize,
	LedgerUpgradeTypeTxExpirationPeriod,
}

var ledgerUpgradeTypeMap = map[int32]string{
	0: "LedgerUpgradeTypeNone",
	1: "LedgerUpgradeTypeVersion",
	2: "LedgerUpgradeTypeMaxTxSetSize",
	3: "LedgerUpgradeTypeTxExpirationPeriod",
}

var ledgerUpgradeTypeShortMap = map[int32]string{
	0: "none",
	1: "version",
	2: "max_tx_set_size",
	3: "tx_expiration_period",
}

var ledgerUpgradeTypeRevMap = map[string]int32{
	"LedgerUpgradeTypeNone":               0,
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
//    case NONE:
//        void;
//    case VERSION:
//        uint32 newLedgerVersion; // update ledgerVersion
//    case MAX_TX_SET_SIZE:
//        uint32 newMaxTxSetSize; // update maxTxSetSize
//    case TX_EXPIRATION_PERIOD:
//        uint64 newTxExpirationPeriod;
//    };
//
type LedgerUpgrade struct {
	Type                  LedgerUpgradeType `json:"type,omitempty"`
	NewLedgerVersion      *Uint32           `json:"newLedgerVersion,omitempty"`
	NewMaxTxSetSize       *Uint32           `json:"newMaxTxSetSize,omitempty"`
	NewTxExpirationPeriod *Uint64           `json:"newTxExpirationPeriod,omitempty"`
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
	case LedgerUpgradeTypeNone:
		return "", true
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
	case LedgerUpgradeTypeNone:
		// void
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
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
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
func (u LedgerUpgrade) MustNewTxExpirationPeriod() Uint64 {
	val, ok := u.GetNewTxExpirationPeriod()

	if !ok {
		panic("arm NewTxExpirationPeriod is not set")
	}

	return val
}

// GetNewTxExpirationPeriod retrieves the NewTxExpirationPeriod value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewTxExpirationPeriod() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewTxExpirationPeriod" {
		result = *u.NewTxExpirationPeriod
		ok = true
	}

	return
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
//        Hash txSetHash;          // hash of transactions' hashes
//        Hash txSetResultHash;    // the TransactionResultSet that led to this ledger
//
//        uint32 ledgerSeq; // sequence number of this ledger
//        uint64 closeTime; // network close time
//
//        IdGenerator idGenerators<>; // generators of ids
//        LedgerUpgrade upgrade; // upgrade in current ledger (usually none), only one upgrade in one closed ledger is enough
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
	TxSetHash          Hash            `json:"txSetHash,omitempty"`
	TxSetResultHash    Hash            `json:"txSetResultHash,omitempty"`
	LedgerSeq          Uint32          `json:"ledgerSeq,omitempty"`
	CloseTime          Uint64          `json:"closeTime,omitempty"`
	IdGenerators       []IdGenerator   `json:"idGenerators,omitempty"`
	Upgrade            LedgerUpgrade   `json:"upgrade,omitempty"`
	Ext                LedgerHeaderExt `json:"ext,omitempty"`
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

// ChangeAccountRolesOp is an XDR Struct defines as:
//
//   struct ChangeAccountRolesOp
//    {
//        AccountID destinationAccount;
//
//        //: ID of account role that will be attached to `destinationAccount`
//        uint64 rolesToSet<>;
//        //: Arbitrary stringified json object that can be used to attach data to be reviewed by an admin
//        longstring details;
//
//        EmptyExt ext;
//    };
//
type ChangeAccountRolesOp struct {
	DestinationAccount AccountId  `json:"destinationAccount,omitempty"`
	RolesToSet         []Uint64   `json:"rolesToSet,omitempty"`
	Details            Longstring `json:"details,omitempty"`
	Ext                EmptyExt   `json:"ext,omitempty"`
}

// ChangeAccountRolesResultCode is an XDR Enum defines as:
//
//   enum ChangeAccountRolesResultCode
//    {
//        SUCCESS = 0,
//
//        INVALID_DETAILS = -1,
//        ACCOUNT_NOT_FOUND = -2,
//        TOO_MANY_ROLES = -3,
//        NO_SUCH_ROLE = -4,
//        NO_ROLE_IDS = -5,
//        ROLE_ID_DUPLICATION = -6
//    };
//
type ChangeAccountRolesResultCode int32

const (
	ChangeAccountRolesResultCodeSuccess           ChangeAccountRolesResultCode = 0
	ChangeAccountRolesResultCodeInvalidDetails    ChangeAccountRolesResultCode = -1
	ChangeAccountRolesResultCodeAccountNotFound   ChangeAccountRolesResultCode = -2
	ChangeAccountRolesResultCodeTooManyRoles      ChangeAccountRolesResultCode = -3
	ChangeAccountRolesResultCodeNoSuchRole        ChangeAccountRolesResultCode = -4
	ChangeAccountRolesResultCodeNoRoleIds         ChangeAccountRolesResultCode = -5
	ChangeAccountRolesResultCodeRoleIdDuplication ChangeAccountRolesResultCode = -6
)

var ChangeAccountRolesResultCodeAll = []ChangeAccountRolesResultCode{
	ChangeAccountRolesResultCodeSuccess,
	ChangeAccountRolesResultCodeInvalidDetails,
	ChangeAccountRolesResultCodeAccountNotFound,
	ChangeAccountRolesResultCodeTooManyRoles,
	ChangeAccountRolesResultCodeNoSuchRole,
	ChangeAccountRolesResultCodeNoRoleIds,
	ChangeAccountRolesResultCodeRoleIdDuplication,
}

var changeAccountRolesResultCodeMap = map[int32]string{
	0:  "ChangeAccountRolesResultCodeSuccess",
	-1: "ChangeAccountRolesResultCodeInvalidDetails",
	-2: "ChangeAccountRolesResultCodeAccountNotFound",
	-3: "ChangeAccountRolesResultCodeTooManyRoles",
	-4: "ChangeAccountRolesResultCodeNoSuchRole",
	-5: "ChangeAccountRolesResultCodeNoRoleIds",
	-6: "ChangeAccountRolesResultCodeRoleIdDuplication",
}

var changeAccountRolesResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_details",
	-2: "account_not_found",
	-3: "too_many_roles",
	-4: "no_such_role",
	-5: "no_role_ids",
	-6: "role_id_duplication",
}

var changeAccountRolesResultCodeRevMap = map[string]int32{
	"ChangeAccountRolesResultCodeSuccess":           0,
	"ChangeAccountRolesResultCodeInvalidDetails":    -1,
	"ChangeAccountRolesResultCodeAccountNotFound":   -2,
	"ChangeAccountRolesResultCodeTooManyRoles":      -3,
	"ChangeAccountRolesResultCodeNoSuchRole":        -4,
	"ChangeAccountRolesResultCodeNoRoleIds":         -5,
	"ChangeAccountRolesResultCodeRoleIdDuplication": -6,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ChangeAccountRolesResultCode
func (e ChangeAccountRolesResultCode) ValidEnum(v int32) bool {
	_, ok := changeAccountRolesResultCodeMap[v]
	return ok
}
func (e ChangeAccountRolesResultCode) isFlag() bool {
	for i := len(ChangeAccountRolesResultCodeAll) - 1; i >= 0; i-- {
		expected := ChangeAccountRolesResultCode(2) << uint64(len(ChangeAccountRolesResultCodeAll)-1) >> uint64(len(ChangeAccountRolesResultCodeAll)-i)
		if expected != ChangeAccountRolesResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ChangeAccountRolesResultCode) String() string {
	name, _ := changeAccountRolesResultCodeMap[int32(e)]
	return name
}

func (e ChangeAccountRolesResultCode) ShortString() string {
	name, _ := changeAccountRolesResultCodeShortMap[int32(e)]
	return name
}

func (e ChangeAccountRolesResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range ChangeAccountRolesResultCodeAll {
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

func (e *ChangeAccountRolesResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ChangeAccountRolesResultCode(t.Value)
	return nil
}

// ChangeAccountRolesResult is an XDR Union defines as:
//
//   union ChangeAccountRolesResult switch (ChangeAccountRolesResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    case TOO_MANY_ROLES:
//        uint32 maxRolesCount;
//    case NO_SUCH_ROLE:
//    case ROLE_ID_DUPLICATION:
//        uint64 roleID;
//    default:
//        void;
//    };
//
type ChangeAccountRolesResult struct {
	Code          ChangeAccountRolesResultCode `json:"code,omitempty"`
	Ext           *EmptyExt                    `json:"ext,omitempty"`
	MaxRolesCount *Uint32                      `json:"maxRolesCount,omitempty"`
	RoleId        *Uint64                      `json:"roleID,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ChangeAccountRolesResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ChangeAccountRolesResult
func (u ChangeAccountRolesResult) ArmForSwitch(sw int32) (string, bool) {
	switch ChangeAccountRolesResultCode(sw) {
	case ChangeAccountRolesResultCodeSuccess:
		return "Ext", true
	case ChangeAccountRolesResultCodeTooManyRoles:
		return "MaxRolesCount", true
	case ChangeAccountRolesResultCodeNoSuchRole:
		return "RoleId", true
	case ChangeAccountRolesResultCodeRoleIdDuplication:
		return "RoleId", true
	default:
		return "", true
	}
}

// NewChangeAccountRolesResult creates a new  ChangeAccountRolesResult.
func NewChangeAccountRolesResult(code ChangeAccountRolesResultCode, value interface{}) (result ChangeAccountRolesResult, err error) {
	result.Code = code
	switch ChangeAccountRolesResultCode(code) {
	case ChangeAccountRolesResultCodeSuccess:
		tv, ok := value.(EmptyExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be EmptyExt")
			return
		}
		result.Ext = &tv
	case ChangeAccountRolesResultCodeTooManyRoles:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.MaxRolesCount = &tv
	case ChangeAccountRolesResultCodeNoSuchRole:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RoleId = &tv
	case ChangeAccountRolesResultCodeRoleIdDuplication:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RoleId = &tv
	default:
		// void
	}
	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u ChangeAccountRolesResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ChangeAccountRolesResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// MustMaxRolesCount retrieves the MaxRolesCount value from the union,
// panicing if the value is not set.
func (u ChangeAccountRolesResult) MustMaxRolesCount() Uint32 {
	val, ok := u.GetMaxRolesCount()

	if !ok {
		panic("arm MaxRolesCount is not set")
	}

	return val
}

// GetMaxRolesCount retrieves the MaxRolesCount value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ChangeAccountRolesResult) GetMaxRolesCount() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "MaxRolesCount" {
		result = *u.MaxRolesCount
		ok = true
	}

	return
}

// MustRoleId retrieves the RoleId value from the union,
// panicing if the value is not set.
func (u ChangeAccountRolesResult) MustRoleId() Uint64 {
	val, ok := u.GetRoleId()

	if !ok {
		panic("arm RoleId is not set")
	}

	return val
}

// GetRoleId retrieves the RoleId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ChangeAccountRolesResult) GetRoleId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RoleId" {
		result = *u.RoleId
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
//        uint64 roleIDs<>;
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
	RoleIDs     []Uint64           `json:"roleIDs,omitempty"`
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
//        NO_SIGNER_DATA = -6, // empty signer data array not allowed
//        NO_ROLE_IDS = -7,
//        ROLE_ID_DUPLICATION = -8,
//        TOO_MANY_ROLES = -9
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
	CreateAccountResultCodeNoRoleIds          CreateAccountResultCode = -7
	CreateAccountResultCodeRoleIdDuplication  CreateAccountResultCode = -8
	CreateAccountResultCodeTooManyRoles       CreateAccountResultCode = -9
)

var CreateAccountResultCodeAll = []CreateAccountResultCode{
	CreateAccountResultCodeSuccess,
	CreateAccountResultCodeInvalidDestination,
	CreateAccountResultCodeAlreadyExists,
	CreateAccountResultCodeInvalidWeight,
	CreateAccountResultCodeNoSuchRole,
	CreateAccountResultCodeInvalidSignerData,
	CreateAccountResultCodeNoSignerData,
	CreateAccountResultCodeNoRoleIds,
	CreateAccountResultCodeRoleIdDuplication,
	CreateAccountResultCodeTooManyRoles,
}

var createAccountResultCodeMap = map[int32]string{
	0:  "CreateAccountResultCodeSuccess",
	-1: "CreateAccountResultCodeInvalidDestination",
	-2: "CreateAccountResultCodeAlreadyExists",
	-3: "CreateAccountResultCodeInvalidWeight",
	-4: "CreateAccountResultCodeNoSuchRole",
	-5: "CreateAccountResultCodeInvalidSignerData",
	-6: "CreateAccountResultCodeNoSignerData",
	-7: "CreateAccountResultCodeNoRoleIds",
	-8: "CreateAccountResultCodeRoleIdDuplication",
	-9: "CreateAccountResultCodeTooManyRoles",
}

var createAccountResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_destination",
	-2: "already_exists",
	-3: "invalid_weight",
	-4: "no_such_role",
	-5: "invalid_signer_data",
	-6: "no_signer_data",
	-7: "no_role_ids",
	-8: "role_id_duplication",
	-9: "too_many_roles",
}

var createAccountResultCodeRevMap = map[string]int32{
	"CreateAccountResultCodeSuccess":            0,
	"CreateAccountResultCodeInvalidDestination": -1,
	"CreateAccountResultCodeAlreadyExists":      -2,
	"CreateAccountResultCodeInvalidWeight":      -3,
	"CreateAccountResultCodeNoSuchRole":         -4,
	"CreateAccountResultCodeInvalidSignerData":  -5,
	"CreateAccountResultCodeNoSignerData":       -6,
	"CreateAccountResultCodeNoRoleIds":          -7,
	"CreateAccountResultCodeRoleIdDuplication":  -8,
	"CreateAccountResultCodeTooManyRoles":       -9,
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
//        CreateSignerResultCode createSignerErrorCode;
//    case NO_SUCH_ROLE:
//    case ROLE_ID_DUPLICATION:
//        uint64 roleID;
//    case TOO_MANY_ROLES:
//        uint32 maxRolesCount;
//    default:
//        void;
//    };
//
type CreateAccountResult struct {
	Code                  CreateAccountResultCode `json:"code,omitempty"`
	Success               *CreateAccountSuccess   `json:"success,omitempty"`
	CreateSignerErrorCode *CreateSignerResultCode `json:"createSignerErrorCode,omitempty"`
	RoleId                *Uint64                 `json:"roleID,omitempty"`
	MaxRolesCount         *Uint32                 `json:"maxRolesCount,omitempty"`
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
	case CreateAccountResultCodeNoSuchRole:
		return "RoleId", true
	case CreateAccountResultCodeRoleIdDuplication:
		return "RoleId", true
	case CreateAccountResultCodeTooManyRoles:
		return "MaxRolesCount", true
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
		tv, ok := value.(CreateSignerResultCode)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateSignerResultCode")
			return
		}
		result.CreateSignerErrorCode = &tv
	case CreateAccountResultCodeNoSuchRole:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RoleId = &tv
	case CreateAccountResultCodeRoleIdDuplication:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RoleId = &tv
	case CreateAccountResultCodeTooManyRoles:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.MaxRolesCount = &tv
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
func (u CreateAccountResult) MustCreateSignerErrorCode() CreateSignerResultCode {
	val, ok := u.GetCreateSignerErrorCode()

	if !ok {
		panic("arm CreateSignerErrorCode is not set")
	}

	return val
}

// GetCreateSignerErrorCode retrieves the CreateSignerErrorCode value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAccountResult) GetCreateSignerErrorCode() (result CreateSignerResultCode, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "CreateSignerErrorCode" {
		result = *u.CreateSignerErrorCode
		ok = true
	}

	return
}

// MustRoleId retrieves the RoleId value from the union,
// panicing if the value is not set.
func (u CreateAccountResult) MustRoleId() Uint64 {
	val, ok := u.GetRoleId()

	if !ok {
		panic("arm RoleId is not set")
	}

	return val
}

// GetRoleId retrieves the RoleId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAccountResult) GetRoleId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RoleId" {
		result = *u.RoleId
		ok = true
	}

	return
}

// MustMaxRolesCount retrieves the MaxRolesCount value from the union,
// panicing if the value is not set.
func (u CreateAccountResult) MustMaxRolesCount() Uint32 {
	val, ok := u.GetMaxRolesCount()

	if !ok {
		panic("arm MaxRolesCount is not set")
	}

	return val
}

// GetMaxRolesCount retrieves the MaxRolesCount value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAccountResult) GetMaxRolesCount() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "MaxRolesCount" {
		result = *u.MaxRolesCount
		ok = true
	}

	return
}

// CreateAssetOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateAssetOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAssetOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAssetOpExt
func (u CreateAssetOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAssetOpExt creates a new  CreateAssetOpExt.
func NewCreateAssetOpExt(v LedgerVersion, value interface{}) (result CreateAssetOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAssetOp is an XDR Struct defines as:
//
//   struct CreateAssetOp
//    {
//        AssetCode code;
//
//        uint32 securityType; // use instead policies that limit usage, use in account rules
//        uint32 state;
//
//    	uint64 maxIssuanceAmount; // max number of tokens to be issued
//
//        uint32 trailingDigitsCount;
//
//        longstring details;
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
type CreateAssetOp struct {
	Code                AssetCode        `json:"code,omitempty"`
	SecurityType        Uint32           `json:"securityType,omitempty"`
	State               Uint32           `json:"state,omitempty"`
	MaxIssuanceAmount   Uint64           `json:"maxIssuanceAmount,omitempty"`
	TrailingDigitsCount Uint32           `json:"trailingDigitsCount,omitempty"`
	Details             Longstring       `json:"details,omitempty"`
	Ext                 CreateAssetOpExt `json:"ext,omitempty"`
}

// CreateAssetResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageAssetOp
//    enum CreateAssetResultCode
//    {
//        //: Specified action in `data` of ManageSignerOp was successfully performed
//        SUCCESS = 0,                       // request was successfully created/updated/canceled
//
//        //: It is not allowed to create an asset with a code that is already used for another asset
//        ASSET_ALREADY_EXISTS = -1,	      // asset with such code already exist
//        //: It is not allowed to set max issuance amount that is
//        //: less than the sum of issued, pending issuance and available for issuance amounts
//        INVALID_MAX_ISSUANCE_AMOUNT = -2, // max issuance amount is 0
//        //: It is not allowed to use an asset code that is empty or contains space
//        INVALID_CODE = -3,                // asset code is invalid (empty or contains space)
//        //: It is not allowed to use details with invalid json structure
//        INVALID_CREATOR_DETAILS = -4,                        // details must be a valid json
//        //: It is not allowed to set a trailing digits count greater than the maximum trailing digits count (6 at the moment)
//        INVALID_TRAILING_DIGITS_COUNT = -5,          // invalid number of trailing digits
//        //: Maximum issuance amount precision and asset precision are mismatched
//        INVALID_MAX_ISSUANCE_AMOUNT_PRECISION = -6
//    };
//
type CreateAssetResultCode int32

const (
	CreateAssetResultCodeSuccess                           CreateAssetResultCode = 0
	CreateAssetResultCodeAssetAlreadyExists                CreateAssetResultCode = -1
	CreateAssetResultCodeInvalidMaxIssuanceAmount          CreateAssetResultCode = -2
	CreateAssetResultCodeInvalidCode                       CreateAssetResultCode = -3
	CreateAssetResultCodeInvalidCreatorDetails             CreateAssetResultCode = -4
	CreateAssetResultCodeInvalidTrailingDigitsCount        CreateAssetResultCode = -5
	CreateAssetResultCodeInvalidMaxIssuanceAmountPrecision CreateAssetResultCode = -6
)

var CreateAssetResultCodeAll = []CreateAssetResultCode{
	CreateAssetResultCodeSuccess,
	CreateAssetResultCodeAssetAlreadyExists,
	CreateAssetResultCodeInvalidMaxIssuanceAmount,
	CreateAssetResultCodeInvalidCode,
	CreateAssetResultCodeInvalidCreatorDetails,
	CreateAssetResultCodeInvalidTrailingDigitsCount,
	CreateAssetResultCodeInvalidMaxIssuanceAmountPrecision,
}

var createAssetResultCodeMap = map[int32]string{
	0:  "CreateAssetResultCodeSuccess",
	-1: "CreateAssetResultCodeAssetAlreadyExists",
	-2: "CreateAssetResultCodeInvalidMaxIssuanceAmount",
	-3: "CreateAssetResultCodeInvalidCode",
	-4: "CreateAssetResultCodeInvalidCreatorDetails",
	-5: "CreateAssetResultCodeInvalidTrailingDigitsCount",
	-6: "CreateAssetResultCodeInvalidMaxIssuanceAmountPrecision",
}

var createAssetResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "asset_already_exists",
	-2: "invalid_max_issuance_amount",
	-3: "invalid_code",
	-4: "invalid_creator_details",
	-5: "invalid_trailing_digits_count",
	-6: "invalid_max_issuance_amount_precision",
}

var createAssetResultCodeRevMap = map[string]int32{
	"CreateAssetResultCodeSuccess":                           0,
	"CreateAssetResultCodeAssetAlreadyExists":                -1,
	"CreateAssetResultCodeInvalidMaxIssuanceAmount":          -2,
	"CreateAssetResultCodeInvalidCode":                       -3,
	"CreateAssetResultCodeInvalidCreatorDetails":             -4,
	"CreateAssetResultCodeInvalidTrailingDigitsCount":        -5,
	"CreateAssetResultCodeInvalidMaxIssuanceAmountPrecision": -6,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateAssetResultCode
func (e CreateAssetResultCode) ValidEnum(v int32) bool {
	_, ok := createAssetResultCodeMap[v]
	return ok
}
func (e CreateAssetResultCode) isFlag() bool {
	for i := len(CreateAssetResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateAssetResultCode(2) << uint64(len(CreateAssetResultCodeAll)-1) >> uint64(len(CreateAssetResultCodeAll)-i)
		if expected != CreateAssetResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateAssetResultCode) String() string {
	name, _ := createAssetResultCodeMap[int32(e)]
	return name
}

func (e CreateAssetResultCode) ShortString() string {
	name, _ := createAssetResultCodeShortMap[int32(e)]
	return name
}

func (e CreateAssetResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateAssetResultCodeAll {
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

func (e *CreateAssetResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateAssetResultCode(t.Value)
	return nil
}

// CreateAssetResult is an XDR Union defines as:
//
//   //: Is used to return the result of operation application
//    union CreateAssetResult switch (CreateAssetResultCode code)
//    {
//    case SUCCESS:
//        //: Result of successful operation application
//        EmptyExt ext;
//    default:
//        void;
//    };
//
type CreateAssetResult struct {
	Code CreateAssetResultCode `json:"code,omitempty"`
	Ext  *EmptyExt             `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAssetResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAssetResult
func (u CreateAssetResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateAssetResultCode(sw) {
	case CreateAssetResultCodeSuccess:
		return "Ext", true
	default:
		return "", true
	}
}

// NewCreateAssetResult creates a new  CreateAssetResult.
func NewCreateAssetResult(code CreateAssetResultCode, value interface{}) (result CreateAssetResult, err error) {
	result.Code = code
	switch CreateAssetResultCode(code) {
	case CreateAssetResultCodeSuccess:
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
func (u CreateAssetResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAssetResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// CreateBalanceOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateBalanceOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateBalanceOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateBalanceOpExt
func (u CreateBalanceOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateBalanceOpExt creates a new  CreateBalanceOpExt.
func NewCreateBalanceOpExt(v LedgerVersion, value interface{}) (result CreateBalanceOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateBalanceOp is an XDR Struct defines as:
//
//   //: `ManageBalanceOp` applies an `action` of the `ManageBalanceAction` type on the balance of a particular `asset` (referenced to by its AssetCode)
//    //: of the `destination` account (referenced to by its AccountID)
//    struct CreateBalanceOp
//    {
//        //: Defines an account whose balance will be managed
//        AccountID destination;
//        //: Defines an asset code of the balance to which `action` will be applied
//        AssetCode asset;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateBalanceOp struct {
	Destination AccountId          `json:"destination,omitempty"`
	Asset       AssetCode          `json:"asset,omitempty"`
	Ext         CreateBalanceOpExt `json:"ext,omitempty"`
}

// CreateBalanceResultCode is an XDR Enum defines as:
//
//   //: Result codes for the ManageBalance operation
//    enum CreateBalanceResultCode
//    {
//        // codes considered as "success" for the operation
//        //: Indicates that `ManageBalanceOp` is successfully applied
//        SUCCESS = 0,
//
//        //: AssetCode `asset` is invalid (e.g. `AssetCode` does not consist of alphanumeric symbols)
//        INVALID_ASSET = -1,
//        //: Cannot find an asset with a provided asset code
//        ASSET_NOT_FOUND = -2,
//        //: Cannot find an account provided by the `destination` AccountID
//        DESTINATION_NOT_FOUND = -3
//    };
//
type CreateBalanceResultCode int32

const (
	CreateBalanceResultCodeSuccess             CreateBalanceResultCode = 0
	CreateBalanceResultCodeInvalidAsset        CreateBalanceResultCode = -1
	CreateBalanceResultCodeAssetNotFound       CreateBalanceResultCode = -2
	CreateBalanceResultCodeDestinationNotFound CreateBalanceResultCode = -3
)

var CreateBalanceResultCodeAll = []CreateBalanceResultCode{
	CreateBalanceResultCodeSuccess,
	CreateBalanceResultCodeInvalidAsset,
	CreateBalanceResultCodeAssetNotFound,
	CreateBalanceResultCodeDestinationNotFound,
}

var createBalanceResultCodeMap = map[int32]string{
	0:  "CreateBalanceResultCodeSuccess",
	-1: "CreateBalanceResultCodeInvalidAsset",
	-2: "CreateBalanceResultCodeAssetNotFound",
	-3: "CreateBalanceResultCodeDestinationNotFound",
}

var createBalanceResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_asset",
	-2: "asset_not_found",
	-3: "destination_not_found",
}

var createBalanceResultCodeRevMap = map[string]int32{
	"CreateBalanceResultCodeSuccess":             0,
	"CreateBalanceResultCodeInvalidAsset":        -1,
	"CreateBalanceResultCodeAssetNotFound":       -2,
	"CreateBalanceResultCodeDestinationNotFound": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateBalanceResultCode
func (e CreateBalanceResultCode) ValidEnum(v int32) bool {
	_, ok := createBalanceResultCodeMap[v]
	return ok
}
func (e CreateBalanceResultCode) isFlag() bool {
	for i := len(CreateBalanceResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateBalanceResultCode(2) << uint64(len(CreateBalanceResultCodeAll)-1) >> uint64(len(CreateBalanceResultCodeAll)-i)
		if expected != CreateBalanceResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateBalanceResultCode) String() string {
	name, _ := createBalanceResultCodeMap[int32(e)]
	return name
}

func (e CreateBalanceResultCode) ShortString() string {
	name, _ := createBalanceResultCodeShortMap[int32(e)]
	return name
}

func (e CreateBalanceResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateBalanceResultCodeAll {
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

func (e *CreateBalanceResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateBalanceResultCode(t.Value)
	return nil
}

// CreateBalanceSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateBalanceSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateBalanceSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateBalanceSuccessExt
func (u CreateBalanceSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateBalanceSuccessExt creates a new  CreateBalanceSuccessExt.
func NewCreateBalanceSuccessExt(v LedgerVersion, value interface{}) (result CreateBalanceSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateBalanceSuccess is an XDR Struct defines as:
//
//   struct CreateBalanceSuccess {
//        //: ID of the balance that was managed
//        BalanceID balanceID;
//        //: reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateBalanceSuccess struct {
	BalanceId BalanceId               `json:"balanceID,omitempty"`
	Ext       CreateBalanceSuccessExt `json:"ext,omitempty"`
}

// CreateBalanceResult is an XDR Union defines as:
//
//   union CreateBalanceResult switch (CreateBalanceResultCode code)
//    {
//    case SUCCESS:
//        CreateBalanceSuccess success;
//    default:
//        void;
//    };
//
type CreateBalanceResult struct {
	Code    CreateBalanceResultCode `json:"code,omitempty"`
	Success *CreateBalanceSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateBalanceResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateBalanceResult
func (u CreateBalanceResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateBalanceResultCode(sw) {
	case CreateBalanceResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateBalanceResult creates a new  CreateBalanceResult.
func NewCreateBalanceResult(code CreateBalanceResultCode, value interface{}) (result CreateBalanceResult, err error) {
	result.Code = code
	switch CreateBalanceResultCode(code) {
	case CreateBalanceResultCodeSuccess:
		tv, ok := value.(CreateBalanceSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateBalanceSuccess")
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
func (u CreateBalanceResult) MustSuccess() CreateBalanceSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateBalanceResult) GetSuccess() (result CreateBalanceSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateDataOp is an XDR Struct defines as:
//
//   struct CreateDataOp
//    {
//        uint32 securityType;
//        longstring value;
//
//        EmptyExt ext;
//    };
//
type CreateDataOp struct {
	SecurityType Uint32     `json:"securityType,omitempty"`
	Value        Longstring `json:"value,omitempty"`
	Ext          EmptyExt   `json:"ext,omitempty"`
}

// CreateDataResultCode is an XDR Enum defines as:
//
//   enum CreateDataResultCode
//    {
//        SUCCESS = 0,
//
//        INVALID_DATA = -1
//    };
//
type CreateDataResultCode int32

const (
	CreateDataResultCodeSuccess     CreateDataResultCode = 0
	CreateDataResultCodeInvalidData CreateDataResultCode = -1
)

var CreateDataResultCodeAll = []CreateDataResultCode{
	CreateDataResultCodeSuccess,
	CreateDataResultCodeInvalidData,
}

var createDataResultCodeMap = map[int32]string{
	0:  "CreateDataResultCodeSuccess",
	-1: "CreateDataResultCodeInvalidData",
}

var createDataResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_data",
}

var createDataResultCodeRevMap = map[string]int32{
	"CreateDataResultCodeSuccess":     0,
	"CreateDataResultCodeInvalidData": -1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateDataResultCode
func (e CreateDataResultCode) ValidEnum(v int32) bool {
	_, ok := createDataResultCodeMap[v]
	return ok
}
func (e CreateDataResultCode) isFlag() bool {
	for i := len(CreateDataResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateDataResultCode(2) << uint64(len(CreateDataResultCodeAll)-1) >> uint64(len(CreateDataResultCodeAll)-i)
		if expected != CreateDataResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateDataResultCode) String() string {
	name, _ := createDataResultCodeMap[int32(e)]
	return name
}

func (e CreateDataResultCode) ShortString() string {
	name, _ := createDataResultCodeShortMap[int32(e)]
	return name
}

func (e CreateDataResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateDataResultCodeAll {
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

func (e *CreateDataResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateDataResultCode(t.Value)
	return nil
}

// CreateDataSuccess is an XDR Struct defines as:
//
//   struct CreateDataSuccess
//    {
//        uint64 dataID;
//
//        EmptyExt ext;
//    };
//
type CreateDataSuccess struct {
	DataId Uint64   `json:"dataID,omitempty"`
	Ext    EmptyExt `json:"ext,omitempty"`
}

// CreateDataResult is an XDR Union defines as:
//
//   union CreateDataResult switch (CreateDataResultCode code)
//    {
//        case SUCCESS:
//            CreateDataSuccess success;
//        default:
//            void;
//    };
//
type CreateDataResult struct {
	Code    CreateDataResultCode `json:"code,omitempty"`
	Success *CreateDataSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateDataResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateDataResult
func (u CreateDataResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateDataResultCode(sw) {
	case CreateDataResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateDataResult creates a new  CreateDataResult.
func NewCreateDataResult(code CreateDataResultCode, value interface{}) (result CreateDataResult, err error) {
	result.Code = code
	switch CreateDataResultCode(code) {
	case CreateDataResultCodeSuccess:
		tv, ok := value.(CreateDataSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateDataSuccess")
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
func (u CreateDataResult) MustSuccess() CreateDataSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateDataResult) GetSuccess() (result CreateDataSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateReviewableRequestOp is an XDR Struct defines as:
//
//   struct CreateReviewableRequestOp
//    {
//        uint32 securityType;
//        ReviewableRequestOperation operations<>;
//
//        EmptyExt ext;
//    };
//
type CreateReviewableRequestOp struct {
	SecurityType Uint32                       `json:"securityType,omitempty"`
	Operations   []ReviewableRequestOperation `json:"operations,omitempty"`
	Ext          EmptyExt                     `json:"ext,omitempty"`
}

// CreateReviewableRequestResultCode is an XDR Enum defines as:
//
//   enum CreateReviewableRequestResultCode
//    {
//        SUCCESS = 0,
//
//        INVALID_OPERATION = -1,
//        TASKS_NOT_FOUND = -2,
//        TOO_MANY_OPERATIONS = -3
//    };
//
type CreateReviewableRequestResultCode int32

const (
	CreateReviewableRequestResultCodeSuccess           CreateReviewableRequestResultCode = 0
	CreateReviewableRequestResultCodeInvalidOperation  CreateReviewableRequestResultCode = -1
	CreateReviewableRequestResultCodeTasksNotFound     CreateReviewableRequestResultCode = -2
	CreateReviewableRequestResultCodeTooManyOperations CreateReviewableRequestResultCode = -3
)

var CreateReviewableRequestResultCodeAll = []CreateReviewableRequestResultCode{
	CreateReviewableRequestResultCodeSuccess,
	CreateReviewableRequestResultCodeInvalidOperation,
	CreateReviewableRequestResultCodeTasksNotFound,
	CreateReviewableRequestResultCodeTooManyOperations,
}

var createReviewableRequestResultCodeMap = map[int32]string{
	0:  "CreateReviewableRequestResultCodeSuccess",
	-1: "CreateReviewableRequestResultCodeInvalidOperation",
	-2: "CreateReviewableRequestResultCodeTasksNotFound",
	-3: "CreateReviewableRequestResultCodeTooManyOperations",
}

var createReviewableRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_operation",
	-2: "tasks_not_found",
	-3: "too_many_operations",
}

var createReviewableRequestResultCodeRevMap = map[string]int32{
	"CreateReviewableRequestResultCodeSuccess":           0,
	"CreateReviewableRequestResultCodeInvalidOperation":  -1,
	"CreateReviewableRequestResultCodeTasksNotFound":     -2,
	"CreateReviewableRequestResultCodeTooManyOperations": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateReviewableRequestResultCode
func (e CreateReviewableRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createReviewableRequestResultCodeMap[v]
	return ok
}
func (e CreateReviewableRequestResultCode) isFlag() bool {
	for i := len(CreateReviewableRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateReviewableRequestResultCode(2) << uint64(len(CreateReviewableRequestResultCodeAll)-1) >> uint64(len(CreateReviewableRequestResultCodeAll)-i)
		if expected != CreateReviewableRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateReviewableRequestResultCode) String() string {
	name, _ := createReviewableRequestResultCodeMap[int32(e)]
	return name
}

func (e CreateReviewableRequestResultCode) ShortString() string {
	name, _ := createReviewableRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreateReviewableRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateReviewableRequestResultCodeAll {
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

func (e *CreateReviewableRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateReviewableRequestResultCode(t.Value)
	return nil
}

// CreateReviewableRequestSuccessResult is an XDR Struct defines as:
//
//   struct CreateReviewableRequestSuccessResult
//    {
//        uint64 requestID;
//
//        ExtendedResult extendedResult;
//
//        EmptyExt ext;
//    };
//
type CreateReviewableRequestSuccessResult struct {
	RequestId      Uint64         `json:"requestID,omitempty"`
	ExtendedResult ExtendedResult `json:"extendedResult,omitempty"`
	Ext            EmptyExt       `json:"ext,omitempty"`
}

// CreateReviewableRequestResult is an XDR Union defines as:
//
//   union CreateReviewableRequestResult switch (CreateReviewableRequestResultCode code)
//    {
//    case SUCCESS:
//        CreateReviewableRequestSuccessResult success;
//    case INVALID_OPERATION:
//        OperationResult operationResult;
//    case TOO_MANY_OPERATIONS:
//        uint32 maxOperationsCount;
//    default:
//        void;
//    };
//
type CreateReviewableRequestResult struct {
	Code               CreateReviewableRequestResultCode     `json:"code,omitempty"`
	Success            *CreateReviewableRequestSuccessResult `json:"success,omitempty"`
	OperationResult    *OperationResult                      `json:"operationResult,omitempty"`
	MaxOperationsCount *Uint32                               `json:"maxOperationsCount,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateReviewableRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateReviewableRequestResult
func (u CreateReviewableRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateReviewableRequestResultCode(sw) {
	case CreateReviewableRequestResultCodeSuccess:
		return "Success", true
	case CreateReviewableRequestResultCodeInvalidOperation:
		return "OperationResult", true
	case CreateReviewableRequestResultCodeTooManyOperations:
		return "MaxOperationsCount", true
	default:
		return "", true
	}
}

// NewCreateReviewableRequestResult creates a new  CreateReviewableRequestResult.
func NewCreateReviewableRequestResult(code CreateReviewableRequestResultCode, value interface{}) (result CreateReviewableRequestResult, err error) {
	result.Code = code
	switch CreateReviewableRequestResultCode(code) {
	case CreateReviewableRequestResultCodeSuccess:
		tv, ok := value.(CreateReviewableRequestSuccessResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateReviewableRequestSuccessResult")
			return
		}
		result.Success = &tv
	case CreateReviewableRequestResultCodeInvalidOperation:
		tv, ok := value.(OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be OperationResult")
			return
		}
		result.OperationResult = &tv
	case CreateReviewableRequestResultCodeTooManyOperations:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.MaxOperationsCount = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateReviewableRequestResult) MustSuccess() CreateReviewableRequestSuccessResult {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateReviewableRequestResult) GetSuccess() (result CreateReviewableRequestSuccessResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustOperationResult retrieves the OperationResult value from the union,
// panicing if the value is not set.
func (u CreateReviewableRequestResult) MustOperationResult() OperationResult {
	val, ok := u.GetOperationResult()

	if !ok {
		panic("arm OperationResult is not set")
	}

	return val
}

// GetOperationResult retrieves the OperationResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateReviewableRequestResult) GetOperationResult() (result OperationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "OperationResult" {
		result = *u.OperationResult
		ok = true
	}

	return
}

// MustMaxOperationsCount retrieves the MaxOperationsCount value from the union,
// panicing if the value is not set.
func (u CreateReviewableRequestResult) MustMaxOperationsCount() Uint32 {
	val, ok := u.GetMaxOperationsCount()

	if !ok {
		panic("arm MaxOperationsCount is not set")
	}

	return val
}

// GetMaxOperationsCount retrieves the MaxOperationsCount value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateReviewableRequestResult) GetMaxOperationsCount() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "MaxOperationsCount" {
		result = *u.MaxOperationsCount
		ok = true
	}

	return
}

// CreateRoleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateRoleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateRoleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateRoleOpExt
func (u CreateRoleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateRoleOpExt creates a new  CreateRoleOpExt.
func NewCreateRoleOpExt(v LedgerVersion, value interface{}) (result CreateRoleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateRoleOp is an XDR Struct defines as:
//
//   //: CreateSignerRoleData is used to pass necessary params to create a new signer role
//    struct CreateRoleOp
//    {
//        //: Array of ids of existing, unique and not default rules
//        uint64 ruleIDs<>;
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
type CreateRoleOp struct {
	RuleIDs []Uint64        `json:"ruleIDs,omitempty"`
	Details Longstring      `json:"details,omitempty"`
	Ext     CreateRoleOpExt `json:"ext,omitempty"`
}

// CreateRoleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerRoleResultCode
//    enum CreateRoleResultCode
//    {
//        //: Means that the specified action in `data` of ManageSignerRoleOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: Passed details have invalid json structure
//        INVALID_DETAILS = -1,
//        //: There is no rule with id passed through `ruleIDs`
//        NO_SUCH_RULE = -2,
//        //: It is not allowed to duplicate ids in `ruleIDs` array
//        RULE_ID_DUPLICATION = -3,
//        //: It is not allowed to pass ruleIDs that are more than maxSignerRuleCount (by default, 128)
//        TOO_MANY_RULE_IDS = -4
//    };
//
type CreateRoleResultCode int32

const (
	CreateRoleResultCodeSuccess           CreateRoleResultCode = 0
	CreateRoleResultCodeInvalidDetails    CreateRoleResultCode = -1
	CreateRoleResultCodeNoSuchRule        CreateRoleResultCode = -2
	CreateRoleResultCodeRuleIdDuplication CreateRoleResultCode = -3
	CreateRoleResultCodeTooManyRuleIds    CreateRoleResultCode = -4
)

var CreateRoleResultCodeAll = []CreateRoleResultCode{
	CreateRoleResultCodeSuccess,
	CreateRoleResultCodeInvalidDetails,
	CreateRoleResultCodeNoSuchRule,
	CreateRoleResultCodeRuleIdDuplication,
	CreateRoleResultCodeTooManyRuleIds,
}

var createRoleResultCodeMap = map[int32]string{
	0:  "CreateRoleResultCodeSuccess",
	-1: "CreateRoleResultCodeInvalidDetails",
	-2: "CreateRoleResultCodeNoSuchRule",
	-3: "CreateRoleResultCodeRuleIdDuplication",
	-4: "CreateRoleResultCodeTooManyRuleIds",
}

var createRoleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_details",
	-2: "no_such_rule",
	-3: "rule_id_duplication",
	-4: "too_many_rule_ids",
}

var createRoleResultCodeRevMap = map[string]int32{
	"CreateRoleResultCodeSuccess":           0,
	"CreateRoleResultCodeInvalidDetails":    -1,
	"CreateRoleResultCodeNoSuchRule":        -2,
	"CreateRoleResultCodeRuleIdDuplication": -3,
	"CreateRoleResultCodeTooManyRuleIds":    -4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateRoleResultCode
func (e CreateRoleResultCode) ValidEnum(v int32) bool {
	_, ok := createRoleResultCodeMap[v]
	return ok
}
func (e CreateRoleResultCode) isFlag() bool {
	for i := len(CreateRoleResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateRoleResultCode(2) << uint64(len(CreateRoleResultCodeAll)-1) >> uint64(len(CreateRoleResultCodeAll)-i)
		if expected != CreateRoleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateRoleResultCode) String() string {
	name, _ := createRoleResultCodeMap[int32(e)]
	return name
}

func (e CreateRoleResultCode) ShortString() string {
	name, _ := createRoleResultCodeShortMap[int32(e)]
	return name
}

func (e CreateRoleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateRoleResultCodeAll {
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

func (e *CreateRoleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateRoleResultCode(t.Value)
	return nil
}

// CreateRoleResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type CreateRoleResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateRoleResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateRoleResultSuccessExt
func (u CreateRoleResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateRoleResultSuccessExt creates a new  CreateRoleResultSuccessExt.
func NewCreateRoleResultSuccessExt(v LedgerVersion, value interface{}) (result CreateRoleResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateRoleResultSuccess is an XDR NestedStruct defines as:
//
//   struct
//        {
//            //: id of a role that was managed
//            uint64 roleID;
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
type CreateRoleResultSuccess struct {
	RoleId Uint64                     `json:"roleID,omitempty"`
	Ext    CreateRoleResultSuccessExt `json:"ext,omitempty"`
}

// CreateRoleResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union CreateRoleResult switch (CreateRoleResultCode code)
//    {
//    case SUCCESS:
//        struct
//        {
//            //: id of a role that was managed
//            uint64 roleID;
//
//            //: reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } success;
//    case RULE_ID_DUPLICATION:
//    case NO_SUCH_RULE:
//        //: ID of a rule that was either duplicated or is default or does not exist
//        uint64 ruleID;
//    case TOO_MANY_RULE_IDS:
//        //: max count of rule ids that can be passed in `ruleIDs` array
//        uint32 maxRuleIDsCount;
//    default:
//        void;
//    };
//
type CreateRoleResult struct {
	Code            CreateRoleResultCode     `json:"code,omitempty"`
	Success         *CreateRoleResultSuccess `json:"success,omitempty"`
	RuleId          *Uint64                  `json:"ruleID,omitempty"`
	MaxRuleIDsCount *Uint32                  `json:"maxRuleIDsCount,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateRoleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateRoleResult
func (u CreateRoleResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateRoleResultCode(sw) {
	case CreateRoleResultCodeSuccess:
		return "Success", true
	case CreateRoleResultCodeRuleIdDuplication:
		return "RuleId", true
	case CreateRoleResultCodeNoSuchRule:
		return "RuleId", true
	case CreateRoleResultCodeTooManyRuleIds:
		return "MaxRuleIDsCount", true
	default:
		return "", true
	}
}

// NewCreateRoleResult creates a new  CreateRoleResult.
func NewCreateRoleResult(code CreateRoleResultCode, value interface{}) (result CreateRoleResult, err error) {
	result.Code = code
	switch CreateRoleResultCode(code) {
	case CreateRoleResultCodeSuccess:
		tv, ok := value.(CreateRoleResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateRoleResultSuccess")
			return
		}
		result.Success = &tv
	case CreateRoleResultCodeRuleIdDuplication:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RuleId = &tv
	case CreateRoleResultCodeNoSuchRule:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RuleId = &tv
	case CreateRoleResultCodeTooManyRuleIds:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
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
func (u CreateRoleResult) MustSuccess() CreateRoleResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateRoleResult) GetSuccess() (result CreateRoleResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustRuleId retrieves the RuleId value from the union,
// panicing if the value is not set.
func (u CreateRoleResult) MustRuleId() Uint64 {
	val, ok := u.GetRuleId()

	if !ok {
		panic("arm RuleId is not set")
	}

	return val
}

// GetRuleId retrieves the RuleId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateRoleResult) GetRuleId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RuleId" {
		result = *u.RuleId
		ok = true
	}

	return
}

// MustMaxRuleIDsCount retrieves the MaxRuleIDsCount value from the union,
// panicing if the value is not set.
func (u CreateRoleResult) MustMaxRuleIDsCount() Uint32 {
	val, ok := u.GetMaxRuleIDsCount()

	if !ok {
		panic("arm MaxRuleIDsCount is not set")
	}

	return val
}

// GetMaxRuleIDsCount retrieves the MaxRuleIDsCount value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateRoleResult) GetMaxRuleIDsCount() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "MaxRuleIDsCount" {
		result = *u.MaxRuleIDsCount
		ok = true
	}

	return
}

// CreateRuleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateRuleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateRuleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateRuleOpExt
func (u CreateRuleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateRuleOpExt creates a new  CreateRuleOpExt.
func NewCreateRuleOpExt(v LedgerVersion, value interface{}) (result CreateRuleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateRuleOp is an XDR Struct defines as:
//
//   //: CreateSignerRuleData is used to pass necessary params to create a new signer rule
//    struct CreateRuleOp
//    {
//        //: Resource is used to specify an entity (for some, with properties) that can be managed through operations
//        RuleResource resource;
//        //: Value from enum that can be applied to `resource`
//        RuleAction action;
//        //: Indicate whether or not an `action` on the provided `resource` is prohibited
//        bool forbids;
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
type CreateRuleOp struct {
	Resource RuleResource    `json:"resource,omitempty"`
	Action   RuleAction      `json:"action,omitempty"`
	Forbids  bool            `json:"forbids,omitempty"`
	Details  Longstring      `json:"details,omitempty"`
	Ext      CreateRuleOpExt `json:"ext,omitempty"`
}

// CreateRuleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerRuleOp
//    enum CreateRuleResultCode
//    {
//        //: Specified action in `data` of ManageSignerRuleOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: Passed details have invalid json structure
//        INVALID_DETAILS = -1,
//
//        INVALID_CUSTOM_ACTION = -2,
//        INVALID_CUSTOM_RESOURCE = -3
//    };
//
type CreateRuleResultCode int32

const (
	CreateRuleResultCodeSuccess               CreateRuleResultCode = 0
	CreateRuleResultCodeInvalidDetails        CreateRuleResultCode = -1
	CreateRuleResultCodeInvalidCustomAction   CreateRuleResultCode = -2
	CreateRuleResultCodeInvalidCustomResource CreateRuleResultCode = -3
)

var CreateRuleResultCodeAll = []CreateRuleResultCode{
	CreateRuleResultCodeSuccess,
	CreateRuleResultCodeInvalidDetails,
	CreateRuleResultCodeInvalidCustomAction,
	CreateRuleResultCodeInvalidCustomResource,
}

var createRuleResultCodeMap = map[int32]string{
	0:  "CreateRuleResultCodeSuccess",
	-1: "CreateRuleResultCodeInvalidDetails",
	-2: "CreateRuleResultCodeInvalidCustomAction",
	-3: "CreateRuleResultCodeInvalidCustomResource",
}

var createRuleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_details",
	-2: "invalid_custom_action",
	-3: "invalid_custom_resource",
}

var createRuleResultCodeRevMap = map[string]int32{
	"CreateRuleResultCodeSuccess":               0,
	"CreateRuleResultCodeInvalidDetails":        -1,
	"CreateRuleResultCodeInvalidCustomAction":   -2,
	"CreateRuleResultCodeInvalidCustomResource": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateRuleResultCode
func (e CreateRuleResultCode) ValidEnum(v int32) bool {
	_, ok := createRuleResultCodeMap[v]
	return ok
}
func (e CreateRuleResultCode) isFlag() bool {
	for i := len(CreateRuleResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateRuleResultCode(2) << uint64(len(CreateRuleResultCodeAll)-1) >> uint64(len(CreateRuleResultCodeAll)-i)
		if expected != CreateRuleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateRuleResultCode) String() string {
	name, _ := createRuleResultCodeMap[int32(e)]
	return name
}

func (e CreateRuleResultCode) ShortString() string {
	name, _ := createRuleResultCodeShortMap[int32(e)]
	return name
}

func (e CreateRuleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateRuleResultCodeAll {
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

func (e *CreateRuleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateRuleResultCode(t.Value)
	return nil
}

// CreateRuleResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type CreateRuleResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateRuleResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateRuleResultSuccessExt
func (u CreateRuleResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateRuleResultSuccessExt creates a new  CreateRuleResultSuccessExt.
func NewCreateRuleResultSuccessExt(v LedgerVersion, value interface{}) (result CreateRuleResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateRuleResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//            //: id of the rule that was managed
//            uint64 ruleID;
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
type CreateRuleResultSuccess struct {
	RuleId Uint64                     `json:"ruleID,omitempty"`
	Ext    CreateRuleResultSuccessExt `json:"ext,omitempty"`
}

// CreateRuleResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union CreateRuleResult switch (CreateRuleResultCode code)
//    {
//    case SUCCESS:
//        struct {
//            //: id of the rule that was managed
//            uint64 ruleID;
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
type CreateRuleResult struct {
	Code    CreateRuleResultCode     `json:"code,omitempty"`
	Success *CreateRuleResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateRuleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateRuleResult
func (u CreateRuleResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateRuleResultCode(sw) {
	case CreateRuleResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateRuleResult creates a new  CreateRuleResult.
func NewCreateRuleResult(code CreateRuleResultCode, value interface{}) (result CreateRuleResult, err error) {
	result.Code = code
	switch CreateRuleResultCode(code) {
	case CreateRuleResultCodeSuccess:
		tv, ok := value.(CreateRuleResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateRuleResultSuccess")
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
func (u CreateRuleResult) MustSuccess() CreateRuleResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateRuleResult) GetSuccess() (result CreateRuleResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// SignerData is an XDR Struct defines as:
//
//   struct SignerData
//    {
//        //: Public key of a signer
//        PublicKey publicKey;
//        //: id of the role that will be attached to a signer
//        uint64 roleIDs<>;
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
	RoleIDs   []Uint64   `json:"roleIDs,omitempty"`
	Weight    Uint32     `json:"weight,omitempty"`
	Identity  Uint32     `json:"identity,omitempty"`
	Details   Longstring `json:"details,omitempty"`
	Ext       EmptyExt   `json:"ext,omitempty"`
}

// CreateSignerOp is an XDR Struct defines as:
//
//   struct CreateSignerOp
//    {
//        SignerData data;
//
//        //: reserved for future extension
//        EmptyExt ext;
//    };
//
type CreateSignerOp struct {
	Data SignerData `json:"data,omitempty"`
	Ext  EmptyExt   `json:"ext,omitempty"`
}

// CreateSignerResultCode is an XDR Enum defines as:
//
//   enum CreateSignerResultCode
//    {
//        SUCCESS = 0,
//
//        //: Passed details have invalid json structure
//        INVALID_DETAILS = -1, // invalid json details
//        //: Signer with such public key is already attached to the source account
//        ALREADY_EXISTS = -2, // signer already exist
//        //: There is no role with such id
//        NO_SUCH_ROLE = -3,
//        //: It is not allowed to set weight more than 1000
//        INVALID_WEIGHT = -4, // more than 1000
//        NO_ROLE_IDS = -5,
//        ROLE_ID_DUPLICATION = -6,
//        TOO_MANY_ROLES = -7
//    };
//
type CreateSignerResultCode int32

const (
	CreateSignerResultCodeSuccess           CreateSignerResultCode = 0
	CreateSignerResultCodeInvalidDetails    CreateSignerResultCode = -1
	CreateSignerResultCodeAlreadyExists     CreateSignerResultCode = -2
	CreateSignerResultCodeNoSuchRole        CreateSignerResultCode = -3
	CreateSignerResultCodeInvalidWeight     CreateSignerResultCode = -4
	CreateSignerResultCodeNoRoleIds         CreateSignerResultCode = -5
	CreateSignerResultCodeRoleIdDuplication CreateSignerResultCode = -6
	CreateSignerResultCodeTooManyRoles      CreateSignerResultCode = -7
)

var CreateSignerResultCodeAll = []CreateSignerResultCode{
	CreateSignerResultCodeSuccess,
	CreateSignerResultCodeInvalidDetails,
	CreateSignerResultCodeAlreadyExists,
	CreateSignerResultCodeNoSuchRole,
	CreateSignerResultCodeInvalidWeight,
	CreateSignerResultCodeNoRoleIds,
	CreateSignerResultCodeRoleIdDuplication,
	CreateSignerResultCodeTooManyRoles,
}

var createSignerResultCodeMap = map[int32]string{
	0:  "CreateSignerResultCodeSuccess",
	-1: "CreateSignerResultCodeInvalidDetails",
	-2: "CreateSignerResultCodeAlreadyExists",
	-3: "CreateSignerResultCodeNoSuchRole",
	-4: "CreateSignerResultCodeInvalidWeight",
	-5: "CreateSignerResultCodeNoRoleIds",
	-6: "CreateSignerResultCodeRoleIdDuplication",
	-7: "CreateSignerResultCodeTooManyRoles",
}

var createSignerResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_details",
	-2: "already_exists",
	-3: "no_such_role",
	-4: "invalid_weight",
	-5: "no_role_ids",
	-6: "role_id_duplication",
	-7: "too_many_roles",
}

var createSignerResultCodeRevMap = map[string]int32{
	"CreateSignerResultCodeSuccess":           0,
	"CreateSignerResultCodeInvalidDetails":    -1,
	"CreateSignerResultCodeAlreadyExists":     -2,
	"CreateSignerResultCodeNoSuchRole":        -3,
	"CreateSignerResultCodeInvalidWeight":     -4,
	"CreateSignerResultCodeNoRoleIds":         -5,
	"CreateSignerResultCodeRoleIdDuplication": -6,
	"CreateSignerResultCodeTooManyRoles":      -7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateSignerResultCode
func (e CreateSignerResultCode) ValidEnum(v int32) bool {
	_, ok := createSignerResultCodeMap[v]
	return ok
}
func (e CreateSignerResultCode) isFlag() bool {
	for i := len(CreateSignerResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateSignerResultCode(2) << uint64(len(CreateSignerResultCodeAll)-1) >> uint64(len(CreateSignerResultCodeAll)-i)
		if expected != CreateSignerResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateSignerResultCode) String() string {
	name, _ := createSignerResultCodeMap[int32(e)]
	return name
}

func (e CreateSignerResultCode) ShortString() string {
	name, _ := createSignerResultCodeShortMap[int32(e)]
	return name
}

func (e CreateSignerResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range CreateSignerResultCodeAll {
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

func (e *CreateSignerResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateSignerResultCode(t.Value)
	return nil
}

// CreateSignerResult is an XDR Union defines as:
//
//   union CreateSignerResult switch (CreateSignerResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    case NO_SUCH_ROLE:
//    case ROLE_ID_DUPLICATION:
//        uint64 roleID;
//    case TOO_MANY_ROLES:
//        uint32 maxRolesCount;
//    default:
//        void;
//    };
//
type CreateSignerResult struct {
	Code          CreateSignerResultCode `json:"code,omitempty"`
	Ext           *EmptyExt              `json:"ext,omitempty"`
	RoleId        *Uint64                `json:"roleID,omitempty"`
	MaxRolesCount *Uint32                `json:"maxRolesCount,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateSignerResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateSignerResult
func (u CreateSignerResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateSignerResultCode(sw) {
	case CreateSignerResultCodeSuccess:
		return "Ext", true
	case CreateSignerResultCodeNoSuchRole:
		return "RoleId", true
	case CreateSignerResultCodeRoleIdDuplication:
		return "RoleId", true
	case CreateSignerResultCodeTooManyRoles:
		return "MaxRolesCount", true
	default:
		return "", true
	}
}

// NewCreateSignerResult creates a new  CreateSignerResult.
func NewCreateSignerResult(code CreateSignerResultCode, value interface{}) (result CreateSignerResult, err error) {
	result.Code = code
	switch CreateSignerResultCode(code) {
	case CreateSignerResultCodeSuccess:
		tv, ok := value.(EmptyExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be EmptyExt")
			return
		}
		result.Ext = &tv
	case CreateSignerResultCodeNoSuchRole:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RoleId = &tv
	case CreateSignerResultCodeRoleIdDuplication:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RoleId = &tv
	case CreateSignerResultCodeTooManyRoles:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.MaxRolesCount = &tv
	default:
		// void
	}
	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u CreateSignerResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateSignerResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// MustRoleId retrieves the RoleId value from the union,
// panicing if the value is not set.
func (u CreateSignerResult) MustRoleId() Uint64 {
	val, ok := u.GetRoleId()

	if !ok {
		panic("arm RoleId is not set")
	}

	return val
}

// GetRoleId retrieves the RoleId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateSignerResult) GetRoleId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RoleId" {
		result = *u.RoleId
		ok = true
	}

	return
}

// MustMaxRolesCount retrieves the MaxRolesCount value from the union,
// panicing if the value is not set.
func (u CreateSignerResult) MustMaxRolesCount() Uint32 {
	val, ok := u.GetMaxRolesCount()

	if !ok {
		panic("arm MaxRolesCount is not set")
	}

	return val
}

// GetMaxRolesCount retrieves the MaxRolesCount value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateSignerResult) GetMaxRolesCount() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "MaxRolesCount" {
		result = *u.MaxRolesCount
		ok = true
	}

	return
}

// DestructionOp is an XDR Struct defines as:
//
//   //: Destruction operation will charge the specified amount from balance
//    struct DestructionOp
//    {
//        //: security type
//        uint32 securityType;
//        //: Balance to withdraw from
//        BalanceID balance; // balance id from which withdrawal will be performed
//        //: Amount to withdraw
//        uint64 amount; // amount to be withdrawn
//
//        longstring reference;
//
//        //: Total fee to pay, contains fixed amount and calculated percent of the withdrawn amount
//        Fee fee; // expected fee to be paid
//        //: Arbitrary stringified json object that can be used to attach data to be reviewed by an admin
//        longstring creatorDetails; // details set by requester
//
//        //: Reserved for future use
//        EmptyExt ext;
//    };
//
type DestructionOp struct {
	SecurityType   Uint32     `json:"securityType,omitempty"`
	Balance        BalanceId  `json:"balance,omitempty"`
	Amount         Uint64     `json:"amount,omitempty"`
	Reference      Longstring `json:"reference,omitempty"`
	Fee            Fee        `json:"fee,omitempty"`
	CreatorDetails Longstring `json:"creatorDetails,omitempty"`
	Ext            EmptyExt   `json:"ext,omitempty"`
}

// DestructionResultCode is an XDR Enum defines as:
//
//   //: Destruction operation result codes
//    enum DestructionResultCode
//    {
//        // codes considered as "success" for the operation
//        //: Destruction operation successfully applied
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: Trying to create a withdrawal with a 0 amount
//        INVALID_AMOUNT = -1,
//        //: Creator details are not in a valid JSON format
//        INVALID_CREATOR_DETAILS = -2,
//        //: Source balance to withdraw from is not found
//        BALANCE_NOT_FOUND = -3, // balance not found
//        //: Asset cannot be withdrawn because AssetPolicy::WITHDRAWABLE is not set
//        ASSET_IS_NOT_WITHDRAWABLE = -4,
//        //: Expected fee and actual fee mismatch
//        FEE_MISMATCHED = -5,
//        //: Trying to lock balance, locked amount would exceed UINT64_MAX
//        BALANCE_LOCK_OVERFLOW = -6,
//        //: Insufficient balance to withdraw the requested amount
//        UNDERFUNDED = -7,
//        //: Applying operation would overflow statistics
//        STATS_OVERFLOW = -8,
//        //: Applying operation would exceed limits set in the system
//        LIMITS_EXCEEDED = -9,
//        //: Amount withdrawn is smaller than the minimal withdrawable amount set in the system
//        LOWER_BOUND_NOT_EXCEEDED = -10,
//        REFERENCE_DUPLICATION = -11
//    };
//
type DestructionResultCode int32

const (
	DestructionResultCodeSuccess                DestructionResultCode = 0
	DestructionResultCodeInvalidAmount          DestructionResultCode = -1
	DestructionResultCodeInvalidCreatorDetails  DestructionResultCode = -2
	DestructionResultCodeBalanceNotFound        DestructionResultCode = -3
	DestructionResultCodeAssetIsNotWithdrawable DestructionResultCode = -4
	DestructionResultCodeFeeMismatched          DestructionResultCode = -5
	DestructionResultCodeBalanceLockOverflow    DestructionResultCode = -6
	DestructionResultCodeUnderfunded            DestructionResultCode = -7
	DestructionResultCodeStatsOverflow          DestructionResultCode = -8
	DestructionResultCodeLimitsExceeded         DestructionResultCode = -9
	DestructionResultCodeLowerBoundNotExceeded  DestructionResultCode = -10
	DestructionResultCodeReferenceDuplication   DestructionResultCode = -11
)

var DestructionResultCodeAll = []DestructionResultCode{
	DestructionResultCodeSuccess,
	DestructionResultCodeInvalidAmount,
	DestructionResultCodeInvalidCreatorDetails,
	DestructionResultCodeBalanceNotFound,
	DestructionResultCodeAssetIsNotWithdrawable,
	DestructionResultCodeFeeMismatched,
	DestructionResultCodeBalanceLockOverflow,
	DestructionResultCodeUnderfunded,
	DestructionResultCodeStatsOverflow,
	DestructionResultCodeLimitsExceeded,
	DestructionResultCodeLowerBoundNotExceeded,
	DestructionResultCodeReferenceDuplication,
}

var destructionResultCodeMap = map[int32]string{
	0:   "DestructionResultCodeSuccess",
	-1:  "DestructionResultCodeInvalidAmount",
	-2:  "DestructionResultCodeInvalidCreatorDetails",
	-3:  "DestructionResultCodeBalanceNotFound",
	-4:  "DestructionResultCodeAssetIsNotWithdrawable",
	-5:  "DestructionResultCodeFeeMismatched",
	-6:  "DestructionResultCodeBalanceLockOverflow",
	-7:  "DestructionResultCodeUnderfunded",
	-8:  "DestructionResultCodeStatsOverflow",
	-9:  "DestructionResultCodeLimitsExceeded",
	-10: "DestructionResultCodeLowerBoundNotExceeded",
	-11: "DestructionResultCodeReferenceDuplication",
}

var destructionResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "invalid_amount",
	-2:  "invalid_creator_details",
	-3:  "balance_not_found",
	-4:  "asset_is_not_withdrawable",
	-5:  "fee_mismatched",
	-6:  "balance_lock_overflow",
	-7:  "underfunded",
	-8:  "stats_overflow",
	-9:  "limits_exceeded",
	-10: "lower_bound_not_exceeded",
	-11: "reference_duplication",
}

var destructionResultCodeRevMap = map[string]int32{
	"DestructionResultCodeSuccess":                0,
	"DestructionResultCodeInvalidAmount":          -1,
	"DestructionResultCodeInvalidCreatorDetails":  -2,
	"DestructionResultCodeBalanceNotFound":        -3,
	"DestructionResultCodeAssetIsNotWithdrawable": -4,
	"DestructionResultCodeFeeMismatched":          -5,
	"DestructionResultCodeBalanceLockOverflow":    -6,
	"DestructionResultCodeUnderfunded":            -7,
	"DestructionResultCodeStatsOverflow":          -8,
	"DestructionResultCodeLimitsExceeded":         -9,
	"DestructionResultCodeLowerBoundNotExceeded":  -10,
	"DestructionResultCodeReferenceDuplication":   -11,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for DestructionResultCode
func (e DestructionResultCode) ValidEnum(v int32) bool {
	_, ok := destructionResultCodeMap[v]
	return ok
}
func (e DestructionResultCode) isFlag() bool {
	for i := len(DestructionResultCodeAll) - 1; i >= 0; i-- {
		expected := DestructionResultCode(2) << uint64(len(DestructionResultCodeAll)-1) >> uint64(len(DestructionResultCodeAll)-i)
		if expected != DestructionResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e DestructionResultCode) String() string {
	name, _ := destructionResultCodeMap[int32(e)]
	return name
}

func (e DestructionResultCode) ShortString() string {
	name, _ := destructionResultCodeShortMap[int32(e)]
	return name
}

func (e DestructionResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range DestructionResultCodeAll {
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

func (e *DestructionResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = DestructionResultCode(t.Value)
	return nil
}

// DestructionSuccess is an XDR Struct defines as:
//
//   //: Result of the successful withdrawal request creation
//    struct DestructionSuccess {
//        //: Account address of the receiver
//        AccountID targetAccount;
//        BalanceID targetBalance;
//
//        uint64 actualAmount;
//        //: Paid fee
//        Fee fee;
//
//        EmptyExt ext;
//    };
//
type DestructionSuccess struct {
	TargetAccount AccountId `json:"targetAccount,omitempty"`
	TargetBalance BalanceId `json:"targetBalance,omitempty"`
	ActualAmount  Uint64    `json:"actualAmount,omitempty"`
	Fee           Fee       `json:"fee,omitempty"`
	Ext           EmptyExt  `json:"ext,omitempty"`
}

// DestructionResult is an XDR Union defines as:
//
//   //: Result of applying Destruction operation along with the result code
//    union DestructionResult switch (DestructionResultCode code)
//    {
//        case SUCCESS:
//            DestructionSuccess success;
//        default:
//            void;
//    };
//
type DestructionResult struct {
	Code    DestructionResultCode `json:"code,omitempty"`
	Success *DestructionSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u DestructionResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of DestructionResult
func (u DestructionResult) ArmForSwitch(sw int32) (string, bool) {
	switch DestructionResultCode(sw) {
	case DestructionResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewDestructionResult creates a new  DestructionResult.
func NewDestructionResult(code DestructionResultCode, value interface{}) (result DestructionResult, err error) {
	result.Code = code
	switch DestructionResultCode(code) {
	case DestructionResultCodeSuccess:
		tv, ok := value.(DestructionSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be DestructionSuccess")
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
func (u DestructionResult) MustSuccess() DestructionSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u DestructionResult) GetSuccess() (result DestructionSuccess, ok bool) {
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

// IssuanceOp is an XDR Struct defines as:
//
//   //: IssuanceOp is used to issuance specified amount of asset to a receiver's balance
//    struct IssuanceOp
//    {
//        //: security type
//        uint32 securityType;
//
//        //: Code of an asset to issuance
//        AssetCode asset;
//        //: Amount to issuance
//        uint64 amount;
//
//        MovementDestination destination;
//
//        longstring reference;
//
//        //: Arbitrary stringified json object that can be used to attach data to be reviewed by an admin
//        longstring creatorDetails; // details of the issuance (External system id, etc.)
//        //: Total fee to pay, consists of fixed fee and percent fee, calculated automatically
//        Fee fee; //totalFee to be payed (calculated automatically)
//        //: Reserved for future use
//        EmptyExt ext;
//    };
//
type IssuanceOp struct {
	SecurityType   Uint32              `json:"securityType,omitempty"`
	Asset          AssetCode           `json:"asset,omitempty"`
	Amount         Uint64              `json:"amount,omitempty"`
	Destination    MovementDestination `json:"destination,omitempty"`
	Reference      Longstring          `json:"reference,omitempty"`
	CreatorDetails Longstring          `json:"creatorDetails,omitempty"`
	Fee            Fee                 `json:"fee,omitempty"`
	Ext            EmptyExt            `json:"ext,omitempty"`
}

// IssuanceResultCode is an XDR Enum defines as:
//
//   //: Result codes of the IssuanceOp
//    enum IssuanceResultCode
//    {
//        // codes considered as "success" for the operation
//        //: Issuance operation application was successful
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: Asset to issuance is not found
//        ASSET_NOT_FOUND = -1,
//        //: Trying to create an issuance request with negative/zero amount
//        INVALID_AMOUNT = -2,
//        //: Either the target balance is not found or there is a mismatch between the target balance asset and an asset in the request
//        NO_COUNTERPARTY = -4,
//        //: Source of operation is not an owner of the asset
//        NOT_AUTHORIZED = -5,
//        //: Issuanced amount plus amount to issuance will exceed max issuance amount
//        EXCEEDS_MAX_ISSUANCE_AMOUNT = -6,
//        //: Amount to issuance plus amount on balance would exceed UINT64_MAX
//        RECEIVER_FULL_LINE = -7,
//        //: Creator details are not valid JSON
//        INVALID_CREATOR_DETAILS = -8,
//        //: Fee is greater than the amount to issuance
//        FEE_EXCEEDS_AMOUNT = -9,
//        INVALID_AMOUNT_PRECISION = -10
//    };
//
type IssuanceResultCode int32

const (
	IssuanceResultCodeSuccess                  IssuanceResultCode = 0
	IssuanceResultCodeAssetNotFound            IssuanceResultCode = -1
	IssuanceResultCodeInvalidAmount            IssuanceResultCode = -2
	IssuanceResultCodeNoCounterparty           IssuanceResultCode = -4
	IssuanceResultCodeNotAuthorized            IssuanceResultCode = -5
	IssuanceResultCodeExceedsMaxIssuanceAmount IssuanceResultCode = -6
	IssuanceResultCodeReceiverFullLine         IssuanceResultCode = -7
	IssuanceResultCodeInvalidCreatorDetails    IssuanceResultCode = -8
	IssuanceResultCodeFeeExceedsAmount         IssuanceResultCode = -9
	IssuanceResultCodeInvalidAmountPrecision   IssuanceResultCode = -10
)

var IssuanceResultCodeAll = []IssuanceResultCode{
	IssuanceResultCodeSuccess,
	IssuanceResultCodeAssetNotFound,
	IssuanceResultCodeInvalidAmount,
	IssuanceResultCodeNoCounterparty,
	IssuanceResultCodeNotAuthorized,
	IssuanceResultCodeExceedsMaxIssuanceAmount,
	IssuanceResultCodeReceiverFullLine,
	IssuanceResultCodeInvalidCreatorDetails,
	IssuanceResultCodeFeeExceedsAmount,
	IssuanceResultCodeInvalidAmountPrecision,
}

var issuanceResultCodeMap = map[int32]string{
	0:   "IssuanceResultCodeSuccess",
	-1:  "IssuanceResultCodeAssetNotFound",
	-2:  "IssuanceResultCodeInvalidAmount",
	-4:  "IssuanceResultCodeNoCounterparty",
	-5:  "IssuanceResultCodeNotAuthorized",
	-6:  "IssuanceResultCodeExceedsMaxIssuanceAmount",
	-7:  "IssuanceResultCodeReceiverFullLine",
	-8:  "IssuanceResultCodeInvalidCreatorDetails",
	-9:  "IssuanceResultCodeFeeExceedsAmount",
	-10: "IssuanceResultCodeInvalidAmountPrecision",
}

var issuanceResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "asset_not_found",
	-2:  "invalid_amount",
	-4:  "no_counterparty",
	-5:  "not_authorized",
	-6:  "exceeds_max_issuance_amount",
	-7:  "receiver_full_line",
	-8:  "invalid_creator_details",
	-9:  "fee_exceeds_amount",
	-10: "invalid_amount_precision",
}

var issuanceResultCodeRevMap = map[string]int32{
	"IssuanceResultCodeSuccess":                  0,
	"IssuanceResultCodeAssetNotFound":            -1,
	"IssuanceResultCodeInvalidAmount":            -2,
	"IssuanceResultCodeNoCounterparty":           -4,
	"IssuanceResultCodeNotAuthorized":            -5,
	"IssuanceResultCodeExceedsMaxIssuanceAmount": -6,
	"IssuanceResultCodeReceiverFullLine":         -7,
	"IssuanceResultCodeInvalidCreatorDetails":    -8,
	"IssuanceResultCodeFeeExceedsAmount":         -9,
	"IssuanceResultCodeInvalidAmountPrecision":   -10,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for IssuanceResultCode
func (e IssuanceResultCode) ValidEnum(v int32) bool {
	_, ok := issuanceResultCodeMap[v]
	return ok
}
func (e IssuanceResultCode) isFlag() bool {
	for i := len(IssuanceResultCodeAll) - 1; i >= 0; i-- {
		expected := IssuanceResultCode(2) << uint64(len(IssuanceResultCodeAll)-1) >> uint64(len(IssuanceResultCodeAll)-i)
		if expected != IssuanceResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e IssuanceResultCode) String() string {
	name, _ := issuanceResultCodeMap[int32(e)]
	return name
}

func (e IssuanceResultCode) ShortString() string {
	name, _ := issuanceResultCodeShortMap[int32(e)]
	return name
}

func (e IssuanceResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range IssuanceResultCodeAll {
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

func (e *IssuanceResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = IssuanceResultCode(t.Value)
	return nil
}

// IssuanceSuccess is an XDR Struct defines as:
//
//   //:Result of successful application of Issuance operation
//    struct IssuanceSuccess {
//
//        //: Account address of the receiver
//        AccountID receiver;
//        BalanceID receiverBalance;
//
//        //: Paid fee
//        Fee fee;
//        //: Reserved for future use
//        EmptyExt ext;
//    };
//
type IssuanceSuccess struct {
	Receiver        AccountId `json:"receiver,omitempty"`
	ReceiverBalance BalanceId `json:"receiverBalance,omitempty"`
	Fee             Fee       `json:"fee,omitempty"`
	Ext             EmptyExt  `json:"ext,omitempty"`
}

// IssuanceResult is an XDR Union defines as:
//
//   //: Create issuance request result with result code
//    union IssuanceResult switch (IssuanceResultCode code)
//    {
//    case SUCCESS:
//        IssuanceSuccess success;
//    default:
//        void;
//    };
//
type IssuanceResult struct {
	Code    IssuanceResultCode `json:"code,omitempty"`
	Success *IssuanceSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u IssuanceResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of IssuanceResult
func (u IssuanceResult) ArmForSwitch(sw int32) (string, bool) {
	switch IssuanceResultCode(sw) {
	case IssuanceResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewIssuanceResult creates a new  IssuanceResult.
func NewIssuanceResult(code IssuanceResultCode, value interface{}) (result IssuanceResult, err error) {
	result.Code = code
	switch IssuanceResultCode(code) {
	case IssuanceResultCodeSuccess:
		tv, ok := value.(IssuanceSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be IssuanceSuccess")
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
func (u IssuanceResult) MustSuccess() IssuanceSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u IssuanceResult) GetSuccess() (result IssuanceSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// KycRecoveryOp is an XDR Struct defines as:
//
//   //: KYCRecoveryOp to create KYC recovery request and set new signers for account
//    struct KYCRecoveryOp
//    {
//        //: Account for which signers will be set
//        AccountID targetAccount;
//        //: New signers to set
//        SignerData signersData<>;
//
//         //: Arbitrary stringified json object that can be used to attach data to be reviewed by an admin
//        longstring creatorDetails; // details set by requester
//
//        EmptyExt ext;
//    };
//
type KycRecoveryOp struct {
	TargetAccount  AccountId    `json:"targetAccount,omitempty"`
	SignersData    []SignerData `json:"signersData,omitempty"`
	CreatorDetails Longstring   `json:"creatorDetails,omitempty"`
	Ext            EmptyExt     `json:"ext,omitempty"`
}

// KycRecoveryResultCode is an XDR Enum defines as:
//
//   //: Result codes of KYCRecoveryOp
//    enum KYCRecoveryResultCode
//    {
//        //: KYC Recovery request was successfully created
//        SUCCESS = 0,
//
//        //: Creator details are not in a valid JSON format
//        INVALID_CREATOR_DETAILS = -1,
//        //: Not allowed to provide empty slice of signers
//        NO_SIGNER_DATA = -2,
//        //: SignerData contains duplicates
//        SIGNER_DUPLICATION = -3,
//        //: Signer has weight > threshold
//        INVALID_WEIGHT = -4,
//        //: Signer has invalid details
//        INVALID_DETAILS = -5,
//        //: Account with provided account address does not exist
//        TARGET_ACCOUNT_NOT_FOUND = -9,
//        //: System configuration forbids KYC recovery
//        RECOVERY_NOT_ALLOWED = -10
//    };
//
type KycRecoveryResultCode int32

const (
	KycRecoveryResultCodeSuccess               KycRecoveryResultCode = 0
	KycRecoveryResultCodeInvalidCreatorDetails KycRecoveryResultCode = -1
	KycRecoveryResultCodeNoSignerData          KycRecoveryResultCode = -2
	KycRecoveryResultCodeSignerDuplication     KycRecoveryResultCode = -3
	KycRecoveryResultCodeInvalidWeight         KycRecoveryResultCode = -4
	KycRecoveryResultCodeInvalidDetails        KycRecoveryResultCode = -5
	KycRecoveryResultCodeTargetAccountNotFound KycRecoveryResultCode = -9
	KycRecoveryResultCodeRecoveryNotAllowed    KycRecoveryResultCode = -10
)

var KycRecoveryResultCodeAll = []KycRecoveryResultCode{
	KycRecoveryResultCodeSuccess,
	KycRecoveryResultCodeInvalidCreatorDetails,
	KycRecoveryResultCodeNoSignerData,
	KycRecoveryResultCodeSignerDuplication,
	KycRecoveryResultCodeInvalidWeight,
	KycRecoveryResultCodeInvalidDetails,
	KycRecoveryResultCodeTargetAccountNotFound,
	KycRecoveryResultCodeRecoveryNotAllowed,
}

var kycRecoveryResultCodeMap = map[int32]string{
	0:   "KycRecoveryResultCodeSuccess",
	-1:  "KycRecoveryResultCodeInvalidCreatorDetails",
	-2:  "KycRecoveryResultCodeNoSignerData",
	-3:  "KycRecoveryResultCodeSignerDuplication",
	-4:  "KycRecoveryResultCodeInvalidWeight",
	-5:  "KycRecoveryResultCodeInvalidDetails",
	-9:  "KycRecoveryResultCodeTargetAccountNotFound",
	-10: "KycRecoveryResultCodeRecoveryNotAllowed",
}

var kycRecoveryResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "invalid_creator_details",
	-2:  "no_signer_data",
	-3:  "signer_duplication",
	-4:  "invalid_weight",
	-5:  "invalid_details",
	-9:  "target_account_not_found",
	-10: "recovery_not_allowed",
}

var kycRecoveryResultCodeRevMap = map[string]int32{
	"KycRecoveryResultCodeSuccess":               0,
	"KycRecoveryResultCodeInvalidCreatorDetails": -1,
	"KycRecoveryResultCodeNoSignerData":          -2,
	"KycRecoveryResultCodeSignerDuplication":     -3,
	"KycRecoveryResultCodeInvalidWeight":         -4,
	"KycRecoveryResultCodeInvalidDetails":        -5,
	"KycRecoveryResultCodeTargetAccountNotFound": -9,
	"KycRecoveryResultCodeRecoveryNotAllowed":    -10,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for KycRecoveryResultCode
func (e KycRecoveryResultCode) ValidEnum(v int32) bool {
	_, ok := kycRecoveryResultCodeMap[v]
	return ok
}
func (e KycRecoveryResultCode) isFlag() bool {
	for i := len(KycRecoveryResultCodeAll) - 1; i >= 0; i-- {
		expected := KycRecoveryResultCode(2) << uint64(len(KycRecoveryResultCodeAll)-1) >> uint64(len(KycRecoveryResultCodeAll)-i)
		if expected != KycRecoveryResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e KycRecoveryResultCode) String() string {
	name, _ := kycRecoveryResultCodeMap[int32(e)]
	return name
}

func (e KycRecoveryResultCode) ShortString() string {
	name, _ := kycRecoveryResultCodeShortMap[int32(e)]
	return name
}

func (e KycRecoveryResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range KycRecoveryResultCodeAll {
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

func (e *KycRecoveryResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = KycRecoveryResultCode(t.Value)
	return nil
}

// KycRecoveryResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//            //: reserved for future use
//            EmptyExt ext;
//        }
//
type KycRecoveryResultSuccess struct {
	Ext EmptyExt `json:"ext,omitempty"`
}

// KycRecoveryResult is an XDR Union defines as:
//
//   //: Result of operation applying
//    union KYCRecoveryResult switch (KYCRecoveryResultCode code)
//    {
//    case SUCCESS:
//        //: Is used to pass useful params if operation is success
//        struct {
//            //: reserved for future use
//            EmptyExt ext;
//        } success;
//    default:
//        void;
//    };
//
type KycRecoveryResult struct {
	Code    KycRecoveryResultCode     `json:"code,omitempty"`
	Success *KycRecoveryResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u KycRecoveryResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of KycRecoveryResult
func (u KycRecoveryResult) ArmForSwitch(sw int32) (string, bool) {
	switch KycRecoveryResultCode(sw) {
	case KycRecoveryResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewKycRecoveryResult creates a new  KycRecoveryResult.
func NewKycRecoveryResult(code KycRecoveryResultCode, value interface{}) (result KycRecoveryResult, err error) {
	result.Code = code
	switch KycRecoveryResultCode(code) {
	case KycRecoveryResultCodeSuccess:
		tv, ok := value.(KycRecoveryResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be KycRecoveryResultSuccess")
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
func (u KycRecoveryResult) MustSuccess() KycRecoveryResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u KycRecoveryResult) GetSuccess() (result KycRecoveryResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// PaymentFeeDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PaymentFeeDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentFeeDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentFeeDataExt
func (u PaymentFeeDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentFeeDataExt creates a new  PaymentFeeDataExt.
func NewPaymentFeeDataExt(v LedgerVersion, value interface{}) (result PaymentFeeDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentFeeData is an XDR Struct defines as:
//
//   struct PaymentFeeData {
//        //: Fee to pay by source balance
//        Fee sourceFee;
//        //: Fee kept from destination account/balance
//        Fee destinationFee;
//        //: Indicates whether or not the source of payment pays the destination fee
//        bool sourcePaysForDest;
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
type PaymentFeeData struct {
	SourceFee         Fee               `json:"sourceFee,omitempty"`
	DestinationFee    Fee               `json:"destinationFee,omitempty"`
	SourcePaysForDest bool              `json:"sourcePaysForDest,omitempty"`
	Ext               PaymentFeeDataExt `json:"ext,omitempty"`
}

// PaymentOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PaymentOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentOpExt
func (u PaymentOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentOpExt creates a new  PaymentOpExt.
func NewPaymentOpExt(v LedgerVersion, value interface{}) (result PaymentOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentOp is an XDR Struct defines as:
//
//   //: PaymentOp is used to transfer some amount of asset from the source balance to destination account/balance
//    struct PaymentOp
//    {
//        //: ID of the source balance of payment
//        BalanceID sourceBalanceID;
//
//        uint32 securityType;
//
//        //: `destination` defines the type of instance that receives the payment based on given PaymentDestinationType
//        MovementDestination destination;
//
//        //: Amount of payment
//        uint64 amount;
//
//        //: `feeData` defines all data about the payment fee
//        PaymentFeeData feeData;
//
//        //: `subject` is a user-provided info about the real-life purpose of payment
//        longstring subject;
//        //: `reference` is a string formed by a payment sender. `Reference-sender account` pair is unique.
//        longstring reference;
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
type PaymentOp struct {
	SourceBalanceId BalanceId           `json:"sourceBalanceID,omitempty"`
	SecurityType    Uint32              `json:"securityType,omitempty"`
	Destination     MovementDestination `json:"destination,omitempty"`
	Amount          Uint64              `json:"amount,omitempty"`
	FeeData         PaymentFeeData      `json:"feeData,omitempty"`
	Subject         Longstring          `json:"subject,omitempty"`
	Reference       Longstring          `json:"reference,omitempty"`
	Ext             PaymentOpExt        `json:"ext,omitempty"`
}

// PaymentResultCode is an XDR Enum defines as:
//
//   enum PaymentResultCode
//    {
//        // codes considered as "success" for the operation
//        //: Payment was successfully completed
//        SUCCESS = 0, // payment successfully completed
//
//        // codes considered as "failure" for the operation
//        //: Payment sender balance ID and payment receiver balance ID are equal or reference is longer than 64 symbols
//        MALFORMED = -1,
//        //: Not enough funds in the source account
//        UNDERFUNDED = -2,
//        //: After the payment fulfillment, the destination balance will exceed the limit (total amount on the balance will be greater than UINT64_MAX)
//        LINE_FULL = -3,
//        //: There is no balance found with an ID provided in `destinations.balanceID`
//        DESTINATION_BALANCE_NOT_FOUND = -4,
//        //: Sender balance asset and receiver balance asset are not equal
//        BALANCE_ASSETS_MISMATCHED = -5,
//        //: There is no balance found with ID provided in `sourceBalanceID`
//        SRC_BALANCE_NOT_FOUND = -6,
//        //: Pair `reference-sender account` of the payment is not unique
//        REFERENCE_DUPLICATION = -7,
//        //: Stats entry exceeded account limits
//        STATS_OVERFLOW = -8,
//        //: Account will exceed its limits after the payment is fulfilled
//        LIMITS_EXCEEDED = -9,
//        //: Payment asset does not have a `TRANSFERABLE` policy set
//        NOT_ALLOWED_BY_ASSET_POLICY = -10,
//        //: Overflow during total fee calculation
//        INVALID_DESTINATION_FEE = -11,
//        //: Payment fee amount is insufficient
//        INSUFFICIENT_FEE_AMOUNT = -12,
//        //: Fee charged from destination balance is greater than the payment amount
//        PAYMENT_AMOUNT_IS_LESS_THAN_DEST_FEE = -13,
//        //: There is no account found with an ID provided in `destination.accountID`
//        DESTINATION_ACCOUNT_NOT_FOUND = -14,
//        //: Amount precision and asset precision are mismatched
//        INCORRECT_AMOUNT_PRECISION = -15
//    };
//
type PaymentResultCode int32

const (
	PaymentResultCodeSuccess                        PaymentResultCode = 0
	PaymentResultCodeMalformed                      PaymentResultCode = -1
	PaymentResultCodeUnderfunded                    PaymentResultCode = -2
	PaymentResultCodeLineFull                       PaymentResultCode = -3
	PaymentResultCodeDestinationBalanceNotFound     PaymentResultCode = -4
	PaymentResultCodeBalanceAssetsMismatched        PaymentResultCode = -5
	PaymentResultCodeSrcBalanceNotFound             PaymentResultCode = -6
	PaymentResultCodeReferenceDuplication           PaymentResultCode = -7
	PaymentResultCodeStatsOverflow                  PaymentResultCode = -8
	PaymentResultCodeLimitsExceeded                 PaymentResultCode = -9
	PaymentResultCodeNotAllowedByAssetPolicy        PaymentResultCode = -10
	PaymentResultCodeInvalidDestinationFee          PaymentResultCode = -11
	PaymentResultCodeInsufficientFeeAmount          PaymentResultCode = -12
	PaymentResultCodePaymentAmountIsLessThanDestFee PaymentResultCode = -13
	PaymentResultCodeDestinationAccountNotFound     PaymentResultCode = -14
	PaymentResultCodeIncorrectAmountPrecision       PaymentResultCode = -15
)

var PaymentResultCodeAll = []PaymentResultCode{
	PaymentResultCodeSuccess,
	PaymentResultCodeMalformed,
	PaymentResultCodeUnderfunded,
	PaymentResultCodeLineFull,
	PaymentResultCodeDestinationBalanceNotFound,
	PaymentResultCodeBalanceAssetsMismatched,
	PaymentResultCodeSrcBalanceNotFound,
	PaymentResultCodeReferenceDuplication,
	PaymentResultCodeStatsOverflow,
	PaymentResultCodeLimitsExceeded,
	PaymentResultCodeNotAllowedByAssetPolicy,
	PaymentResultCodeInvalidDestinationFee,
	PaymentResultCodeInsufficientFeeAmount,
	PaymentResultCodePaymentAmountIsLessThanDestFee,
	PaymentResultCodeDestinationAccountNotFound,
	PaymentResultCodeIncorrectAmountPrecision,
}

var paymentResultCodeMap = map[int32]string{
	0:   "PaymentResultCodeSuccess",
	-1:  "PaymentResultCodeMalformed",
	-2:  "PaymentResultCodeUnderfunded",
	-3:  "PaymentResultCodeLineFull",
	-4:  "PaymentResultCodeDestinationBalanceNotFound",
	-5:  "PaymentResultCodeBalanceAssetsMismatched",
	-6:  "PaymentResultCodeSrcBalanceNotFound",
	-7:  "PaymentResultCodeReferenceDuplication",
	-8:  "PaymentResultCodeStatsOverflow",
	-9:  "PaymentResultCodeLimitsExceeded",
	-10: "PaymentResultCodeNotAllowedByAssetPolicy",
	-11: "PaymentResultCodeInvalidDestinationFee",
	-12: "PaymentResultCodeInsufficientFeeAmount",
	-13: "PaymentResultCodePaymentAmountIsLessThanDestFee",
	-14: "PaymentResultCodeDestinationAccountNotFound",
	-15: "PaymentResultCodeIncorrectAmountPrecision",
}

var paymentResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "malformed",
	-2:  "underfunded",
	-3:  "line_full",
	-4:  "destination_balance_not_found",
	-5:  "balance_assets_mismatched",
	-6:  "src_balance_not_found",
	-7:  "reference_duplication",
	-8:  "stats_overflow",
	-9:  "limits_exceeded",
	-10: "not_allowed_by_asset_policy",
	-11: "invalid_destination_fee",
	-12: "insufficient_fee_amount",
	-13: "payment_amount_is_less_than_dest_fee",
	-14: "destination_account_not_found",
	-15: "incorrect_amount_precision",
}

var paymentResultCodeRevMap = map[string]int32{
	"PaymentResultCodeSuccess":                        0,
	"PaymentResultCodeMalformed":                      -1,
	"PaymentResultCodeUnderfunded":                    -2,
	"PaymentResultCodeLineFull":                       -3,
	"PaymentResultCodeDestinationBalanceNotFound":     -4,
	"PaymentResultCodeBalanceAssetsMismatched":        -5,
	"PaymentResultCodeSrcBalanceNotFound":             -6,
	"PaymentResultCodeReferenceDuplication":           -7,
	"PaymentResultCodeStatsOverflow":                  -8,
	"PaymentResultCodeLimitsExceeded":                 -9,
	"PaymentResultCodeNotAllowedByAssetPolicy":        -10,
	"PaymentResultCodeInvalidDestinationFee":          -11,
	"PaymentResultCodeInsufficientFeeAmount":          -12,
	"PaymentResultCodePaymentAmountIsLessThanDestFee": -13,
	"PaymentResultCodeDestinationAccountNotFound":     -14,
	"PaymentResultCodeIncorrectAmountPrecision":       -15,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PaymentResultCode
func (e PaymentResultCode) ValidEnum(v int32) bool {
	_, ok := paymentResultCodeMap[v]
	return ok
}
func (e PaymentResultCode) isFlag() bool {
	for i := len(PaymentResultCodeAll) - 1; i >= 0; i-- {
		expected := PaymentResultCode(2) << uint64(len(PaymentResultCodeAll)-1) >> uint64(len(PaymentResultCodeAll)-i)
		if expected != PaymentResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PaymentResultCode) String() string {
	name, _ := paymentResultCodeMap[int32(e)]
	return name
}

func (e PaymentResultCode) ShortString() string {
	name, _ := paymentResultCodeShortMap[int32(e)]
	return name
}

func (e PaymentResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range PaymentResultCodeAll {
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

func (e *PaymentResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PaymentResultCode(t.Value)
	return nil
}

// PaymentSuccess is an XDR Struct defines as:
//
//   //: `PaymentResponse` defines the response on the corresponding PaymentOp
//    struct PaymentSuccess {
//        //: ID of the destination account
//        AccountID destination;
//        //: ID of the destination balance
//        BalanceID destinationBalanceID;
//
//        //: Code of an asset used in payment
//        AssetCode asset;
//        //: Actual amount received
//        uint64 amountReceived;
//
//        //: Fee charged from the source balance
//        Fee actualSourcePaymentFee;
//        //: Fee charged from the destination balance
//        Fee actualDestinationPaymentFee;
//
//        //: reserved for future use
//        EmptyExt ext;
//    };
//
type PaymentSuccess struct {
	Destination                 AccountId `json:"destination,omitempty"`
	DestinationBalanceId        BalanceId `json:"destinationBalanceID,omitempty"`
	Asset                       AssetCode `json:"asset,omitempty"`
	AmountReceived              Uint64    `json:"amountReceived,omitempty"`
	ActualSourcePaymentFee      Fee       `json:"actualSourcePaymentFee,omitempty"`
	ActualDestinationPaymentFee Fee       `json:"actualDestinationPaymentFee,omitempty"`
	Ext                         EmptyExt  `json:"ext,omitempty"`
}

// PaymentResult is an XDR Union defines as:
//
//   union PaymentResult switch (PaymentResultCode code)
//    {
//    case SUCCESS:
//        PaymentSuccess paymentSuccess;
//    default:
//        void;
//    };
//
type PaymentResult struct {
	Code           PaymentResultCode `json:"code,omitempty"`
	PaymentSuccess *PaymentSuccess   `json:"paymentSuccess,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentResult
func (u PaymentResult) ArmForSwitch(sw int32) (string, bool) {
	switch PaymentResultCode(sw) {
	case PaymentResultCodeSuccess:
		return "PaymentSuccess", true
	default:
		return "", true
	}
}

// NewPaymentResult creates a new  PaymentResult.
func NewPaymentResult(code PaymentResultCode, value interface{}) (result PaymentResult, err error) {
	result.Code = code
	switch PaymentResultCode(code) {
	case PaymentResultCodeSuccess:
		tv, ok := value.(PaymentSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentSuccess")
			return
		}
		result.PaymentSuccess = &tv
	default:
		// void
	}
	return
}

// MustPaymentSuccess retrieves the PaymentSuccess value from the union,
// panicing if the value is not set.
func (u PaymentResult) MustPaymentSuccess() PaymentSuccess {
	val, ok := u.GetPaymentSuccess()

	if !ok {
		panic("arm PaymentSuccess is not set")
	}

	return val
}

// GetPaymentSuccess retrieves the PaymentSuccess value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PaymentResult) GetPaymentSuccess() (result PaymentSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "PaymentSuccess" {
		result = *u.PaymentSuccess
		ok = true
	}

	return
}

// PutKeyValueOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type PutKeyValueOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PutKeyValueOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PutKeyValueOpExt
func (u PutKeyValueOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPutKeyValueOpExt creates a new  PutKeyValueOpExt.
func NewPutKeyValueOpExt(v LedgerVersion, value interface{}) (result PutKeyValueOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PutKeyValueOp is an XDR Struct defines as:
//
//   //: `PutKeyValueOp` is used to update the key-value entry present in the system
//        struct PutKeyValueOp
//        {
//            //: `key` is the key for KeyValueEntry
//            longstring key;
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
type PutKeyValueOp struct {
	Key   Longstring         `json:"key,omitempty"`
	Value KeyValueEntryValue `json:"value,omitempty"`
	Ext   PutKeyValueOpExt   `json:"ext,omitempty"`
}

// PutKeyValueSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type PutKeyValueSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PutKeyValueSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PutKeyValueSuccessExt
func (u PutKeyValueSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPutKeyValueSuccessExt creates a new  PutKeyValueSuccessExt.
func NewPutKeyValueSuccessExt(v LedgerVersion, value interface{}) (result PutKeyValueSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PutKeyValueSuccess is an XDR Struct defines as:
//
//   //: `PutKeyValueSuccess` represents details returned after the successful application of `PutKeyValueOp`
//        struct PutKeyValueSuccess
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
type PutKeyValueSuccess struct {
	Ext PutKeyValueSuccessExt `json:"ext,omitempty"`
}

// PutKeyValueResultCode is an XDR Enum defines as:
//
//   //: Result codes for `PutKeyValueOp`
//        enum PutKeyValueResultCode
//        {
//            //: `PutKeyValueOp` is applied successfully
//            SUCCESS = 0,
//            //: Value type of the key-value entry is forbidden for the provided key
//            INVALID_TYPE = -1,
//            //: value is forbidden for the provided key
//            VALUE_NOT_ALLOWED = -2
//        };
//
type PutKeyValueResultCode int32

const (
	PutKeyValueResultCodeSuccess         PutKeyValueResultCode = 0
	PutKeyValueResultCodeInvalidType     PutKeyValueResultCode = -1
	PutKeyValueResultCodeValueNotAllowed PutKeyValueResultCode = -2
)

var PutKeyValueResultCodeAll = []PutKeyValueResultCode{
	PutKeyValueResultCodeSuccess,
	PutKeyValueResultCodeInvalidType,
	PutKeyValueResultCodeValueNotAllowed,
}

var putKeyValueResultCodeMap = map[int32]string{
	0:  "PutKeyValueResultCodeSuccess",
	-1: "PutKeyValueResultCodeInvalidType",
	-2: "PutKeyValueResultCodeValueNotAllowed",
}

var putKeyValueResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_type",
	-2: "value_not_allowed",
}

var putKeyValueResultCodeRevMap = map[string]int32{
	"PutKeyValueResultCodeSuccess":         0,
	"PutKeyValueResultCodeInvalidType":     -1,
	"PutKeyValueResultCodeValueNotAllowed": -2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PutKeyValueResultCode
func (e PutKeyValueResultCode) ValidEnum(v int32) bool {
	_, ok := putKeyValueResultCodeMap[v]
	return ok
}
func (e PutKeyValueResultCode) isFlag() bool {
	for i := len(PutKeyValueResultCodeAll) - 1; i >= 0; i-- {
		expected := PutKeyValueResultCode(2) << uint64(len(PutKeyValueResultCodeAll)-1) >> uint64(len(PutKeyValueResultCodeAll)-i)
		if expected != PutKeyValueResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PutKeyValueResultCode) String() string {
	name, _ := putKeyValueResultCodeMap[int32(e)]
	return name
}

func (e PutKeyValueResultCode) ShortString() string {
	name, _ := putKeyValueResultCodeShortMap[int32(e)]
	return name
}

func (e PutKeyValueResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range PutKeyValueResultCodeAll {
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

func (e *PutKeyValueResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PutKeyValueResultCode(t.Value)
	return nil
}

// PutKeyValueResult is an XDR Union defines as:
//
//   //: `PutKeyValueResult` represents the result of PutKeyValueOp
//        union PutKeyValueResult switch (PutKeyValueResultCode code)
//        {
//            case SUCCESS:
//                PutKeyValueSuccess success;
//            default:
//                void;
//        };
//
type PutKeyValueResult struct {
	Code    PutKeyValueResultCode `json:"code,omitempty"`
	Success *PutKeyValueSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PutKeyValueResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PutKeyValueResult
func (u PutKeyValueResult) ArmForSwitch(sw int32) (string, bool) {
	switch PutKeyValueResultCode(sw) {
	case PutKeyValueResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewPutKeyValueResult creates a new  PutKeyValueResult.
func NewPutKeyValueResult(code PutKeyValueResultCode, value interface{}) (result PutKeyValueResult, err error) {
	result.Code = code
	switch PutKeyValueResultCode(code) {
	case PutKeyValueResultCodeSuccess:
		tv, ok := value.(PutKeyValueSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be PutKeyValueSuccess")
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
func (u PutKeyValueResult) MustSuccess() PutKeyValueSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PutKeyValueResult) GetSuccess() (result PutKeyValueSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// RemoveDataOp is an XDR Struct defines as:
//
//   struct RemoveDataOp
//    {
//        uint64 dataID;
//
//        EmptyExt ext;
//    };
//
type RemoveDataOp struct {
	DataId Uint64   `json:"dataID,omitempty"`
	Ext    EmptyExt `json:"ext,omitempty"`
}

// RemoveDataResultCode is an XDR Enum defines as:
//
//   enum RemoveDataResultCode
//    {
//        SUCCESS = 0,
//
//        NOT_FOUND = -1,
//        NOT_ALLOWED = -2
//    };
//
type RemoveDataResultCode int32

const (
	RemoveDataResultCodeSuccess    RemoveDataResultCode = 0
	RemoveDataResultCodeNotFound   RemoveDataResultCode = -1
	RemoveDataResultCodeNotAllowed RemoveDataResultCode = -2
)

var RemoveDataResultCodeAll = []RemoveDataResultCode{
	RemoveDataResultCodeSuccess,
	RemoveDataResultCodeNotFound,
	RemoveDataResultCodeNotAllowed,
}

var removeDataResultCodeMap = map[int32]string{
	0:  "RemoveDataResultCodeSuccess",
	-1: "RemoveDataResultCodeNotFound",
	-2: "RemoveDataResultCodeNotAllowed",
}

var removeDataResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "not_allowed",
}

var removeDataResultCodeRevMap = map[string]int32{
	"RemoveDataResultCodeSuccess":    0,
	"RemoveDataResultCodeNotFound":   -1,
	"RemoveDataResultCodeNotAllowed": -2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RemoveDataResultCode
func (e RemoveDataResultCode) ValidEnum(v int32) bool {
	_, ok := removeDataResultCodeMap[v]
	return ok
}
func (e RemoveDataResultCode) isFlag() bool {
	for i := len(RemoveDataResultCodeAll) - 1; i >= 0; i-- {
		expected := RemoveDataResultCode(2) << uint64(len(RemoveDataResultCodeAll)-1) >> uint64(len(RemoveDataResultCodeAll)-i)
		if expected != RemoveDataResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RemoveDataResultCode) String() string {
	name, _ := removeDataResultCodeMap[int32(e)]
	return name
}

func (e RemoveDataResultCode) ShortString() string {
	name, _ := removeDataResultCodeShortMap[int32(e)]
	return name
}

func (e RemoveDataResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range RemoveDataResultCodeAll {
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

func (e *RemoveDataResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RemoveDataResultCode(t.Value)
	return nil
}

// RemoveDataResult is an XDR Union defines as:
//
//   union RemoveDataResult switch (RemoveDataResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    default:
//        void;
//    };
//
type RemoveDataResult struct {
	Code RemoveDataResultCode `json:"code,omitempty"`
	Ext  *EmptyExt            `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveDataResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveDataResult
func (u RemoveDataResult) ArmForSwitch(sw int32) (string, bool) {
	switch RemoveDataResultCode(sw) {
	case RemoveDataResultCodeSuccess:
		return "Ext", true
	default:
		return "", true
	}
}

// NewRemoveDataResult creates a new  RemoveDataResult.
func NewRemoveDataResult(code RemoveDataResultCode, value interface{}) (result RemoveDataResult, err error) {
	result.Code = code
	switch RemoveDataResultCode(code) {
	case RemoveDataResultCodeSuccess:
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
func (u RemoveDataResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RemoveDataResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// RemoveKeyValueOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type RemoveKeyValueOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveKeyValueOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveKeyValueOpExt
func (u RemoveKeyValueOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRemoveKeyValueOpExt creates a new  RemoveKeyValueOpExt.
func NewRemoveKeyValueOpExt(v LedgerVersion, value interface{}) (result RemoveKeyValueOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RemoveKeyValueOp is an XDR Struct defines as:
//
//   //: `RemoveKeyValueOp` is used to remove key-value entry present in the system by key
//        struct RemoveKeyValueOp
//        {
//            //: `key` is the key for KeyValueEntry
//            longstring key;
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
type RemoveKeyValueOp struct {
	Key Longstring          `json:"key,omitempty"`
	Ext RemoveKeyValueOpExt `json:"ext,omitempty"`
}

// RemoveKeyValueSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type RemoveKeyValueSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveKeyValueSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveKeyValueSuccessExt
func (u RemoveKeyValueSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRemoveKeyValueSuccessExt creates a new  RemoveKeyValueSuccessExt.
func NewRemoveKeyValueSuccessExt(v LedgerVersion, value interface{}) (result RemoveKeyValueSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RemoveKeyValueSuccess is an XDR Struct defines as:
//
//   //: `RemoveKeyValueSuccess` represents details returned after the successful application of `RemoveKeyValueOp`
//        struct RemoveKeyValueSuccess
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
type RemoveKeyValueSuccess struct {
	Ext RemoveKeyValueSuccessExt `json:"ext,omitempty"`
}

// RemoveKeyValueResultCode is an XDR Enum defines as:
//
//   //: Result codes for `RemoveKeyValueOp`
//        enum RemoveKeyValueResultCode
//        {
//            //: `RemoveKeyValueOp` is applied successfully
//            SUCCESS = 0,
//            //: There is no key value with such key
//            NOT_FOUND = -1
//        };
//
type RemoveKeyValueResultCode int32

const (
	RemoveKeyValueResultCodeSuccess  RemoveKeyValueResultCode = 0
	RemoveKeyValueResultCodeNotFound RemoveKeyValueResultCode = -1
)

var RemoveKeyValueResultCodeAll = []RemoveKeyValueResultCode{
	RemoveKeyValueResultCodeSuccess,
	RemoveKeyValueResultCodeNotFound,
}

var removeKeyValueResultCodeMap = map[int32]string{
	0:  "RemoveKeyValueResultCodeSuccess",
	-1: "RemoveKeyValueResultCodeNotFound",
}

var removeKeyValueResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
}

var removeKeyValueResultCodeRevMap = map[string]int32{
	"RemoveKeyValueResultCodeSuccess":  0,
	"RemoveKeyValueResultCodeNotFound": -1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RemoveKeyValueResultCode
func (e RemoveKeyValueResultCode) ValidEnum(v int32) bool {
	_, ok := removeKeyValueResultCodeMap[v]
	return ok
}
func (e RemoveKeyValueResultCode) isFlag() bool {
	for i := len(RemoveKeyValueResultCodeAll) - 1; i >= 0; i-- {
		expected := RemoveKeyValueResultCode(2) << uint64(len(RemoveKeyValueResultCodeAll)-1) >> uint64(len(RemoveKeyValueResultCodeAll)-i)
		if expected != RemoveKeyValueResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RemoveKeyValueResultCode) String() string {
	name, _ := removeKeyValueResultCodeMap[int32(e)]
	return name
}

func (e RemoveKeyValueResultCode) ShortString() string {
	name, _ := removeKeyValueResultCodeShortMap[int32(e)]
	return name
}

func (e RemoveKeyValueResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range RemoveKeyValueResultCodeAll {
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

func (e *RemoveKeyValueResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RemoveKeyValueResultCode(t.Value)
	return nil
}

// RemoveKeyValueResult is an XDR Union defines as:
//
//   //: `RemoveKeyValueResult` represents the result of RemoveKeyValueOp
//        union RemoveKeyValueResult switch (RemoveKeyValueResultCode code)
//        {
//            case SUCCESS:
//                RemoveKeyValueSuccess success;
//            default:
//                void;
//        };
//
type RemoveKeyValueResult struct {
	Code    RemoveKeyValueResultCode `json:"code,omitempty"`
	Success *RemoveKeyValueSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveKeyValueResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveKeyValueResult
func (u RemoveKeyValueResult) ArmForSwitch(sw int32) (string, bool) {
	switch RemoveKeyValueResultCode(sw) {
	case RemoveKeyValueResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewRemoveKeyValueResult creates a new  RemoveKeyValueResult.
func NewRemoveKeyValueResult(code RemoveKeyValueResultCode, value interface{}) (result RemoveKeyValueResult, err error) {
	result.Code = code
	switch RemoveKeyValueResultCode(code) {
	case RemoveKeyValueResultCodeSuccess:
		tv, ok := value.(RemoveKeyValueSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveKeyValueSuccess")
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
func (u RemoveKeyValueResult) MustSuccess() RemoveKeyValueSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RemoveKeyValueResult) GetSuccess() (result RemoveKeyValueSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// RemoveReviewableRequestOp is an XDR Struct defines as:
//
//   struct RemoveReviewableRequestOp
//    {
//        uint64 requestID;
//
//        EmptyExt ext;
//    };
//
type RemoveReviewableRequestOp struct {
	RequestId Uint64   `json:"requestID,omitempty"`
	Ext       EmptyExt `json:"ext,omitempty"`
}

// RemoveReviewableRequestResultCode is an XDR Enum defines as:
//
//   enum RemoveReviewableRequestResultCode
//    {
//        SUCCESS = 0,
//
//
//        NOT_FOUND = -1
//    };
//
type RemoveReviewableRequestResultCode int32

const (
	RemoveReviewableRequestResultCodeSuccess  RemoveReviewableRequestResultCode = 0
	RemoveReviewableRequestResultCodeNotFound RemoveReviewableRequestResultCode = -1
)

var RemoveReviewableRequestResultCodeAll = []RemoveReviewableRequestResultCode{
	RemoveReviewableRequestResultCodeSuccess,
	RemoveReviewableRequestResultCodeNotFound,
}

var removeReviewableRequestResultCodeMap = map[int32]string{
	0:  "RemoveReviewableRequestResultCodeSuccess",
	-1: "RemoveReviewableRequestResultCodeNotFound",
}

var removeReviewableRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
}

var removeReviewableRequestResultCodeRevMap = map[string]int32{
	"RemoveReviewableRequestResultCodeSuccess":  0,
	"RemoveReviewableRequestResultCodeNotFound": -1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RemoveReviewableRequestResultCode
func (e RemoveReviewableRequestResultCode) ValidEnum(v int32) bool {
	_, ok := removeReviewableRequestResultCodeMap[v]
	return ok
}
func (e RemoveReviewableRequestResultCode) isFlag() bool {
	for i := len(RemoveReviewableRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := RemoveReviewableRequestResultCode(2) << uint64(len(RemoveReviewableRequestResultCodeAll)-1) >> uint64(len(RemoveReviewableRequestResultCodeAll)-i)
		if expected != RemoveReviewableRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RemoveReviewableRequestResultCode) String() string {
	name, _ := removeReviewableRequestResultCodeMap[int32(e)]
	return name
}

func (e RemoveReviewableRequestResultCode) ShortString() string {
	name, _ := removeReviewableRequestResultCodeShortMap[int32(e)]
	return name
}

func (e RemoveReviewableRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range RemoveReviewableRequestResultCodeAll {
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

func (e *RemoveReviewableRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RemoveReviewableRequestResultCode(t.Value)
	return nil
}

// RemoveReviewableRequestResult is an XDR Union defines as:
//
//   union RemoveReviewableRequestResult switch (RemoveReviewableRequestResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    default:
//        void;
//    };
//
type RemoveReviewableRequestResult struct {
	Code RemoveReviewableRequestResultCode `json:"code,omitempty"`
	Ext  *EmptyExt                         `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveReviewableRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveReviewableRequestResult
func (u RemoveReviewableRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch RemoveReviewableRequestResultCode(sw) {
	case RemoveReviewableRequestResultCodeSuccess:
		return "Ext", true
	default:
		return "", true
	}
}

// NewRemoveReviewableRequestResult creates a new  RemoveReviewableRequestResult.
func NewRemoveReviewableRequestResult(code RemoveReviewableRequestResultCode, value interface{}) (result RemoveReviewableRequestResult, err error) {
	result.Code = code
	switch RemoveReviewableRequestResultCode(code) {
	case RemoveReviewableRequestResultCodeSuccess:
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
func (u RemoveReviewableRequestResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RemoveReviewableRequestResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// RemoveRoleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type RemoveRoleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveRoleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveRoleOpExt
func (u RemoveRoleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRemoveRoleOpExt creates a new  RemoveRoleOpExt.
func NewRemoveRoleOpExt(v LedgerVersion, value interface{}) (result RemoveRoleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RemoveRoleOp is an XDR Struct defines as:
//
//   //: RemoveSignerRoleData is used to pass necessary params to remove existing signer role
//    struct RemoveRoleOp
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
type RemoveRoleOp struct {
	RoleId Uint64          `json:"roleID,omitempty"`
	Ext    RemoveRoleOpExt `json:"ext,omitempty"`
}

// RemoveRoleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerRoleResultCode
//    enum RemoveRoleResultCode
//    {
//        //: Means that the specified action in `data` of ManageSignerRoleOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: There is no signer role with such id or the source cannot manage a role
//        NOT_FOUND = -1, // does not exist or owner mismatched
//        //: It is not allowed to remove role if it is attached to at least one singer
//        ROLE_IS_USED = -2
//    };
//
type RemoveRoleResultCode int32

const (
	RemoveRoleResultCodeSuccess    RemoveRoleResultCode = 0
	RemoveRoleResultCodeNotFound   RemoveRoleResultCode = -1
	RemoveRoleResultCodeRoleIsUsed RemoveRoleResultCode = -2
)

var RemoveRoleResultCodeAll = []RemoveRoleResultCode{
	RemoveRoleResultCodeSuccess,
	RemoveRoleResultCodeNotFound,
	RemoveRoleResultCodeRoleIsUsed,
}

var removeRoleResultCodeMap = map[int32]string{
	0:  "RemoveRoleResultCodeSuccess",
	-1: "RemoveRoleResultCodeNotFound",
	-2: "RemoveRoleResultCodeRoleIsUsed",
}

var removeRoleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "role_is_used",
}

var removeRoleResultCodeRevMap = map[string]int32{
	"RemoveRoleResultCodeSuccess":    0,
	"RemoveRoleResultCodeNotFound":   -1,
	"RemoveRoleResultCodeRoleIsUsed": -2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RemoveRoleResultCode
func (e RemoveRoleResultCode) ValidEnum(v int32) bool {
	_, ok := removeRoleResultCodeMap[v]
	return ok
}
func (e RemoveRoleResultCode) isFlag() bool {
	for i := len(RemoveRoleResultCodeAll) - 1; i >= 0; i-- {
		expected := RemoveRoleResultCode(2) << uint64(len(RemoveRoleResultCodeAll)-1) >> uint64(len(RemoveRoleResultCodeAll)-i)
		if expected != RemoveRoleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RemoveRoleResultCode) String() string {
	name, _ := removeRoleResultCodeMap[int32(e)]
	return name
}

func (e RemoveRoleResultCode) ShortString() string {
	name, _ := removeRoleResultCodeShortMap[int32(e)]
	return name
}

func (e RemoveRoleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range RemoveRoleResultCodeAll {
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

func (e *RemoveRoleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RemoveRoleResultCode(t.Value)
	return nil
}

// RemoveRoleResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union RemoveRoleResult switch (RemoveRoleResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    default:
//        void;
//    };
//
type RemoveRoleResult struct {
	Code RemoveRoleResultCode `json:"code,omitempty"`
	Ext  *EmptyExt            `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveRoleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveRoleResult
func (u RemoveRoleResult) ArmForSwitch(sw int32) (string, bool) {
	switch RemoveRoleResultCode(sw) {
	case RemoveRoleResultCodeSuccess:
		return "Ext", true
	default:
		return "", true
	}
}

// NewRemoveRoleResult creates a new  RemoveRoleResult.
func NewRemoveRoleResult(code RemoveRoleResultCode, value interface{}) (result RemoveRoleResult, err error) {
	result.Code = code
	switch RemoveRoleResultCode(code) {
	case RemoveRoleResultCodeSuccess:
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
func (u RemoveRoleResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RemoveRoleResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// RemoveRuleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type RemoveRuleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveRuleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveRuleOpExt
func (u RemoveRuleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRemoveRuleOpExt creates a new  RemoveRuleOpExt.
func NewRemoveRuleOpExt(v LedgerVersion, value interface{}) (result RemoveRuleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RemoveRuleOp is an XDR Struct defines as:
//
//   //: RemoveSignerRuleData is used to pass necessary params to remove existing signer rule
//    struct RemoveRuleOp
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
type RemoveRuleOp struct {
	RuleId Uint64          `json:"ruleID,omitempty"`
	Ext    RemoveRuleOpExt `json:"ext,omitempty"`
}

// RemoveRuleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerRuleOp
//    enum RemoveRuleResultCode
//    {
//        //: Specified action in `data` of ManageSignerRuleOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: There is no signer rule with such id or source cannot manage the rule
//        NOT_FOUND = -1, // does not exists or owner mismatched
//        //: It is not allowed to remove the rule if it is attached to at least one role
//        RULE_IS_USED = -2
//    };
//
type RemoveRuleResultCode int32

const (
	RemoveRuleResultCodeSuccess    RemoveRuleResultCode = 0
	RemoveRuleResultCodeNotFound   RemoveRuleResultCode = -1
	RemoveRuleResultCodeRuleIsUsed RemoveRuleResultCode = -2
)

var RemoveRuleResultCodeAll = []RemoveRuleResultCode{
	RemoveRuleResultCodeSuccess,
	RemoveRuleResultCodeNotFound,
	RemoveRuleResultCodeRuleIsUsed,
}

var removeRuleResultCodeMap = map[int32]string{
	0:  "RemoveRuleResultCodeSuccess",
	-1: "RemoveRuleResultCodeNotFound",
	-2: "RemoveRuleResultCodeRuleIsUsed",
}

var removeRuleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "rule_is_used",
}

var removeRuleResultCodeRevMap = map[string]int32{
	"RemoveRuleResultCodeSuccess":    0,
	"RemoveRuleResultCodeNotFound":   -1,
	"RemoveRuleResultCodeRuleIsUsed": -2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RemoveRuleResultCode
func (e RemoveRuleResultCode) ValidEnum(v int32) bool {
	_, ok := removeRuleResultCodeMap[v]
	return ok
}
func (e RemoveRuleResultCode) isFlag() bool {
	for i := len(RemoveRuleResultCodeAll) - 1; i >= 0; i-- {
		expected := RemoveRuleResultCode(2) << uint64(len(RemoveRuleResultCodeAll)-1) >> uint64(len(RemoveRuleResultCodeAll)-i)
		if expected != RemoveRuleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RemoveRuleResultCode) String() string {
	name, _ := removeRuleResultCodeMap[int32(e)]
	return name
}

func (e RemoveRuleResultCode) ShortString() string {
	name, _ := removeRuleResultCodeShortMap[int32(e)]
	return name
}

func (e RemoveRuleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range RemoveRuleResultCodeAll {
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

func (e *RemoveRuleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RemoveRuleResultCode(t.Value)
	return nil
}

// RemoveRuleResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union RemoveRuleResult switch (RemoveRuleResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    case RULE_IS_USED:
//        //: ids of roles which use a rule that cannot be removed
//        uint64 roleIDs<>;
//    default:
//        void;
//    };
//
type RemoveRuleResult struct {
	Code    RemoveRuleResultCode `json:"code,omitempty"`
	Ext     *EmptyExt            `json:"ext,omitempty"`
	RoleIDs *[]Uint64            `json:"roleIDs,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveRuleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveRuleResult
func (u RemoveRuleResult) ArmForSwitch(sw int32) (string, bool) {
	switch RemoveRuleResultCode(sw) {
	case RemoveRuleResultCodeSuccess:
		return "Ext", true
	case RemoveRuleResultCodeRuleIsUsed:
		return "RoleIDs", true
	default:
		return "", true
	}
}

// NewRemoveRuleResult creates a new  RemoveRuleResult.
func NewRemoveRuleResult(code RemoveRuleResultCode, value interface{}) (result RemoveRuleResult, err error) {
	result.Code = code
	switch RemoveRuleResultCode(code) {
	case RemoveRuleResultCodeSuccess:
		tv, ok := value.(EmptyExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be EmptyExt")
			return
		}
		result.Ext = &tv
	case RemoveRuleResultCodeRuleIsUsed:
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

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u RemoveRuleResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RemoveRuleResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// MustRoleIDs retrieves the RoleIDs value from the union,
// panicing if the value is not set.
func (u RemoveRuleResult) MustRoleIDs() []Uint64 {
	val, ok := u.GetRoleIDs()

	if !ok {
		panic("arm RoleIDs is not set")
	}

	return val
}

// GetRoleIDs retrieves the RoleIDs value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RemoveRuleResult) GetRoleIDs() (result []Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RoleIDs" {
		result = *u.RoleIDs
		ok = true
	}

	return
}

// RemoveSignerOp is an XDR Struct defines as:
//
//   //: RemoveSignerData is used to pass necessary data to remove a signer
//    struct RemoveSignerOp
//    {
//        //: Public key of an existing signer
//        PublicKey publicKey;
//
//        //: reserved for future extension
//        EmptyExt ext;
//    };
//
type RemoveSignerOp struct {
	PublicKey PublicKey `json:"publicKey,omitempty"`
	Ext       EmptyExt  `json:"ext,omitempty"`
}

// RemoveSignerResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerOp
//    enum RemoveSignerResultCode
//    {
//        //: Specified action in `data` of ManageSignerOp was successfully executed
//        SUCCESS = 0,
//
//        //: Source account does not have a signer with the provided public key
//        NOT_FOUND = -1 // there is no signer with such public key
//    };
//
type RemoveSignerResultCode int32

const (
	RemoveSignerResultCodeSuccess  RemoveSignerResultCode = 0
	RemoveSignerResultCodeNotFound RemoveSignerResultCode = -1
)

var RemoveSignerResultCodeAll = []RemoveSignerResultCode{
	RemoveSignerResultCodeSuccess,
	RemoveSignerResultCodeNotFound,
}

var removeSignerResultCodeMap = map[int32]string{
	0:  "RemoveSignerResultCodeSuccess",
	-1: "RemoveSignerResultCodeNotFound",
}

var removeSignerResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
}

var removeSignerResultCodeRevMap = map[string]int32{
	"RemoveSignerResultCodeSuccess":  0,
	"RemoveSignerResultCodeNotFound": -1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RemoveSignerResultCode
func (e RemoveSignerResultCode) ValidEnum(v int32) bool {
	_, ok := removeSignerResultCodeMap[v]
	return ok
}
func (e RemoveSignerResultCode) isFlag() bool {
	for i := len(RemoveSignerResultCodeAll) - 1; i >= 0; i-- {
		expected := RemoveSignerResultCode(2) << uint64(len(RemoveSignerResultCodeAll)-1) >> uint64(len(RemoveSignerResultCodeAll)-i)
		if expected != RemoveSignerResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RemoveSignerResultCode) String() string {
	name, _ := removeSignerResultCodeMap[int32(e)]
	return name
}

func (e RemoveSignerResultCode) ShortString() string {
	name, _ := removeSignerResultCodeShortMap[int32(e)]
	return name
}

func (e RemoveSignerResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range RemoveSignerResultCodeAll {
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

func (e *RemoveSignerResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RemoveSignerResultCode(t.Value)
	return nil
}

// RemoveSignerResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union RemoveSignerResult switch (RemoveSignerResultCode code)
//    {
//    case SUCCESS:
//        //: reserved for future extension
//        EmptyExt ext;
//    default:
//        void;
//    };
//
type RemoveSignerResult struct {
	Code RemoveSignerResultCode `json:"code,omitempty"`
	Ext  *EmptyExt              `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RemoveSignerResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RemoveSignerResult
func (u RemoveSignerResult) ArmForSwitch(sw int32) (string, bool) {
	switch RemoveSignerResultCode(sw) {
	case RemoveSignerResultCodeSuccess:
		return "Ext", true
	default:
		return "", true
	}
}

// NewRemoveSignerResult creates a new  RemoveSignerResult.
func NewRemoveSignerResult(code RemoveSignerResultCode, value interface{}) (result RemoveSignerResult, err error) {
	result.Code = code
	switch RemoveSignerResultCode(code) {
	case RemoveSignerResultCodeSuccess:
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
func (u RemoveSignerResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RemoveSignerResult) GetExt() (result EmptyExt, ok bool) {
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
//
//        //: Review action defines an action performed on the pending ReviewableRequest
//        ReviewRequestOpAction action;
//        //: Contains reject reason
//        longstring reason;
//
//        //: Tasks to add to pending
//        uint64 tasksToAdd;
//        //: Tasks to remove from pending
//        uint64 tasksToRemove;
//        //: Details of the current review
//        longstring externalDetails;
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
	RequestId       Uint64                `json:"requestID,omitempty"`
	RequestHash     Hash                  `json:"requestHash,omitempty"`
	Action          ReviewRequestOpAction `json:"action,omitempty"`
	Reason          Longstring            `json:"reason,omitempty"`
	TasksToAdd      Uint64                `json:"tasksToAdd,omitempty"`
	TasksToRemove   Uint64                `json:"tasksToRemove,omitempty"`
	ExternalDetails Longstring            `json:"externalDetails,omitempty"`
	Ext             ReviewRequestOpExt    `json:"ext,omitempty"`
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
//        REMOVING_NOT_SET_TASKS = -10,// cannot remove tasks which are not set
//        //: CheckValid or Confirm of operation is failed
//        INVALID_OPERATION = -11
//    };
//
type ReviewRequestResultCode int32

const (
	ReviewRequestResultCodeSuccess                   ReviewRequestResultCode = 0
	ReviewRequestResultCodeInvalidReason             ReviewRequestResultCode = -1
	ReviewRequestResultCodeInvalidAction             ReviewRequestResultCode = -2
	ReviewRequestResultCodeHashMismatched            ReviewRequestResultCode = -3
	ReviewRequestResultCodeNotFound                  ReviewRequestResultCode = -4
	ReviewRequestResultCodeTypeMismatched            ReviewRequestResultCode = -5
	ReviewRequestResultCodeRejectNotAllowed          ReviewRequestResultCode = -6
	ReviewRequestResultCodeInvalidExternalDetails    ReviewRequestResultCode = -7
	ReviewRequestResultCodeRequestorIsBlocked        ReviewRequestResultCode = -8
	ReviewRequestResultCodePermanentRejectNotAllowed ReviewRequestResultCode = -9
	ReviewRequestResultCodeRemovingNotSetTasks       ReviewRequestResultCode = -10
	ReviewRequestResultCodeInvalidOperation          ReviewRequestResultCode = -11
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
	ReviewRequestResultCodeInvalidOperation,
}

var reviewRequestResultCodeMap = map[int32]string{
	0:   "ReviewRequestResultCodeSuccess",
	-1:  "ReviewRequestResultCodeInvalidReason",
	-2:  "ReviewRequestResultCodeInvalidAction",
	-3:  "ReviewRequestResultCodeHashMismatched",
	-4:  "ReviewRequestResultCodeNotFound",
	-5:  "ReviewRequestResultCodeTypeMismatched",
	-6:  "ReviewRequestResultCodeRejectNotAllowed",
	-7:  "ReviewRequestResultCodeInvalidExternalDetails",
	-8:  "ReviewRequestResultCodeRequestorIsBlocked",
	-9:  "ReviewRequestResultCodePermanentRejectNotAllowed",
	-10: "ReviewRequestResultCodeRemovingNotSetTasks",
	-11: "ReviewRequestResultCodeInvalidOperation",
}

var reviewRequestResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "invalid_reason",
	-2:  "invalid_action",
	-3:  "hash_mismatched",
	-4:  "not_found",
	-5:  "type_mismatched",
	-6:  "reject_not_allowed",
	-7:  "invalid_external_details",
	-8:  "requestor_is_blocked",
	-9:  "permanent_reject_not_allowed",
	-10: "removing_not_set_tasks",
	-11: "invalid_operation",
}

var reviewRequestResultCodeRevMap = map[string]int32{
	"ReviewRequestResultCodeSuccess":                   0,
	"ReviewRequestResultCodeInvalidReason":             -1,
	"ReviewRequestResultCodeInvalidAction":             -2,
	"ReviewRequestResultCodeHashMismatched":            -3,
	"ReviewRequestResultCodeNotFound":                  -4,
	"ReviewRequestResultCodeTypeMismatched":            -5,
	"ReviewRequestResultCodeRejectNotAllowed":          -6,
	"ReviewRequestResultCodeInvalidExternalDetails":    -7,
	"ReviewRequestResultCodeRequestorIsBlocked":        -8,
	"ReviewRequestResultCodePermanentRejectNotAllowed": -9,
	"ReviewRequestResultCodeRemovingNotSetTasks":       -10,
	"ReviewRequestResultCodeInvalidOperation":          -11,
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
//
//        OperationResult operationResults<>;
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
type ExtendedResult struct {
	Fulfilled        bool              `json:"fulfilled,omitempty"`
	OperationResults []OperationResult `json:"operationResults,omitempty"`
	Ext              ExtendedResultExt `json:"ext,omitempty"`
}

// ReviewRequestResult is an XDR Union defines as:
//
//   //: Result of applying the review request with result code
//    union ReviewRequestResult switch (ReviewRequestResultCode code)
//    {
//    case SUCCESS:
//        ExtendedResult success;
//    case INVALID_OPERATION:
//        OperationResult operationResult;
//    default:
//        void;
//    };
//
type ReviewRequestResult struct {
	Code            ReviewRequestResultCode `json:"code,omitempty"`
	Success         *ExtendedResult         `json:"success,omitempty"`
	OperationResult *OperationResult        `json:"operationResult,omitempty"`
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
	case ReviewRequestResultCodeInvalidOperation:
		return "OperationResult", true
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
	case ReviewRequestResultCodeInvalidOperation:
		tv, ok := value.(OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be OperationResult")
			return
		}
		result.OperationResult = &tv
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

// MustOperationResult retrieves the OperationResult value from the union,
// panicing if the value is not set.
func (u ReviewRequestResult) MustOperationResult() OperationResult {
	val, ok := u.GetOperationResult()

	if !ok {
		panic("arm OperationResult is not set")
	}

	return val
}

// GetOperationResult retrieves the OperationResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestResult) GetOperationResult() (result OperationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "OperationResult" {
		result = *u.OperationResult
		ok = true
	}

	return
}

// UpdateAssetOp is an XDR Struct defines as:
//
//   struct UpdateAssetOp
//    {
//        AssetCode code;
//        longstring *details;
//    	uint64 *maxIssuanceAmount;
//    	uint32 *state;
//
//        EmptyExt ext;
//    };
//
type UpdateAssetOp struct {
	Code              AssetCode   `json:"code,omitempty"`
	Details           *Longstring `json:"details,omitempty"`
	MaxIssuanceAmount *Uint64     `json:"maxIssuanceAmount,omitempty"`
	State             *Uint32     `json:"state,omitempty"`
	Ext               EmptyExt    `json:"ext,omitempty"`
}

// UpdateAssetResultCode is an XDR Enum defines as:
//
//   enum UpdateAssetResultCode
//    {
//        SUCCESS = 0,
//
//        NOT_FOUND = -1,
//        INVALID_DETAILS = -2,
//        UNSUFFICIENT_MAX_ISSUANCE_AMOUNT = -3,
//        NOT_DEFINED_UPDATE = -4
//    };
//
type UpdateAssetResultCode int32

const (
	UpdateAssetResultCodeSuccess                       UpdateAssetResultCode = 0
	UpdateAssetResultCodeNotFound                      UpdateAssetResultCode = -1
	UpdateAssetResultCodeInvalidDetails                UpdateAssetResultCode = -2
	UpdateAssetResultCodeUnsufficientMaxIssuanceAmount UpdateAssetResultCode = -3
	UpdateAssetResultCodeNotDefinedUpdate              UpdateAssetResultCode = -4
)

var UpdateAssetResultCodeAll = []UpdateAssetResultCode{
	UpdateAssetResultCodeSuccess,
	UpdateAssetResultCodeNotFound,
	UpdateAssetResultCodeInvalidDetails,
	UpdateAssetResultCodeUnsufficientMaxIssuanceAmount,
	UpdateAssetResultCodeNotDefinedUpdate,
}

var updateAssetResultCodeMap = map[int32]string{
	0:  "UpdateAssetResultCodeSuccess",
	-1: "UpdateAssetResultCodeNotFound",
	-2: "UpdateAssetResultCodeInvalidDetails",
	-3: "UpdateAssetResultCodeUnsufficientMaxIssuanceAmount",
	-4: "UpdateAssetResultCodeNotDefinedUpdate",
}

var updateAssetResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "invalid_details",
	-3: "unsufficient_max_issuance_amount",
	-4: "not_defined_update",
}

var updateAssetResultCodeRevMap = map[string]int32{
	"UpdateAssetResultCodeSuccess":                       0,
	"UpdateAssetResultCodeNotFound":                      -1,
	"UpdateAssetResultCodeInvalidDetails":                -2,
	"UpdateAssetResultCodeUnsufficientMaxIssuanceAmount": -3,
	"UpdateAssetResultCodeNotDefinedUpdate":              -4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for UpdateAssetResultCode
func (e UpdateAssetResultCode) ValidEnum(v int32) bool {
	_, ok := updateAssetResultCodeMap[v]
	return ok
}
func (e UpdateAssetResultCode) isFlag() bool {
	for i := len(UpdateAssetResultCodeAll) - 1; i >= 0; i-- {
		expected := UpdateAssetResultCode(2) << uint64(len(UpdateAssetResultCodeAll)-1) >> uint64(len(UpdateAssetResultCodeAll)-i)
		if expected != UpdateAssetResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e UpdateAssetResultCode) String() string {
	name, _ := updateAssetResultCodeMap[int32(e)]
	return name
}

func (e UpdateAssetResultCode) ShortString() string {
	name, _ := updateAssetResultCodeShortMap[int32(e)]
	return name
}

func (e UpdateAssetResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range UpdateAssetResultCodeAll {
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

func (e *UpdateAssetResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = UpdateAssetResultCode(t.Value)
	return nil
}

// UpdateAssetResult is an XDR Union defines as:
//
//   union UpdateAssetResult switch(UpdateAssetResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    default:
//        void;
//    };
//
type UpdateAssetResult struct {
	Code UpdateAssetResultCode `json:"code,omitempty"`
	Ext  *EmptyExt             `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateAssetResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateAssetResult
func (u UpdateAssetResult) ArmForSwitch(sw int32) (string, bool) {
	switch UpdateAssetResultCode(sw) {
	case UpdateAssetResultCodeSuccess:
		return "Ext", true
	default:
		return "", true
	}
}

// NewUpdateAssetResult creates a new  UpdateAssetResult.
func NewUpdateAssetResult(code UpdateAssetResultCode, value interface{}) (result UpdateAssetResult, err error) {
	result.Code = code
	switch UpdateAssetResultCode(code) {
	case UpdateAssetResultCodeSuccess:
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
func (u UpdateAssetResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateAssetResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// UpdateDataOp is an XDR Struct defines as:
//
//   struct UpdateDataOp
//    {
//        uint64 dataID;
//
//        longstring value;
//
//        EmptyExt ext;
//    };
//
type UpdateDataOp struct {
	DataId Uint64     `json:"dataID,omitempty"`
	Value  Longstring `json:"value,omitempty"`
	Ext    EmptyExt   `json:"ext,omitempty"`
}

// UpdateDataResultCode is an XDR Enum defines as:
//
//   enum UpdateDataResultCode
//    {
//        SUCCESS = 0,
//
//        INVALID_DATA = -1,
//        NOT_FOUND = -2,
//        NOT_ALLOWED = -3
//    };
//
type UpdateDataResultCode int32

const (
	UpdateDataResultCodeSuccess     UpdateDataResultCode = 0
	UpdateDataResultCodeInvalidData UpdateDataResultCode = -1
	UpdateDataResultCodeNotFound    UpdateDataResultCode = -2
	UpdateDataResultCodeNotAllowed  UpdateDataResultCode = -3
)

var UpdateDataResultCodeAll = []UpdateDataResultCode{
	UpdateDataResultCodeSuccess,
	UpdateDataResultCodeInvalidData,
	UpdateDataResultCodeNotFound,
	UpdateDataResultCodeNotAllowed,
}

var updateDataResultCodeMap = map[int32]string{
	0:  "UpdateDataResultCodeSuccess",
	-1: "UpdateDataResultCodeInvalidData",
	-2: "UpdateDataResultCodeNotFound",
	-3: "UpdateDataResultCodeNotAllowed",
}

var updateDataResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_data",
	-2: "not_found",
	-3: "not_allowed",
}

var updateDataResultCodeRevMap = map[string]int32{
	"UpdateDataResultCodeSuccess":     0,
	"UpdateDataResultCodeInvalidData": -1,
	"UpdateDataResultCodeNotFound":    -2,
	"UpdateDataResultCodeNotAllowed":  -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for UpdateDataResultCode
func (e UpdateDataResultCode) ValidEnum(v int32) bool {
	_, ok := updateDataResultCodeMap[v]
	return ok
}
func (e UpdateDataResultCode) isFlag() bool {
	for i := len(UpdateDataResultCodeAll) - 1; i >= 0; i-- {
		expected := UpdateDataResultCode(2) << uint64(len(UpdateDataResultCodeAll)-1) >> uint64(len(UpdateDataResultCodeAll)-i)
		if expected != UpdateDataResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e UpdateDataResultCode) String() string {
	name, _ := updateDataResultCodeMap[int32(e)]
	return name
}

func (e UpdateDataResultCode) ShortString() string {
	name, _ := updateDataResultCodeShortMap[int32(e)]
	return name
}

func (e UpdateDataResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range UpdateDataResultCodeAll {
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

func (e *UpdateDataResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = UpdateDataResultCode(t.Value)
	return nil
}

// UpdateDataResult is an XDR Union defines as:
//
//   union UpdateDataResult switch (UpdateDataResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    default:
//        void;
//    };
//
type UpdateDataResult struct {
	Code UpdateDataResultCode `json:"code,omitempty"`
	Ext  *EmptyExt            `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateDataResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateDataResult
func (u UpdateDataResult) ArmForSwitch(sw int32) (string, bool) {
	switch UpdateDataResultCode(sw) {
	case UpdateDataResultCodeSuccess:
		return "Ext", true
	default:
		return "", true
	}
}

// NewUpdateDataResult creates a new  UpdateDataResult.
func NewUpdateDataResult(code UpdateDataResultCode, value interface{}) (result UpdateDataResult, err error) {
	result.Code = code
	switch UpdateDataResultCode(code) {
	case UpdateDataResultCodeSuccess:
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
func (u UpdateDataResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateDataResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// UpdateReviewableRequestOp is an XDR Struct defines as:
//
//   struct UpdateReviewableRequestOp
//    {
//        uint64 requestID;
//
//        ReviewableRequestOperation operations<>;
//
//        EmptyExt ext;
//    };
//
type UpdateReviewableRequestOp struct {
	RequestId  Uint64                       `json:"requestID,omitempty"`
	Operations []ReviewableRequestOperation `json:"operations,omitempty"`
	Ext        EmptyExt                     `json:"ext,omitempty"`
}

// UpdateReviewableRequestResultCode is an XDR Enum defines as:
//
//   enum UpdateReviewableRequestResultCode
//    {
//        SUCCESS = 0,
//
//        INVALID_OPERATION = -1,
//        TASKS_NOT_FOUND = -2,
//        TOO_MANY_OPERATIONS = -3,
//        NOT_FOUND = -4
//    };
//
type UpdateReviewableRequestResultCode int32

const (
	UpdateReviewableRequestResultCodeSuccess           UpdateReviewableRequestResultCode = 0
	UpdateReviewableRequestResultCodeInvalidOperation  UpdateReviewableRequestResultCode = -1
	UpdateReviewableRequestResultCodeTasksNotFound     UpdateReviewableRequestResultCode = -2
	UpdateReviewableRequestResultCodeTooManyOperations UpdateReviewableRequestResultCode = -3
	UpdateReviewableRequestResultCodeNotFound          UpdateReviewableRequestResultCode = -4
)

var UpdateReviewableRequestResultCodeAll = []UpdateReviewableRequestResultCode{
	UpdateReviewableRequestResultCodeSuccess,
	UpdateReviewableRequestResultCodeInvalidOperation,
	UpdateReviewableRequestResultCodeTasksNotFound,
	UpdateReviewableRequestResultCodeTooManyOperations,
	UpdateReviewableRequestResultCodeNotFound,
}

var updateReviewableRequestResultCodeMap = map[int32]string{
	0:  "UpdateReviewableRequestResultCodeSuccess",
	-1: "UpdateReviewableRequestResultCodeInvalidOperation",
	-2: "UpdateReviewableRequestResultCodeTasksNotFound",
	-3: "UpdateReviewableRequestResultCodeTooManyOperations",
	-4: "UpdateReviewableRequestResultCodeNotFound",
}

var updateReviewableRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_operation",
	-2: "tasks_not_found",
	-3: "too_many_operations",
	-4: "not_found",
}

var updateReviewableRequestResultCodeRevMap = map[string]int32{
	"UpdateReviewableRequestResultCodeSuccess":           0,
	"UpdateReviewableRequestResultCodeInvalidOperation":  -1,
	"UpdateReviewableRequestResultCodeTasksNotFound":     -2,
	"UpdateReviewableRequestResultCodeTooManyOperations": -3,
	"UpdateReviewableRequestResultCodeNotFound":          -4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for UpdateReviewableRequestResultCode
func (e UpdateReviewableRequestResultCode) ValidEnum(v int32) bool {
	_, ok := updateReviewableRequestResultCodeMap[v]
	return ok
}
func (e UpdateReviewableRequestResultCode) isFlag() bool {
	for i := len(UpdateReviewableRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := UpdateReviewableRequestResultCode(2) << uint64(len(UpdateReviewableRequestResultCodeAll)-1) >> uint64(len(UpdateReviewableRequestResultCodeAll)-i)
		if expected != UpdateReviewableRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e UpdateReviewableRequestResultCode) String() string {
	name, _ := updateReviewableRequestResultCodeMap[int32(e)]
	return name
}

func (e UpdateReviewableRequestResultCode) ShortString() string {
	name, _ := updateReviewableRequestResultCodeShortMap[int32(e)]
	return name
}

func (e UpdateReviewableRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range UpdateReviewableRequestResultCodeAll {
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

func (e *UpdateReviewableRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = UpdateReviewableRequestResultCode(t.Value)
	return nil
}

// UpdateReviewableRequestResult is an XDR Union defines as:
//
//   union UpdateReviewableRequestResult switch (UpdateReviewableRequestResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    case INVALID_OPERATION:
//        OperationResult operationResult;
//    case TOO_MANY_OPERATIONS:
//        uint32 maxOperationsCount;
//    default:
//        void;
//    };
//
type UpdateReviewableRequestResult struct {
	Code               UpdateReviewableRequestResultCode `json:"code,omitempty"`
	Ext                *EmptyExt                         `json:"ext,omitempty"`
	OperationResult    *OperationResult                  `json:"operationResult,omitempty"`
	MaxOperationsCount *Uint32                           `json:"maxOperationsCount,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateReviewableRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateReviewableRequestResult
func (u UpdateReviewableRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch UpdateReviewableRequestResultCode(sw) {
	case UpdateReviewableRequestResultCodeSuccess:
		return "Ext", true
	case UpdateReviewableRequestResultCodeInvalidOperation:
		return "OperationResult", true
	case UpdateReviewableRequestResultCodeTooManyOperations:
		return "MaxOperationsCount", true
	default:
		return "", true
	}
}

// NewUpdateReviewableRequestResult creates a new  UpdateReviewableRequestResult.
func NewUpdateReviewableRequestResult(code UpdateReviewableRequestResultCode, value interface{}) (result UpdateReviewableRequestResult, err error) {
	result.Code = code
	switch UpdateReviewableRequestResultCode(code) {
	case UpdateReviewableRequestResultCodeSuccess:
		tv, ok := value.(EmptyExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be EmptyExt")
			return
		}
		result.Ext = &tv
	case UpdateReviewableRequestResultCodeInvalidOperation:
		tv, ok := value.(OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be OperationResult")
			return
		}
		result.OperationResult = &tv
	case UpdateReviewableRequestResultCodeTooManyOperations:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.MaxOperationsCount = &tv
	default:
		// void
	}
	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u UpdateReviewableRequestResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateReviewableRequestResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// MustOperationResult retrieves the OperationResult value from the union,
// panicing if the value is not set.
func (u UpdateReviewableRequestResult) MustOperationResult() OperationResult {
	val, ok := u.GetOperationResult()

	if !ok {
		panic("arm OperationResult is not set")
	}

	return val
}

// GetOperationResult retrieves the OperationResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateReviewableRequestResult) GetOperationResult() (result OperationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "OperationResult" {
		result = *u.OperationResult
		ok = true
	}

	return
}

// MustMaxOperationsCount retrieves the MaxOperationsCount value from the union,
// panicing if the value is not set.
func (u UpdateReviewableRequestResult) MustMaxOperationsCount() Uint32 {
	val, ok := u.GetMaxOperationsCount()

	if !ok {
		panic("arm MaxOperationsCount is not set")
	}

	return val
}

// GetMaxOperationsCount retrieves the MaxOperationsCount value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateReviewableRequestResult) GetMaxOperationsCount() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "MaxOperationsCount" {
		result = *u.MaxOperationsCount
		ok = true
	}

	return
}

// UpdateRoleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateRoleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateRoleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateRoleOpExt
func (u UpdateRoleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateRoleOpExt creates a new  UpdateRoleOpExt.
func NewUpdateRoleOpExt(v LedgerVersion, value interface{}) (result UpdateRoleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateRoleOp is an XDR Struct defines as:
//
//   //: UpdateSignerRoleData is used to pass necessary params to update an existing signer role
//    struct UpdateRoleOp
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
type UpdateRoleOp struct {
	RoleId  Uint64          `json:"roleID,omitempty"`
	RuleIDs []Uint64        `json:"ruleIDs,omitempty"`
	Details Longstring      `json:"details,omitempty"`
	Ext     UpdateRoleOpExt `json:"ext,omitempty"`
}

// UpdateRoleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerRoleResultCode
//    enum UpdateRoleResultCode
//    {
//        //: Means that the specified action in `data` of ManageSignerRoleOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: There is no signer role with such id or the source cannot manage a role
//        NOT_FOUND = -1, // does not exist or owner mismatched
//        //: Passed details have invalid json structure
//        INVALID_DETAILS = -2,
//        //: There is no rule with id passed through `ruleIDs`
//        NO_SUCH_RULE = -3,
//        //: It is not allowed to duplicate ids in `ruleIDs` array
//        RULE_ID_DUPLICATION = -4,
//        //: It is not allowed to pass ruleIDs that are more than maxSignerRuleCount (by default, 128)
//        TOO_MANY_RULE_IDS = -5
//    };
//
type UpdateRoleResultCode int32

const (
	UpdateRoleResultCodeSuccess           UpdateRoleResultCode = 0
	UpdateRoleResultCodeNotFound          UpdateRoleResultCode = -1
	UpdateRoleResultCodeInvalidDetails    UpdateRoleResultCode = -2
	UpdateRoleResultCodeNoSuchRule        UpdateRoleResultCode = -3
	UpdateRoleResultCodeRuleIdDuplication UpdateRoleResultCode = -4
	UpdateRoleResultCodeTooManyRuleIds    UpdateRoleResultCode = -5
)

var UpdateRoleResultCodeAll = []UpdateRoleResultCode{
	UpdateRoleResultCodeSuccess,
	UpdateRoleResultCodeNotFound,
	UpdateRoleResultCodeInvalidDetails,
	UpdateRoleResultCodeNoSuchRule,
	UpdateRoleResultCodeRuleIdDuplication,
	UpdateRoleResultCodeTooManyRuleIds,
}

var updateRoleResultCodeMap = map[int32]string{
	0:  "UpdateRoleResultCodeSuccess",
	-1: "UpdateRoleResultCodeNotFound",
	-2: "UpdateRoleResultCodeInvalidDetails",
	-3: "UpdateRoleResultCodeNoSuchRule",
	-4: "UpdateRoleResultCodeRuleIdDuplication",
	-5: "UpdateRoleResultCodeTooManyRuleIds",
}

var updateRoleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "invalid_details",
	-3: "no_such_rule",
	-4: "rule_id_duplication",
	-5: "too_many_rule_ids",
}

var updateRoleResultCodeRevMap = map[string]int32{
	"UpdateRoleResultCodeSuccess":           0,
	"UpdateRoleResultCodeNotFound":          -1,
	"UpdateRoleResultCodeInvalidDetails":    -2,
	"UpdateRoleResultCodeNoSuchRule":        -3,
	"UpdateRoleResultCodeRuleIdDuplication": -4,
	"UpdateRoleResultCodeTooManyRuleIds":    -5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for UpdateRoleResultCode
func (e UpdateRoleResultCode) ValidEnum(v int32) bool {
	_, ok := updateRoleResultCodeMap[v]
	return ok
}
func (e UpdateRoleResultCode) isFlag() bool {
	for i := len(UpdateRoleResultCodeAll) - 1; i >= 0; i-- {
		expected := UpdateRoleResultCode(2) << uint64(len(UpdateRoleResultCodeAll)-1) >> uint64(len(UpdateRoleResultCodeAll)-i)
		if expected != UpdateRoleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e UpdateRoleResultCode) String() string {
	name, _ := updateRoleResultCodeMap[int32(e)]
	return name
}

func (e UpdateRoleResultCode) ShortString() string {
	name, _ := updateRoleResultCodeShortMap[int32(e)]
	return name
}

func (e UpdateRoleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range UpdateRoleResultCodeAll {
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

func (e *UpdateRoleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = UpdateRoleResultCode(t.Value)
	return nil
}

// UpdateRoleResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union UpdateRoleResult switch (UpdateRoleResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    case RULE_ID_DUPLICATION:
//    case NO_SUCH_RULE:
//        //: ID of a rule that was either duplicated or is default or does not exist
//        uint64 ruleID;
//    case TOO_MANY_RULE_IDS:
//        //: max count of rule ids that can be passed in `ruleIDs` array
//        uint32 maxRuleIDsCount;
//    default:
//        void;
//    };
//
type UpdateRoleResult struct {
	Code            UpdateRoleResultCode `json:"code,omitempty"`
	Ext             *EmptyExt            `json:"ext,omitempty"`
	RuleId          *Uint64              `json:"ruleID,omitempty"`
	MaxRuleIDsCount *Uint32              `json:"maxRuleIDsCount,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateRoleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateRoleResult
func (u UpdateRoleResult) ArmForSwitch(sw int32) (string, bool) {
	switch UpdateRoleResultCode(sw) {
	case UpdateRoleResultCodeSuccess:
		return "Ext", true
	case UpdateRoleResultCodeRuleIdDuplication:
		return "RuleId", true
	case UpdateRoleResultCodeNoSuchRule:
		return "RuleId", true
	case UpdateRoleResultCodeTooManyRuleIds:
		return "MaxRuleIDsCount", true
	default:
		return "", true
	}
}

// NewUpdateRoleResult creates a new  UpdateRoleResult.
func NewUpdateRoleResult(code UpdateRoleResultCode, value interface{}) (result UpdateRoleResult, err error) {
	result.Code = code
	switch UpdateRoleResultCode(code) {
	case UpdateRoleResultCodeSuccess:
		tv, ok := value.(EmptyExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be EmptyExt")
			return
		}
		result.Ext = &tv
	case UpdateRoleResultCodeRuleIdDuplication:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RuleId = &tv
	case UpdateRoleResultCodeNoSuchRule:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RuleId = &tv
	case UpdateRoleResultCodeTooManyRuleIds:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.MaxRuleIDsCount = &tv
	default:
		// void
	}
	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u UpdateRoleResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateRoleResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// MustRuleId retrieves the RuleId value from the union,
// panicing if the value is not set.
func (u UpdateRoleResult) MustRuleId() Uint64 {
	val, ok := u.GetRuleId()

	if !ok {
		panic("arm RuleId is not set")
	}

	return val
}

// GetRuleId retrieves the RuleId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateRoleResult) GetRuleId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RuleId" {
		result = *u.RuleId
		ok = true
	}

	return
}

// MustMaxRuleIDsCount retrieves the MaxRuleIDsCount value from the union,
// panicing if the value is not set.
func (u UpdateRoleResult) MustMaxRuleIDsCount() Uint32 {
	val, ok := u.GetMaxRuleIDsCount()

	if !ok {
		panic("arm MaxRuleIDsCount is not set")
	}

	return val
}

// GetMaxRuleIDsCount retrieves the MaxRuleIDsCount value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateRoleResult) GetMaxRuleIDsCount() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "MaxRuleIDsCount" {
		result = *u.MaxRuleIDsCount
		ok = true
	}

	return
}

// UpdateRuleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateRuleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateRuleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateRuleOpExt
func (u UpdateRuleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateRuleOpExt creates a new  UpdateRuleOpExt.
func NewUpdateRuleOpExt(v LedgerVersion, value interface{}) (result UpdateRuleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateRuleOp is an XDR Struct defines as:
//
//   //: UpdateSignerRuleData is used to pass necessary params to update an existing signer rule
//    struct UpdateRuleOp
//    {
//        //: Identifier of an existing signer rule
//        uint64 ruleID;
//        //: Resource is used to specify entity (for some, with properties) that can be managed through operations
//        RuleResource resource;
//        //: Value from enum that can be applied to `resource`
//        RuleAction action;
//        //: True means that such rule will be automatically added to each new or updated signer role
//        bool forbids;
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
type UpdateRuleOp struct {
	RuleId   Uint64          `json:"ruleID,omitempty"`
	Resource RuleResource    `json:"resource,omitempty"`
	Action   RuleAction      `json:"action,omitempty"`
	Forbids  bool            `json:"forbids,omitempty"`
	Details  Longstring      `json:"details,omitempty"`
	Ext      UpdateRuleOpExt `json:"ext,omitempty"`
}

// UpdateRuleResultCode is an XDR Enum defines as:
//
//   //: Result codes of ManageSignerRuleOp
//    enum UpdateRuleResultCode
//    {
//        //: Specified action in `data` of ManageSignerRuleOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: There is no signer rule with such id or source cannot manage the rule
//        NOT_FOUND = -1, // does not exists or owner mismatched
//        //: Passed details have invalid json structure
//        INVALID_DETAILS = -2,
//
//        INVALID_CUSTOM_ACTION = -3,
//        INVALID_CUSTOM_RESOURCE = -4
//    };
//
type UpdateRuleResultCode int32

const (
	UpdateRuleResultCodeSuccess               UpdateRuleResultCode = 0
	UpdateRuleResultCodeNotFound              UpdateRuleResultCode = -1
	UpdateRuleResultCodeInvalidDetails        UpdateRuleResultCode = -2
	UpdateRuleResultCodeInvalidCustomAction   UpdateRuleResultCode = -3
	UpdateRuleResultCodeInvalidCustomResource UpdateRuleResultCode = -4
)

var UpdateRuleResultCodeAll = []UpdateRuleResultCode{
	UpdateRuleResultCodeSuccess,
	UpdateRuleResultCodeNotFound,
	UpdateRuleResultCodeInvalidDetails,
	UpdateRuleResultCodeInvalidCustomAction,
	UpdateRuleResultCodeInvalidCustomResource,
}

var updateRuleResultCodeMap = map[int32]string{
	0:  "UpdateRuleResultCodeSuccess",
	-1: "UpdateRuleResultCodeNotFound",
	-2: "UpdateRuleResultCodeInvalidDetails",
	-3: "UpdateRuleResultCodeInvalidCustomAction",
	-4: "UpdateRuleResultCodeInvalidCustomResource",
}

var updateRuleResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "invalid_details",
	-3: "invalid_custom_action",
	-4: "invalid_custom_resource",
}

var updateRuleResultCodeRevMap = map[string]int32{
	"UpdateRuleResultCodeSuccess":               0,
	"UpdateRuleResultCodeNotFound":              -1,
	"UpdateRuleResultCodeInvalidDetails":        -2,
	"UpdateRuleResultCodeInvalidCustomAction":   -3,
	"UpdateRuleResultCodeInvalidCustomResource": -4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for UpdateRuleResultCode
func (e UpdateRuleResultCode) ValidEnum(v int32) bool {
	_, ok := updateRuleResultCodeMap[v]
	return ok
}
func (e UpdateRuleResultCode) isFlag() bool {
	for i := len(UpdateRuleResultCodeAll) - 1; i >= 0; i-- {
		expected := UpdateRuleResultCode(2) << uint64(len(UpdateRuleResultCodeAll)-1) >> uint64(len(UpdateRuleResultCodeAll)-i)
		if expected != UpdateRuleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e UpdateRuleResultCode) String() string {
	name, _ := updateRuleResultCodeMap[int32(e)]
	return name
}

func (e UpdateRuleResultCode) ShortString() string {
	name, _ := updateRuleResultCodeShortMap[int32(e)]
	return name
}

func (e UpdateRuleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range UpdateRuleResultCodeAll {
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

func (e *UpdateRuleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = UpdateRuleResultCode(t.Value)
	return nil
}

// UpdateRuleResult is an XDR Union defines as:
//
//   //: Result of operation application
//    union UpdateRuleResult switch (UpdateRuleResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    default:
//        void;
//    };
//
type UpdateRuleResult struct {
	Code UpdateRuleResultCode `json:"code,omitempty"`
	Ext  *EmptyExt            `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateRuleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateRuleResult
func (u UpdateRuleResult) ArmForSwitch(sw int32) (string, bool) {
	switch UpdateRuleResultCode(sw) {
	case UpdateRuleResultCodeSuccess:
		return "Ext", true
	default:
		return "", true
	}
}

// NewUpdateRuleResult creates a new  UpdateRuleResult.
func NewUpdateRuleResult(code UpdateRuleResultCode, value interface{}) (result UpdateRuleResult, err error) {
	result.Code = code
	switch UpdateRuleResultCode(code) {
	case UpdateRuleResultCodeSuccess:
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
func (u UpdateRuleResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateRuleResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// UpdateSignerOp is an XDR Struct defines as:
//
//   struct UpdateSignerOp
//    {
//        SignerData data;
//
//        EmptyExt ext;
//    };
//
type UpdateSignerOp struct {
	Data SignerData `json:"data,omitempty"`
	Ext  EmptyExt   `json:"ext,omitempty"`
}

// UpdateSignerResultCode is an XDR Enum defines as:
//
//   enum UpdateSignerResultCode
//    {
//        //: Specified action in `data` of ManageSignerOp was successfully executed
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        //: Passed details have invalid json structure
//        INVALID_DETAILS = -1, // invalid json details
//        //: Source account does not have a signer with the provided public key
//        NOT_FOUND = -2, // there is no signer with such public key
//        //: There is no role with such id
//        NO_SUCH_ROLE = -3,
//        //: It is not allowed to set weight more than 1000
//        INVALID_WEIGHT = -4, // more than 1000
//        NO_ROLE_IDS = -5,
//        ROLE_ID_DUPLICATION = -6,
//        TOO_MANY_ROLES = -7
//    };
//
type UpdateSignerResultCode int32

const (
	UpdateSignerResultCodeSuccess           UpdateSignerResultCode = 0
	UpdateSignerResultCodeInvalidDetails    UpdateSignerResultCode = -1
	UpdateSignerResultCodeNotFound          UpdateSignerResultCode = -2
	UpdateSignerResultCodeNoSuchRole        UpdateSignerResultCode = -3
	UpdateSignerResultCodeInvalidWeight     UpdateSignerResultCode = -4
	UpdateSignerResultCodeNoRoleIds         UpdateSignerResultCode = -5
	UpdateSignerResultCodeRoleIdDuplication UpdateSignerResultCode = -6
	UpdateSignerResultCodeTooManyRoles      UpdateSignerResultCode = -7
)

var UpdateSignerResultCodeAll = []UpdateSignerResultCode{
	UpdateSignerResultCodeSuccess,
	UpdateSignerResultCodeInvalidDetails,
	UpdateSignerResultCodeNotFound,
	UpdateSignerResultCodeNoSuchRole,
	UpdateSignerResultCodeInvalidWeight,
	UpdateSignerResultCodeNoRoleIds,
	UpdateSignerResultCodeRoleIdDuplication,
	UpdateSignerResultCodeTooManyRoles,
}

var updateSignerResultCodeMap = map[int32]string{
	0:  "UpdateSignerResultCodeSuccess",
	-1: "UpdateSignerResultCodeInvalidDetails",
	-2: "UpdateSignerResultCodeNotFound",
	-3: "UpdateSignerResultCodeNoSuchRole",
	-4: "UpdateSignerResultCodeInvalidWeight",
	-5: "UpdateSignerResultCodeNoRoleIds",
	-6: "UpdateSignerResultCodeRoleIdDuplication",
	-7: "UpdateSignerResultCodeTooManyRoles",
}

var updateSignerResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_details",
	-2: "not_found",
	-3: "no_such_role",
	-4: "invalid_weight",
	-5: "no_role_ids",
	-6: "role_id_duplication",
	-7: "too_many_roles",
}

var updateSignerResultCodeRevMap = map[string]int32{
	"UpdateSignerResultCodeSuccess":           0,
	"UpdateSignerResultCodeInvalidDetails":    -1,
	"UpdateSignerResultCodeNotFound":          -2,
	"UpdateSignerResultCodeNoSuchRole":        -3,
	"UpdateSignerResultCodeInvalidWeight":     -4,
	"UpdateSignerResultCodeNoRoleIds":         -5,
	"UpdateSignerResultCodeRoleIdDuplication": -6,
	"UpdateSignerResultCodeTooManyRoles":      -7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for UpdateSignerResultCode
func (e UpdateSignerResultCode) ValidEnum(v int32) bool {
	_, ok := updateSignerResultCodeMap[v]
	return ok
}
func (e UpdateSignerResultCode) isFlag() bool {
	for i := len(UpdateSignerResultCodeAll) - 1; i >= 0; i-- {
		expected := UpdateSignerResultCode(2) << uint64(len(UpdateSignerResultCodeAll)-1) >> uint64(len(UpdateSignerResultCodeAll)-i)
		if expected != UpdateSignerResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e UpdateSignerResultCode) String() string {
	name, _ := updateSignerResultCodeMap[int32(e)]
	return name
}

func (e UpdateSignerResultCode) ShortString() string {
	name, _ := updateSignerResultCodeShortMap[int32(e)]
	return name
}

func (e UpdateSignerResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range UpdateSignerResultCodeAll {
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

func (e *UpdateSignerResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = UpdateSignerResultCode(t.Value)
	return nil
}

// UpdateSignerResult is an XDR Union defines as:
//
//   union UpdateSignerResult switch (UpdateSignerResultCode code)
//    {
//    case SUCCESS:
//        EmptyExt ext;
//    case NO_SUCH_ROLE:
//    case ROLE_ID_DUPLICATION:
//        uint64 roleID;
//    case TOO_MANY_ROLES:
//        uint32 maxRolesCount;
//    default:
//        void;
//    };
//
type UpdateSignerResult struct {
	Code          UpdateSignerResultCode `json:"code,omitempty"`
	Ext           *EmptyExt              `json:"ext,omitempty"`
	RoleId        *Uint64                `json:"roleID,omitempty"`
	MaxRolesCount *Uint32                `json:"maxRolesCount,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateSignerResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateSignerResult
func (u UpdateSignerResult) ArmForSwitch(sw int32) (string, bool) {
	switch UpdateSignerResultCode(sw) {
	case UpdateSignerResultCodeSuccess:
		return "Ext", true
	case UpdateSignerResultCodeNoSuchRole:
		return "RoleId", true
	case UpdateSignerResultCodeRoleIdDuplication:
		return "RoleId", true
	case UpdateSignerResultCodeTooManyRoles:
		return "MaxRolesCount", true
	default:
		return "", true
	}
}

// NewUpdateSignerResult creates a new  UpdateSignerResult.
func NewUpdateSignerResult(code UpdateSignerResultCode, value interface{}) (result UpdateSignerResult, err error) {
	result.Code = code
	switch UpdateSignerResultCode(code) {
	case UpdateSignerResultCodeSuccess:
		tv, ok := value.(EmptyExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be EmptyExt")
			return
		}
		result.Ext = &tv
	case UpdateSignerResultCodeNoSuchRole:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RoleId = &tv
	case UpdateSignerResultCodeRoleIdDuplication:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RoleId = &tv
	case UpdateSignerResultCodeTooManyRoles:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.MaxRolesCount = &tv
	default:
		// void
	}
	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u UpdateSignerResult) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateSignerResult) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// MustRoleId retrieves the RoleId value from the union,
// panicing if the value is not set.
func (u UpdateSignerResult) MustRoleId() Uint64 {
	val, ok := u.GetRoleId()

	if !ok {
		panic("arm RoleId is not set")
	}

	return val
}

// GetRoleId retrieves the RoleId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateSignerResult) GetRoleId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "RoleId" {
		result = *u.RoleId
		ok = true
	}

	return
}

// MustMaxRolesCount retrieves the MaxRolesCount value from the union,
// panicing if the value is not set.
func (u UpdateSignerResult) MustMaxRolesCount() Uint32 {
	val, ok := u.GetMaxRolesCount()

	if !ok {
		panic("arm MaxRolesCount is not set")
	}

	return val
}

// GetMaxRolesCount retrieves the MaxRolesCount value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u UpdateSignerResult) GetMaxRolesCount() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "MaxRolesCount" {
		result = *u.MaxRolesCount
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
//    //case SCP_QUORUMSET:
//    //    SCPQuorumSet qSet;
//    //case SCP_MESSAGE:
//    //    SCPEnvelope envelope;
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

// CustomRuleAction is an XDR Struct defines as:
//
//   struct CustomRuleAction
//    {
//        longstring actionType;
//
//        longstring actionPayload;
//    };
//
type CustomRuleAction struct {
	ActionType    Longstring `json:"actionType,omitempty"`
	ActionPayload Longstring `json:"actionPayload,omitempty"`
}

// CustomRuleResource is an XDR Struct defines as:
//
//   struct CustomRuleResource
//    {
//        longstring resourceType;
//        longstring resourcePayload;
//    };
//
type CustomRuleResource struct {
	ResourceType    Longstring `json:"resourceType,omitempty"`
	ResourcePayload Longstring `json:"resourcePayload,omitempty"`
}

// RuleResourceType is an XDR Enum defines as:
//
//   enum RuleResourceType
//    {
//        LEDGER_ENTRY = 0,
//        CUSTOM = 1
//    };
//
type RuleResourceType int32

const (
	RuleResourceTypeLedgerEntry RuleResourceType = 0
	RuleResourceTypeCustom      RuleResourceType = 1
)

var RuleResourceTypeAll = []RuleResourceType{
	RuleResourceTypeLedgerEntry,
	RuleResourceTypeCustom,
}

var ruleResourceTypeMap = map[int32]string{
	0: "RuleResourceTypeLedgerEntry",
	1: "RuleResourceTypeCustom",
}

var ruleResourceTypeShortMap = map[int32]string{
	0: "ledger_entry",
	1: "custom",
}

var ruleResourceTypeRevMap = map[string]int32{
	"RuleResourceTypeLedgerEntry": 0,
	"RuleResourceTypeCustom":      1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RuleResourceType
func (e RuleResourceType) ValidEnum(v int32) bool {
	_, ok := ruleResourceTypeMap[v]
	return ok
}
func (e RuleResourceType) isFlag() bool {
	for i := len(RuleResourceTypeAll) - 1; i >= 0; i-- {
		expected := RuleResourceType(2) << uint64(len(RuleResourceTypeAll)-1) >> uint64(len(RuleResourceTypeAll)-i)
		if expected != RuleResourceTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RuleResourceType) String() string {
	name, _ := ruleResourceTypeMap[int32(e)]
	return name
}

func (e RuleResourceType) ShortString() string {
	name, _ := ruleResourceTypeShortMap[int32(e)]
	return name
}

func (e RuleResourceType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range RuleResourceTypeAll {
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

func (e *RuleResourceType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RuleResourceType(t.Value)
	return nil
}

// RuleResource is an XDR Union defines as:
//
//   union RuleResource switch(RuleResourceType resourceType)
//    {
//        case LEDGER_ENTRY:
//            InternalRuleResource internalRuleResource;
//        case CUSTOM:
//            CustomRuleResource customRuleResource;
//    };
//
type RuleResource struct {
	ResourceType         RuleResourceType      `json:"resourceType,omitempty"`
	InternalRuleResource *InternalRuleResource `json:"internalRuleResource,omitempty"`
	CustomRuleResource   *CustomRuleResource   `json:"customRuleResource,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RuleResource) SwitchFieldName() string {
	return "ResourceType"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RuleResource
func (u RuleResource) ArmForSwitch(sw int32) (string, bool) {
	switch RuleResourceType(sw) {
	case RuleResourceTypeLedgerEntry:
		return "InternalRuleResource", true
	case RuleResourceTypeCustom:
		return "CustomRuleResource", true
	}
	return "-", false
}

// NewRuleResource creates a new  RuleResource.
func NewRuleResource(resourceType RuleResourceType, value interface{}) (result RuleResource, err error) {
	result.ResourceType = resourceType
	switch RuleResourceType(resourceType) {
	case RuleResourceTypeLedgerEntry:
		tv, ok := value.(InternalRuleResource)
		if !ok {
			err = fmt.Errorf("invalid value, must be InternalRuleResource")
			return
		}
		result.InternalRuleResource = &tv
	case RuleResourceTypeCustom:
		tv, ok := value.(CustomRuleResource)
		if !ok {
			err = fmt.Errorf("invalid value, must be CustomRuleResource")
			return
		}
		result.CustomRuleResource = &tv
	}
	return
}

// MustInternalRuleResource retrieves the InternalRuleResource value from the union,
// panicing if the value is not set.
func (u RuleResource) MustInternalRuleResource() InternalRuleResource {
	val, ok := u.GetInternalRuleResource()

	if !ok {
		panic("arm InternalRuleResource is not set")
	}

	return val
}

// GetInternalRuleResource retrieves the InternalRuleResource value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleResource) GetInternalRuleResource() (result InternalRuleResource, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.ResourceType))

	if armName == "InternalRuleResource" {
		result = *u.InternalRuleResource
		ok = true
	}

	return
}

// MustCustomRuleResource retrieves the CustomRuleResource value from the union,
// panicing if the value is not set.
func (u RuleResource) MustCustomRuleResource() CustomRuleResource {
	val, ok := u.GetCustomRuleResource()

	if !ok {
		panic("arm CustomRuleResource is not set")
	}

	return val
}

// GetCustomRuleResource retrieves the CustomRuleResource value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleResource) GetCustomRuleResource() (result CustomRuleResource, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.ResourceType))

	if armName == "CustomRuleResource" {
		result = *u.CustomRuleResource
		ok = true
	}

	return
}

// InternalRuleResourceReviewableRequest is an XDR NestedStruct defines as:
//
//   struct
//        {
//            ReviewableRequestOperationRule opRules<>;
//
//            uint32 securityType;
//
//            EmptyExt ext;
//        }
//
type InternalRuleResourceReviewableRequest struct {
	OpRules      []ReviewableRequestOperationRule `json:"opRules,omitempty"`
	SecurityType Uint32                           `json:"securityType,omitempty"`
	Ext          EmptyExt                         `json:"ext,omitempty"`
}

// InternalRuleResourceAsset is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AssetCode assetCode;
//            uint32 securityType;
//            uint32 state;
//
//            EmptyExt ext;
//        }
//
type InternalRuleResourceAsset struct {
	AssetCode    AssetCode `json:"assetCode,omitempty"`
	SecurityType Uint32    `json:"securityType,omitempty"`
	State        Uint32    `json:"state,omitempty"`
	Ext          EmptyExt  `json:"ext,omitempty"`
}

// InternalRuleResourceRole is an XDR NestedStruct defines as:
//
//   struct
//        {
//            //: For signer role creating resource will be triggered if `roleID` equals `0`
//            uint64 roleID;
//
//            EmptyExt ext;
//        }
//
type InternalRuleResourceRole struct {
	RoleId Uint64   `json:"roleID,omitempty"`
	Ext    EmptyExt `json:"ext,omitempty"`
}

// InternalRuleResourceSigner is an XDR NestedStruct defines as:
//
//   struct
//        {
//            uint64 roleIDs<>;
//
//            EmptyExt ext;
//        }
//
type InternalRuleResourceSigner struct {
	RoleIDs []Uint64 `json:"roleIDs,omitempty"`
	Ext     EmptyExt `json:"ext,omitempty"`
}

// InternalRuleResourceKeyValue is an XDR NestedStruct defines as:
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
type InternalRuleResourceKeyValue struct {
	KeyPrefix Longstring `json:"keyPrefix,omitempty"`
	Ext       EmptyExt   `json:"ext,omitempty"`
}

// InternalRuleResourceData is an XDR NestedStruct defines as:
//
//   struct
//        {
//            uint32 securityType;
//            EmptyExt ext;
//        }
//
type InternalRuleResourceData struct {
	SecurityType Uint32   `json:"securityType,omitempty"`
	Ext          EmptyExt `json:"ext,omitempty"`
}

// InternalRuleResource is an XDR Union defines as:
//
//   //: Describes properties of some entries that can be used to restrict the usage of entries
//    union InternalRuleResource switch (LedgerEntryType type)
//    {
//    case REVIEWABLE_REQUEST:
//        //: Describes properties that are equal to managed reviewable request entry fields
//        struct
//        {
//            ReviewableRequestOperationRule opRules<>;
//
//            uint32 securityType;
//
//            EmptyExt ext;
//        } reviewableRequest;
//    case ASSET:
//        //: Describes properties that are equal to managed asset entry fields
//        struct
//        {
//            AssetCode assetCode;
//            uint32 securityType;
//            uint32 state;
//
//            EmptyExt ext;
//        } asset;
//    case ANY:
//        void;
//    case ROLE:
//        //: Describes properties that are equal to managed signer role entry fields
//        struct
//        {
//            //: For signer role creating resource will be triggered if `roleID` equals `0`
//            uint64 roleID;
//
//            EmptyExt ext;
//        } role;
//    case SIGNER:
//        //: Describes properties that are equal to managed signer entry fields
//        struct
//        {
//            uint64 roleIDs<>;
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
//    case DATA:
//        struct
//        {
//            uint32 securityType;
//            EmptyExt ext;
//        } data;
//    default:
//        //: reserved for future extension
//        EmptyExt ext;
//    };
//
type InternalRuleResource struct {
	Type              LedgerEntryType                        `json:"type,omitempty"`
	ReviewableRequest *InternalRuleResourceReviewableRequest `json:"reviewableRequest,omitempty"`
	Asset             *InternalRuleResourceAsset             `json:"asset,omitempty"`
	Role              *InternalRuleResourceRole              `json:"role,omitempty"`
	Signer            *InternalRuleResourceSigner            `json:"signer,omitempty"`
	KeyValue          *InternalRuleResourceKeyValue          `json:"keyValue,omitempty"`
	Data              *InternalRuleResourceData              `json:"data,omitempty"`
	Ext               *EmptyExt                              `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InternalRuleResource) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InternalRuleResource
func (u InternalRuleResource) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeAsset:
		return "Asset", true
	case LedgerEntryTypeAny:
		return "", true
	case LedgerEntryTypeRole:
		return "Role", true
	case LedgerEntryTypeSigner:
		return "Signer", true
	case LedgerEntryTypeKeyValue:
		return "KeyValue", true
	case LedgerEntryTypeData:
		return "Data", true
	default:
		return "Ext", true
	}
}

// NewInternalRuleResource creates a new  InternalRuleResource.
func NewInternalRuleResource(aType LedgerEntryType, value interface{}) (result InternalRuleResource, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeReviewableRequest:
		tv, ok := value.(InternalRuleResourceReviewableRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be InternalRuleResourceReviewableRequest")
			return
		}
		result.ReviewableRequest = &tv
	case LedgerEntryTypeAsset:
		tv, ok := value.(InternalRuleResourceAsset)
		if !ok {
			err = fmt.Errorf("invalid value, must be InternalRuleResourceAsset")
			return
		}
		result.Asset = &tv
	case LedgerEntryTypeAny:
		// void
	case LedgerEntryTypeRole:
		tv, ok := value.(InternalRuleResourceRole)
		if !ok {
			err = fmt.Errorf("invalid value, must be InternalRuleResourceRole")
			return
		}
		result.Role = &tv
	case LedgerEntryTypeSigner:
		tv, ok := value.(InternalRuleResourceSigner)
		if !ok {
			err = fmt.Errorf("invalid value, must be InternalRuleResourceSigner")
			return
		}
		result.Signer = &tv
	case LedgerEntryTypeKeyValue:
		tv, ok := value.(InternalRuleResourceKeyValue)
		if !ok {
			err = fmt.Errorf("invalid value, must be InternalRuleResourceKeyValue")
			return
		}
		result.KeyValue = &tv
	case LedgerEntryTypeData:
		tv, ok := value.(InternalRuleResourceData)
		if !ok {
			err = fmt.Errorf("invalid value, must be InternalRuleResourceData")
			return
		}
		result.Data = &tv
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
func (u InternalRuleResource) MustReviewableRequest() InternalRuleResourceReviewableRequest {
	val, ok := u.GetReviewableRequest()

	if !ok {
		panic("arm ReviewableRequest is not set")
	}

	return val
}

// GetReviewableRequest retrieves the ReviewableRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InternalRuleResource) GetReviewableRequest() (result InternalRuleResourceReviewableRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewableRequest" {
		result = *u.ReviewableRequest
		ok = true
	}

	return
}

// MustAsset retrieves the Asset value from the union,
// panicing if the value is not set.
func (u InternalRuleResource) MustAsset() InternalRuleResourceAsset {
	val, ok := u.GetAsset()

	if !ok {
		panic("arm Asset is not set")
	}

	return val
}

// GetAsset retrieves the Asset value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InternalRuleResource) GetAsset() (result InternalRuleResourceAsset, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Asset" {
		result = *u.Asset
		ok = true
	}

	return
}

// MustRole retrieves the Role value from the union,
// panicing if the value is not set.
func (u InternalRuleResource) MustRole() InternalRuleResourceRole {
	val, ok := u.GetRole()

	if !ok {
		panic("arm Role is not set")
	}

	return val
}

// GetRole retrieves the Role value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InternalRuleResource) GetRole() (result InternalRuleResourceRole, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Role" {
		result = *u.Role
		ok = true
	}

	return
}

// MustSigner retrieves the Signer value from the union,
// panicing if the value is not set.
func (u InternalRuleResource) MustSigner() InternalRuleResourceSigner {
	val, ok := u.GetSigner()

	if !ok {
		panic("arm Signer is not set")
	}

	return val
}

// GetSigner retrieves the Signer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InternalRuleResource) GetSigner() (result InternalRuleResourceSigner, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Signer" {
		result = *u.Signer
		ok = true
	}

	return
}

// MustKeyValue retrieves the KeyValue value from the union,
// panicing if the value is not set.
func (u InternalRuleResource) MustKeyValue() InternalRuleResourceKeyValue {
	val, ok := u.GetKeyValue()

	if !ok {
		panic("arm KeyValue is not set")
	}

	return val
}

// GetKeyValue retrieves the KeyValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InternalRuleResource) GetKeyValue() (result InternalRuleResourceKeyValue, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KeyValue" {
		result = *u.KeyValue
		ok = true
	}

	return
}

// MustData retrieves the Data value from the union,
// panicing if the value is not set.
func (u InternalRuleResource) MustData() InternalRuleResourceData {
	val, ok := u.GetData()

	if !ok {
		panic("arm Data is not set")
	}

	return val
}

// GetData retrieves the Data value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InternalRuleResource) GetData() (result InternalRuleResourceData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Data" {
		result = *u.Data
		ok = true
	}

	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u InternalRuleResource) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u InternalRuleResource) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// RuleActionType is an XDR Enum defines as:
//
//   //: Actions that can be applied to a signer rule resource
//    enum RuleActionType
//    {
//        ANY = 1,
//        CREATE = 2,
//        UPDATE = 4,
//        ISSUE = 5,
//        SEND = 6,
//        REMOVE = 7,
//        DESTROY = 8,
//        REVIEW = 9,
//        CHANGE_ROLES = 10,
//        INITIATE_RECOVERY = 11,
//        RECOVER = 12,
//        UPDATE_MAX_ISSUANCE = 13,
//        UPDATE_STATE = 14,
//        RECEIVE = 19,
//        RECEIVE_ISSUANCE = 20,
//        CUSTOM = 21
//    };
//
type RuleActionType int32

const (
	RuleActionTypeAny               RuleActionType = 1
	RuleActionTypeCreate            RuleActionType = 2
	RuleActionTypeUpdate            RuleActionType = 4
	RuleActionTypeIssue             RuleActionType = 5
	RuleActionTypeSend              RuleActionType = 6
	RuleActionTypeRemove            RuleActionType = 7
	RuleActionTypeDestroy           RuleActionType = 8
	RuleActionTypeReview            RuleActionType = 9
	RuleActionTypeChangeRoles       RuleActionType = 10
	RuleActionTypeInitiateRecovery  RuleActionType = 11
	RuleActionTypeRecover           RuleActionType = 12
	RuleActionTypeUpdateMaxIssuance RuleActionType = 13
	RuleActionTypeUpdateState       RuleActionType = 14
	RuleActionTypeReceive           RuleActionType = 19
	RuleActionTypeReceiveIssuance   RuleActionType = 20
	RuleActionTypeCustom            RuleActionType = 21
)

var RuleActionTypeAll = []RuleActionType{
	RuleActionTypeAny,
	RuleActionTypeCreate,
	RuleActionTypeUpdate,
	RuleActionTypeIssue,
	RuleActionTypeSend,
	RuleActionTypeRemove,
	RuleActionTypeDestroy,
	RuleActionTypeReview,
	RuleActionTypeChangeRoles,
	RuleActionTypeInitiateRecovery,
	RuleActionTypeRecover,
	RuleActionTypeUpdateMaxIssuance,
	RuleActionTypeUpdateState,
	RuleActionTypeReceive,
	RuleActionTypeReceiveIssuance,
	RuleActionTypeCustom,
}

var ruleActionTypeMap = map[int32]string{
	1:  "RuleActionTypeAny",
	2:  "RuleActionTypeCreate",
	4:  "RuleActionTypeUpdate",
	5:  "RuleActionTypeIssue",
	6:  "RuleActionTypeSend",
	7:  "RuleActionTypeRemove",
	8:  "RuleActionTypeDestroy",
	9:  "RuleActionTypeReview",
	10: "RuleActionTypeChangeRoles",
	11: "RuleActionTypeInitiateRecovery",
	12: "RuleActionTypeRecover",
	13: "RuleActionTypeUpdateMaxIssuance",
	14: "RuleActionTypeUpdateState",
	19: "RuleActionTypeReceive",
	20: "RuleActionTypeReceiveIssuance",
	21: "RuleActionTypeCustom",
}

var ruleActionTypeShortMap = map[int32]string{
	1:  "any",
	2:  "create",
	4:  "update",
	5:  "issue",
	6:  "send",
	7:  "remove",
	8:  "destroy",
	9:  "review",
	10: "change_roles",
	11: "initiate_recovery",
	12: "recover",
	13: "update_max_issuance",
	14: "update_state",
	19: "receive",
	20: "receive_issuance",
	21: "custom",
}

var ruleActionTypeRevMap = map[string]int32{
	"RuleActionTypeAny":               1,
	"RuleActionTypeCreate":            2,
	"RuleActionTypeUpdate":            4,
	"RuleActionTypeIssue":             5,
	"RuleActionTypeSend":              6,
	"RuleActionTypeRemove":            7,
	"RuleActionTypeDestroy":           8,
	"RuleActionTypeReview":            9,
	"RuleActionTypeChangeRoles":       10,
	"RuleActionTypeInitiateRecovery":  11,
	"RuleActionTypeRecover":           12,
	"RuleActionTypeUpdateMaxIssuance": 13,
	"RuleActionTypeUpdateState":       14,
	"RuleActionTypeReceive":           19,
	"RuleActionTypeReceiveIssuance":   20,
	"RuleActionTypeCustom":            21,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RuleActionType
func (e RuleActionType) ValidEnum(v int32) bool {
	_, ok := ruleActionTypeMap[v]
	return ok
}
func (e RuleActionType) isFlag() bool {
	for i := len(RuleActionTypeAll) - 1; i >= 0; i-- {
		expected := RuleActionType(2) << uint64(len(RuleActionTypeAll)-1) >> uint64(len(RuleActionTypeAll)-i)
		if expected != RuleActionTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RuleActionType) String() string {
	name, _ := ruleActionTypeMap[int32(e)]
	return name
}

func (e RuleActionType) ShortString() string {
	name, _ := ruleActionTypeShortMap[int32(e)]
	return name
}

func (e RuleActionType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range RuleActionTypeAll {
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

func (e *RuleActionType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RuleActionType(t.Value)
	return nil
}

// RuleActionCreate is an XDR NestedStruct defines as:
//
//   struct {
//            bool forOther;
//
//            EmptyExt ext;
//        }
//
type RuleActionCreate struct {
	ForOther bool     `json:"forOther,omitempty"`
	Ext      EmptyExt `json:"ext,omitempty"`
}

// RuleActionUpdate is an XDR NestedStruct defines as:
//
//   struct {
//            bool forOther;
//
//            EmptyExt ext;
//        }
//
type RuleActionUpdate struct {
	ForOther bool     `json:"forOther,omitempty"`
	Ext      EmptyExt `json:"ext,omitempty"`
}

// RuleActionIssue is an XDR NestedStruct defines as:
//
//   struct {
//            uint32 securityType;
//
//            EmptyExt ext;
//        }
//
type RuleActionIssue struct {
	SecurityType Uint32   `json:"securityType,omitempty"`
	Ext          EmptyExt `json:"ext,omitempty"`
}

// RuleActionDestroy is an XDR NestedStruct defines as:
//
//   struct {
//            uint32 securityType;
//            bool forOther;
//
//            EmptyExt ext;
//        }
//
type RuleActionDestroy struct {
	SecurityType Uint32   `json:"securityType,omitempty"`
	ForOther     bool     `json:"forOther,omitempty"`
	Ext          EmptyExt `json:"ext,omitempty"`
}

// RuleActionSend is an XDR NestedStruct defines as:
//
//   struct {
//            uint32 securityType;
//
//            EmptyExt ext;
//        }
//
type RuleActionSend struct {
	SecurityType Uint32   `json:"securityType,omitempty"`
	Ext          EmptyExt `json:"ext,omitempty"`
}

// RuleActionReceive is an XDR NestedStruct defines as:
//
//   struct {
//            uint32 securityType;
//
//            EmptyExt ext;
//        }
//
type RuleActionReceive struct {
	SecurityType Uint32   `json:"securityType,omitempty"`
	Ext          EmptyExt `json:"ext,omitempty"`
}

// RuleActionReceiveIssuance is an XDR NestedStruct defines as:
//
//   struct {
//            uint32 securityType;
//
//            EmptyExt ext;
//        }
//
type RuleActionReceiveIssuance struct {
	SecurityType Uint32   `json:"securityType,omitempty"`
	Ext          EmptyExt `json:"ext,omitempty"`
}

// RuleActionChangeRoles is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 roleIDs<>; // if roleIDsToSet (from operation body) the same, action will triggered
//            bool forOther;
//
//            EmptyExt ext;
//        }
//
type RuleActionChangeRoles struct {
	RoleIDs  []Uint64 `json:"roleIDs,omitempty"`
	ForOther bool     `json:"forOther,omitempty"`
	Ext      EmptyExt `json:"ext,omitempty"`
}

// RuleActionInitiateRecovery is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 roleIDs<>;
//
//            EmptyExt ext;
//        }
//
type RuleActionInitiateRecovery struct {
	RoleIDs []Uint64 `json:"roleIDs,omitempty"`
	Ext     EmptyExt `json:"ext,omitempty"`
}

// RuleActionReview is an XDR NestedStruct defines as:
//
//   struct {
//            //: Bit mask of tasks that is allowed to add to reviewable request pending tasks
//            uint64 tasksToAdd;
//            //: Bit mask of tasks that is allowed to remove from reviewable request pending tasks
//            uint64 tasksToRemove;
//            EmptyExt ext;
//        }
//
type RuleActionReview struct {
	TasksToAdd    Uint64   `json:"tasksToAdd,omitempty"`
	TasksToRemove Uint64   `json:"tasksToRemove,omitempty"`
	Ext           EmptyExt `json:"ext,omitempty"`
}

// RuleAction is an XDR Union defines as:
//
//   union RuleAction switch (RuleActionType type)
//    {
//    case CREATE:
//        struct {
//            bool forOther;
//
//            EmptyExt ext;
//        } create;
//    case UPDATE:
//        struct {
//            bool forOther;
//
//            EmptyExt ext;
//        } update;
//    case ISSUE:
//        struct {
//            uint32 securityType;
//
//            EmptyExt ext;
//        } issue;
//    case DESTROY:
//        struct {
//            uint32 securityType;
//            bool forOther;
//
//            EmptyExt ext;
//        } destroy;
//    case SEND:
//        struct {
//            uint32 securityType;
//
//            EmptyExt ext;
//        } send;
//    case RECEIVE:
//        struct {
//            uint32 securityType;
//
//            EmptyExt ext;
//        } receive;
//    case RECEIVE_ISSUANCE:
//        struct {
//            uint32 securityType;
//
//            EmptyExt ext;
//        } receiveIssuance;
//    case CHANGE_ROLES:
//        struct {
//            uint64 roleIDs<>; // if roleIDsToSet (from operation body) the same, action will triggered
//            bool forOther;
//
//            EmptyExt ext;
//        } changeRoles;
//    case INITIATE_RECOVERY:
//        struct {
//            uint64 roleIDs<>;
//
//            EmptyExt ext;
//        } initiateRecovery;
//    case REVIEW:
//        struct {
//            //: Bit mask of tasks that is allowed to add to reviewable request pending tasks
//            uint64 tasksToAdd;
//            //: Bit mask of tasks that is allowed to remove from reviewable request pending tasks
//            uint64 tasksToRemove;
//            EmptyExt ext;
//        } review;
//    case CUSTOM:
//        CustomRuleAction customRuleAction;
//    default:
//        EmptyExt ext;
//    };
//
type RuleAction struct {
	Type             RuleActionType              `json:"type,omitempty"`
	Create           *RuleActionCreate           `json:"create,omitempty"`
	Update           *RuleActionUpdate           `json:"update,omitempty"`
	Issue            *RuleActionIssue            `json:"issue,omitempty"`
	Destroy          *RuleActionDestroy          `json:"destroy,omitempty"`
	Send             *RuleActionSend             `json:"send,omitempty"`
	Receive          *RuleActionReceive          `json:"receive,omitempty"`
	ReceiveIssuance  *RuleActionReceiveIssuance  `json:"receiveIssuance,omitempty"`
	ChangeRoles      *RuleActionChangeRoles      `json:"changeRoles,omitempty"`
	InitiateRecovery *RuleActionInitiateRecovery `json:"initiateRecovery,omitempty"`
	Review           *RuleActionReview           `json:"review,omitempty"`
	CustomRuleAction *CustomRuleAction           `json:"customRuleAction,omitempty"`
	Ext              *EmptyExt                   `json:"ext,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RuleAction) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RuleAction
func (u RuleAction) ArmForSwitch(sw int32) (string, bool) {
	switch RuleActionType(sw) {
	case RuleActionTypeCreate:
		return "Create", true
	case RuleActionTypeUpdate:
		return "Update", true
	case RuleActionTypeIssue:
		return "Issue", true
	case RuleActionTypeDestroy:
		return "Destroy", true
	case RuleActionTypeSend:
		return "Send", true
	case RuleActionTypeReceive:
		return "Receive", true
	case RuleActionTypeReceiveIssuance:
		return "ReceiveIssuance", true
	case RuleActionTypeChangeRoles:
		return "ChangeRoles", true
	case RuleActionTypeInitiateRecovery:
		return "InitiateRecovery", true
	case RuleActionTypeReview:
		return "Review", true
	case RuleActionTypeCustom:
		return "CustomRuleAction", true
	default:
		return "Ext", true
	}
}

// NewRuleAction creates a new  RuleAction.
func NewRuleAction(aType RuleActionType, value interface{}) (result RuleAction, err error) {
	result.Type = aType
	switch RuleActionType(aType) {
	case RuleActionTypeCreate:
		tv, ok := value.(RuleActionCreate)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionCreate")
			return
		}
		result.Create = &tv
	case RuleActionTypeUpdate:
		tv, ok := value.(RuleActionUpdate)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionUpdate")
			return
		}
		result.Update = &tv
	case RuleActionTypeIssue:
		tv, ok := value.(RuleActionIssue)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionIssue")
			return
		}
		result.Issue = &tv
	case RuleActionTypeDestroy:
		tv, ok := value.(RuleActionDestroy)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionDestroy")
			return
		}
		result.Destroy = &tv
	case RuleActionTypeSend:
		tv, ok := value.(RuleActionSend)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionSend")
			return
		}
		result.Send = &tv
	case RuleActionTypeReceive:
		tv, ok := value.(RuleActionReceive)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionReceive")
			return
		}
		result.Receive = &tv
	case RuleActionTypeReceiveIssuance:
		tv, ok := value.(RuleActionReceiveIssuance)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionReceiveIssuance")
			return
		}
		result.ReceiveIssuance = &tv
	case RuleActionTypeChangeRoles:
		tv, ok := value.(RuleActionChangeRoles)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionChangeRoles")
			return
		}
		result.ChangeRoles = &tv
	case RuleActionTypeInitiateRecovery:
		tv, ok := value.(RuleActionInitiateRecovery)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionInitiateRecovery")
			return
		}
		result.InitiateRecovery = &tv
	case RuleActionTypeReview:
		tv, ok := value.(RuleActionReview)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleActionReview")
			return
		}
		result.Review = &tv
	case RuleActionTypeCustom:
		tv, ok := value.(CustomRuleAction)
		if !ok {
			err = fmt.Errorf("invalid value, must be CustomRuleAction")
			return
		}
		result.CustomRuleAction = &tv
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

// MustCreate retrieves the Create value from the union,
// panicing if the value is not set.
func (u RuleAction) MustCreate() RuleActionCreate {
	val, ok := u.GetCreate()

	if !ok {
		panic("arm Create is not set")
	}

	return val
}

// GetCreate retrieves the Create value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetCreate() (result RuleActionCreate, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Create" {
		result = *u.Create
		ok = true
	}

	return
}

// MustUpdate retrieves the Update value from the union,
// panicing if the value is not set.
func (u RuleAction) MustUpdate() RuleActionUpdate {
	val, ok := u.GetUpdate()

	if !ok {
		panic("arm Update is not set")
	}

	return val
}

// GetUpdate retrieves the Update value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetUpdate() (result RuleActionUpdate, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Update" {
		result = *u.Update
		ok = true
	}

	return
}

// MustIssue retrieves the Issue value from the union,
// panicing if the value is not set.
func (u RuleAction) MustIssue() RuleActionIssue {
	val, ok := u.GetIssue()

	if !ok {
		panic("arm Issue is not set")
	}

	return val
}

// GetIssue retrieves the Issue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetIssue() (result RuleActionIssue, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Issue" {
		result = *u.Issue
		ok = true
	}

	return
}

// MustDestroy retrieves the Destroy value from the union,
// panicing if the value is not set.
func (u RuleAction) MustDestroy() RuleActionDestroy {
	val, ok := u.GetDestroy()

	if !ok {
		panic("arm Destroy is not set")
	}

	return val
}

// GetDestroy retrieves the Destroy value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetDestroy() (result RuleActionDestroy, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Destroy" {
		result = *u.Destroy
		ok = true
	}

	return
}

// MustSend retrieves the Send value from the union,
// panicing if the value is not set.
func (u RuleAction) MustSend() RuleActionSend {
	val, ok := u.GetSend()

	if !ok {
		panic("arm Send is not set")
	}

	return val
}

// GetSend retrieves the Send value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetSend() (result RuleActionSend, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Send" {
		result = *u.Send
		ok = true
	}

	return
}

// MustReceive retrieves the Receive value from the union,
// panicing if the value is not set.
func (u RuleAction) MustReceive() RuleActionReceive {
	val, ok := u.GetReceive()

	if !ok {
		panic("arm Receive is not set")
	}

	return val
}

// GetReceive retrieves the Receive value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetReceive() (result RuleActionReceive, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Receive" {
		result = *u.Receive
		ok = true
	}

	return
}

// MustReceiveIssuance retrieves the ReceiveIssuance value from the union,
// panicing if the value is not set.
func (u RuleAction) MustReceiveIssuance() RuleActionReceiveIssuance {
	val, ok := u.GetReceiveIssuance()

	if !ok {
		panic("arm ReceiveIssuance is not set")
	}

	return val
}

// GetReceiveIssuance retrieves the ReceiveIssuance value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetReceiveIssuance() (result RuleActionReceiveIssuance, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReceiveIssuance" {
		result = *u.ReceiveIssuance
		ok = true
	}

	return
}

// MustChangeRoles retrieves the ChangeRoles value from the union,
// panicing if the value is not set.
func (u RuleAction) MustChangeRoles() RuleActionChangeRoles {
	val, ok := u.GetChangeRoles()

	if !ok {
		panic("arm ChangeRoles is not set")
	}

	return val
}

// GetChangeRoles retrieves the ChangeRoles value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetChangeRoles() (result RuleActionChangeRoles, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChangeRoles" {
		result = *u.ChangeRoles
		ok = true
	}

	return
}

// MustInitiateRecovery retrieves the InitiateRecovery value from the union,
// panicing if the value is not set.
func (u RuleAction) MustInitiateRecovery() RuleActionInitiateRecovery {
	val, ok := u.GetInitiateRecovery()

	if !ok {
		panic("arm InitiateRecovery is not set")
	}

	return val
}

// GetInitiateRecovery retrieves the InitiateRecovery value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetInitiateRecovery() (result RuleActionInitiateRecovery, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "InitiateRecovery" {
		result = *u.InitiateRecovery
		ok = true
	}

	return
}

// MustReview retrieves the Review value from the union,
// panicing if the value is not set.
func (u RuleAction) MustReview() RuleActionReview {
	val, ok := u.GetReview()

	if !ok {
		panic("arm Review is not set")
	}

	return val
}

// GetReview retrieves the Review value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetReview() (result RuleActionReview, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Review" {
		result = *u.Review
		ok = true
	}

	return
}

// MustCustomRuleAction retrieves the CustomRuleAction value from the union,
// panicing if the value is not set.
func (u RuleAction) MustCustomRuleAction() CustomRuleAction {
	val, ok := u.GetCustomRuleAction()

	if !ok {
		panic("arm CustomRuleAction is not set")
	}

	return val
}

// GetCustomRuleAction retrieves the CustomRuleAction value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetCustomRuleAction() (result CustomRuleAction, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CustomRuleAction" {
		result = *u.CustomRuleAction
		ok = true
	}

	return
}

// MustExt retrieves the Ext value from the union,
// panicing if the value is not set.
func (u RuleAction) MustExt() EmptyExt {
	val, ok := u.GetExt()

	if !ok {
		panic("arm Ext is not set")
	}

	return val
}

// GetExt retrieves the Ext value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RuleAction) GetExt() (result EmptyExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ext" {
		result = *u.Ext
		ok = true
	}

	return
}

// ReviewableRequestOperationRule is an XDR Struct defines as:
//
//   struct ReviewableRequestOperationRule
//    {
//        InternalRuleResource resource;
//
//        RuleAction action;
//
//        EmptyExt ext;
//    };
//
type ReviewableRequestOperationRule struct {
	Resource InternalRuleResource `json:"resource,omitempty"`
	Action   RuleAction           `json:"action,omitempty"`
	Ext      EmptyExt             `json:"ext,omitempty"`
}

// OperationBody is an XDR NestedUnion defines as:
//
//   union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountOp createAccountOp;
//    	case DESTRUCTION:
//    		DestructionOp destructionOp;
//    	case CREATE_BALANCE:
//    		CreateBalanceOp createBalanceOp;
//        case CREATE_ASSET:
//            CreateAssetOp createAssetOp;
//        case UPDATE_ASSET:
//            UpdateAssetOp updateAssetOp;
//        case CREATE_DATA:
//            CreateDataOp createDataOp;
//        case UPDATE_DATA:
//            UpdateDataOp updateDataOp;
//        case REMOVE_DATA:
//            RemoveDataOp removeDataOp;
//        case REVIEW_REQUEST:
//    		ReviewRequestOp reviewRequestOp;
//    	case PUT_KEY_VALUE:
//    	    PutKeyValueOp putKeyValueOp;
//        case REMOVE_KEY_VALUE:
//    	    RemoveKeyValueOp removeKeyValueOp;
//    	case CHANGE_ACCOUNT_ROLES:
//    		ChangeAccountRolesOp changeAccountRolesOp;
//        case PAYMENT:
//            PaymentOp paymentOp;
//        case CREATE_SIGNER:
//            CreateSignerOp createSignerOp;
//        case UPDATE_SIGNER:
//            UpdateSignerOp updateSignerOp;
//        case REMOVE_SIGNER:
//            RemoveSignerOp removeSignerOp;
//        case CREATE_ROLE:
//            CreateRoleOp createRoleOp;
//        case UPDATE_ROLE:
//            UpdateRoleOp updateRoleOp;
//        case REMOVE_ROLE:
//            RemoveRoleOp removeRoleOp;
//        case CREATE_RULE:
//            CreateRuleOp createRuleOp;
//        case UPDATE_RULE:
//            UpdateRuleOp updateRuleOp;
//        case REMOVE_RULE:
//            RemoveRuleOp removeRuleOp;
//        case CREATE_REVIEWABLE_REQUEST:
//            CreateReviewableRequestOp createReviewableRequestOp;
//        case UPDATE_REVIEWABLE_REQUEST:
//            UpdateReviewableRequestOp updateReviewableRequestOp;
//        case REMOVE_REVIEWABLE_REQUEST:
//            RemoveReviewableRequestOp removeReviewableRequestOp;
//        case INITIATE_KYC_RECOVERY:
//            InitiateKYCRecoveryOp initiateKYCRecoveryOp;
//        case KYC_RECOVERY:
//            KYCRecoveryOp kycRecoveryOp;
//        case ISSUANCE:
//            IssuanceOp issuanceOp;
//        }
//
type OperationBody struct {
	Type                      OperationType              `json:"type,omitempty"`
	CreateAccountOp           *CreateAccountOp           `json:"createAccountOp,omitempty"`
	DestructionOp             *DestructionOp             `json:"destructionOp,omitempty"`
	CreateBalanceOp           *CreateBalanceOp           `json:"createBalanceOp,omitempty"`
	CreateAssetOp             *CreateAssetOp             `json:"createAssetOp,omitempty"`
	UpdateAssetOp             *UpdateAssetOp             `json:"updateAssetOp,omitempty"`
	CreateDataOp              *CreateDataOp              `json:"createDataOp,omitempty"`
	UpdateDataOp              *UpdateDataOp              `json:"updateDataOp,omitempty"`
	RemoveDataOp              *RemoveDataOp              `json:"removeDataOp,omitempty"`
	ReviewRequestOp           *ReviewRequestOp           `json:"reviewRequestOp,omitempty"`
	PutKeyValueOp             *PutKeyValueOp             `json:"putKeyValueOp,omitempty"`
	RemoveKeyValueOp          *RemoveKeyValueOp          `json:"removeKeyValueOp,omitempty"`
	ChangeAccountRolesOp      *ChangeAccountRolesOp      `json:"changeAccountRolesOp,omitempty"`
	PaymentOp                 *PaymentOp                 `json:"paymentOp,omitempty"`
	CreateSignerOp            *CreateSignerOp            `json:"createSignerOp,omitempty"`
	UpdateSignerOp            *UpdateSignerOp            `json:"updateSignerOp,omitempty"`
	RemoveSignerOp            *RemoveSignerOp            `json:"removeSignerOp,omitempty"`
	CreateRoleOp              *CreateRoleOp              `json:"createRoleOp,omitempty"`
	UpdateRoleOp              *UpdateRoleOp              `json:"updateRoleOp,omitempty"`
	RemoveRoleOp              *RemoveRoleOp              `json:"removeRoleOp,omitempty"`
	CreateRuleOp              *CreateRuleOp              `json:"createRuleOp,omitempty"`
	UpdateRuleOp              *UpdateRuleOp              `json:"updateRuleOp,omitempty"`
	RemoveRuleOp              *RemoveRuleOp              `json:"removeRuleOp,omitempty"`
	CreateReviewableRequestOp *CreateReviewableRequestOp `json:"createReviewableRequestOp,omitempty"`
	UpdateReviewableRequestOp *UpdateReviewableRequestOp `json:"updateReviewableRequestOp,omitempty"`
	RemoveReviewableRequestOp *RemoveReviewableRequestOp `json:"removeReviewableRequestOp,omitempty"`
	InitiateKycRecoveryOp     *InitiateKycRecoveryOp     `json:"initiateKYCRecoveryOp,omitempty"`
	KycRecoveryOp             *KycRecoveryOp             `json:"kycRecoveryOp,omitempty"`
	IssuanceOp                *IssuanceOp                `json:"issuanceOp,omitempty"`
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
	case OperationTypeDestruction:
		return "DestructionOp", true
	case OperationTypeCreateBalance:
		return "CreateBalanceOp", true
	case OperationTypeCreateAsset:
		return "CreateAssetOp", true
	case OperationTypeUpdateAsset:
		return "UpdateAssetOp", true
	case OperationTypeCreateData:
		return "CreateDataOp", true
	case OperationTypeUpdateData:
		return "UpdateDataOp", true
	case OperationTypeRemoveData:
		return "RemoveDataOp", true
	case OperationTypeReviewRequest:
		return "ReviewRequestOp", true
	case OperationTypePutKeyValue:
		return "PutKeyValueOp", true
	case OperationTypeRemoveKeyValue:
		return "RemoveKeyValueOp", true
	case OperationTypeChangeAccountRoles:
		return "ChangeAccountRolesOp", true
	case OperationTypePayment:
		return "PaymentOp", true
	case OperationTypeCreateSigner:
		return "CreateSignerOp", true
	case OperationTypeUpdateSigner:
		return "UpdateSignerOp", true
	case OperationTypeRemoveSigner:
		return "RemoveSignerOp", true
	case OperationTypeCreateRole:
		return "CreateRoleOp", true
	case OperationTypeUpdateRole:
		return "UpdateRoleOp", true
	case OperationTypeRemoveRole:
		return "RemoveRoleOp", true
	case OperationTypeCreateRule:
		return "CreateRuleOp", true
	case OperationTypeUpdateRule:
		return "UpdateRuleOp", true
	case OperationTypeRemoveRule:
		return "RemoveRuleOp", true
	case OperationTypeCreateReviewableRequest:
		return "CreateReviewableRequestOp", true
	case OperationTypeUpdateReviewableRequest:
		return "UpdateReviewableRequestOp", true
	case OperationTypeRemoveReviewableRequest:
		return "RemoveReviewableRequestOp", true
	case OperationTypeInitiateKycRecovery:
		return "InitiateKycRecoveryOp", true
	case OperationTypeKycRecovery:
		return "KycRecoveryOp", true
	case OperationTypeIssuance:
		return "IssuanceOp", true
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
	case OperationTypeDestruction:
		tv, ok := value.(DestructionOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be DestructionOp")
			return
		}
		result.DestructionOp = &tv
	case OperationTypeCreateBalance:
		tv, ok := value.(CreateBalanceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateBalanceOp")
			return
		}
		result.CreateBalanceOp = &tv
	case OperationTypeCreateAsset:
		tv, ok := value.(CreateAssetOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAssetOp")
			return
		}
		result.CreateAssetOp = &tv
	case OperationTypeUpdateAsset:
		tv, ok := value.(UpdateAssetOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateAssetOp")
			return
		}
		result.UpdateAssetOp = &tv
	case OperationTypeCreateData:
		tv, ok := value.(CreateDataOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateDataOp")
			return
		}
		result.CreateDataOp = &tv
	case OperationTypeUpdateData:
		tv, ok := value.(UpdateDataOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateDataOp")
			return
		}
		result.UpdateDataOp = &tv
	case OperationTypeRemoveData:
		tv, ok := value.(RemoveDataOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveDataOp")
			return
		}
		result.RemoveDataOp = &tv
	case OperationTypeReviewRequest:
		tv, ok := value.(ReviewRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestOp")
			return
		}
		result.ReviewRequestOp = &tv
	case OperationTypePutKeyValue:
		tv, ok := value.(PutKeyValueOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PutKeyValueOp")
			return
		}
		result.PutKeyValueOp = &tv
	case OperationTypeRemoveKeyValue:
		tv, ok := value.(RemoveKeyValueOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveKeyValueOp")
			return
		}
		result.RemoveKeyValueOp = &tv
	case OperationTypeChangeAccountRoles:
		tv, ok := value.(ChangeAccountRolesOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ChangeAccountRolesOp")
			return
		}
		result.ChangeAccountRolesOp = &tv
	case OperationTypePayment:
		tv, ok := value.(PaymentOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentOp")
			return
		}
		result.PaymentOp = &tv
	case OperationTypeCreateSigner:
		tv, ok := value.(CreateSignerOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateSignerOp")
			return
		}
		result.CreateSignerOp = &tv
	case OperationTypeUpdateSigner:
		tv, ok := value.(UpdateSignerOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateSignerOp")
			return
		}
		result.UpdateSignerOp = &tv
	case OperationTypeRemoveSigner:
		tv, ok := value.(RemoveSignerOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveSignerOp")
			return
		}
		result.RemoveSignerOp = &tv
	case OperationTypeCreateRole:
		tv, ok := value.(CreateRoleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateRoleOp")
			return
		}
		result.CreateRoleOp = &tv
	case OperationTypeUpdateRole:
		tv, ok := value.(UpdateRoleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateRoleOp")
			return
		}
		result.UpdateRoleOp = &tv
	case OperationTypeRemoveRole:
		tv, ok := value.(RemoveRoleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveRoleOp")
			return
		}
		result.RemoveRoleOp = &tv
	case OperationTypeCreateRule:
		tv, ok := value.(CreateRuleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateRuleOp")
			return
		}
		result.CreateRuleOp = &tv
	case OperationTypeUpdateRule:
		tv, ok := value.(UpdateRuleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateRuleOp")
			return
		}
		result.UpdateRuleOp = &tv
	case OperationTypeRemoveRule:
		tv, ok := value.(RemoveRuleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveRuleOp")
			return
		}
		result.RemoveRuleOp = &tv
	case OperationTypeCreateReviewableRequest:
		tv, ok := value.(CreateReviewableRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateReviewableRequestOp")
			return
		}
		result.CreateReviewableRequestOp = &tv
	case OperationTypeUpdateReviewableRequest:
		tv, ok := value.(UpdateReviewableRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateReviewableRequestOp")
			return
		}
		result.UpdateReviewableRequestOp = &tv
	case OperationTypeRemoveReviewableRequest:
		tv, ok := value.(RemoveReviewableRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveReviewableRequestOp")
			return
		}
		result.RemoveReviewableRequestOp = &tv
	case OperationTypeInitiateKycRecovery:
		tv, ok := value.(InitiateKycRecoveryOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be InitiateKycRecoveryOp")
			return
		}
		result.InitiateKycRecoveryOp = &tv
	case OperationTypeKycRecovery:
		tv, ok := value.(KycRecoveryOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be KycRecoveryOp")
			return
		}
		result.KycRecoveryOp = &tv
	case OperationTypeIssuance:
		tv, ok := value.(IssuanceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be IssuanceOp")
			return
		}
		result.IssuanceOp = &tv
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

// MustDestructionOp retrieves the DestructionOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustDestructionOp() DestructionOp {
	val, ok := u.GetDestructionOp()

	if !ok {
		panic("arm DestructionOp is not set")
	}

	return val
}

// GetDestructionOp retrieves the DestructionOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetDestructionOp() (result DestructionOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DestructionOp" {
		result = *u.DestructionOp
		ok = true
	}

	return
}

// MustCreateBalanceOp retrieves the CreateBalanceOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateBalanceOp() CreateBalanceOp {
	val, ok := u.GetCreateBalanceOp()

	if !ok {
		panic("arm CreateBalanceOp is not set")
	}

	return val
}

// GetCreateBalanceOp retrieves the CreateBalanceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateBalanceOp() (result CreateBalanceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateBalanceOp" {
		result = *u.CreateBalanceOp
		ok = true
	}

	return
}

// MustCreateAssetOp retrieves the CreateAssetOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateAssetOp() CreateAssetOp {
	val, ok := u.GetCreateAssetOp()

	if !ok {
		panic("arm CreateAssetOp is not set")
	}

	return val
}

// GetCreateAssetOp retrieves the CreateAssetOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateAssetOp() (result CreateAssetOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAssetOp" {
		result = *u.CreateAssetOp
		ok = true
	}

	return
}

// MustUpdateAssetOp retrieves the UpdateAssetOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustUpdateAssetOp() UpdateAssetOp {
	val, ok := u.GetUpdateAssetOp()

	if !ok {
		panic("arm UpdateAssetOp is not set")
	}

	return val
}

// GetUpdateAssetOp retrieves the UpdateAssetOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetUpdateAssetOp() (result UpdateAssetOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateAssetOp" {
		result = *u.UpdateAssetOp
		ok = true
	}

	return
}

// MustCreateDataOp retrieves the CreateDataOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateDataOp() CreateDataOp {
	val, ok := u.GetCreateDataOp()

	if !ok {
		panic("arm CreateDataOp is not set")
	}

	return val
}

// GetCreateDataOp retrieves the CreateDataOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateDataOp() (result CreateDataOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateDataOp" {
		result = *u.CreateDataOp
		ok = true
	}

	return
}

// MustUpdateDataOp retrieves the UpdateDataOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustUpdateDataOp() UpdateDataOp {
	val, ok := u.GetUpdateDataOp()

	if !ok {
		panic("arm UpdateDataOp is not set")
	}

	return val
}

// GetUpdateDataOp retrieves the UpdateDataOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetUpdateDataOp() (result UpdateDataOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateDataOp" {
		result = *u.UpdateDataOp
		ok = true
	}

	return
}

// MustRemoveDataOp retrieves the RemoveDataOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRemoveDataOp() RemoveDataOp {
	val, ok := u.GetRemoveDataOp()

	if !ok {
		panic("arm RemoveDataOp is not set")
	}

	return val
}

// GetRemoveDataOp retrieves the RemoveDataOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRemoveDataOp() (result RemoveDataOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveDataOp" {
		result = *u.RemoveDataOp
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

// MustPutKeyValueOp retrieves the PutKeyValueOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustPutKeyValueOp() PutKeyValueOp {
	val, ok := u.GetPutKeyValueOp()

	if !ok {
		panic("arm PutKeyValueOp is not set")
	}

	return val
}

// GetPutKeyValueOp retrieves the PutKeyValueOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetPutKeyValueOp() (result PutKeyValueOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PutKeyValueOp" {
		result = *u.PutKeyValueOp
		ok = true
	}

	return
}

// MustRemoveKeyValueOp retrieves the RemoveKeyValueOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRemoveKeyValueOp() RemoveKeyValueOp {
	val, ok := u.GetRemoveKeyValueOp()

	if !ok {
		panic("arm RemoveKeyValueOp is not set")
	}

	return val
}

// GetRemoveKeyValueOp retrieves the RemoveKeyValueOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRemoveKeyValueOp() (result RemoveKeyValueOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveKeyValueOp" {
		result = *u.RemoveKeyValueOp
		ok = true
	}

	return
}

// MustChangeAccountRolesOp retrieves the ChangeAccountRolesOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustChangeAccountRolesOp() ChangeAccountRolesOp {
	val, ok := u.GetChangeAccountRolesOp()

	if !ok {
		panic("arm ChangeAccountRolesOp is not set")
	}

	return val
}

// GetChangeAccountRolesOp retrieves the ChangeAccountRolesOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetChangeAccountRolesOp() (result ChangeAccountRolesOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChangeAccountRolesOp" {
		result = *u.ChangeAccountRolesOp
		ok = true
	}

	return
}

// MustPaymentOp retrieves the PaymentOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustPaymentOp() PaymentOp {
	val, ok := u.GetPaymentOp()

	if !ok {
		panic("arm PaymentOp is not set")
	}

	return val
}

// GetPaymentOp retrieves the PaymentOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetPaymentOp() (result PaymentOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentOp" {
		result = *u.PaymentOp
		ok = true
	}

	return
}

// MustCreateSignerOp retrieves the CreateSignerOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateSignerOp() CreateSignerOp {
	val, ok := u.GetCreateSignerOp()

	if !ok {
		panic("arm CreateSignerOp is not set")
	}

	return val
}

// GetCreateSignerOp retrieves the CreateSignerOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateSignerOp() (result CreateSignerOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateSignerOp" {
		result = *u.CreateSignerOp
		ok = true
	}

	return
}

// MustUpdateSignerOp retrieves the UpdateSignerOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustUpdateSignerOp() UpdateSignerOp {
	val, ok := u.GetUpdateSignerOp()

	if !ok {
		panic("arm UpdateSignerOp is not set")
	}

	return val
}

// GetUpdateSignerOp retrieves the UpdateSignerOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetUpdateSignerOp() (result UpdateSignerOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateSignerOp" {
		result = *u.UpdateSignerOp
		ok = true
	}

	return
}

// MustRemoveSignerOp retrieves the RemoveSignerOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRemoveSignerOp() RemoveSignerOp {
	val, ok := u.GetRemoveSignerOp()

	if !ok {
		panic("arm RemoveSignerOp is not set")
	}

	return val
}

// GetRemoveSignerOp retrieves the RemoveSignerOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRemoveSignerOp() (result RemoveSignerOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveSignerOp" {
		result = *u.RemoveSignerOp
		ok = true
	}

	return
}

// MustCreateRoleOp retrieves the CreateRoleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateRoleOp() CreateRoleOp {
	val, ok := u.GetCreateRoleOp()

	if !ok {
		panic("arm CreateRoleOp is not set")
	}

	return val
}

// GetCreateRoleOp retrieves the CreateRoleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateRoleOp() (result CreateRoleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateRoleOp" {
		result = *u.CreateRoleOp
		ok = true
	}

	return
}

// MustUpdateRoleOp retrieves the UpdateRoleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustUpdateRoleOp() UpdateRoleOp {
	val, ok := u.GetUpdateRoleOp()

	if !ok {
		panic("arm UpdateRoleOp is not set")
	}

	return val
}

// GetUpdateRoleOp retrieves the UpdateRoleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetUpdateRoleOp() (result UpdateRoleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateRoleOp" {
		result = *u.UpdateRoleOp
		ok = true
	}

	return
}

// MustRemoveRoleOp retrieves the RemoveRoleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRemoveRoleOp() RemoveRoleOp {
	val, ok := u.GetRemoveRoleOp()

	if !ok {
		panic("arm RemoveRoleOp is not set")
	}

	return val
}

// GetRemoveRoleOp retrieves the RemoveRoleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRemoveRoleOp() (result RemoveRoleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveRoleOp" {
		result = *u.RemoveRoleOp
		ok = true
	}

	return
}

// MustCreateRuleOp retrieves the CreateRuleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateRuleOp() CreateRuleOp {
	val, ok := u.GetCreateRuleOp()

	if !ok {
		panic("arm CreateRuleOp is not set")
	}

	return val
}

// GetCreateRuleOp retrieves the CreateRuleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateRuleOp() (result CreateRuleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateRuleOp" {
		result = *u.CreateRuleOp
		ok = true
	}

	return
}

// MustUpdateRuleOp retrieves the UpdateRuleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustUpdateRuleOp() UpdateRuleOp {
	val, ok := u.GetUpdateRuleOp()

	if !ok {
		panic("arm UpdateRuleOp is not set")
	}

	return val
}

// GetUpdateRuleOp retrieves the UpdateRuleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetUpdateRuleOp() (result UpdateRuleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateRuleOp" {
		result = *u.UpdateRuleOp
		ok = true
	}

	return
}

// MustRemoveRuleOp retrieves the RemoveRuleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRemoveRuleOp() RemoveRuleOp {
	val, ok := u.GetRemoveRuleOp()

	if !ok {
		panic("arm RemoveRuleOp is not set")
	}

	return val
}

// GetRemoveRuleOp retrieves the RemoveRuleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRemoveRuleOp() (result RemoveRuleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveRuleOp" {
		result = *u.RemoveRuleOp
		ok = true
	}

	return
}

// MustCreateReviewableRequestOp retrieves the CreateReviewableRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateReviewableRequestOp() CreateReviewableRequestOp {
	val, ok := u.GetCreateReviewableRequestOp()

	if !ok {
		panic("arm CreateReviewableRequestOp is not set")
	}

	return val
}

// GetCreateReviewableRequestOp retrieves the CreateReviewableRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateReviewableRequestOp() (result CreateReviewableRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateReviewableRequestOp" {
		result = *u.CreateReviewableRequestOp
		ok = true
	}

	return
}

// MustUpdateReviewableRequestOp retrieves the UpdateReviewableRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustUpdateReviewableRequestOp() UpdateReviewableRequestOp {
	val, ok := u.GetUpdateReviewableRequestOp()

	if !ok {
		panic("arm UpdateReviewableRequestOp is not set")
	}

	return val
}

// GetUpdateReviewableRequestOp retrieves the UpdateReviewableRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetUpdateReviewableRequestOp() (result UpdateReviewableRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateReviewableRequestOp" {
		result = *u.UpdateReviewableRequestOp
		ok = true
	}

	return
}

// MustRemoveReviewableRequestOp retrieves the RemoveReviewableRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRemoveReviewableRequestOp() RemoveReviewableRequestOp {
	val, ok := u.GetRemoveReviewableRequestOp()

	if !ok {
		panic("arm RemoveReviewableRequestOp is not set")
	}

	return val
}

// GetRemoveReviewableRequestOp retrieves the RemoveReviewableRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRemoveReviewableRequestOp() (result RemoveReviewableRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveReviewableRequestOp" {
		result = *u.RemoveReviewableRequestOp
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

// MustKycRecoveryOp retrieves the KycRecoveryOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustKycRecoveryOp() KycRecoveryOp {
	val, ok := u.GetKycRecoveryOp()

	if !ok {
		panic("arm KycRecoveryOp is not set")
	}

	return val
}

// GetKycRecoveryOp retrieves the KycRecoveryOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetKycRecoveryOp() (result KycRecoveryOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KycRecoveryOp" {
		result = *u.KycRecoveryOp
		ok = true
	}

	return
}

// MustIssuanceOp retrieves the IssuanceOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustIssuanceOp() IssuanceOp {
	val, ok := u.GetIssuanceOp()

	if !ok {
		panic("arm IssuanceOp is not set")
	}

	return val
}

// GetIssuanceOp retrieves the IssuanceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetIssuanceOp() (result IssuanceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "IssuanceOp" {
		result = *u.IssuanceOp
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
//    	case DESTRUCTION:
//    		DestructionOp destructionOp;
//    	case CREATE_BALANCE:
//    		CreateBalanceOp createBalanceOp;
//        case CREATE_ASSET:
//            CreateAssetOp createAssetOp;
//        case UPDATE_ASSET:
//            UpdateAssetOp updateAssetOp;
//        case CREATE_DATA:
//            CreateDataOp createDataOp;
//        case UPDATE_DATA:
//            UpdateDataOp updateDataOp;
//        case REMOVE_DATA:
//            RemoveDataOp removeDataOp;
//        case REVIEW_REQUEST:
//    		ReviewRequestOp reviewRequestOp;
//    	case PUT_KEY_VALUE:
//    	    PutKeyValueOp putKeyValueOp;
//        case REMOVE_KEY_VALUE:
//    	    RemoveKeyValueOp removeKeyValueOp;
//    	case CHANGE_ACCOUNT_ROLES:
//    		ChangeAccountRolesOp changeAccountRolesOp;
//        case PAYMENT:
//            PaymentOp paymentOp;
//        case CREATE_SIGNER:
//            CreateSignerOp createSignerOp;
//        case UPDATE_SIGNER:
//            UpdateSignerOp updateSignerOp;
//        case REMOVE_SIGNER:
//            RemoveSignerOp removeSignerOp;
//        case CREATE_ROLE:
//            CreateRoleOp createRoleOp;
//        case UPDATE_ROLE:
//            UpdateRoleOp updateRoleOp;
//        case REMOVE_ROLE:
//            RemoveRoleOp removeRoleOp;
//        case CREATE_RULE:
//            CreateRuleOp createRuleOp;
//        case UPDATE_RULE:
//            UpdateRuleOp updateRuleOp;
//        case REMOVE_RULE:
//            RemoveRuleOp removeRuleOp;
//        case CREATE_REVIEWABLE_REQUEST:
//            CreateReviewableRequestOp createReviewableRequestOp;
//        case UPDATE_REVIEWABLE_REQUEST:
//            UpdateReviewableRequestOp updateReviewableRequestOp;
//        case REMOVE_REVIEWABLE_REQUEST:
//            RemoveReviewableRequestOp removeReviewableRequestOp;
//        case INITIATE_KYC_RECOVERY:
//            InitiateKYCRecoveryOp initiateKYCRecoveryOp;
//        case KYC_RECOVERY:
//            KYCRecoveryOp kycRecoveryOp;
//        case ISSUANCE:
//            IssuanceOp issuanceOp;
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
//        opBAD_AUTH_EXTRA = -8,
//        opNO_ROLE_PERMISSION = -9, // not allowed for this role of source account
//        opNO_ENTRY = -10,
//        opNOT_SUPPORTED = -11,
//        opLICENSE_VIOLATION = -12, // number of admins is greater than allowed
//        //: operation was skipped cause of failure validation of previous operation
//        opSKIPPED = -13
//    };
//
type OperationResultCode int32

const (
	OperationResultCodeOpInner            OperationResultCode = 0
	OperationResultCodeOpBadAuth          OperationResultCode = -1
	OperationResultCodeOpNoAccount        OperationResultCode = -2
	OperationResultCodeOpNotAllowed       OperationResultCode = -3
	OperationResultCodeOpAccountBlocked   OperationResultCode = -4
	OperationResultCodeOpBadAuthExtra     OperationResultCode = -8
	OperationResultCodeOpNoRolePermission OperationResultCode = -9
	OperationResultCodeOpNoEntry          OperationResultCode = -10
	OperationResultCodeOpNotSupported     OperationResultCode = -11
	OperationResultCodeOpLicenseViolation OperationResultCode = -12
	OperationResultCodeOpSkipped          OperationResultCode = -13
)

var OperationResultCodeAll = []OperationResultCode{
	OperationResultCodeOpInner,
	OperationResultCodeOpBadAuth,
	OperationResultCodeOpNoAccount,
	OperationResultCodeOpNotAllowed,
	OperationResultCodeOpAccountBlocked,
	OperationResultCodeOpBadAuthExtra,
	OperationResultCodeOpNoRolePermission,
	OperationResultCodeOpNoEntry,
	OperationResultCodeOpNotSupported,
	OperationResultCodeOpLicenseViolation,
	OperationResultCodeOpSkipped,
}

var operationResultCodeMap = map[int32]string{
	0:   "OperationResultCodeOpInner",
	-1:  "OperationResultCodeOpBadAuth",
	-2:  "OperationResultCodeOpNoAccount",
	-3:  "OperationResultCodeOpNotAllowed",
	-4:  "OperationResultCodeOpAccountBlocked",
	-8:  "OperationResultCodeOpBadAuthExtra",
	-9:  "OperationResultCodeOpNoRolePermission",
	-10: "OperationResultCodeOpNoEntry",
	-11: "OperationResultCodeOpNotSupported",
	-12: "OperationResultCodeOpLicenseViolation",
	-13: "OperationResultCodeOpSkipped",
}

var operationResultCodeShortMap = map[int32]string{
	0:   "op_inner",
	-1:  "op_bad_auth",
	-2:  "op_no_account",
	-3:  "op_not_allowed",
	-4:  "op_account_blocked",
	-8:  "op_bad_auth_extra",
	-9:  "op_no_role_permission",
	-10: "op_no_entry",
	-11: "op_not_supported",
	-12: "op_license_violation",
	-13: "op_skipped",
}

var operationResultCodeRevMap = map[string]int32{
	"OperationResultCodeOpInner":            0,
	"OperationResultCodeOpBadAuth":          -1,
	"OperationResultCodeOpNoAccount":        -2,
	"OperationResultCodeOpNotAllowed":       -3,
	"OperationResultCodeOpAccountBlocked":   -4,
	"OperationResultCodeOpBadAuthExtra":     -8,
	"OperationResultCodeOpNoRolePermission": -9,
	"OperationResultCodeOpNoEntry":          -10,
	"OperationResultCodeOpNotSupported":     -11,
	"OperationResultCodeOpLicenseViolation": -12,
	"OperationResultCodeOpSkipped":          -13,
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

// RuleRequirement is an XDR Struct defines as:
//
//   //: Defines requirements for tx or operation which were not fulfilled
//    struct RuleRequirement
//    {
//    	//: defines resources to which access was denied
//        RuleResource resource;
//    	//: defines action which was denied
//        RuleAction action;
//    	//: defines account for which requirements were not met
//    	AccountID account;
//
//    	//: reserved for future extension
//        EmptyExt ext;
//    };
//
type RuleRequirement struct {
	Resource RuleResource `json:"resource,omitempty"`
	Action   RuleAction   `json:"action,omitempty"`
	Account  AccountId    `json:"account,omitempty"`
	Ext      EmptyExt     `json:"ext,omitempty"`
}

// OperationResultTr is an XDR Union defines as:
//
//   union OperationResultTr switch (OperationType type)
//    {
//    case CREATE_ACCOUNT:
//        CreateAccountResult createAccountResult;
//    case DESTRUCTION:
//        DestructionResult destructionResult;
//    case CREATE_BALANCE:
//        CreateBalanceResult createBalanceResult;
//    case CREATE_ASSET:
//        CreateAssetResult createAssetResult;
//    case UPDATE_ASSET:
//        UpdateAssetResult updateAssetResult;
//    case CREATE_DATA:
//        CreateDataResult createDataResult;
//    case UPDATE_DATA:
//        UpdateDataResult updateDataResult;
//    case REMOVE_DATA:
//        RemoveDataResult removeDataResult;
//    case REVIEW_REQUEST:
//        ReviewRequestResult reviewRequestResult;
//    case PUT_KEY_VALUE:
//        PutKeyValueResult putKeyValueResult;
//    case REMOVE_KEY_VALUE:
//        RemoveKeyValueResult removeKeyValueResult;
//    case CHANGE_ACCOUNT_ROLES:
//        ChangeAccountRolesResult changeAccountRolesResult;
//    case PAYMENT:
//        PaymentResult paymentResult;
//    case CREATE_SIGNER:
//        CreateSignerResult createSignerResult;
//    case UPDATE_SIGNER:
//        UpdateSignerResult updateSignerResult;
//    case REMOVE_SIGNER:
//        RemoveSignerResult removeSignerResult;
//    case CREATE_ROLE:
//        CreateRoleResult createRoleResult;
//    case UPDATE_ROLE:
//        UpdateRoleResult updateRoleResult;
//    case REMOVE_ROLE:
//        RemoveRoleResult removeRoleResult;
//    case CREATE_RULE:
//        CreateRuleResult createRuleResult;
//    case UPDATE_RULE:
//        UpdateRuleResult updateRuleResult;
//    case REMOVE_RULE:
//        RemoveRuleResult removeRuleResult;
//    case CREATE_REVIEWABLE_REQUEST:
//        CreateReviewableRequestResult createReviewableRequestResult;
//    case UPDATE_REVIEWABLE_REQUEST:
//        UpdateReviewableRequestResult updateReviewableRequestResult;
//    case REMOVE_REVIEWABLE_REQUEST:
//        RemoveReviewableRequestResult removeReviewableRequestResult;
//    case KYC_RECOVERY:
//        KYCRecoveryResult kycRecoveryResult;
//    case INITIATE_KYC_RECOVERY:
//        InitiateKYCRecoveryResult initiateKYCRecoveryResult;
//    case ISSUANCE:
//        IssuanceResult issuanceResult;
//    };
//
type OperationResultTr struct {
	Type                          OperationType                  `json:"type,omitempty"`
	CreateAccountResult           *CreateAccountResult           `json:"createAccountResult,omitempty"`
	DestructionResult             *DestructionResult             `json:"destructionResult,omitempty"`
	CreateBalanceResult           *CreateBalanceResult           `json:"createBalanceResult,omitempty"`
	CreateAssetResult             *CreateAssetResult             `json:"createAssetResult,omitempty"`
	UpdateAssetResult             *UpdateAssetResult             `json:"updateAssetResult,omitempty"`
	CreateDataResult              *CreateDataResult              `json:"createDataResult,omitempty"`
	UpdateDataResult              *UpdateDataResult              `json:"updateDataResult,omitempty"`
	RemoveDataResult              *RemoveDataResult              `json:"removeDataResult,omitempty"`
	ReviewRequestResult           *ReviewRequestResult           `json:"reviewRequestResult,omitempty"`
	PutKeyValueResult             *PutKeyValueResult             `json:"putKeyValueResult,omitempty"`
	RemoveKeyValueResult          *RemoveKeyValueResult          `json:"removeKeyValueResult,omitempty"`
	ChangeAccountRolesResult      *ChangeAccountRolesResult      `json:"changeAccountRolesResult,omitempty"`
	PaymentResult                 *PaymentResult                 `json:"paymentResult,omitempty"`
	CreateSignerResult            *CreateSignerResult            `json:"createSignerResult,omitempty"`
	UpdateSignerResult            *UpdateSignerResult            `json:"updateSignerResult,omitempty"`
	RemoveSignerResult            *RemoveSignerResult            `json:"removeSignerResult,omitempty"`
	CreateRoleResult              *CreateRoleResult              `json:"createRoleResult,omitempty"`
	UpdateRoleResult              *UpdateRoleResult              `json:"updateRoleResult,omitempty"`
	RemoveRoleResult              *RemoveRoleResult              `json:"removeRoleResult,omitempty"`
	CreateRuleResult              *CreateRuleResult              `json:"createRuleResult,omitempty"`
	UpdateRuleResult              *UpdateRuleResult              `json:"updateRuleResult,omitempty"`
	RemoveRuleResult              *RemoveRuleResult              `json:"removeRuleResult,omitempty"`
	CreateReviewableRequestResult *CreateReviewableRequestResult `json:"createReviewableRequestResult,omitempty"`
	UpdateReviewableRequestResult *UpdateReviewableRequestResult `json:"updateReviewableRequestResult,omitempty"`
	RemoveReviewableRequestResult *RemoveReviewableRequestResult `json:"removeReviewableRequestResult,omitempty"`
	KycRecoveryResult             *KycRecoveryResult             `json:"kycRecoveryResult,omitempty"`
	InitiateKycRecoveryResult     *InitiateKycRecoveryResult     `json:"initiateKYCRecoveryResult,omitempty"`
	IssuanceResult                *IssuanceResult                `json:"issuanceResult,omitempty"`
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
	case OperationTypeDestruction:
		return "DestructionResult", true
	case OperationTypeCreateBalance:
		return "CreateBalanceResult", true
	case OperationTypeCreateAsset:
		return "CreateAssetResult", true
	case OperationTypeUpdateAsset:
		return "UpdateAssetResult", true
	case OperationTypeCreateData:
		return "CreateDataResult", true
	case OperationTypeUpdateData:
		return "UpdateDataResult", true
	case OperationTypeRemoveData:
		return "RemoveDataResult", true
	case OperationTypeReviewRequest:
		return "ReviewRequestResult", true
	case OperationTypePutKeyValue:
		return "PutKeyValueResult", true
	case OperationTypeRemoveKeyValue:
		return "RemoveKeyValueResult", true
	case OperationTypeChangeAccountRoles:
		return "ChangeAccountRolesResult", true
	case OperationTypePayment:
		return "PaymentResult", true
	case OperationTypeCreateSigner:
		return "CreateSignerResult", true
	case OperationTypeUpdateSigner:
		return "UpdateSignerResult", true
	case OperationTypeRemoveSigner:
		return "RemoveSignerResult", true
	case OperationTypeCreateRole:
		return "CreateRoleResult", true
	case OperationTypeUpdateRole:
		return "UpdateRoleResult", true
	case OperationTypeRemoveRole:
		return "RemoveRoleResult", true
	case OperationTypeCreateRule:
		return "CreateRuleResult", true
	case OperationTypeUpdateRule:
		return "UpdateRuleResult", true
	case OperationTypeRemoveRule:
		return "RemoveRuleResult", true
	case OperationTypeCreateReviewableRequest:
		return "CreateReviewableRequestResult", true
	case OperationTypeUpdateReviewableRequest:
		return "UpdateReviewableRequestResult", true
	case OperationTypeRemoveReviewableRequest:
		return "RemoveReviewableRequestResult", true
	case OperationTypeKycRecovery:
		return "KycRecoveryResult", true
	case OperationTypeInitiateKycRecovery:
		return "InitiateKycRecoveryResult", true
	case OperationTypeIssuance:
		return "IssuanceResult", true
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
	case OperationTypeDestruction:
		tv, ok := value.(DestructionResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be DestructionResult")
			return
		}
		result.DestructionResult = &tv
	case OperationTypeCreateBalance:
		tv, ok := value.(CreateBalanceResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateBalanceResult")
			return
		}
		result.CreateBalanceResult = &tv
	case OperationTypeCreateAsset:
		tv, ok := value.(CreateAssetResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAssetResult")
			return
		}
		result.CreateAssetResult = &tv
	case OperationTypeUpdateAsset:
		tv, ok := value.(UpdateAssetResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateAssetResult")
			return
		}
		result.UpdateAssetResult = &tv
	case OperationTypeCreateData:
		tv, ok := value.(CreateDataResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateDataResult")
			return
		}
		result.CreateDataResult = &tv
	case OperationTypeUpdateData:
		tv, ok := value.(UpdateDataResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateDataResult")
			return
		}
		result.UpdateDataResult = &tv
	case OperationTypeRemoveData:
		tv, ok := value.(RemoveDataResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveDataResult")
			return
		}
		result.RemoveDataResult = &tv
	case OperationTypeReviewRequest:
		tv, ok := value.(ReviewRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestResult")
			return
		}
		result.ReviewRequestResult = &tv
	case OperationTypePutKeyValue:
		tv, ok := value.(PutKeyValueResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be PutKeyValueResult")
			return
		}
		result.PutKeyValueResult = &tv
	case OperationTypeRemoveKeyValue:
		tv, ok := value.(RemoveKeyValueResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveKeyValueResult")
			return
		}
		result.RemoveKeyValueResult = &tv
	case OperationTypeChangeAccountRoles:
		tv, ok := value.(ChangeAccountRolesResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ChangeAccountRolesResult")
			return
		}
		result.ChangeAccountRolesResult = &tv
	case OperationTypePayment:
		tv, ok := value.(PaymentResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentResult")
			return
		}
		result.PaymentResult = &tv
	case OperationTypeCreateSigner:
		tv, ok := value.(CreateSignerResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateSignerResult")
			return
		}
		result.CreateSignerResult = &tv
	case OperationTypeUpdateSigner:
		tv, ok := value.(UpdateSignerResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateSignerResult")
			return
		}
		result.UpdateSignerResult = &tv
	case OperationTypeRemoveSigner:
		tv, ok := value.(RemoveSignerResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveSignerResult")
			return
		}
		result.RemoveSignerResult = &tv
	case OperationTypeCreateRole:
		tv, ok := value.(CreateRoleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateRoleResult")
			return
		}
		result.CreateRoleResult = &tv
	case OperationTypeUpdateRole:
		tv, ok := value.(UpdateRoleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateRoleResult")
			return
		}
		result.UpdateRoleResult = &tv
	case OperationTypeRemoveRole:
		tv, ok := value.(RemoveRoleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveRoleResult")
			return
		}
		result.RemoveRoleResult = &tv
	case OperationTypeCreateRule:
		tv, ok := value.(CreateRuleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateRuleResult")
			return
		}
		result.CreateRuleResult = &tv
	case OperationTypeUpdateRule:
		tv, ok := value.(UpdateRuleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateRuleResult")
			return
		}
		result.UpdateRuleResult = &tv
	case OperationTypeRemoveRule:
		tv, ok := value.(RemoveRuleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveRuleResult")
			return
		}
		result.RemoveRuleResult = &tv
	case OperationTypeCreateReviewableRequest:
		tv, ok := value.(CreateReviewableRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateReviewableRequestResult")
			return
		}
		result.CreateReviewableRequestResult = &tv
	case OperationTypeUpdateReviewableRequest:
		tv, ok := value.(UpdateReviewableRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateReviewableRequestResult")
			return
		}
		result.UpdateReviewableRequestResult = &tv
	case OperationTypeRemoveReviewableRequest:
		tv, ok := value.(RemoveReviewableRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RemoveReviewableRequestResult")
			return
		}
		result.RemoveReviewableRequestResult = &tv
	case OperationTypeKycRecovery:
		tv, ok := value.(KycRecoveryResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be KycRecoveryResult")
			return
		}
		result.KycRecoveryResult = &tv
	case OperationTypeInitiateKycRecovery:
		tv, ok := value.(InitiateKycRecoveryResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be InitiateKycRecoveryResult")
			return
		}
		result.InitiateKycRecoveryResult = &tv
	case OperationTypeIssuance:
		tv, ok := value.(IssuanceResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be IssuanceResult")
			return
		}
		result.IssuanceResult = &tv
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

// MustDestructionResult retrieves the DestructionResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustDestructionResult() DestructionResult {
	val, ok := u.GetDestructionResult()

	if !ok {
		panic("arm DestructionResult is not set")
	}

	return val
}

// GetDestructionResult retrieves the DestructionResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetDestructionResult() (result DestructionResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DestructionResult" {
		result = *u.DestructionResult
		ok = true
	}

	return
}

// MustCreateBalanceResult retrieves the CreateBalanceResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateBalanceResult() CreateBalanceResult {
	val, ok := u.GetCreateBalanceResult()

	if !ok {
		panic("arm CreateBalanceResult is not set")
	}

	return val
}

// GetCreateBalanceResult retrieves the CreateBalanceResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateBalanceResult() (result CreateBalanceResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateBalanceResult" {
		result = *u.CreateBalanceResult
		ok = true
	}

	return
}

// MustCreateAssetResult retrieves the CreateAssetResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateAssetResult() CreateAssetResult {
	val, ok := u.GetCreateAssetResult()

	if !ok {
		panic("arm CreateAssetResult is not set")
	}

	return val
}

// GetCreateAssetResult retrieves the CreateAssetResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateAssetResult() (result CreateAssetResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAssetResult" {
		result = *u.CreateAssetResult
		ok = true
	}

	return
}

// MustUpdateAssetResult retrieves the UpdateAssetResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustUpdateAssetResult() UpdateAssetResult {
	val, ok := u.GetUpdateAssetResult()

	if !ok {
		panic("arm UpdateAssetResult is not set")
	}

	return val
}

// GetUpdateAssetResult retrieves the UpdateAssetResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetUpdateAssetResult() (result UpdateAssetResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateAssetResult" {
		result = *u.UpdateAssetResult
		ok = true
	}

	return
}

// MustCreateDataResult retrieves the CreateDataResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateDataResult() CreateDataResult {
	val, ok := u.GetCreateDataResult()

	if !ok {
		panic("arm CreateDataResult is not set")
	}

	return val
}

// GetCreateDataResult retrieves the CreateDataResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateDataResult() (result CreateDataResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateDataResult" {
		result = *u.CreateDataResult
		ok = true
	}

	return
}

// MustUpdateDataResult retrieves the UpdateDataResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustUpdateDataResult() UpdateDataResult {
	val, ok := u.GetUpdateDataResult()

	if !ok {
		panic("arm UpdateDataResult is not set")
	}

	return val
}

// GetUpdateDataResult retrieves the UpdateDataResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetUpdateDataResult() (result UpdateDataResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateDataResult" {
		result = *u.UpdateDataResult
		ok = true
	}

	return
}

// MustRemoveDataResult retrieves the RemoveDataResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRemoveDataResult() RemoveDataResult {
	val, ok := u.GetRemoveDataResult()

	if !ok {
		panic("arm RemoveDataResult is not set")
	}

	return val
}

// GetRemoveDataResult retrieves the RemoveDataResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRemoveDataResult() (result RemoveDataResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveDataResult" {
		result = *u.RemoveDataResult
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

// MustPutKeyValueResult retrieves the PutKeyValueResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustPutKeyValueResult() PutKeyValueResult {
	val, ok := u.GetPutKeyValueResult()

	if !ok {
		panic("arm PutKeyValueResult is not set")
	}

	return val
}

// GetPutKeyValueResult retrieves the PutKeyValueResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetPutKeyValueResult() (result PutKeyValueResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PutKeyValueResult" {
		result = *u.PutKeyValueResult
		ok = true
	}

	return
}

// MustRemoveKeyValueResult retrieves the RemoveKeyValueResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRemoveKeyValueResult() RemoveKeyValueResult {
	val, ok := u.GetRemoveKeyValueResult()

	if !ok {
		panic("arm RemoveKeyValueResult is not set")
	}

	return val
}

// GetRemoveKeyValueResult retrieves the RemoveKeyValueResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRemoveKeyValueResult() (result RemoveKeyValueResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveKeyValueResult" {
		result = *u.RemoveKeyValueResult
		ok = true
	}

	return
}

// MustChangeAccountRolesResult retrieves the ChangeAccountRolesResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustChangeAccountRolesResult() ChangeAccountRolesResult {
	val, ok := u.GetChangeAccountRolesResult()

	if !ok {
		panic("arm ChangeAccountRolesResult is not set")
	}

	return val
}

// GetChangeAccountRolesResult retrieves the ChangeAccountRolesResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetChangeAccountRolesResult() (result ChangeAccountRolesResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ChangeAccountRolesResult" {
		result = *u.ChangeAccountRolesResult
		ok = true
	}

	return
}

// MustPaymentResult retrieves the PaymentResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustPaymentResult() PaymentResult {
	val, ok := u.GetPaymentResult()

	if !ok {
		panic("arm PaymentResult is not set")
	}

	return val
}

// GetPaymentResult retrieves the PaymentResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetPaymentResult() (result PaymentResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentResult" {
		result = *u.PaymentResult
		ok = true
	}

	return
}

// MustCreateSignerResult retrieves the CreateSignerResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateSignerResult() CreateSignerResult {
	val, ok := u.GetCreateSignerResult()

	if !ok {
		panic("arm CreateSignerResult is not set")
	}

	return val
}

// GetCreateSignerResult retrieves the CreateSignerResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateSignerResult() (result CreateSignerResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateSignerResult" {
		result = *u.CreateSignerResult
		ok = true
	}

	return
}

// MustUpdateSignerResult retrieves the UpdateSignerResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustUpdateSignerResult() UpdateSignerResult {
	val, ok := u.GetUpdateSignerResult()

	if !ok {
		panic("arm UpdateSignerResult is not set")
	}

	return val
}

// GetUpdateSignerResult retrieves the UpdateSignerResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetUpdateSignerResult() (result UpdateSignerResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateSignerResult" {
		result = *u.UpdateSignerResult
		ok = true
	}

	return
}

// MustRemoveSignerResult retrieves the RemoveSignerResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRemoveSignerResult() RemoveSignerResult {
	val, ok := u.GetRemoveSignerResult()

	if !ok {
		panic("arm RemoveSignerResult is not set")
	}

	return val
}

// GetRemoveSignerResult retrieves the RemoveSignerResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRemoveSignerResult() (result RemoveSignerResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveSignerResult" {
		result = *u.RemoveSignerResult
		ok = true
	}

	return
}

// MustCreateRoleResult retrieves the CreateRoleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateRoleResult() CreateRoleResult {
	val, ok := u.GetCreateRoleResult()

	if !ok {
		panic("arm CreateRoleResult is not set")
	}

	return val
}

// GetCreateRoleResult retrieves the CreateRoleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateRoleResult() (result CreateRoleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateRoleResult" {
		result = *u.CreateRoleResult
		ok = true
	}

	return
}

// MustUpdateRoleResult retrieves the UpdateRoleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustUpdateRoleResult() UpdateRoleResult {
	val, ok := u.GetUpdateRoleResult()

	if !ok {
		panic("arm UpdateRoleResult is not set")
	}

	return val
}

// GetUpdateRoleResult retrieves the UpdateRoleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetUpdateRoleResult() (result UpdateRoleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateRoleResult" {
		result = *u.UpdateRoleResult
		ok = true
	}

	return
}

// MustRemoveRoleResult retrieves the RemoveRoleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRemoveRoleResult() RemoveRoleResult {
	val, ok := u.GetRemoveRoleResult()

	if !ok {
		panic("arm RemoveRoleResult is not set")
	}

	return val
}

// GetRemoveRoleResult retrieves the RemoveRoleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRemoveRoleResult() (result RemoveRoleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveRoleResult" {
		result = *u.RemoveRoleResult
		ok = true
	}

	return
}

// MustCreateRuleResult retrieves the CreateRuleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateRuleResult() CreateRuleResult {
	val, ok := u.GetCreateRuleResult()

	if !ok {
		panic("arm CreateRuleResult is not set")
	}

	return val
}

// GetCreateRuleResult retrieves the CreateRuleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateRuleResult() (result CreateRuleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateRuleResult" {
		result = *u.CreateRuleResult
		ok = true
	}

	return
}

// MustUpdateRuleResult retrieves the UpdateRuleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustUpdateRuleResult() UpdateRuleResult {
	val, ok := u.GetUpdateRuleResult()

	if !ok {
		panic("arm UpdateRuleResult is not set")
	}

	return val
}

// GetUpdateRuleResult retrieves the UpdateRuleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetUpdateRuleResult() (result UpdateRuleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateRuleResult" {
		result = *u.UpdateRuleResult
		ok = true
	}

	return
}

// MustRemoveRuleResult retrieves the RemoveRuleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRemoveRuleResult() RemoveRuleResult {
	val, ok := u.GetRemoveRuleResult()

	if !ok {
		panic("arm RemoveRuleResult is not set")
	}

	return val
}

// GetRemoveRuleResult retrieves the RemoveRuleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRemoveRuleResult() (result RemoveRuleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveRuleResult" {
		result = *u.RemoveRuleResult
		ok = true
	}

	return
}

// MustCreateReviewableRequestResult retrieves the CreateReviewableRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateReviewableRequestResult() CreateReviewableRequestResult {
	val, ok := u.GetCreateReviewableRequestResult()

	if !ok {
		panic("arm CreateReviewableRequestResult is not set")
	}

	return val
}

// GetCreateReviewableRequestResult retrieves the CreateReviewableRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateReviewableRequestResult() (result CreateReviewableRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateReviewableRequestResult" {
		result = *u.CreateReviewableRequestResult
		ok = true
	}

	return
}

// MustUpdateReviewableRequestResult retrieves the UpdateReviewableRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustUpdateReviewableRequestResult() UpdateReviewableRequestResult {
	val, ok := u.GetUpdateReviewableRequestResult()

	if !ok {
		panic("arm UpdateReviewableRequestResult is not set")
	}

	return val
}

// GetUpdateReviewableRequestResult retrieves the UpdateReviewableRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetUpdateReviewableRequestResult() (result UpdateReviewableRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateReviewableRequestResult" {
		result = *u.UpdateReviewableRequestResult
		ok = true
	}

	return
}

// MustRemoveReviewableRequestResult retrieves the RemoveReviewableRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRemoveReviewableRequestResult() RemoveReviewableRequestResult {
	val, ok := u.GetRemoveReviewableRequestResult()

	if !ok {
		panic("arm RemoveReviewableRequestResult is not set")
	}

	return val
}

// GetRemoveReviewableRequestResult retrieves the RemoveReviewableRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRemoveReviewableRequestResult() (result RemoveReviewableRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RemoveReviewableRequestResult" {
		result = *u.RemoveReviewableRequestResult
		ok = true
	}

	return
}

// MustKycRecoveryResult retrieves the KycRecoveryResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustKycRecoveryResult() KycRecoveryResult {
	val, ok := u.GetKycRecoveryResult()

	if !ok {
		panic("arm KycRecoveryResult is not set")
	}

	return val
}

// GetKycRecoveryResult retrieves the KycRecoveryResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetKycRecoveryResult() (result KycRecoveryResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KycRecoveryResult" {
		result = *u.KycRecoveryResult
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

// MustIssuanceResult retrieves the IssuanceResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustIssuanceResult() IssuanceResult {
	val, ok := u.GetIssuanceResult()

	if !ok {
		panic("arm IssuanceResult is not set")
	}

	return val
}

// GetIssuanceResult retrieves the IssuanceResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetIssuanceResult() (result IssuanceResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "IssuanceResult" {
		result = *u.IssuanceResult
		ok = true
	}

	return
}

// OperationResult is an XDR Union defines as:
//
//   union OperationResult switch (OperationResultCode code)
//    {
//    case opINNER:
//        OperationResultTr tr;
//    case opNO_ENTRY:
//        LedgerKey entryKey;
//    case opNO_ROLE_PERMISSION:
//    case opBAD_AUTH:
//        RuleRequirement requirement;
//    default:
//        void;
//    };
//
type OperationResult struct {
	Code        OperationResultCode `json:"code,omitempty"`
	Tr          *OperationResultTr  `json:"tr,omitempty"`
	EntryKey    *LedgerKey          `json:"entryKey,omitempty"`
	Requirement *RuleRequirement    `json:"requirement,omitempty"`
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
		return "EntryKey", true
	case OperationResultCodeOpNoRolePermission:
		return "Requirement", true
	case OperationResultCodeOpBadAuth:
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
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.EntryKey = &tv
	case OperationResultCodeOpNoRolePermission:
		tv, ok := value.(RuleRequirement)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleRequirement")
			return
		}
		result.Requirement = &tv
	case OperationResultCodeOpBadAuth:
		tv, ok := value.(RuleRequirement)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleRequirement")
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

// MustEntryKey retrieves the EntryKey value from the union,
// panicing if the value is not set.
func (u OperationResult) MustEntryKey() LedgerKey {
	val, ok := u.GetEntryKey()

	if !ok {
		panic("arm EntryKey is not set")
	}

	return val
}

// GetEntryKey retrieves the EntryKey value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResult) GetEntryKey() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "EntryKey" {
		result = *u.EntryKey
		ok = true
	}

	return
}

// MustRequirement retrieves the Requirement value from the union,
// panicing if the value is not set.
func (u OperationResult) MustRequirement() RuleRequirement {
	val, ok := u.GetRequirement()

	if !ok {
		panic("arm Requirement is not set")
	}

	return val
}

// GetRequirement retrieves the Requirement value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResult) GetRequirement() (result RuleRequirement, ok bool) {
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
//        txINSUFFICIENT_FEE = -11,          // the actual total fee amount is greater than the max total fee amount, provided by the source
//        txSOURCE_UNDERFUNDED = -12,        // not enough tx fee asset on source balance
//        txCOMMISSION_LINE_FULL = -13,      // commission tx fee asset balance amount overflow
//        txFEE_INCORRECT_PRECISION = -14,   // fee amount is incompatible with asset precision
//        txNO_ROLE_PERMISSION = -15         // account role has not rule that allows send transaction
//    };
//
type TransactionResultCode int32

const (
	TransactionResultCodeTxSuccess               TransactionResultCode = 0
	TransactionResultCodeTxFailed                TransactionResultCode = -1
	TransactionResultCodeTxTooEarly              TransactionResultCode = -2
	TransactionResultCodeTxTooLate               TransactionResultCode = -3
	TransactionResultCodeTxMissingOperation      TransactionResultCode = -4
	TransactionResultCodeTxBadAuth               TransactionResultCode = -5
	TransactionResultCodeTxNoAccount             TransactionResultCode = -6
	TransactionResultCodeTxBadAuthExtra          TransactionResultCode = -7
	TransactionResultCodeTxInternalError         TransactionResultCode = -8
	TransactionResultCodeTxAccountBlocked        TransactionResultCode = -9
	TransactionResultCodeTxDuplication           TransactionResultCode = -10
	TransactionResultCodeTxInsufficientFee       TransactionResultCode = -11
	TransactionResultCodeTxSourceUnderfunded     TransactionResultCode = -12
	TransactionResultCodeTxCommissionLineFull    TransactionResultCode = -13
	TransactionResultCodeTxFeeIncorrectPrecision TransactionResultCode = -14
	TransactionResultCodeTxNoRolePermission      TransactionResultCode = -15
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
	TransactionResultCodeTxInsufficientFee,
	TransactionResultCodeTxSourceUnderfunded,
	TransactionResultCodeTxCommissionLineFull,
	TransactionResultCodeTxFeeIncorrectPrecision,
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
	-11: "TransactionResultCodeTxInsufficientFee",
	-12: "TransactionResultCodeTxSourceUnderfunded",
	-13: "TransactionResultCodeTxCommissionLineFull",
	-14: "TransactionResultCodeTxFeeIncorrectPrecision",
	-15: "TransactionResultCodeTxNoRolePermission",
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
	-11: "tx_insufficient_fee",
	-12: "tx_source_underfunded",
	-13: "tx_commission_line_full",
	-14: "tx_fee_incorrect_precision",
	-15: "tx_no_role_permission",
}

var transactionResultCodeRevMap = map[string]int32{
	"TransactionResultCodeTxSuccess":               0,
	"TransactionResultCodeTxFailed":                -1,
	"TransactionResultCodeTxTooEarly":              -2,
	"TransactionResultCodeTxTooLate":               -3,
	"TransactionResultCodeTxMissingOperation":      -4,
	"TransactionResultCodeTxBadAuth":               -5,
	"TransactionResultCodeTxNoAccount":             -6,
	"TransactionResultCodeTxBadAuthExtra":          -7,
	"TransactionResultCodeTxInternalError":         -8,
	"TransactionResultCodeTxAccountBlocked":        -9,
	"TransactionResultCodeTxDuplication":           -10,
	"TransactionResultCodeTxInsufficientFee":       -11,
	"TransactionResultCodeTxSourceUnderfunded":     -12,
	"TransactionResultCodeTxCommissionLineFull":    -13,
	"TransactionResultCodeTxFeeIncorrectPrecision": -14,
	"TransactionResultCodeTxNoRolePermission":      -15,
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

// OperationFeeExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type OperationFeeExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationFeeExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationFeeExt
func (u OperationFeeExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewOperationFeeExt creates a new  OperationFeeExt.
func NewOperationFeeExt(v LedgerVersion, value interface{}) (result OperationFeeExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// OperationFee is an XDR Struct defines as:
//
//   struct OperationFee
//    {
//        OperationType operationType;
//        uint64 amount;
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
type OperationFee struct {
	OperationType OperationType   `json:"operationType,omitempty"`
	Amount        Uint64          `json:"amount,omitempty"`
	Ext           OperationFeeExt `json:"ext,omitempty"`
}

// TransactionResultResult is an XDR NestedUnion defines as:
//
//   union switch (TransactionResultCode code)
//        {
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        case txNO_ROLE_PERMISSION:
//            RuleRequirement requirement;
//        default:
//            void;
//        }
//
type TransactionResultResult struct {
	Code        TransactionResultCode `json:"code,omitempty"`
	Results     *[]OperationResult    `json:"results,omitempty"`
	Requirement *RuleRequirement      `json:"requirement,omitempty"`
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
		tv, ok := value.(RuleRequirement)
		if !ok {
			err = fmt.Errorf("invalid value, must be RuleRequirement")
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
func (u TransactionResultResult) MustRequirement() RuleRequirement {
	val, ok := u.GetRequirement()

	if !ok {
		panic("arm Requirement is not set")
	}

	return val
}

// GetRequirement retrieves the Requirement value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionResultResult) GetRequirement() (result RuleRequirement, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Requirement" {
		result = *u.Requirement
		ok = true
	}

	return
}

// TransactionResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultExt
func (u TransactionResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionResultExt creates a new  TransactionResultExt.
func NewTransactionResultExt(v LedgerVersion, value interface{}) (result TransactionResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionResult is an XDR Struct defines as:
//
//   struct TransactionResult
//    {
//        int64 feeCharged; // actual fee charged for the transaction
//
//        union switch (TransactionResultCode code)
//        {
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        case txNO_ROLE_PERMISSION:
//            RuleRequirement requirement;
//        default:
//            void;
//        }
//        result;
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
type TransactionResult struct {
	FeeCharged Int64                   `json:"feeCharged,omitempty"`
	Result     TransactionResultResult `json:"result,omitempty"`
	Ext        TransactionResultExt    `json:"ext,omitempty"`
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
//        TRANSACTION = 0, // is used for account rule resource
//        ANY = 1, // is used for rules
//        ACCOUNT = 2,
//        SIGNER = 3,
//        BALANCE = 5,
//        DATA = 6,
//        ASSET = 7,
//        REFERENCE_ENTRY = 8,
//        REVIEWABLE_REQUEST = 15,
//    	ACCOUNT_KYC = 18,
//        KEY_VALUE = 20,
//        RULE = 30,
//        ROLE = 31
//    };
//
type LedgerEntryType int32

const (
	LedgerEntryTypeTransaction       LedgerEntryType = 0
	LedgerEntryTypeAny               LedgerEntryType = 1
	LedgerEntryTypeAccount           LedgerEntryType = 2
	LedgerEntryTypeSigner            LedgerEntryType = 3
	LedgerEntryTypeBalance           LedgerEntryType = 5
	LedgerEntryTypeData              LedgerEntryType = 6
	LedgerEntryTypeAsset             LedgerEntryType = 7
	LedgerEntryTypeReferenceEntry    LedgerEntryType = 8
	LedgerEntryTypeReviewableRequest LedgerEntryType = 15
	LedgerEntryTypeAccountKyc        LedgerEntryType = 18
	LedgerEntryTypeKeyValue          LedgerEntryType = 20
	LedgerEntryTypeRule              LedgerEntryType = 30
	LedgerEntryTypeRole              LedgerEntryType = 31
)

var LedgerEntryTypeAll = []LedgerEntryType{
	LedgerEntryTypeTransaction,
	LedgerEntryTypeAny,
	LedgerEntryTypeAccount,
	LedgerEntryTypeSigner,
	LedgerEntryTypeBalance,
	LedgerEntryTypeData,
	LedgerEntryTypeAsset,
	LedgerEntryTypeReferenceEntry,
	LedgerEntryTypeReviewableRequest,
	LedgerEntryTypeAccountKyc,
	LedgerEntryTypeKeyValue,
	LedgerEntryTypeRule,
	LedgerEntryTypeRole,
}

var ledgerEntryTypeMap = map[int32]string{
	0:  "LedgerEntryTypeTransaction",
	1:  "LedgerEntryTypeAny",
	2:  "LedgerEntryTypeAccount",
	3:  "LedgerEntryTypeSigner",
	5:  "LedgerEntryTypeBalance",
	6:  "LedgerEntryTypeData",
	7:  "LedgerEntryTypeAsset",
	8:  "LedgerEntryTypeReferenceEntry",
	15: "LedgerEntryTypeReviewableRequest",
	18: "LedgerEntryTypeAccountKyc",
	20: "LedgerEntryTypeKeyValue",
	30: "LedgerEntryTypeRule",
	31: "LedgerEntryTypeRole",
}

var ledgerEntryTypeShortMap = map[int32]string{
	0:  "transaction",
	1:  "any",
	2:  "account",
	3:  "signer",
	5:  "balance",
	6:  "data",
	7:  "asset",
	8:  "reference_entry",
	15: "reviewable_request",
	18: "account_kyc",
	20: "key_value",
	30: "rule",
	31: "role",
}

var ledgerEntryTypeRevMap = map[string]int32{
	"LedgerEntryTypeTransaction":       0,
	"LedgerEntryTypeAny":               1,
	"LedgerEntryTypeAccount":           2,
	"LedgerEntryTypeSigner":            3,
	"LedgerEntryTypeBalance":           5,
	"LedgerEntryTypeData":              6,
	"LedgerEntryTypeAsset":             7,
	"LedgerEntryTypeReferenceEntry":    8,
	"LedgerEntryTypeReviewableRequest": 15,
	"LedgerEntryTypeAccountKyc":        18,
	"LedgerEntryTypeKeyValue":          20,
	"LedgerEntryTypeRule":              30,
	"LedgerEntryTypeRole":              31,
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

// BalanceId is an XDR Typedef defines as:
//
//   typedef PublicKey BalanceID;
//
type BalanceId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BalanceId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u BalanceId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewBalanceId creates a new  BalanceId.
func NewBalanceId(aType CryptoKeyType, value interface{}) (result BalanceId, err error) {
	u, err := NewPublicKey(aType, value)
	result = BalanceId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u BalanceId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BalanceId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

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

// AssetCode is an XDR Typedef defines as:
//
//   typedef string AssetCode<16>;
//
type AssetCode string

// XDRMaxSize implements the Sized interface for AssetCode
func (e AssetCode) XDRMaxSize() int {
	return 16
}

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

// FeeExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//        {
//            case EMPTY_VERSION:
//                void;
//        }
//
type FeeExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u FeeExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of FeeExt
func (u FeeExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewFeeExt creates a new  FeeExt.
func NewFeeExt(v LedgerVersion, value interface{}) (result FeeExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Fee is an XDR Struct defines as:
//
//   //: `Fee` is used to unite fixed and percent fee amounts
//    struct Fee {
//        //: Fixed amount to pay for the operation
//    	uint64 fixed;
//    	//: Part of the managed amount in percents
//    	uint64 percent;
//
//        //: reserved for future use
//        union switch(LedgerVersion v)
//        {
//            case EMPTY_VERSION:
//                void;
//        }
//        ext;
//    };
//
type Fee struct {
	Fixed   Uint64 `json:"fixed,omitempty"`
	Percent Uint64 `json:"percent,omitempty"`
	Ext     FeeExt `json:"ext,omitempty"`
}

// OperationType is an XDR Enum defines as:
//
//   enum OperationType
//    {
//        CREATE_ACCOUNT = 1,
//        ISSUANCE = 3,
//        DESTRUCTION = 7,
//        CREATE_BALANCE = 9,
//        CREATE_ASSET = 11,
//        UPDATE_ASSET = 12,
//        CREATE_DATA = 14,
//        UPDATE_DATA = 15,
//        REMOVE_DATA = 16,
//        REVIEW_REQUEST = 18,
//    	CHANGE_ACCOUNT_ROLES = 22,
//        PAYMENT = 23,
//        PUT_KEY_VALUE = 27,
//        REMOVE_KEY_VALUE = 28,
//        CREATE_SIGNER = 30,
//        UPDATE_SIGNER = 31,
//        REMOVE_SIGNER = 32,
//        CREATE_ROLE = 39,
//        UPDATE_ROLE = 40,
//        REMOVE_ROLE = 41,
//        CREATE_RULE = 42,
//        UPDATE_RULE = 43,
//        REMOVE_RULE = 44,
//        CREATE_REVIEWABLE_REQUEST = 45,
//        UPDATE_REVIEWABLE_REQUEST = 46,
//        REMOVE_REVIEWABLE_REQUEST = 47,
//        INITIATE_KYC_RECOVERY = 48,
//        KYC_RECOVERY = 49
//    };
//
type OperationType int32

const (
	OperationTypeCreateAccount           OperationType = 1
	OperationTypeIssuance                OperationType = 3
	OperationTypeDestruction             OperationType = 7
	OperationTypeCreateBalance           OperationType = 9
	OperationTypeCreateAsset             OperationType = 11
	OperationTypeUpdateAsset             OperationType = 12
	OperationTypeCreateData              OperationType = 14
	OperationTypeUpdateData              OperationType = 15
	OperationTypeRemoveData              OperationType = 16
	OperationTypeReviewRequest           OperationType = 18
	OperationTypeChangeAccountRoles      OperationType = 22
	OperationTypePayment                 OperationType = 23
	OperationTypePutKeyValue             OperationType = 27
	OperationTypeRemoveKeyValue          OperationType = 28
	OperationTypeCreateSigner            OperationType = 30
	OperationTypeUpdateSigner            OperationType = 31
	OperationTypeRemoveSigner            OperationType = 32
	OperationTypeCreateRole              OperationType = 39
	OperationTypeUpdateRole              OperationType = 40
	OperationTypeRemoveRole              OperationType = 41
	OperationTypeCreateRule              OperationType = 42
	OperationTypeUpdateRule              OperationType = 43
	OperationTypeRemoveRule              OperationType = 44
	OperationTypeCreateReviewableRequest OperationType = 45
	OperationTypeUpdateReviewableRequest OperationType = 46
	OperationTypeRemoveReviewableRequest OperationType = 47
	OperationTypeInitiateKycRecovery     OperationType = 48
	OperationTypeKycRecovery             OperationType = 49
)

var OperationTypeAll = []OperationType{
	OperationTypeCreateAccount,
	OperationTypeIssuance,
	OperationTypeDestruction,
	OperationTypeCreateBalance,
	OperationTypeCreateAsset,
	OperationTypeUpdateAsset,
	OperationTypeCreateData,
	OperationTypeUpdateData,
	OperationTypeRemoveData,
	OperationTypeReviewRequest,
	OperationTypeChangeAccountRoles,
	OperationTypePayment,
	OperationTypePutKeyValue,
	OperationTypeRemoveKeyValue,
	OperationTypeCreateSigner,
	OperationTypeUpdateSigner,
	OperationTypeRemoveSigner,
	OperationTypeCreateRole,
	OperationTypeUpdateRole,
	OperationTypeRemoveRole,
	OperationTypeCreateRule,
	OperationTypeUpdateRule,
	OperationTypeRemoveRule,
	OperationTypeCreateReviewableRequest,
	OperationTypeUpdateReviewableRequest,
	OperationTypeRemoveReviewableRequest,
	OperationTypeInitiateKycRecovery,
	OperationTypeKycRecovery,
}

var operationTypeMap = map[int32]string{
	1:  "OperationTypeCreateAccount",
	3:  "OperationTypeIssuance",
	7:  "OperationTypeDestruction",
	9:  "OperationTypeCreateBalance",
	11: "OperationTypeCreateAsset",
	12: "OperationTypeUpdateAsset",
	14: "OperationTypeCreateData",
	15: "OperationTypeUpdateData",
	16: "OperationTypeRemoveData",
	18: "OperationTypeReviewRequest",
	22: "OperationTypeChangeAccountRoles",
	23: "OperationTypePayment",
	27: "OperationTypePutKeyValue",
	28: "OperationTypeRemoveKeyValue",
	30: "OperationTypeCreateSigner",
	31: "OperationTypeUpdateSigner",
	32: "OperationTypeRemoveSigner",
	39: "OperationTypeCreateRole",
	40: "OperationTypeUpdateRole",
	41: "OperationTypeRemoveRole",
	42: "OperationTypeCreateRule",
	43: "OperationTypeUpdateRule",
	44: "OperationTypeRemoveRule",
	45: "OperationTypeCreateReviewableRequest",
	46: "OperationTypeUpdateReviewableRequest",
	47: "OperationTypeRemoveReviewableRequest",
	48: "OperationTypeInitiateKycRecovery",
	49: "OperationTypeKycRecovery",
}

var operationTypeShortMap = map[int32]string{
	1:  "create_account",
	3:  "issuance",
	7:  "destruction",
	9:  "create_balance",
	11: "create_asset",
	12: "update_asset",
	14: "create_data",
	15: "update_data",
	16: "remove_data",
	18: "review_request",
	22: "change_account_roles",
	23: "payment",
	27: "put_key_value",
	28: "remove_key_value",
	30: "create_signer",
	31: "update_signer",
	32: "remove_signer",
	39: "create_role",
	40: "update_role",
	41: "remove_role",
	42: "create_rule",
	43: "update_rule",
	44: "remove_rule",
	45: "create_reviewable_request",
	46: "update_reviewable_request",
	47: "remove_reviewable_request",
	48: "initiate_kyc_recovery",
	49: "kyc_recovery",
}

var operationTypeRevMap = map[string]int32{
	"OperationTypeCreateAccount":           1,
	"OperationTypeIssuance":                3,
	"OperationTypeDestruction":             7,
	"OperationTypeCreateBalance":           9,
	"OperationTypeCreateAsset":             11,
	"OperationTypeUpdateAsset":             12,
	"OperationTypeCreateData":              14,
	"OperationTypeUpdateData":              15,
	"OperationTypeRemoveData":              16,
	"OperationTypeReviewRequest":           18,
	"OperationTypeChangeAccountRoles":      22,
	"OperationTypePayment":                 23,
	"OperationTypePutKeyValue":             27,
	"OperationTypeRemoveKeyValue":          28,
	"OperationTypeCreateSigner":            30,
	"OperationTypeUpdateSigner":            31,
	"OperationTypeRemoveSigner":            32,
	"OperationTypeCreateRole":              39,
	"OperationTypeUpdateRole":              40,
	"OperationTypeRemoveRole":              41,
	"OperationTypeCreateRule":              42,
	"OperationTypeUpdateRule":              43,
	"OperationTypeRemoveRule":              44,
	"OperationTypeCreateReviewableRequest": 45,
	"OperationTypeUpdateReviewableRequest": 46,
	"OperationTypeRemoveReviewableRequest": 47,
	"OperationTypeInitiateKycRecovery":     48,
	"OperationTypeKycRecovery":             49,
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

// DestinationType is an XDR Enum defines as:
//
//   //: Defines the type of destination for operation
//    enum DestinationType {
//        ACCOUNT = 0,
//        BALANCE = 1
//    };
//
type DestinationType int32

const (
	DestinationTypeAccount DestinationType = 0
	DestinationTypeBalance DestinationType = 1
)

var DestinationTypeAll = []DestinationType{
	DestinationTypeAccount,
	DestinationTypeBalance,
}

var destinationTypeMap = map[int32]string{
	0: "DestinationTypeAccount",
	1: "DestinationTypeBalance",
}

var destinationTypeShortMap = map[int32]string{
	0: "account",
	1: "balance",
}

var destinationTypeRevMap = map[string]int32{
	"DestinationTypeAccount": 0,
	"DestinationTypeBalance": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for DestinationType
func (e DestinationType) ValidEnum(v int32) bool {
	_, ok := destinationTypeMap[v]
	return ok
}
func (e DestinationType) isFlag() bool {
	for i := len(DestinationTypeAll) - 1; i >= 0; i-- {
		expected := DestinationType(2) << uint64(len(DestinationTypeAll)-1) >> uint64(len(DestinationTypeAll)-i)
		if expected != DestinationTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e DestinationType) String() string {
	name, _ := destinationTypeMap[int32(e)]
	return name
}

func (e DestinationType) ShortString() string {
	name, _ := destinationTypeShortMap[int32(e)]
	return name
}

func (e DestinationType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
			Flags: make([]flagValue, 0),
		}
		for _, value := range DestinationTypeAll {
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

func (e *DestinationType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = DestinationType(t.Value)
	return nil
}

// MovementDestination is an XDR Union defines as:
//
//   union MovementDestination switch (DestinationType type) {
//        case ACCOUNT:
//            AccountID accountID;
//        case BALANCE:
//            BalanceID balanceID;
//    };
//
type MovementDestination struct {
	Type      DestinationType `json:"type,omitempty"`
	AccountId *AccountId      `json:"accountID,omitempty"`
	BalanceId *BalanceId      `json:"balanceID,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u MovementDestination) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of MovementDestination
func (u MovementDestination) ArmForSwitch(sw int32) (string, bool) {
	switch DestinationType(sw) {
	case DestinationTypeAccount:
		return "AccountId", true
	case DestinationTypeBalance:
		return "BalanceId", true
	}
	return "-", false
}

// NewMovementDestination creates a new  MovementDestination.
func NewMovementDestination(aType DestinationType, value interface{}) (result MovementDestination, err error) {
	result.Type = aType
	switch DestinationType(aType) {
	case DestinationTypeAccount:
		tv, ok := value.(AccountId)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountId")
			return
		}
		result.AccountId = &tv
	case DestinationTypeBalance:
		tv, ok := value.(BalanceId)
		if !ok {
			err = fmt.Errorf("invalid value, must be BalanceId")
			return
		}
		result.BalanceId = &tv
	}
	return
}

// MustAccountId retrieves the AccountId value from the union,
// panicing if the value is not set.
func (u MovementDestination) MustAccountId() AccountId {
	val, ok := u.GetAccountId()

	if !ok {
		panic("arm AccountId is not set")
	}

	return val
}

// GetAccountId retrieves the AccountId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u MovementDestination) GetAccountId() (result AccountId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountId" {
		result = *u.AccountId
		ok = true
	}

	return
}

// MustBalanceId retrieves the BalanceId value from the union,
// panicing if the value is not set.
func (u MovementDestination) MustBalanceId() BalanceId {
	val, ok := u.GetBalanceId()

	if !ok {
		panic("arm BalanceId is not set")
	}

	return val
}

// GetBalanceId retrieves the BalanceId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u MovementDestination) GetBalanceId() (result BalanceId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BalanceId" {
		result = *u.BalanceId
		ok = true
	}

	return
}

var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
var Revision = "2efa51c931dc71565fcedc696e425d0bd065599a"
