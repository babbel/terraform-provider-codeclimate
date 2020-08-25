package codeclimateclient

import (
	"fmt"
	"net/http"
	"testing"
)

const (
	repositorySlug         = "testorg/testarepo"
	expectedTestReporterID = "0c89092bc2c088d667612ddd1a992ec62f643ded331f40783bcf6b847561234d"
	organizationID         = "testorg"
	repositoryUrl          = "https://github.com/testorg/testarepo"
)

func TestClient_GetRepository(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("github_slug") != repositorySlug {
			t.Fatal(fmt.Errorf("received slug doesn match `%s`", repositorySlug))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getFixture("repositories/repository.json"))
	})

	repository, err := client.GetRepository(repositorySlug)

	if err != nil {
		t.Fatal(err)
	}

	if repository.TestReporterId != expectedTestReporterID {
		t.Errorf("Expected test_reporter_id to be '%s', got: '%s'", expectedTestReporterID, repository.TestReporterId)
	}

	if repository.GithubSlug != repositorySlug {
		t.Errorf("Expected github slug to be '%s', got: '%s'", repositorySlug, repository.GithubSlug)
	}
}

func TestClient_CreateOrganizationRepository(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/orgs/%s/repos", organizationID), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getFixture("repositories/create_repository_response.json"))
	})

	repository, err := client.CreateOrganizationRepository(organizationID, repositoryUrl)

	if err != nil {
		t.Fatal(err)
	}

	if repository.TestReporterId != expectedTestReporterID {
		t.Errorf("Expected test_reporter_id to be '%s', got: '%s'", expectedTestReporterID, repository.TestReporterId)
	}

	if repository.GithubSlug != repositorySlug {
		t.Errorf("Expected github slug to be '%s', got: '%s'", repositorySlug, repository.GithubSlug)
	}
}
