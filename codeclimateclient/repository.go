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
	Data []struct {
		ID         string `json:"id"`
		Attributes struct {
			TestReporterID string `json:"test_reporter_id"`
		} `json:"attributes"`
	} `json:"data"`
}

func (client *Client) GetRepository(repositorySlug string) (*Repository, error) {
	var repositoryData readRepositoryResponse

	data, err := client.makeRequest(fmt.Sprintf("repos?github_slug=%s", repositorySlug))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &repositoryData)
	if err != nil {
		return nil, err
	}

	// TODO: check size of data

	repository := &Repository{
		Id:             repositoryData.Data[0].ID,
		TestReporterId: repositoryData.Data[0].Attributes.TestReporterID,
	}

	return repository, nil
}
