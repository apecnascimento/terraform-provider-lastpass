package lastpass

import (
	"context"
	"errors"
	"strconv"

	"github.com/apecnascimento/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceSecret describes our lastpass secret data source
func DataSourceSecret() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceSecretRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"last_modified_gmt": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_touch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"note": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func fillResourceData(d *schema.ResourceData, laspassSecret api.Secret) {
	d.SetId(laspassSecret.ID)
	d.Set("name", laspassSecret.Name)
	d.Set("username", laspassSecret.Username)
	d.Set("password", laspassSecret.Password)
	d.Set("last_modified_gmt", laspassSecret.LastModifiedGmt)
	d.Set("last_touch", laspassSecret.LastTouch)
	d.Set("group", laspassSecret.Group)
	d.Set("url", laspassSecret.URL)
	d.Set("note", laspassSecret.Note)
}

// DataSourceSecretRead reads resource from upstream/lastpass
func DataSourceSecretRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	var diags diag.Diagnostics

	id := d.Get("id").(string)
	name := d.Get("name").(string)

	if id != "" {

		if _, err := strconv.Atoi(id); err != nil {
			err := errors.New("not a valid Lastpass ID")
			return diag.FromErr(err)
		}
		secret, err := client.GetByID(id)
		if err != nil {
			return diag.FromErr(err)
		}
		fillResourceData(d, secret)

	} else if name != "" {
		secret, err := client.GetByName(name)
		if err != nil {
			return diag.FromErr(err)
		}
		fillResourceData(d, secret)
	} else {
		d.SetId("")
	}

	return diags
}
