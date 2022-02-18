package lastpass

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/opszak/api"
)

// Provider is the terraform element of the lastpass provider
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"lastpass_secret": DataSourceSecret(),
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Lastpass login e-mail",
				DefaultFunc: schema.EnvDefaultFunc("LASTPASS_USER", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Lastpass login password",
				DefaultFunc: schema.EnvDefaultFunc("LASTPASS_PASSWORD", nil),
			},
			"totp": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Lastpass totp token",
				DefaultFunc: schema.EnvDefaultFunc("LASTPASS_TOTP", nil),
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	client := api.Client{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Totp:     d.Get("totp").(string),
	}
	return &client, diags
}
