package codeclimateclient

import (
	"encoding/json"
	"fmt"
	"log"
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

type createRepositoryResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			TestReporterID string `json:"test_reporter_id"`
		} `json:"attributes"`
	} `json:"data"`
}

type createRepositoryRequest struct {
	Data Data `json:"data"`
}

type Data struct {
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	URL string `json:"url"`
}

func (client *Client) GetRepository(repositorySlug string) (*Repository, error) {
	var repositoryData readRepositoryResponse

	data, err := client.makeRequest("GET", fmt.Sprintf("repos?github_slug=%s", repositorySlug), nil)

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

func (client *Client) CreateRepository(organizationId string, repositoryUrl string) (*Repository, error) {
	var repositoryData createRepositoryResponse

	createRepositoryBody := &createRepositoryRequest{
		Data: Data{
			Type: "repos",
			Attributes: Attributes{
				URL: repositoryUrl,
			},
		},
	}

	createRepositoryJson, err := json.Marshal(createRepositoryBody)
	log.Println(string(createRepositoryJson))
	if err != nil {
		return nil, err
	}

	data, err := client.makeRequest("POST", fmt.Sprintf("orgs/%s/repos", organizationId), createRepositoryJson)
	// TODO: Add check for the status code here or in the client

	if err != nil {
		return nil, err
	}

	log.Println("asfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasfasf")
	log.Println(string(data))

	err = json.Unmarshal(data, &repositoryData)
	log.Println(string(data))
	if err != nil {
		return nil, err
	}

	repository := &Repository{
		Id:             repositoryData.Data.ID,
		TestReporterId: repositoryData.Data.Attributes.TestReporterID,
	}

	return repository, nil
}
