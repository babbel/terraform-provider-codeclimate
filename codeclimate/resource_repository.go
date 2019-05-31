package codeclimate

import (
	"fmt"

	"github.com/babbel/terraform-provider-codeclimate/codeclimateclient"
	"github.com/hashicorp/terraform/helper/schema"
)

const lessonnineGithubApiHost string = "https://github.com/lessonnine/"

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceRepositoryCreate,
		Read:   resourceRepositoryRead,
		Update: resourceRepositoryUpdate,
		Delete: resourceRepositoryDelete,
		Importer: &schema.ResourceImporter{
			State: resourceRepositoryImporter,
		},

		Schema: map[string]*schema.Schema{
			"organization_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
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

func resourceRepositoryCreate(d *schema.ResourceData, client interface{}) error {
	organizationId := d.Get("organization_id").(string)
	repositorySlug := d.Get("repository_slug").(string)
	repositoryUrl := fmt.Sprintf("%s%s", lessonnineGithubApiHost, repositorySlug)

	c := client.(*codeclimateclient.Client)
	repository, err := c.CreateRepository(organizationId, repositoryUrl)

	d.SetId(repository.Id)
	d.Set("test_reporter_id", repository.TestReporterId)

	return err
}

func resourceRepositoryRead(d *schema.ResourceData, client interface{}) error {
	return nil
}

func resourceRepositoryUpdate(d *schema.ResourceData, client interface{}) error {
	return nil
}

func resourceRepositoryDelete(d *schema.ResourceData, client interface{}) error {
	return nil
}

func resourceRepositoryImporter(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
	return nil, nil
}

// func dataSourceRepositoryRead(d *schema.ResourceData, client interface{}) error {
// 	repositorySlug := d.Get("repository_slug").(string)
//
// 	c := client.(*codeclimateclient.Client)
// 	repository, err := c.GetRepository(repositorySlug)
// 	if err != nil {
// 		return err
// 	}
//
// 	d.SetId(repository.Id)
// 	d.Set("test_reporter_id", repository.TestReporterId)
//
// 	return err
// }
