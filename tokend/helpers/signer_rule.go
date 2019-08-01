package helpers

import (
	"fmt"
	"math"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type SignerEntryFunc func(d *schema.ResourceData) (*xdr.SignerRuleResource, error)

var SignerEntries = map[string]SignerEntryFunc{
	"signer":             signerRuleResourceSigner,
	"transaction":        signerRuleResourceTransaction,
	"key_value":          signerRuleResourceKeyValue,
	"reviewable_request": signerRuleResourceReviewableRequest,
}

func SignerRuleEntry(d *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	tpe := d.Get("entry_type").(string)
	createEntry, ok := SignerEntries[tpe]
	if !ok {
		return nil, fmt.Errorf(`entry_type "%s" is not supported`, tpe)
	}
	resource, err := createEntry(d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create signer rule resource")
	}

	return resource, nil
}

func signerRuleResourceSigner(_ *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	return &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeSigner,
		Ext:  &xdr.EmptyExt{},
	}, nil
}

func signerRuleResourceTransaction(_ *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	return &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeTransaction,
		Ext:  &xdr.EmptyExt{},
	}, nil
}

func signerRuleResourceKeyValue(_ *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	return &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeKeyValue,
		KeyValue: &xdr.SignerRuleResourceKeyValue{
			KeyPrefix: "",
		},
	}, nil
}

func signerRuleResourceReviewableRequest(d *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	var resource xdr.SignerRuleResource
	resource.Type = xdr.LedgerEntryTypeReviewableRequest
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

	resource.ReviewableRequest = &xdr.SignerRuleResourceReviewableRequest{
		TasksToAdd:    xdr.Uint64(math.MaxUint64), // TODO
		TasksToRemove: xdr.Uint64(math.MaxUint64), // TODO
		AllTasks:      xdr.Uint64(math.MaxUint64), // TODO
		RequestType:   requestType,
	}

	return &resource, nil
}
