package regources

import "gitlab.com/tokend/go/xdr"

// Fee - describes fee happened on balance. Direction of fee depends on the operation (depending on effect might be
// charged, locked, unlocked, for all incoming effects but unlocked it's always charged)
type Fee struct {
	Fixed             Amount `json:"fixed"`
	CalculatedPercent Amount `json:"calculated_percent"`
}

type FeeRecord struct {
	Key
	Attributes    FeeAttrs         `json:"attributes"`
	Relationships FeeRelationships `json:"relationships"`
}

type FeesResponse struct {
	Links    *Links      `json:"links"`
	Data     []FeeRecord `json:"data"`
	Included Included    `json:"included"`
}

type FeeRelationships struct {
	Account     *Relation `json:"account,omitempty"`
	AccountRole *Relation `json:"account_role,omitempty"`
	Asset       *Relation `json:"asset,omitempty"`
}

type FeeAttrs struct {
	Fixed     Amount       `json:"fixed"`
	Percent   Amount       `json:"percent"`
	AppliedTo FeeAppliedTo `json:"applied_to"`
}

type FeeAppliedTo struct {
	Asset           string      `json:"asset"`
	Subtype         int64       `json:"subtype"`
	FeeType         int32       `json:"fee_type"`
	FeeTypeExtended xdr.FeeType `json:"fee_type_extended"`
	LowerBound      Amount      `json:"lower_bound"`
	UpperBound      Amount      `json:"upper_bound"`
}

type CalculatedFee struct {
	Key
	Attributes Fee `json:"attributes"`
}

type CalculatedFeeResponse struct {
	Data CalculatedFee `json:"data"`
}
