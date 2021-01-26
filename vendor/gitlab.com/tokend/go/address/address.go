package address

import (
	"net/http"

	"github.com/go-chi/chi"

	. "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/strkey"
)

var (
	ErrAddressInvalid = Errors{"address": errors.New("address is invalid")}
)

type Address string

func FromRequest(r *http.Request, paramName string) (*Address, error) {
	a := chi.URLParam(r, paramName)
	if a == "" {
		return nil, Errors{paramName: errors.New("empty string")}
	}

	addr := Address(a)

	return &addr, addr.Validate()
}

func (a Address) String() string {
	return string(a)
}

func (a Address) Validate() error {
	_, err := strkey.Decode(strkey.VersionByteAccountID, string(a))
	if err != nil {
		return ErrAddressInvalid
	}
	return nil
}

var IsAddress = &isAddress{}

type isAddress struct{}

func (ia *isAddress) Validate(value interface{}) error {
	a, err := cast.ToStringE(value)
	if err != nil {
		return err
	}
	address := Address(a)
	return address.Validate()
}
