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
	repoId := d.Get("repository_id").(string)
	repositoryData, err := getRepository(m.(Config).apiKey, repoId)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(repositoryData.(ReadRepositoryResponse).Data.ID)
	// TODO: Check that repositoryData.Data.Attributes.TestReporterID exists
	d.Set("test_reporter_id", repositoryData.(ReadRepositoryResponse).Data.Attributes.TestReporterID)

	return err
}

func getRepository(apiKey string, repoId string) (interface{}, error) {
	var repositoryData ReadRepositoryResponse

	// TODO: Extract into a client
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/repos/%s", codeClimateApiHost, repoId), nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Accept", `W/"application/vnd.api+json"`)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(data, &repositoryData)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return repositoryData, nil
}
