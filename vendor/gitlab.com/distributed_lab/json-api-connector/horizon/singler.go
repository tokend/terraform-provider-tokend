package horizon

import (
	"net/url"
	"path"

	connector "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Singler struct {
	*Getter
	postRetrievers []func(dst interface{}, receivedErr error) error
}

func NewSingler(endpoint string, client *Client, pathParams types.PathParamer) connector.Singler {
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	u.Path = path.Join(u.Path, pathParams.Path())

	return &Singler{
		Getter: &Getter{
			Endpoint: u,
			Client:   client,
		},
		postRetrievers: []func(interface{}, error) error{
			checkErr,
		},
	}
}

func (r *Singler) ValidateResponses(shouldValidate bool) connector.Singler {
	r.shouldValidateResponses = shouldValidate
	return r
}

func checkErr(_ interface{}, err error) error {
	switch err {
	case ErrNotFound, nil:
		return err
	default:
		return errors.Wrap(err, "failed to retrieve resource")
	}
}

func (r *Singler) Get(dst interface{}, query ...types.QueryParamer) error {
	r.Getter.Endpoint.RawQuery = encodeQuery(query...)

	err := r.retrieve(dst)
	if err != nil {
		return err
	}

	for _, retriever := range r.postRetrievers {
		err := retriever(dst, err)
		if err != nil {
			return err
		}
	}
	return nil
}

func encodePath(params types.PathParamer) string {
	return params.Path()
}
