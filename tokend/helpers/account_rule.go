package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type AccountEntryFunc func(d *schema.ResourceData) (*xdr.AccountRuleResource, error)

var AccountEntries = map[string]AccountEntryFunc{
	"transaction":                           accountRuleResourceTransaction,
	"signer":                                accountRuleResourceSigner,
	"reviewable_request":                    accountRuleResourceReviewableRequest,
}

func AccountRuleEntry(d *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	tpe := d.Get("entry_type").(string)
	createEntry, ok := AccountEntries[tpe]
	if !ok {
		return nil, fmt.Errorf(`entry_type "%s" is not supported`, tpe)
	}
	resource, err := createEntry(d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create account rule resource")
	}

	return resource, nil
}


func accountRuleResourceTransaction(_ *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	return &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeTransaction,
		Ext:  &xdr.EmptyExt{},
	}, nil
}

func accountRuleResourceSigner(_ *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	return &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeSigner,
		Ext:  &xdr.EmptyExt{},
	}, nil
}

func accountRuleResourceReviewableRequest(d *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	requestTypeRaw := d.Get("entry.request_type").(string)
	var requestType xdr.ReviewableRequestType
	if requestTypeRaw == "*" {
		requestType = xdr.ReviewableRequestTypeAny
	} else {
		for _, guess := range xdr.ReviewableRequestTypeAll {
			if guess.ShortString() == requestTypeRaw {
				requestType = guess
			}
		}
		if requestType == 0 {
			return nil, fmt.Errorf("unknown request type: %s", requestTypeRaw)
		}
	}
	return &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeReviewableRequest,
		ReviewableRequest: &xdr.AccountRuleResourceReviewableRequest{
			RequestType: requestType,
		},
	}, nil
}

