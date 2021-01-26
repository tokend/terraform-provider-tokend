package asset

import (
	"fmt"

	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/tokend/horizon-connector/internal"
	"gitlab.com/tokend/regources"
)

type Q struct {
	client internal.Client
}

func NewQ(client internal.Client) *Q {
	return &Q{
		client,
	}
}

func (q Q) ByCode(code string) (*regources.Asset, error) {
	endpoint := fmt.Sprintf("/assets/%s", code)
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var asset regources.Asset
	if err := json.Unmarshal(response, &asset); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	return &asset, nil
}

func (q Q) Index() ([]regources.Asset, error) {
	endpoint := "/assets"
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var assets []regources.Asset
	if err := json.Unmarshal(response, &assets); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	return assets, nil
}

func (q Q) Pairs() ([]regources.AssetPair, error) {
	endpoint := "/asset_pairs"
	response, err := q.client.Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if response == nil {
		return nil, nil
	}

	var assetPairs []regources.AssetPair
	if err := json.Unmarshal(response, &assetPairs); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	return assetPairs, nil
}

func (q *Q) AssetPairByCode(base, quote string) (assetPair *regources.AssetPair, err error) {
	assetPairs, err := q.Pairs()
	if err != nil {
		return assetPair, errors.Wrap(err, "failed to get asset pairs")
	}
	for _, ap := range assetPairs {
		if ap.Base == base && ap.Quote == quote {
			return &ap, nil
		}
	}
	return nil, nil
}
