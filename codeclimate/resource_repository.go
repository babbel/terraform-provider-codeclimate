package codeclimate

import (
	"github.com/babbel/terraform-provider-codeclimate/codeclimateclient"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		Read: resourceRepositoryRead,

		Schema: map[string]*schema.Schema{
			"repository_slug": {
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

func resourceRepositoryRead(d *schema.ResourceData, client interface{}) error {
	repositorySlug := d.Get("repository_slug").(string)

	c := client.(*codeclimateclient.Client)
	repository, err := c.GetRepository(repositorySlug)
	if err != nil {
		return err
	}

	d.SetId(repository.Id)
	err = d.Set("test_reporter_id", repository.TestReporterId)

	return err
}
