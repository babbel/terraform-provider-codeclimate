package codeclimateclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Repository struct {
	Id             string
	TestReporterId string
	GithubSlug     string
}

// The structure describes just what we need from the response.
//  For the full description look at: https://developer.codeclimate.com/?shell#get-repository
type readRepositoriesResponse struct {
	Data []struct {
		ID         string `json:"id"`
		Attributes struct {
			TestReporterID string `json:"test_reporter_id"`
			GithubSlug     string `json:"github_slug"`
		} `json:"attributes"`
	} `json:"data"`
}

type createRepositoryResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			TestReporterID string `json:"test_reporter_id"`
			GithubSlug     string `json:"github_slug"`
		} `json:"attributes"`
	} `json:"data"`
}

type errorResponse struct {
	Errors []struct {
		Detail string `json:"detail"`
		Title  string `json:"title"`
	} `json:"errors"`
}

func (client *Client) GetRepository(repositorySlug string) (*Repository, error) {
	var repositoryData readRepositoriesResponse

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
		GithubSlug:     repositoryData.Data[0].Attributes.GithubSlug,
	}

	return repository, nil
}

func (client *Client) CreateOrganizationRepository(organizationID string, url string) (*Repository, error) {
	var repositoryData createRepositoryResponse

	payload := fmt.Sprintf(`
		{
			"data": {
				"type": "repos",
				"attributes": {
					"url": "%s"
				}
			}
		}`, url)

	data, err := client.makeRequest(http.MethodPost, fmt.Sprintf("orgs/%s/repos", organizationID), strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &repositoryData)

	if err != nil {
		return nil, err
	}

	if repositoryData.Data.ID == "" {
		return nil, handleErrorResponse(data)
	}

	repository := &Repository{
		Id:             repositoryData.Data.ID,
		TestReporterId: repositoryData.Data.Attributes.TestReporterID,
		GithubSlug:     repositoryData.Data.Attributes.GithubSlug,
	}

	return repository, nil

}

func (client *Client) DeleteOrganizationRepository(repositoryID string) error {
	_, err := client.makeRequest(http.MethodDelete, fmt.Sprintf("repos/%s", repositoryID), nil)

	if err != nil {
		return fmt.Errorf("The repository couldn't be deleted: %s", err)
	}

	return nil
}

func handleErrorResponse(data []byte) error {
	var errorResponse errorResponse
	err := json.Unmarshal(data, &errorResponse)
	if err != nil {
		return err
	}
	return fmt.Errorf("Invalid json response %s", errorResponse.Errors)
}
