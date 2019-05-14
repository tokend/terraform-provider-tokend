package responses

import (
	"net/url"

	"strconv"

	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/horizon-connector/internal/resources"
)

type Wallets struct {
	Data  []resources.Wallet `json:"data"`
	Links Links              `json:"links"`
}

//FIXME when urlval will be a separate project
func (w Wallets) GetPage() (int32, error) {
	url, err := url.Parse(w.Links.Next)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to parse link")
	}

	query := url.Query()
	page := query.Get("page")

	result, err := strconv.ParseInt(page, 10, 32)
	if err != nil {
		return 0, errors.Wrap(err, "Failed cast page to int32")
	}

	return int32(result), nil
}
