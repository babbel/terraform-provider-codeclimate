package codeclimate

import (
	"github.com/babbel/terraform-provider-codeclimate/codeclimate_client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRepositoryRead,

		Schema: map[string]*schema.Schema{
			"repository_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"test_reporter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRepositoryRead(d *schema.ResourceData, client interface{}) error {
	repositoryId := d.Get("repository_id").(string)

	c := client.(codeclimate_client.Client)
	repositoryData, err := c.GetRepository(repositoryId)
	if err != nil {
		return err
	}

	d.SetId(repositoryData.(codeclimate_client.Repository).Id)
	d.Set("test_reporter_id", repositoryData.(codeclimate_client.Repository).TestReporterId)

	return err
}
