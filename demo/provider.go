package demo

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http://127.0.0.1:8888/",
				DefaultFunc: schema.EnvDefaultFunc("GW_ENDPOINT", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if value == "" {
						errors = append(errors, fmt.Errorf("Endpoint must not be an empty string"))
					}

					return
				},
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"yunjigjl_demo": resourceDemo(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"yunjigjl_account": dataSourceYunjiAccount(),
		},
	}
	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return configureProvider(d)
	}
	return provider
}

type Configuration struct {
	endpoint string
}

func configureProvider(d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	endpoint := d.Get("endpoint").(string)

	_, err := http.Get(endpoint)
	if err != nil {
		return nil, diag.Errorf("Error connect to gateway ")
	}
	return &Configuration{
		endpoint: endpoint,
	}, nil
}
