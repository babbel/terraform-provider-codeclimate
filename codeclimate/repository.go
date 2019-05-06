package codeclimate

import (
	"encoding/json"
	"fmt"
	"log"
)

// The structure describes just what we need from the response.
//  For the full description look at: https://developer.codeclimate.com/?shell#get-repository
type ReadRepositoryResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			TestReporterID string `json:"test_reporter_id"`
		} `json:"attributes"`
	} `json:"data"`
}

func getRepository(client Client, repoId string) (interface{}, error) {
	var repositoryData ReadRepositoryResponse

	data, err := client.makeRequest(fmt.Sprintf("/repos/%s", repoId))

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
