package helpers

import (
	"fmt"
	"math"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type SignerRuleEntryFunc func(d *schema.ResourceData) (*xdr.SignerRuleResource, error)

var SignerRuleEntries = map[string]SignerRuleEntryFunc{
	"signer":             signerRuleResourceSigner,
	"transaction":        signerRuleResourceTransaction,
	"limits":             signerRuleResourceLimits,
	"fee":                signerRuleResourceFee,
	"key_value":          signerRuleResourceKeyValue,
	"sale":               signerRuleResourceSale,
	"asset":              signerRuleResourceAsset,
	"reviewable_request": signerRuleResourceReviewableRequest,
	"stamp":              signerRuleResourceStamp,
	"license":            signerRuleResourceLicense,
}

func SignerRuleEntry(d *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	tpe := d.Get("entry_type").(string)
	createEntry, ok := SignerRuleEntries[tpe]
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
		Details: xdr.ReviewableRequestResource{
			RequestType: requestType,
			Ext:         &xdr.EmptyExt{},
		},
	}

	switch requestType {
	case xdr.ReviewableRequestTypeCreateIssuance:
		resource.ReviewableRequest.Details.CreateIssuance = &xdr.ReviewableRequestResourceCreateIssuance{
			AssetCode: "*",            // TODO
			AssetType: math.MaxUint64, // TODO
		}
	}

	return &resource, nil
}

func signerRuleResourceFee(_ *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	return &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeFee,
		Ext:  &xdr.EmptyExt{},
	}, nil
}

func signerRuleResourceLimits(_ *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	return &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeLimitsV2,
		Ext:  &xdr.EmptyExt{},
	}, nil
}

func signerRuleResourceSale(d *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	var resource xdr.SignerRuleResource
	resource.Type = xdr.LedgerEntryTypeSale
	entry := d.Get("entry").(map[string]interface{})
	saleIdRaw := entry["sale_id"].(string)
	saleTypeRaw := entry["sale_type"].(string)

	saleID, err := WildCardUintFromRaw(saleIdRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast sale_id")
	}

	saleType, err := WildCardUintFromRaw(saleTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to to cast sale_type")
	}

	resource.Sale = &xdr.SignerRuleResourceSale{
		SaleId:   xdr.Uint64(saleID),
		SaleType: xdr.Uint64(saleType),
	}
	return &resource, nil
}

func signerRuleResourceAsset(d *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	var resource xdr.SignerRuleResource
	resource.Type = xdr.LedgerEntryTypeAsset
	entry := d.Get("entry").(map[string]interface{})
	assetCode := entry["asset_code"].(string)
	assetTypeRaw := entry["asset_type"].(string)

	assetType, err := WildCardUintFromRaw(assetTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast asset_type")
	}

	resource.Asset = &xdr.SignerRuleResourceAsset{
		AssetCode: xdr.AssetCode(assetCode),
		AssetType: xdr.Uint64(assetType),
	}
	return &resource, nil
}

func signerRuleResourceStamp(_ *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	return &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeStamp,
		Ext:  &xdr.EmptyExt{},
	}, nil
}

func signerRuleResourceLicense(_ *schema.ResourceData) (*xdr.SignerRuleResource, error) {
	return &xdr.SignerRuleResource{
		Type: xdr.LedgerEntryTypeLicense,
		Ext:  &xdr.EmptyExt{},
	}, nil
}
