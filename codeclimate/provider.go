package codeclimate

import (
	"github.com/babbel/terraform-provider-codeclimate/codeclimateclient"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

const codeClimateApiHost string = "https://api.codeclimate.com/v1"

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
	client := codeclimateclient.Client{
		ApiKey:  d.Get("api_key").(string),
		BaseUrl: codeClimateApiHost,
	}

	return client, nil
}
