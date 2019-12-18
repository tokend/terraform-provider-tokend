package tokend

import (
	"github.com/tokend/terraform-provider-tokend/tokend/connector"
	"github.com/tokend/terraform-provider-tokend/tokend/data"
	"gitlab.com/distributed_lab/json-api-connector/signed"
	"net/http"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers/validation"
	"gitlab.com/tokend/connectors/submit"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.ValidateSource,
			},
			"signer": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.ValidateSigner,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"tokend_account":                       resourceAccount(),
			"tokend_rule":                          resourceRule(),
			"tokend_role":                          resourceRole(),
			"tokend_key_value":                     resourceKeyValue(),
			"tokend_asset":                         resourceAsset(),
			"tokend_account_signer":                resourceAccountSigner(),
			"tokend_reviewable_request_operations": resourceReviewableRequestOperations(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			endpoint, err := url.Parse(d.Get("endpoint").(string))
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse endpoint")
			}
			signer, err := keypair.ParseSeed(d.Get("signer").(string))
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse signer")
			}
			source, err := keypair.ParseAddress(d.Get("account").(string))

			client := signed.NewClient(http.DefaultClient, endpoint).WithSigner(source, signer)

			submitter := submit.New(client)
			builder, err := submitter.TXBuilder()
			if err != nil {
				return nil, errors.Wrap(err, "failed to create builder")
			}

			return Meta{
				Connector: connector.NewConnector(client),
				Submitter: submitter,
				Builder:   *builder,
				Source:    source,
				Signer:    signer,
			}, nil
		},
	}
}

type Meta struct {
	Connector data.Connector
	Submitter *submit.Transactor
	Signer    keypair.Full
	Source    keypair.Address
	Builder   xdrbuild.Builder
}
