package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mailjet/mailjet-apiv3-go/v4"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: configSchema(),
			/*DataSourcesMap: map[string]*schema.Resource{
				"scaffolding_data_source": dataSourceScaffolding(),
			},*/
			ResourcesMap: map[string]*schema.Resource{
				"mailjet_subaccount": resourceMailjetSubaccount(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"public_key": {
			Type:     schema.TypeString,
			Description: "Mailjet public API Key (main account)",
			Optional: false,
			Required: true,
		},
		"private_key": {
			Type:     schema.TypeString,
			Description: "Mailjet private API Key (main account)",
			Optional: false,
			Required: true,
		},
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent

		publicKey := d.Get("public_key").(string)
		privateKey := d.Get("private_key").(string)

		mailjetClient := mailjet.NewMailjetClient(publicKey, privateKey)

		return &mailjetClient, nil
	}
}
