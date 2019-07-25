package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

type AccountEntryFunc func(d *schema.ResourceData) (*xdr.AccountRuleResource, error)

var AccountEntries = map[string]AccountEntryFunc{
	"offer":                                 accountRuleResourceOffer,
	"transaction":                           accountRuleResourceTransaction,
	"balance":                               accountRuleResourceBalance,
	"signer":                                accountRuleResourceSigner,
	"reviewable_request":                    accountRuleResourceReviewableRequest,
	"sale":                                  accountRuleResourceSale,
	"asset":                                 accountRuleResourceAsset,
	"external_system_account_id_pool_entry": accountRuleResourceExternalSystemAccountIdPool,
	"vote":                                  accountRuleResourceVote,
	"poll":                                  accountRuleResourcePoll,
	"atomic_swap_ask":                       accountRuleResourceAtomicSwapAsk,
	"asset_pair":                            accountRuleResourceAssetPair,
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

func accountRuleResourceAssetPair(_ *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	return &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeAssetPair,
		Ext:  &xdr.EmptyExt{},
	}, nil
}

func accountRuleResourceExternalSystemAccountIdPool(_ *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	return &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeExternalSystemAccountIdPoolEntry,
		Ext:  &xdr.EmptyExt{},
	}, nil
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

func accountRuleResourceBalance(_ *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	return &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeBalance,
		Ext:  &xdr.EmptyExt{},
	}, nil
}

func accountRuleResourceReviewableRequest(_ *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	return &xdr.AccountRuleResource{
		Type: xdr.LedgerEntryTypeReviewableRequest,
		ReviewableRequest: &xdr.AccountRuleResourceReviewableRequest{
			Details: xdr.ReviewableRequestResource{
				RequestType: xdr.ReviewableRequestTypeAny, // TODO
				Ext:         &xdr.EmptyExt{},
			},
		},
	}, nil
}

func accountRuleResourceOffer(d *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	var resource xdr.AccountRuleResource
	isBuyRaw := d.Get("entry.is_buy")
	isBuy, err := cast.ToBoolE(isBuyRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed cast is_buy")
	}
	baseType, err := cast.ToUint64E(d.Get("entry.base_asset_type"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast base_asset_type")
	}
	quoteType, err := cast.ToUint64E(d.Get("entry.quote_asset_type"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast quote_asset_type")
	}
	resource.Type = xdr.LedgerEntryTypeOfferEntry
	resource.Offer = &xdr.AccountRuleResourceOffer{
		BaseAssetType:  xdr.Uint64(baseType),
		QuoteAssetType: xdr.Uint64(quoteType),
		BaseAssetCode:  "*",
		QuoteAssetCode: "*",
		IsBuy:          isBuy,
	}
	return &resource, nil
}

func accountRuleResourceSale(d *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	var resource xdr.AccountRuleResource
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

	resource.Sale = &xdr.AccountRuleResourceSale{
		SaleId:   xdr.Uint64(saleID),
		SaleType: xdr.Uint64(saleType),
	}
	return &resource, nil
}

func accountRuleResourceAsset(d *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	var resource xdr.AccountRuleResource
	resource.Type = xdr.LedgerEntryTypeAsset
	entry := d.Get("entry").(map[string]interface{})
	assetCode := entry["asset_code"].(string)
	assetTypeRaw := entry["asset_type"].(string)
	assetType, err := WildCardUintFromRaw(assetTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast asset_type")
	}

	resource.Asset = &xdr.AccountRuleResourceAsset{
		AssetCode: xdr.AssetCode(assetCode),
		AssetType: xdr.Uint64(assetType),
	}
	return &resource, nil
}

func accountRuleResourceVote(d *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	var resource xdr.AccountRuleResource
	resource.Type = xdr.LedgerEntryTypeVote
	entry := d.Get("entry").(map[string]interface{})
	pollIDRaw := entry["poll_id"].(string)
	permissionTypeRaw := entry["permission_type"].(string)
	pollID, err := WildCardUintFromRaw(pollIDRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast poll_id")
	}
	permissionType, err := WildCardUintFromRaw(permissionTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast permission_type")
	}

	resource.Vote = &xdr.AccountRuleResourceVote{
		PollId:         xdr.Uint64(pollID),
		PermissionType: xdr.Uint32(permissionType),
	}
	return &resource, nil
}

func accountRuleResourcePoll(d *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	var resource xdr.AccountRuleResource
	resource.Type = xdr.LedgerEntryTypePoll
	entry := d.Get("entry").(map[string]interface{})
	pollIDRaw := entry["poll_id"].(string)
	permissionTypeRaw := entry["permission_type"].(string)
	pollID, err := WildCardUintFromRaw(pollIDRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast poll_id")
	}
	permissionType, err := WildCardUintFromRaw(permissionTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast permission_type")
	}

	resource.Poll = &xdr.AccountRuleResourcePoll{
		PollId:         xdr.Uint64(pollID),
		PermissionType: xdr.Uint32(permissionType),
	}
	return &resource, nil
}

func accountRuleResourceAtomicSwapAsk(d *schema.ResourceData) (*xdr.AccountRuleResource, error) {
	var resource xdr.AccountRuleResource
	resource.Type = xdr.LedgerEntryTypeAtomicSwapAsk
	entry := d.Get("entry").(map[string]interface{})
	assetCode := entry["asset_code"].(string)
	assetTypeRaw := entry["asset_type"].(string)
	assetType, err := WildCardUintFromRaw(assetTypeRaw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to cast asset_type")
	}

	resource.AtomicSwapAsk = &xdr.AccountRuleResourceAtomicSwapAsk{
		AssetCode: xdr.AssetCode(assetCode),
		AssetType: xdr.Uint64(assetType),
	}
	return &resource, nil
}
