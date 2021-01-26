package __old

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

// ManageOfferOp should not be constructed directly,
// one should use CreateOffer or DeleteOffer function
// for constructing a ManageOfferOp instance.
type ManageOfferOp struct {
	BaseBalance  string
	QuoteBalance string
	IsBuy        bool
	Amount       int64
	Price        int64
	Fee          int64
	OfferID      uint64
}

func (op ManageOfferOp) Validate() error {
	if op.OfferID == 0 {
		// Create Offer
		return validation.ValidateStruct(&op,
			// TODO validate BaseBalance and QuoteBalance are addresses
			validation.Field(&op.BaseBalance, validation.Required),
			validation.Field(&op.QuoteBalance, validation.Required),

			validation.Field(&op.Amount, validation.Required, validation.Min(int64(1))),
			validation.Field(&op.Price, validation.Required, validation.Min(int64(1))),
			validation.Field(&op.Fee, validation.Min(int64(0))),
		)
	} else {
		// Delete Offer
		return validation.ValidateStruct(&op,
			// Must both be empty
			validation.Field(&op.BaseBalance, validation.RuneLength(0, 0)),
			validation.Field(&op.QuoteBalance, validation.RuneLength(0, 0)),

			validation.Field(&op.Amount, validation.Max(int64(0))),
			validation.Field(&op.Price, validation.Max(int64(0))),
			validation.Field(&op.Fee, validation.Max(int64(0))),
		)
	}
}

func (op ManageOfferOp) XDR() (*xdr.Operation, error) {
	if len(op.BaseBalance) == 0 {
		// Just a random BalanceID, during Delete Offer the value will not be taken into account, but won't work without a valid BalanceID.
		op.BaseBalance = "BBQJLW43UPKJSS67OXGAHJKZ4RM7JMQ6P7FCJEOMBC5GMYT3XQARZI54"
	}
	if len(op.QuoteBalance) == 0 {
		// Just a random BalanceID, during Delete Offer the value will not be taken into account, but won't work without a valid BalanceID.
		op.QuoteBalance = "BBQJLW43UPKJSS67OXGAHJKZ4RM7JMQ6P7FCJEOMBC5GMYT3XQARZI54"
	}

	var baseBalance, quoteBalance xdr.BalanceId
	if err := baseBalance.SetString(op.BaseBalance); err != nil {
		return nil, errors.Wrap(err, "failed to set base BalanceID")
	}
	if err := quoteBalance.SetString(op.QuoteBalance); err != nil {
		return nil, errors.Wrap(err, "failed to set quote BalanceID")
	}

	xdrOp := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageOffer,
			ManageOfferOp: &xdr.ManageOfferOp{
				BaseBalance:  baseBalance,
				QuoteBalance: quoteBalance,
				IsBuy:        op.IsBuy,
				Amount:       xdr.Int64(op.Amount),
				Price:        xdr.Int64(op.Price),
				Fee:          xdr.Int64(op.Fee),
				OfferId:      xdr.Uint64(op.OfferID),
				// OrderBookID 0 means not a sale by general OrderBook
				OrderBookId: xdr.Uint64(0),
			},
		},
	}

	return xdrOp, nil
}

func (op ManageOfferOp) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"base_balance":  op.BaseBalance,
		"quote_balance": op.QuoteBalance,
		"is_buy":        op.IsBuy,
		"amount":        op.Amount,
		"price":         op.Price,
		"fee":           op.Fee,
		"offer_id":      op.OfferID,
	}
}

// CreateOffer is a constructor for ManageOfferOp for creating new Offer.
func CreateOffer(baseBalance, quoteBalance string, isBuy bool, amount int64, price int64, fee int64) *ManageOfferOp {
	return &ManageOfferOp{
		BaseBalance:  baseBalance,
		QuoteBalance: quoteBalance,
		IsBuy:        isBuy,
		Amount:       amount,
		Price:        price,
		Fee:          fee,
	}
}

// DeleteOffer is a constructor for ManageOfferOp
// for deleting the Offer with provided offerID.
func DeleteOffer(offerID uint64) *ManageOfferOp {
	return &ManageOfferOp{
		OfferID: offerID,
	}
}
