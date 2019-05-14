package types

import (
	"fmt"
	"net/url"
)

type GetOpts struct {
	Verified *bool
	Page     *int32
}

func (o GetOpts) Encode() string {
	query := url.Values{}
	if o.Verified != nil {
		query.Set("verified", fmt.Sprintf("%t", *o.Verified))
	}

	if o.Page != nil {
		query.Set("page", fmt.Sprintf("%d", *o.Page))
	}

	return fmt.Sprintf("/wallets?%s", query.Encode())
}
