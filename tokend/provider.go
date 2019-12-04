package tokend

import (
	"net/http"
	"net/url"

	"github.com/tokend/terraform-provider-tokend/tokend/data"

	"github.com/tokend/terraform-provider-tokend/tokend/connector"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers/validation"
	"gitlab.com/distributed_lab/json-api-connector/horizon"
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
			"tokend_account":        resourceAccount(),
			"tokend_rule":           resourceRule(),
			"tokend_role":           resourceRole(),
			"tokend_key_value":      resourceKeyValue(),
			"tokend_asset":          resourceAsset(),
			"tokend_account_signer": resourceAccountSigner(),
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
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse source")
			}
			client := horizon.NewClient(http.DefaultClient, endpoint).WithSigner(signer)
			hrz := horizon.NewConnector(client, false)
			builder, err := hrz.TXBuilder()
			if err != nil {
				return nil, errors.Wrap(err, "failed to init builder")
			}
			return Meta{
				Horizon:   hrz,
				Connector: connector.NewConnector(client),
				Submitter: connector.NewSubmitter(hrz),
				Builder:   *builder,
				Source:    source,
				Signer:    signer,
			}, nil
		},
	}
}

type Meta struct {
	Horizon   *horizon.Connector
	Connector data.Connector
	Submitter connector.Submitter
	Signer    keypair.Full
	Source    keypair.Address
	Builder   xdrbuild.Builder
}
