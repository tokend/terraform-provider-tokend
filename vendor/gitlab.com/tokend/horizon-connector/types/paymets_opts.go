package types

import "net/url"

type PaymentsOpts struct {
	Since  string
	Cursor string
}

func (o PaymentsOpts) Encode() string {
	query := url.Values{}

	if o.Since != "" {
		query.Set("since", o.Since)
	}

	if o.Cursor != "" {
		query.Set("cursor", o.Cursor)
	}

	return query.Encode()
}
