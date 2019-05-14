package types

import (
	"net/url"
)

type UpdateKYCRequestsOpts struct {
	AccountToUpdateKYC string
	Order              string
	Cursor             string
}

func (o UpdateKYCRequestsOpts) Encode() string {
	query := url.Values{}

	if o.Order != "" {
		query.Set("order", o.Order)
	}

	if o.Cursor != "" {
		query.Set("cursor", o.Cursor)
	}

	if o.AccountToUpdateKYC != "" {
		query.Set("account_to_update_kyc", o.AccountToUpdateKYC)
	}

	return query.Encode()
}
