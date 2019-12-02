package types

import (
	"net/url"

	"github.com/spf13/cast"
)

type PageParams struct {
	Number *uint32
	Cursor *uint32
	Limit  *uint32
	Order  *string
}

func (p *PageParams) Encode() string {
	u := url.Values{}

	if p.Number != nil {
		u.Add("page[number]", cast.ToString(*p.Number))
	}
	if p.Cursor != nil {
		u.Add("page[cursor]", cast.ToString(*p.Cursor))
	}
	if p.Limit != nil {
		u.Add("page[limit]", cast.ToString(*p.Limit))
	}
	if p.Order != nil {
		u.Add("page[order]", *p.Order)
	}

	return u.Encode()
}
