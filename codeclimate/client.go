package codeclimate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ReadRepositoryResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			TestReporterID string `json:"test_reporter_id"`
		} `json:"attributes"`
	} `json:"data"`
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
