package connector

import (
	"github.com/fatih/structs"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/json-api-connector/client"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/url"
	"reflect"
)

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
	Self  string `json:"self"`
}

type Encoder interface {
	Encode() string
}

type StubEncoder struct{}

func (s StubEncoder) Encode() string {
	return ""
}

var ErrNotFound = errors.New("resource not found")
var ErrDataEmpty = errors.New("no data on page")

type Streamer struct {
	connector        *Connector
	currentPageLinks *Links

	params Encoder

	endpoint *url.URL

	shouldValidateResponses bool
	postRetrievers          []func(arg interface{}) error
}

func NewStreamer(client client.Client, endpoint string, params Encoder) *Streamer {
	u, _ := url.Parse(endpoint)

	nexter := &Streamer{
		params:           params,
		endpoint:         u,
		connector:        NewConnector(client),
		currentPageLinks: &Links{},
	}

	nexter.postRetrievers = []func(arg interface{}) error{
		nexter.tryUpdateLinks,
	}
	return nexter
}

func (r *Streamer) ValidateResponses(shouldValidate bool) *Streamer {
	r.shouldValidateResponses = shouldValidate
	return r
}

func (r *Streamer) Next(dst interface{}) error {
	err := r.prepareNext()
	if err != nil {
		return errors.Wrap(err, "failed to prepare endpoint")
	}

	if err = r.connector.Get(r.endpoint, dst); err != nil {
		return err
	}

	for _, retriever := range r.postRetrievers {
		if err = retriever(dst); err != nil {
			return err
		}
	}

	return nil
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

func (r *Streamer) tryUpdateLinks(from interface{}) error {
	fromVal := reflect.ValueOf(from)

	if fromVal.Kind() == reflect.Ptr {
		fromVal = fromVal.Elem()
	}

	links := fromVal.FieldByName("Links")
	if !links.IsValid() {
		return errors.New("src does not have links field")
	}

	nxtLinks, err := createLinks(links.Interface())
	if err != nil {
		return errors.Wrap(err, "failed to create new links for nexter")
	}

	if nxtLinks.Next != "" {
		r.currentPageLinks = nxtLinks
	}

	return nil
}

func createLinks(src interface{}) (*Links, error) {
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

	return &Links{
		First: res.First,
		Last:  res.Last,
		Next:  res.Next,
		Prev:  res.Prev,
		Self:  res.Self,
	}, nil
}

func (r *Streamer) prepareNext() error {
	if r.currentPageLinks == nil ||
		(r.currentPageLinks.Next == "" && r.currentPageLinks.Self == "") {
		r.endpoint.RawQuery = r.params.Encode()
		return nil
	}

	var err error
	if r.currentPageLinks.Next != "" {
		r.endpoint, err = url.Parse(r.currentPageLinks.Next)
		if err != nil { // normally should never happen
			return errors.Wrap(err, "failed to parse links.Next", logan.F{
				"next": r.currentPageLinks.Next,
			})
		}
		return nil
	}

	r.endpoint, err = url.Parse(r.currentPageLinks.Self)
	if err != nil {
		return errors.Wrap(err, "failed to parse links.Self", logan.F{
			"self": r.currentPageLinks.Self,
		})
	}

	return nil
}
