package codeclimate

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

type Config struct {
	apiKey string
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Token for the CodeClimate API.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"codeclimate_repository": dataSourceRepository(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		apiKey: d.Get("api_key").(string),
	}

	return config, nil
}
