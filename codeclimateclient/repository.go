package codeclimateclient

import (
	"encoding/json"
	"fmt"
)

type Repository struct {
	Id             string
	TestReporterId string
}

// The structure describes just what we need from the response.
//  For the full description look at: https://developer.codeclimate.com/?shell#get-repository
type readRepositoryResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			TestReporterID string `json:"test_reporter_id"`
		} `json:"attributes"`
	} `json:"data"`
}

func (client *Client) GetRepository(repoId string) (*Repository, error) {
	var repositoryData readRepositoryResponse

	data, err := client.makeRequest(fmt.Sprintf("/repos/%s", repoId))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &repositoryData)
	if err != nil {
		return nil, err
	}

	repository := &Repository{
		Id:             repositoryData.Data.ID,
		TestReporterId: repositoryData.Data.Attributes.TestReporterID,
	}

	return repository, nil
}
