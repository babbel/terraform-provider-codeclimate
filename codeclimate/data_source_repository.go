package codeclimate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	codeClimateApiHost = "https://api.codeclimate.com/v1"
)

// It's not the full structure, here is descibed only the part we require.
type ReadRepositoryResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			TestReporterID string `json:"test_reporter_id"`
		} `json:"attributes"`
	} `json:"data"`
}

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

func dataSourceRepositoryRead(d *schema.ResourceData, m interface{}) error {
	var repositoryData ReadRepositoryResponse

	client := &http.Client{}
	repoId := d.Get("repository_id").(string)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/repos/%s", codeClimateApiHost, repoId), nil)

	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Add("Accept", `W/"application/vnd.api+json"`)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", m.(Config).apiKey))

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return err
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(data, &repositoryData)
	log.Printf("The id is: %s", repositoryData.Data.Attributes.TestReporterID)
	log.Printf("Repdata: %s", repositoryData)

	log.Printf("%s", data)

	d.SetId(repositoryData.Data.ID)
	// TODO: Check that repositoryData.Data.Attributes.TestReporterID exists
	d.Set("test_reporter_id", repositoryData.Data.Attributes.TestReporterID)
	log.Printf("The test_reporter_id is: %s", d.Get("test_reporter_id").(string))
	if err != nil {
		return err
	}
	return err
}
