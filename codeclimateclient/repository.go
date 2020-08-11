package codeclimateclient

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	data, err := client.makeRequest(http.MethodGet, fmt.Sprintf("repos?github_slug=%s", repositorySlug), nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &repositoryData)
	if err != nil {
		return nil, err
	}

	numberOfReposFound := len(repositoryData.Data)

	if numberOfReposFound != 1 {
		return nil, fmt.Errorf(
			"The response for %s returned %v repositories (should have been 1)",
			repositorySlug, numberOfReposFound,
		)
	}

	repository := &Repository{
		Id:             repositoryData.Data[0].ID,
		TestReporterId: repositoryData.Data[0].Attributes.TestReporterID,
	}

	return repository, nil
}
