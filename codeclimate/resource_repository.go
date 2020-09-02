package codeclimate

import (
	"github.com/babbel/terraform-provider-codeclimate/codeclimateclient"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRepositoryRead,
		Create: resourceRepositoryCreateForOrganization,
		Delete: resourceRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"codeclimate_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"test_reporter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repository_url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceRepositoryRead(d *schema.ResourceData, client interface{}) error {
	repositorySlug := d.Id()

	c := client.(*codeclimateclient.Client)
	repository, err := c.GetRepository(repositorySlug)
	if err != nil {
		return err
	}

	d.SetId(repositorySlug)
	err = d.Set("test_reporter_id", repository.TestReporterId)
	if err != nil {
		return err
	}
	err = d.Set("codeclimate_id", repository.Id)
	if err != nil {
		return err
	}
	err = d.Set("repository_url", repository.RepositoryURL)
	if err != nil {
		return err
	}
	err = d.Set("organization_id", repository.Organization)
	return err
}

func resourceRepositoryCreateForOrganization(d *schema.ResourceData, client interface{}) error {
	repositoryUrl := d.Get("repository_url").(string)
	organizationId := d.Get("organization_id").(string)

	c := client.(*codeclimateclient.Client)

	repository, err := c.CreateOrganizationRepository(organizationId, repositoryUrl)
	if err != nil {
		return err
	}

	d.SetId(repository.GithubSlug)
	err = d.Set("test_reporter_id", repository.TestReporterId)
	if err != nil {
		return err
	}
	err = d.Set("codeclimate_id", repository.Id)

	return err
}

func resourceRepositoryDelete(d *schema.ResourceData, client interface{}) error {
	repositoryID := d.Get("codeclimate_id").(string)

	c := client.(*codeclimateclient.Client)
	err := c.DeleteOrganizationRepository(repositoryID)
	if err != nil {
		return err
	}

	d.SetId("")

	return err
}
