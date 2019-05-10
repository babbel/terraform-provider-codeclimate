package codeclimate

import (
	"github.com/babbel/terraform-provider-codeclimate/codeclimateclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRepositoryRead,

		Schema: map[string]*schema.Schema{
			"repository_slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"test_reporter_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRepositoryRead(d *schema.ResourceData, client interface{}) error {
	repositorySlug := d.Get("repository_slug").(string)

	c := client.(*codeclimateclient.Client)
	repository, err := c.GetRepository(repositorySlug)
	if err != nil {
		return err
	}

	d.SetId(repository.Id)
	d.Set("test_reporter_id", repository.TestReporterId)

	return err
}
