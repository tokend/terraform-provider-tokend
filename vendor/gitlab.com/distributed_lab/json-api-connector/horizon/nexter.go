package horizon

import (
	"net/url"
	"path"
	"reflect"
	"strings"

	"github.com/spf13/cast"

	"gitlab.com/distributed_lab/figure"

	"github.com/fatih/structs"
	connector "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/regources"
)

type Nexter struct {
	*Getter
	currentPageLinks *regources.Links
	postRetrievers   []func(arg interface{}) error
}

func NewNexter(endpoint string, client *Client) *Nexter {
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	nexter := &Nexter{
		Getter:           NewGetter(u, client),
		currentPageLinks: &regources.Links{},
	}

	nexter.postRetrievers = []func(arg interface{}) error{
		checkDataContents,
		nexter.tryUpdateLinks,
	}
	return nexter
}

func (r *Nexter) ValidateResponses(shouldValidate bool) connector.Nexter {
	r.shouldValidateResponses = shouldValidate
	return r
}

func (r *Nexter) Next(dst interface{}, params ...types.QueryParamer) error {
	err := r.prepareEndpoint(params...)
	if err != nil {
		return errors.Wrap(err, "failed to prepare endpoint")
	}

	if err = r.retrieve(dst); err != nil {
		return err
	}

	for _, retriever := range r.postRetrievers {
		if err = retriever(dst); err != nil {
			return err
		}
	}

	return nil
}

func (r *Nexter) WithPathParams(pathParams types.PathParamer) connector.Nexter {
	r.Endpoint.Path = path.Join(r.Endpoint.Path, pathParams.Path())
	return r
}

func checkDataContents(container interface{}) error {
	containerVal := reflect.ValueOf(container)

	if containerVal.Kind() == reflect.Ptr {
		containerVal = containerVal.Elem()
	}

	data := containerVal.FieldByName("Data")
	if !data.IsValid() || data.Kind() != reflect.Slice {
		return errors.New("no Data field or not a slice")
	}

	if data.Len() == 0 {
		return ErrDataEmpty
	}

	return nil
}

func (r *Nexter) tryUpdateLinks(from interface{}) error {
	fromVal := reflect.ValueOf(from)

	if fromVal.Kind() == reflect.Ptr {
		fromVal = fromVal.Elem()
	}

	links := fromVal.FieldByName("Links")
	if !links.IsValid() {
		return errors.New("src does not have links field")
	}

	var err error
	if r.currentPageLinks, err = createLinks(links.Interface()); err != nil {
		return errors.Wrap(err, "failed to create new links for nexter")
	}

	if r.currentPageLinks == nil {
		switch from.(type) { // FIXME: temporary hotfix until proper links in account/{id}/signers would arrive
		case *regources.SignerListResponse:
			return nil
		default:
			return ErrLinksNull
		}
	}

	return nil
}

func createLinks(src interface{}) (*regources.Links, error) {
	if reflect.ValueOf(src).IsNil() {
		return nil, nil
	}

	var res struct {
		First string `fig:"First"`
		Last  string `fig:"Last"`
		Next  string `fig:"Next"`
		Prev  string `fig:"Prev"`
		Self  string `fig:"Self"`
	}

	stringOrStrPtrHook := func(value interface{}) (reflect.Value, error) {
		switch value.(type) {
		case *string:
			if reflect.ValueOf(value).IsNil() {
				value = ""
			}
		default:
		}

		res, err := cast.ToStringE(value)
		if err != nil {
			return reflect.Value{}, errors.Wrap(err, "failed to parse string")
		}

		return reflect.ValueOf(res), nil
	}

	if err := figure.Out(&res).With(figure.Hooks{
		"string":  stringOrStrPtrHook,
		"*string": stringOrStrPtrHook,
	}).From(structs.Map(src)).Please(); err != nil {
		return nil, errors.Wrap(err, "failed to figure out links")
	}

	return &regources.Links{
		First: res.First,
		Last:  res.Last,
		Next:  res.Next,
		Prev:  res.Prev,
		Self:  res.Self,
	}, nil
}

func (r *Nexter) prepareEndpoint(params ...types.QueryParamer) error {
	if r.currentPageLinks == nil || r.currentPageLinks.Next == "" {
		r.Endpoint.RawQuery = encodeQuery(params...)
		return nil
	}

	var err error
	r.Endpoint, err = url.Parse(r.currentPageLinks.Next)
	if err != nil { // normally should never happen
		return errors.Wrap(err, "failed to parse links.Next", logan.F{
			"next": r.currentPageLinks.Next,
		})
	}

	return nil
}

func encodeQuery(params ...types.QueryParamer) string {
	encodedParams := make([]string, 0, len(params))
	for _, v := range params {
		encodedParams = append(encodedParams, v.Encode())
	}
	return strings.Join(encodedParams, "&")
}
