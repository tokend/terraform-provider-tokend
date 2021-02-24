package tokend

import (
	"net/url"

	"github.com/tokend/terraform-provider-tokend/tokend/data"

	"github.com/tokend/terraform-provider-tokend/tokend/connector"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
	"github.com/tokend/terraform-provider-tokend/tokend/helpers/validation"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/horizon-connector"
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
			"tokend_account":                    resourceAccount(),
			"tokend_account_rule":               resourceAccountRule(),
			"tokend_account_role":               resourceAccountRole(),
			"tokend_key_value":                  resourceKeyValue(),
			"tokend_asset":                      resourceAsset(),
			"tokend_signer_rule":                resourceSignerRule(),
			"tokend_signer_role":                resourceSignerRole(),
			"tokend_asset_pair":                 resourceAssetPair(),
			"tokend_account_signer":             resourceAccountSigner(),
			"tokend_external_system_pool_entry": resourceExternalSystemPoolEntry(),
			"tokend_data":                       resourceData(),
			"tokend_initial_admin":              resourceInitialAdmin(),
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

			hrz := horizon.NewConnector(endpoint, source).WithSigner(signer)
			builder, err := hrz.TXBuilder()
			if err != nil {
				return nil, errors.Wrap(err, "failed to init builder")
			}

			tasksSvc := horizon.NewConnector(endpoint, source).WithSigner(signer)

			return Meta{
				Horizon:   hrz,
				TasksSvc:  tasksSvc,
				Connector: connector.NewConnector(hrz.Client()),
				Builder:   *builder,
				Source:    source,
				Signer:    signer,
			}, nil
		},
	}
}

type Meta struct {
	Horizon   *horizon.Connector
	TasksSvc  *horizon.Connector
	Connector data.Connector
	Signer    keypair.Full
	Source    keypair.Address
	Builder   xdrbuild.Builder
}
